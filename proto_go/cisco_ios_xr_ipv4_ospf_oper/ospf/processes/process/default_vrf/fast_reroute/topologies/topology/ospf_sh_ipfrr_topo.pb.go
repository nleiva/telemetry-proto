// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_ipfrr_topo.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_fast_reroute_topologies_topology

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

// OSPF IPFRR Topology Information
type OspfShIpfrrTopo_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	RouterId             string   `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIpfrrTopo_KEYS) Reset()         { *m = OspfShIpfrrTopo_KEYS{} }
func (m *OspfShIpfrrTopo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo_KEYS) ProtoMessage()    {}
func (*OspfShIpfrrTopo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_ipfrr_topo_d43eda00d54d9db0, []int{0}
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShIpfrrTopo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopo_KEYS.Merge(dst, src)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Size(m)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopo_KEYS proto.InternalMessageInfo

func (m *OspfShIpfrrTopo_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type OspfShIpfrrTopo struct {
	// Area ID string in decimal or dotted decimal format
	IpfrrTopoAreaId string `protobuf:"bytes,50,opt,name=ipfrr_topo_area_id,json=ipfrrTopoAreaId,proto3" json:"ipfrr_topo_area_id,omitempty"`
	// OSPF Router ID
	IpfrrRouterId string `protobuf:"bytes,51,opt,name=ipfrr_router_id,json=ipfrrRouterId,proto3" json:"ipfrr_router_id,omitempty"`
	// IPFRR Topology Revision
	IpfrrAreaRevision uint32 `protobuf:"varint,52,opt,name=ipfrr_area_revision,json=ipfrrAreaRevision,proto3" json:"ipfrr_area_revision,omitempty"`
	// IPFRR Topology entries
	IpfrrTopo            []*OspfShIpfrrTopoEntry `protobuf:"bytes,53,rep,name=ipfrr_topo,json=ipfrrTopo,proto3" json:"ipfrr_topo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OspfShIpfrrTopo) Reset()         { *m = OspfShIpfrrTopo{} }
func (m *OspfShIpfrrTopo) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo) ProtoMessage()    {}
func (*OspfShIpfrrTopo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_ipfrr_topo_d43eda00d54d9db0, []int{1}
}
func (m *OspfShIpfrrTopo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopo.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopo.Marshal(b, m, deterministic)
}
func (dst *OspfShIpfrrTopo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopo.Merge(dst, src)
}
func (m *OspfShIpfrrTopo) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopo.Size(m)
}
func (m *OspfShIpfrrTopo) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopo.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopo proto.InternalMessageInfo

func (m *OspfShIpfrrTopo) GetIpfrrTopoAreaId() string {
	if m != nil {
		return m.IpfrrTopoAreaId
	}
	return ""
}

func (m *OspfShIpfrrTopo) GetIpfrrRouterId() string {
	if m != nil {
		return m.IpfrrRouterId
	}
	return ""
}

func (m *OspfShIpfrrTopo) GetIpfrrAreaRevision() uint32 {
	if m != nil {
		return m.IpfrrAreaRevision
	}
	return 0
}

func (m *OspfShIpfrrTopo) GetIpfrrTopo() []*OspfShIpfrrTopoEntry {
	if m != nil {
		return m.IpfrrTopo
	}
	return nil
}

