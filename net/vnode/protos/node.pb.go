// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package protos

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

type Node struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Hostname             []byte   `protobuf:"bytes,2,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	HostType             uint32   `protobuf:"varint,3,opt,name=HostType,proto3" json:"HostType,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	Net                  uint32   `protobuf:"varint,5,opt,name=Net,proto3" json:"Net,omitempty"`
	Ext                  []byte   `protobuf:"bytes,6,opt,name=Ext,proto3" json:"Ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Node) GetHostname() []byte {
	if m != nil {
		return m.Hostname
	}
	return nil
}

func (m *Node) GetHostType() uint32 {
	if m != nil {
		return m.HostType
	}
	return 0
}

func (m *Node) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetNet() uint32 {
	if m != nil {
		return m.Net
	}
	return 0
}

func (m *Node) GetExt() []byte {
	if m != nil {
		return m.Ext
	}
	return nil
}

type EndPoint struct {
	Host                 []byte   `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	HostType             int32    `protobuf:"varint,3,opt,name=HostType,proto3" json:"HostType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndPoint) Reset()         { *m = EndPoint{} }
func (m *EndPoint) String() string { return proto.CompactTextString(m) }
func (*EndPoint) ProtoMessage()    {}
func (*EndPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *EndPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndPoint.Unmarshal(m, b)
}
func (m *EndPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndPoint.Marshal(b, m, deterministic)
}
func (m *EndPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndPoint.Merge(m, src)
}
func (m *EndPoint) XXX_Size() int {
	return xxx_messageInfo_EndPoint.Size(m)
}
func (m *EndPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_EndPoint.DiscardUnknown(m)
}

var xxx_messageInfo_EndPoint proto.InternalMessageInfo

func (m *EndPoint) GetHost() []byte {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *EndPoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EndPoint) GetHostType() int32 {
	if m != nil {
		return m.HostType
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*EndPoint)(nil), "protos.EndPoint")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x6d, 0x8c, 0x5c, 0x2c,
	0x7e, 0xf9, 0x29, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x52, 0x5c, 0x1c, 0x1e, 0xf9, 0xc5, 0x25, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x4c, 0x60, 0x51, 0x38, 0x1f, 0x26, 0x17, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x1b, 0x04, 0xe7, 0x0b, 0x09, 0x71, 0xb1, 0x04, 0xe4, 0x17, 0x95, 0x48, 0xb0, 0x80, 0xc5,
	0xc1, 0x6c, 0x21, 0x01, 0x2e, 0x66, 0xbf, 0xd4, 0x12, 0x09, 0x56, 0xb0, 0x10, 0x88, 0x09, 0x12,
	0x71, 0xad, 0x28, 0x91, 0x60, 0x03, 0x1b, 0x0c, 0x62, 0x2a, 0xf9, 0x71, 0x71, 0xb8, 0xe6, 0xa5,
	0x04, 0xe4, 0x67, 0xe6, 0x95, 0x80, 0xcc, 0x00, 0x99, 0x07, 0x75, 0x0d, 0x98, 0x0d, 0x37, 0x17,
	0xe4, 0x16, 0x56, 0xa8, 0xb9, 0xe8, 0xee, 0x60, 0x45, 0xb8, 0x23, 0x09, 0xe2, 0x41, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x32, 0x07, 0x4b, 0xf5, 0x00, 0x00, 0x00,
}