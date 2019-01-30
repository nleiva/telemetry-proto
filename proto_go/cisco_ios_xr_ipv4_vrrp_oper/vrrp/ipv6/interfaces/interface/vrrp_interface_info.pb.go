// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vrrp_interface_info.proto

package cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface

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

// VRRP Interface statistics
type VrrpInterfaceInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VrrpInterfaceInfo_KEYS) Reset()         { *m = VrrpInterfaceInfo_KEYS{} }
func (m *VrrpInterfaceInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*VrrpInterfaceInfo_KEYS) ProtoMessage()    {}
func (*VrrpInterfaceInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_76aa910590caf3c8, []int{0}
}

func (m *VrrpInterfaceInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrrpInterfaceInfo_KEYS.Unmarshal(m, b)
}
func (m *VrrpInterfaceInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrrpInterfaceInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *VrrpInterfaceInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrrpInterfaceInfo_KEYS.Merge(m, src)
}
func (m *VrrpInterfaceInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_VrrpInterfaceInfo_KEYS.Size(m)
}
func (m *VrrpInterfaceInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VrrpInterfaceInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VrrpInterfaceInfo_KEYS proto.InternalMessageInfo

func (m *VrrpInterfaceInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type VrrpInterfaceInfo struct {
	// IM Interface
	Interface string `protobuf:"bytes,50,opt,name=interface,proto3" json:"interface,omitempty"`
	// Invalid checksum
	InvalidChecksumCount uint32 `protobuf:"varint,51,opt,name=invalid_checksum_count,json=invalidChecksumCount,proto3" json:"invalid_checksum_count,omitempty"`
	// Unknown/unsupported version
	InvalidVersionCount uint32 `protobuf:"varint,52,opt,name=invalid_version_count,json=invalidVersionCount,proto3" json:"invalid_version_count,omitempty"`
	// Invalid vrID
	InvalidVridCount uint32 `protobuf:"varint,53,opt,name=invalid_vrid_count,json=invalidVridCount,proto3" json:"invalid_vrid_count,omitempty"`
	// Bad packet lengths
	InvalidPacketLengthCount uint32   `protobuf:"varint,54,opt,name=invalid_packet_length_count,json=invalidPacketLengthCount,proto3" json:"invalid_packet_length_count,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *VrrpInterfaceInfo) Reset()         { *m = VrrpInterfaceInfo{} }
func (m *VrrpInterfaceInfo) String() string { return proto.CompactTextString(m) }
func (*VrrpInterfaceInfo) ProtoMessage()    {}
func (*VrrpInterfaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76aa910590caf3c8, []int{1}
}

func (m *VrrpInterfaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrrpInterfaceInfo.Unmarshal(m, b)
}
func (m *VrrpInterfaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrrpInterfaceInfo.Marshal(b, m, deterministic)
}
func (m *VrrpInterfaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrrpInterfaceInfo.Merge(m, src)
}
func (m *VrrpInterfaceInfo) XXX_Size() int {
	return xxx_messageInfo_VrrpInterfaceInfo.Size(m)
}
func (m *VrrpInterfaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VrrpInterfaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VrrpInterfaceInfo proto.InternalMessageInfo

func (m *VrrpInterfaceInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *VrrpInterfaceInfo) GetInvalidChecksumCount() uint32 {
	if m != nil {
		return m.InvalidChecksumCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidVersionCount() uint32 {
	if m != nil {
		return m.InvalidVersionCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidVridCount() uint32 {
	if m != nil {
		return m.InvalidVridCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidPacketLengthCount() uint32 {
	if m != nil {
		return m.InvalidPacketLengthCount
	}
	return 0
}

func init() {
	proto.RegisterType((*VrrpInterfaceInfo_KEYS)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv6.interfaces.interface.vrrp_interface_info_KEYS")
	proto.RegisterType((*VrrpInterfaceInfo)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv6.interfaces.interface.vrrp_interface_info")
}

func init() { proto.RegisterFile("vrrp_interface_info.proto", fileDescriptor_76aa910590caf3c8) }

var fileDescriptor_76aa910590caf3c8 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0x06, 0x70, 0xfa, 0x3f, 0xfc, 0x61, 0x81, 0x89, 0x64, 0x2a, 0x15, 0x3d, 0x8c, 0x81, 0xb0,
	0x83, 0xf4, 0xb0, 0xd5, 0x1d, 0x04, 0x0f, 0x32, 0x3c, 0x29, 0x22, 0x13, 0x04, 0x4f, 0x2f, 0x31,
	0xcb, 0xdc, 0xcb, 0xd6, 0x24, 0x24, 0x59, 0xf0, 0x33, 0xf8, 0xa9, 0xa5, 0x6f, 0xd3, 0xf6, 0xb2,
	0x5b, 0x78, 0x9e, 0xe7, 0xf7, 0xf6, 0x50, 0x76, 0x19, 0x9d, 0xb3, 0x80, 0x3a, 0x28, 0xb7, 0x11,
	0x52, 0x01, 0xea, 0x8d, 0x29, 0xac, 0x33, 0xc1, 0xf0, 0x7b, 0x89, 0x5e, 0x1a, 0x40, 0xe3, 0xe1,
	0xc7, 0x01, 0xda, 0x58, 0x02, 0x8d, 0x8d, 0x55, 0xae, 0xa8, 0x5f, 0x05, 0xda, 0xb8, 0x28, 0x3a,
	0xeb, 0xfb, 0xe7, 0xe4, 0x91, 0xe5, 0x47, 0x0e, 0xc3, 0xf3, 0xd3, 0xe7, 0x3b, 0xbf, 0x61, 0x27,
	0x7d, 0xac, 0x45, 0xa5, 0xf2, 0x6c, 0x9c, 0x4d, 0x07, 0xab, 0x61, 0x97, 0xbe, 0x8a, 0x4a, 0x4d,
	0x7e, 0xff, 0xb1, 0xd1, 0x91, 0x1b, 0xfc, 0x9a, 0x0d, 0xba, 0x24, 0x9f, 0x91, 0xec, 0x03, 0x5e,
	0xb2, 0x0b, 0xd4, 0x51, 0xec, 0x71, 0x0d, 0x72, 0xab, 0xe4, 0xce, 0x1f, 0x2a, 0x90, 0xe6, 0xa0,
	0x43, 0x3e, 0x1f, 0x67, 0xd3, 0xe1, 0xea, 0x2c, 0xb5, 0xcb, 0x54, 0x2e, 0xeb, 0x8e, 0xcf, 0xd8,
	0x79, 0xab, 0xa2, 0x72, 0x1e, 0x8d, 0x4e, 0xa8, 0x24, 0x34, 0x4a, 0xe5, 0x47, 0xd3, 0x35, 0xe6,
	0x96, 0xf1, 0xce, 0xb8, 0xfa, 0x73, 0x04, 0xee, 0x08, 0x9c, 0xb6, 0xc0, 0xe1, 0xba, 0x59, 0x3f,
	0xb0, 0xab, 0x76, 0x6d, 0x85, 0xdc, 0xa9, 0x00, 0x7b, 0xa5, 0xbf, 0xc3, 0x36, 0xb1, 0x05, 0xb1,
	0x3c, 0x4d, 0xde, 0x68, 0xf1, 0x42, 0x03, 0xe2, 0x5f, 0xff, 0xe9, 0x97, 0xcc, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x32, 0x14, 0xcd, 0xaf, 0x01, 0x00, 0x00,
}
