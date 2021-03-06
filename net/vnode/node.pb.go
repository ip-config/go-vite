// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package vnode

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

type PNode struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname             []byte   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	HostType             uint32   `protobuf:"varint,3,opt,name=hostType,proto3" json:"hostType,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Net                  uint32   `protobuf:"varint,5,opt,name=net,proto3" json:"net,omitempty"`
	Ext                  []byte   `protobuf:"bytes,6,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PNode) Reset()         { *m = PNode{} }
func (m *PNode) String() string { return proto.CompactTextString(m) }
func (*PNode) ProtoMessage()    {}
func (*PNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *PNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PNode.Unmarshal(m, b)
}
func (m *PNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PNode.Marshal(b, m, deterministic)
}
func (m *PNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PNode.Merge(m, src)
}
func (m *PNode) XXX_Size() int {
	return xxx_messageInfo_PNode.Size(m)
}
func (m *PNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PNode.DiscardUnknown(m)
}

var xxx_messageInfo_PNode proto.InternalMessageInfo

func (m *PNode) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PNode) GetHostname() []byte {
	if m != nil {
		return m.Hostname
	}
	return nil
}

func (m *PNode) GetHostType() uint32 {
	if m != nil {
		return m.HostType
	}
	return 0
}

func (m *PNode) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PNode) GetNet() uint32 {
	if m != nil {
		return m.Net
	}
	return 0
}

func (m *PNode) GetExt() []byte {
	if m != nil {
		return m.Ext
	}
	return nil
}

type PEndPoint struct {
	Host                 []byte   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	HostType             int32    `protobuf:"varint,3,opt,name=hostType,proto3" json:"hostType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PEndPoint) Reset()         { *m = PEndPoint{} }
func (m *PEndPoint) String() string { return proto.CompactTextString(m) }
func (*PEndPoint) ProtoMessage()    {}
func (*PEndPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *PEndPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PEndPoint.Unmarshal(m, b)
}
func (m *PEndPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PEndPoint.Marshal(b, m, deterministic)
}
func (m *PEndPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PEndPoint.Merge(m, src)
}
func (m *PEndPoint) XXX_Size() int {
	return xxx_messageInfo_PEndPoint.Size(m)
}
func (m *PEndPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PEndPoint.DiscardUnknown(m)
}

var xxx_messageInfo_PEndPoint proto.InternalMessageInfo

func (m *PEndPoint) GetHost() []byte {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *PEndPoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PEndPoint) GetHostType() int32 {
	if m != nil {
		return m.HostType
	}
	return 0
}

func init() {
	proto.RegisterType((*PNode)(nil), "vnode.pNode")
	proto.RegisterType((*PEndPoint)(nil), "vnode.pEndPoint")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x03, 0x71, 0x94, 0xda, 0x19, 0xb9, 0x58,
	0x0b, 0xfc, 0xf2, 0x53, 0x52, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x78, 0x82, 0x98, 0x32, 0x53, 0x84, 0xa4, 0xb8, 0x38, 0x32, 0xf2, 0x8b, 0x4b, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x98, 0xc0, 0xa2, 0x70, 0x3e, 0x4c, 0x2e, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x59, 0x81,
	0x51, 0x83, 0x37, 0x08, 0xce, 0x17, 0x12, 0xe2, 0x62, 0x29, 0xc8, 0x2f, 0x2a, 0x91, 0x60, 0x01,
	0x8b, 0x83, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa9, 0x25, 0x12, 0xac, 0x60, 0x21, 0x10, 0x13,
	0x24, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0x06, 0x36, 0x18, 0xc4, 0x54, 0xf2, 0xe7, 0xe2, 0x2c, 0x70,
	0xcd, 0x4b, 0x09, 0xc8, 0xcf, 0xcc, 0x2b, 0x01, 0x19, 0x02, 0x32, 0x10, 0xea, 0x1c, 0x30, 0x1b,
	0x6e, 0x30, 0xc8, 0x31, 0xac, 0x50, 0x83, 0xd1, 0x1d, 0xc2, 0x8a, 0x70, 0x48, 0x12, 0x1b, 0xd8,
	0xa3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x07, 0x37, 0x1a, 0xf6, 0x00, 0x00, 0x00,
}
