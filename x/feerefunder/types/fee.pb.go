// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/feerefunder/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Fee defines the ICS29 receive, acknowledgement and timeout fees
type Fee struct {
	// the packet receive fee
	RecvFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=recv_fee,json=recvFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"recv_fee" yaml:"recv_fee"`
	// the packet acknowledgement fee
	AckFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"ack_fee" yaml:"ack_fee"`
	// the packet timeout fee
	TimeoutFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"timeout_fee" yaml:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_08c389f5f82f1076, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetRecvFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.RecvFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

type PacketID struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId    string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *PacketID) Reset()         { *m = PacketID{} }
func (m *PacketID) String() string { return proto.CompactTextString(m) }
func (*PacketID) ProtoMessage()    {}
func (*PacketID) Descriptor() ([]byte, []int) {
	return fileDescriptor_08c389f5f82f1076, []int{1}
}
func (m *PacketID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketID.Merge(m, src)
}
func (m *PacketID) XXX_Size() int {
	return m.Size()
}
func (m *PacketID) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketID.DiscardUnknown(m)
}

var xxx_messageInfo_PacketID proto.InternalMessageInfo

func (m *PacketID) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *PacketID) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *PacketID) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Fee)(nil), "neutron.feerefunder.Fee")
	proto.RegisterType((*PacketID)(nil), "neutron.feerefunder.PacketID")
}

func init() { proto.RegisterFile("neutron/feerefunder/fee.proto", fileDescriptor_08c389f5f82f1076) }

var fileDescriptor_08c389f5f82f1076 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0x6e, 0xe2, 0x40,
	0x14, 0x85, 0x3d, 0xb0, 0xe2, 0x67, 0x90, 0x76, 0x25, 0xef, 0x4a, 0xcb, 0x22, 0x61, 0x90, 0x2b,
	0x37, 0x78, 0xc4, 0xfe, 0x34, 0x5b, 0x42, 0x84, 0x44, 0x95, 0x88, 0x32, 0x45, 0x90, 0x3d, 0xbe,
	0x18, 0xcb, 0x78, 0x86, 0xd8, 0x63, 0x2b, 0xb4, 0x79, 0x82, 0x3c, 0x47, 0x9e, 0x84, 0x92, 0x32,
	0x15, 0x89, 0xe0, 0x0d, 0xd2, 0x47, 0x8a, 0xc6, 0x1e, 0x22, 0x52, 0x21, 0xaa, 0x39, 0x33, 0x47,
	0xe7, 0x7c, 0x77, 0xa4, 0x8b, 0xdb, 0x0c, 0x52, 0x11, 0x73, 0x46, 0x66, 0x00, 0x31, 0xcc, 0x52,
	0xe6, 0x41, 0x2c, 0xb5, 0xbd, 0x8c, 0xb9, 0xe0, 0xfa, 0x77, 0x65, 0xdb, 0x47, 0x76, 0xcb, 0xa0,
	0x3c, 0x89, 0x78, 0x42, 0x5c, 0x27, 0x01, 0x92, 0xf5, 0x5d, 0x10, 0x4e, 0x9f, 0x50, 0x1e, 0xb0,
	0x22, 0xd4, 0xfa, 0xe1, 0x73, 0x9f, 0xe7, 0x92, 0x48, 0x55, 0xbc, 0x9a, 0x6f, 0x25, 0x5c, 0x1e,
	0x01, 0xe8, 0x2b, 0x5c, 0x8b, 0x81, 0x66, 0xd3, 0x19, 0x40, 0x13, 0x75, 0xcb, 0x56, 0xe3, 0xf7,
	0x2f, 0xbb, 0x28, 0xb4, 0x65, 0xa1, 0xad, 0x0a, 0xed, 0x21, 0x0f, 0xd8, 0x60, 0xb8, 0xde, 0x76,
	0xb4, 0xd7, 0x6d, 0xe7, 0xdb, 0xca, 0x89, 0x16, 0xff, 0xcd, 0x43, 0xd0, 0x7c, 0x7c, 0xee, 0x58,
	0x7e, 0x20, 0xe6, 0xa9, 0x6b, 0x53, 0x1e, 0x11, 0x35, 0x50, 0x71, 0xf4, 0x12, 0x2f, 0x24, 0x62,
	0xb5, 0x84, 0x24, 0xef, 0x48, 0x26, 0x55, 0x19, 0x93, 0xe8, 0x0c, 0x57, 0x1d, 0x1a, 0xe6, 0xe4,
	0xd2, 0x29, 0xf2, 0x40, 0x91, 0xbf, 0x16, 0x64, 0x95, 0x3b, 0x0f, 0x5c, 0x71, 0x68, 0x28, 0xb9,
	0xf7, 0x08, 0x37, 0x44, 0x10, 0x01, 0x4f, 0x45, 0x0e, 0x2f, 0x9f, 0x82, 0x8f, 0x14, 0x5c, 0x2f,
	0xe0, 0x47, 0xd9, 0xf3, 0x06, 0xc0, 0x2a, 0x39, 0x02, 0x30, 0x6f, 0x70, 0xed, 0xca, 0xa1, 0x21,
	0x88, 0xf1, 0x85, 0xde, 0xc6, 0x98, 0xce, 0x1d, 0xc6, 0x60, 0x31, 0x0d, 0xbc, 0x26, 0xea, 0x22,
	0xab, 0x3e, 0xa9, 0xab, 0x97, 0xb1, 0xa7, 0xff, 0xc4, 0xd5, 0x25, 0x8f, 0x85, 0xf4, 0x4a, 0xb9,
	0x57, 0x91, 0xd7, 0xb1, 0xa7, 0xb7, 0x70, 0x2d, 0x81, 0xdb, 0x14, 0x18, 0x95, 0x9f, 0x40, 0xd6,
	0x97, 0xc9, 0xc7, 0x7d, 0x70, 0xb9, 0xde, 0x19, 0x68, 0xb3, 0x33, 0xd0, 0xcb, 0xce, 0x40, 0x0f,
	0x7b, 0x43, 0xdb, 0xec, 0x0d, 0xed, 0x69, 0x6f, 0x68, 0xd7, 0xff, 0x8e, 0xe6, 0x55, 0xfb, 0xd4,
	0xe3, 0xb1, 0x7f, 0xd0, 0x24, 0xfb, 0x4b, 0xee, 0x3e, 0xed, 0x5f, 0xfe, 0x05, 0xb7, 0x92, 0xef,
	0xcd, 0x9f, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x9c, 0xd3, 0x7d, 0xa3, 0x02, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RecvFee) > 0 {
		for iNdEx := len(m.RecvFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecvFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PacketID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintFee(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintFee(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintFee(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RecvFee) > 0 {
		for _, e := range m.RecvFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovFee(uint64(m.Sequence))
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecvFee = append(m.RecvFee, types.Coin{})
			if err := m.RecvFee[len(m.RecvFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *PacketID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: PacketID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
