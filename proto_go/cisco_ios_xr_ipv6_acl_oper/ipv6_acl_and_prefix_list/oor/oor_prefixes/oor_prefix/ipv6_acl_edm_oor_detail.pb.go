// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_acl_edm_oor_detail.proto

package cisco_ios_xr_ipv6_acl_oper_ipv6_acl_and_prefix_list_oor_oor_prefixes_oor_prefix

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
type Ipv6AclEdmOorDetail_KEYS struct {
	PrefixListName       string   `protobuf:"bytes,1,opt,name=prefix_list_name,json=prefixListName,proto3" json:"prefix_list_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6AclEdmOorDetail_KEYS) Reset()         { *m = Ipv6AclEdmOorDetail_KEYS{} }
func (m *Ipv6AclEdmOorDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6AclEdmOorDetail_KEYS) ProtoMessage()    {}
func (*Ipv6AclEdmOorDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_acl_edm_oor_detail_95b7ceeb6a5a3153, []int{0}
}
func (m *Ipv6AclEdmOorDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS.Unmarshal(m, b)
}
func (m *Ipv6AclEdmOorDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv6AclEdmOorDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS.Merge(dst, src)
}
func (m *Ipv6AclEdmOorDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS.Size(m)
}
func (m *Ipv6AclEdmOorDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6AclEdmOorDetail_KEYS proto.InternalMessageInfo

func (m *Ipv6AclEdmOorDetail_KEYS) GetPrefixListName() string {
	if m != nil {
		return m.PrefixListName
	}
	return ""
}

type Ipv6AclEdmOorDetail struct {
	// default max configurable acls
	IsDefaultMaximumConfigurableAcLs uint32 `protobuf:"varint,50,opt,name=is_default_maximum_configurable_ac_ls,json=isDefaultMaximumConfigurableAcLs,proto3" json:"is_default_maximum_configurable_ac_ls,omitempty"`
	// default max configurable aces
	IsDefaultMaximumConfigurableAcEs uint32 `protobuf:"varint,51,opt,name=is_default_maximum_configurable_ac_es,json=isDefaultMaximumConfigurableAcEs,proto3" json:"is_default_maximum_configurable_ac_es,omitempty"`
	// Current configured acls
	IsCurrentConfiguredAcLs uint32 `protobuf:"varint,52,opt,name=is_current_configured_ac_ls,json=isCurrentConfiguredAcLs,proto3" json:"is_current_configured_ac_ls,omitempty"`
	// Current configured aces
	IsCurrentConfiguredAces uint32 `protobuf:"varint,53,opt,name=is_current_configured_aces,json=isCurrentConfiguredAces,proto3" json:"is_current_configured_aces,omitempty"`
	// Current max configurable acls
	IsCurrentMaximumConfigurableAcls uint32 `protobuf:"varint,54,opt,name=is_current_maximum_configurable_acls,json=isCurrentMaximumConfigurableAcls,proto3" json:"is_current_maximum_configurable_acls,omitempty"`
	// Current max configurable aces
	IsCurrentMaximumConfigurableAces uint32 `protobuf:"varint,55,opt,name=is_current_maximum_configurable_aces,json=isCurrentMaximumConfigurableAces,proto3" json:"is_current_maximum_configurable_aces,omitempty"`
	// max configurable acls
	IsMaximumConfigurableAcLs uint32 `protobuf:"varint,56,opt,name=is_maximum_configurable_ac_ls,json=isMaximumConfigurableAcLs,proto3" json:"is_maximum_configurable_ac_ls,omitempty"`
	// max configurable aces
	IsMaximumConfigurableAcEs uint32   `protobuf:"varint,57,opt,name=is_maximum_configurable_ac_es,json=isMaximumConfigurableAcEs,proto3" json:"is_maximum_configurable_ac_es,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Ipv6AclEdmOorDetail) Reset()         { *m = Ipv6AclEdmOorDetail{} }
func (m *Ipv6AclEdmOorDetail) String() string { return proto.CompactTextString(m) }
func (*Ipv6AclEdmOorDetail) ProtoMessage()    {}
func (*Ipv6AclEdmOorDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_acl_edm_oor_detail_95b7ceeb6a5a3153, []int{1}
}
func (m *Ipv6AclEdmOorDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6AclEdmOorDetail.Unmarshal(m, b)
}
func (m *Ipv6AclEdmOorDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6AclEdmOorDetail.Marshal(b, m, deterministic)
}
func (dst *Ipv6AclEdmOorDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6AclEdmOorDetail.Merge(dst, src)
}
func (m *Ipv6AclEdmOorDetail) XXX_Size() int {
	return xxx_messageInfo_Ipv6AclEdmOorDetail.Size(m)
}
func (m *Ipv6AclEdmOorDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6AclEdmOorDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6AclEdmOorDetail proto.InternalMessageInfo

func (m *Ipv6AclEdmOorDetail) GetIsDefaultMaximumConfigurableAcLs() uint32 {
	if m != nil {
		return m.IsDefaultMaximumConfigurableAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsDefaultMaximumConfigurableAcEs() uint32 {
	if m != nil {
		return m.IsDefaultMaximumConfigurableAcEs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentConfiguredAcLs() uint32 {
	if m != nil {
		return m.IsCurrentConfiguredAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentConfiguredAces() uint32 {
	if m != nil {
		return m.IsCurrentConfiguredAces
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentMaximumConfigurableAcls() uint32 {
	if m != nil {
		return m.IsCurrentMaximumConfigurableAcls
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentMaximumConfigurableAces() uint32 {
	if m != nil {
		return m.IsCurrentMaximumConfigurableAces
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsMaximumConfigurableAcLs() uint32 {
	if m != nil {
		return m.IsMaximumConfigurableAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsMaximumConfigurableAcEs() uint32 {
	if m != nil {
		return m.IsMaximumConfigurableAcEs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6AclEdmOorDetail_KEYS)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.oor.oor_prefixes.oor_prefix.ipv6_acl_edm_oor_detail_KEYS")
	proto.RegisterType((*Ipv6AclEdmOorDetail)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.oor.oor_prefixes.oor_prefix.ipv6_acl_edm_oor_detail")
}

func init() {
	proto.RegisterFile("ipv6_acl_edm_oor_detail.proto", fileDescriptor_ipv6_acl_edm_oor_detail_95b7ceeb6a5a3153)
}

var fileDescriptor_ipv6_acl_edm_oor_detail_95b7ceeb6a5a3153 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0x29, 0xbc, 0xbc, 0xe0, 0x82, 0x22, 0xb9, 0x74, 0xfd, 0x53, 0x28, 0x45, 0xa1, 0xa7,
	0x1e, 0xac, 0x56, 0x45, 0x0f, 0x4a, 0x2d, 0x08, 0xd6, 0x16, 0xea, 0xc9, 0xd3, 0xb0, 0xdd, 0x4c,
	0x65, 0x60, 0x37, 0x1b, 0x76, 0x36, 0xd2, 0xcf, 0xea, 0xa7, 0x11, 0x93, 0x1a, 0x72, 0x30, 0x21,
	0x87, 0x1c, 0x66, 0x78, 0xe6, 0xc9, 0xfc, 0x98, 0x15, 0x3d, 0x4a, 0x3f, 0x27, 0xa0, 0xb4, 0x01,
	0x8c, 0x2d, 0x38, 0xe7, 0x21, 0xc6, 0xa0, 0xc8, 0x8c, 0x52, 0xef, 0x82, 0x8b, 0x96, 0x9a, 0x58,
	0x3b, 0x20, 0xc7, 0xb0, 0xf5, 0x50, 0xb2, 0x2e, 0x45, 0x3f, 0x2a, 0x2b, 0x95, 0xc4, 0x90, 0x7a,
	0xdc, 0xd0, 0x16, 0x0c, 0x71, 0x18, 0x39, 0xe7, 0x7f, 0xbe, 0x5d, 0x0f, 0xb9, 0x52, 0x0c, 0x9e,
	0xc5, 0x69, 0xcd, 0x1f, 0xe1, 0x65, 0xf6, 0xfe, 0x16, 0x0d, 0xc5, 0x61, 0x45, 0x05, 0x89, 0xb2,
	0x28, 0x3b, 0xfd, 0xce, 0x70, 0x6f, 0x75, 0x50, 0xf4, 0xe7, 0xc4, 0x61, 0xa1, 0x2c, 0x0e, 0xbe,
	0xfe, 0x89, 0x6e, 0x8d, 0x2a, 0x5a, 0x8a, 0x73, 0x62, 0x88, 0x71, 0xa3, 0x32, 0x13, 0xc0, 0xaa,
	0x2d, 0xd9, 0xcc, 0x82, 0x76, 0xc9, 0x86, 0x3e, 0x32, 0xaf, 0xd6, 0x06, 0x41, 0x69, 0x30, 0x2c,
	0x2f, 0xfa, 0x9d, 0xe1, 0xfe, 0xaa, 0x4f, 0xfc, 0x54, 0xb0, 0xaf, 0x05, 0x3a, 0xad, 0x90, 0x8f,
	0x7a, 0xce, 0x2d, 0x85, 0xc8, 0x72, 0xdc, 0x46, 0x38, 0xe3, 0xe8, 0x5e, 0x9c, 0x10, 0x83, 0xce,
	0xbc, 0xc7, 0x24, 0x94, 0x22, 0x8c, 0x77, 0x7b, 0x5d, 0xe6, 0x9a, 0x2e, 0xf1, 0xb4, 0x20, 0xa6,
	0x25, 0x90, 0xaf, 0x73, 0x27, 0x8e, 0xeb, 0xa6, 0x91, 0xe5, 0x55, 0xc3, 0x30, 0x72, 0xb4, 0x10,
	0x67, 0x95, 0xe1, 0x9a, 0x2c, 0x86, 0xe5, 0xe4, 0x37, 0xca, 0x4e, 0xf3, 0x67, 0x14, 0xd3, 0xd2,
	0x87, 0x2c, 0xaf, 0xdb, 0xf8, 0x90, 0xa3, 0x07, 0xd1, 0x23, 0x6e, 0x3a, 0xda, 0x4d, 0x2e, 0x3a,
	0x22, 0xae, 0xbb, 0x56, 0xb3, 0x01, 0x59, 0xde, 0x36, 0x1a, 0x66, 0xbc, 0xfe, 0x9f, 0x3f, 0xff,
	0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x4f, 0x62, 0x2c, 0x1f, 0x03, 0x00, 0x00,
}
