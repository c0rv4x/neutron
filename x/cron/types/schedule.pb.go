// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/cron/schedule.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ExecutionStage defines when messages will be executed in the block
type ExecutionStage int32

const (
	// Execution at the end of the block
	ExecutionStage_EXECUTION_STAGE_END_BLOCKER ExecutionStage = 0
	// Execution at the beginning of the block
	ExecutionStage_EXECUTION_STAGE_BEGIN_BLOCKER ExecutionStage = 1
)

var ExecutionStage_name = map[int32]string{
	0: "EXECUTION_STAGE_END_BLOCKER",
	1: "EXECUTION_STAGE_BEGIN_BLOCKER",
}

var ExecutionStage_value = map[string]int32{
	"EXECUTION_STAGE_END_BLOCKER":   0,
	"EXECUTION_STAGE_BEGIN_BLOCKER": 1,
}

func (x ExecutionStage) String() string {
	return proto.EnumName(ExecutionStage_name, int32(x))
}

func (ExecutionStage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{0}
}

// Schedule defines the schedule for execution
type Schedule struct {
	// Name of schedule
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Period in blocks
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// Msgs that will be executed every certain number of blocks, specified in the `period` field
	Msgs []MsgExecuteContract `protobuf:"bytes,3,rep,name=msgs,proto3" json:"msgs"`
	// Last execution's block height
	LastExecuteHeight uint64 `protobuf:"varint,4,opt,name=last_execute_height,json=lastExecuteHeight,proto3" json:"last_execute_height,omitempty"`
	// Execution stages when messages will be executed
	ExecutionStages []ExecutionStage `protobuf:"varint,5,rep,packed,name=execution_stages,json=executionStages,proto3,enum=neutron.cron.ExecutionStage" json:"execution_stages,omitempty"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{0}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return m.Size()
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schedule) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Schedule) GetMsgs() []MsgExecuteContract {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func (m *Schedule) GetLastExecuteHeight() uint64 {
	if m != nil {
		return m.LastExecuteHeight
	}
	return 0
}

func (m *Schedule) GetExecutionStages() []ExecutionStage {
	if m != nil {
		return m.ExecutionStages
	}
	return nil
}

// MsgExecuteContract defines the contract and the message to pass
type MsgExecuteContract struct {
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Msg is json encoded message to be passed to the contract
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *MsgExecuteContract) Reset()         { *m = MsgExecuteContract{} }
func (m *MsgExecuteContract) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContract) ProtoMessage()    {}
func (*MsgExecuteContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{1}
}
func (m *MsgExecuteContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContract.Merge(m, src)
}
func (m *MsgExecuteContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContract proto.InternalMessageInfo

func (m *MsgExecuteContract) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *MsgExecuteContract) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// ScheduleCount defines the number of current schedules
type ScheduleCount struct {
	// Count is the number of current schedules
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *ScheduleCount) Reset()         { *m = ScheduleCount{} }
func (m *ScheduleCount) String() string { return proto.CompactTextString(m) }
func (*ScheduleCount) ProtoMessage()    {}
func (*ScheduleCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{2}
}
func (m *ScheduleCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduleCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduleCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleCount.Merge(m, src)
}
func (m *ScheduleCount) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleCount.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleCount proto.InternalMessageInfo

func (m *ScheduleCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("neutron.cron.ExecutionStage", ExecutionStage_name, ExecutionStage_value)
	proto.RegisterType((*Schedule)(nil), "neutron.cron.Schedule")
	proto.RegisterType((*MsgExecuteContract)(nil), "neutron.cron.MsgExecuteContract")
	proto.RegisterType((*ScheduleCount)(nil), "neutron.cron.ScheduleCount")
}

func init() { proto.RegisterFile("neutron/cron/schedule.proto", fileDescriptor_49ace1b59de613ef) }

var fileDescriptor_49ace1b59de613ef = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xcd, 0x98, 0x74, 0xd9, 0x1d, 0x75, 0xad, 0xe3, 0x22, 0x61, 0xab, 0x69, 0x2c, 0x08, 0x41,
	0x30, 0x81, 0xea, 0xc9, 0x9b, 0x89, 0x43, 0x5b, 0xb4, 0x2d, 0xa4, 0x15, 0xc4, 0x4b, 0x48, 0xd3,
	0x61, 0x12, 0x68, 0x32, 0x25, 0x33, 0x91, 0xfa, 0x2f, 0xfc, 0x59, 0x3d, 0xf6, 0xe8, 0x49, 0xa4,
	0xfd, 0x05, 0xfe, 0x03, 0xc9, 0x64, 0x5a, 0xac, 0x7b, 0x19, 0xde, 0x37, 0xef, 0x3d, 0xde, 0xc7,
	0xe3, 0x83, 0x9d, 0x82, 0x54, 0xa2, 0x64, 0x85, 0x97, 0xd4, 0x0f, 0x4f, 0x52, 0xb2, 0xac, 0x56,
	0xc4, 0x5d, 0x97, 0x4c, 0x30, 0xf4, 0x40, 0x91, 0x6e, 0x4d, 0xde, 0xde, 0x50, 0x46, 0x99, 0x24,
	0xbc, 0x1a, 0x35, 0x9a, 0xde, 0x1f, 0x00, 0x2f, 0x67, 0xca, 0x86, 0x10, 0x34, 0x8a, 0x38, 0x27,
	0x26, 0xb0, 0x81, 0x73, 0x15, 0x4a, 0x8c, 0x9e, 0xc2, 0x8b, 0x35, 0x29, 0x33, 0xb6, 0x34, 0xef,
	0xd9, 0xc0, 0x31, 0x42, 0x35, 0xa1, 0x77, 0xd0, 0xc8, 0x39, 0xe5, 0xa6, 0x6e, 0xeb, 0xce, 0xfd,
	0xbe, 0xed, 0xfe, 0x9b, 0xe5, 0x8e, 0x39, 0xc5, 0x1b, 0x92, 0x54, 0x82, 0x04, 0xac, 0x10, 0x65,
	0x9c, 0x08, 0xdf, 0xd8, 0xfe, 0xea, 0x6a, 0xa1, 0xf4, 0x20, 0x17, 0x3e, 0x59, 0xc5, 0x5c, 0x44,
	0xa4, 0xd1, 0x44, 0x29, 0xc9, 0x68, 0x2a, 0x4c, 0x43, 0x06, 0x3c, 0xae, 0x29, 0xe5, 0x1e, 0x4a,
	0x02, 0x8d, 0x61, 0xbb, 0x91, 0x66, 0xac, 0x88, 0xb8, 0x88, 0x29, 0xe1, 0x66, 0xcb, 0xd6, 0x9d,
	0xeb, 0xfe, 0xb3, 0xf3, 0x5c, 0x7c, 0x54, 0xcd, 0x6a, 0x91, 0xca, 0x7c, 0x44, 0xce, 0x7e, 0x79,
	0xcf, 0x87, 0xe8, 0xee, 0x82, 0xe8, 0x16, 0x5e, 0x26, 0x0a, 0xab, 0x02, 0x4e, 0x33, 0x6a, 0x43,
	0x3d, 0xe7, 0x54, 0x36, 0x70, 0x15, 0xd6, 0xb0, 0xf7, 0x12, 0x3e, 0x3c, 0xd6, 0x16, 0xb0, 0xaa,
	0x10, 0xe8, 0x06, 0xb6, 0x92, 0x1a, 0x48, 0x6f, 0x2b, 0x6c, 0x86, 0x57, 0x73, 0x78, 0x7d, 0xbe,
	0x13, 0xea, 0xc2, 0x0e, 0xfe, 0x82, 0x83, 0xcf, 0xf3, 0xd1, 0x74, 0x12, 0xcd, 0xe6, 0xef, 0x07,
	0x38, 0xc2, 0x93, 0x0f, 0x91, 0xff, 0x69, 0x1a, 0x7c, 0xc4, 0x61, 0x5b, 0x43, 0x2f, 0xe0, 0xf3,
	0xff, 0x05, 0x3e, 0x1e, 0x8c, 0x26, 0x27, 0x09, 0xf0, 0x87, 0xdb, 0xbd, 0x05, 0x76, 0x7b, 0x0b,
	0xfc, 0xde, 0x5b, 0xe0, 0xc7, 0xc1, 0xd2, 0x76, 0x07, 0x4b, 0xfb, 0x79, 0xb0, 0xb4, 0xaf, 0x2e,
	0xcd, 0x44, 0x5a, 0x2d, 0xdc, 0x84, 0xe5, 0x9e, 0x6a, 0xe6, 0x35, 0x2b, 0xe9, 0x11, 0x7b, 0xdf,
	0xde, 0x7a, 0x9b, 0xe6, 0x56, 0xc4, 0xf7, 0x35, 0xe1, 0x8b, 0x0b, 0x79, 0x05, 0x6f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x0a, 0x42, 0x3f, 0x48, 0x02, 0x00, 0x00,
}

func (m *Schedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExecutionStages) > 0 {
		dAtA2 := make([]byte, len(m.ExecutionStages)*10)
		var j1 int
		for _, num := range m.ExecutionStages {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintSchedule(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastExecuteHeight != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.LastExecuteHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSchedule(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Period != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduleCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchedule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	if m.Period != 0 {
		n += 1 + sovSchedule(uint64(m.Period))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovSchedule(uint64(l))
		}
	}
	if m.LastExecuteHeight != 0 {
		n += 1 + sovSchedule(uint64(m.LastExecuteHeight))
	}
	if len(m.ExecutionStages) > 0 {
		l = 0
		for _, e := range m.ExecutionStages {
			l += sovSchedule(uint64(e))
		}
		n += 1 + sovSchedule(uint64(l)) + l
	}
	return n
}

func (m *MsgExecuteContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *ScheduleCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovSchedule(uint64(m.Count))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, MsgExecuteContract{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastExecuteHeight", wireType)
			}
			m.LastExecuteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastExecuteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType == 0 {
				var v ExecutionStage
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ExecutionStage(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExecutionStages = append(m.ExecutionStages, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSchedule
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSchedule
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.ExecutionStages) == 0 {
					m.ExecutionStages = make([]ExecutionStage, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ExecutionStage
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSchedule
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ExecutionStage(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExecutionStages = append(m.ExecutionStages, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionStages", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgExecuteContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgExecuteContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchedule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchedule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchedule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchedule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchedule = fmt.Errorf("proto: unexpected end of group")
)
