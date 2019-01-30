// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srms_mi_t_b.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_backup_policy_mi

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

// SRMS show bag
type SrmsMiTB_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	MiId                 string   `protobuf:"bytes,2,opt,name=mi_id,json=miId,proto3" json:"mi_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB_KEYS) Reset()         { *m = SrmsMiTB_KEYS{} }
func (m *SrmsMiTB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB_KEYS) ProtoMessage()    {}
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{0}
}

func (m *SrmsMiTB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB_KEYS.Unmarshal(m, b)
}
func (m *SrmsMiTB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB_KEYS.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB_KEYS.Merge(m, src)
}
func (m *SrmsMiTB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB_KEYS.Size(m)
}
func (m *SrmsMiTB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB_KEYS proto.InternalMessageInfo

func (m *SrmsMiTB_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *SrmsMiTB_KEYS) GetMiId() string {
	if m != nil {
		return m.MiId
	}
	return ""
}

type SrmsMiTB struct {
	Src string `protobuf:"bytes,50,opt,name=src,proto3" json:"src,omitempty"`
	// Router ID
	Router string `protobuf:"bytes,51,opt,name=router,proto3" json:"router,omitempty"`
	// Area (OSPF) or Level (ISIS)
	Area string `protobuf:"bytes,52,opt,name=area,proto3" json:"area,omitempty"`
	Addr *Addr  `protobuf:"bytes,53,opt,name=addr,proto3" json:"addr,omitempty"`
	// Prefix length
	Prefix uint32 `protobuf:"varint,54,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Starting SID
	SidStart uint32 `protobuf:"varint,55,opt,name=sid_start,json=sidStart,proto3" json:"sid_start,omitempty"`
	// SID range
	SidCount uint32 `protobuf:"varint,56,opt,name=sid_count,json=sidCount,proto3" json:"sid_count,omitempty"`
	// Last IP Prefix
	LastPrefix string `protobuf:"bytes,57,opt,name=last_prefix,json=lastPrefix,proto3" json:"last_prefix,omitempty"`
	// Last SID Index
	LastSidIndex uint32 `protobuf:"varint,58,opt,name=last_sid_index,json=lastSidIndex,proto3" json:"last_sid_index,omitempty"`
	// Attached flag
	FlagAttached         string   `protobuf:"bytes,59,opt,name=flag_attached,json=flagAttached,proto3" json:"flag_attached,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB) Reset()         { *m = SrmsMiTB{} }
func (m *SrmsMiTB) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB) ProtoMessage()    {}
func (*SrmsMiTB) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{1}
}

func (m *SrmsMiTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB.Unmarshal(m, b)
}
func (m *SrmsMiTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB.Merge(m, src)
}
func (m *SrmsMiTB) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB.Size(m)
}
func (m *SrmsMiTB) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB proto.InternalMessageInfo

func (m *SrmsMiTB) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *SrmsMiTB) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *SrmsMiTB) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *SrmsMiTB) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *SrmsMiTB) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SrmsMiTB) GetSidStart() uint32 {
	if m != nil {
		return m.SidStart
	}
	return 0
}

func (m *SrmsMiTB) GetSidCount() uint32 {
	if m != nil {
		return m.SidCount
	}
	return 0
}

func (m *SrmsMiTB) GetLastPrefix() string {
	if m != nil {
		return m.LastPrefix
	}
	return ""
}

func (m *SrmsMiTB) GetLastSidIndex() uint32 {
	if m != nil {
		return m.LastSidIndex
	}
	return 0
}

func (m *SrmsMiTB) GetFlagAttached() string {
	if m != nil {
		return m.FlagAttached
	}
	return ""
}

type In6AddrTB struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In6AddrTB) Reset()         { *m = In6AddrTB{} }
func (m *In6AddrTB) String() string { return proto.CompactTextString(m) }
func (*In6AddrTB) ProtoMessage()    {}
func (*In6AddrTB) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{2}
}

func (m *In6AddrTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In6AddrTB.Unmarshal(m, b)
}
func (m *In6AddrTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In6AddrTB.Marshal(b, m, deterministic)
}
func (m *In6AddrTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In6AddrTB.Merge(m, src)
}
func (m *In6AddrTB) XXX_Size() int {
	return xxx_messageInfo_In6AddrTB.Size(m)
}
func (m *In6AddrTB) XXX_DiscardUnknown() {
	xxx_messageInfo_In6AddrTB.DiscardUnknown(m)
}

var xxx_messageInfo_In6AddrTB proto.InternalMessageInfo

func (m *In6AddrTB) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Addr struct {
	Af string `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	// IPv4
	Ipv4 string `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// IPv6
	Ipv6                 *In6AddrTB `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{3}
}

func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *Addr) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Addr) GetIpv6() *In6AddrTB {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

func init() {
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b")
	proto.RegisterType((*In6AddrTB)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.in6_addr_t_b")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.addr")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor_bb21d79d36617d23) }

