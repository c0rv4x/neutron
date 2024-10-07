# syntax=docker/dockerfile:1

ARG GO_VERSION="1.22"
ARG RUNNER_IMAGE="gcr.io/distroless/static:nonroot"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine3.20 as builder

ARG GIT_VERSION
ARG GIT_COMMIT
ARG BUILD_TAGS
ARG ENABLED_PROPOSALS

ENV GOTOOLCHAIN go1.22.6

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers \
    wget

# Download go dependencies
WORKDIR /neutron
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Cosmwasm - Download and verify libwasmvm version
RUN WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm/v2 | cut -d ' ' -f 2) && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$(uname -m).a \
      -O /lib/libwasmvm_muslc.$(uname -m).a && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.$(uname -m).a | grep $(grep libwasmvm_muslc.$(uname -m) /tmp/checksums.txt | cut -d ' ' -f 1)

# Copy the remaining files
COPY . .

# Build neutrond binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go build \
      -mod=readonly \
      -tags ${BUILD_TAGS} \
      -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=neutron \
              -X github.com/cosmos/cosmos-sdk/version.AppName=neutrond \
              -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
              -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
              -X github.com/cosmos/cosmos-sdk/version.BuildTags='${BUILD_TAGS}' \
              -X github.com/neutron-org/neutron/app.EnableSpecificProposals=${ENABLED_PROPOSALS} \
              -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
      -trimpath \
      -o /neutron/build/neutrond \
      /neutron/cmd/neutrond

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ${RUNNER_IMAGE}

COPY --from=builder /neutron/build/neutrond /bin/neutrond

ENV HOME /neutron
WORKDIR $HOME

USER nonroot

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["neutrond"]
