// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_topology_internal.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_internal_routes_internal_route

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

// OSPFv3 Topology Internal Information
type Ospfv3EdmTopologyInternal_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopologyInternal_KEYS) Reset()         { *m = Ospfv3EdmTopologyInternal_KEYS{} }
func (m *Ospfv3EdmTopologyInternal_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyInternal_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmTopologyInternal_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8df910defcd84a2, []int{0}
}

func (m *Ospfv3EdmTopologyInternal_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyInternal_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyInternal_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmTopologyInternal_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS.Size(m)
}
func (m *Ospfv3EdmTopologyInternal_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyInternal_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyInternal_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmTopologyInternal_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ospfv3EdmTopologyInternal_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ospfv3EdmTopologyInternal struct {
	// Common Route topology information
	RouteTopology *Ospfv3EdmTopology `protobuf:"bytes,50,opt,name=route_topology,json=routeTopology,proto3" json:"route_topology,omitempty"`
	// Route area ID
	RouteAreaId uint32 `protobuf:"varint,51,opt,name=route_area_id,json=routeAreaId,proto3" json:"route_area_id,omitempty"`
	// List of topology source information
	RouteSource []*Ospfv3EdmTopSource `protobuf:"bytes,52,rep,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	// List of paths to this route
	RoutePathList        []*Ospfv3EdmTopPath `protobuf:"bytes,53,rep,name=route_path_list,json=routePathList,proto3" json:"route_path_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ospfv3EdmTopologyInternal) Reset()         { *m = Ospfv3EdmTopologyInternal{} }
func (m *Ospfv3EdmTopologyInternal) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyInternal) ProtoMessage()    {}
func (*Ospfv3EdmTopologyInternal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8df910defcd84a2, []int{1}
}

func (m *Ospfv3EdmTopologyInternal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyInternal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyInternal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyInternal.Merge(m, src)
}
func (m *Ospfv3EdmTopologyInternal) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyInternal.Size(m)
}
func (m *Ospfv3EdmTopologyInternal) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyInternal.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyInternal proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyInternal) GetRouteTopology() *Ospfv3EdmTopology {
	if m != nil {
		return m.RouteTopology
	}
	return nil
}

func (m *Ospfv3EdmTopologyInternal) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *Ospfv3EdmTopologyInternal) GetRouteSource() []*Ospfv3EdmTopSource {
	if m != nil {
		return m.RouteSource
	}
	return nil
}

