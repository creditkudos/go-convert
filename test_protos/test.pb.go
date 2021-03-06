// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test_protos

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

type Master struct {
	A                    string    `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32     `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    int64     `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D                    uint32    `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E                    uint64    `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F                    float32   `protobuf:"fixed32,6,opt,name=f,proto3" json:"f,omitempty"`
	G                    float64   `protobuf:"fixed64,7,opt,name=g,proto3" json:"g,omitempty"`
	H                    bool      `protobuf:"varint,8,opt,name=h,proto3" json:"h,omitempty"`
	I                    []*Minion `protobuf:"bytes,9,rep,name=i,proto3" json:"i,omitempty"`
	J                    *Minion   `protobuf:"bytes,10,opt,name=j,proto3" json:"j,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Master) Reset()         { *m = Master{} }
func (m *Master) String() string { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()    {}
func (*Master) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Master) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Master.Unmarshal(m, b)
}
func (m *Master) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Master.Marshal(b, m, deterministic)
}
func (m *Master) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Master.Merge(m, src)
}
func (m *Master) XXX_Size() int {
	return xxx_messageInfo_Master.Size(m)
}
func (m *Master) XXX_DiscardUnknown() {
	xxx_messageInfo_Master.DiscardUnknown(m)
}

var xxx_messageInfo_Master proto.InternalMessageInfo

func (m *Master) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *Master) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Master) GetC() int64 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *Master) GetD() uint32 {
	if m != nil {
		return m.D
	}
	return 0
}

func (m *Master) GetE() uint64 {
	if m != nil {
		return m.E
	}
	return 0
}

func (m *Master) GetF() float32 {
	if m != nil {
		return m.F
	}
	return 0
}

func (m *Master) GetG() float64 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *Master) GetH() bool {
	if m != nil {
		return m.H
	}
	return false
}

func (m *Master) GetI() []*Minion {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *Master) GetJ() *Minion {
	if m != nil {
		return m.J
	}
	return nil
}

