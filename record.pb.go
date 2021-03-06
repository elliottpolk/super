// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

package super

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RecordInfo_Status int32

const (
	RecordInfo_draft    RecordInfo_Status = 0
	RecordInfo_active   RecordInfo_Status = 1
	RecordInfo_archived RecordInfo_Status = 2
	RecordInfo_invalid  RecordInfo_Status = 3
)

var RecordInfo_Status_name = map[int32]string{
	0: "draft",
	1: "active",
	2: "archived",
	3: "invalid",
}

var RecordInfo_Status_value = map[string]int32{
	"draft":    0,
	"active":   1,
	"archived": 2,
	"invalid":  3,
}

func (x RecordInfo_Status) String() string {
	return proto.EnumName(RecordInfo_Status_name, int32(x))
}

func (RecordInfo_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0, 0}
}

type RecordInfo struct {
	Created              *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy            string               `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	UpdatedBy            string               `protobuf:"bytes,4,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Status               RecordInfo_Status    `protobuf:"varint,5,opt,name=status,proto3,enum=super.RecordInfo_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecordInfo) Reset()         { *m = RecordInfo{} }
func (m *RecordInfo) String() string { return proto.CompactTextString(m) }
func (*RecordInfo) ProtoMessage()    {}
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

func (m *RecordInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordInfo.Unmarshal(m, b)
}
func (m *RecordInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordInfo.Marshal(b, m, deterministic)
}
func (m *RecordInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordInfo.Merge(m, src)
}
func (m *RecordInfo) XXX_Size() int {
	return xxx_messageInfo_RecordInfo.Size(m)
}
func (m *RecordInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecordInfo proto.InternalMessageInfo

func (m *RecordInfo) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *RecordInfo) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *RecordInfo) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *RecordInfo) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *RecordInfo) GetStatus() RecordInfo_Status {
	if m != nil {
		return m.Status
	}
	return RecordInfo_draft
}

func init() {
	proto.RegisterEnum("super.RecordInfo_Status", RecordInfo_Status_name, RecordInfo_Status_value)
	proto.RegisterType((*RecordInfo)(nil), "super.RecordInfo")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xcd, 0xe6, 0x3a, 0xf7, 0x6d, 0x48, 0x09, 0x08, 0x65, 0x28, 0x8e, 0x9d, 0x7a, 0x4a,
	0x65, 0x7a, 0xf2, 0x58, 0xbc, 0x78, 0x1b, 0xd5, 0x93, 0x17, 0x49, 0x93, 0xb4, 0x0b, 0xb6, 0x4b,
	0x49, 0xd3, 0x42, 0xff, 0x8f, 0x3f, 0x54, 0x9a, 0xa4, 0xe8, 0xcd, 0xdb, 0xc7, 0x9b, 0x27, 0xcf,
	0xcb, 0x0b, 0x1b, 0x2d, 0x98, 0xd2, 0x9c, 0x34, 0x5a, 0x19, 0x85, 0x17, 0x6d, 0xd7, 0x08, 0xbd,
	0xbd, 0x2f, 0x95, 0x2a, 0x2b, 0x91, 0xd8, 0x30, 0xef, 0x8a, 0xc4, 0xc8, 0x5a, 0xb4, 0x86, 0xd6,
	0x8d, 0xe3, 0xf6, 0xdf, 0x33, 0x80, 0xcc, 0x7e, 0x7c, 0x3d, 0x17, 0x0a, 0x3f, 0xc1, 0x92, 0x69,
	0x41, 0x8d, 0xe0, 0x11, 0xda, 0xa1, 0x78, 0x7d, 0xd8, 0x12, 0x67, 0x20, 0x93, 0x81, 0xbc, 0x4f,
	0x86, 0x6c, 0x42, 0xf1, 0x1d, 0x80, 0x3f, 0x3f, 0xf3, 0x21, 0x9a, 0xed, 0x50, 0xbc, 0xca, 0x56,
	0x3e, 0x49, 0x87, 0x51, 0xda, 0x35, 0xdc, 0x4a, 0xe7, 0xff, 0x4b, 0x3d, 0x3a, 0x4a, 0xfd, 0x39,
	0x4a, 0x2f, 0x9d, 0xd4, 0x27, 0xe9, 0x80, 0x1f, 0x20, 0x68, 0x0d, 0x35, 0x5d, 0x1b, 0x2d, 0x76,
	0x28, 0xbe, 0x3e, 0x44, 0xc4, 0x2e, 0x26, 0xbf, 0x63, 0xc8, 0x9b, 0x7d, 0xcf, 0x3c, 0xb7, 0x7f,
	0x86, 0xc0, 0x25, 0x78, 0x05, 0x0b, 0xae, 0x69, 0x61, 0xc2, 0x0b, 0x0c, 0x10, 0x50, 0x66, 0x64,
	0x2f, 0x42, 0x84, 0x37, 0x70, 0x45, 0x35, 0x3b, 0xc9, 0x5e, 0xf0, 0x70, 0x86, 0xd7, 0xb0, 0x94,
	0xe7, 0x9e, 0x56, 0x92, 0x87, 0xf3, 0xf4, 0x05, 0x6e, 0x98, 0xaa, 0x89, 0xa8, 0x2a, 0xa9, 0x8c,
	0x69, 0x54, 0xf5, 0xe5, 0xea, 0xd2, 0xb5, 0xeb, 0x3b, 0x8e, 0x43, 0x8e, 0xe8, 0xe3, 0xb6, 0x94,
	0xe6, 0xd4, 0xe5, 0x84, 0xa9, 0x3a, 0xf9, 0x03, 0x27, 0x16, 0xce, 0x03, 0xbb, 0xf7, 0xf1, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0xc1, 0xb8, 0x13, 0xab, 0x01, 0x00, 0x00,
}
