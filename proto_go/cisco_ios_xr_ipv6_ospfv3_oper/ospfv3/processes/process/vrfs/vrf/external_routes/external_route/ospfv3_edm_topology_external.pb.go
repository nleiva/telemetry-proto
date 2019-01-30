// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_topology_external.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_external_routes_external_route

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

// OSPFv3 Topology External Information
type Ospfv3EdmTopologyExternal_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopologyExternal_KEYS) Reset()         { *m = Ospfv3EdmTopologyExternal_KEYS{} }
func (m *Ospfv3EdmTopologyExternal_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyExternal_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmTopologyExternal_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{0}
}

func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Size(m)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyExternal_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ospfv3EdmTopologyExternal struct {
	// Common Route topology information
	RouteTopology *Ospfv3EdmTopology `protobuf:"bytes,50,opt,name=route_topology,json=routeTopology,proto3" json:"route_topology,omitempty"`
	// List of paths to this route
	RoutePathList []*Ospfv3EdmTopPath `protobuf:"bytes,51,rep,name=route_path_list,json=routePathList,proto3" json:"route_path_list,omitempty"`
	// Extended communities in the route
	RouteExtendedCommunity *Ospfv3ShRouteExtendedComm `protobuf:"bytes,52,opt,name=route_extended_community,json=routeExtendedCommunity,proto3" json:"route_extended_community,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                   `json:"-"`
	XXX_unrecognized       []byte                     `json:"-"`
	XXX_sizecache          int32                      `json:"-"`
}

func (m *Ospfv3EdmTopologyExternal) Reset()         { *m = Ospfv3EdmTopologyExternal{} }
func (m *Ospfv3EdmTopologyExternal) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyExternal) ProtoMessage()    {}
func (*Ospfv3EdmTopologyExternal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{1}
}

func (m *Ospfv3EdmTopologyExternal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyExternal.Merge(m, src)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Size(m)
}
func (m *Ospfv3EdmTopologyExternal) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyExternal.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyExternal proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyExternal) GetRouteTopology() *Ospfv3EdmTopology {
	if m != nil {
		return m.RouteTopology
	}
	return nil
}

func (m *Ospfv3EdmTopologyExternal) GetRoutePathList() []*Ospfv3EdmTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

func (m *Ospfv3EdmTopologyExternal) GetRouteExtendedCommunity() *Ospfv3ShRouteExtendedComm {
	if m != nil {
		return m.RouteExtendedCommunity
	}
	return nil
}

// OSPFv3 Route Backup Path Information
type Ospfv3ShBackupPath struct {
	// Next hop Interface
	BackupRouteInterfaceName string `protobuf:"bytes,1,opt,name=backup_route_interface_name,json=backupRouteInterfaceName,proto3" json:"backup_route_interface_name,omitempty"`
	// Nexthop IP address
	BackupRouteNextHopAddress string `protobuf:"bytes,2,opt,name=backup_route_next_hop_address,json=backupRouteNextHopAddress,proto3" json:"backup_route_next_hop_address,omitempty"`
	// IP address of source of route
	BackupRouteSource string `protobuf:"bytes,3,opt,name=backup_route_source,json=backupRouteSource,proto3" json:"backup_route_source,omitempty"`
	// Metric
	BackupMetric uint32 `protobuf:"varint,4,opt,name=backup_metric,json=backupMetric,proto3" json:"backup_metric,omitempty"`
	// Primary Path
	PrimaryPath bool `protobuf:"varint,5,opt,name=primary_path,json=primaryPath,proto3" json:"primary_path,omitempty"`
	// Line Card Disjoint
	LineCardDisjoint bool `protobuf:"varint,6,opt,name=line_card_disjoint,json=lineCardDisjoint,proto3" json:"line_card_disjoint,omitempty"`
	// Downstream
	Downstream bool `protobuf:"varint,7,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// Node Protect
	NodeProtect bool `protobuf:"varint,8,opt,name=node_protect,json=nodeProtect,proto3" json:"node_protect,omitempty"`
	// SRLG Disjoint
	SrlgDisjoint         bool     `protobuf:"varint,9,opt,name=srlg_disjoint,json=srlgDisjoint,proto3" json:"srlg_disjoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3ShBackupPath) Reset()         { *m = Ospfv3ShBackupPath{} }
func (m *Ospfv3ShBackupPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShBackupPath) ProtoMessage()    {}
func (*Ospfv3ShBackupPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{2}
}

func (m *Ospfv3ShBackupPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShBackupPath.Unmarshal(m, b)
}
func (m *Ospfv3ShBackupPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShBackupPath.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShBackupPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShBackupPath.Merge(m, src)
}
func (m *Ospfv3ShBackupPath) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShBackupPath.Size(m)
}
func (m *Ospfv3ShBackupPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShBackupPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShBackupPath proto.InternalMessageInfo

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
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Route path next hop
	RoutePathNextHop string `protobuf:"bytes,2,opt,name=route_path_next_hop,json=routePathNextHop,proto3" json:"route_path_next_hop,omitempty"`
	// Path ID of path
	RoutePathId uint32 `protobuf:"varint,3,opt,name=route_path_id,json=routePathId,proto3" json:"route_path_id,omitempty"`
	// Backup Path Info
	RouteBackupPath      *Ospfv3ShBackupPath `protobuf:"bytes,4,opt,name=route_backup_path,json=routeBackupPath,proto3" json:"route_backup_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ospfv3EdmTopPath) Reset()         { *m = Ospfv3EdmTopPath{} }
func (m *Ospfv3EdmTopPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopPath) ProtoMessage()    {}
func (*Ospfv3EdmTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{3}
}

func (m *Ospfv3EdmTopPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopPath.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopPath.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopPath.Merge(m, src)
}
func (m *Ospfv3EdmTopPath) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopPath.Size(m)
}
func (m *Ospfv3EdmTopPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopPath proto.InternalMessageInfo

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

// OSPF External Route Extended Community Information
type Ospfv3ShRouteExtendedComm struct {
	// Domain ID value
	ExtendedCommunityDomainIdValue []byte `protobuf:"bytes,1,opt,name=extended_community_domain_id_value,json=extendedCommunityDomainIdValue,proto3" json:"extended_community_domain_id_value,omitempty"`
	// Domain ID type
	ExtendedCommunitylDomainIdType uint32 `protobuf:"varint,2,opt,name=extended_communityl_domain_id_type,json=extendedCommunitylDomainIdType,proto3" json:"extended_communityl_domain_id_type,omitempty"`
	// Area id
	ExtendedCommunityAreaId uint32 `protobuf:"varint,3,opt,name=extended_community_area_id,json=extendedCommunityAreaId,proto3" json:"extended_community_area_id,omitempty"`
	// Router id
	ExtendedCommunityRouterId string `protobuf:"bytes,4,opt,name=extended_community_router_id,json=extendedCommunityRouterId,proto3" json:"extended_community_router_id,omitempty"`
	// Route type
	ExtendedCommunityRouteType uint32 `protobuf:"varint,5,opt,name=extended_community_route_type,json=extendedCommunityRouteType,proto3" json:"extended_community_route_type,omitempty"`
	// Route Options
	ExtendedCommunityOptions uint32   `protobuf:"varint,6,opt,name=extended_community_options,json=extendedCommunityOptions,proto3" json:"extended_community_options,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Ospfv3ShRouteExtendedComm) Reset()         { *m = Ospfv3ShRouteExtendedComm{} }
func (m *Ospfv3ShRouteExtendedComm) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShRouteExtendedComm) ProtoMessage()    {}
func (*Ospfv3ShRouteExtendedComm) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{4}
}

func (m *Ospfv3ShRouteExtendedComm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Unmarshal(m, b)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShRouteExtendedComm.Merge(m, src)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Size(m)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShRouteExtendedComm.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShRouteExtendedComm proto.InternalMessageInfo

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityDomainIdValue() []byte {
	if m != nil {
		return m.ExtendedCommunityDomainIdValue
	}
	return nil
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunitylDomainIdType() uint32 {
	if m != nil {
		return m.ExtendedCommunitylDomainIdType
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityAreaId() uint32 {
	if m != nil {
		return m.ExtendedCommunityAreaId
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityRouterId() string {
	if m != nil {
		return m.ExtendedCommunityRouterId
	}
	return ""
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityRouteType() uint32 {
	if m != nil {
		return m.ExtendedCommunityRouteType
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityOptions() uint32 {
	if m != nil {
		return m.ExtendedCommunityOptions
	}
	return 0
}

// OSPFv3 Topology Information
type Ospfv3EdmTopology struct {
	// Route ID
	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	// Route distance
	RouteDistance uint32 `protobuf:"varint,2,opt,name=route_distance,json=routeDistance,proto3" json:"route_distance,omitempty"`
	// Route cost
	RouteCost uint32 `protobuf:"varint,3,opt,name=route_cost,json=routeCost,proto3" json:"route_cost,omitempty"`
	// Route type
	RouteType            uint32   `protobuf:"varint,4,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopology) Reset()         { *m = Ospfv3EdmTopology{} }
func (m *Ospfv3EdmTopology) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopology) ProtoMessage()    {}
func (*Ospfv3EdmTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{5}
}

func (m *Ospfv3EdmTopology) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopology.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopology) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopology.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopology) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopology.Merge(m, src)
}
func (m *Ospfv3EdmTopology) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopology.Size(m)
}
func (m *Ospfv3EdmTopology) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopology.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopology proto.InternalMessageInfo

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
	proto.RegisterType((*Ospfv3EdmTopologyExternal_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology_external_KEYS")
	proto.RegisterType((*Ospfv3EdmTopologyExternal)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology_external")
	proto.RegisterType((*Ospfv3ShBackupPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_sh_backup_path")
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3ShRouteExtendedComm)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_sh_route_extended_comm")
	proto.RegisterType((*Ospfv3EdmTopology)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology")
}

func init() { proto.RegisterFile("ospfv3_edm_topology_external.proto", fileDescriptor_b36ab0046955ae4e) }

var fileDescriptor_b36ab0046955ae4e = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4d, 0x4f, 0x1b, 0x47,
	0x18, 0xd6, 0x62, 0x03, 0x66, 0xec, 0xa5, 0x30, 0xa8, 0x74, 0xa1, 0x18, 0x99, 0xad, 0x2a, 0xf9,
	0xd0, 0xfa, 0x00, 0x55, 0x2f, 0x6d, 0xd5, 0x52, 0x40, 0xaa, 0x5b, 0x4a, 0xd1, 0x82, 0x2a, 0xe5,
	0x92, 0xd1, 0xb0, 0x33, 0xc6, 0x93, 0xec, 0xee, 0xac, 0x66, 0xc6, 0x8e, 0x7d, 0xc8, 0xaf, 0x88,
	0x72, 0x49, 0xa4, 0x48, 0xb9, 0xe4, 0x0f, 0x44, 0x4a, 0xfe, 0x5e, 0x34, 0x1f, 0x8b, 0x17, 0xbc,
	0xe1, 0x16, 0x2e, 0x68, 0xe7, 0x79, 0x9f, 0xf7, 0x9d, 0xe7, 0xfd, 0x1a, 0x0c, 0x42, 0x2e, 0xf3,
	0xc1, 0xf8, 0x00, 0x51, 0x92, 0x22, 0xc5, 0x73, 0x9e, 0xf0, 0xeb, 0x29, 0xa2, 0x13, 0x45, 0x45,
	0x86, 0x93, 0x5e, 0x2e, 0xb8, 0xe2, 0xf0, 0x71, 0xcc, 0x64, 0xcc, 0x11, 0xe3, 0x12, 0x4d, 0x04,
	0x62, 0xf9, 0xf8, 0x67, 0xe4, 0xbc, 0x78, 0x4e, 0x45, 0xcf, 0x7e, 0x6b, 0x6e, 0x4c, 0xa5, 0xa4,
	0xb2, 0xf8, 0xea, 0x8d, 0xc5, 0xc0, 0xfc, 0xe9, 0x15, 0xf1, 0x90, 0xe0, 0x23, 0x45, 0xe5, 0x9d,
	0x73, 0xf8, 0xc6, 0x03, 0x7b, 0xf7, 0xc9, 0x40, 0xff, 0x9c, 0x3c, 0xba, 0x80, 0x7b, 0xa0, 0xe5,
	0x02, 0xa3, 0x0c, 0xa7, 0x34, 0xf0, 0x3a, 0x5e, 0x77, 0x25, 0x6a, 0x3a, 0xec, 0x0c, 0xa7, 0x14,
	0x6e, 0x81, 0xc6, 0x58, 0x0c, 0xac, 0x79, 0xc1, 0x98, 0x97, 0xc7, 0x62, 0x60, 0x4c, 0x9b, 0x60,
	0x29, 0x17, 0x74, 0xc0, 0x26, 0x41, 0xcd, 0x18, 0xdc, 0x09, 0x7e, 0x07, 0x7c, 0xfb, 0x85, 0x12,
	0x9a, 0x5d, 0xab, 0x61, 0x50, 0xef, 0x78, 0x5d, 0x3f, 0x6a, 0x59, 0xf0, 0xd4, 0x60, 0xe1, 0x8b,
	0x3a, 0xd8, 0xb9, 0x4f, 0x20, 0x7c, 0xe5, 0x81, 0x55, 0x93, 0xcb, 0x8d, 0x2d, 0xd8, 0xef, 0x78,
	0xdd, 0xe6, 0xbe, 0xec, 0x7d, 0xd9, 0xda, 0xf5, 0x2a, 0x64, 0x45, 0xbe, 0x31, 0x5d, 0xba, 0x23,
	0x7c, 0xed, 0x81, 0xaf, 0xac, 0xb8, 0x1c, 0xab, 0x21, 0x4a, 0x98, 0x54, 0xc1, 0x41, 0xa7, 0xf6,
	0xf0, 0xea, 0xcc, 0xfd, 0x4e, 0xdd, 0x39, 0x56, 0xc3, 0x53, 0x26, 0x15, 0xfc, 0xe8, 0x81, 0xc0,
	0xaa, 0xd3, 0xbe, 0x19, 0xa1, 0x04, 0xc5, 0x3c, 0x4d, 0x47, 0x19, 0x53, 0xd3, 0xe0, 0x27, 0x53,
	0xc4, 0xe7, 0x0f, 0x24, 0x53, 0x0e, 0x51, 0x85, 0x92, 0x68, 0xd3, 0x80, 0x27, 0x0e, 0x3b, 0x2a,
	0xc4, 0x85, 0xef, 0x6a, 0xe0, 0xeb, 0x99, 0xe7, 0x15, 0x8e, 0x9f, 0x8e, 0x6c, 0x8a, 0xf0, 0x37,
	0xf0, 0xad, 0x3b, 0xda, 0x78, 0x2c, 0x53, 0x54, 0x0c, 0x70, 0x4c, 0xcb, 0x93, 0x1b, 0x58, 0x4a,
	0xa4, 0x19, 0xfd, 0x82, 0x60, 0x66, 0xf5, 0x0f, 0xd0, 0xbe, 0xe5, 0x9e, 0xd1, 0x89, 0x42, 0x43,
	0x9e, 0x23, 0x4c, 0x88, 0xa0, 0x52, 0xba, 0xd9, 0xde, 0x2a, 0x05, 0x38, 0xa3, 0x13, 0xf5, 0x17,
	0xcf, 0x0f, 0x2d, 0x01, 0xf6, 0xc0, 0xc6, 0xad, 0x08, 0x92, 0x8f, 0x44, 0x4c, 0xdd, 0xe8, 0xaf,
	0x97, 0xfc, 0x2e, 0x8c, 0x41, 0x6f, 0x81, 0xe3, 0xa7, 0x54, 0x09, 0x16, 0x17, 0x5b, 0x60, 0xc1,
	0x7f, 0x0d, 0x66, 0x17, 0x90, 0xa5, 0x58, 0x4c, 0x4d, 0x96, 0xc1, 0x62, 0xc7, 0xeb, 0x36, 0xf4,
	0x02, 0x1a, 0x4c, 0x37, 0x14, 0xfe, 0x00, 0x60, 0xc2, 0x32, 0x8a, 0x62, 0x2c, 0x08, 0x22, 0x4c,
	0x3e, 0xe1, 0x2c, 0x53, 0xc1, 0x92, 0x21, 0xae, 0x69, 0xcb, 0x11, 0x16, 0xe4, 0xd8, 0xe1, 0x70,
	0x17, 0x00, 0xc2, 0x9f, 0x65, 0x52, 0x09, 0x8a, 0xd3, 0x60, 0xd9, 0xb0, 0x4a, 0x88, 0xbe, 0x30,
	0xe3, 0x84, 0x22, 0xfd, 0x0a, 0xd1, 0x58, 0x05, 0x0d, 0x7b, 0xa1, 0xc6, 0xce, 0x2d, 0xa4, 0x85,
	0x4b, 0x91, 0x5c, 0xcf, 0xee, 0x5a, 0x31, 0x9c, 0x96, 0x06, 0x8b, 0x7b, 0xc2, 0xf7, 0x0b, 0x60,
	0xa3, 0x62, 0x12, 0xe1, 0xf7, 0x60, 0xb5, 0xb2, 0x33, 0x3e, 0xbb, 0xd5, 0x8e, 0x1f, 0xc1, 0x46,
	0x69, 0x7d, 0x8a, 0x66, 0xb8, 0x26, 0xac, 0xdd, 0x4c, 0xb3, 0x6b, 0x01, 0x0c, 0x81, 0x5f, 0xa2,
	0x33, 0x62, 0xaa, 0xee, 0x47, 0xcd, 0x1b, 0x62, 0x9f, 0xc0, 0xb7, 0x1e, 0x58, 0xb7, 0xa4, 0xd2,
	0xd8, 0x98, 0xa2, 0x37, 0xf7, 0x47, 0x0f, 0x37, 0xed, 0xa5, 0xcb, 0x23, 0xfb, 0x44, 0xfc, 0x69,
	0x10, 0xad, 0x32, 0xfc, 0x50, 0x03, 0xed, 0x7b, 0x17, 0x03, 0xfe, 0x0d, 0xc2, 0xf9, 0x9d, 0x45,
	0x84, 0xa7, 0x98, 0x65, 0x88, 0x11, 0x34, 0xc6, 0xc9, 0xc8, 0xd6, 0xb4, 0x15, 0xed, 0xd2, 0xbb,
	0xfb, 0x73, 0x6c, 0x78, 0x7d, 0xf2, 0xbf, 0x66, 0x55, 0xc7, 0x4a, 0x4a, 0xc1, 0xd4, 0x34, 0xb7,
	0x8f, 0xba, 0x5f, 0x11, 0x2b, 0x29, 0x82, 0x5d, 0x4e, 0x73, 0x0a, 0x7f, 0x01, 0xdb, 0x15, 0xba,
	0xb0, 0xa0, 0x78, 0xd6, 0x8e, 0x6f, 0xe6, 0x62, 0x1c, 0x0a, 0x8a, 0xfb, 0x04, 0xfe, 0x0e, 0x76,
	0x2a, 0x9c, 0x4d, 0xfa, 0x42, 0xbb, 0xd7, 0xed, 0xee, 0xcd, 0xb9, 0x9b, 0x75, 0x12, 0x7d, 0x02,
	0x0f, 0x41, 0xfb, 0x73, 0x01, 0x6c, 0x12, 0x8b, 0x46, 0xc0, 0x76, 0x75, 0x04, 0x93, 0xc0, 0xaf,
	0x95, 0x09, 0xf0, 0x5c, 0x31, 0x9e, 0x49, 0xb3, 0x4e, 0x7e, 0x14, 0xcc, 0xf9, 0xff, 0x67, 0xed,
	0xe1, 0x4b, 0xef, 0xee, 0xb8, 0xdb, 0xff, 0x03, 0x5b, 0xa0, 0xe1, 0x9e, 0x23, 0xe2, 0x06, 0x7d,
	0xd9, 0x9c, 0xfb, 0x44, 0x6f, 0x82, 0x35, 0x11, 0x26, 0x15, 0xce, 0xe2, 0xa2, 0xd2, 0x76, 0x92,
	0x8f, 0x1d, 0x08, 0xdb, 0x00, 0x58, 0x5a, 0xcc, 0xa5, 0x72, 0x85, 0x5c, 0x31, 0xc8, 0x11, 0x97,
	0x6a, 0x66, 0x36, 0x69, 0xd6, 0x4b, 0x66, 0x9d, 0xd5, 0xd5, 0x92, 0xf9, 0x35, 0x71, 0xf0, 0x29,
	0x00, 0x00, 0xff, 0xff, 0x58, 0x0b, 0x57, 0xcc, 0x73, 0x08, 0x00, 0x00,
}
