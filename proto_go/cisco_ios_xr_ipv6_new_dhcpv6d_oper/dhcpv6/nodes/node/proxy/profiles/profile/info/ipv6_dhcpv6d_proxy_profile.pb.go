// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_dhcpv6d_proxy_profile.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_profiles_profile_info

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

// DHCPv6 proxy profile parameters
type Ipv6Dhcpv6DProxyProfile_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProfileName          string   `protobuf:"bytes,2,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyProfile_KEYS) Reset()         { *m = Ipv6Dhcpv6DProxyProfile_KEYS{} }
func (m *Ipv6Dhcpv6DProxyProfile_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyProfile_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyProfile_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{0}
}

func (m *Ipv6Dhcpv6DProxyProfile_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyProfile_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyProfile_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyProfile_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DProxyProfile_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyProfile_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyProfile_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyProfile_KEYS) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

type Ipv6Dhcpv6DProxyProfile struct {
	// Proxy profile name
	ProfileName string `protobuf:"bytes,50,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	// Helper addresses
	ProfileHelperAddress []*IPV6AddressType `protobuf:"bytes,51,rep,name=profile_helper_address,json=profileHelperAddress,proto3" json:"profile_helper_address,omitempty"`
	// VRF names
	VrfName []*StringVrf `protobuf:"bytes,52,rep,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	// Interface names
	InterfaceName []*StringIfname `protobuf:"bytes,53,rep,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Remote id
	RemoteId string `protobuf:"bytes,54,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	// Link address
	ProfileLinkAddress string `protobuf:"bytes,55,opt,name=profile_link_address,json=profileLinkAddress,proto3" json:"profile_link_address,omitempty"`
	// Interface id references
	InterfaceIdReferences *Ipv6Dhcpv6DProxyIidReference `protobuf:"bytes,56,opt,name=interface_id_references,json=interfaceIdReferences,proto3" json:"interface_id_references,omitempty"`
	// VRF references
	VrfReferences *Ipv6Dhcpv6DProxyVrfReference `protobuf:"bytes,57,opt,name=vrf_references,json=vrfReferences,proto3" json:"vrf_references,omitempty"`
	// Interface references
	InterfaceReferences  *Ipv6Dhcpv6DProxyInterfaceReference `protobuf:"bytes,58,opt,name=interface_references,json=interfaceReferences,proto3" json:"interface_references,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyProfile) Reset()         { *m = Ipv6Dhcpv6DProxyProfile{} }
func (m *Ipv6Dhcpv6DProxyProfile) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyProfile) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{1}
}

