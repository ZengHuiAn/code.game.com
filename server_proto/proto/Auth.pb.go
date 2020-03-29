// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Auth.proto

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

type AuthRequest struct {
	AuthToken            []byte   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0be2fdcb97a9c409, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetAuthToken() []byte {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

type AuthRespone struct {
	ErrorCode            uint32   `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	AuthToken            []byte   `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRespone) Reset()         { *m = AuthRespone{} }
func (m *AuthRespone) String() string { return proto.CompactTextString(m) }
func (*AuthRespone) ProtoMessage()    {}
func (*AuthRespone) Descriptor() ([]byte, []int) {
	return fileDescriptor_0be2fdcb97a9c409, []int{1}
}

func (m *AuthRespone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRespone.Unmarshal(m, b)
}
func (m *AuthRespone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRespone.Marshal(b, m, deterministic)
}
func (m *AuthRespone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRespone.Merge(m, src)
}
func (m *AuthRespone) XXX_Size() int {
	return xxx_messageInfo_AuthRespone.Size(m)
}
func (m *AuthRespone) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRespone.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRespone proto.InternalMessageInfo

func (m *AuthRespone) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *AuthRespone) GetAuthToken() []byte {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "proto.AuthRequest")
	proto.RegisterType((*AuthRespone)(nil), "proto.AuthRespone")
}

func init() {
	proto.RegisterFile("Auth.proto", fileDescriptor_0be2fdcb97a9c409)
}

var fileDescriptor_0be2fdcb97a9c409 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x72, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xda, 0x5c, 0xdc, 0x20, 0xc1,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e, 0xce, 0xc4, 0xd2, 0x92, 0x8c, 0x90,
	0xfc, 0xec, 0xd4, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x84, 0x80, 0x92, 0x27, 0x4c,
	0x71, 0x71, 0x41, 0x7e, 0x5e, 0x2a, 0x48, 0x71, 0x6a, 0x51, 0x51, 0x7e, 0x91, 0x73, 0x7e, 0x4a,
	0x2a, 0x58, 0x31, 0x6f, 0x10, 0x42, 0x00, 0xd5, 0x28, 0x26, 0x34, 0xa3, 0x8c, 0xec, 0xb8, 0x58,
	0x40, 0x46, 0x09, 0x99, 0x41, 0x1c, 0xe5, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0x22, 0x24, 0x04, 0x71,
	0x9c, 0x1e, 0x92, 0x93, 0xa4, 0x50, 0xc5, 0xc0, 0x36, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x05, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x97, 0x3b, 0xb4, 0xd3, 0x00, 0x00, 0x00,
}
