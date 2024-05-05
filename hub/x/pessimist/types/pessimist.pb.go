// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hub/pessimist/pessimist.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClientState struct {
	DependentClientId string `protobuf:"bytes,1,opt,name=dependent_client_id,json=dependentClientId,proto3" json:"dependent_client_id,omitempty"`
	//ibc.core.client.v1.Height latest_height = 2 [(gogoproto.nullable) = false];
	LatestHeight Height `protobuf:"bytes,2,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
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

// TODO: Add app hash here too!
type CommitteeProposal struct {
	Height      Height        `protobuf:"bytes,1,opt,name=height,proto3" json:"height"`
	Commitments []*Commitment `protobuf:"bytes,2,rep,name=commitments,proto3" json:"commitments,omitempty"`
}

func (m *CommitteeProposal) Reset()         { *m = CommitteeProposal{} }
func (m *CommitteeProposal) String() string { return proto.CompactTextString(m) }
func (*CommitteeProposal) ProtoMessage()    {}
func (*CommitteeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_cce33860bba2a04d, []int{1}
}
func (m *CommitteeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitteeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitteeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitteeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitteeProposal.Merge(m, src)
}
func (m *CommitteeProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommitteeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitteeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommitteeProposal proto.InternalMessageInfo

func (m *CommitteeProposal) GetHeight() Height {
	if m != nil {
		return m.Height
	}
	return Height{}
}

func (m *CommitteeProposal) GetCommitments() []*Commitment {
	if m != nil {
		return m.Commitments
	}
	return nil
}

type Commitment struct {
	ClientId      string    `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Height        Height    `protobuf:"bytes,2,opt,name=height,proto3" json:"height"`
	Timestamp     time.Time `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ValidatorAddr string    `protobuf:"bytes,4,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
	Signature     []byte    `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_cce33860bba2a04d, []int{2}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return m.Size()
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

func (m *Commitment) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Commitment) GetHeight() Height {
	if m != nil {
		return m.Height
	}
	return Height{}
}

func (m *Commitment) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Commitment) GetValidatorAddr() string {
	if m != nil {
		return m.ValidatorAddr
	}
	return ""
}

func (m *Commitment) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Copied from IBC because it was hard to import and then ignite and pulsar to work together (and removing pulsar broke ignite)
type Height struct {
	// the revision that the client is currently on
	RevisionNumber uint64 `protobuf:"varint,1,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	// the height within the given revision
	RevisionHeight uint64 `protobuf:"varint,2,opt,name=revision_height,json=revisionHeight,proto3" json:"revision_height,omitempty"`
}

func (m *Height) Reset()         { *m = Height{} }
func (m *Height) String() string { return proto.CompactTextString(m) }
func (*Height) ProtoMessage()    {}
func (*Height) Descriptor() ([]byte, []int) {
	return fileDescriptor_cce33860bba2a04d, []int{3}
}
func (m *Height) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Height) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Height.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Height) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Height.Merge(m, src)
}
func (m *Height) XXX_Size() int {
	return m.Size()
}
func (m *Height) XXX_DiscardUnknown() {
	xxx_messageInfo_Height.DiscardUnknown(m)
}

var xxx_messageInfo_Height proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "hub.pessimist.ClientState")
	proto.RegisterType((*CommitteeProposal)(nil), "hub.pessimist.CommitteeProposal")
	proto.RegisterType((*Commitment)(nil), "hub.pessimist.Commitment")
	proto.RegisterType((*Height)(nil), "hub.pessimist.Height")
}

func init() { proto.RegisterFile("hub/pessimist/pessimist.proto", fileDescriptor_cce33860bba2a04d) }

