package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	tendermintLightClientTypes "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

const (
	LabelRegisterInterchainQuery = "register_interchain_query"
)

type (
	Keeper struct {
		cdc                   codec.BinaryCodec
		storeKey              storetypes.StoreKey
		memKey                storetypes.StoreKey
		paramstore            paramtypes.Subspace
		ibcKeeper             *ibckeeper.Keeper
		bank                  types.BankKeeper
		contractManagerKeeper types.ContractManagerKeeper
		headerVerifier        types.HeaderVerifier
		transactionVerifier   types.TransactionVerifier
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ibcKeeper *ibckeeper.Keeper,
	bank types.BankKeeper,
	contractManagerKeeper types.ContractManagerKeeper,
	headerVerifier types.HeaderVerifier,
	transactionVerifier types.TransactionVerifier,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                   cdc,
		storeKey:              storeKey,
		memKey:                memKey,
		paramstore:            ps,
		ibcKeeper:             ibcKeeper,
		bank:                  bank,
		contractManagerKeeper: contractManagerKeeper,
		headerVerifier:        headerVerifier,
		transactionVerifier:   transactionVerifier,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetLastRegisteredQueryKey(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.LastRegisteredQueryIDKey)
	if bytes == nil {
		k.Logger(ctx).Debug("Last registered query key don't exists, GetLastRegisteredQueryKey returns 0")
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (k Keeper) SetLastRegisteredQueryKey(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastRegisteredQueryIDKey, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) SaveQuery(ctx sdk.Context, query *types.RegisteredQuery) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.Marshal(query)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal registered query: %v", err)
	}

	store.Set(types.GetRegisteredQueryByIDKey(query.Id), bz)
	k.Logger(ctx).Debug("SaveQuery successful", "query", query)

	return nil
}

func (k Keeper) GetQueryByID(ctx sdk.Context, id uint64) (*types.RegisteredQuery, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredQueryByIDKey(id))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "there is no query with id: %v", id)
	}

	var query types.RegisteredQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return &query, nil
}

// GetAllRegisteredQueries returns all registered queries
func (k Keeper) GetAllRegisteredQueries(ctx sdk.Context) []*types.RegisteredQuery {
	var (
		store   = prefix.NewStore(ctx.KVStore(k.storeKey), types.RegisteredQueryKey)
		queries []*types.RegisteredQuery
	)

	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		query := types.RegisteredQuery{}
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		queries = append(queries, &query)
	}

	return queries
}

// RemoveQuery removes the given query and relative result data from the store. For a KV query it
// deletes the *types.QueryResult stored by the query ID, for a TX query it stores the query ID to
// the list of queries to be removed so the ICQ module can remove the query hashes later.
func (k Keeper) RemoveQuery(ctx sdk.Context, query *types.RegisteredQuery) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRegisteredQueryByIDKey(query.Id))
	queryType := types.InterchainQueryType(query.GetQueryType())
	switch {
	case queryType.IsKV():
		store.Delete(types.GetRegisteredQueryResultByIDKey(query.Id))
	case queryType.IsTX():
		store.Set(types.GetTxQueryToRemoveByIDKey(query.Id), []byte{})
	}
}

// TxQueriesCleanup cleans the module store from obsolete registered TX queries and relative
// stored transaction hashes. Cleans up to params.TxQueryRemovalLimit hashes at a time or all
// the hashes if params.TxQueryRemovalLimit is 0.
func (k Keeper) TxQueriesCleanup(ctx sdk.Context) {
	rmLimit := k.GetParams(ctx).TxQueryRemovalLimit
	limited := rmLimit != 0

	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.TxQueryToRemoveKey)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		removed, complete := k.removeTxQueryHashes(ctx, sdk.BigEndianToUint64(iterator.Key()), rmLimit)

		if complete {
			prefixStore.Delete(iterator.Key())
		}
		if limited {
			rmLimit -= removed
			if rmLimit <= 0 {
				break
			}
		}
	}
}

