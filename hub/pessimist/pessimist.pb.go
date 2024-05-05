// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hub/pessimist/pessimist.proto

package hub_pessimist

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
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

type ClientState struct {
	DepedentClientId string       `protobuf:"bytes,1,opt,name=depedentClientId,proto3" json:"depedentClientId,omitempty"`
	LatestHeight     types.Height `protobuf:"bytes,7,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cce33860bba2a04d, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "hub.pessimist.ClientState")
}

func init() { proto.RegisterFile("hub/pessimist/pessimist.proto", fileDescriptor_cce33860bba2a04d) }

var fileDescriptor_cce33860bba2a04d = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x28, 0x4d, 0xd2,
	0x2f, 0x48, 0x2d, 0x2e, 0xce, 0xcc, 0xcd, 0x2c, 0x2e, 0x41, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x78, 0x33, 0x4a, 0x93, 0xf4, 0xe0, 0x82, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60,
	0x19, 0x7d, 0x10, 0x0b, 0xa2, 0x48, 0x4a, 0x3e, 0x33, 0x29, 0x59, 0x3f, 0x39, 0xbf, 0x28, 0x55,
	0x3f, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x44, 0xbf, 0xcc, 0x10, 0xca, 0x82, 0x28, 0x50, 0x6a, 0x63,
	0xe4, 0xe2, 0x76, 0x06, 0x0b, 0x04, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x69, 0x71, 0x09, 0xa4, 0xa4,
	0x16, 0xa4, 0xa6, 0xa4, 0xe6, 0x95, 0x40, 0x84, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x30, 0xc4, 0x85, 0x5c, 0xb9, 0x78, 0x73, 0x12, 0x4b, 0x52, 0x8b, 0x4b, 0xe2, 0x33, 0x52,
	0x33, 0xd3, 0x33, 0x4a, 0x24, 0xd8, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x32, 0x93, 0x92,
	0xf5, 0x40, 0x96, 0xea, 0x41, 0xad, 0x2a, 0x33, 0xd4, 0xf3, 0x00, 0xab, 0x70, 0x62, 0x39, 0x71,
	0x4f, 0x9e, 0x21, 0x88, 0x07, 0xa2, 0x0d, 0x22, 0x66, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0x49,
	0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x2e, 0x35, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x5d, 0xe3, 0xa6, 0x10, 0x01, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPessimist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.DepedentClientId) > 0 {
		i -= len(m.DepedentClientId)
		copy(dAtA[i:], m.DepedentClientId)
		i = encodeVarintPessimist(dAtA, i, uint64(len(m.DepedentClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPessimist(dAtA []byte, offset int, v uint64) int {
	offset -= sovPessimist(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepedentClientId)
	if l > 0 {
		n += 1 + l + sovPessimist(uint64(l))
	}
	l = m.LatestHeight.Size()
	n += 1 + l + sovPessimist(uint64(l))
	return n
}

func sovPessimist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPessimist(x uint64) (n int) {
	return sovPessimist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPessimist
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepedentClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPessimist
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
				return ErrInvalidLengthPessimist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPessimist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepedentClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPessimist
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
				return ErrInvalidLengthPessimist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPessimist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPessimist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPessimist
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
func skipPessimist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPessimist
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
					return 0, ErrIntOverflowPessimist
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
					return 0, ErrIntOverflowPessimist
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
				return 0, ErrInvalidLengthPessimist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPessimist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPessimist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPessimist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPessimist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPessimist = fmt.Errorf("proto: unexpected end of group")
)