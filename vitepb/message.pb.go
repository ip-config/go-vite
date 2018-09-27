// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vitepb/message.proto

package vitepb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Handshake struct {
	NetID                uint64   `protobuf:"varint,1,opt,name=NetID,proto3" json:"NetID,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	Current              []byte   `protobuf:"bytes,4,opt,name=Current,proto3" json:"Current,omitempty"`
	Genesis              []byte   `protobuf:"bytes,5,opt,name=Genesis,proto3" json:"Genesis,omitempty"`
	Port                 uint32   `protobuf:"varint,6,opt,name=Port,proto3" json:"Port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{0}
}

func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (m *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(m, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func (m *Handshake) GetNetID() uint64 {
	if m != nil {
		return m.NetID
	}
	return 0
}

func (m *Handshake) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Handshake) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Handshake) GetCurrent() []byte {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *Handshake) GetGenesis() []byte {
	if m != nil {
		return m.Genesis
	}
	return nil
}

func (m *Handshake) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type BlockID struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{1}
}

func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=End,proto3" json:"End,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{2}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *File) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type FileList struct {
	Files                []*File  `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty"`
	Start                uint64   `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  uint64   `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"`
	Nonce                uint64   `protobuf:"varint,4,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileList) Reset()         { *m = FileList{} }
func (m *FileList) String() string { return proto.CompactTextString(m) }
func (*FileList) ProtoMessage()    {}
func (*FileList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{3}
}

func (m *FileList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileList.Unmarshal(m, b)
}
func (m *FileList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileList.Marshal(b, m, deterministic)
}
func (m *FileList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileList.Merge(m, src)
}
func (m *FileList) XXX_Size() int {
	return xxx_messageInfo_FileList.Size(m)
}
func (m *FileList) XXX_DiscardUnknown() {
	xxx_messageInfo_FileList.DiscardUnknown(m)
}

var xxx_messageInfo_FileList proto.InternalMessageInfo

func (m *FileList) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *FileList) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *FileList) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *FileList) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type RequestFile struct {
	File                 string   `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	Nonce                uint64   `protobuf:"varint,2,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFile) Reset()         { *m = RequestFile{} }
func (m *RequestFile) String() string { return proto.CompactTextString(m) }
func (*RequestFile) ProtoMessage()    {}
func (*RequestFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{4}
}

func (m *RequestFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFile.Unmarshal(m, b)
}
func (m *RequestFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFile.Marshal(b, m, deterministic)
}
func (m *RequestFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFile.Merge(m, src)
}
func (m *RequestFile) XXX_Size() int {
	return xxx_messageInfo_RequestFile.Size(m)
}
func (m *RequestFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFile.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFile proto.InternalMessageInfo

func (m *RequestFile) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *RequestFile) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type SubLedger struct {
	SBlocks              []*SnapshotBlock `protobuf:"bytes,1,rep,name=SBlocks,proto3" json:"SBlocks,omitempty"`
	ABlocks              []*AccountBlock  `protobuf:"bytes,2,rep,name=ABlocks,proto3" json:"ABlocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubLedger) Reset()         { *m = SubLedger{} }
func (m *SubLedger) String() string { return proto.CompactTextString(m) }
func (*SubLedger) ProtoMessage()    {}
func (*SubLedger) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{5}
}

func (m *SubLedger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubLedger.Unmarshal(m, b)
}
func (m *SubLedger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubLedger.Marshal(b, m, deterministic)
}
func (m *SubLedger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubLedger.Merge(m, src)
}
func (m *SubLedger) XXX_Size() int {
	return xxx_messageInfo_SubLedger.Size(m)
}
func (m *SubLedger) XXX_DiscardUnknown() {
	xxx_messageInfo_SubLedger.DiscardUnknown(m)
}

var xxx_messageInfo_SubLedger proto.InternalMessageInfo

func (m *SubLedger) GetSBlocks() []*SnapshotBlock {
	if m != nil {
		return m.SBlocks
	}
	return nil
}

