// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/resource/v1/resource.proto

package v1

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

type Resource struct {
	Header *ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9a40e43d3305340, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetHeader() *ResourceHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Resource) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResourceHeader struct {
	CollectionId      string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Id                string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ResourceType      string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	MediaType         string `protobuf:"bytes,5,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	Created           string `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Checksum          []byte `protobuf:"bytes,7,opt,name=checksum,proto3" json:"checksum,omitempty"`
	PreviousVersionId string `protobuf:"bytes,8,opt,name=previous_version_id,json=previousVersionId,proto3" json:"previous_version_id,omitempty"`
	NextVersionId     string `protobuf:"bytes,9,opt,name=next_version_id,json=nextVersionId,proto3" json:"next_version_id,omitempty"`
}

func (m *ResourceHeader) Reset()         { *m = ResourceHeader{} }
func (m *ResourceHeader) String() string { return proto.CompactTextString(m) }
func (*ResourceHeader) ProtoMessage()    {}
func (*ResourceHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9a40e43d3305340, []int{1}
}
func (m *ResourceHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHeader.Merge(m, src)
}
func (m *ResourceHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResourceHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHeader proto.InternalMessageInfo

func (m *ResourceHeader) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *ResourceHeader) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHeader) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceHeader) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *ResourceHeader) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *ResourceHeader) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *ResourceHeader) GetPreviousVersionId() string {
	if m != nil {
		return m.PreviousVersionId
	}
	return ""
}

func (m *ResourceHeader) GetNextVersionId() string {
	if m != nil {
		return m.NextVersionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "cheqdid.cheqdnode.resource.v1.Resource")
	proto.RegisterType((*ResourceHeader)(nil), "cheqdid.cheqdnode.resource.v1.ResourceHeader")
}

func init() { proto.RegisterFile("cheqd/resource/v1/resource.proto", fileDescriptor_f9a40e43d3305340) }

var fileDescriptor_f9a40e43d3305340 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x69, 0x2f, 0x17, 0xe8, 0x5c, 0xe0, 0xe6, 0xce, 0xdd, 0x34, 0x26, 0x34, 0x04, 0x13,
	0xc3, 0x86, 0x69, 0xd0, 0x37, 0x30, 0x31, 0x91, 0x6d, 0x63, 0x5c, 0xb8, 0x21, 0x65, 0xe6, 0xc4,
	0x4e, 0xa4, 0x9d, 0x3a, 0x9d, 0x36, 0xf0, 0x16, 0xbe, 0x87, 0x2f, 0xe2, 0x92, 0xa5, 0x4b, 0x03,
	0x2f, 0x62, 0x7a, 0x4a, 0x8b, 0x6e, 0xdc, 0xb4, 0xe7, 0xfc, 0xe7, 0xfb, 0xff, 0xf6, 0x64, 0x86,
	0x8c, 0x79, 0x04, 0xcf, 0xc2, 0xd7, 0x90, 0xa9, 0x5c, 0x73, 0xf0, 0x8b, 0x79, 0x53, 0xb3, 0x54,
	0x2b, 0xa3, 0xe8, 0x08, 0x09, 0x29, 0x18, 0xbe, 0x13, 0x25, 0x80, 0x35, 0x44, 0x31, 0x9f, 0x00,
	0xe9, 0x05, 0xc7, 0x96, 0xde, 0x90, 0x4e, 0x04, 0xa1, 0x00, 0xed, 0x5a, 0x63, 0x6b, 0xfa, 0xe7,
	0x72, 0xc6, 0x7e, 0xf4, 0xb2, 0xda, 0x78, 0x8b, 0xa6, 0xe0, 0x68, 0xa6, 0x94, 0xb4, 0x45, 0x68,
	0x42, 0xd7, 0x1e, 0x5b, 0xd3, 0x7e, 0x80, 0xf5, 0xe4, 0xd5, 0x26, 0xc3, 0xef, 0x38, 0x3d, 0x27,
	0x03, 0xae, 0xd6, 0x6b, 0xe0, 0x46, 0xaa, 0x64, 0x29, 0x05, 0x7e, 0xd4, 0x09, 0xfa, 0x27, 0x71,
	0x21, 0xe8, 0x90, 0xd8, 0x52, 0x60, 0x92, 0x13, 0xd8, 0x52, 0x94, 0xd9, 0x49, 0x18, 0x83, 0xfb,
	0x0b, 0x15, 0xac, 0xcb, 0xa0, 0xfa, 0xaf, 0x96, 0x66, 0x9b, 0x82, 0xdb, 0xae, 0x82, 0x6a, 0xf1,
	0x6e, 0x9b, 0x02, 0x1d, 0x11, 0x12, 0x83, 0x90, 0x61, 0x45, 0xfc, 0x46, 0xc2, 0x41, 0x05, 0xc7,
	0x2e, 0xe9, 0x72, 0x0d, 0xa1, 0x01, 0xe1, 0x76, 0x70, 0x56, 0xb7, 0xf4, 0x8c, 0xf4, 0x78, 0x04,
	0xfc, 0x29, 0xcb, 0x63, 0xb7, 0x8b, 0x1b, 0x35, 0x3d, 0x65, 0xe4, 0x7f, 0xaa, 0xa1, 0x90, 0x2a,
	0xcf, 0x96, 0x05, 0xe8, 0xec, 0xb8, 0x48, 0x0f, 0x13, 0xfe, 0xd5, 0xa3, 0xfb, 0x6a, 0xb2, 0x10,
	0xf4, 0x82, 0xfc, 0x4d, 0x60, 0x63, 0xbe, 0xb2, 0x0e, 0xb2, 0x83, 0x52, 0x6e, 0xb8, 0xeb, 0xc5,
	0xdb, 0xde, 0xb3, 0x76, 0x7b, 0xcf, 0xfa, 0xd8, 0x7b, 0xd6, 0xcb, 0xc1, 0x6b, 0xed, 0x0e, 0x5e,
	0xeb, 0xfd, 0xe0, 0xb5, 0x1e, 0xfc, 0x47, 0x69, 0xa2, 0x7c, 0xc5, 0xb8, 0x8a, 0xfd, 0xea, 0xe8,
	0xf1, 0x39, 0x2b, 0xcf, 0xc6, 0xdf, 0x9c, 0xee, 0x41, 0xb9, 0x65, 0xe6, 0x17, 0xf3, 0x55, 0x07,
	0x6f, 0xc1, 0xd5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x6a, 0xa6, 0xe7, 0x29, 0x02, 0x00,
	0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextVersionId) > 0 {
		i -= len(m.NextVersionId)
		copy(dAtA[i:], m.NextVersionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.NextVersionId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.PreviousVersionId) > 0 {
		i -= len(m.PreviousVersionId)
		copy(dAtA[i:], m.PreviousVersionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.PreviousVersionId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MediaType) > 0 {
		i -= len(m.MediaType)
		copy(dAtA[i:], m.MediaType)
		i = encodeVarintResource(dAtA, i, uint64(len(m.MediaType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ResourceType) > 0 {
		i -= len(m.ResourceType)
		copy(dAtA[i:], m.ResourceType)
		i = encodeVarintResource(dAtA, i, uint64(len(m.ResourceType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CollectionId) > 0 {
		i -= len(m.CollectionId)
		copy(dAtA[i:], m.CollectionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.CollectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *ResourceHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollectionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.ResourceType)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.MediaType)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.PreviousVersionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.NextVersionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResourceHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *ResourceHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: ResourceHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MediaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousVersionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousVersionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextVersionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextVersionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)