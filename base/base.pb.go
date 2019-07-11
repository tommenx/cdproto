// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package base

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BaseResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResp) Reset()         { *m = BaseResp{} }
func (m *BaseResp) String() string { return proto.CompactTextString(m) }
func (*BaseResp) ProtoMessage()    {}
func (*BaseResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *BaseResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResp.Unmarshal(m, b)
}
func (m *BaseResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResp.Marshal(b, m, deterministic)
}
func (m *BaseResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResp.Merge(m, src)
}
func (m *BaseResp) XXX_Size() int {
	return xxx_messageInfo_BaseResp.Size(m)
}
func (m *BaseResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResp.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResp proto.InternalMessageInfo

func (m *BaseResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Base struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Base) Reset()         { *m = Base{} }
func (m *Base) String() string { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()    {}
func (*Base) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *Base) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base.Unmarshal(m, b)
}
func (m *Base) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base.Marshal(b, m, deterministic)
}
func (m *Base) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base.Merge(m, src)
}
func (m *Base) XXX_Size() int {
	return xxx_messageInfo_Base.Size(m)
}
func (m *Base) XXX_DiscardUnknown() {
	xxx_messageInfo_Base.DiscardUnknown(m)
}

var xxx_messageInfo_Base proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseResp)(nil), "base.BaseResp")
	proto.RegisterType((*Base)(nil), "base.Base")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x2c, 0xb8, 0x38, 0x9c, 0x12,
	0x8b, 0x53, 0x83, 0x52, 0x8b, 0x0b, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4,
	0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x89, 0x8d, 0x8b, 0x05, 0xa4,
	0xd3, 0x49, 0x31, 0x4a, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf,
	0x24, 0x3f, 0x37, 0x37, 0x35, 0xaf, 0x42, 0x3f, 0x39, 0x05, 0x6c, 0x8b, 0x3e, 0xc8, 0x92, 0x24,
	0x36, 0x30, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x3f, 0xa6, 0x43, 0x7f, 0x00, 0x00,
	0x00,
}
