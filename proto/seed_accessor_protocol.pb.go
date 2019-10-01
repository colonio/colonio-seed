// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/seed_accessor_protocol.proto

package proto

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

type SeedAccessor struct {
	DstNid               *NodeID  `protobuf:"bytes,1,opt,name=dst_nid,json=dstNid,proto3" json:"dst_nid,omitempty"`
	SrcNid               *NodeID  `protobuf:"bytes,2,opt,name=src_nid,json=srcNid,proto3" json:"src_nid,omitempty"`
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Mode                 uint32   `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
	Channel              uint32   `protobuf:"varint,5,opt,name=channel,proto3" json:"channel,omitempty"`
	CommandId            uint32   `protobuf:"varint,6,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Content              []byte   `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeedAccessor) Reset()         { *m = SeedAccessor{} }
func (m *SeedAccessor) String() string { return proto.CompactTextString(m) }
func (*SeedAccessor) ProtoMessage()    {}
func (*SeedAccessor) Descriptor() ([]byte, []int) {
	return fileDescriptor_64bbe10f5f81d39f, []int{0}
}

func (m *SeedAccessor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeedAccessor.Unmarshal(m, b)
}
func (m *SeedAccessor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeedAccessor.Marshal(b, m, deterministic)
}
func (m *SeedAccessor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeedAccessor.Merge(m, src)
}
func (m *SeedAccessor) XXX_Size() int {
	return xxx_messageInfo_SeedAccessor.Size(m)
}
func (m *SeedAccessor) XXX_DiscardUnknown() {
	xxx_messageInfo_SeedAccessor.DiscardUnknown(m)
}

var xxx_messageInfo_SeedAccessor proto.InternalMessageInfo

func (m *SeedAccessor) GetDstNid() *NodeID {
	if m != nil {
		return m.DstNid
	}
	return nil
}

func (m *SeedAccessor) GetSrcNid() *NodeID {
	if m != nil {
		return m.SrcNid
	}
	return nil
}

func (m *SeedAccessor) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SeedAccessor) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *SeedAccessor) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *SeedAccessor) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

func (m *SeedAccessor) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Auth struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Hint                 uint32   `protobuf:"varint,2,opt,name=hint,proto3" json:"hint,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_64bbe10f5f81d39f, []int{1}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Auth) GetHint() uint32 {
	if m != nil {
		return m.Hint
	}
	return 0
}

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthSuccess struct {
	Config               string   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSuccess) Reset()         { *m = AuthSuccess{} }
func (m *AuthSuccess) String() string { return proto.CompactTextString(m) }
func (*AuthSuccess) ProtoMessage()    {}
func (*AuthSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_64bbe10f5f81d39f, []int{2}
}

func (m *AuthSuccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSuccess.Unmarshal(m, b)
}
func (m *AuthSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSuccess.Marshal(b, m, deterministic)
}
func (m *AuthSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSuccess.Merge(m, src)
}
func (m *AuthSuccess) XXX_Size() int {
	return xxx_messageInfo_AuthSuccess.Size(m)
}
func (m *AuthSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSuccess proto.InternalMessageInfo

func (m *AuthSuccess) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type Hint struct {
	Hint                 uint32   `protobuf:"varint,1,opt,name=hint,proto3" json:"hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hint) Reset()         { *m = Hint{} }
func (m *Hint) String() string { return proto.CompactTextString(m) }
func (*Hint) ProtoMessage()    {}
func (*Hint) Descriptor() ([]byte, []int) {
	return fileDescriptor_64bbe10f5f81d39f, []int{3}
}

func (m *Hint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hint.Unmarshal(m, b)
}
func (m *Hint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hint.Marshal(b, m, deterministic)
}
func (m *Hint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hint.Merge(m, src)
}
func (m *Hint) XXX_Size() int {
	return xxx_messageInfo_Hint.Size(m)
}
func (m *Hint) XXX_DiscardUnknown() {
	xxx_messageInfo_Hint.DiscardUnknown(m)
}

var xxx_messageInfo_Hint proto.InternalMessageInfo

func (m *Hint) GetHint() uint32 {
	if m != nil {
		return m.Hint
	}
	return 0
}