func (m *Ipv6Dhcpv6DProxyProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyProfile.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyProfile) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyProfile.Size(m)
}
func (m *Ipv6Dhcpv6DProxyProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyProfile.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyProfile proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyProfile) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyProfile) GetProfileHelperAddress() []*IPV6AddressType {
	if m != nil {
		return m.ProfileHelperAddress
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyProfile) GetVrfName() []*StringVrf {
	if m != nil {
		return m.VrfName
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyProfile) GetInterfaceName() []*StringIfname {
	if m != nil {
		return m.InterfaceName
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyProfile) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyProfile) GetProfileLinkAddress() string {
	if m != nil {
		return m.ProfileLinkAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyProfile) GetInterfaceIdReferences() *Ipv6Dhcpv6DProxyIidReference {
	if m != nil {
		return m.InterfaceIdReferences
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyProfile) GetVrfReferences() *Ipv6Dhcpv6DProxyVrfReference {
	if m != nil {
		return m.VrfReferences
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyProfile) GetInterfaceReferences() *Ipv6Dhcpv6DProxyInterfaceReference {
	if m != nil {
		return m.InterfaceReferences
	}
	return nil
}

// IPV6 Address type
type IPV6AddressType struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPV6AddressType) Reset()         { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()    {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{2}
}

func (m *IPV6AddressType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPV6AddressType.Unmarshal(m, b)
}
func (m *IPV6AddressType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPV6AddressType.Marshal(b, m, deterministic)
}
func (m *IPV6AddressType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPV6AddressType.Merge(m, src)
}
func (m *IPV6AddressType) XXX_Size() int {
	return xxx_messageInfo_IPV6AddressType.Size(m)
}
func (m *IPV6AddressType) XXX_DiscardUnknown() {
	xxx_messageInfo_IPV6AddressType.DiscardUnknown(m)
}

var xxx_messageInfo_IPV6AddressType proto.InternalMessageInfo

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringVrf struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringVrf) Reset()         { *m = StringVrf{} }
func (m *StringVrf) String() string { return proto.CompactTextString(m) }
func (*StringVrf) ProtoMessage()    {}
func (*StringVrf) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{3}
}

func (m *StringVrf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringVrf.Unmarshal(m, b)
}
func (m *StringVrf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringVrf.Marshal(b, m, deterministic)
}
func (m *StringVrf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringVrf.Merge(m, src)
}
func (m *StringVrf) XXX_Size() int {
	return xxx_messageInfo_StringVrf.Size(m)
}
func (m *StringVrf) XXX_DiscardUnknown() {
	xxx_messageInfo_StringVrf.DiscardUnknown(m)
}

var xxx_messageInfo_StringVrf proto.InternalMessageInfo

func (m *StringVrf) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringIfname struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringIfname) Reset()         { *m = StringIfname{} }
func (m *StringIfname) String() string { return proto.CompactTextString(m) }
func (*StringIfname) ProtoMessage()    {}
func (*StringIfname) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{4}
}

func (m *StringIfname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringIfname.Unmarshal(m, b)
}
func (m *StringIfname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringIfname.Marshal(b, m, deterministic)
}
func (m *StringIfname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringIfname.Merge(m, src)
}
func (m *StringIfname) XXX_Size() int {
	return xxx_messageInfo_StringIfname.Size(m)
}
func (m *StringIfname) XXX_DiscardUnknown() {
	xxx_messageInfo_StringIfname.DiscardUnknown(m)
}

var xxx_messageInfo_StringIfname proto.InternalMessageInfo

func (m *StringIfname) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// DHCPv6 proxy profile reference
type Ipv6Dhcpv6DProxyVrfReference struct {
	// Next vrf reference information
	Ipv6Dhcpv6DProxyVrfReference []*Ipv6Dhcpv6DProxyVrfReferenceItem `protobuf:"bytes,1,rep,name=ipv6_dhcpv6d_proxy_vrf_reference,json=ipv6Dhcpv6dProxyVrfReference,proto3" json:"ipv6_dhcpv6d_proxy_vrf_reference,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                            `json:"-"`
	XXX_unrecognized             []byte                              `json:"-"`
	XXX_sizecache                int32                               `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyVrfReference) Reset()         { *m = Ipv6Dhcpv6DProxyVrfReference{} }
func (m *Ipv6Dhcpv6DProxyVrfReference) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyVrfReference) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyVrfReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{5}
}

func (m *Ipv6Dhcpv6DProxyVrfReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyVrfReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyVrfReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyVrfReference) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference.Size(m)
}
func (m *Ipv6Dhcpv6DProxyVrfReference) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReference proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyVrfReference) GetIpv6Dhcpv6DProxyVrfReference() []*Ipv6Dhcpv6DProxyVrfReferenceItem {
	if m != nil {
		return m.Ipv6Dhcpv6DProxyVrfReference
	}
	return nil
}

type Ipv6Dhcpv6DProxyVrfReferenceItem struct {
	// VRF name
	ProxyReferenceVrfName string   `protobuf:"bytes,1,opt,name=proxy_reference_vrf_name,json=proxyReferenceVrfName,proto3" json:"proxy_reference_vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) Reset()         { *m = Ipv6Dhcpv6DProxyVrfReferenceItem{} }
func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyVrfReferenceItem) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyVrfReferenceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{6}
}

func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem.Size(m)
}
func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyVrfReferenceItem proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyVrfReferenceItem) GetProxyReferenceVrfName() string {
	if m != nil {
		return m.ProxyReferenceVrfName
	}
	return ""
}

// DHCPv6 proxy profile reference
type Ipv6Dhcpv6DProxyInterfaceReference struct {
	// Next interface reference information
	Ipv6Dhcpv6DProxyInterfaceReference []*Ipv6Dhcpv6DProxyInterfaceReferenceItem `protobuf:"bytes,1,rep,name=ipv6_dhcpv6d_proxy_interface_reference,json=ipv6Dhcpv6dProxyInterfaceReference,proto3" json:"ipv6_dhcpv6d_proxy_interface_reference,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}                                  `json:"-"`
	XXX_unrecognized                   []byte                                    `json:"-"`
	XXX_sizecache                      int32                                     `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyInterfaceReference) Reset()         { *m = Ipv6Dhcpv6DProxyInterfaceReference{} }
func (m *Ipv6Dhcpv6DProxyInterfaceReference) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyInterfaceReference) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyInterfaceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{7}
}

func (m *Ipv6Dhcpv6DProxyInterfaceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReference) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference.Size(m)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReference) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReference proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyInterfaceReference) GetIpv6Dhcpv6DProxyInterfaceReference() []*Ipv6Dhcpv6DProxyInterfaceReferenceItem {
	if m != nil {
		return m.Ipv6Dhcpv6DProxyInterfaceReference
	}
	return nil
}

type Ipv6Dhcpv6DProxyInterfaceReferenceItem struct {
	// Interface name
	ProxyReferenceInterfaceName string   `protobuf:"bytes,1,opt,name=proxy_reference_interface_name,json=proxyReferenceInterfaceName,proto3" json:"proxy_reference_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) Reset() {
	*m = Ipv6Dhcpv6DProxyInterfaceReferenceItem{}
}
func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyInterfaceReferenceItem) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyInterfaceReferenceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{8}
}