func (m *SubLedger) GetABlocks() []*AccountBlock {
	if m != nil {
		return m.ABlocks
	}
	return nil
}

type GetSnapshotBlocks struct {
	From                 *BlockID `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	Forward              bool     `protobuf:"varint,3,opt,name=Forward,proto3" json:"Forward,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSnapshotBlocks) Reset()         { *m = GetSnapshotBlocks{} }
func (m *GetSnapshotBlocks) String() string { return proto.CompactTextString(m) }
func (*GetSnapshotBlocks) ProtoMessage()    {}
func (*GetSnapshotBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{6}
}

func (m *GetSnapshotBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSnapshotBlocks.Unmarshal(m, b)
}
func (m *GetSnapshotBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSnapshotBlocks.Marshal(b, m, deterministic)
}
func (m *GetSnapshotBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSnapshotBlocks.Merge(m, src)
}
func (m *GetSnapshotBlocks) XXX_Size() int {
	return xxx_messageInfo_GetSnapshotBlocks.Size(m)
}
func (m *GetSnapshotBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSnapshotBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_GetSnapshotBlocks proto.InternalMessageInfo

func (m *GetSnapshotBlocks) GetFrom() *BlockID {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetSnapshotBlocks) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetSnapshotBlocks) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

type SnapshotBlocks struct {
	Blocks               []*SnapshotBlock `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SnapshotBlocks) Reset()         { *m = SnapshotBlocks{} }
func (m *SnapshotBlocks) String() string { return proto.CompactTextString(m) }
func (*SnapshotBlocks) ProtoMessage()    {}
func (*SnapshotBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{7}
}

func (m *SnapshotBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBlocks.Unmarshal(m, b)
}
func (m *SnapshotBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBlocks.Marshal(b, m, deterministic)
}
func (m *SnapshotBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBlocks.Merge(m, src)
}
func (m *SnapshotBlocks) XXX_Size() int {
	return xxx_messageInfo_SnapshotBlocks.Size(m)
}
func (m *SnapshotBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBlocks proto.InternalMessageInfo

func (m *SnapshotBlocks) GetBlocks() []*SnapshotBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type GetAccountBlockByHeight struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	From                 *BlockID `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	Count                uint64   `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
	Forward              bool     `protobuf:"varint,4,opt,name=Forward,proto3" json:"Forward,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountBlockByHeight) Reset()         { *m = GetAccountBlockByHeight{} }
func (m *GetAccountBlockByHeight) String() string { return proto.CompactTextString(m) }
func (*GetAccountBlockByHeight) ProtoMessage()    {}
func (*GetAccountBlockByHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{8}
}

func (m *GetAccountBlockByHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountBlockByHeight.Unmarshal(m, b)
}
func (m *GetAccountBlockByHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountBlockByHeight.Marshal(b, m, deterministic)
}
func (m *GetAccountBlockByHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountBlockByHeight.Merge(m, src)
}
func (m *GetAccountBlockByHeight) XXX_Size() int {
	return xxx_messageInfo_GetAccountBlockByHeight.Size(m)
}
func (m *GetAccountBlockByHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountBlockByHeight.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountBlockByHeight proto.InternalMessageInfo

func (m *GetAccountBlockByHeight) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *GetAccountBlockByHeight) GetFrom() *BlockID {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetAccountBlockByHeight) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetAccountBlockByHeight) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

type AccountBlocksMsg struct {
	Address              []byte          `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Blocks               []*AccountBlock `protobuf:"bytes,3,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccountBlocksMsg) Reset()         { *m = AccountBlocksMsg{} }
func (m *AccountBlocksMsg) String() string { return proto.CompactTextString(m) }
func (*AccountBlocksMsg) ProtoMessage()    {}
func (*AccountBlocksMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{9}
}

func (m *AccountBlocksMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBlocksMsg.Unmarshal(m, b)
}
func (m *AccountBlocksMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBlocksMsg.Marshal(b, m, deterministic)
}
func (m *AccountBlocksMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBlocksMsg.Merge(m, src)
}
func (m *AccountBlocksMsg) XXX_Size() int {
	return xxx_messageInfo_AccountBlocksMsg.Size(m)
}
func (m *AccountBlocksMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBlocksMsg.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBlocksMsg proto.InternalMessageInfo

func (m *AccountBlocksMsg) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AccountBlocksMsg) GetBlocks() []*AccountBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*Handshake)(nil), "vitepb.Handshake")
	proto.RegisterType((*BlockID)(nil), "vitepb.BlockID")
	proto.RegisterType((*File)(nil), "vitepb.File")
	proto.RegisterType((*FileList)(nil), "vitepb.FileList")
	proto.RegisterType((*RequestFile)(nil), "vitepb.RequestFile")
	proto.RegisterType((*SubLedger)(nil), "vitepb.SubLedger")
	proto.RegisterType((*GetSnapshotBlocks)(nil), "vitepb.GetSnapshotBlocks")
	proto.RegisterType((*SnapshotBlocks)(nil), "vitepb.SnapshotBlocks")
	proto.RegisterType((*GetAccountBlockByHeight)(nil), "vitepb.GetAccountBlockByHeight")
	proto.RegisterType((*AccountBlocksMsg)(nil), "vitepb.AccountBlocksMsg")
}

