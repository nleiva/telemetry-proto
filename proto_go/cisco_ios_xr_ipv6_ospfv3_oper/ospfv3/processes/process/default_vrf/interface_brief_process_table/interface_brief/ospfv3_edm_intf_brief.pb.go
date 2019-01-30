// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_intf_brief.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_interface_brief_process_table_interface_brief

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

// OSPFv3 brief interface information
type Ospfv3EdmIntfBrief_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmIntfBrief_KEYS) Reset()         { *m = Ospfv3EdmIntfBrief_KEYS{} }
func (m *Ospfv3EdmIntfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmIntfBrief_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmIntfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f427722ca06b929e, []int{0}
}

func (m *Ospfv3EdmIntfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmIntfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmIntfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmIntfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS.Size(m)
}
func (m *Ospfv3EdmIntfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmIntfBrief_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmIntfBrief_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmIntfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmIntfBrief struct {
	// Interface IP address
	InterfaceAddress string `protobuf:"bytes,50,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,51,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,52,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	// Total number of neighbors
	InterfaceNeighbors uint32 `protobuf:"varint,53,opt,name=interface_neighbors,json=interfaceNeighbors,proto3" json:"interface_neighbors,omitempty"`
	// Total number of adjacent neighbors
	InterfaceAdjacentNeighbors uint32 `protobuf:"varint,54,opt,name=interface_adjacent_neighbors,json=interfaceAdjacentNeighbors,proto3" json:"interface_adjacent_neighbors,omitempty"`
	// Network type
	NetworkType          string   `protobuf:"bytes,55,opt,name=network_type,json=networkType,proto3" json:"network_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmIntfBrief) Reset()         { *m = Ospfv3EdmIntfBrief{} }
func (m *Ospfv3EdmIntfBrief) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmIntfBrief) ProtoMessage()    {}
func (*Ospfv3EdmIntfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_f427722ca06b929e, []int{1}
}

func (m *Ospfv3EdmIntfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmIntfBrief.Unmarshal(m, b)
}
func (m *Ospfv3EdmIntfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmIntfBrief.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmIntfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmIntfBrief.Merge(m, src)
}
func (m *Ospfv3EdmIntfBrief) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmIntfBrief.Size(m)
}
func (m *Ospfv3EdmIntfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmIntfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmIntfBrief proto.InternalMessageInfo

func (m *Ospfv3EdmIntfBrief) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *Ospfv3EdmIntfBrief) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *Ospfv3EdmIntfBrief) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *Ospfv3EdmIntfBrief) GetInterfaceNeighbors() uint32 {
	if m != nil {
		return m.InterfaceNeighbors
	}
	return 0
}

func (m *Ospfv3EdmIntfBrief) GetInterfaceAdjacentNeighbors() uint32 {
	if m != nil {
		return m.InterfaceAdjacentNeighbors
	}
	return 0
}

func (m *Ospfv3EdmIntfBrief) GetNetworkType() string {
	if m != nil {
		return m.NetworkType
	}
	return ""
}

