// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_topology.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_multicast_intact_route_table_route

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

// OSPF Route Information
type OspfShTopology_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTopology_KEYS) Reset()         { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()    {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed, []int{0}
}
func (m *OspfShTopology_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopology_KEYS.Unmarshal(m, b)
}
func (m *OspfShTopology_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopology_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShTopology_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopology_KEYS.Merge(dst, src)
}
func (m *OspfShTopology_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShTopology_KEYS.Size(m)
}
func (m *OspfShTopology_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopology_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopology_KEYS proto.InternalMessageInfo

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength,proto3" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo,proto3" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList        []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList,proto3" json:"route_path_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OspfShTopology) Reset()         { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()    {}
func (*OspfShTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed, []int{1}
}
func (m *OspfShTopology) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopology.Unmarshal(m, b)
}
func (m *OspfShTopology) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopology.Marshal(b, m, deterministic)
}
func (dst *OspfShTopology) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopology.Merge(dst, src)
}
func (m *OspfShTopology) XXX_Size() int {
	return xxx_messageInfo_OspfShTopology.Size(m)
}
func (m *OspfShTopology) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopology.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopology proto.InternalMessageInfo

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

type OspfShTime struct {
	Second               uint32   `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
	Nanosecond           uint32   `protobuf:"varint,2,opt,name=nanosecond,proto3" json:"nanosecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTime) Reset()         { *m = OspfShTime{} }
func (m *OspfShTime) String() string { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()    {}
func (*OspfShTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed, []int{2}
}
func (m *OspfShTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTime.Unmarshal(m, b)
}
func (m *OspfShTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTime.Marshal(b, m, deterministic)
}
func (dst *OspfShTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTime.Merge(dst, src)
}
func (m *OspfShTime) XXX_Size() int {
	return xxx_messageInfo_OspfShTime.Size(m)
}
func (m *OspfShTime) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTime.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTime proto.InternalMessageInfo

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

// OSPF Common Route Information
type OspfShTopCommon struct {
	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId,proto3" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric,proto3" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion,proto3" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion,proto3" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance,proto3" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime,proto3" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime,proto3" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority,proto3" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded,proto3" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered,proto3" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32   `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered,proto3" json:"route_srte_nbr_registered,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *OspfShTopCommon) Reset()         { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()    {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed, []int{3}
}
func (m *OspfShTopCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopCommon.Unmarshal(m, b)
}
func (m *OspfShTopCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopCommon.Marshal(b, m, deterministic)
}
func (dst *OspfShTopCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopCommon.Merge(dst, src)
}
func (m *OspfShTopCommon) XXX_Size() int {
	return xxx_messageInfo_OspfShTopCommon.Size(m)
}
func (m *OspfShTopCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopCommon.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopCommon proto.InternalMessageInfo

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName,proto3" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress,proto3" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid,proto3" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact,proto3" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath,proto3" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat           bool     `protobuf:"varint,10,opt,name=area_format,json=areaFormat,proto3" json:"area_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTopPath) Reset()         { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()    {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed, []int{4}
}
func (m *OspfShTopPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopPath.Unmarshal(m, b)
}
func (m *OspfShTopPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopPath.Marshal(b, m, deterministic)
}
func (dst *OspfShTopPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopPath.Merge(dst, src)
}
func (m *OspfShTopPath) XXX_Size() int {
	return xxx_messageInfo_OspfShTopPath.Size(m)
}
func (m *OspfShTopPath) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopPath.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopPath proto.InternalMessageInfo

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_route_table.route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_route_table.route.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_route_table.route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_route_table.route.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_route_table.route.ospf_sh_top_path")
}

func init() {
	proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed)
}

