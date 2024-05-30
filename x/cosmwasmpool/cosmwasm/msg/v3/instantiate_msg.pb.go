// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/v3/instantiate_msg.proto

package v3

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

// InstantiateMsg defines the message to instantiate a new alloyed asset
// transmuter v3 pool.
type InstantiateMsg struct {
	PoolAssetConfigs                []AssetConfig `protobuf:"bytes,1,rep,name=pool_asset_configs,json=poolAssetConfigs,proto3" json:"pool_asset_configs"`
	AlloyedAssetSubdenom            string        `protobuf:"bytes,2,opt,name=alloyed_asset_subdenom,json=alloyedAssetSubdenom,proto3" json:"alloyed_asset_subdenom,omitempty"`
	AlloyedAssetNormalizationFactor string        `protobuf:"bytes,3,opt,name=alloyed_asset_normalization_factor,json=alloyedAssetNormalizationFactor,proto3" json:"alloyed_asset_normalization_factor,omitempty"`
	Admin                           string        `protobuf:"bytes,4,opt,name=admin,proto3" json:"admin,omitempty"`
	Moderator                       string        `protobuf:"bytes,5,opt,name=moderator,proto3" json:"moderator,omitempty"`
}

func (m *InstantiateMsg) Reset()         { *m = InstantiateMsg{} }
func (m *InstantiateMsg) String() string { return proto.CompactTextString(m) }
func (*InstantiateMsg) ProtoMessage()    {}
func (*InstantiateMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e95cde1b816aa470, []int{0}
}
func (m *InstantiateMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstantiateMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstantiateMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstantiateMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstantiateMsg.Merge(m, src)
}
func (m *InstantiateMsg) XXX_Size() int {
	return m.Size()
}
func (m *InstantiateMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_InstantiateMsg.DiscardUnknown(m)
}

var xxx_messageInfo_InstantiateMsg proto.InternalMessageInfo

func (m *InstantiateMsg) GetPoolAssetConfigs() []AssetConfig {
	if m != nil {
		return m.PoolAssetConfigs
	}
	return nil
}

func (m *InstantiateMsg) GetAlloyedAssetSubdenom() string {
	if m != nil {
		return m.AlloyedAssetSubdenom
	}
	return ""
}

func (m *InstantiateMsg) GetAlloyedAssetNormalizationFactor() string {
	if m != nil {
		return m.AlloyedAssetNormalizationFactor
	}
	return ""
}

func (m *InstantiateMsg) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *InstantiateMsg) GetModerator() string {
	if m != nil {
		return m.Moderator
	}
	return ""
}

func init() {
	proto.RegisterType((*InstantiateMsg)(nil), "osmosis.cosmwasmpool.v1beta1.model.v3.InstantiateMsg")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/v3/instantiate_msg.proto", fileDescriptor_e95cde1b816aa470)
}

