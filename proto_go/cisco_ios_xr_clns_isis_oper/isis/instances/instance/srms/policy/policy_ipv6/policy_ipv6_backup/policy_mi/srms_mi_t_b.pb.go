// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srms_mi_t_b.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_srms_policy_policy_ipv6_policy_ipv6_backup_policy_mi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SRMS show bag
type SrmsMiTB_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	MiId                 string   `protobuf:"bytes,2,opt,name=mi_id,json=miId,proto3" json:"mi_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB_KEYS) Reset()         { *m = SrmsMiTB_KEYS{} }
func (m *SrmsMiTB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB_KEYS) ProtoMessage()    {}
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb, []int{0}
}
func (m *SrmsMiTB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB_KEYS.Unmarshal(m, b)
}
func (m *SrmsMiTB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB_KEYS.Marshal(b, m, deterministic)
}
func (dst *SrmsMiTB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB_KEYS.Merge(dst, src)
}
func (m *SrmsMiTB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB_KEYS.Size(m)
}
func (m *SrmsMiTB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB_KEYS proto.InternalMessageInfo

func (m *SrmsMiTB_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
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
	return fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb, []int{1}
}
func (m *SrmsMiTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB.Unmarshal(m, b)
}
func (m *SrmsMiTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB.Marshal(b, m, deterministic)
}
func (dst *SrmsMiTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB.Merge(dst, src)
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
	return fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb, []int{2}
}
func (m *In6AddrTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In6AddrTB.Unmarshal(m, b)
}
func (m *In6AddrTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In6AddrTB.Marshal(b, m, deterministic)
}
func (dst *In6AddrTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In6AddrTB.Merge(dst, src)
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
	return fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb, []int{3}
}
func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (dst *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(dst, src)
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
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv6.policy_ipv6_backup.policy_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv6.policy_ipv6_backup.policy_mi.srms_mi_t_b")
	proto.RegisterType((*In6AddrTB)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv6.policy_ipv6_backup.policy_mi.in6_addr_t_b")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv6.policy_ipv6_backup.policy_mi.addr")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb) }

var fileDescriptor_srms_mi_t_b_4bdeda4fd3ded3bb = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xbd, 0x8e, 0x13, 0x31,
	0x10, 0xd6, 0x26, 0xb9, 0x13, 0x37, 0xc9, 0x9d, 0x0e, 0x83, 0x90, 0x25, 0x0a, 0xa2, 0x70, 0x45,
	0xaa, 0x2d, 0xee, 0x8e, 0xe5, 0xaf, 0x42, 0x88, 0xe2, 0x04, 0x42, 0x28, 0xa9, 0xa8, 0x46, 0x8e,
	0xed, 0x70, 0x23, 0x76, 0xed, 0x95, 0xed, 0x44, 0xa1, 0xa5, 0xa4, 0xe0, 0x75, 0x78, 0x3d, 0x34,
	0xce, 0x2e, 0xe4, 0x09, 0x52, 0xed, 0xf7, 0x63, 0x7d, 0xe3, 0xf1, 0xa7, 0x85, 0x87, 0x31, 0x34,
	0x11, 0x1b, 0xc2, 0x84, 0xab, 0xb2, 0x0d, 0x3e, 0x79, 0x71, 0xaf, 0x29, 0x6a, 0x8f, 0xe4, 0x23,
	0xee, 0x02, 0xea, 0xda, 0x45, 0xa4, 0x48, 0x11, 0x7d, 0x6b, 0x43, 0xc9, 0xa8, 0x24, 0x17, 0x93,
	0x72, 0xda, 0xfe, 0x47, 0x25, 0xc7, 0x94, 0xad, 0xaf, 0x49, 0xff, 0xe8, 0x3e, 0x48, 0xed, 0xb6,
	0x3a, 0xc4, 0xb8, 0x52, 0xfa, 0xfb, 0xa6, 0xed, 0xa5, 0x86, 0x66, 0x9f, 0xe0, 0xf2, 0x60, 0x3c,
	0x7e, 0xfc, 0xf0, 0x75, 0x29, 0x9e, 0xc3, 0x79, 0x1f, 0x8a, 0x4e, 0x35, 0x56, 0x16, 0xd3, 0x62,
	0x7e, 0xb6, 0x98, 0xf4, 0xe2, 0x67, 0xd5, 0x58, 0xf1, 0x08, 0x4e, 0x1a, 0x42, 0x32, 0x72, 0x90,
	0xcd, 0x51, 0x43, 0x77, 0x66, 0xf6, 0x7b, 0x08, 0xe3, 0x83, 0x38, 0x71, 0x09, 0xc3, 0x18, 0xb4,
	0xbc, 0xce, 0x47, 0x18, 0x8a, 0x27, 0x70, 0x1a, 0xfc, 0x26, 0xd9, 0x20, 0x6f, 0xb2, 0xd8, 0x31,
	0x21, 0x60, 0xa4, 0x82, 0x55, 0xf2, 0x76, 0x9f, 0xc6, 0x58, 0xfc, 0x2c, 0x60, 0xa4, 0x8c, 0x09,
	0xf2, 0xc5, 0xb4, 0x98, 0x8f, 0xaf, 0x5d, 0x79, 0xac, 0x57, 0x29, 0x79, 0xea, 0x22, 0xcf, 0xe6,
	0x0b, 0xb7, 0xc1, 0xae, 0x69, 0x27, 0xab, 0x69, 0x31, 0x3f, 0x5f, 0x74, 0x4c, 0x3c, 0x85, 0xb3,
	0x48, 0x06, 0x63, 0x52, 0x21, 0xc9, 0x97, 0xd9, 0x7a, 0x10, 0xc9, 0x2c, 0x99, 0xf7, 0xa6, 0xf6,
	0x1b, 0x97, 0xe4, 0xab, 0x7f, 0xe6, 0x7b, 0xe6, 0xe2, 0x19, 0x8c, 0x6b, 0x15, 0x13, 0x76, 0xb1,
	0xaf, 0xf3, 0xc6, 0xc0, 0xd2, 0x97, 0x7d, 0xf4, 0x15, 0x5c, 0xe4, 0x03, 0x1c, 0x41, 0xce, 0xd8,
	0x9d, 0x7c, 0x93, 0x23, 0x26, 0xac, 0x2e, 0xc9, 0xdc, 0xb1, 0xc6, 0x2d, 0xad, 0x6b, 0xf5, 0x0d,
	0x55, 0x4a, 0x4a, 0xdf, 0x5b, 0x23, 0xdf, 0xee, 0x5b, 0x62, 0xf1, 0x5d, 0xa7, 0xcd, 0xae, 0x60,
	0x42, 0xae, 0x42, 0xde, 0x24, 0x17, 0xf2, 0x18, 0x4e, 0xb6, 0xaa, 0xde, 0xf4, 0x95, 0xee, 0xc9,
	0xec, 0x4f, 0xf7, 0xd0, 0xe2, 0x02, 0x06, 0x6a, 0xdd, 0x79, 0x03, 0xb5, 0xe6, 0x56, 0xa8, 0xdd,
	0xde, 0xf6, 0x1d, 0x33, 0x16, 0xbf, 0x8a, 0x2c, 0x56, 0x72, 0x98, 0x5b, 0xd9, 0x1e, 0xaf, 0x95,
	0xc3, 0x4d, 0xf2, 0x65, 0xaa, 0xd5, 0x69, 0xfe, 0x5f, 0x6e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0xf7, 0xb2, 0xfc, 0x44, 0x03, 0x00, 0x00,
}
