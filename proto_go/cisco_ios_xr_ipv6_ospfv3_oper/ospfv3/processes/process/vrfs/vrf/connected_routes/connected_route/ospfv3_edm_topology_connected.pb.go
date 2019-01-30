// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_topology_connected.proto

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_connected_routes_connected_route is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_topology_connected.proto

It has these top-level messages:
	Ospfv3EdmTopologyConnected_KEYS
	Ospfv3EdmTopologyConnected
	Ospfv3ShBackupPath
	Ospfv3EdmTopPath
	Ospfv3EdmTopology
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_connected_routes_connected_route

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

// OSPFv3 Topology Connected Information
type Ospfv3EdmTopologyConnected_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	Prefix       string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ospfv3EdmTopologyConnected_KEYS) Reset()                    { *m = Ospfv3EdmTopologyConnected_KEYS{} }
func (m *Ospfv3EdmTopologyConnected_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyConnected_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmTopologyConnected_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmTopologyConnected_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmTopologyConnected_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmTopologyConnected_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ospfv3EdmTopologyConnected_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ospfv3EdmTopologyConnected struct {
	// Common Route topology information
	RouteTopology *Ospfv3EdmTopology `protobuf:"bytes,50,opt,name=route_topology,json=routeTopology" json:"route_topology,omitempty"`
	// List of paths to this route
	RoutePathList []*Ospfv3EdmTopPath `protobuf:"bytes,51,rep,name=route_path_list,json=routePathList" json:"route_path_list,omitempty"`
}

func (m *Ospfv3EdmTopologyConnected) Reset()                    { *m = Ospfv3EdmTopologyConnected{} }
func (m *Ospfv3EdmTopologyConnected) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyConnected) ProtoMessage()               {}
func (*Ospfv3EdmTopologyConnected) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmTopologyConnected) GetRouteTopology() *Ospfv3EdmTopology {
	if m != nil {
		return m.RouteTopology
	}
	return nil
}

func (m *Ospfv3EdmTopologyConnected) GetRoutePathList() []*Ospfv3EdmTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

// OSPFv3 Route Backup Path Information
type Ospfv3ShBackupPath struct {
	// Next hop Interface
	BackupRouteInterfaceName string `protobuf:"bytes,1,opt,name=backup_route_interface_name,json=backupRouteInterfaceName" json:"backup_route_interface_name,omitempty"`
	// Nexthop IP address
	BackupRouteNextHopAddress string `protobuf:"bytes,2,opt,name=backup_route_next_hop_address,json=backupRouteNextHopAddress" json:"backup_route_next_hop_address,omitempty"`
	// IP address of source of route
	BackupRouteSource string `protobuf:"bytes,3,opt,name=backup_route_source,json=backupRouteSource" json:"backup_route_source,omitempty"`
	// Metric
	BackupMetric uint32 `protobuf:"varint,4,opt,name=backup_metric,json=backupMetric" json:"backup_metric,omitempty"`
	// Primary Path
	PrimaryPath bool `protobuf:"varint,5,opt,name=primary_path,json=primaryPath" json:"primary_path,omitempty"`
	// Line Card Disjoint
	LineCardDisjoint bool `protobuf:"varint,6,opt,name=line_card_disjoint,json=lineCardDisjoint" json:"line_card_disjoint,omitempty"`
	// Downstream
	Downstream bool `protobuf:"varint,7,opt,name=downstream" json:"downstream,omitempty"`
	// Node Protect
	NodeProtect bool `protobuf:"varint,8,opt,name=node_protect,json=nodeProtect" json:"node_protect,omitempty"`
	// SRLG Disjoint
	SrlgDisjoint bool `protobuf:"varint,9,opt,name=srlg_disjoint,json=srlgDisjoint" json:"srlg_disjoint,omitempty"`
}