var fileDescriptor_e95cde1b816aa470 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x93, 0xbe, 0xfc, 0xa1, 0xf9, 0x83, 0x48, 0x28, 0x12, 0x8a, 0xa4, 0xa5, 0x20, 0xf4,
	0xe2, 0x2e, 0x6d, 0xf4, 0xa2, 0x27, 0x2b, 0x08, 0x22, 0x7a, 0xa8, 0x37, 0x11, 0xc2, 0xe6, 0xa5,
	0x71, 0x21, 0x9b, 0xa9, 0x99, 0x6d, 0xb4, 0x7e, 0x0a, 0xbf, 0x89, 0x5f, 0xa3, 0xc7, 0x1e, 0x3d,
	0x89, 0xb4, 0x5f, 0x44, 0xb2, 0x4d, 0x69, 0x7a, 0xeb, 0x2d, 0x33, 0xf3, 0xfc, 0x9e, 0x67, 0xb2,
	0x63, 0x5c, 0x02, 0x0a, 0x40, 0x8e, 0xd4, 0x07, 0x14, 0x6f, 0x0c, 0xc5, 0x04, 0x20, 0xa6, 0x59,
	0xdf, 0x0b, 0x25, 0xeb, 0x53, 0x01, 0x41, 0x18, 0xd3, 0xcc, 0xa1, 0x3c, 0x41, 0xc9, 0x12, 0xc9,
	0x99, 0x0c, 0x5d, 0x81, 0x11, 0x99, 0xa4, 0x20, 0xc1, 0x3c, 0x29, 0x60, 0x52, 0x86, 0x49, 0x01,
	0x13, 0x05, 0x93, 0xcc, 0x69, 0x35, 0x23, 0x88, 0x40, 0x11, 0x34, 0xff, 0x5a, 0xc3, 0xad, 0x8b,
	0xfd, 0x92, 0xf3, 0xae, 0xfb, 0x3a, 0x0d, 0xd3, 0xd9, 0x36, 0xb8, 0xfb, 0x55, 0x31, 0x0e, 0x6e,
	0xb7, 0x2b, 0xdd, 0x63, 0x64, 0x8e, 0x0d, 0x53, 0x49, 0x19, 0x62, 0x28, 0x5d, 0x1f, 0x92, 0x31,
	0x8f, 0xd0, 0xd2, 0x3b, 0xd5, 0xde, 0xff, 0xc1, 0x80, 0xec, 0xb5, 0x28, 0xb9, 0xca, 0xd9, 0x6b,
	0x85, 0x0e, 0x6b, 0xf3, 0x9f, 0xb6, 0x36, 0x3a, 0xcc, 0x85, 0xa5, 0x36, 0x9a, 0x67, 0xc6, 0x11,
	0x8b, 0x63, 0x98, 0x85, 0x41, 0x11, 0x85, 0x53, 0x2f, 0x08, 0x13, 0x10, 0x56, 0xa5, 0xa3, 0xf7,
	0x1a, 0xa3, 0x66, 0x31, 0x55, 0xd0, 0x63, 0x31, 0x33, 0xef, 0x8c, 0xee, 0x2e, 0x95, 0x40, 0x2a,
	0x58, 0xcc, 0x3f, 0x98, 0xe4, 0x90, 0xb8, 0x63, 0xe6, 0x4b, 0x48, 0xad, 0xaa, 0x72, 0x68, 0x97,
	0x1d, 0x1e, 0xca, 0xba, 0x1b, 0x25, 0x33, 0x9b, 0x46, 0x9d, 0x05, 0x82, 0x27, 0x56, 0x4d, 0xe9,
	0xd7, 0x85, 0x79, 0x6c, 0x34, 0xf2, 0x1f, 0x49, 0x59, 0xee, 0x54, 0x57, 0x93, 0x6d, 0x63, 0xf8,
	0x3c, 0x5f, 0xda, 0xfa, 0x62, 0x69, 0xeb, 0xbf, 0x4b, 0x5b, 0xff, 0x5c, 0xd9, 0xda, 0x62, 0x65,
	0x6b, 0xdf, 0x2b, 0x5b, 0x7b, 0x1a, 0x46, 0x5c, 0xbe, 0x4c, 0x3d, 0xe2, 0x83, 0xa0, 0xc5, 0x33,
	0x9d, 0xc6, 0xcc, 0xc3, 0x4d, 0x41, 0xb3, 0xc1, 0x39, 0x7d, 0xdf, 0xbd, 0xd2, 0xa6, 0xa0, 0x02,
	0x23, 0x9a, 0x39, 0xde, 0x3f, 0x75, 0x16, 0xe7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x59, 0xfb,
	0x99, 0x4e, 0x02, 0x00, 0x00,
}

func (m *InstantiateMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstantiateMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstantiateMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Moderator) > 0 {
		i -= len(m.Moderator)
		copy(dAtA[i:], m.Moderator)
		i = encodeVarintInstantiateMsg(dAtA, i, uint64(len(m.Moderator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintInstantiateMsg(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AlloyedAssetNormalizationFactor) > 0 {
		i -= len(m.AlloyedAssetNormalizationFactor)
		copy(dAtA[i:], m.AlloyedAssetNormalizationFactor)
		i = encodeVarintInstantiateMsg(dAtA, i, uint64(len(m.AlloyedAssetNormalizationFactor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AlloyedAssetSubdenom) > 0 {
		i -= len(m.AlloyedAssetSubdenom)
		copy(dAtA[i:], m.AlloyedAssetSubdenom)
		i = encodeVarintInstantiateMsg(dAtA, i, uint64(len(m.AlloyedAssetSubdenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolAssetConfigs) > 0 {
		for iNdEx := len(m.PoolAssetConfigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssetConfigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInstantiateMsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintInstantiateMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovInstantiateMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InstantiateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolAssetConfigs) > 0 {
		for _, e := range m.PoolAssetConfigs {
			l = e.Size()
			n += 1 + l + sovInstantiateMsg(uint64(l))
		}
	}
	l = len(m.AlloyedAssetSubdenom)
	if l > 0 {
		n += 1 + l + sovInstantiateMsg(uint64(l))
	}
	l = len(m.AlloyedAssetNormalizationFactor)
	if l > 0 {
		n += 1 + l + sovInstantiateMsg(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovInstantiateMsg(uint64(l))
	}
	l = len(m.Moderator)
	if l > 0 {
		n += 1 + l + sovInstantiateMsg(uint64(l))
	}
	return n
}

func sovInstantiateMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInstantiateMsg(x uint64) (n int) {
	return sovInstantiateMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InstantiateMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInstantiateMsg
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
			return fmt.Errorf("proto: InstantiateMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstantiateMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssetConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssetConfigs = append(m.PoolAssetConfigs, AssetConfig{})
			if err := m.PoolAssetConfigs[len(m.PoolAssetConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlloyedAssetSubdenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlloyedAssetSubdenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlloyedAssetNormalizationFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlloyedAssetNormalizationFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moderator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moderator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInstantiateMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInstantiateMsg
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
func skipInstantiateMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInstantiateMsg
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
					return 0, ErrIntOverflowInstantiateMsg
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
					return 0, ErrIntOverflowInstantiateMsg
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
				return 0, ErrInvalidLengthInstantiateMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInstantiateMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInstantiateMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInstantiateMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInstantiateMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInstantiateMsg = fmt.Errorf("proto: unexpected end of group")
)