func init() {
	proto.RegisterType((*SeedAccessor)(nil), "colonio.SeedAccessorProtocol.SeedAccessor")
	proto.RegisterType((*Auth)(nil), "colonio.SeedAccessorProtocol.Auth")
	proto.RegisterType((*AuthSuccess)(nil), "colonio.SeedAccessorProtocol.AuthSuccess")
	proto.RegisterType((*Hint)(nil), "colonio.SeedAccessorProtocol.Hint")
}

func init() { proto.RegisterFile("core/seed_accessor_protocol.proto", fileDescriptor_64bbe10f5f81d39f) }

var fileDescriptor_64bbe10f5f81d39f = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x9b, 0xa6, 0x64, 0xda, 0x7a, 0x58, 0x45, 0x96, 0xa2, 0x50, 0x03, 0x42, 0x4f,
	0x11, 0xf5, 0x09, 0x2a, 0x1e, 0x5a, 0x0f, 0x45, 0xd2, 0x9b, 0x97, 0x50, 0x77, 0x46, 0xbb, 0xd8,
	0xec, 0x48, 0x76, 0xeb, 0x4b, 0xfb, 0x12, 0x92, 0x31, 0x41, 0x3d, 0x79, 0xda, 0xf9, 0x67, 0xff,
	0xff, 0x63, 0x66, 0x17, 0x2e, 0x0c, 0xd7, 0x74, 0xe5, 0x89, 0xb0, 0xdc, 0x1a, 0x43, 0xde, 0x73,
	0x5d, 0xbe, 0xd7, 0x1c, 0xd8, 0xf0, 0x3e, 0x97, 0x42, 0x9d, 0x19, 0xde, 0xb3, 0xb3, 0x9c, 0x6f,
	0x88, 0x70, 0xd1, 0x9a, 0x1e, 0x5b, 0xcf, 0xf4, 0x58, 0x00, 0x7f, 0x23, 0xd9, 0x67, 0x04, 0xe3,
	0xdf, 0x6e, 0x75, 0x0d, 0x43, 0xf4, 0xa1, 0x74, 0x16, 0x75, 0x34, 0x8b, 0xe6, 0xa3, 0x1b, 0x9d,
	0x77, 0xd4, 0x8e, 0x94, 0xaf, 0x19, 0x69, 0x75, 0x5f, 0x24, 0xe8, 0xc3, 0xda, 0x62, 0x13, 0xf1,
	0xb5, 0x91, 0x48, 0xef, 0xbf, 0x88, 0xaf, 0x4d, 0x13, 0x39, 0x82, 0x9e, 0x45, 0xdd, 0x9f, 0x45,
	0xf3, 0x49, 0xd1, 0xb3, 0xa8, 0x14, 0xc4, 0x15, 0x23, 0xe9, 0x58, 0x3a, 0x52, 0x2b, 0x0d, 0x43,
	0xb3, 0xdb, 0x3a, 0x47, 0x7b, 0x3d, 0x90, 0x76, 0x27, 0xd5, 0x39, 0x80, 0xe1, 0xaa, 0xda, 0x3a,
	0x2c, 0x2d, 0xea, 0x44, 0x2e, 0xd3, 0xb6, 0xb3, 0x42, 0x09, 0xb2, 0x0b, 0xe4, 0x82, 0x1e, 0xce,
	0xa2, 0xf9, 0xb8, 0xe8, 0x64, 0xf6, 0x00, 0xf1, 0xe2, 0x10, 0x76, 0x8d, 0xe3, 0x83, 0x6a, 0x6f,
	0xd9, 0xc9, 0x92, 0x69, 0xd1, 0xc9, 0x66, 0x90, 0x9d, 0x75, 0x41, 0x16, 0x99, 0x14, 0x52, 0xab,
	0x13, 0x18, 0x04, 0x7e, 0x23, 0x27, 0xf3, 0xa6, 0xc5, 0xb7, 0xc8, 0x2e, 0x61, 0xd4, 0xb0, 0x36,
	0x07, 0x79, 0x39, 0x75, 0x0a, 0x89, 0x61, 0xf7, 0x62, 0x5f, 0x5b, 0x62, 0xab, 0xb2, 0x29, 0xc4,
	0xcb, 0x06, 0xd2, 0x81, 0xa3, 0x1f, 0xf0, 0x5d, 0xba, 0xec, 0x3f, 0x0d, 0xe4, 0x1f, 0x9e, 0x13,
	0x39, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x33, 0x95, 0x2b, 0xe6, 0x01, 0x00, 0x00,
}