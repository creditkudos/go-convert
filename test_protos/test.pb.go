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

func init() {
	proto.RegisterType((*Master)(nil), "test_protos.Master")
	proto.RegisterType((*Minion)(nil), "test_protos.Minion")
	proto.RegisterType((*Child)(nil), "test_protos.Child")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xbd, 0x6e, 0xea, 0x40,
	0x10, 0x05, 0x60, 0x9d, 0x35, 0xf6, 0x85, 0xe1, 0x92, 0x62, 0x53, 0x64, 0xca, 0x89, 0xab, 0xad,
	0x5c, 0x24, 0x6f, 0x90, 0xd4, 0x34, 0x53, 0xa6, 0x89, 0xfc, 0x07, 0x36, 0x8a, 0x31, 0xc2, 0x7e,
	0xbe, 0x3c, 0x5b, 0x34, 0x8b, 0x88, 0x48, 0x41, 0xb5, 0xf3, 0x49, 0x47, 0x3b, 0x7b, 0x96, 0x68,
	0x6e, 0xa7, 0xb9, 0x38, 0x9d, 0xc7, 0x79, 0xf4, 0x6b, 0x9b, 0x3f, 0xe3, 0x3c, 0xe5, 0xdf, 0xa0,
	0x6c, 0x5b, 0x4e, 0x73, 0x7b, 0xf6, 0xff, 0x09, 0x25, 0x43, 0x10, 0x56, 0x8a, 0xd2, 0x54, 0xb1,
	0x13, 0x84, 0x54, 0x51, 0x99, 0x6a, 0x4e, 0x04, 0x21, 0x51, 0xd4, 0xa6, 0x86, 0x17, 0x82, 0xb0,
	0x51, 0x34, 0xa6, 0x96, 0x53, 0x41, 0x58, 0x28, 0x5a, 0xd3, 0x8e, 0x33, 0x41, 0x70, 0x8a, 0x9d,
	0x69, 0xcf, 0xff, 0x04, 0x01, 0x8a, 0xbd, 0xa9, 0xe3, 0xa5, 0x20, 0x2c, 0x15, 0x9d, 0x7f, 0x26,
	0xf4, 0xbc, 0x92, 0x24, 0xac, 0x5f, 0x1e, 0x8b, 0x9b, 0x37, 0x15, 0xdb, 0xfe, 0xd8, 0x8f, 0x47,
	0x45, 0x6f, 0x91, 0x03, 0x93, 0xe0, 0x6e, 0xe4, 0x90, 0x9f, 0x28, 0xbb, 0xc0, 0x3f, 0x90, 0x1b,
	0xae, 0x05, 0xdc, 0x50, 0x46, 0x5f, 0x2b, 0xb8, 0xa1, 0xf2, 0x39, 0xb9, 0xc1, 0x4a, 0xd8, 0x42,
	0xff, 0xe7, 0xb6, 0xf7, 0xae, 0xff, 0x6a, 0xd4, 0x0d, 0x75, 0xcc, 0x5c, 0xaa, 0xdd, 0xcb, 0x34,
	0xf9, 0x13, 0xa5, 0x11, 0xb6, 0xa0, 0xfe, 0x5d, 0x58, 0x97, 0x6f, 0x9b, 0x8f, 0xdb, 0xaf, 0xad,
	0xb2, 0x78, 0xbe, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x57, 0x5f, 0x14, 0x8a, 0x7c, 0x01, 0x00,
	0x00,
}