func (m *Ospfv3ShBackupPath) Reset()                    { *m = Ospfv3ShBackupPath{} }
func (m *Ospfv3ShBackupPath) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3ShBackupPath) ProtoMessage()               {}
func (*Ospfv3ShBackupPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3ShBackupPath) GetBackupRouteInterfaceName() string {
	if m != nil {
		return m.BackupRouteInterfaceName
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupRouteNextHopAddress() string {
	if m != nil {
		return m.BackupRouteNextHopAddress
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupRouteSource() string {
	if m != nil {
		return m.BackupRouteSource
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupMetric() uint32 {
	if m != nil {
		return m.BackupMetric
	}
	return 0
}

func (m *Ospfv3ShBackupPath) GetPrimaryPath() bool {
	if m != nil {
		return m.PrimaryPath
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetLineCardDisjoint() bool {
	if m != nil {
		return m.LineCardDisjoint
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetNodeProtect() bool {
	if m != nil {
		return m.NodeProtect
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetSrlgDisjoint() bool {
	if m != nil {
		return m.SrlgDisjoint
	}
	return false
}

// OSPFv3 topology path information
type Ospfv3EdmTopPath struct {
	// Route path interface name
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Route path next hop
	RoutePathNextHop string `protobuf:"bytes,2,opt,name=route_path_next_hop,json=routePathNextHop" json:"route_path_next_hop,omitempty"`
	// Path ID of path
	RoutePathId uint32 `protobuf:"varint,3,opt,name=route_path_id,json=routePathId" json:"route_path_id,omitempty"`
	// Backup Path Info
	RouteBackupPath *Ospfv3ShBackupPath `protobuf:"bytes,4,opt,name=route_backup_path,json=routeBackupPath" json:"route_backup_path,omitempty"`
}

func (m *Ospfv3EdmTopPath) Reset()                    { *m = Ospfv3EdmTopPath{} }
func (m *Ospfv3EdmTopPath) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopPath) ProtoMessage()               {}
func (*Ospfv3EdmTopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ospfv3EdmTopPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmTopPath) GetRoutePathNextHop() string {
	if m != nil {
		return m.RoutePathNextHop
	}
	return ""
}

func (m *Ospfv3EdmTopPath) GetRoutePathId() uint32 {
	if m != nil {
		return m.RoutePathId
	}
	return 0
}

func (m *Ospfv3EdmTopPath) GetRouteBackupPath() *Ospfv3ShBackupPath {
	if m != nil {
		return m.RouteBackupPath
	}
	return nil
}

// OSPFv3 Topology Information
type Ospfv3EdmTopology struct {
	// Route ID
	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId" json:"route_id,omitempty"`
	// Route distance
	RouteDistance uint32 `protobuf:"varint,2,opt,name=route_distance,json=routeDistance" json:"route_distance,omitempty"`
	// Route cost
	RouteCost uint32 `protobuf:"varint,3,opt,name=route_cost,json=routeCost" json:"route_cost,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,4,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
}

func (m *Ospfv3EdmTopology) Reset()                    { *m = Ospfv3EdmTopology{} }
func (m *Ospfv3EdmTopology) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopology) ProtoMessage()               {}
func (*Ospfv3EdmTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ospfv3EdmTopology) GetRouteId() string {
	if m != nil {
		return m.RouteId
	}
	return ""
}

func (m *Ospfv3EdmTopology) GetRouteDistance() uint32 {
	if m != nil {
		return m.RouteDistance
	}
	return 0
}

func (m *Ospfv3EdmTopology) GetRouteCost() uint32 {
	if m != nil {
		return m.RouteCost
	}
	return 0
}

func (m *Ospfv3EdmTopology) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmTopologyConnected_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.connected_routes.connected_route.ospfv3_edm_topology_connected_KEYS")
	proto.RegisterType((*Ospfv3EdmTopologyConnected)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.connected_routes.connected_route.ospfv3_edm_topology_connected")
	proto.RegisterType((*Ospfv3ShBackupPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.connected_routes.connected_route.ospfv3_sh_backup_path")
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.connected_routes.connected_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3EdmTopology)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.connected_routes.connected_route.ospfv3_edm_topology")
}

func init() { proto.RegisterFile("ospfv3_edm_topology_connected.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4f, 0x4f, 0x14, 0x31,
	0x1c, 0xcd, 0x80, 0xc2, 0xf2, 0x5b, 0x06, 0xa1, 0x44, 0x33, 0xc4, 0x60, 0xd6, 0x21, 0x26, 0x7b,
	0xd0, 0x39, 0x40, 0xe2, 0xcd, 0x44, 0x05, 0x13, 0x89, 0x48, 0xc8, 0xc2, 0xc5, 0x53, 0x2d, 0xed,
	0x6f, 0xd9, 0xea, 0xee, 0x74, 0xd2, 0x76, 0x97, 0xdd, 0x0f, 0xe2, 0x4d, 0xe3, 0xc1, 0xc4, 0x6f,
	0xe0, 0xd7, 0xf1, 0xb3, 0x98, 0xfe, 0xd9, 0x65, 0x96, 0x10, 0x6e, 0xea, 0x65, 0xd2, 0xbe, 0xf7,
	0xda, 0xbe, 0xf9, 0xfd, 0x5e, 0x0b, 0x3b, 0xca, 0x54, 0xdd, 0xd1, 0x1e, 0x45, 0x31, 0xa0, 0x56,
	0x55, 0xaa, 0xaf, 0x2e, 0x26, 0x94, 0xab, 0xb2, 0x44, 0x6e, 0x51, 0x14, 0x95, 0x56, 0x56, 0x91,
	0x8f, 0x5c, 0x1a, 0xae, 0xa8, 0x54, 0x86, 0x8e, 0x35, 0x95, 0xd5, 0xe8, 0x39, 0x8d, 0xcb, 0x54,
	0x85, 0xba, 0x08, 0x63, 0xa7, 0xe5, 0x68, 0x0c, 0x9a, 0xe9, 0xa8, 0x18, 0xe9, 0xae, 0xff, 0x14,
	0xb3, 0x0d, 0xa9, 0x56, 0x43, 0x8b, 0xe6, 0x3a, 0x90, 0x7f, 0x4f, 0x20, 0xbf, 0xd5, 0x09, 0x7d,
	0xf7, 0xe6, 0xc3, 0x29, 0x79, 0x0c, 0xab, 0x71, 0x6f, 0x5a, 0xb2, 0x01, 0x66, 0x49, 0x2b, 0x69,
	0xaf, 0x74, 0x9a, 0x11, 0x3b, 0x66, 0x03, 0x24, 0x5b, 0xd0, 0x18, 0xe9, 0x6e, 0xa0, 0x17, 0x3c,
	0xbd, 0x3c, 0xd2, 0x5d, 0x4f, 0x3d, 0x80, 0xa5, 0x4a, 0x63, 0x57, 0x8e, 0xb3, 0x45, 0x4f, 0xc4,
	0x19, 0xd9, 0x81, 0x34, 0x8c, 0x68, 0x1f, 0xcb, 0x0b, 0xdb, 0xcb, 0xee, 0xb4, 0x92, 0x76, 0xda,
	0x59, 0x0d, 0xe0, 0x91, 0xc7, 0xf2, 0xdf, 0x0b, 0xb0, 0x7d, 0xab, 0x43, 0xf2, 0x35, 0x81, 0x35,
	0xff, 0x37, 0x33, 0x32, 0xdb, 0x6d, 0x25, 0xed, 0xe6, 0xee, 0xb0, 0xf8, 0xdb, 0xf5, 0x2b, 0x6e,
	0x70, 0xd6, 0x49, 0x3d, 0x75, 0x16, 0xa7, 0xe4, 0x5b, 0x02, 0xf7, 0x82, 0xbd, 0x8a, 0xd9, 0x1e,
	0xed, 0x4b, 0x63, 0xb3, 0xbd, 0xd6, 0xe2, 0xff, 0xf0, 0xe7, 0x1d, 0x44, 0x7f, 0x27, 0xcc, 0xf6,
	0x8e, 0xa4, 0xb1, 0xf9, 0xcf, 0x45, 0xb8, 0x1f, 0x65, 0xa6, 0x47, 0xcf, 0x19, 0xff, 0x3c, 0x0c,
	0x42, 0xf2, 0x02, 0x1e, 0xc6, 0x69, 0xf0, 0x2f, 0x4b, 0x8b, 0xba, 0xcb, 0x38, 0xd6, 0x43, 0x90,
	0x05, 0x49, 0xc7, 0x29, 0x0e, 0xa7, 0x02, 0xdf, 0xf6, 0x97, 0xb0, 0x3d, 0xb7, 0xbc, 0xc4, 0xb1,
	0xa5, 0x3d, 0x55, 0x51, 0x26, 0x84, 0x46, 0x63, 0x62, 0x4c, 0xb6, 0x6a, 0x1b, 0x1c, 0xe3, 0xd8,
	0xbe, 0x55, 0xd5, 0xab, 0x20, 0x20, 0x05, 0x6c, 0xce, 0xed, 0x60, 0xd4, 0x50, 0x73, 0x8c, 0x29,
	0xda, 0xa8, 0xad, 0x3b, 0xf5, 0x84, 0x0b, 0x54, 0xd4, 0x0f, 0xd0, 0x6a, 0xc9, 0xa7, 0x81, 0x0a,
	0xe0, 0x7b, 0x8f, 0x85, 0x2c, 0xcb, 0x01, 0xd3, 0x13, 0xff, 0x97, 0xd9, 0xdd, 0x56, 0xd2, 0x6e,
	0xb8, 0x2c, 0x7b, 0xcc, 0x95, 0x85, 0x3c, 0x05, 0xd2, 0x97, 0x25, 0x52, 0xce, 0xb4, 0xa0, 0x42,
	0x9a, 0x4f, 0x4a, 0x96, 0x36, 0x5b, 0xf2, 0xc2, 0x75, 0xc7, 0xec, 0x33, 0x2d, 0x0e, 0x22, 0x4e,
	0x1e, 0x01, 0x08, 0x75, 0x59, 0x1a, 0xab, 0x91, 0x0d, 0xb2, 0x65, 0xaf, 0xaa, 0x21, 0xee, 0xc0,
	0x52, 0x09, 0xa4, 0xee, 0x4e, 0x23, 0xb7, 0x59, 0x23, 0x1c, 0xe8, 0xb0, 0x93, 0x00, 0x39, 0xe3,
	0x46, 0xf7, 0x2f, 0xae, 0xce, 0x5a, 0xf1, 0x9a, 0x55, 0x07, 0x4e, 0xcf, 0xc9, 0x7f, 0x2d, 0xc0,
	0xe6, 0x0d, 0xfd, 0x24, 0x4f, 0x60, 0xed, 0xc6, 0xce, 0xa4, 0x72, 0xae, 0x1d, 0xcf, 0x60, 0xb3,
	0x16, 0xc3, 0x69, 0x33, 0x62, 0x13, 0xd6, 0x67, 0x99, 0x88, 0x2d, 0x20, 0x39, 0xa4, 0x35, 0xb9,
	0x14, 0xbe, 0xea, 0x69, 0xa7, 0x39, 0x13, 0x1e, 0x0a, 0xf2, 0x23, 0x81, 0x8d, 0x20, 0xaa, 0xc5,
	0xc6, 0x17, 0xbd, 0xb9, 0x7b, 0xf9, 0xcf, 0xc2, 0x3d, 0x9f, 0xda, 0x4e, 0xb8, 0x6c, 0xaf, 0x3d,
	0xe2, 0x7c, 0xe6, 0x5f, 0x92, 0xeb, 0x75, 0x0b, 0x17, 0x73, 0x0b, 0x1a, 0x31, 0xd7, 0x22, 0x56,
	0x6c, 0xd9, 0xcf, 0x0f, 0x85, 0x2b, 0x69, 0xa0, 0x84, 0x34, 0x96, 0x95, 0x3c, 0x3c, 0x69, 0x69,
	0xbc, 0x3a, 0x07, 0x11, 0x24, 0xdb, 0x00, 0x41, 0xc6, 0x95, 0xb1, 0xb1, 0x40, 0x2b, 0x1e, 0xd9,
	0x57, 0xc6, 0x5e, 0xd1, 0x76, 0x52, 0x61, 0xcc, 0x62, 0xa0, 0xcf, 0x26, 0x15, 0x9e, 0x2f, 0xf9,
	0x47, 0x7e, 0xef, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x03, 0xd4, 0xe9, 0x0b, 0x06, 0x00,
	0x00,
}
