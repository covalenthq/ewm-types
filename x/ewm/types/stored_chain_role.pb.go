// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ewm/stored_chain_role.proto

package types

import (
	fmt "fmt"
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

type StoredChainRole struct {
	Index           string            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ChainId         int32             `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ProofSequencer  string            `protobuf:"bytes,3,opt,name=proofSequencer,proto3" json:"proofSequencer,omitempty"`
	ProofOperator   string            `protobuf:"bytes,4,opt,name=proofOperator,proto3" json:"proofOperator,omitempty"`
	ProofValidator  string            `protobuf:"bytes,5,opt,name=proofValidator,proto3" json:"proofValidator,omitempty"`
	ProofType       string            `protobuf:"bytes,6,opt,name=proofType,proto3" json:"proofType,omitempty"`
	Creator         string            `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	ValidatorStatus map[string]bool   `protobuf:"bytes,8,rep,name=validatorStatus,proto3" json:"validatorStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	OpValPair       map[string]string `protobuf:"bytes,9,rep,name=opValPair,proto3" json:"opValPair,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp       string            `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *StoredChainRole) Reset()         { *m = StoredChainRole{} }
func (m *StoredChainRole) String() string { return proto.CompactTextString(m) }
func (*StoredChainRole) ProtoMessage()    {}
func (*StoredChainRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_8172ed37ca41bfed, []int{0}
}
func (m *StoredChainRole) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredChainRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredChainRole.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredChainRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredChainRole.Merge(m, src)
}
func (m *StoredChainRole) XXX_Size() int {
	return m.Size()
}
func (m *StoredChainRole) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredChainRole.DiscardUnknown(m)
}

var xxx_messageInfo_StoredChainRole proto.InternalMessageInfo

func (m *StoredChainRole) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredChainRole) GetChainId() int32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *StoredChainRole) GetProofSequencer() string {
	if m != nil {
		return m.ProofSequencer
	}
	return ""
}

func (m *StoredChainRole) GetProofOperator() string {
	if m != nil {
		return m.ProofOperator
	}
	return ""
}

func (m *StoredChainRole) GetProofValidator() string {
	if m != nil {
		return m.ProofValidator
	}
	return ""
}

func (m *StoredChainRole) GetProofType() string {
	if m != nil {
		return m.ProofType
	}
	return ""
}

func (m *StoredChainRole) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StoredChainRole) GetValidatorStatus() map[string]bool {
	if m != nil {
		return m.ValidatorStatus
	}
	return nil
}

func (m *StoredChainRole) GetOpValPair() map[string]string {
	if m != nil {
		return m.OpValPair
	}
	return nil
}

func (m *StoredChainRole) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredChainRole)(nil), "covalenthq.covenet.covenet.StoredChainRole")
	proto.RegisterMapType((map[string]string)(nil), "covalenthq.covenet.covenet.StoredChainRole.OpValPairEntry")
	proto.RegisterMapType((map[string]bool)(nil), "covalenthq.covenet.covenet.StoredChainRole.ValidatorStatusEntry")
}

func init() { proto.RegisterFile("ewm/stored_chain_role.proto", fileDescriptor_8172ed37ca41bfed) }

