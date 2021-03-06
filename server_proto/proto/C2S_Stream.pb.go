// Code generated by protoc-gen-go. DO NOT EDIT.
// source: C2S_Stream.proto

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

type C2S_RouterRequest struct {
	BodyMsg              []byte   `protobuf:"bytes,1,opt,name=bodyMsg,proto3" json:"bodyMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_RouterRequest) Reset()         { *m = C2S_RouterRequest{} }
func (m *C2S_RouterRequest) String() string { return proto.CompactTextString(m) }
func (*C2S_RouterRequest) ProtoMessage()    {}
func (*C2S_RouterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45418ee92d27b7dd, []int{0}
}

func (m *C2S_RouterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_RouterRequest.Unmarshal(m, b)
}
func (m *C2S_RouterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_RouterRequest.Marshal(b, m, deterministic)
}
func (m *C2S_RouterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_RouterRequest.Merge(m, src)
}
func (m *C2S_RouterRequest) XXX_Size() int {
	return xxx_messageInfo_C2S_RouterRequest.Size(m)
}
func (m *C2S_RouterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_RouterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_RouterRequest proto.InternalMessageInfo

func (m *C2S_RouterRequest) GetBodyMsg() []byte {
	if m != nil {
		return m.BodyMsg
	}
	return nil
}

type S2C_RouterRespone struct {
	Pids                 []uint32 `protobuf:"varint,1,rep,packed,name=pids,proto3" json:"pids,omitempty"`
	BodyMsg              []byte   `protobuf:"bytes,2,opt,name=bodyMsg,proto3" json:"bodyMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_RouterRespone) Reset()         { *m = S2C_RouterRespone{} }
func (m *S2C_RouterRespone) String() string { return proto.CompactTextString(m) }
func (*S2C_RouterRespone) ProtoMessage()    {}
func (*S2C_RouterRespone) Descriptor() ([]byte, []int) {
	return fileDescriptor_45418ee92d27b7dd, []int{1}
}

func (m *S2C_RouterRespone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_RouterRespone.Unmarshal(m, b)
}
func (m *S2C_RouterRespone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_RouterRespone.Marshal(b, m, deterministic)
}
func (m *S2C_RouterRespone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_RouterRespone.Merge(m, src)
}
func (m *S2C_RouterRespone) XXX_Size() int {
	return xxx_messageInfo_S2C_RouterRespone.Size(m)
}
func (m *S2C_RouterRespone) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_RouterRespone.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_RouterRespone proto.InternalMessageInfo

func (m *S2C_RouterRespone) GetPids() []uint32 {
	if m != nil {
		return m.Pids
	}
	return nil
}

func (m *S2C_RouterRespone) GetBodyMsg() []byte {
	if m != nil {
		return m.BodyMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S_RouterRequest)(nil), "proto.C2S_RouterRequest")
	proto.RegisterType((*S2C_RouterRespone)(nil), "proto.S2C_RouterRespone")
}

func init() {
	proto.RegisterFile("C2S_Stream.proto", fileDescriptor_45418ee92d27b7dd)
}

var fileDescriptor_45418ee92d27b7dd = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x70, 0x36, 0x0a, 0x8e,
	0x0f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0xba, 0x5c, 0x82, 0x20, 0xa9, 0xa0, 0xfc, 0xd2, 0x92, 0xd4, 0xa2, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xa4, 0xfc, 0x94, 0x4a, 0xdf, 0xe2, 0x74, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0xc9, 0x91, 0x4b, 0x30, 0xd8, 0xc8, 0x19, 0xae, 0xbc,
	0xb8, 0x20, 0x3f, 0x2f, 0x55, 0x48, 0x88, 0x8b, 0xa5, 0x20, 0x33, 0xa5, 0x58, 0x82, 0x51, 0x81,
	0x59, 0x83, 0x37, 0x08, 0xcc, 0x46, 0x36, 0x82, 0x09, 0xc5, 0x08, 0xa3, 0x10, 0x2e, 0x2e, 0x84,
	0x8d, 0x42, 0x6e, 0x5c, 0x5c, 0x8e, 0xa5, 0x25, 0x19, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x25, 0x42,
	0x12, 0x10, 0xc7, 0xe9, 0x61, 0x38, 0x49, 0x0a, 0x26, 0x83, 0x61, 0xbb, 0x12, 0x83, 0x06, 0xa3,
	0x01, 0x63, 0x12, 0x1b, 0x58, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x72, 0xd7, 0xd8,
	0xe9, 0x00, 0x00, 0x00,
}