// SaveKVQueryResult saves the result of the query and updates the query's local and remote heights
// of last result submission. The result's height must be greater than the current remote height of
// the last query result submission, otherwise operation fails.
func (k Keeper) SaveKVQueryResult(ctx sdk.Context, queryID uint64, result *types.QueryResult) error {
	query, err := k.getRegisteredQueryByID(ctx, queryID)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get registered query")
	}

	return k.saveKVQueryResult(ctx, query, result)
}

// SaveTransactionAsProcessed simply stores a key (SubmittedTxKey + bigEndianBytes(queryID) + tx_hash) with
// mock data. This key can be used to check whether a certain transaction was already submitted for a specific
// transaction query.
func (k Keeper) SaveTransactionAsProcessed(ctx sdk.Context, queryID uint64, txHash []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubmittedTransactionIDForQueryKey(queryID, txHash)

	store.Set(key, []byte{})
}

func (k Keeper) CheckTransactionIsAlreadyProcessed(ctx sdk.Context, queryID uint64, txHash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubmittedTransactionIDForQueryKey(queryID, txHash)

	return store.Has(key)
}

// GetQueryResultByID returns a QueryResult for query with id
func (k Keeper) GetQueryResultByID(ctx sdk.Context, id uint64) (*types.QueryResult, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredQueryResultByIDKey(id))
	if bz == nil {
		return nil, types.ErrNoQueryResult
	}

	var query types.QueryResult
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return &query, nil
}

func (k Keeper) UpdateLastLocalHeight(ctx sdk.Context, queryID uint64, newLocalHeight uint64) error {
	query, err := k.getRegisteredQueryByID(ctx, queryID)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get registered query")
	}

	query.LastSubmittedResultLocalHeight = newLocalHeight
	return k.SaveQuery(ctx, query)
}

// UpdateLastRemoteHeight updates the relative query's remote height of the last result submission.
// The height must be greater than the current remote height of the last query result submission,
// otherwise operation fails.
func (k Keeper) UpdateLastRemoteHeight(ctx sdk.Context, queryID uint64, newRemoteHeight uint64) error {
	query, err := k.getRegisteredQueryByID(ctx, queryID)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get registered query")
	}

	if err := k.checkLastRemoteHeight(ctx, *query, newRemoteHeight); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidHeight, err.Error())
	}
	k.updateLastRemoteHeight(ctx, query, newRemoteHeight)
	return k.SaveQuery(ctx, query)
}

// saveKVQueryResult saves the result of the query and updates the query's local and remote heights
// of last result submission. The result's height must be greater than the current remote height of
// the last query result submission, otherwise operation fails.
func (k Keeper) saveKVQueryResult(ctx sdk.Context, query *types.RegisteredQuery, result *types.QueryResult) error {
	store := ctx.KVStore(k.storeKey)
	cleanResult := clearQueryResult(result)
	bz, err := k.cdc.Marshal(&cleanResult)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal result result: %v", err)
	}
	store.Set(types.GetRegisteredQueryResultByIDKey(query.Id), bz)

	k.updateLastRemoteHeight(ctx, query, result.Height)
	k.updateLastLocalHeight(ctx, query, uint64(ctx.BlockHeight()))
	if err := k.SaveQuery(ctx, query); err != nil {
		return sdkerrors.Wrapf(err, "failed to save query %d: %v", query.Id, err)
	}

	k.Logger(ctx).Debug("Successfully saved query result", "result", &result)
	return nil
}

// updateLastLocalHeight updates the query's local height of the last result submission.
func (k Keeper) updateLastLocalHeight(ctx sdk.Context, query *types.RegisteredQuery, height uint64) {
	query.LastSubmittedResultLocalHeight = height
	k.Logger(ctx).Debug("Updated last local height on given query", "queryID", query.Id, "new_local_height", height)
}

