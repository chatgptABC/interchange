// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchange/dex/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId            string          `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	SellOrderBookList []SellOrderBook `protobuf:"bytes,3,rep,name=sellOrderBookList,proto3" json:"sellOrderBookList"`
	BuyOrderBookList  []BuyOrderBook  `protobuf:"bytes,4,rep,name=buyOrderBookList,proto3" json:"buyOrderBookList"`
	DenomTraceList    []DenomTrace    `protobuf:"bytes,5,rep,name=denomTraceList,proto3" json:"denomTraceList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_369e5344ab79db30, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetSellOrderBookList() []SellOrderBook {
	if m != nil {
		return m.SellOrderBookList
	}
	return nil
}

func (m *GenesisState) GetBuyOrderBookList() []BuyOrderBook {
	if m != nil {
		return m.BuyOrderBookList
	}
	return nil
}

func (m *GenesisState) GetDenomTraceList() []DenomTrace {
	if m != nil {
		return m.DenomTraceList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "interchange.dex.GenesisState")
}

func init() { proto.RegisterFile("interchange/dex/genesis.proto", fileDescriptor_369e5344ab79db30) }

var fileDescriptor_369e5344ab79db30 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xdb, 0x6d, 0x4e, 0xcc, 0xc4, 0x97, 0x20, 0x6c, 0x4c, 0x17, 0xa7, 0x28, 0xec, 0xd4,
	0xe1, 0xc4, 0x2f, 0x50, 0x04, 0x19, 0x08, 0x93, 0xcd, 0x93, 0x97, 0xd2, 0x2e, 0x0f, 0xb5, 0x6c,
	0x6b, 0x4a, 0x9a, 0xc1, 0xf6, 0x2d, 0xfc, 0x58, 0x3b, 0xc9, 0x8e, 0x9e, 0x44, 0xd6, 0x2f, 0x22,
	0x4d, 0x83, 0xd4, 0xe4, 0xd6, 0xf2, 0xff, 0x3d, 0xbf, 0x3c, 0x2f, 0xa8, 0x13, 0xc5, 0x02, 0xf8,
	0xf4, 0xdd, 0x8f, 0x43, 0xe8, 0x53, 0x58, 0xf5, 0x43, 0x88, 0x21, 0x8d, 0x52, 0x27, 0xe1, 0x4c,
	0x30, 0x7c, 0x5c, 0x8a, 0x1d, 0x0a, 0xab, 0xf6, 0x59, 0xc8, 0x42, 0x26, 0xb3, 0x7e, 0xfe, 0x55,
	0x60, 0xed, 0x0b, 0xdd, 0x92, 0xf8, 0xdc, 0x5f, 0x28, 0x49, 0xfb, 0x56, 0x4f, 0x53, 0x98, 0xcf,
	0x3d, 0xc6, 0x29, 0x70, 0x2f, 0x60, 0x6c, 0xa6, 0xb0, 0x1b, 0x1d, 0x0b, 0x96, 0x6b, 0x93, 0xba,
	0xd2, 0x29, 0x0a, 0x31, 0x5b, 0x78, 0x82, 0xfb, 0x53, 0x28, 0x90, 0xeb, 0xcf, 0x0a, 0x3a, 0x7c,
	0x2a, 0xc6, 0x98, 0x08, 0x5f, 0x00, 0x7e, 0x40, 0xf5, 0xa2, 0xa1, 0x96, 0xdd, 0xb5, 0x7b, 0x8d,
	0x41, 0xd3, 0xd1, 0xc6, 0x72, 0x5e, 0x64, 0xec, 0xd6, 0x36, 0xdf, 0x97, 0xd6, 0x58, 0xc1, 0xb8,
	0x89, 0xf6, 0x13, 0xc6, 0x85, 0x17, 0xd1, 0x56, 0xa5, 0x6b, 0xf7, 0x0e, 0xc6, 0xf5, 0xfc, 0x77,
	0x48, 0xf1, 0x18, 0x9d, 0xe6, 0x23, 0x8c, 0xf2, 0xde, 0x5c, 0xc6, 0x66, 0xcf, 0x51, 0x2a, 0x5a,
	0xd5, 0x6e, 0xb5, 0xd7, 0x18, 0x10, 0x43, 0x3d, 0x29, 0x93, 0xea, 0x05, 0xb3, 0x1c, 0x8f, 0xd0,
	0x49, 0xb0, 0x5c, 0xff, 0x57, 0xd6, 0xa4, 0xb2, 0x63, 0x28, 0xdd, 0x12, 0xa8, 0x8c, 0x46, 0x31,
	0x1e, 0xa2, 0x23, 0xb9, 0x9a, 0xd7, 0x7c, 0x33, 0x52, 0xb7, 0x27, 0x75, 0xe7, 0x86, 0xee, 0xf1,
	0x0f, 0x53, 0x32, 0xad, 0xd0, 0xbd, 0xdb, 0xec, 0x88, 0xbd, 0xdd, 0x11, 0xfb, 0x67, 0x47, 0xec,
	0x8f, 0x8c, 0x58, 0xdb, 0x8c, 0x58, 0x5f, 0x19, 0xb1, 0xde, 0x9a, 0xe5, 0x6b, 0xac, 0xe4, 0x3d,
	0xc4, 0x3a, 0x81, 0x34, 0xa8, 0xcb, 0x53, 0xdc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x67, 0xc3,
	0x33, 0x75, 0x60, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomTraceList) > 0 {
		for iNdEx := len(m.DenomTraceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomTraceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BuyOrderBookList) > 0 {
		for iNdEx := len(m.BuyOrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BuyOrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SellOrderBookList) > 0 {
		for iNdEx := len(m.SellOrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.SellOrderBookList) > 0 {
		for _, e := range m.SellOrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BuyOrderBookList) > 0 {
		for _, e := range m.BuyOrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenomTraceList) > 0 {
		for _, e := range m.DenomTraceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellOrderBookList = append(m.SellOrderBookList, SellOrderBook{})
			if err := m.SellOrderBookList[len(m.SellOrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyOrderBookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyOrderBookList = append(m.BuyOrderBookList, BuyOrderBook{})
			if err := m.BuyOrderBookList[len(m.BuyOrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTraceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomTraceList = append(m.DenomTraceList, DenomTrace{})
			if err := m.DenomTraceList[len(m.DenomTraceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)