func init() {
	proto.RegisterType((*Ospfv3EdmIntfBrief_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.interface_brief_process_table.interface_brief.ospfv3_edm_intf_brief_KEYS")
	proto.RegisterType((*Ospfv3EdmIntfBrief)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.interface_brief_process_table.interface_brief.ospfv3_edm_intf_brief")
}

func init() { proto.RegisterFile("ospfv3_edm_intf_brief.proto", fileDescriptor_f427722ca06b929e) }

var fileDescriptor_f427722ca06b929e = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xd9, 0x0e, 0x82, 0x71, 0x13, 0x8d, 0x0a, 0x65, 0x7a, 0xd8, 0x06, 0xc2, 0x40, 0xa8,
	0xe2, 0x74, 0x5e, 0x15, 0xf1, 0x20, 0x8a, 0x87, 0xcd, 0x8b, 0xa7, 0x3f, 0x69, 0xfb, 0x8f, 0xc6,
	0x6d, 0x49, 0x48, 0x62, 0x75, 0xdf, 0xcd, 0x0f, 0x27, 0x4d, 0xdb, 0xa5, 0xc8, 0x6e, 0xe1, 0xbd,
	0xdf, 0x7b, 0x2f, 0x29, 0x25, 0xc7, 0xca, 0x6a, 0x9e, 0x8f, 0x01, 0xb3, 0x25, 0x08, 0xe9, 0x38,
	0x24, 0x46, 0x20, 0x8f, 0xb5, 0x51, 0x4e, 0x51, 0x9d, 0x0a, 0x9b, 0x2a, 0x10, 0xca, 0xc2, 0x8f,
	0x01, 0xa1, 0xf3, 0x09, 0x54, 0xb8, 0xd2, 0x68, 0xe2, 0xf2, 0x5c, 0xb0, 0x29, 0x5a, 0x8b, 0xb6,
	0x3e, 0xc5, 0x19, 0x72, 0xf6, 0xb5, 0x70, 0x90, 0x1b, 0x1e, 0x0b, 0xe9, 0xd0, 0x70, 0x96, 0x62,
	0xd9, 0x0c, 0x15, 0x03, 0x8e, 0x25, 0x0b, 0xfc, 0xef, 0x0e, 0x39, 0xe9, 0x6d, 0xbc, 0x10, 0x3c,
	0x3d, 0xbc, 0xcd, 0xe8, 0x80, 0x74, 0xea, 0xb8, 0x64, 0x4b, 0x8c, 0x5a, 0xfd, 0xd6, 0x68, 0x7b,
	0xba, 0x53, 0x69, 0x2f, 0x6c, 0x89, 0xf4, 0x94, 0xec, 0x86, 0x4e, 0x0f, 0xb5, 0x3d, 0xd4, 0x5d,
	0xab, 0x05, 0x36, 0xfc, 0x6d, 0x93, 0xa3, 0x8d, 0x43, 0xf4, 0x8c, 0xec, 0x87, 0x02, 0x96, 0x65,
	0x06, 0xad, 0x8d, 0x2e, 0x7d, 0xc7, 0xde, 0xda, 0xb8, 0x2b, 0x75, 0x1a, 0x93, 0x83, 0x00, 0x2f,
	0x84, 0x9c, 0x43, 0xaa, 0xac, 0x8b, 0xc6, 0xfd, 0xd6, 0xa8, 0x3b, 0x0d, 0x3d, 0xcf, 0x42, 0xce,
	0xef, 0x95, 0x75, 0xf4, 0x82, 0x1c, 0x16, 0xab, 0x10, 0x42, 0xd6, 0x31, 0x87, 0xd1, 0x95, 0xef,
	0xa7, 0x85, 0xf7, 0x58, 0x5b, 0xb3, 0xc2, 0xa1, 0xe7, 0xcd, 0x05, 0x89, 0xe2, 0xfd, 0x23, 0x51,
	0xc6, 0x46, 0xd7, 0x7e, 0x81, 0x86, 0x47, 0xd5, 0x0e, 0xbd, 0x25, 0x27, 0xcd, 0xfb, 0x7f, 0xb2,
	0x14, 0xa5, 0x6b, 0x24, 0x27, 0x3e, 0xd9, 0x6b, 0x3c, 0xa5, 0x44, 0x42, 0xc3, 0x80, 0x74, 0x24,
	0xba, 0x6f, 0x65, 0xe6, 0xe0, 0x56, 0x1a, 0xa3, 0x9b, 0xf2, 0x2b, 0x57, 0xda, 0xeb, 0x4a, 0x63,
	0xb2, 0xe5, 0xff, 0x8f, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x24, 0xa2, 0x49, 0x3e,
	0x02, 0x00, 0x00,
}