// OSPF_IPFRR Topology Entry
type OspfShIpfrrTopoEntry struct {
	// IPFRR Topology Node ID
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// IPFRR Topology Distance
	Distance uint32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// IPFRR Topology Reverse Distance
	DistanceReverse uint32 `protobuf:"varint,3,opt,name=distance_reverse,json=distanceReverse,proto3" json:"distance_reverse,omitempty"`
	// IPFRR Topoogy Type-4 entry
	Type4 bool `protobuf:"varint,4,opt,name=type4,proto3" json:"type4,omitempty"`
	// IPFRR Topology Revision
	Revision uint32 `protobuf:"varint,5,opt,name=revision,proto3" json:"revision,omitempty"`
	// IPFRR Topology Neighbor Sourced
	NeighborSourced bool `protobuf:"varint,6,opt,name=neighbor_sourced,json=neighborSourced,proto3" json:"neighbor_sourced,omitempty"`
	// IPFRR Topology DR entry
	Dr bool `protobuf:"varint,7,opt,name=dr,proto3" json:"dr,omitempty"`
	// IPFRR Topology rSPT poison
	Poison               bool     `protobuf:"varint,8,opt,name=poison,proto3" json:"poison,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIpfrrTopoEntry) Reset()         { *m = OspfShIpfrrTopoEntry{} }
func (m *OspfShIpfrrTopoEntry) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopoEntry) ProtoMessage()    {}
func (*OspfShIpfrrTopoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_ipfrr_topo_d43eda00d54d9db0, []int{2}
}
func (m *OspfShIpfrrTopoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Marshal(b, m, deterministic)
}
func (dst *OspfShIpfrrTopoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopoEntry.Merge(dst, src)
}
func (m *OspfShIpfrrTopoEntry) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Size(m)
}
func (m *OspfShIpfrrTopoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopoEntry proto.InternalMessageInfo

func (m *OspfShIpfrrTopoEntry) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *OspfShIpfrrTopoEntry) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetDistanceReverse() uint32 {
	if m != nil {
		return m.DistanceReverse
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetType4() bool {
	if m != nil {
		return m.Type4
	}
	return false
}

func (m *OspfShIpfrrTopoEntry) GetRevision() uint32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetNeighborSourced() bool {
	if m != nil {
		return m.NeighborSourced
	}
	return false
}

func (m *OspfShIpfrrTopoEntry) GetDr() bool {
	if m != nil {
		return m.Dr
	}
	return false
}

func (m *OspfShIpfrrTopoEntry) GetPoison() bool {
	if m != nil {
		return m.Poison
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShIpfrrTopo_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_KEYS")
	proto.RegisterType((*OspfShIpfrrTopo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo")
	proto.RegisterType((*OspfShIpfrrTopoEntry)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_entry")
}

func init() {
	proto.RegisterFile("ospf_sh_ipfrr_topo.proto", fileDescriptor_ospf_sh_ipfrr_topo_d43eda00d54d9db0)
}

var fileDescriptor_ospf_sh_ipfrr_topo_d43eda00d54d9db0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0xea, 0x26, 0x53, 0x42, 0x60, 0x41, 0x74, 0x05, 0x97, 0x90, 0x03, 0x0a, 0x42,
	0xf2, 0xa1, 0x0d, 0x1f, 0xc0, 0x81, 0x43, 0x84, 0xc4, 0xc1, 0xe5, 0xc2, 0x69, 0xe5, 0x7a, 0xc7,
	0xed, 0x4a, 0xad, 0x67, 0x35, 0xeb, 0x44, 0xf8, 0x4f, 0xf8, 0x10, 0xfe, 0x0f, 0xb4, 0xb3, 0x76,
	0x8a, 0xd4, 0x5e, 0xb9, 0xcd, 0x7b, 0x6f, 0xf6, 0xcd, 0x64, 0x5e, 0x0c, 0x9a, 0x82, 0x6f, 0x4c,
	0xb8, 0x31, 0xce, 0x37, 0xcc, 0xa6, 0x23, 0x4f, 0x85, 0x67, 0xea, 0x48, 0x99, 0xda, 0x85, 0x9a,
	0x8c, 0xa3, 0x60, 0x7e, 0xb2, 0x71, 0x7e, 0xbf, 0x31, 0xd2, 0x4b, 0x1e, 0xb9, 0x88, 0x55, 0xec,
	0xab, 0x31, 0x04, 0x0c, 0x63, 0x55, 0x58, 0x6c, 0xaa, 0xdd, 0x6d, 0x67, 0xf6, 0xdc, 0x14, 0x4d,
	0x15, 0x3a, 0xc3, 0xc8, 0xb4, 0xeb, 0xb0, 0x88, 0xbe, 0xb7, 0x74, 0xed, 0x30, 0x8c, 0x65, 0xbf,
	0x62, 0x38, 0x7b, 0x38, 0xdc, 0x7c, 0xfd, 0xf2, 0xe3, 0x52, 0xbd, 0x83, 0xa7, 0x83, 0xa5, 0x69,
	0xab, 0x3b, 0xd4, 0xd9, 0x32, 0x5b, 0xcf, 0xca, 0xd3, 0x81, 0xfb, 0x56, 0xdd, 0xa1, 0x7a, 0x0b,
	0x33, 0xb1, 0x66, 0xe3, 0xac, 0x9e, 0x88, 0x3e, 0x4d, 0xc4, 0xd6, 0xaa, 0x33, 0x38, 0xa9, 0x18,
	0xab, 0x28, 0x1d, 0x2d, 0xb3, 0xf5, 0xbc, 0xcc, 0x23, 0xdc, 0xda, 0xd5, 0xef, 0x09, 0xa8, 0x87,
	0x43, 0xd5, 0x47, 0x50, 0xff, 0xac, 0x30, 0x3e, 0x3d, 0x17, 0xd7, 0x85, 0x28, 0xdf, 0xc9, 0xd3,
	0x67, 0xf1, 0x50, 0xef, 0x21, 0x51, 0xe6, 0x7e, 0xfe, 0x85, 0x74, 0xce, 0x85, 0x2e, 0xc7, 0x25,
	0x0a, 0x78, 0x99, 0xfa, 0xc4, 0x8f, 0x71, 0xef, 0x82, 0xa3, 0x56, 0x6f, 0x64, 0xa1, 0x17, 0x22,
	0x45, 0xc7, 0x72, 0x10, 0xd4, 0xaf, 0x0c, 0xe0, 0x7e, 0x0b, 0xfd, 0x69, 0x79, 0xb4, 0x3e, 0x3d,
	0xef, 0x8b, 0xff, 0x1c, 0x43, 0xf1, 0x48, 0x06, 0xd8, 0x76, 0xdc, 0x97, 0xb3, 0xc3, 0x0f, 0x5f,
	0xfd, 0xc9, 0x1e, 0xfb, 0xa3, 0xa4, 0xbe, 0x78, 0xec, 0x96, 0x2c, 0xc6, 0x3b, 0xa4, 0x9c, 0xf2,
	0x08, 0xb7, 0x56, 0xbd, 0x81, 0xa9, 0x75, 0xa1, 0xab, 0xda, 0x1a, 0x25, 0xa1, 0x79, 0x79, 0xc0,
	0xea, 0x03, 0x3c, 0x1f, 0xeb, 0x78, 0x1a, 0xe4, 0x80, 0x43, 0x54, 0x8b, 0x91, 0x2f, 0x13, 0xad,
	0x5e, 0xc1, 0x71, 0xd7, 0x7b, 0xdc, 0xe8, 0x27, 0xcb, 0x6c, 0x3d, 0x2d, 0x13, 0x88, 0xe6, 0x87,
	0x93, 0x1e, 0x27, 0xf3, 0x11, 0x47, 0xf3, 0x16, 0xdd, 0xf5, 0xcd, 0x15, 0xb1, 0x09, 0xb4, 0xe3,
	0x1a, 0xad, 0xce, 0xe5, 0xf1, 0x62, 0xe4, 0x2f, 0x13, 0xad, 0x9e, 0xc1, 0xc4, 0xb2, 0x3e, 0x11,
	0x71, 0x62, 0x59, 0xbd, 0x86, 0xdc, 0x93, 0x0b, 0xd4, 0xea, 0xa9, 0x70, 0x03, 0xba, 0xca, 0xe5,
	0xa3, 0xb8, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x82, 0xd5, 0x4d, 0x30, 0x03, 0x00, 0x00,
}