func (m *Ospfv3EdmTopologyInternal) GetRoutePathList() []*Ospfv3EdmTopPath {
	if m != nil {
		return m.RoutePathList
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
	return fileDescriptor_a8df910defcd84a2, []int{2}
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
	return fileDescriptor_a8df910defcd84a2, []int{3}
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

// OSPFv3 topology source information
type Ospfv3EdmTopSource struct {
	// Route source of the advertising router
	RouteSourceAdverstingRouter string `protobuf:"bytes,1,opt,name=route_source_adversting_router,json=routeSourceAdverstingRouter,proto3" json:"route_source_adversting_router,omitempty"`
	// Route source ID
	RouteSourceId string `protobuf:"bytes,2,opt,name=route_source_id,json=routeSourceId,proto3" json:"route_source_id,omitempty"`
	// Type of LSA advertising the prefix, see RFC5340
	RouteSourceLsaType   uint32   `protobuf:"varint,3,opt,name=route_source_lsa_type,json=routeSourceLsaType,proto3" json:"route_source_lsa_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopSource) Reset()         { *m = Ospfv3EdmTopSource{} }
func (m *Ospfv3EdmTopSource) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopSource) ProtoMessage()    {}
func (*Ospfv3EdmTopSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8df910defcd84a2, []int{4}
}

func (m *Ospfv3EdmTopSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopSource.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopSource.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopSource.Merge(m, src)
}
func (m *Ospfv3EdmTopSource) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopSource.Size(m)
}
func (m *Ospfv3EdmTopSource) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopSource.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopSource proto.InternalMessageInfo

func (m *Ospfv3EdmTopSource) GetRouteSourceAdverstingRouter() string {
	if m != nil {
		return m.RouteSourceAdverstingRouter
	}
	return ""
}

func (m *Ospfv3EdmTopSource) GetRouteSourceId() string {
	if m != nil {
		return m.RouteSourceId
	}
	return ""
}

func (m *Ospfv3EdmTopSource) GetRouteSourceLsaType() uint32 {
	if m != nil {
		return m.RouteSourceLsaType
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
	return fileDescriptor_a8df910defcd84a2, []int{5}
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
	proto.RegisterType((*Ospfv3EdmTopologyInternal_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_edm_topology_internal_KEYS")
	proto.RegisterType((*Ospfv3EdmTopologyInternal)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_edm_topology_internal")
	proto.RegisterType((*Ospfv3ShBackupPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_sh_backup_path")
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3EdmTopSource)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_edm_top_source")
	proto.RegisterType((*Ospfv3EdmTopology)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.internal_routes.internal_route.ospfv3_edm_topology")
}

func init() { proto.RegisterFile("ospfv3_edm_topology_internal.proto", fileDescriptor_a8df910defcd84a2) }

var fileDescriptor_a8df910defcd84a2 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xce, 0x52, 0x5e, 0x3e, 0xa6, 0x2d, 0x2f, 0x0c, 0xc1, 0x2c, 0x41, 0x48, 0x59, 0xa2, 0xe9,
	0x85, 0x6e, 0x22, 0xa8, 0x77, 0x26, 0x22, 0x98, 0xd8, 0x88, 0x84, 0x14, 0x6e, 0xbc, 0x9a, 0x0c,
	0x3b, 0xd3, 0x76, 0x74, 0xbb, 0xb3, 0x99, 0x99, 0xd6, 0xf6, 0xda, 0xdf, 0xe0, 0x95, 0xd1, 0xc4,
	0xc4, 0xf8, 0x13, 0xf4, 0xef, 0x99, 0x3d, 0x73, 0x5a, 0xb6, 0x84, 0x70, 0x27, 0x77, 0x7b, 0x9e,
	0xf3, 0xcc, 0x39, 0xcf, 0x9c, 0x8f, 0x59, 0x12, 0x69, 0x9b, 0x77, 0x86, 0x07, 0x4c, 0x8a, 0x3e,
	0x73, 0x3a, 0xd7, 0xa9, 0xee, 0x8e, 0x99, 0xca, 0x9c, 0x34, 0x19, 0x4f, 0xe3, 0xdc, 0x68, 0xa7,
	0x29, 0x4f, 0x94, 0x4d, 0x34, 0x53, 0xda, 0xb2, 0x91, 0x61, 0x2a, 0x1f, 0x3e, 0x67, 0x78, 0x4a,
	0xe7, 0xd2, 0xc4, 0xfe, 0xbb, 0xe0, 0x26, 0xd2, 0x5a, 0x69, 0x27, 0x5f, 0xb1, 0x90, 0x1d, 0x3e,
	0x48, 0x1d, 0x1b, 0x9a, 0x4e, 0x3c, 0x09, 0xc9, 0x8c, 0x1e, 0x38, 0x69, 0xaf, 0xd9, 0xd1, 0xe7,
	0x80, 0xec, 0xde, 0xa6, 0x84, 0xbd, 0x7d, 0xfd, 0xfe, 0x9c, 0xee, 0x92, 0x1a, 0xc6, 0x66, 0x19,
	0xef, 0xcb, 0x30, 0x68, 0x04, 0xcd, 0xe5, 0x76, 0x15, 0xb1, 0x53, 0xde, 0x97, 0xf4, 0x1e, 0x59,
	0xc8, 0x8d, 0xec, 0xa8, 0x51, 0x38, 0x07, 0x4e, 0xb4, 0xe8, 0x1e, 0xa9, 0xfb, 0x2f, 0x96, 0xca,
	0xac, 0xeb, 0x7a, 0x61, 0xa5, 0x11, 0x34, 0xeb, 0xed, 0x9a, 0x07, 0x4f, 0x00, 0x8b, 0x7e, 0xcc,
	0x93, 0xfb, 0xb7, 0xa9, 0xa0, 0xdf, 0x02, 0xb2, 0x02, 0x82, 0xa7, 0xbe, 0x70, 0xbf, 0x11, 0x34,
	0xab, 0xfb, 0xc3, 0xf8, 0x9f, 0xd7, 0x28, 0xbe, 0x41, 0x59, 0xbb, 0x0e, 0xae, 0x0b, 0x34, 0x69,
	0x44, 0x3c, 0xc0, 0xb8, 0x91, 0x9c, 0x29, 0x11, 0x1e, 0xc0, 0x2d, 0xab, 0x00, 0x1e, 0x1a, 0xc9,
	0x5b, 0x82, 0x7e, 0x0d, 0x48, 0xcd, 0x93, 0xac, 0x1e, 0x98, 0x44, 0x86, 0x4f, 0x1b, 0x95, 0x66,
	0x75, 0x7f, 0x74, 0xe7, 0x37, 0xc0, 0xfc, 0xa8, 0xee, 0x1c, 0x0c, 0xfa, 0x3d, 0x20, 0xff, 0x7b,
	0x75, 0x39, 0x77, 0x3d, 0x96, 0x2a, 0xeb, 0xc2, 0x67, 0x20, 0xf0, 0xee, 0x4b, 0x0c, 0x12, 0xb0,
	0xc4, 0x67, 0xdc, 0xf5, 0x4e, 0x94, 0x75, 0xd1, 0xaf, 0x0a, 0xd9, 0x40, 0x9a, 0xed, 0xb1, 0x4b,
	0x9e, 0x7c, 0x1c, 0x78, 0x22, 0x7d, 0x41, 0xb6, 0xd0, 0xf4, 0x17, 0x80, 0xd8, 0x1d, 0x9e, 0xc8,
	0xf2, 0xb0, 0x86, 0x9e, 0xd2, 0x2e, 0x18, 0xad, 0x09, 0x01, 0x26, 0xf7, 0x25, 0xd9, 0x9e, 0x39,
	0x9e, 0xc9, 0x91, 0x63, 0x3d, 0x9d, 0x33, 0x2e, 0x84, 0x91, 0xd6, 0xe2, 0x40, 0x6f, 0x96, 0x02,
	0x9c, 0xca, 0x91, 0x7b, 0xa3, 0xf3, 0x43, 0x4f, 0xa0, 0x31, 0x59, 0x9f, 0x89, 0x80, 0xfd, 0xad,
	0xc0, 0xb9, 0xb5, 0xd2, 0x39, 0xac, 0xf5, 0x1e, 0xa9, 0x23, 0xbf, 0x2f, 0x9d, 0x51, 0x49, 0x38,
	0xef, 0x77, 0xc2, 0x83, 0xef, 0x00, 0xf3, 0x3b, 0xa7, 0xfa, 0xdc, 0x8c, 0xe1, 0x96, 0xe1, 0x7f,
	0x8d, 0xa0, 0xb9, 0x54, 0xec, 0x1c, 0x60, 0x45, 0x59, 0xe8, 0x23, 0x42, 0x53, 0x95, 0x49, 0x96,
	0x70, 0x23, 0x98, 0x50, 0xf6, 0x83, 0x56, 0x99, 0x0b, 0x17, 0x80, 0xb8, 0x5a, 0x78, 0x8e, 0xb8,
	0x11, 0xc7, 0x88, 0xd3, 0x1d, 0x42, 0x84, 0xfe, 0x94, 0x59, 0x67, 0x24, 0xef, 0x87, 0x8b, 0xc0,
	0x2a, 0x21, 0x45, 0xc2, 0x4c, 0x0b, 0xc9, 0x8a, 0xb7, 0x47, 0x26, 0x2e, 0x5c, 0xf2, 0x09, 0x0b,
	0xec, 0xcc, 0x43, 0x85, 0x70, 0x6b, 0xd2, 0xee, 0x55, 0xae, 0x65, 0xe0, 0xd4, 0x0a, 0x70, 0x92,
	0x27, 0xfa, 0x3d, 0x47, 0xd6, 0x6f, 0xe8, 0x27, 0x7d, 0x40, 0x56, 0x6e, 0xec, 0x4c, 0x5d, 0xcd,
	0xb4, 0xe3, 0x31, 0x59, 0x2f, 0xcd, 0xe1, 0xa4, 0x19, 0xd8, 0x84, 0xd5, 0xe9, 0x4c, 0x60, 0x0b,
	0xae, 0x36, 0x0f, 0xe8, 0x4a, 0xe0, 0xfb, 0x52, 0x9d, 0x12, 0x5b, 0x82, 0xfe, 0x0c, 0xc8, 0x9a,
	0x27, 0x95, 0xc6, 0x06, 0x8a, 0x7e, 0xa7, 0xeb, 0x37, 0x3b, 0xb6, 0x6d, 0xbf, 0x6e, 0xaf, 0x00,
	0x29, 0x84, 0x46, 0x7f, 0x82, 0xe9, 0x84, 0xcf, 0x6e, 0x2a, 0x3d, 0x22, 0x3b, 0xe5, 0xc9, 0x62,
	0x5c, 0x0c, 0xa5, 0xb1, 0x4e, 0x65, 0x5d, 0x9f, 0xc1, 0x60, 0x29, 0xb7, 0x4a, 0x1b, 0x7d, 0x38,
	0xe5, 0xc0, 0xe8, 0x19, 0xfa, 0x70, 0xb2, 0xe0, 0x18, 0x44, 0x09, 0x2c, 0x6a, 0xbd, 0x74, 0xaa,
	0x25, 0xe8, 0x13, 0xb2, 0x31, 0xc3, 0x4b, 0x2d, 0x67, 0x6e, 0x9c, 0x4b, 0xac, 0x2c, 0x2d, 0xb1,
	0x4f, 0x2c, 0xbf, 0x18, 0xe7, 0x32, 0xfa, 0x12, 0x5c, 0x6f, 0xb9, 0x7f, 0x16, 0x37, 0xc9, 0x12,
	0xae, 0xa4, 0x40, 0x85, 0x8b, 0x60, 0xb7, 0x44, 0x31, 0x0d, 0xde, 0x25, 0x94, 0x75, 0x3c, 0x4b,
	0x24, 0x88, 0xa9, 0xa3, 0x98, 0x63, 0x04, 0xe9, 0x36, 0x21, 0x9e, 0x96, 0x68, 0xeb, 0x50, 0xc1,
	0x32, 0x20, 0x47, 0xda, 0xba, 0x2b, 0x37, 0x08, 0x9c, 0x2f, 0xb9, 0x0b, 0x5d, 0x97, 0x0b, 0xf0,
	0x1f, 0x3d, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x52, 0xba, 0xf1, 0xf0, 0x6d, 0x07, 0x00, 0x00,
}
