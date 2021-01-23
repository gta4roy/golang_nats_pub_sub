// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

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

type GetSystemTime struct {
	Systemtime           string   `protobuf:"bytes,1,opt,name=systemtime,proto3" json:"systemtime,omitempty"`
	Systemdate           string   `protobuf:"bytes,2,opt,name=systemdate,proto3" json:"systemdate,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Serverip             string   `protobuf:"bytes,4,opt,name=serverip,proto3" json:"serverip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSystemTime) Reset()         { *m = GetSystemTime{} }
func (m *GetSystemTime) String() string { return proto.CompactTextString(m) }
func (*GetSystemTime) ProtoMessage()    {}
func (*GetSystemTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *GetSystemTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSystemTime.Unmarshal(m, b)
}
func (m *GetSystemTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSystemTime.Marshal(b, m, deterministic)
}
func (m *GetSystemTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSystemTime.Merge(m, src)
}
func (m *GetSystemTime) XXX_Size() int {
	return xxx_messageInfo_GetSystemTime.Size(m)
}
func (m *GetSystemTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSystemTime.DiscardUnknown(m)
}

var xxx_messageInfo_GetSystemTime proto.InternalMessageInfo

func (m *GetSystemTime) GetSystemtime() string {
	if m != nil {
		return m.Systemtime
	}
	return ""
}

func (m *GetSystemTime) GetSystemdate() string {
	if m != nil {
		return m.Systemdate
	}
	return ""
}

func (m *GetSystemTime) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetSystemTime) GetServerip() string {
	if m != nil {
		return m.Serverip
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSystemTime)(nil), "order.GetSystemTime")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xda, 0x19, 0xb9, 0x78,
	0xdd, 0x53, 0x4b, 0x82, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0x43, 0x32, 0x73, 0x53, 0x85, 0xe4, 0xb8,
	0xb8, 0x8a, 0xc1, 0xbc, 0x12, 0x20, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0x04,
	0x21, 0x9f, 0x92, 0x58, 0x92, 0x2a, 0xc1, 0x84, 0x2c, 0x0f, 0x12, 0x11, 0x92, 0xe2, 0xe2, 0x28,
	0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0x04, 0xea, 0x66, 0x06, 0xcb, 0xc2, 0xf9, 0x20, 0x39, 0x20, 0xb3,
	0x2c, 0xb5, 0x28, 0xb3, 0x40, 0x82, 0x05, 0x22, 0x07, 0xe3, 0x27, 0xb1, 0x81, 0xdd, 0x65, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x66, 0x0d, 0x19, 0xa6, 0x00, 0x00, 0x00,
}