func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem.Size(m)
}
func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyInterfaceReferenceItem proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyInterfaceReferenceItem) GetProxyReferenceInterfaceName() string {
	if m != nil {
		return m.ProxyReferenceInterfaceName
	}
	return ""
}

// DHCPv6 proxy profile iid reference
type Ipv6Dhcpv6DProxyIidReference struct {
	// Next interface iid reference information
	Ipv6Dhcpv6DProxyIidReference []*Ipv6Dhcpv6DProxyIidReferenceItem `protobuf:"bytes,1,rep,name=ipv6_dhcpv6d_proxy_iid_reference,json=ipv6Dhcpv6dProxyIidReference,proto3" json:"ipv6_dhcpv6d_proxy_iid_reference,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                            `json:"-"`
	XXX_unrecognized             []byte                              `json:"-"`
	XXX_sizecache                int32                               `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyIidReference) Reset()         { *m = Ipv6Dhcpv6DProxyIidReference{} }
func (m *Ipv6Dhcpv6DProxyIidReference) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyIidReference) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyIidReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{9}
}

func (m *Ipv6Dhcpv6DProxyIidReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyIidReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyIidReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyIidReference) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference.Size(m)
}
func (m *Ipv6Dhcpv6DProxyIidReference) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyIidReference proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyIidReference) GetIpv6Dhcpv6DProxyIidReference() []*Ipv6Dhcpv6DProxyIidReferenceItem {
	if m != nil {
		return m.Ipv6Dhcpv6DProxyIidReference
	}
	return nil
}