// checkLastRemoteHeight checks whether the given height is greater than the query's remote height
// of the last result submission.
func (k Keeper) checkLastRemoteHeight(ctx sdk.Context, query types.RegisteredQuery, height uint64) error {
	if query.LastSubmittedResultRemoteHeight >= height {
		return fmt.Errorf("result's remote height %d is less than or equal to last result's remote height %d", height, query.LastSubmittedResultRemoteHeight)
	}
	return nil
}

// updateLastRemoteHeight updates query's remote height of the last result submission.
func (k Keeper) updateLastRemoteHeight(ctx sdk.Context, query *types.RegisteredQuery, height uint64) {
	query.LastSubmittedResultRemoteHeight = height
	k.Logger(ctx).Debug("Updated last remote height on given query", "queryID", query.Id, "new_remote_height", height)
}

// getRegisteredQueryByID loads a query by the given ID from the store.
func (k Keeper) getRegisteredQueryByID(ctx sdk.Context, queryID uint64) (*types.RegisteredQuery, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRegisteredQueryByIDKey(queryID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "query with ID %d not found", queryID)
	}

	var query types.RegisteredQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}
	return &query, nil
}

// We don't need to store proofs or transactions, so we just remove unnecessary fields
func clearQueryResult(result *types.QueryResult) types.QueryResult {
	storageValues := make([]*types.StorageValue, 0, len(result.KvResults))
	for _, v := range result.KvResults {
		storageValues = append(storageValues, &types.StorageValue{
			StoragePrefix: v.StoragePrefix,
			Key:           v.Key,
			Value:         v.Value,
			Proof:         nil,
		})
	}

	cleanResult := types.QueryResult{
		KvResults: storageValues,
		Block:     nil,
		Height:    result.Height,
	}

	return cleanResult
}

func (k Keeper) checkRegisteredQueryExists(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetRegisteredQueryByIDKey(id))
}

func (k Keeper) GetClientState(ctx sdk.Context, clientID string) (*tendermintLightClientTypes.ClientState, error) {
	clientStateResponse, ok := k.ibcKeeper.ClientKeeper.GetClientState(ctx, clientID)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrInvalidClientID, "could not find a ClientState with client id: %s", clientID)
	}

	clientState, ok := clientStateResponse.(*tendermintLightClientTypes.ClientState)
	if !ok {
		return nil, sdkerrors.Wrapf(ibcclienttypes.ErrInvalidClientType, "cannot cast ClientState interface into ClientState type")
	}

	return clientState, nil
}

func (k *Keeper) CollectDeposit(ctx sdk.Context, queryInfo types.RegisteredQuery) error {
	owner, err := queryInfo.GetOwnerAddress()
	if err != nil {
		panic(err.Error())
	}

	err = k.bank.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, queryInfo.Deposit)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) MustPayOutDeposit(ctx sdk.Context, deposit sdk.Coins, sender sdk.AccAddress) {
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, deposit)
	if err != nil {
		panic(err.Error())
	}
}

// GetTxQueriesToRemove retrieves the list of TX queries registered to be removed.
func (k Keeper) GetTxQueriesToRemove(ctx sdk.Context) []uint64 {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.TxQueryToRemoveKey)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()
	ids := make([]uint64, 0, 100)
	for ; iterator.Valid(); iterator.Next() {
		ids = append(ids, sdk.BigEndianToUint64(iterator.Key()))
	}
	if len(ids) == 0 {
		return nil
	}
	return ids
}

// removeTxQueryHashes removes up to limit tx hashes related to the query with the given ID from the
// Keeper's store and returns the number of removed tx hashes and if all the query's hashes have
// been removed during the call. If limit is 0, it removes all the hashes for the given query.
func (k Keeper) removeTxQueryHashes(ctx sdk.Context, queryID, limit uint64) (removed uint64, complete bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSubmittedTransactionIDForQueryKeyPrefix(queryID))
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		txHashKey := iterator.Key()
		prefixStore.Delete(txHashKey)
		removed++
		if limit != 0 && removed >= limit {
			return removed, !iterator.Valid()
		}
	}
	return removed, true
}
