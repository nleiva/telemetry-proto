// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_neighbor.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_adjacency_information_neighbors_neighbor

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

// OSPF Neighbor Summary Information
type OspfShNeighbor_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighbor_KEYS) Reset()         { *m = OspfShNeighbor_KEYS{} }
func (m *OspfShNeighbor_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor_KEYS) ProtoMessage()    {}
func (*OspfShNeighbor_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_0b83212aba965586, []int{0}
}
func (m *OspfShNeighbor_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Unmarshal(m, b)
}
func (m *OspfShNeighbor_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighbor_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor_KEYS.Merge(dst, src)
}
func (m *OspfShNeighbor_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Size(m)
}
func (m *OspfShNeighbor_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor_KEYS proto.InternalMessageInfo

func (m *OspfShNeighbor_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShNeighbor_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShNeighbor_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShNeighbor struct {
	// Neighbor ID
	NeighborId string `protobuf:"bytes,50,opt,name=neighbor_id,json=neighborId,proto3" json:"neighbor_id,omitempty"`
	// Neighbor IP Address
	NeighborAddress string `protobuf:"bytes,51,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	// Interface via which the neighbor is seen
	NeighborInterfaceName string `protobuf:"bytes,52,opt,name=neighbor_interface_name,json=neighborInterfaceName,proto3" json:"neighbor_interface_name,omitempty"`
	// Neighbor's DR priority
	NeighborDrPriority uint32 `protobuf:"varint,53,opt,name=neighbor_dr_priority,json=neighborDrPriority,proto3" json:"neighbor_dr_priority,omitempty"`
	// Neighbor's state
	NeighborState string `protobuf:"bytes,54,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	// Designated Router state
	DrBdrState string `protobuf:"bytes,55,opt,name=dr_bdr_state,json=drBdrState,proto3" json:"dr_bdr_state,omitempty"`
	// Time until neighbor's dead timer expires (s)
	NeighborDeadTimer uint32 `protobuf:"varint,56,opt,name=neighbor_dead_timer,json=neighborDeadTimer,proto3" json:"neighbor_dead_timer,omitempty"`
	// Amount of time since the adjacency is up (s)
	NeighborUpTime uint32 `protobuf:"varint,57,opt,name=neighbor_up_time,json=neighborUpTime,proto3" json:"neighbor_up_time,omitempty"`
	// Interface is MADJ
	NeighborMadjInterface bool `protobuf:"varint,58,opt,name=neighbor_madj_interface,json=neighborMadjInterface,proto3" json:"neighbor_madj_interface,omitempty"`
	// Neighbor BFD information
	NeighborBfdInformation *OspfShNeighborBfd `protobuf:"bytes,59,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShNeighbor) Reset()         { *m = OspfShNeighbor{} }
func (m *OspfShNeighbor) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor) ProtoMessage()    {}
func (*OspfShNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_0b83212aba965586, []int{1}
}
func (m *OspfShNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor.Unmarshal(m, b)
}
func (m *OspfShNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor.Merge(dst, src)
}
func (m *OspfShNeighbor) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor.Size(m)
}
func (m *OspfShNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor proto.InternalMessageInfo

func (m *OspfShNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDrPriority() uint32 {
	if m != nil {
		return m.NeighborDrPriority
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *OspfShNeighbor) GetDrBdrState() string {
	if m != nil {
		return m.DrBdrState
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDeadTimer() uint32 {
	if m != nil {
		return m.NeighborDeadTimer
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborUpTime() uint32 {
	if m != nil {
		return m.NeighborUpTime
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborMadjInterface() bool {
	if m != nil {
		return m.NeighborMadjInterface
	}
	return false
}

func (m *OspfShNeighbor) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

// OSPF Neighbor BFD information
type OspfShNeighborBfd struct {
	// BFD enable mode - Default/Strict
	BfdIntfEnableMode uint32 `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	// Status of the BFD Session
	BfdStatusFlag        uint32   `protobuf:"varint,2,opt,name=bfd_status_flag,json=bfdStatusFlag,proto3" json:"bfd_status_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborBfd) Reset()         { *m = OspfShNeighborBfd{} }
func (m *OspfShNeighborBfd) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborBfd) ProtoMessage()    {}
func (*OspfShNeighborBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_0b83212aba965586, []int{2}
}
func (m *OspfShNeighborBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborBfd.Unmarshal(m, b)
}
func (m *OspfShNeighborBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborBfd.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighborBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborBfd.Merge(dst, src)
}
func (m *OspfShNeighborBfd) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborBfd.Size(m)
}
func (m *OspfShNeighborBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborBfd.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborBfd proto.InternalMessageInfo

func (m *OspfShNeighborBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *OspfShNeighborBfd) GetBfdStatusFlag() uint32 {
	if m != nil {
		return m.BfdStatusFlag
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShNeighbor_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor_KEYS")
	proto.RegisterType((*OspfShNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor")
	proto.RegisterType((*OspfShNeighborBfd)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor_bfd")
}

func init() {
	proto.RegisterFile("ospf_sh_neighbor.proto", fileDescriptor_ospf_sh_neighbor_0b83212aba965586)
}

var fileDescriptor_ospf_sh_neighbor_0b83212aba965586 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0x56, 0x60, 0x82, 0xe1, 0x2e, 0x6c, 0x33, 0xdd, 0x30, 0x27, 0x42, 0x25, 0x50, 0xb8, 0x04,
	0xb4, 0x8d, 0xf1, 0xef, 0xc4, 0xc4, 0x90, 0x26, 0x34, 0x84, 0x52, 0x38, 0x70, 0xb2, 0x9c, 0xfc,
	0xec, 0xce, 0x55, 0x13, 0x47, 0xb6, 0x57, 0x31, 0x5e, 0x88, 0x17, 0xe0, 0x69, 0x78, 0x1a, 0x64,
	0xa7, 0x71, 0xd6, 0xd2, 0xf3, 0x2e, 0x95, 0xfb, 0xfd, 0xf1, 0xf7, 0xf3, 0x17, 0x1b, 0xed, 0x2b,
	0xd3, 0x08, 0x6a, 0x2e, 0x68, 0xcd, 0xe5, 0xe4, 0xa2, 0x50, 0x3a, 0x6b, 0xb4, 0xb2, 0x0a, 0x43,
	0x29, 0x4d, 0xa9, 0xa8, 0x54, 0x86, 0xfe, 0xd4, 0x54, 0x36, 0xf3, 0x23, 0xea, 0x95, 0xaa, 0xe1,
	0x3a, 0x73, 0x2b, 0xa7, 0x2b, 0xb9, 0x31, 0xdc, 0x74, 0xab, 0x6c, 0xae, 0x85, 0xff, 0xc9, 0x18,
	0x4c, 0x59, 0xc9, 0xeb, 0xf2, 0x8a, 0xca, 0x5a, 0x28, 0x5d, 0x31, 0x2b, 0x55, 0x9d, 0x75, 0x09,
	0x26, 0xac, 0x46, 0xbf, 0x23, 0xb4, 0xb7, 0x3a, 0x00, 0xfd, 0x7c, 0xfa, 0x63, 0x8c, 0x9f, 0xa0,
	0xad, 0xc5, 0xb6, 0xb4, 0x66, 0x15, 0x27, 0x51, 0x12, 0xa5, 0xf7, 0xf2, 0xc1, 0x02, 0xfb, 0xc2,
	0x2a, 0x8e, 0x1f, 0xa1, 0xcd, 0xb9, 0x16, 0x2d, 0x7d, 0xcb, 0xd3, 0x77, 0xe7, 0x5a, 0x78, 0xea,
	0x29, 0xba, 0x2f, 0x6b, 0xcb, 0xb5, 0x60, 0x25, 0x6f, 0x05, 0xb7, 0xbd, 0x20, 0x0e, 0xa8, 0x97,
	0x3d, 0x47, 0x3b, 0x21, 0x95, 0x01, 0x68, 0x6e, 0x0c, 0xd9, 0xf0, 0xc2, 0xed, 0x0e, 0xff, 0xd0,
	0xc2, 0xa3, 0xbf, 0x1b, 0x68, 0x67, 0x75, 0x52, 0xfc, 0x18, 0x0d, 0x82, 0x5f, 0x02, 0x39, 0xf0,
	0x56, 0xd4, 0x41, 0x67, 0xb0, 0x36, 0xe0, 0x70, 0x6d, 0x00, 0x3e, 0x46, 0x0f, 0xfb, 0xbd, 0x96,
	0x67, 0x3f, 0xf2, 0x8e, 0xbd, 0xb0, 0xef, 0xd2, 0x19, 0x5e, 0xa2, 0x61, 0xf0, 0x81, 0xa6, 0x8d,
	0x96, 0x4a, 0x4b, 0x7b, 0x45, 0x5e, 0x25, 0x51, 0x1a, 0xe7, 0xb8, 0xe3, 0x3e, 0xea, 0xaf, 0x0b,
	0xc6, 0x95, 0x13, 0x1c, 0xc6, 0x32, 0xcb, 0xc9, 0x71, 0x5b, 0x4e, 0x87, 0x8e, 0x1d, 0x88, 0x13,
	0xb4, 0x05, 0x9a, 0x16, 0xd0, 0x89, 0x5e, 0xb7, 0xa7, 0x03, 0x7d, 0x02, 0x0b, 0x45, 0x86, 0x1e,
	0xf4, 0xd1, 0x9c, 0x01, 0xb5, 0xb2, 0xe2, 0x9a, 0xbc, 0xf1, 0xc9, 0xbb, 0x21, 0x99, 0x33, 0xf8,
	0xe6, 0x08, 0x9c, 0x5e, 0x6b, 0xe3, 0xb2, 0xf1, 0x6a, 0xf2, 0xd6, 0x8b, 0xc3, 0x40, 0xdf, 0x1b,
	0x27, 0x5d, 0x2a, 0xa3, 0x62, 0x30, 0xed, 0x1b, 0x21, 0xef, 0x92, 0x28, 0xdd, 0xec, 0xcb, 0x38,
	0x67, 0x30, 0x0d, 0x85, 0xe0, 0x3f, 0x11, 0x22, 0xc1, 0x58, 0x08, 0xb8, 0x7e, 0xff, 0xc8, 0xfb,
	0x24, 0x4a, 0x07, 0x07, 0xbf, 0xb2, 0x9b, 0xb8, 0xd9, 0xd9, 0x7f, 0xb7, 0xba, 0x10, 0x90, 0xef,
	0x77, 0xff, 0x4e, 0x04, 0x9c, 0xf5, 0xfe, 0x91, 0x42, 0xc3, 0x75, 0x7a, 0xfc, 0x02, 0x0d, 0xdb,
	0x43, 0x58, 0x41, 0x79, 0xcd, 0x8a, 0x19, 0xa7, 0x95, 0x82, 0xf6, 0x31, 0xc4, 0xf9, 0x6e, 0xe1,
	0x76, 0xb1, 0xe2, 0xd4, 0x33, 0xe7, 0x0a, 0x38, 0x7e, 0x86, 0xb6, 0x9d, 0xc1, 0x7d, 0xb0, 0x4b,
	0x43, 0xc5, 0x8c, 0x4d, 0xfc, 0xcb, 0x88, 0xf3, 0xb8, 0x10, 0x30, 0xf6, 0xe8, 0xa7, 0x19, 0x9b,
	0x14, 0x77, 0xfc, 0x23, 0x3f, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x30, 0xa4, 0xb8, 0xd7, 0xfe,
	0x03, 0x00, 0x00,
}