func init() { proto.RegisterFile("vitepb/message.proto", fileDescriptor_2a6a8486deb9ab39) }

var fileDescriptor_2a6a8486deb9ab39 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x8e, 0x12, 0x41,
	0x10, 0xcd, 0xc0, 0x30, 0xec, 0x14, 0xa8, 0x6b, 0x07, 0x75, 0x82, 0x17, 0x32, 0x5e, 0x38, 0x28,
	0x9b, 0xac, 0x31, 0x1e, 0x0d, 0xec, 0x0a, 0x6c, 0xb2, 0x6e, 0x4c, 0x93, 0x78, 0xf0, 0x62, 0x06,
	0xa6, 0xc2, 0x4c, 0x16, 0xa6, 0xb1, 0xbb, 0xd1, 0xf8, 0x07, 0xfe, 0x83, 0x3f, 0x6b, 0xba, 0xba,
	0x1b, 0x07, 0xa3, 0x9b, 0x3d, 0x51, 0xaf, 0x5f, 0xbd, 0x9a, 0xf7, 0xaa, 0x1b, 0xe8, 0x7d, 0x2b,
	0x35, 0xee, 0x96, 0x67, 0x5b, 0x54, 0x2a, 0x5b, 0xe3, 0x68, 0x27, 0x85, 0x16, 0x2c, 0xb2, 0xa7,
	0xfd, 0xbe, 0x63, 0xb3, 0xd5, 0x4a, 0xec, 0x2b, 0xfd, 0x65, 0xb9, 0x11, 0xab, 0x5b, 0xdb, 0xd3,
	0x7f, 0xee, 0x38, 0x55, 0x65, 0x3b, 0x55, 0x88, 0x23, 0x32, 0xfd, 0x15, 0x40, 0x3c, 0xcf, 0xaa,
	0x5c, 0x15, 0xd9, 0x2d, 0xb2, 0x1e, 0xb4, 0x6e, 0x50, 0x5f, 0x5d, 0x26, 0xc1, 0x20, 0x18, 0x86,
	0xdc, 0x02, 0x96, 0x40, 0xfb, 0x13, 0x4a, 0x55, 0x8a, 0x2a, 0x69, 0xd0, 0xb9, 0x87, 0xec, 0x29,
	0x44, 0x73, 0x2c, 0xd7, 0x85, 0x4e, 0x9a, 0x44, 0x38, 0x64, 0x14, 0x17, 0x7b, 0x29, 0xb1, 0xd2,
	0x49, 0x38, 0x08, 0x86, 0x5d, 0xee, 0xa1, 0x61, 0x66, 0x58, 0xa1, 0x2a, 0x55, 0xd2, 0xb2, 0x8c,
	0x83, 0x8c, 0x41, 0xf8, 0x51, 0x48, 0x9d, 0x44, 0x83, 0x60, 0xf8, 0x80, 0x53, 0x9d, 0xbe, 0x81,
	0xf6, 0xc4, 0x98, 0xbd, 0xba, 0x34, 0xf4, 0x3c, 0x53, 0x05, 0x39, 0xeb, 0x72, 0xaa, 0x6b, 0x9f,
	0x6f, 0xd4, 0x3f, 0x9f, 0x4e, 0x20, 0x9c, 0x96, 0x1b, 0x34, 0x9a, 0x9b, 0x6c, 0x8b, 0xa4, 0x89,
	0x39, 0xd5, 0x26, 0xe2, 0x42, 0x67, 0xd2, 0x4a, 0x62, 0x6e, 0x01, 0x3b, 0x85, 0xe6, 0xfb, 0x2a,
	0xa7, 0x14, 0x31, 0x37, 0x65, 0xba, 0x81, 0x13, 0x33, 0xe3, 0xba, 0x54, 0x9a, 0xa5, 0xd0, 0x32,
	0xb5, 0x4a, 0x82, 0x41, 0x73, 0xd8, 0x39, 0xef, 0x8e, 0xec, 0x46, 0x47, 0xe6, 0x90, 0x5b, 0xea,
	0x78, 0x6e, 0xf8, 0x8f, 0xb9, 0x21, 0xcd, 0xa5, 0x15, 0x8b, 0x6a, 0x85, 0xb4, 0x18, 0xb3, 0x62,
	0x03, 0xd2, 0xb7, 0xd0, 0xe1, 0xf8, 0x75, 0x8f, 0x4a, 0x7b, 0xe3, 0xe6, 0xd7, 0x1b, 0xa7, 0xb3,
	0x83, 0xb0, 0x51, 0x17, 0x6e, 0x20, 0x5e, 0xec, 0x97, 0xd7, 0x98, 0xaf, 0x51, 0xb2, 0x33, 0x68,
	0x2f, 0x68, 0x5f, 0xde, 0xe9, 0x13, 0xef, 0x74, 0xe1, 0xee, 0x9e, 0x58, 0xee, 0xbb, 0xd8, 0x08,
	0xda, 0x63, 0x27, 0x68, 0x90, 0xa0, 0xe7, 0x05, 0x63, 0xfb, 0x90, 0x5c, 0xbf, 0x6b, 0x4a, 0x0b,
	0x78, 0x3c, 0x43, 0x7d, 0x34, 0x4c, 0xb1, 0x17, 0x10, 0x4e, 0xa5, 0xd8, 0x92, 0xd9, 0xce, 0xf9,
	0x23, 0x3f, 0xc1, 0x5d, 0x1c, 0x27, 0xd2, 0xb8, 0xbf, 0x30, 0x03, 0xbd, 0x7b, 0x02, 0xe6, 0x35,
	0x4c, 0x85, 0xfc, 0x9e, 0x49, 0xbb, 0xa2, 0x13, 0xee, 0x61, 0xfa, 0x0e, 0x1e, 0xfe, 0xf5, 0x99,
	0x57, 0x10, 0xdd, 0x27, 0x9b, 0x6b, 0x4a, 0x7f, 0x06, 0xf0, 0x6c, 0x86, 0xba, 0x9e, 0x63, 0xf2,
	0xe3, 0xcf, 0xf3, 0x1c, 0xe7, 0xb9, 0x44, 0xa5, 0xdc, 0x73, 0xf2, 0xf0, 0x90, 0xa5, 0x71, 0xaf,
	0x2c, 0xcd, 0xff, 0x64, 0x09, 0x8f, 0xb3, 0x7c, 0x86, 0xd3, 0xba, 0x0d, 0xf5, 0x41, 0xad, 0xef,
	0xb0, 0xf0, 0xf2, 0x90, 0xb3, 0x79, 0xc7, 0x95, 0xb8, 0x9e, 0x65, 0x44, 0x7f, 0xe3, 0xd7, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x48, 0xd2, 0xf2, 0xa2, 0x1f, 0x04, 0x00, 0x00,
}