var fileDescriptor_bb21d79d36617d23 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x26, 0x69, 0x45, 0x27, 0x69, 0x55, 0x0c, 0x42, 0x96, 0x38, 0x10, 0x42, 0x0f, 0x39,
	0xed, 0xa1, 0x2d, 0xcb, 0xd7, 0x09, 0x21, 0x0e, 0x05, 0x09, 0xa1, 0xe4, 0xc4, 0x69, 0xe4, 0xd8,
	0xde, 0x62, 0xb1, 0x5e, 0xaf, 0x6c, 0x6f, 0x15, 0x8e, 0x1c, 0x39, 0xf0, 0x77, 0xf8, 0x7d, 0x68,
	0x1c, 0x2f, 0xda, 0x5f, 0xd0, 0x53, 0xde, 0xbc, 0x37, 0x79, 0xe3, 0x99, 0xa7, 0x85, 0x87, 0xc1,
	0xdb, 0x80, 0xd6, 0x60, 0xc4, 0x5d, 0xd9, 0x79, 0x17, 0x1d, 0xbb, 0x95, 0x26, 0x48, 0x87, 0xc6,
	0x05, 0xdc, 0x7b, 0x34, 0xdd, 0xdd, 0x35, 0xba, 0xd0, 0xd5, 0xe8, 0x3a, 0xed, 0x4b, 0x42, 0xd4,
	0x27, 0x75, 0x08, 0x3a, 0x0c, 0xa8, 0x24, 0x97, 0xb2, 0x73, 0x8d, 0x91, 0x3f, 0xf3, 0x4f, 0xfa,
	0xe3, 0x18, 0xe3, 0x4e, 0xc8, 0x1f, 0x7d, 0x37, 0x50, 0xd6, 0xac, 0x3e, 0xc1, 0xf9, 0x68, 0x3a,
	0x7e, 0xfe, 0xf8, 0x6d, 0xcb, 0x9e, 0xc3, 0x22, 0x7b, 0x62, 0x2b, 0xac, 0xe6, 0xc5, 0xb2, 0x58,
	0x9f, 0x6c, 0xe6, 0x99, 0xfb, 0x22, 0xac, 0x66, 0x8f, 0xe0, 0xc8, 0x1a, 0x34, 0x8a, 0x4f, 0x92,
	0x36, 0xb3, 0xe6, 0x46, 0xad, 0xfe, 0x4c, 0x61, 0x3e, 0x32, 0x63, 0xe7, 0x30, 0x0d, 0x5e, 0xf2,
	0xcb, 0xd4, 0x42, 0x90, 0x3d, 0x81, 0x63, 0xef, 0xfa, 0xa8, 0x3d, 0xbf, 0x4a, 0x64, 0xae, 0x18,
	0x83, 0x99, 0xf0, 0x5a, 0xf0, 0xeb, 0x83, 0x1b, 0x61, 0xf6, 0xab, 0x80, 0x99, 0x50, 0xca, 0xf3,
	0x97, 0xcb, 0x62, 0x3d, 0xbf, 0xb4, 0xe5, 0x3d, 0x9d, 0xa4, 0xa4, 0xa1, 0x9b, 0x34, 0x9a, 0xde,
	0xdb, 0x79, 0x5d, 0x9b, 0x3d, 0xaf, 0x96, 0xc5, 0xfa, 0x74, 0x93, 0x2b, 0xf6, 0x14, 0x4e, 0x82,
	0x51, 0x18, 0xa2, 0xf0, 0x91, 0xbf, 0x4a, 0xd2, 0x83, 0x60, 0xd4, 0x96, 0xea, 0x41, 0x94, 0xae,
	0x6f, 0x23, 0x7f, 0xfd, 0x5f, 0xfc, 0x40, 0x35, 0x7b, 0x06, 0xf3, 0x46, 0x84, 0x88, 0xd9, 0xf6,
	0x4d, 0x5a, 0x18, 0x88, 0xfa, 0x7a, 0xb0, 0xbe, 0x80, 0xb3, 0xd4, 0x40, 0x16, 0xa6, 0x55, 0x7a,
	0xcf, 0xdf, 0x26, 0x8b, 0x05, 0xb1, 0x5b, 0xa3, 0x6e, 0x88, 0x63, 0x2f, 0xe0, 0xb4, 0x6e, 0xc4,
	0x2d, 0x8a, 0x18, 0x85, 0xfc, 0xae, 0x15, 0x7f, 0x97, 0x8c, 0x16, 0x44, 0xbe, 0xcf, 0xdc, 0xea,
	0x02, 0x16, 0xa6, 0xad, 0x90, 0x36, 0x49, 0x79, 0x3c, 0x86, 0xa3, 0x3b, 0xd1, 0xf4, 0x43, 0xa0,
	0x87, 0x62, 0xf5, 0x37, 0xdf, 0x99, 0x9d, 0xc1, 0x44, 0xd4, 0x59, 0x9b, 0x88, 0x9a, 0x42, 0xa1,
	0x03, 0x0d, 0x11, 0x13, 0x66, 0xbf, 0x8b, 0x44, 0x56, 0x7c, 0x9a, 0x42, 0xe9, 0xef, 0x2d, 0x94,
	0xf1, 0x22, 0xe9, 0x2d, 0xd5, 0xee, 0x38, 0x7d, 0x2a, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x8b, 0xcb, 0xb0, 0x3f, 0x03, 0x00, 0x00,
}
