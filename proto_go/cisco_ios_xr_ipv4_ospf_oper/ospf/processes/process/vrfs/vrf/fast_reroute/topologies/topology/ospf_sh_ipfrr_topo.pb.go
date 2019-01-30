// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_ipfrr_topo.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_fast_reroute_topologies_topology is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_ipfrr_topo.proto

It has these top-level messages:
	OspfShIpfrrTopo_KEYS
	OspfShIpfrrTopo
	OspfShIpfrrTopoEntry
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_fast_reroute_topologies_topology

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
	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName     string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	RouterId    string `protobuf:"bytes,3,opt,name=router_id,json=routerId" json:"router_id,omitempty"`
	AreaId      uint32 `protobuf:"varint,4,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
}

func (m *OspfShIpfrrTopo_KEYS) Reset()                    { *m = OspfShIpfrrTopo_KEYS{} }
func (m *OspfShIpfrrTopo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo_KEYS) ProtoMessage()               {}
func (*OspfShIpfrrTopo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShIpfrrTopo_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	IpfrrTopoAreaId string `protobuf:"bytes,50,opt,name=ipfrr_topo_area_id,json=ipfrrTopoAreaId" json:"ipfrr_topo_area_id,omitempty"`
	// OSPF Router ID
	IpfrrRouterId string `protobuf:"bytes,51,opt,name=ipfrr_router_id,json=ipfrrRouterId" json:"ipfrr_router_id,omitempty"`
	// IPFRR Topology Revision
	IpfrrAreaRevision uint32 `protobuf:"varint,52,opt,name=ipfrr_area_revision,json=ipfrrAreaRevision" json:"ipfrr_area_revision,omitempty"`
	// IPFRR Topology entries
	IpfrrTopo []*OspfShIpfrrTopoEntry `protobuf:"bytes,53,rep,name=ipfrr_topo,json=ipfrrTopo" json:"ipfrr_topo,omitempty"`
}

func (m *OspfShIpfrrTopo) Reset()                    { *m = OspfShIpfrrTopo{} }
func (m *OspfShIpfrrTopo) String() string            { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo) ProtoMessage()               {}
func (*OspfShIpfrrTopo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	// IPFRR Topology Distance
	Distance uint32 `protobuf:"varint,2,opt,name=distance" json:"distance,omitempty"`
	// IPFRR Topology Reverse Distance
	DistanceReverse uint32 `protobuf:"varint,3,opt,name=distance_reverse,json=distanceReverse" json:"distance_reverse,omitempty"`
	// IPFRR Topoogy Type-4 entry
	Type4 bool `protobuf:"varint,4,opt,name=type4" json:"type4,omitempty"`
	// IPFRR Topology Revision
	Revision uint32 `protobuf:"varint,5,opt,name=revision" json:"revision,omitempty"`
	// IPFRR Topology Neighbor Sourced
	NeighborSourced bool `protobuf:"varint,6,opt,name=neighbor_sourced,json=neighborSourced" json:"neighbor_sourced,omitempty"`
	// IPFRR Topology DR entry
	Dr bool `protobuf:"varint,7,opt,name=dr" json:"dr,omitempty"`
	// IPFRR Topology rSPT poison
	Poison bool `protobuf:"varint,8,opt,name=poison" json:"poison,omitempty"`
}

func (m *OspfShIpfrrTopoEntry) Reset()                    { *m = OspfShIpfrrTopoEntry{} }
func (m *OspfShIpfrrTopoEntry) String() string            { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopoEntry) ProtoMessage()               {}
func (*OspfShIpfrrTopoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	proto.RegisterType((*OspfShIpfrrTopo_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_KEYS")
	proto.RegisterType((*OspfShIpfrrTopo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo")
	proto.RegisterType((*OspfShIpfrrTopoEntry)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_entry")
}

func init() { proto.RegisterFile("ospf_sh_ipfrr_topo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0xea, 0x3a, 0x53, 0x42, 0x60, 0x41, 0xd4, 0xc0, 0x25, 0xe4, 0x80, 0x82, 0x90,
	0x7c, 0x68, 0xc3, 0x07, 0x70, 0xe0, 0x10, 0x21, 0x71, 0xd8, 0x72, 0x41, 0x42, 0x5a, 0xb9, 0xde,
	0x71, 0xbb, 0x12, 0xf5, 0xac, 0x66, 0x5d, 0x8b, 0x7c, 0x04, 0x47, 0x3e, 0x84, 0x2f, 0x04, 0xed,
	0x6c, 0x9c, 0x22, 0xb5, 0x57, 0x2e, 0xd6, 0xbc, 0xf7, 0x66, 0xdf, 0x8e, 0xe7, 0xd9, 0x50, 0x52,
	0xf0, 0xad, 0x09, 0x57, 0xc6, 0xf9, 0x96, 0xd9, 0xf4, 0xe4, 0xa9, 0xf2, 0x4c, 0x3d, 0xa9, 0x6f,
	0x8d, 0x0b, 0x0d, 0x19, 0x47, 0xc1, 0xfc, 0x60, 0xe3, 0xfc, 0xb0, 0x36, 0xd2, 0x4b, 0x1e, 0xb9,
	0x8a, 0x55, 0xec, 0x6b, 0x30, 0x04, 0x0c, 0x63, 0x55, 0x0d, 0xdc, 0xca, 0xa3, 0x6a, 0xeb, 0xd0,
	0x1b, 0x46, 0xa6, 0x9b, 0x1e, 0xab, 0x68, 0xfa, 0x9d, 0x2e, 0x1d, 0x86, 0xb1, 0xdc, 0x2e, 0x7f,
	0x66, 0x70, 0x72, 0xf7, 0x6a, 0xf3, 0xe9, 0xe3, 0xd7, 0x73, 0xf5, 0x1a, 0x1e, 0xee, 0x0c, 0x4d,
	0x57, 0x5f, 0x63, 0x99, 0x2d, 0xb2, 0xd5, 0x54, 0x1f, 0xef, 0xb8, 0xcf, 0xf5, 0x35, 0xaa, 0x17,
	0x50, 0x0c, 0xdc, 0x26, 0x79, 0x22, 0xf2, 0xd1, 0xc0, 0xad, 0x48, 0xaf, 0x60, 0x2a, 0xd7, 0xb2,
	0x71, 0xb6, 0x3c, 0x10, 0xad, 0x48, 0xc4, 0xc6, 0xaa, 0x13, 0x38, 0xaa, 0x19, 0xeb, 0x28, 0x3d,
	0x58, 0x64, 0xab, 0x99, 0xce, 0x23, 0xdc, 0xd8, 0xe5, 0xef, 0x09, 0xa8, 0xbb, 0xf3, 0xa8, 0x77,
	0xa0, 0xfe, 0x99, 0x6e, 0x3c, 0x7a, 0x2a, 0xae, 0x73, 0x51, 0xbe, 0x90, 0xa7, 0x0f, 0xe2, 0xa1,
	0xde, 0x40, 0xa2, 0xcc, 0xed, 0xfd, 0x67, 0xd2, 0x39, 0x13, 0x5a, 0x8f, 0x43, 0x54, 0xf0, 0x34,
	0xf5, 0x89, 0x1f, 0xe3, 0xe0, 0x82, 0xa3, 0xae, 0x5c, 0xcb, 0x40, 0x4f, 0x44, 0x8a, 0x8e, 0x7a,
	0x27, 0xa8, 0x5f, 0x19, 0xc0, 0xed, 0x14, 0xe5, 0xfb, 0xc5, 0xc1, 0xea, 0xf8, 0x74, 0xa8, 0xfe,
	0x67, 0x3e, 0xd5, 0x3d, 0xd9, 0x60, 0xd7, 0xf3, 0x56, 0x4f, 0xf7, 0x6f, 0xbd, 0xfc, 0x93, 0xdd,
	0xf7, 0xf9, 0xa4, 0xbe, 0xb8, 0xe9, 0x8e, 0x2c, 0xc6, 0x25, 0xa4, 0xfc, 0xf2, 0x08, 0x37, 0x56,
	0xbd, 0x84, 0xc2, 0xba, 0xd0, 0xd7, 0x5d, 0x93, 0xa2, 0x9b, 0xe9, 0x3d, 0x56, 0x6f, 0xe1, 0xf1,
	0x58, 0xc7, 0xbd, 0x20, 0x07, 0x94, 0x08, 0x67, 0x7a, 0x3e, 0xf2, 0x3a, 0xd1, 0xea, 0x19, 0x1c,
	0xf6, 0x5b, 0x8f, 0x6b, 0xc9, 0xb1, 0xd0, 0x09, 0x44, 0xf3, 0xfd, 0x3e, 0x0f, 0x93, 0xf9, 0x88,
	0xa3, 0x79, 0x87, 0xee, 0xf2, 0xea, 0x82, 0xd8, 0x04, 0xba, 0xe1, 0x06, 0x6d, 0x99, 0xcb, 0xe1,
	0xf9, 0xc8, 0x9f, 0x27, 0x5a, 0x3d, 0x82, 0x89, 0xe5, 0xf2, 0x48, 0xc4, 0x89, 0x65, 0xf5, 0x1c,
	0x72, 0x4f, 0x2e, 0x50, 0x57, 0x16, 0xc2, 0xed, 0xd0, 0x45, 0x2e, 0xbf, 0xca, 0xd9, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x42, 0xcf, 0xfd, 0x46, 0x46, 0x03, 0x00, 0x00,
}
