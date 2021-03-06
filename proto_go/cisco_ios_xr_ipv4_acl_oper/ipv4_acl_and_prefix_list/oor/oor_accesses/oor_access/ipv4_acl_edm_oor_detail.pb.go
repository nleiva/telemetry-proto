// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_acl_edm_oor_detail.proto

package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_oor_oor_accesses_oor_access

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

// Oor deatil config BAG
type Ipv4AclEdmOorDetail_KEYS struct {
	AccessListName       string   `protobuf:"bytes,1,opt,name=access_list_name,json=accessListName,proto3" json:"access_list_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclEdmOorDetail_KEYS) Reset()         { *m = Ipv4AclEdmOorDetail_KEYS{} }
func (m *Ipv4AclEdmOorDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail_KEYS) ProtoMessage()    {}
func (*Ipv4AclEdmOorDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_oor_detail_a653f72ed624921f, []int{0}
}
func (m *Ipv4AclEdmOorDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS.Unmarshal(m, b)
}
func (m *Ipv4AclEdmOorDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclEdmOorDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS.Merge(dst, src)
}
func (m *Ipv4AclEdmOorDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS.Size(m)
}
func (m *Ipv4AclEdmOorDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmOorDetail_KEYS proto.InternalMessageInfo

func (m *Ipv4AclEdmOorDetail_KEYS) GetAccessListName() string {
	if m != nil {
		return m.AccessListName
	}
	return ""
}

type Ipv4AclEdmOorDetail struct {
	// default max configurable acls
	DefaultMaxAcLs uint32 `protobuf:"varint,50,opt,name=default_max_ac_ls,json=defaultMaxAcLs,proto3" json:"default_max_ac_ls,omitempty"`
	// default max configurable aces
	DefaultMaxAcEs uint32 `protobuf:"varint,51,opt,name=default_max_ac_es,json=defaultMaxAcEs,proto3" json:"default_max_ac_es,omitempty"`
	// Current configured acls
	CurrentConfiguredAcLs uint32 `protobuf:"varint,52,opt,name=current_configured_ac_ls,json=currentConfiguredAcLs,proto3" json:"current_configured_ac_ls,omitempty"`
	// Current configured aces
	CurrentConfiguredAcEs uint32 `protobuf:"varint,53,opt,name=current_configured_ac_es,json=currentConfiguredAcEs,proto3" json:"current_configured_ac_es,omitempty"`
	// Current max configurable acls
	CurrentMaxConfigurableAcLs uint32 `protobuf:"varint,54,opt,name=current_max_configurable_ac_ls,json=currentMaxConfigurableAcLs,proto3" json:"current_max_configurable_ac_ls,omitempty"`
	// Current max configurable aces
	CurrentMaxConfigurableAcEs uint32 `protobuf:"varint,55,opt,name=current_max_configurable_ac_es,json=currentMaxConfigurableAcEs,proto3" json:"current_max_configurable_ac_es,omitempty"`
	// max configurable acls
	MaxConfigurableAcLs uint32 `protobuf:"varint,56,opt,name=max_configurable_ac_ls,json=maxConfigurableAcLs,proto3" json:"max_configurable_ac_ls,omitempty"`
	// max configurable aces
	MaxConfigurableAcEs  uint32   `protobuf:"varint,57,opt,name=max_configurable_ac_es,json=maxConfigurableAcEs,proto3" json:"max_configurable_ac_es,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclEdmOorDetail) Reset()         { *m = Ipv4AclEdmOorDetail{} }
func (m *Ipv4AclEdmOorDetail) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail) ProtoMessage()    {}
func (*Ipv4AclEdmOorDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_oor_detail_a653f72ed624921f, []int{1}
}
func (m *Ipv4AclEdmOorDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmOorDetail.Unmarshal(m, b)
}
func (m *Ipv4AclEdmOorDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmOorDetail.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclEdmOorDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmOorDetail.Merge(dst, src)
}
func (m *Ipv4AclEdmOorDetail) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmOorDetail.Size(m)
}
func (m *Ipv4AclEdmOorDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmOorDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmOorDetail proto.InternalMessageInfo

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcLs() uint32 {
	if m != nil {
		return m.DefaultMaxAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcEs() uint32 {
	if m != nil {
		return m.DefaultMaxAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcLs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcEs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcEs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4AclEdmOorDetail_KEYS)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.oor_accesses.oor_access.ipv4_acl_edm_oor_detail_KEYS")
	proto.RegisterType((*Ipv4AclEdmOorDetail)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.oor_accesses.oor_access.ipv4_acl_edm_oor_detail")
}

func init() {
	proto.RegisterFile("ipv4_acl_edm_oor_detail.proto", fileDescriptor_ipv4_acl_edm_oor_detail_a653f72ed624921f)
}

var fileDescriptor_ipv4_acl_edm_oor_detail_a653f72ed624921f = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0x29, 0x1f, 0x7c, 0x60, 0xc0, 0xa2, 0x2b, 0x6a, 0x10, 0x95, 0xd2, 0x53, 0xbd, 0xf4,
	0x60, 0xab, 0xd5, 0xa3, 0x96, 0x05, 0xc1, 0x56, 0xa1, 0x9e, 0x3c, 0x0d, 0xd3, 0xec, 0x54, 0x02,
	0xd9, 0xcd, 0x92, 0x49, 0x65, 0xff, 0xa7, 0x7f, 0x48, 0x9a, 0x6e, 0x97, 0x82, 0x8d, 0x1e, 0x72,
	0x98, 0xcc, 0xfb, 0xbc, 0x4f, 0x0e, 0x11, 0x17, 0xba, 0xfc, 0x1c, 0x02, 0x2a, 0x03, 0x94, 0xe5,
	0x60, 0xad, 0x83, 0x8c, 0x3c, 0x6a, 0xd3, 0x2f, 0x9d, 0xf5, 0x36, 0x79, 0x55, 0x9a, 0x95, 0x05,
	0x6d, 0x19, 0x2a, 0x07, 0x4d, 0xd6, 0x96, 0xe4, 0xfa, 0xcd, 0x84, 0x45, 0x06, 0xa5, 0xa3, 0x85,
	0xae, 0xc0, 0x68, 0xf6, 0x7d, 0x6b, 0xdd, 0xea, 0x00, 0x2a, 0x45, 0xcc, 0xc4, 0x5b, 0x43, 0xf7,
	0x49, 0x9c, 0x47, 0x8c, 0xf0, 0x9c, 0xbe, 0xbf, 0x25, 0x3d, 0x71, 0xb0, 0x4e, 0x86, 0x2a, 0x28,
	0x30, 0x27, 0xd9, 0xea, 0xb4, 0x7a, 0x7b, 0xb3, 0xf6, 0xfa, 0x7e, 0xa2, 0xd9, 0xbf, 0x60, 0x4e,
	0xdd, 0xaf, 0x7f, 0xe2, 0x34, 0x52, 0x95, 0x5c, 0x89, 0xc3, 0x8c, 0x16, 0xb8, 0x34, 0x1e, 0x72,
	0xac, 0x00, 0x15, 0x18, 0x96, 0xd7, 0x9d, 0x56, 0x6f, 0x7f, 0xd6, 0xae, 0x17, 0x53, 0xac, 0x1e,
	0xd4, 0x84, 0x77, 0x44, 0x89, 0xe5, 0xe0, 0x67, 0x34, 0xe5, 0x64, 0x24, 0xa4, 0x5a, 0x3a, 0x47,
	0x85, 0x07, 0x65, 0x8b, 0x85, 0xfe, 0x58, 0x3a, 0xca, 0xea, 0xf2, 0x61, 0x20, 0x8e, 0xeb, 0xfd,
	0xb8, 0x59, 0x07, 0x47, 0x14, 0x24, 0x96, 0x37, 0x51, 0x30, 0xe5, 0xe4, 0x51, 0x5c, 0x6e, 0xc0,
	0xd5, 0xe3, 0x36, 0x30, 0xce, 0x0d, 0xd5, 0xde, 0xdb, 0x80, 0x9f, 0xd5, 0xa9, 0x29, 0x56, 0xe3,
	0xad, 0x4c, 0x90, 0xff, 0xd1, 0x41, 0x2c, 0x47, 0xbf, 0x77, 0xa4, 0x9c, 0x0c, 0xc4, 0x49, 0xc4,
	0x7f, 0x17, 0xd8, 0xa3, 0x7c, 0x87, 0x38, 0x02, 0x11, 0xcb, 0xfb, 0x08, 0x94, 0xf2, 0xfc, 0x7f,
	0xf8, 0x77, 0x83, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x6d, 0x66, 0x22, 0x98, 0x02, 0x00,
	0x00,
}