var fileDescriptor_cce33860bba2a04d = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbf, 0x6e, 0xd3, 0x40,
	0x18, 0xf7, 0xa5, 0x21, 0x6a, 0x2e, 0x4d, 0x51, 0x0d, 0x95, 0xdc, 0x88, 0x3a, 0x21, 0x0b, 0x59,
	0xb0, 0xa5, 0x76, 0x83, 0x05, 0x9c, 0xa5, 0x2c, 0x08, 0xb9, 0x88, 0x81, 0xc5, 0xd8, 0xf1, 0x87,
	0x73, 0x92, 0xed, 0xb3, 0xee, 0x3e, 0x47, 0xf0, 0x00, 0x48, 0xb0, 0xf5, 0x11, 0x78, 0x08, 0x1e,
	0xa2, 0x63, 0xc5, 0xc4, 0x04, 0x28, 0x79, 0x11, 0xe4, 0x3b, 0xc7, 0x71, 0x59, 0x60, 0xbb, 0xfb,
	0xfd, 0xf9, 0xf4, 0xbb, 0xef, 0xa7, 0xa3, 0xa7, 0xcb, 0x32, 0x72, 0x0b, 0x90, 0x92, 0x65, 0x4c,
	0xe2, 0xee, 0xe4, 0x14, 0x82, 0x23, 0x37, 0x87, 0xcb, 0x32, 0x72, 0x1a, 0x70, 0x74, 0x3f, 0xe1,
	0x09, 0x57, 0x8c, 0x5b, 0x9d, 0xb4, 0x68, 0x34, 0x4e, 0x38, 0x4f, 0x52, 0x70, 0xd5, 0x2d, 0x2a,
	0xdf, 0xbb, 0xc8, 0x32, 0x90, 0x18, 0x66, 0x45, 0x2d, 0x38, 0x59, 0x70, 0x99, 0x71, 0x19, 0x68,
	0xa7, 0xbe, 0x68, 0x6a, 0xfa, 0x89, 0xd0, 0xc1, 0x3c, 0x65, 0x90, 0xe3, 0x25, 0x86, 0x08, 0xa6,
	0x43, 0xef, 0xc5, 0x50, 0x40, 0x1e, 0x43, 0x8e, 0xc1, 0x42, 0x11, 0x01, 0x8b, 0x2d, 0x32, 0x21,
	0xb3, 0xbe, 0x7f, 0xd4, 0x50, 0xda, 0xf2, 0x22, 0x36, 0x9f, 0xd1, 0x61, 0x1a, 0x22, 0x48, 0x0c,
	0x96, 0xc0, 0x92, 0x25, 0x5a, 0x9d, 0x09, 0x99, 0x0d, 0xce, 0x8e, 0x9d, 0x5b, 0xc1, 0x9d, 0x0b,
	0x45, 0x7a, 0xdd, 0xeb, 0x9f, 0x63, 0xc3, 0x3f, 0xd0, 0x0e, 0x8d, 0x3d, 0xe9, 0x7e, 0xfe, 0x3a,
	0x36, 0xaa, 0x1c, 0x47, 0x73, 0x9e, 0x65, 0x0c, 0x11, 0xe0, 0x95, 0xe0, 0x05, 0x97, 0x61, 0x6a,
	0x9e, 0xd3, 0x5e, 0x3d, 0x96, 0xfc, 0x7b, 0x6c, 0x2d, 0x35, 0x9f, 0xd2, 0xc1, 0x42, 0x4d, 0xca,
	0x20, 0x47, 0x69, 0x75, 0x26, 0x7b, 0xb3, 0xc1, 0xd9, 0xc9, 0x5f, 0xce, 0x79, 0xa3, 0xf0, 0xdb,
	0xea, 0xe9, 0x97, 0x0e, 0xa5, 0x3b, 0xce, 0x1c, 0xd1, 0xfd, 0x45, 0xfd, 0xd4, 0x7a, 0x07, 0xcd,
	0xbd, 0x15, 0xae, 0xf3, 0xff, 0xe1, 0x3c, 0xda, 0x6f, 0xda, 0xb1, 0xf6, 0x94, 0x6f, 0xe4, 0xe8,
	0xfe, 0x9c, 0x6d, 0x7f, 0xce, 0xeb, 0xad, 0xc2, 0xdb, 0xaf, 0xcc, 0x57, 0xbf, 0xc6, 0xc4, 0xdf,
	0xd9, 0xcc, 0x0b, 0x7a, 0xb8, 0x0a, 0x53, 0x16, 0x87, 0xc8, 0x45, 0x10, 0xc6, 0xb1, 0xb0, 0xba,
	0x55, 0x34, 0xef, 0xe1, 0xf7, 0x6f, 0x8f, 0x4f, 0xeb, 0x76, 0xdf, 0x6c, 0x05, 0xcf, 0xe3, 0x58,
	0x80, 0x94, 0x97, 0x28, 0x58, 0x9e, 0xf8, 0xc3, 0x55, 0x1b, 0x37, 0x1f, 0xd0, 0xbe, 0x64, 0x49,
	0x1e, 0x62, 0x29, 0xc0, 0xba, 0x33, 0x21, 0xb3, 0x03, 0x7f, 0x07, 0x4c, 0xdf, 0xd1, 0x9e, 0x7e,
	0x83, 0xf9, 0x88, 0xde, 0x15, 0xb0, 0x62, 0x92, 0xf1, 0x3c, 0xc8, 0xcb, 0x2c, 0x02, 0xa1, 0xb6,
	0xd1, 0xf5, 0x0f, 0xb7, 0xf0, 0x4b, 0x85, 0xde, 0x12, 0xb6, 0x96, 0xd3, 0x12, 0xb6, 0x5b, 0xf7,
	0xdc, 0xeb, 0xb5, 0x4d, 0x6e, 0xd6, 0x36, 0xf9, 0xbd, 0xb6, 0xc9, 0xd5, 0xc6, 0x36, 0x6e, 0x36,
	0xb6, 0xf1, 0x63, 0x63, 0x1b, 0x6f, 0x8f, 0xab, 0x7f, 0xf1, 0xa1, 0xf5, 0x33, 0xf0, 0x63, 0x01,
	0x32, 0xea, 0xa9, 0x1d, 0x9d, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x13, 0xdf, 0xb5, 0x69, 0x37,
	0x03, 0x00, 0x00,
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
	dAtA[i] = 0x12
	if len(m.DependentClientId) > 0 {
		i -= len(m.DependentClientId)
		copy(dAtA[i:], m.DependentClientId)
		i = encodeVarintPessimist(dAtA, i, uint64(len(m.DependentClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommitteeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitteeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitteeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commitments) > 0 {
		for iNdEx := len(m.Commitments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commitments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPessimist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPessimist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Commitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintPessimist(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintPessimist(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0x22
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintPessimist(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPessimist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintPessimist(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Height) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Height) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Height) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevisionHeight != 0 {
		i = encodeVarintPessimist(dAtA, i, uint64(m.RevisionHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.RevisionNumber != 0 {
		i = encodeVarintPessimist(dAtA, i, uint64(m.RevisionNumber))
		i--
		dAtA[i] = 0x8
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
	l = len(m.DependentClientId)
	if l > 0 {
		n += 1 + l + sovPessimist(uint64(l))
	}
	l = m.LatestHeight.Size()
	n += 1 + l + sovPessimist(uint64(l))
	return n
}

func (m *CommitteeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Height.Size()
	n += 1 + l + sovPessimist(uint64(l))
	if len(m.Commitments) > 0 {
		for _, e := range m.Commitments {
			l = e.Size()
			n += 1 + l + sovPessimist(uint64(l))
		}
	}
	return n
}

func (m *Commitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovPessimist(uint64(l))
	}
	l = m.Height.Size()
	n += 1 + l + sovPessimist(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovPessimist(uint64(l))
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovPessimist(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovPessimist(uint64(l))
	}
	return n
}

func (m *Height) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RevisionNumber != 0 {
		n += 1 + sovPessimist(uint64(m.RevisionNumber))
	}
	if m.RevisionHeight != 0 {
		n += 1 + sovPessimist(uint64(m.RevisionHeight))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field DependentClientId", wireType)
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
			m.DependentClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
func (m *CommitteeProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CommitteeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitteeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
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
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitments", wireType)
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
			m.Commitments = append(m.Commitments, &Commitment{})
			if err := m.Commitments[len(m.Commitments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Commitment) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Commitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
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
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
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
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPessimist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPessimist
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPessimist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
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
func (m *Height) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Height: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Height: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevisionNumber", wireType)
			}
			m.RevisionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPessimist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevisionNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevisionHeight", wireType)
			}
			m.RevisionHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPessimist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevisionHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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