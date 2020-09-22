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
	A                    string             `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32              `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    int64              `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D                    uint32             `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E                    uint64             `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F                    float32            `protobuf:"fixed32,6,opt,name=f,proto3" json:"f,omitempty"`
	G                    float64            `protobuf:"fixed64,7,opt,name=g,proto3" json:"g,omitempty"`
	H                    bool               `protobuf:"varint,8,opt,name=h,proto3" json:"h,omitempty"`
	I                    []*Minion          `protobuf:"bytes,9,rep,name=i,proto3" json:"i,omitempty"`
	J                    *Minion            `protobuf:"bytes,10,opt,name=j,proto3" json:"j,omitempty"`
	K                    []string           `protobuf:"bytes,11,rep,name=k,proto3" json:"k,omitempty"`
	L                    map[string]float32 `protobuf:"bytes,12,rep,name=l,proto3" json:"l,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	M                    map[string]*Minion `protobuf:"bytes,13,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *Master) GetK() []string {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *Master) GetL() map[string]float32 {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *Master) GetM() map[string]*Minion {
	if m != nil {
		return m.M
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
	proto.RegisterMapType((map[string]float32)(nil), "test_protos.Master.LEntry")
	proto.RegisterMapType((map[string]*Minion)(nil), "test_protos.Master.MEntry")
	proto.RegisterType((*Minion)(nil), "test_protos.Minion")
	proto.RegisterType((*Child)(nil), "test_protos.Child")
	proto.RegisterType((*RepeatedFailure)(nil), "test_protos.RepeatedFailure")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x4f, 0xeb, 0x30,
	0x10, 0xc6, 0x75, 0x4e, 0x93, 0xd7, 0x5e, 0xdb, 0xf7, 0x9e, 0x0c, 0x12, 0x56, 0x17, 0x4c, 0x26,
	0xb3, 0x64, 0x00, 0x06, 0xc4, 0x08, 0x02, 0x09, 0x89, 0x2c, 0x1e, 0x59, 0x90, 0x93, 0xb8, 0x6d,
	0xda, 0xa4, 0xad, 0x92, 0x14, 0xa9, 0x7f, 0x0e, 0xff, 0x29, 0x3a, 0xa7, 0x85, 0x76, 0xc8, 0x94,
	0xfb, 0x45, 0xdf, 0x7d, 0xf6, 0xdd, 0x67, 0xc4, 0xc6, 0xd6, 0x4d, 0xb4, 0xa9, 0xd6, 0xcd, 0x9a,
	0x0f, 0xa9, 0xfe, 0x70, 0x75, 0x1d, 0x7e, 0x79, 0x18, 0xc4, 0xa6, 0x6e, 0x6c, 0xc5, 0x47, 0x08,
	0x46, 0x80, 0x04, 0x35, 0xd0, 0x60, 0x88, 0x12, 0xc1, 0x24, 0x28, 0x5f, 0x43, 0x42, 0x94, 0x0a,
	0x4f, 0x82, 0xf2, 0x34, 0xa4, 0x44, 0x99, 0xe8, 0x49, 0x50, 0x63, 0x0d, 0x19, 0x91, 0x15, 0xbe,
	0x04, 0xd5, 0xd3, 0x60, 0x89, 0xa6, 0x22, 0x90, 0xa0, 0x98, 0x86, 0x29, 0xd1, 0x4c, 0xfc, 0x91,
	0xa0, 0x40, 0xc3, 0x8c, 0x68, 0x2e, 0xfa, 0x12, 0x54, 0x5f, 0xc3, 0x9c, 0x5f, 0x21, 0xe4, 0x62,
	0x20, 0x3d, 0x35, 0xbc, 0x39, 0x8b, 0x8e, 0xee, 0x14, 0xc5, 0xf9, 0x2a, 0x5f, 0xaf, 0x34, 0xe4,
	0x24, 0x59, 0x08, 0x94, 0xd0, 0x29, 0x59, 0x90, 0xe7, 0x52, 0x0c, 0xa5, 0x47, 0xb7, 0x5e, 0x72,
	0x85, 0x50, 0x88, 0x91, 0xf3, 0x9c, 0x9c, 0x36, 0xb8, 0x19, 0xa3, 0xb7, 0xe7, 0x55, 0x53, 0xed,
	0x34, 0x14, 0xa4, 0x2c, 0xc5, 0xb8, 0x5b, 0x19, 0xef, 0x95, 0xe5, 0xe4, 0x0e, 0x83, 0xb6, 0x8d,
	0xff, 0x47, 0x6f, 0x69, 0x77, 0xfb, 0x1d, 0x51, 0xc9, 0xcf, 0xd1, 0xff, 0x34, 0xc5, 0xd6, 0xba,
	0x4d, 0x31, 0xdd, 0xc2, 0x03, 0xbb, 0x87, 0xc9, 0x2b, 0x06, 0x71, 0x57, 0xd7, 0xf5, 0x71, 0x57,
	0xc7, 0x68, 0xbf, 0x56, 0xe1, 0x06, 0x83, 0xf6, 0x27, 0xff, 0x8b, 0xac, 0x3c, 0x64, 0xc4, 0x4a,
	0xe3, 0xf8, 0x90, 0x12, 0x2b, 0x13, 0x1e, 0x22, 0x2b, 0x29, 0x27, 0x9a, 0x8a, 0x9f, 0xb8, 0x3e,
	0xcd, 0xf3, 0x22, 0xd3, 0xac, 0x4c, 0x9d, 0xa6, 0x4d, 0xaf, 0x4b, 0x93, 0x85, 0x17, 0xe8, 0x3b,
	0xa0, 0x03, 0xd2, 0x9f, 0x03, 0x53, 0x13, 0x5e, 0xe2, 0x3f, 0x6d, 0x37, 0xd6, 0x34, 0x36, 0x7b,
	0x31, 0x79, 0xb1, 0xad, 0x5c, 0xe0, 0xb5, 0x80, 0x36, 0x80, 0xfa, 0x71, 0xfc, 0x7e, 0xfc, 0xbc,
	0x92, 0xc0, 0x7d, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x92, 0x2b, 0xd5, 0xec, 0x80, 0x02,
	0x00, 0x00,
}