var fileDescriptor_8172ed37ca41bfed = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4b, 0xe3, 0x40,
	0x14, 0xc7, 0x9b, 0xfe, 0xce, 0x2c, 0xdb, 0x2e, 0x43, 0x0f, 0x43, 0x59, 0xb2, 0x65, 0x59, 0x96,
	0x5e, 0x4c, 0x41, 0x2f, 0x52, 0x3c, 0x48, 0x45, 0xc4, 0x53, 0x25, 0x95, 0x22, 0x5e, 0xca, 0x34,
	0x79, 0xda, 0x68, 0x9a, 0x49, 0x27, 0x93, 0xd0, 0xfc, 0x17, 0xfe, 0x59, 0x1e, 0x7b, 0x11, 0x3c,
	0x4a, 0xfb, 0x8f, 0x48, 0x26, 0x4d, 0xfa, 0x03, 0x45, 0x3c, 0x65, 0xde, 0x37, 0xef, 0xfb, 0x79,
	0x33, 0xef, 0x3d, 0xf4, 0xc7, 0x64, 0x21, 0xb8, 0x20, 0x3a, 0xbe, 0x60, 0x1c, 0xac, 0x91, 0x39,
	0xa1, 0xb6, 0x3b, 0xe2, 0xcc, 0x01, 0xdd, 0xe3, 0x4c, 0x30, 0xdc, 0x34, 0x59, 0x48, 0x1d, 0x70,
	0xc5, 0x64, 0xa6, 0xaf, 0x73, 0xd3, 0xef, 0xdf, 0x97, 0x22, 0xaa, 0x0f, 0xa4, 0xef, 0x2c, 0xb6,
	0x19, 0xcc, 0x01, 0xdc, 0x40, 0x25, 0xdb, 0xb5, 0x60, 0x4e, 0x94, 0x96, 0xd2, 0x56, 0x8d, 0x24,
	0xc0, 0x04, 0x55, 0x24, 0xf9, 0xd2, 0x22, 0xf9, 0x96, 0xd2, 0x2e, 0x19, 0x69, 0x88, 0xff, 0xa3,
	0x9a, 0xc7, 0x19, 0xbb, 0x1b, 0xc0, 0x2c, 0x00, 0xd7, 0x04, 0x4e, 0x0a, 0xd2, 0xb8, 0xa7, 0xe2,
	0x7f, 0xe8, 0xa7, 0x54, 0xfa, 0x1e, 0x70, 0x2a, 0x18, 0x27, 0x45, 0x99, 0xb6, 0x2b, 0x66, 0xb4,
	0x21, 0x75, 0x6c, 0x4b, 0xa6, 0x95, 0xb6, 0x68, 0x99, 0x8a, 0x7f, 0x23, 0x55, 0x2a, 0xd7, 0x91,
	0x07, 0xa4, 0x2c, 0x53, 0x36, 0x82, 0xbc, 0x2d, 0x07, 0x69, 0xaf, 0xc8, 0x7f, 0x69, 0x88, 0x1f,
	0x50, 0x3d, 0x4c, 0x21, 0x03, 0x41, 0x45, 0xe0, 0x93, 0x6a, 0xab, 0xd0, 0xfe, 0x71, 0x78, 0xaa,
	0x7f, 0xde, 0x27, 0x7d, 0xaf, 0x47, 0xfa, 0x70, 0x17, 0x71, 0xee, 0x0a, 0x1e, 0x19, 0xfb, 0x60,
	0x7c, 0x83, 0x54, 0xe6, 0x0d, 0xa9, 0x73, 0x45, 0x6d, 0x4e, 0x54, 0x59, 0xa5, 0xfb, 0x9d, 0x2a,
	0xfd, 0xd4, 0x9c, 0xf0, 0x37, 0xb0, 0xf8, 0xf5, 0xc2, 0x9e, 0x82, 0x2f, 0xe8, 0xd4, 0x23, 0x28,
	0x79, 0x7d, 0x26, 0x34, 0x7b, 0xa8, 0xf1, 0xd1, 0x05, 0xf1, 0x2f, 0x54, 0x78, 0x84, 0x68, 0x3d,
	0xd7, 0xf8, 0x18, 0xcf, 0x3a, 0xa4, 0x4e, 0x00, 0x72, 0xa6, 0x55, 0x23, 0x09, 0xba, 0xf9, 0x63,
	0xa5, 0x79, 0x82, 0x6a, 0xbb, 0xe5, 0xbf, 0x72, 0xab, 0x5b, 0xee, 0xde, 0xc5, 0xf3, 0x52, 0x53,
	0x16, 0x4b, 0x4d, 0x79, 0x5b, 0x6a, 0xca, 0xd3, 0x4a, 0xcb, 0x2d, 0x56, 0x5a, 0xee, 0x75, 0xa5,
	0xe5, 0x6e, 0x0f, 0xee, 0x6d, 0x31, 0x09, 0xc6, 0xba, 0xc9, 0xa6, 0x9d, 0x4d, 0x2b, 0x3a, 0xe9,
	0x12, 0xcf, 0xb3, 0x93, 0x88, 0x3c, 0xf0, 0xc7, 0x65, 0xb9, 0xc3, 0x47, 0xef, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x33, 0x43, 0x20, 0xe6, 0x02, 0x00, 0x00,
}

func (m *StoredChainRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredChainRole) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredChainRole) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.OpValPair) > 0 {
		for k := range m.OpValPair {
			v := m.OpValPair[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintStoredChainRole(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintStoredChainRole(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintStoredChainRole(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ValidatorStatus) > 0 {
		for k := range m.ValidatorStatus {
			v := m.ValidatorStatus[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintStoredChainRole(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintStoredChainRole(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ProofType) > 0 {
		i -= len(m.ProofType)
		copy(dAtA[i:], m.ProofType)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.ProofType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ProofValidator) > 0 {
		i -= len(m.ProofValidator)
		copy(dAtA[i:], m.ProofValidator)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.ProofValidator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProofOperator) > 0 {
		i -= len(m.ProofOperator)
		copy(dAtA[i:], m.ProofOperator)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.ProofOperator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProofSequencer) > 0 {
		i -= len(m.ProofSequencer)
		copy(dAtA[i:], m.ProofSequencer)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.ProofSequencer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainId != 0 {
		i = encodeVarintStoredChainRole(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredChainRole(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredChainRole(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredChainRole(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredChainRole) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovStoredChainRole(uint64(m.ChainId))
	}
	l = len(m.ProofSequencer)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	l = len(m.ProofOperator)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	l = len(m.ProofValidator)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	l = len(m.ProofType)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	if len(m.ValidatorStatus) > 0 {
		for k, v := range m.ValidatorStatus {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStoredChainRole(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovStoredChainRole(uint64(mapEntrySize))
		}
	}
	if len(m.OpValPair) > 0 {
		for k, v := range m.OpValPair {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStoredChainRole(uint64(len(k))) + 1 + len(v) + sovStoredChainRole(uint64(len(v)))
			n += mapEntrySize + 1 + sovStoredChainRole(uint64(mapEntrySize))
		}
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovStoredChainRole(uint64(l))
	}
	return n
}

func sovStoredChainRole(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredChainRole(x uint64) (n int) {
	return sovStoredChainRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredChainRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredChainRole
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
			return fmt.Errorf("proto: StoredChainRole: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredChainRole: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofSequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofSequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOperator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofOperator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofValidator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatorStatus == nil {
				m.ValidatorStatus = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStoredChainRole
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStoredChainRole
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStoredChainRole
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStoredChainRole(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ValidatorStatus[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpValPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OpValPair == nil {
				m.OpValPair = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStoredChainRole
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStoredChainRole
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStoredChainRole
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStoredChainRole(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthStoredChainRole
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.OpValPair[mapkey] = mapvalue
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredChainRole
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
				return ErrInvalidLengthStoredChainRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredChainRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredChainRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredChainRole
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
func skipStoredChainRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredChainRole
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
					return 0, ErrIntOverflowStoredChainRole
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
					return 0, ErrIntOverflowStoredChainRole
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
				return 0, ErrInvalidLengthStoredChainRole
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredChainRole
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredChainRole
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredChainRole        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredChainRole          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredChainRole = fmt.Errorf("proto: unexpected end of group")
)