var fileDescriptor_ospf_sh_topology_f0a24612a9cd39ed = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5f, 0x6f, 0x23, 0x35,
	0x10, 0xd7, 0x5e, 0x4b, 0x92, 0x3a, 0x0d, 0x1c, 0xe6, 0xe8, 0x6d, 0x11, 0x70, 0x21, 0x48, 0x10,
	0x21, 0x14, 0xa1, 0xb6, 0xfc, 0x7d, 0x41, 0x15, 0x5c, 0x45, 0x45, 0xaf, 0x3a, 0x6d, 0x7b, 0x48,
	0x3c, 0x59, 0xce, 0xee, 0x6c, 0x63, 0x69, 0x77, 0x6d, 0xd9, 0x4e, 0x48, 0x5f, 0x78, 0x40, 0x7c,
	0x02, 0x3e, 0x03, 0x7f, 0x24, 0xf8, 0x14, 0x7c, 0x33, 0xb4, 0x33, 0x76, 0xb3, 0x77, 0x7d, 0xa7,
	0xf7, 0x12, 0x79, 0xe7, 0x37, 0x1e, 0xff, 0x3c, 0xf3, 0xf3, 0x4c, 0xd8, 0x9e, 0x76, 0xa6, 0x14,
	0x6e, 0x21, 0xbc, 0x36, 0xba, 0xd2, 0x57, 0xd7, 0x33, 0x63, 0xb5, 0xd7, 0xdc, 0xe4, 0xca, 0xe5,
	0x5a, 0x28, 0xed, 0xc4, 0xda, 0x0a, 0x65, 0x56, 0x47, 0x02, 0x3d, 0xb5, 0x01, 0x3b, 0x6b, 0x57,
	0xad, 0x5f, 0x0e, 0xce, 0x81, 0x8b, 0xab, 0xd9, 0xca, 0x96, 0xf8, 0x33, 0xb3, 0x7a, 0xe9, 0x41,
	0xa8, 0xa6, 0xd4, 0xb6, 0x96, 0x5e, 0xe9, 0x66, 0x56, 0x2f, 0x2b, 0xaf, 0x72, 0xe9, 0xbc, 0x50,
	0x8d, 0x97, 0xb9, 0x17, 0xe4, 0xe2, 0xe5, 0xbc, 0x02, 0x72, 0x9f, 0xfc, 0x96, 0xb0, 0x37, 0x5f,
	0x24, 0x23, 0xbe, 0x7f, 0xfc, 0xe3, 0x05, 0x7f, 0x8f, 0xed, 0x86, 0x23, 0x44, 0x23, 0x6b, 0x48,
	0x93, 0x71, 0x32, 0xdd, 0xc9, 0x86, 0xc1, 0x76, 0x2e, 0x6b, 0xe0, 0xfb, 0x6c, 0xb0, 0xb2, 0x25,
	0xc1, 0xf7, 0x10, 0xee, 0xaf, 0x6c, 0x89, 0xd0, 0x1e, 0xeb, 0x19, 0x0b, 0xa5, 0x5a, 0xa7, 0x5b,
	0x08, 0x84, 0x2f, 0xfe, 0x3e, 0x1b, 0xd1, 0x4a, 0x54, 0xd0, 0x5c, 0xf9, 0x45, 0xba, 0x3d, 0x4e,
	0xa6, 0xa3, 0x6c, 0x97, 0x8c, 0x67, 0x68, 0x9b, 0xfc, 0xb1, 0xcd, 0xee, 0xbf, 0x48, 0xaa, 0xe5,
	0x43, 0xf4, 0x43, 0xdc, 0x03, 0xe2, 0x83, 0xb6, 0xa7, 0x14, 0x7c, 0xc6, 0xde, 0xe8, 0xba, 0xc4,
	0x23, 0x0e, 0xf1, 0x88, 0xd7, 0x3b, 0x9e, 0x74, 0xce, 0x26, 0x64, 0x0d, 0xde, 0xaa, 0x3c, 0x3d,
	0x42, 0x47, 0x0a, 0xf9, 0x04, 0x4d, 0xfc, 0x1d, 0xc6, 0x42, 0xd2, 0xae, 0x0d, 0xa4, 0x9f, 0xe2,
	0x99, 0x3b, 0x68, 0xb9, 0xbc, 0x36, 0xc0, 0x3f, 0x64, 0xaf, 0x11, 0x9c, 0xeb, 0xa6, 0x81, 0xdc,
	0x43, 0x91, 0x7e, 0x36, 0x4e, 0xa6, 0x83, 0xec, 0x55, 0x34, 0x7f, 0x13, 0xad, 0xfc, 0xf7, 0x24,
	0x06, 0x6a, 0x0b, 0x94, 0x7e, 0x3e, 0x4e, 0xa6, 0xc3, 0x83, 0x5f, 0x93, 0xd9, 0xff, 0x5d, 0xf0,
	0x59, 0x27, 0xaf, 0x22, 0xd7, 0x75, 0xad, 0x9b, 0x70, 0x9f, 0xd3, 0xa6, 0xd4, 0xfc, 0x9f, 0x24,
	0x5e, 0xc8, 0x48, 0xbf, 0x10, 0x95, 0x72, 0x3e, 0xfd, 0x62, 0xbc, 0x35, 0x1d, 0x1e, 0xfc, 0x72,
	0xc7, 0x5c, 0x5b, 0x3e, 0xd9, 0x88, 0x6a, 0x28, 0xfd, 0xe2, 0x4c, 0x39, 0x3f, 0x39, 0x61, 0xbb,
	0x37, 0x2e, 0x8a, 0x44, 0xe7, 0x20, 0xd7, 0x4d, 0x81, 0x62, 0x1d, 0x65, 0xe1, 0x8b, 0xbf, 0xcb,
	0x58, 0x23, 0x1b, 0x1d, 0xb0, 0x7b, 0x88, 0x75, 0x2c, 0x93, 0x7f, 0x7b, 0x8c, 0xdf, 0xce, 0x0b,
	0x9f, 0x30, 0x3a, 0x4f, 0x48, 0x0b, 0x52, 0xa8, 0x18, 0x95, 0xf4, 0x71, 0x6c, 0x41, 0x9e, 0x16,
	0xfc, 0x83, 0x98, 0xaf, 0x8d, 0x8a, 0x28, 0x3e, 0x6d, 0xbd, 0x8c, 0x3a, 0xfa, 0x88, 0x91, 0xfe,
	0x84, 0x55, 0x73, 0xb1, 0x02, 0xeb, 0x94, 0x6e, 0xf0, 0x69, 0x8c, 0x32, 0x0a, 0x90, 0xa9, 0xf9,
	0x0f, 0x64, 0xde, 0xf8, 0xb6, 0x94, 0xa2, 0x6f, 0xfb, 0x4e, 0xb6, 0x83, 0xef, 0x85, 0x29, 0xa3,
	0xef, 0x11, 0xdb, 0x23, 0xdf, 0x52, 0xdb, 0x9f, 0xa4, 0x2d, 0x44, 0xa1, 0x9c, 0x97, 0x4d, 0x0e,
	0xe9, 0x2b, 0x18, 0xfc, 0x01, 0xa2, 0x27, 0x04, 0x7e, 0x1b, 0xb0, 0x8d, 0xf0, 0x9d, 0x5e, 0xda,
	0x1c, 0xd2, 0x5e, 0xe7, 0x62, 0x17, 0x68, 0xe2, 0x7f, 0x27, 0x91, 0xc5, 0xd2, 0x14, 0xb2, 0xbd,
	0xa0, 0xaa, 0x21, 0xed, 0xa3, 0x6e, 0x7f, 0xbe, 0x43, 0x29, 0xa8, 0x1a, 0x42, 0x16, 0x9e, 0x21,
	0xaf, 0xcb, 0xb6, 0xf0, 0x7f, 0xde, 0xc8, 0xb6, 0x94, 0xaa, 0x22, 0xaa, 0x83, 0x97, 0x82, 0x2a,
	0xc9, 0xe0, 0x44, 0xaa, 0x0a, 0x89, 0x7e, 0xcc, 0xf8, 0xa6, 0xb4, 0xc6, 0x2a, 0x6d, 0x95, 0xbf,
	0x4e, 0x77, 0x30, 0xfd, 0xf7, 0x63, 0x6d, 0x9f, 0x06, 0xfb, 0xa6, 0x9f, 0xc9, 0xa5, 0xd7, 0x02,
	0xd6, 0x79, 0xb5, 0x2c, 0xa0, 0x48, 0x19, 0x76, 0x18, 0xaa, 0xce, 0xf1, 0xd2, 0xeb, 0xc7, 0x01,
	0xe0, 0x5f, 0xb3, 0xb7, 0x43, 0x74, 0xbb, 0x69, 0x82, 0x16, 0xae, 0x94, 0xf3, 0x60, 0xa1, 0x48,
	0x87, 0xb8, 0x71, 0x9f, 0xce, 0xb1, 0xb1, 0x19, 0x66, 0x37, 0x0e, 0xfc, 0x4b, 0xb6, 0xdf, 0x09,
	0xd0, 0xcc, 0x6d, 0x77, 0xf7, 0x2e, 0xb2, 0xdc, 0xbb, 0xd9, 0x7d, 0x3e, 0xb7, 0x9b, 0xad, 0x93,
	0xbf, 0xb6, 0x9e, 0xeb, 0xd9, 0xf8, 0x5e, 0xf9, 0x27, 0xec, 0x41, 0x4c, 0xa1, 0x07, 0x5b, 0xca,
	0x1c, 0xba, 0xb3, 0x84, 0x87, 0xbe, 0x13, 0x20, 0x9c, 0x1b, 0x87, 0x51, 0xcf, 0x0d, 0xac, 0xbd,
	0x58, 0x68, 0x23, 0x64, 0x51, 0x58, 0x70, 0x2e, 0x0c, 0x18, 0x4a, 0xc8, 0x39, 0xac, 0xfd, 0x77,
	0xda, 0x1c, 0x13, 0x74, 0x4b, 0xce, 0x5b, 0x9d, 0xd1, 0x10, 0xe4, 0xfc, 0x88, 0xd1, 0xa7, 0xa8,
	0x9c, 0x54, 0x05, 0xbe, 0xa6, 0x9d, 0x8c, 0x3a, 0xf2, 0x59, 0x6b, 0xe1, 0x5f, 0xb1, 0xb7, 0x3a,
	0x8d, 0x4f, 0x39, 0x51, 0x77, 0xaa, 0x8c, 0x8f, 0x69, 0x10, 0xee, 0xde, 0xb6, 0x9f, 0x53, 0xf7,
	0xa4, 0x85, 0x4f, 0x11, 0xe5, 0x47, 0xec, 0xe1, 0xf3, 0x7b, 0x97, 0x79, 0x4d, 0x19, 0xc0, 0x97,
	0x35, 0x08, 0xac, 0x69, 0xe3, 0xb3, 0xbc, 0x36, 0xed, 0xea, 0xd6, 0xf4, 0xe9, 0xdf, 0x9e, 0x3e,
	0xfb, 0x6c, 0x50, 0x39, 0x49, 0xb3, 0x67, 0x80, 0x70, 0xbf, 0x72, 0x12, 0x27, 0xcf, 0x43, 0xd6,
	0x8f, 0x6d, 0x89, 0xe4, 0xd3, 0x93, 0xd4, 0x91, 0x1e, 0xb1, 0x21, 0x02, 0xa4, 0xd9, 0x20, 0x16,
	0xd6, 0x9a, 0x4e, 0xd0, 0x32, 0xef, 0xe1, 0x7f, 0x8d, 0xc3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x42, 0xcd, 0x01, 0x22, 0x85, 0x08, 0x00, 0x00,
}