type Minion struct {
	Ma                   string   `protobuf:"bytes,1,opt,name=ma,proto3" json:"ma,omitempty"`
	Mb                   int32    `protobuf:"varint,2,opt,name=mb,proto3" json:"mb,omitempty"`
	Mc                   []*Child `protobuf:"bytes,3,rep,name=mc,proto3" json:"mc,omitempty"`
	Md                   *Child   `protobuf:"bytes,4,opt,name=md,proto3" json:"md,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Minion) Reset()         { *m = Minion{} }
func (m *Minion) String() string { return proto.CompactTextString(m) }
func (*Minion) ProtoMessage()    {}
func (*Minion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Minion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Minion.Unmarshal(m, b)
}
func (m *Minion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Minion.Marshal(b, m, deterministic)
}
func (m *Minion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minion.Merge(m, src)
}
func (m *Minion) XXX_Size() int {
	return xxx_messageInfo_Minion.Size(m)
}
func (m *Minion) XXX_DiscardUnknown() {
	xxx_messageInfo_Minion.DiscardUnknown(m)
}

var xxx_messageInfo_Minion proto.InternalMessageInfo

func (m *Minion) GetMa() string {
	if m != nil {
		return m.Ma
	}
	return ""
}

func (m *Minion) GetMb() int32 {
	if m != nil {
		return m.Mb
	}
	return 0
}

func (m *Minion) GetMc() []*Child {
	if m != nil {
		return m.Mc
	}
	return nil
}

func (m *Minion) GetMd() *Child {
	if m != nil {
		return m.Md
	}
	return nil
}

type Child struct {
	Ca                   string   `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *Child) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Child.Unmarshal(m, b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Child.Marshal(b, m, deterministic)
}
func (m *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(m, src)
}
func (m *Child) XXX_Size() int {
	return xxx_messageInfo_Child.Size(m)
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetCa() string {
	if m != nil {
		return m.Ca
	}
	return ""
}

type RepeatedFailure struct {
	S                    []string `protobuf:"bytes,1,rep,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatedFailure) Reset()         { *m = RepeatedFailure{} }
func (m *RepeatedFailure) String() string { return proto.CompactTextString(m) }
func (*RepeatedFailure) ProtoMessage()    {}
func (*RepeatedFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *RepeatedFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedFailure.Unmarshal(m, b)
}
func (m *RepeatedFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedFailure.Marshal(b, m, deterministic)
}
func (m *RepeatedFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedFailure.Merge(m, src)
}
func (m *RepeatedFailure) XXX_Size() int {
	return xxx_messageInfo_RepeatedFailure.Size(m)
}
func (m *RepeatedFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedFailure.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedFailure proto.InternalMessageInfo

func (m *RepeatedFailure) GetS() []string {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterType((*Master)(nil), "test_protos.Master")
	proto.RegisterType((*Minion)(nil), "test_protos.Minion")
	proto.RegisterType((*Child)(nil), "test_protos.Child")
	proto.RegisterType((*RepeatedFailure)(nil), "test_protos.RepeatedFailure")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0x75, 0x4e, 0x1b, 0x5a, 0x97, 0x82, 0x64, 0x06, 0xbc, 0xf1, 0xc8, 0xe4, 0xa9, 0x03,
	0xfc, 0x03, 0x90, 0xd8, 0xba, 0xbc, 0x91, 0x05, 0x39, 0x89, 0xdb, 0xba, 0x6a, 0x9a, 0x28, 0x09,
	0x7f, 0x8f, 0xdf, 0x86, 0x9e, 0xab, 0xa2, 0x32, 0x74, 0xf2, 0x7d, 0xd2, 0xc9, 0xf7, 0xee, 0xb4,
	0x1e, 0xc3, 0x30, 0xae, 0xba, 0xbe, 0x1d, 0x5b, 0xb3, 0x10, 0xfd, 0x95, 0xf4, 0x50, 0xfc, 0x40,
	0xe7, 0x6b, 0x3f, 0x8c, 0xa1, 0x37, 0xb7, 0x1a, 0xde, 0x82, 0xe0, 0xe6, 0x0c, 0x2f, 0x54, 0x5a,
	0x45, 0x70, 0x53, 0x46, 0x29, 0x54, 0xd9, 0x8c, 0xe0, 0x32, 0x46, 0x25, 0x54, 0xdb, 0x09, 0xc1,
	0x2d, 0x19, 0xb5, 0x50, 0xb0, 0x53, 0x82, 0x9b, 0x30, 0x82, 0xd0, 0xc6, 0xe6, 0x04, 0xa7, 0x18,
	0x1b, 0xa1, 0xad, 0xbd, 0x21, 0x38, 0x30, 0xb6, 0x42, 0x3b, 0x3b, 0x23, 0xb8, 0x19, 0x63, 0x67,
	0x9e, 0x35, 0xa2, 0x9d, 0x53, 0xe6, 0x16, 0x2f, 0x0f, 0xab, 0x8b, 0x9b, 0x56, 0xeb, 0x78, 0x8c,
	0xed, 0x91, 0x11, 0xc5, 0xb2, 0xb7, 0x9a, 0x70, 0xd5, 0xb2, 0x2f, 0x3a, 0x9d, 0x9f, 0xc0, 0xdc,
	0x69, 0xd5, 0x9c, 0x0b, 0xa8, 0xc6, 0x27, 0x3e, 0x57, 0x50, 0x4d, 0x69, 0x0a, 0xad, 0x1a, 0x29,
	0x21, 0x81, 0xe6, 0xdf, 0x6f, 0xef, 0xbb, 0x78, 0xa8, 0x59, 0x35, 0x55, 0xf2, 0x9c, 0xaa, 0x5d,
	0xf3, 0xd4, 0xc5, 0xa3, 0x9e, 0x26, 0x90, 0x80, 0xea, 0x2f, 0xb0, 0xf2, 0xc5, 0x93, 0xbe, 0xe7,
	0xd0, 0x05, 0x3f, 0x86, 0xfa, 0xc3, 0xc7, 0xc3, 0x77, 0x9f, 0xd6, 0x18, 0x2c, 0x28, 0x93, 0x4d,
	0x87, 0xb7, 0xe5, 0xe7, 0xe5, 0xf6, 0x65, 0x9e, 0xde, 0xd7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x63, 0x9d, 0x8e, 0x9d, 0x01, 0x00, 0x00,
}
