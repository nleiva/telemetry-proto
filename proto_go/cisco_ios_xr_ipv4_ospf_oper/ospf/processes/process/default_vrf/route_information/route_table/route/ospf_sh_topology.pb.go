// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_topology.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route

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
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTopology_KEYS) Reset()         { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()    {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_topology_e6b63dd28464f3da, []int{0}
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
	return fileDescriptor_ospf_sh_topology_e6b63dd28464f3da, []int{1}
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
	return fileDescriptor_ospf_sh_topology_e6b63dd28464f3da, []int{2}
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
	return fileDescriptor_ospf_sh_topology_e6b63dd28464f3da, []int{3}
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
	return fileDescriptor_ospf_sh_topology_e6b63dd28464f3da, []int{4}
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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_path")
}

func init() {
	proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_ospf_sh_topology_e6b63dd28464f3da)
}

var fileDescriptor_ospf_sh_topology_e6b63dd28464f3da = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5f, 0x6f, 0xe4, 0x34,
	0x10, 0x57, 0xae, 0x65, 0xbb, 0xf5, 0xb6, 0x70, 0x98, 0xa3, 0x97, 0x22, 0xe0, 0x96, 0x45, 0x82,
	0x15, 0x42, 0x2b, 0xd4, 0x96, 0xbf, 0x2f, 0xa8, 0x82, 0xab, 0xa8, 0xe8, 0x55, 0xa7, 0xb4, 0x87,
	0xc4, 0x93, 0xe5, 0x75, 0x26, 0x5d, 0x4b, 0x49, 0x6c, 0xd9, 0x4e, 0x6f, 0xfb, 0x09, 0xf8, 0x04,
	0x48, 0x3c, 0xf0, 0x8a, 0xf8, 0x9a, 0x28, 0x33, 0x4e, 0x93, 0xbb, 0x3e, 0xd3, 0x7b, 0x73, 0x7e,
	0x33, 0x9e, 0x19, 0xcf, 0xfc, 0xfc, 0x73, 0xd8, 0x9e, 0xf1, 0xb6, 0x10, 0x7e, 0x25, 0x82, 0xb1,
	0xa6, 0x34, 0x57, 0x37, 0x0b, 0xeb, 0x4c, 0x30, 0x7c, 0xa9, 0xb4, 0x57, 0x46, 0x68, 0xe3, 0xc5,
	0xda, 0x09, 0x6d, 0xaf, 0x8f, 0x04, 0x7a, 0x1a, 0x0b, 0x6e, 0xd1, 0xae, 0x5a, 0x3f, 0x05, 0xde,
	0x83, 0xef, 0x56, 0x8b, 0x1c, 0x0a, 0xd9, 0x94, 0x41, 0x5c, 0xbb, 0x62, 0xe1, 0x4c, 0x13, 0x40,
	0xe8, 0xba, 0x30, 0xae, 0x92, 0x41, 0x9b, 0x3a, 0x22, 0x41, 0x2e, 0x4b, 0xa0, 0xf5, 0xec, 0x25,
	0x7b, 0xff, 0xf5, 0xec, 0xe2, 0xd7, 0xa7, 0xbf, 0x5f, 0xf0, 0x4f, 0xd8, 0x4e, 0x8c, 0x29, 0x6a,
	0x59, 0x41, 0x9a, 0x4c, 0x93, 0xf9, 0x76, 0x36, 0x89, 0xd8, 0xb9, 0xac, 0x80, 0xef, 0xb1, 0x91,
	0x75, 0x50, 0xe8, 0x75, 0xfa, 0x00, 0x8d, 0xf1, 0x8b, 0x7f, 0xca, 0x76, 0x69, 0x25, 0x4a, 0xa8,
	0xaf, 0xc2, 0x2a, 0xdd, 0x98, 0x26, 0xf3, 0xdd, 0x6c, 0x87, 0xc0, 0x33, 0xc4, 0x66, 0x7f, 0x6c,
	0xb2, 0x87, 0xaf, 0x67, 0x6e, 0x93, 0x52, 0x89, 0x31, 0xee, 0x01, 0x25, 0x45, 0xec, 0x39, 0x05,
	0x5f, 0xb0, 0xf7, 0x86, 0x2e, 0x5d, 0x8a, 0x43, 0x4c, 0xf1, 0xee, 0xc0, 0x93, 0xf2, 0xf4, 0x21,
	0x2b, 0x08, 0x4e, 0xab, 0xf4, 0x08, 0x1d, 0x29, 0xe4, 0x33, 0x84, 0xf8, 0x47, 0x8c, 0xc5, 0xc6,
	0xdc, 0x58, 0x48, 0xbf, 0xc6, 0x9c, 0xdb, 0x88, 0x5c, 0xde, 0x58, 0xe0, 0x9f, 0xb3, 0x77, 0xc8,
	0xac, 0x4c, 0x5d, 0x83, 0x0a, 0x90, 0xa7, 0xdf, 0x4c, 0x93, 0xf9, 0x38, 0x7b, 0x1b, 0xe1, 0x9f,
	0x3a, 0x94, 0xff, 0x99, 0x74, 0x81, 0xda, 0x9e, 0xa7, 0xdf, 0x4e, 0x93, 0xf9, 0xe4, 0xe0, 0x7a,
	0xf1, 0xff, 0x4f, 0x71, 0x31, 0x68, 0xa4, 0x50, 0xa6, 0xaa, 0x4c, 0x1d, 0x0f, 0x70, 0x5a, 0x17,
	0x86, 0xff, 0x9d, 0x74, 0x27, 0xb0, 0x32, 0xac, 0x44, 0xa9, 0x7d, 0x48, 0xbf, 0x9b, 0x6e, 0xcc,
	0x27, 0x07, 0xe1, 0xbe, 0x8b, 0x6b, 0x0b, 0xc8, 0x76, 0x69, 0x4a, 0x32, 0xac, 0xce, 0xb4, 0x0f,
	0xb3, 0x13, 0xb6, 0x73, 0xeb, 0xa2, 0x89, 0x56, 0x1e, 0x94, 0xa9, 0x73, 0xe4, 0xdc, 0x6e, 0x16,
	0xbf, 0xf8, 0xc7, 0x8c, 0xd5, 0xb2, 0x36, 0xd1, 0xf6, 0x00, 0x6d, 0x03, 0x64, 0xf6, 0xcf, 0x88,
	0xf1, 0xbb, 0x8d, 0xe0, 0x33, 0x46, 0xf9, 0x84, 0x74, 0x20, 0x85, 0xee, 0xa2, 0x12, 0x03, 0x8e,
	0x1d, 0xc8, 0xd3, 0x9c, 0x7f, 0xd6, 0x35, 0xa8, 0xe7, 0x09, 0xc5, 0xa7, 0xad, 0x97, 0x1d, 0x53,
	0xbe, 0x60, 0xc4, 0x30, 0xe1, 0xf4, 0x52, 0x5c, 0x83, 0xf3, 0xda, 0xd4, 0x91, 0xdd, 0x14, 0x20,
	0xd3, 0xcb, 0xdf, 0x08, 0xee, 0x7d, 0xdb, 0x92, 0x3a, 0xdf, 0xcd, 0x69, 0x32, 0xdf, 0x8c, 0xbe,
	0x17, 0xb6, 0xe8, 0x7c, 0x8f, 0xd8, 0x1e, 0xf9, 0x16, 0xc6, 0xbd, 0x94, 0x2e, 0x17, 0xb9, 0xf6,
	0x41, 0xd6, 0x0a, 0xd2, 0xb7, 0x30, 0xf8, 0x23, 0xb4, 0x9e, 0x90, 0xf1, 0xe7, 0x68, 0xeb, 0xa9,
	0xed, 0x4d, 0xe3, 0x14, 0xa4, 0xa3, 0xc1, 0xc1, 0x2e, 0x10, 0x6a, 0x47, 0x1f, 0xab, 0x68, 0x6c,
	0x2e, 0xdb, 0x03, 0xea, 0x0a, 0xd2, 0x2d, 0x64, 0xa6, 0xbd, 0xd7, 0xe1, 0xeb, 0x0a, 0xe2, 0xb9,
	0x5f, 0x60, 0x25, 0x97, 0xed, 0xa8, 0xff, 0xba, 0x65, 0x66, 0x21, 0x75, 0x49, 0xc5, 0x8d, 0xdf,
	0x50, 0x71, 0x34, 0xea, 0x13, 0xa9, 0x4b, 0x2c, 0xed, 0x4b, 0xc6, 0xfb, 0xf1, 0x59, 0xa7, 0x8d,
	0xd3, 0xe1, 0x26, 0xdd, 0xc6, 0x16, 0x3f, 0xec, 0xe6, 0xf7, 0x3c, 0xe2, 0xbd, 0x2a, 0xc9, 0x26,
	0x18, 0x01, 0x6b, 0x55, 0x36, 0x39, 0xe4, 0x29, 0x43, 0x9d, 0xa0, 0x09, 0x1c, 0x37, 0xc1, 0x3c,
	0x8d, 0x06, 0xfe, 0x23, 0xfb, 0x30, 0x46, 0x77, 0xbd, 0x94, 0x39, 0xb8, 0xd2, 0x3e, 0x80, 0x83,
	0x3c, 0x9d, 0xe0, 0xc6, 0x7d, 0xca, 0xe3, 0x3a, 0x49, 0xcb, 0x6e, 0x1d, 0xf8, 0xf7, 0x6c, 0x7f,
	0x10, 0xa0, 0x5e, 0xba, 0xe1, 0xee, 0x1d, 0xac, 0x72, 0xef, 0x76, 0xf7, 0xf9, 0xd2, 0xf5, 0x5b,
	0x67, 0xff, 0x6e, 0xbc, 0xa2, 0xbc, 0x78, 0x27, 0xf9, 0x57, 0xec, 0x51, 0xd7, 0xb1, 0x00, 0xae,
	0x90, 0x0a, 0x86, 0xb2, 0xcf, 0xa3, 0x98, 0x44, 0x13, 0xaa, 0xff, 0x61, 0xc7, 0xd9, 0x1a, 0xd6,
	0x41, 0xac, 0x8c, 0x15, 0x32, 0xcf, 0x1d, 0x78, 0x1f, 0x5f, 0x03, 0x6a, 0xc8, 0x39, 0xac, 0xc3,
	0x2f, 0xc6, 0x1e, 0x93, 0xe9, 0x0e, 0x65, 0x37, 0x06, 0x02, 0x1f, 0x29, 0xfb, 0x84, 0xd1, 0xa7,
	0x28, 0xbd, 0xd4, 0x39, 0xde, 0x98, 0xed, 0x8c, 0x74, 0xf5, 0xac, 0x45, 0xf8, 0x0f, 0xec, 0x83,
	0x81, 0x9a, 0x69, 0x2f, 0x2a, 0x25, 0x7d, 0x68, 0x0b, 0x97, 0x2a, 0xe0, 0x85, 0x19, 0xc7, 0xb3,
	0xb7, 0x12, 0x73, 0xea, 0x9f, 0xb5, 0xe6, 0x53, 0xb4, 0xf2, 0x23, 0xf6, 0xf8, 0xd5, 0xbd, 0x8d,
	0xaa, 0xa8, 0x03, 0x78, 0x7b, 0xc6, 0xb1, 0x6a, 0xda, 0xf8, 0x42, 0x55, 0xb6, 0x5d, 0xdd, 0x79,
	0x43, 0xb6, 0xee, 0xbe, 0x21, 0xfb, 0x6c, 0x5c, 0x7a, 0x49, 0x2f, 0xc8, 0x18, 0xcd, 0x5b, 0xa5,
	0x97, 0xf8, 0x7e, 0x3c, 0x66, 0x5b, 0x9d, 0xf4, 0x10, 0x7d, 0x46, 0x92, 0x54, 0xe7, 0x09, 0x9b,
	0xa0, 0x81, 0x28, 0x1a, 0xc9, 0xc2, 0x5a, 0xe8, 0x04, 0x91, 0xe5, 0x08, 0xff, 0x03, 0x0e, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x81, 0x2c, 0x09, 0x4e, 0x21, 0x08, 0x00, 0x00,
}