type Ipv6Dhcpv6DProxyIidReferenceItem struct {
	// Interface name for interface id
	ProxyIidInterfaceName string `protobuf:"bytes,1,opt,name=proxy_iid_interface_name,json=proxyIidInterfaceName,proto3" json:"proxy_iid_interface_name,omitempty"`
	// Interface id
	ProxyInterfaceId     string   `protobuf:"bytes,2,opt,name=proxy_interface_id,json=proxyInterfaceId,proto3" json:"proxy_interface_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyIidReferenceItem) Reset()         { *m = Ipv6Dhcpv6DProxyIidReferenceItem{} }
func (m *Ipv6Dhcpv6DProxyIidReferenceItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyIidReferenceItem) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyIidReferenceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_78aa0d2748cc147d, []int{10}
}

func (m *Ipv6Dhcpv6DProxyIidReferenceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyIidReferenceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyIidReferenceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyIidReferenceItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem.Size(m)
}
func (m *Ipv6Dhcpv6DProxyIidReferenceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyIidReferenceItem proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyIidReferenceItem) GetProxyIidInterfaceName() string {
	if m != nil {
		return m.ProxyIidInterfaceName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyIidReferenceItem) GetProxyInterfaceId() string {
	if m != nil {
		return m.ProxyInterfaceId
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DProxyProfile_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_profile_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DProxyProfile)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_profile")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.IPV6AddressType")
	proto.RegisterType((*StringVrf)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.string_vrf")
	proto.RegisterType((*StringIfname)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.string_ifname")
	proto.RegisterType((*Ipv6Dhcpv6DProxyVrfReference)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_vrf_reference")
	proto.RegisterType((*Ipv6Dhcpv6DProxyVrfReferenceItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_vrf_reference_item")
	proto.RegisterType((*Ipv6Dhcpv6DProxyInterfaceReference)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_interface_reference")
	proto.RegisterType((*Ipv6Dhcpv6DProxyInterfaceReferenceItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_interface_reference_item")
	proto.RegisterType((*Ipv6Dhcpv6DProxyIidReference)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_iid_reference")
	proto.RegisterType((*Ipv6Dhcpv6DProxyIidReferenceItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.profiles.profile.info.ipv6_dhcpv6d_proxy_iid_reference_item")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_proxy_profile.proto", fileDescriptor_78aa0d2748cc147d) }

var fileDescriptor_78aa0d2748cc147d = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x65, 0x14, 0xf5, 0xf5, 0xd6, 0x3e, 0x65, 0xec, 0xd3, 0x60, 0x45, 0x6b, 0xe0, 0x69, 0x41,
	0x09, 0xd2, 0xa7, 0x7d, 0xea, 0x4e, 0x54, 0x30, 0x28, 0x52, 0xa2, 0x14, 0x04, 0x61, 0x8c, 0x9d,
	0x89, 0x6f, 0x78, 0x6d, 0x12, 0x26, 0x35, 0xed, 0x5b, 0xb8, 0x75, 0xe9, 0x56, 0x5c, 0xba, 0xd5,
	0x1f, 0xe1, 0x6f, 0xf0, 0x3f, 0xf8, 0x33, 0x5c, 0xc8, 0xcc, 0xe4, 0xa3, 0x69, 0x1a, 0xa8, 0x50,
	0xeb, 0x26, 0x34, 0x73, 0x4f, 0xef, 0x39, 0x39, 0xf7, 0xe4, 0x12, 0x68, 0xf3, 0x30, 0xee, 0x11,
	0x7a, 0x30, 0x0c, 0xe3, 0x1e, 0x25, 0xa1, 0x08, 0x66, 0x47, 0xf2, 0xea, 0xf1, 0x11, 0xb3, 0x42,
	0x11, 0x4c, 0x02, 0xdc, 0x1f, 0xf2, 0x68, 0x18, 0x10, 0x1e, 0x44, 0x64, 0x26, 0x88, 0x82, 0xfb,
	0x6c, 0x9a, 0xfd, 0x25, 0x08, 0x99, 0xb0, 0xf4, 0x8d, 0xe5, 0x07, 0x94, 0x45, 0xea, 0x6a, 0xa9,
	0x4e, 0x56, 0xd2, 0x29, 0x4a, 0x7f, 0x58, 0xdc, 0xf7, 0x02, 0xd3, 0x85, 0x2b, 0xd5, 0xac, 0xe4,
	0xe9, 0xe3, 0x57, 0x2f, 0x70, 0x0b, 0x6a, 0xb2, 0x0d, 0xf1, 0xdd, 0x31, 0x33, 0x50, 0x1b, 0x75,
	0x6a, 0xce, 0x96, 0x3c, 0x78, 0xee, 0x8e, 0x19, 0xbe, 0x0a, 0xa7, 0x53, 0xb0, 0xaa, 0x1f, 0x53,
	0xf5, 0x7a, 0x72, 0x26, 0x21, 0xe6, 0xd7, 0x2d, 0xb8, 0x58, 0xcd, 0x51, 0xea, 0xd0, 0x2d, 0x75,
	0xc0, 0x9f, 0x11, 0x9c, 0x4f, 0x31, 0x07, 0x6c, 0x14, 0x32, 0x41, 0x5c, 0x4a, 0x05, 0x8b, 0x22,
	0x63, 0xaf, 0x7d, 0xbc, 0x53, 0xef, 0xba, 0xd6, 0xba, 0x8d, 0xb1, 0xec, 0xfe, 0xa0, 0xf7, 0x40,
	0x93, 0xbc, 0x3c, 0x0a, 0x99, 0xd3, 0x4c, 0xaa, 0x4f, 0x14, 0x7f, 0x52, 0xc1, 0x53, 0xd8, 0x8a,
	0x85, 0xa7, 0x85, 0xdf, 0x56, 0x52, 0x5e, 0xaf, 0x5f, 0x4a, 0x34, 0x11, 0xdc, 0x7f, 0x47, 0x62,
	0xe1, 0x39, 0xa7, 0x62, 0xe1, 0x29, 0x4b, 0x3e, 0x22, 0xd8, 0xe6, 0xfe, 0x84, 0x09, 0xcf, 0x1d,
	0x26, 0xc6, 0xdd, 0x51, 0xfc, 0xe4, 0x9f, 0xf1, 0x73, 0x4f, 0xd2, 0x38, 0x8d, 0x8c, 0x56, 0x09,
	0x69, 0x41, 0x4d, 0xb0, 0x71, 0x30, 0x61, 0x84, 0x53, 0xa3, 0xa7, 0xd3, 0xa1, 0x0f, 0x6c, 0x8a,
	0x6f, 0x41, 0x6a, 0x1b, 0x19, 0x71, 0xff, 0x30, 0x9b, 0xda, 0xbe, 0xc2, 0xe1, 0xa4, 0xf6, 0x8c,
	0xfb, 0x87, 0xa9, 0xa1, 0xdf, 0x10, 0x5c, 0xc8, 0x9f, 0x8b, 0x53, 0x22, 0x98, 0xc7, 0x04, 0xf3,
	0x87, 0x2c, 0x32, 0xee, 0xb6, 0x51, 0xa7, 0xde, 0x15, 0xeb, 0x7f, 0xc0, 0x25, 0xe9, 0xe4, 0xf3,
	0xd4, 0xce, 0x4e, 0x26, 0xc9, 0xa6, 0x4e, 0x26, 0x08, 0x7f, 0x41, 0xb0, 0x2d, 0xc7, 0x3f, 0xa7,
	0xf1, 0xde, 0x06, 0x35, 0x16, 0xa8, 0x9d, 0x86, 0xcc, 0x47, 0xae, 0xed, 0x3b, 0x82, 0x66, 0x6e,
	0xe4, 0x9c, 0xc2, 0xfb, 0x4a, 0xe1, 0x6c, 0x33, 0x2e, 0x96, 0x05, 0x38, 0xe7, 0xb2, 0xc3, 0x5c,
	0xad, 0x79, 0x1d, 0xce, 0x2c, 0xbc, 0x70, 0xb8, 0x09, 0x27, 0x62, 0x77, 0xf4, 0x3e, 0x5d, 0x39,
	0xfa, 0xc6, 0x34, 0x01, 0xf2, 0xd7, 0xa1, 0x02, 0xb3, 0x0b, 0x8d, 0x42, 0x64, 0x2b, 0x60, 0xbf,
	0xd0, 0xd2, 0x8d, 0x5b, 0x70, 0x15, 0xff, 0x58, 0x01, 0x64, 0x20, 0xf5, 0xe6, 0x4d, 0x37, 0x3f,
	0x74, 0xc2, 0x27, 0x6c, 0xec, 0x5c, 0x92, 0xb0, 0x47, 0x1a, 0xd5, 0x97, 0xa0, 0xc1, 0x5c, 0x12,
	0xcc, 0x37, 0xb0, 0xbb, 0x52, 0x1b, 0xbc, 0x0f, 0x86, 0xae, 0xe5, 0xe7, 0xd9, 0x6e, 0xd3, 0xc6,
	0xed, 0xa8, 0x7a, 0xd6, 0x7a, 0xa0, 0x77, 0x91, 0xf9, 0x1b, 0xc1, 0xb5, 0xd5, 0x86, 0x8f, 0x7f,
	0xae, 0x0c, 0x4d, 0x4c, 0xfd, 0xf0, 0xbf, 0x72, 0xaa, 0xad, 0x35, 0x17, 0xad, 0xb5, 0x4b, 0xe1,
	0x35, 0x05, 0xdc, 0xf8, 0x8b, 0x96, 0xf8, 0x21, 0x5c, 0x5e, 0xb4, 0x79, 0x61, 0x91, 0x6b, 0xb3,
	0x5b, 0x45, 0xb3, 0xed, 0xf9, 0xad, 0x5b, 0x95, 0xdd, 0xc2, 0xd6, 0xaa, 0xca, 0x6e, 0x01, 0xb4,
	0xd1, 0xec, 0x16, 0x98, 0x2b, 0xb2, 0x6b, 0xf3, 0x7c, 0xc3, 0x9a, 0x9f, 0xd0, 0xd2, 0xf0, 0x96,
	0xfb, 0xe4, 0xe1, 0x95, 0xb5, 0xa5, 0x7e, 0xea, 0xf0, 0xda, 0x9c, 0x16, 0x9c, 0xc4, 0x37, 0x01,
	0x2f, 0x0e, 0x8c, 0xd3, 0xe4, 0x33, 0xe6, 0x6c, 0x58, 0x18, 0xb9, 0x4d, 0xdf, 0x9e, 0x54, 0xdf,
	0x61, 0x7b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x80, 0x07, 0x13, 0xab, 0x09, 0x00, 0x00,
}
