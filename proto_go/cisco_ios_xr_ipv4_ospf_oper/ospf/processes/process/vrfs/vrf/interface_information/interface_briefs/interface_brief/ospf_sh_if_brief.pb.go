// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_if_brief.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_interface_information_interface_briefs_interface_brief

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

// OSPF Interface Brief Information
type OspfShIfBrief_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIfBrief_KEYS) Reset()         { *m = OspfShIfBrief_KEYS{} }
func (m *OspfShIfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShIfBrief_KEYS) ProtoMessage()    {}
func (*OspfShIfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_if_brief_4e8598e99bb6873f, []int{0}
}
func (m *OspfShIfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Unmarshal(m, b)
}
func (m *OspfShIfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShIfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfBrief_KEYS.Merge(dst, src)
}
func (m *OspfShIfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Size(m)
}
func (m *OspfShIfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfBrief_KEYS proto.InternalMessageInfo

func (m *OspfShIfBrief_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShIfBrief struct {
	// Interface name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,51,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	// Interface IP Address
	InterfaceAddress string `protobuf:"bytes,52,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	// Interface IP Mask
	InterfaceMask uint32 `protobuf:"varint,53,opt,name=interface_mask,json=interfaceMask,proto3" json:"interface_mask,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,54,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,55,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	// Interface in fast detect hold down state
	InterfaceFastDetectHoldDown bool `protobuf:"varint,56,opt,name=interface_fast_detect_hold_down,json=interfaceFastDetectHoldDown,proto3" json:"interface_fast_detect_hold_down,omitempty"`
	// Total number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,57,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,58,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	// If true, interface is multi-area
	InterfaceisMadj bool `protobuf:"varint,59,opt,name=interfaceis_madj,json=interfaceisMadj,proto3" json:"interfaceis_madj,omitempty"`
	// Total number of multi-area
	InterfaceMadjCount uint32 `protobuf:"varint,60,opt,name=interface_madj_count,json=interfaceMadjCount,proto3" json:"interface_madj_count,omitempty"`
	// Information for multi-area on the interface
	InterfaceMadjList    []*OspfShInterfaceMadj `protobuf:"bytes,61,rep,name=interface_madj_list,json=interfaceMadjList,proto3" json:"interface_madj_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OspfShIfBrief) Reset()         { *m = OspfShIfBrief{} }
func (m *OspfShIfBrief) String() string { return proto.CompactTextString(m) }
func (*OspfShIfBrief) ProtoMessage()    {}
func (*OspfShIfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_if_brief_4e8598e99bb6873f, []int{1}
}
func (m *OspfShIfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfBrief.Unmarshal(m, b)
}
func (m *OspfShIfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfBrief.Marshal(b, m, deterministic)
}
func (dst *OspfShIfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfBrief.Merge(dst, src)
}
func (m *OspfShIfBrief) XXX_Size() int {
	return xxx_messageInfo_OspfShIfBrief.Size(m)
}
func (m *OspfShIfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfBrief proto.InternalMessageInfo

func (m *OspfShIfBrief) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceMask() uint32 {
	if m != nil {
		return m.InterfaceMask
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShIfBrief) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceFastDetectHoldDown() bool {
	if m != nil {
		return m.InterfaceFastDetectHoldDown
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceisMadj() bool {
	if m != nil {
		return m.InterfaceisMadj
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceMadjCount() uint32 {
	if m != nil {
		return m.InterfaceMadjCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceMadjList() []*OspfShInterfaceMadj {
	if m != nil {
		return m.InterfaceMadjList
	}
	return nil
}

// OSPF Interface Multi-Area Information
type OspfShInterfaceMadj struct {
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,1,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	// Area ID
	MadjAreaId uint32 `protobuf:"varint,2,opt,name=madj_area_id,json=madjAreaId,proto3" json:"madj_area_id,omitempty"`
	// Number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,3,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,4,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,5,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState   string   `protobuf:"bytes,6,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShInterfaceMadj) Reset()         { *m = OspfShInterfaceMadj{} }
func (m *OspfShInterfaceMadj) String() string { return proto.CompactTextString(m) }
func (*OspfShInterfaceMadj) ProtoMessage()    {}
func (*OspfShInterfaceMadj) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_if_brief_4e8598e99bb6873f, []int{2}
}
func (m *OspfShInterfaceMadj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShInterfaceMadj.Unmarshal(m, b)
}
func (m *OspfShInterfaceMadj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShInterfaceMadj.Marshal(b, m, deterministic)
}
func (dst *OspfShInterfaceMadj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShInterfaceMadj.Merge(dst, src)
}
func (m *OspfShInterfaceMadj) XXX_Size() int {
	return xxx_messageInfo_OspfShInterfaceMadj.Size(m)
}
func (m *OspfShInterfaceMadj) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShInterfaceMadj.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShInterfaceMadj proto.InternalMessageInfo

func (m *OspfShInterfaceMadj) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShInterfaceMadj) GetMadjAreaId() uint32 {
	if m != nil {
		return m.MadjAreaId
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShIfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS")
	proto.RegisterType((*OspfShIfBrief)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief")
	proto.RegisterType((*OspfShInterfaceMadj)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_interface_madj")
}

func init() {
	proto.RegisterFile("ospf_sh_if_brief.proto", fileDescriptor_ospf_sh_if_brief_4e8598e99bb6873f)
}

var fileDescriptor_ospf_sh_if_brief_4e8598e99bb6873f = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0xc7, 0xd9, 0xe6, 0xd7, 0xfe, 0xea, 0xb4, 0xd1, 0x3a, 0x6a, 0x99, 0xa2, 0xe0, 0x1a, 0x10,
	0x22, 0xc2, 0x22, 0x6d, 0xd5, 0xfa, 0x0f, 0x91, 0x46, 0xb1, 0xd8, 0x7a, 0x91, 0x5e, 0x79, 0x35,
	0x4c, 0x76, 0x66, 0xcd, 0x6c, 0xb2, 0x3b, 0x61, 0xce, 0x98, 0x7a, 0xe1, 0x13, 0xf8, 0x4a, 0x82,
	0x0f, 0xe5, 0x13, 0xc8, 0x9c, 0xa4, 0xd9, 0xd9, 0x25, 0x55, 0x44, 0xf0, 0x66, 0x59, 0xce, 0xf7,
	0x73, 0xe6, 0x9c, 0xfd, 0xce, 0x39, 0x4b, 0xb6, 0x0d, 0x4c, 0x32, 0x0e, 0x43, 0xae, 0x33, 0x3e,
	0xb0, 0x5a, 0x65, 0xc9, 0xc4, 0x1a, 0x67, 0xa8, 0x4d, 0x35, 0xa4, 0x86, 0x6b, 0x03, 0xfc, 0xb3,
	0xe5, 0x7a, 0x32, 0xdd, 0xe7, 0x48, 0x9a, 0x89, 0xb2, 0x89, 0x7f, 0xf3, 0x5c, 0xaa, 0x00, 0x14,
	0x9c, 0xbf, 0x25, 0x53, 0x9b, 0xe1, 0x23, 0xd1, 0xa5, 0x53, 0x36, 0x13, 0xa9, 0xe2, 0xba, 0xcc,
	0x8c, 0x2d, 0x84, 0xd3, 0xa6, 0x0c, 0xa2, 0x58, 0x08, 0x9a, 0x81, 0xce, 0x17, 0x72, 0xa3, 0xd9,
	0x0d, 0x7f, 0xf7, 0xfa, 0xc3, 0x29, 0xbd, 0x43, 0x36, 0xe7, 0x35, 0x78, 0x29, 0x0a, 0xc5, 0xa2,
	0x38, 0xea, 0x5e, 0xea, 0x6f, 0xcc, 0x63, 0xef, 0x45, 0xa1, 0xe8, 0x0e, 0x59, 0x9f, 0xda, 0x6c,
	0x26, 0xaf, 0xa0, 0xfc, 0xff, 0xd4, 0x66, 0x28, 0xdd, 0x25, 0x97, 0xab, 0x4a, 0x08, 0xb4, 0x10,
	0x68, 0x2f, 0xa2, 0x1e, 0xeb, 0xfc, 0x58, 0x25, 0x5b, 0xcd, 0xf2, 0x4b, 0x72, 0x77, 0x97, 0xe4,
	0xd6, 0x31, 0x61, 0x95, 0x60, 0x7b, 0x0d, 0xec, 0x95, 0x55, 0x82, 0xde, 0x27, 0x57, 0x03, 0x4c,
	0x4a, 0xab, 0x00, 0xd8, 0x3e, 0x92, 0x5b, 0x15, 0x39, 0x8b, 0xd7, 0xcf, 0x2c, 0x04, 0x8c, 0xd8,
	0xc3, 0x38, 0xea, 0xb6, 0x83, 0x33, 0x4f, 0x04, 0x8c, 0x68, 0x42, 0xae, 0x55, 0xd8, 0x58, 0x97,
	0x23, 0x9e, 0x1a, 0x70, 0xec, 0x11, 0xb2, 0x55, 0xb9, 0x63, 0x5d, 0x8e, 0x0e, 0x0d, 0x38, 0xfa,
	0x80, 0x5c, 0xc7, 0xaf, 0xac, 0x92, 0xc0, 0x09, 0xa7, 0xd8, 0x63, 0x6c, 0x83, 0x7a, 0xed, 0xe8,
	0x5c, 0x3a, 0xf5, 0x0a, 0xed, 0x91, 0xdb, 0x15, 0x9c, 0x09, 0x70, 0x5c, 0x2a, 0xa7, 0x52, 0xc7,
	0x87, 0x66, 0x2c, 0xb9, 0x34, 0x67, 0x25, 0x3b, 0x88, 0xa3, 0xee, 0x7a, 0xff, 0xe6, 0x02, 0x7b,
	0x23, 0xc0, 0xf5, 0x10, 0x7a, 0x6b, 0xc6, 0xb2, 0x67, 0xce, 0x4a, 0x7a, 0x40, 0x58, 0xe0, 0xa4,
	0xd2, 0x1f, 0x87, 0x03, 0x63, 0x79, 0x6a, 0x3e, 0x95, 0x8e, 0x3d, 0xc1, 0x66, 0xb7, 0x2b, 0x4f,
	0xe7, 0xf2, 0xa1, 0x57, 0xe9, 0x4b, 0x72, 0x2b, 0x74, 0x2d, 0x6f, 0x66, 0x3f, 0xc5, 0xec, 0x9d,
	0xc0, 0xc0, 0xbc, 0x7e, 0xc0, 0x3d, 0x52, 0xb9, 0xab, 0x81, 0x17, 0x42, 0xe6, 0xec, 0x19, 0x76,
	0x7c, 0x25, 0x88, 0x9f, 0x08, 0x99, 0x7b, 0x77, 0x42, 0xd3, 0x65, 0x3e, 0xaf, 0xf1, 0x1c, 0x6b,
	0xd0, 0xc0, 0x7a, 0x99, 0xcf, 0x0e, 0xff, 0x16, 0x85, 0x17, 0x80, 0x29, 0x63, 0x0d, 0x8e, 0xbd,
	0x88, 0x5b, 0xdd, 0x8d, 0xdd, 0xaf, 0x51, 0xf2, 0xef, 0x17, 0x29, 0x59, 0x8c, 0x71, 0xad, 0xaf,
	0x60, 0x1a, 0x7c, 0xfb, 0xc7, 0x1a, 0x5c, 0xe7, 0xfb, 0x4a, 0xf0, 0x07, 0xa8, 0xd1, 0x4b, 0x66,
	0x3a, 0x5a, 0x36, 0xd3, 0x31, 0xd9, 0xc4, 0x8f, 0xf6, 0x04, 0xd7, 0x12, 0x97, 0xaf, 0xdd, 0x27,
	0x3e, 0xe6, 0xf5, 0x23, 0xf9, 0xcb, 0x9b, 0x6f, 0xfd, 0xd5, 0xcd, 0xff, 0xf7, 0xbb, 0x9b, 0xbf,
	0x60, 0x39, 0x56, 0xff, 0x74, 0x39, 0xd6, 0x2e, 0x5a, 0x8e, 0xc1, 0x1a, 0xfe, 0x2e, 0xf7, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x43, 0x00, 0xef, 0x8b, 0x48, 0x05, 0x00, 0x00,
}
