// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_rpl_as_protocol_routes_protocol_route

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

// Information of a rib route head and rib proto route
type Ipv6RibEdmRoute_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	As                   string   `protobuf:"bytes,5,opt,name=as,proto3" json:"as,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,7,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmRoute_KEYS) Reset()         { *m = Ipv6RibEdmRoute_KEYS{} }
func (m *Ipv6RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{0}
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Unmarshal(m, b)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv6RibEdmRoute_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Merge(dst, src)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmRoute_KEYS.Size(m)
}
func (m *Ipv6RibEdmRoute_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmRoute_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmRoute_KEYS proto.InternalMessageInfo

func (m *Ipv6RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv6RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion,proto3" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance,proto3" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority,proto3" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType,proto3" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags,proto3" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags,proto3" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag,proto3" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance,proto3" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance,proto3" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric,proto3" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount,proto3" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity,proto3" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex,proto3" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence,proto3" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup,proto3" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag,proto3" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass,proto3" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount,proto3" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active,proto3" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion,proto3" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName,proto3" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge,proto3" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version,proto3" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion,proto3" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime,proto3" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath            *Ipv6RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ipv6RibEdmRoute) Reset()         { *m = Ipv6RibEdmRoute{} }
func (m *Ipv6RibEdmRoute) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute) ProtoMessage()    {}
func (*Ipv6RibEdmRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{1}
}
func (m *Ipv6RibEdmRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmRoute.Unmarshal(m, b)
}
func (m *Ipv6RibEdmRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmRoute.Marshal(b, m, deterministic)
}
func (dst *Ipv6RibEdmRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmRoute.Merge(dst, src)
}
func (m *Ipv6RibEdmRoute) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmRoute.Size(m)
}
func (m *Ipv6RibEdmRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmRoute.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmRoute proto.InternalMessageInfo

func (m *Ipv6RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePath() *Ipv6RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

type Ipv6RibEdmAddr struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmAddr) Reset()         { *m = Ipv6RibEdmAddr{} }
func (m *Ipv6RibEdmAddr) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAddr) ProtoMessage()    {}
func (*Ipv6RibEdmAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{2}
}
func (m *Ipv6RibEdmAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmAddr.Unmarshal(m, b)
}
func (m *Ipv6RibEdmAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmAddr.Marshal(b, m, deterministic)
}
func (dst *Ipv6RibEdmAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmAddr.Merge(dst, src)
}
func (m *Ipv6RibEdmAddr) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmAddr.Size(m)
}
func (m *Ipv6RibEdmAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmAddr.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmAddr proto.InternalMessageInfo

func (m *Ipv6RibEdmAddr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Information of a rib path
type Ipv6RibEdmPath struct {
	// Next path
	Ipv6RibEdmPath       []*Ipv6RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv6_rib_edm_path,json=ipv6RibEdmPath,proto3" json:"ipv6_rib_edm_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ipv6RibEdmPath) Reset()         { *m = Ipv6RibEdmPath{} }
func (m *Ipv6RibEdmPath) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPath) ProtoMessage()    {}
func (*Ipv6RibEdmPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{3}
}
func (m *Ipv6RibEdmPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmPath.Unmarshal(m, b)
}
func (m *Ipv6RibEdmPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmPath.Marshal(b, m, deterministic)
}
func (dst *Ipv6RibEdmPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmPath.Merge(dst, src)
}
func (m *Ipv6RibEdmPath) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmPath.Size(m)
}
func (m *Ipv6RibEdmPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmPath proto.InternalMessageInfo

func (m *Ipv6RibEdmPath) GetIpv6RibEdmPath() []*Ipv6RibEdmPathItem {
	if m != nil {
		return m.Ipv6RibEdmPath
	}
	return nil
}

type Ipv6RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource,proto3" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop,proto3" json:"v6_nexthop,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,5,opt,name=metric,proto3" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,6,opt,name=load_metric,json=loadMetric,proto3" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,7,opt,name=flags64,proto3" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,9,opt,name=private_flags,json=privateFlags,proto3" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,10,opt,name=looped,proto3" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,11,opt,name=next_hop_table_id,json=nextHopTableId,proto3" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,12,opt,name=next_hop_vrf_name,json=nextHopVrfName,proto3" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,13,opt,name=next_hop_table_name,json=nextHopTableName,proto3" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,14,opt,name=next_hop_afi,json=nextHopAfi,proto3" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,15,opt,name=next_hop_safi,json=nextHopSafi,proto3" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,16,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,17,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,18,opt,name=pathid,proto3" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,19,opt,name=backup_pathid,json=backupPathid,proto3" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,20,opt,name=ref_cnt_of_backup,json=refCntOfBackup,proto3" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,21,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities,proto3" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,22,opt,name=mvpn_present,json=mvpnPresent,proto3" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,23,opt,name=pathrt_present,json=pathrtPresent,proto3" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,24,opt,name=vrfimportrt_present,json=vrfimportrtPresent,proto3" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,25,opt,name=sourceasrt_present,json=sourceasrtPresent,proto3" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,26,opt,name=sourcerd_present,json=sourcerdPresent,proto3" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,27,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent,proto3" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,28,opt,name=next_hop_id,json=nextHopId,proto3" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,29,opt,name=next_hop_id_refcount,json=nextHopIdRefcount,proto3" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,30,opt,name=ospf_area_id,json=ospfAreaId,proto3" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr []*Ipv6RibEdmAddr `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,32,opt,name=has_labelstk,json=hasLabelstk,proto3" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,33,opt,name=num_labels,json=numLabels,proto3" json:"num_labels,omitempty"`
	// Labels for this path
	Labelstk []uint32 `protobuf:"varint,34,rep,packed,name=labelstk,proto3" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,35,opt,name=binding_label,json=bindingLabel,proto3" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,36,opt,name=nhid_feid,json=nhidFeid,proto3" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid             uint64   `protobuf:"varint,37,opt,name=mpls_feid,json=mplsFeid,proto3" json:"mpls_feid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmPathItem) Reset()         { *m = Ipv6RibEdmPathItem{} }
func (m *Ipv6RibEdmPathItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPathItem) ProtoMessage()    {}
func (*Ipv6RibEdmPathItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{4}
}
func (m *Ipv6RibEdmPathItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Unmarshal(m, b)
}
func (m *Ipv6RibEdmPathItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Marshal(b, m, deterministic)
}
func (dst *Ipv6RibEdmPathItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmPathItem.Merge(dst, src)
}
func (m *Ipv6RibEdmPathItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmPathItem.Size(m)
}
func (m *Ipv6RibEdmPathItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmPathItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmPathItem proto.InternalMessageInfo

func (m *Ipv6RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetRemoteBackupAddr() []*Ipv6RibEdmAddr {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale,proto3" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored,proto3" json:"mirrored,omitempty"`
	// Merge disable
	MergeDisable uint32 `protobuf:"varint,5,opt,name=merge_disable,json=mergeDisable,proto3" json:"merge_disable,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,6,opt,name=redist_only,json=redistOnly,proto3" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,7,opt,name=label,proto3" json:"label,omitempty"`
	// Distance
	Distance             uint32   `protobuf:"varint,8,opt,name=distance,proto3" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmLocalLabel) Reset()         { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()    {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d, []int{5}
}
func (m *RibEdmLocalLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmLocalLabel.Unmarshal(m, b)
}
func (m *RibEdmLocalLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmLocalLabel.Marshal(b, m, deterministic)
}
func (dst *RibEdmLocalLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmLocalLabel.Merge(dst, src)
}
func (m *RibEdmLocalLabel) XXX_Size() int {
	return xxx_messageInfo_RibEdmLocalLabel.Size(m)
}
func (m *RibEdmLocalLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmLocalLabel.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmLocalLabel proto.InternalMessageInfo

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMergeDisable() uint32 {
	if m != nil {
		return m.MergeDisable
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.ipv6_rib_edm_route")
	proto.RegisterType((*Ipv6RibEdmAddr)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.ipv6_rib_edm_addr")
	proto.RegisterType((*Ipv6RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.ipv6_rib_edm_path")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rpl.as.protocol_routes.protocol_route.rib_edm_local_label")
}

func init() {
	proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d)
}

var fileDescriptor_ipv6_rib_edm_route_3b38ef336048fc0d = []byte{
	// 1527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x52, 0x1b, 0xcb,
	0x15, 0x2e, 0x81, 0x11, 0x52, 0x23, 0x61, 0x34, 0x10, 0x68, 0x2e, 0xf7, 0xfa, 0xca, 0xdc, 0xb8,
	0x0a, 0x52, 0x85, 0x9c, 0xb2, 0x1d, 0x92, 0x38, 0xbf, 0x18, 0x63, 0x5b, 0x31, 0x36, 0x44, 0x76,
	0xb9, 0x2a, 0xab, 0xa9, 0xd6, 0x74, 0x8f, 0xd4, 0xe5, 0xf9, 0x73, 0x77, 0x4b, 0x86, 0x65, 0x1e,
	0x26, 0xdb, 0xac, 0xb3, 0xc9, 0x7b, 0x64, 0x9b, 0x55, 0x2a, 0x4f, 0x90, 0x65, 0x52, 0xe7, 0x9c,
	0x9e, 0x91, 0x04, 0xce, 0x03, 0x78, 0x63, 0x73, 0xbe, 0xef, 0xeb, 0x9e, 0xee, 0x73, 0xfa, 0xfc,
	0x88, 0x71, 0x5d, 0x4c, 0x8f, 0x43, 0xa3, 0x87, 0xa1, 0x92, 0x69, 0x68, 0xf2, 0x89, 0x53, 0xbd,
	0xc2, 0xe4, 0x2e, 0x0f, 0xfe, 0x52, 0x8b, 0xb4, 0x8d, 0xf2, 0x50, 0xe7, 0x36, 0xbc, 0x32, 0xa1,
	0x2e, 0x50, 0x85, 0xf2, 0xbc, 0x50, 0xa6, 0x57, 0x2d, 0xb4, 0x4e, 0x0e, 0xaf, 0x7b, 0x53, 0x13,
	0x5b, 0xf8, 0xa7, 0x27, 0x62, 0xdb, 0x13, 0x71, 0xcf, 0xc2, 0xff, 0x56, 0xc4, 0x3d, 0xbf, 0x10,
	0xb7, 0x0e, 0x9d, 0x18, 0x26, 0x2a, 0xcc, 0x44, 0xaa, 0xec, 0xff, 0x23, 0xe8, 0xf3, 0x51, 0x9e,
	0xf4, 0x4c, 0x91, 0xf4, 0x84, 0xad, 0x6c, 0x52, 0xde, 0xb4, 0xf7, 0xff, 0x59, 0x63, 0x3b, 0xb7,
	0x2f, 0x11, 0xbe, 0x3e, 0xfb, 0xd3, 0xbb, 0x60, 0x97, 0x35, 0xa6, 0x26, 0xc6, 0x6d, 0x79, 0xad,
	0x5b, 0x3b, 0x68, 0x0e, 0x56, 0xa7, 0x26, 0x7e, 0x2b, 0x52, 0x15, 0xec, 0xb0, 0x55, 0xe1, 0x99,
	0x25, 0x64, 0xea, 0x82, 0x88, 0x5d, 0xd6, 0xb0, 0x25, 0xb3, 0x4c, 0x6b, 0xac, 0xa7, 0x0e, 0xd8,
	0xc6, 0xcd, 0xd3, 0xf2, 0x3b, 0x28, 0x59, 0x47, 0xfc, 0x3d, 0xc0, 0xa8, 0x5c, 0x67, 0x4b, 0xc2,
	0xf2, 0x15, 0xe4, 0x96, 0x84, 0x0d, 0x38, 0x5b, 0x15, 0x52, 0x1a, 0x65, 0x2d, 0xaf, 0xd3, 0x9e,
	0xde, 0x0c, 0x7e, 0x60, 0xed, 0xc2, 0xa8, 0x58, 0x5f, 0x85, 0x89, 0xca, 0x46, 0x6e, 0xcc, 0x57,
	0xbb, 0xb5, 0x83, 0xf6, 0xa0, 0x45, 0xe0, 0x39, 0x62, 0xfb, 0xff, 0x69, 0xb2, 0xe0, 0xf6, 0x1d,
	0x83, 0x6d, 0x56, 0x27, 0x19, 0x7f, 0x44, 0x57, 0x20, 0xeb, 0xf6, 0x9e, 0x8f, 0x6f, 0xef, 0x09,
	0x22, 0xba, 0xcc, 0x54, 0x19, 0xab, 0xf3, 0x8c, 0x3f, 0x21, 0x11, 0x82, 0x1f, 0x08, 0x0b, 0xbe,
	0x67, 0x6b, 0x95, 0xbb, 0xb5, 0xe4, 0x3f, 0x43, 0x09, 0x2b, 0xa1, 0xbe, 0xa4, 0x4f, 0x79, 0x01,
	0xfa, 0xe3, 0x18, 0x4f, 0xd2, 0x2a, 0x41, 0xf4, 0xc6, 0x37, 0xac, 0xa1, 0x33, 0xeb, 0x44, 0x16,
	0x29, 0xfe, 0x73, 0xe4, 0x2b, 0x3b, 0xd8, 0x63, 0xcd, 0x28, 0xd1, 0x2a, 0x73, 0xb0, 0xff, 0x2f,
	0x70, 0xff, 0x06, 0x01, 0x7d, 0x19, 0x7c, 0xc7, 0x98, 0x77, 0xf8, 0x75, 0xa1, 0xf8, 0x2f, 0x91,
	0x6d, 0x92, 0xab, 0xaf, 0x0b, 0xdc, 0xb7, 0x30, 0x3a, 0x37, 0xda, 0x5d, 0xf3, 0xa7, 0xb4, 0xb4,
	0xb4, 0x31, 0x8c, 0x53, 0x49, 0x0b, 0x7f, 0x85, 0xdc, 0xaa, 0x9d, 0x4a, 0x5c, 0xb6, 0xc5, 0x56,
	0xe2, 0x44, 0x8c, 0x2c, 0xff, 0x35, 0xe2, 0x64, 0x04, 0x0f, 0xd8, 0xba, 0xba, 0x72, 0x2a, 0x93,
	0x4a, 0x86, 0x44, 0xff, 0xa6, 0x5b, 0x3b, 0xb8, 0x33, 0x68, 0x97, 0xe8, 0x0b, 0x94, 0x6d, 0xb0,
	0x65, 0x27, 0x46, 0xfc, 0xb7, 0xb8, 0x14, 0xfe, 0x84, 0x53, 0x48, 0xed, 0x6f, 0xf7, 0x3b, 0x3a,
	0x45, 0x69, 0x07, 0x47, 0x2c, 0x90, 0xda, 0x3b, 0x38, 0xac, 0x54, 0xbf, 0x47, 0x55, 0xa7, 0x62,
	0x9e, 0x97, 0xf2, 0x6d, 0x56, 0x4f, 0x95, 0x33, 0x3a, 0xe2, 0x27, 0x28, 0xf1, 0x16, 0x86, 0x41,
	0xb8, 0xb1, 0x0d, 0xa3, 0x7c, 0x92, 0x39, 0xfe, 0xcc, 0x87, 0x01, 0xa0, 0x53, 0x40, 0xe0, 0x3b,
	0xc2, 0x39, 0xa3, 0x87, 0xe0, 0x2c, 0x2d, 0x55, 0xe6, 0xc0, 0x27, 0xa7, 0xf4, 0x9d, 0x8a, 0xe9,
	0x7b, 0x02, 0xa2, 0xe6, 0x8c, 0x88, 0x63, 0x1d, 0x85, 0x3a, 0x93, 0xea, 0x8a, 0x3f, 0xa7, 0xd8,
	0x7b, 0xb0, 0x0f, 0x58, 0x70, 0x58, 0xbe, 0xf6, 0xc2, 0xa8, 0x48, 0x49, 0x05, 0x27, 0x3f, 0x43,
	0xdd, 0x5d, 0xc4, 0x2f, 0x2b, 0x18, 0x82, 0xf8, 0x29, 0xb7, 0xe1, 0xc8, 0xe4, 0x93, 0x82, 0xbf,
	0x20, 0x1f, 0x7c, 0xca, 0xed, 0x4b, 0xb0, 0x21, 0x12, 0x71, 0x92, 0x7f, 0x0e, 0xc1, 0x6d, 0x2f,
	0x29, 0x12, 0x60, 0xbf, 0x17, 0x23, 0x58, 0x17, 0x7f, 0x96, 0x61, 0x94, 0x08, 0x6b, 0xf9, 0x2b,
	0x5a, 0x17, 0x7f, 0x96, 0xa7, 0x60, 0x03, 0x59, 0xe8, 0xc8, 0x5f, 0xb9, 0xef, 0xc3, 0xab, 0x23,
	0xba, 0xf0, 0x36, 0xab, 0x8b, 0xc8, 0xe9, 0xa9, 0xe2, 0x7f, 0xe8, 0xd6, 0x0e, 0x1a, 0x03, 0x6f,
	0x05, 0xdf, 0xb2, 0x66, 0xe5, 0x56, 0xfe, 0x1a, 0xa9, 0x19, 0x10, 0xfc, 0x94, 0x6d, 0xcd, 0xc2,
	0x81, 0x4f, 0x94, 0x1e, 0xed, 0x39, 0x3e, 0xca, 0x59, 0xa8, 0x2e, 0x81, 0xc2, 0xa7, 0xbb, 0xc7,
	0xe8, 0xbd, 0x85, 0x62, 0xa4, 0xf8, 0x1b, 0x3a, 0x04, 0x02, 0x27, 0x23, 0x05, 0x61, 0x21, 0x32,
	0x11, 0x43, 0x95, 0xf0, 0xb7, 0x14, 0x16, 0x84, 0xce, 0x01, 0x81, 0xb4, 0x2f, 0xcf, 0x72, 0x41,
	0x37, 0x9f, 0xce, 0x12, 0xcb, 0x0d, 0x93, 0x2a, 0xf7, 0x2e, 0xf1, 0xa9, 0x31, 0x37, 0x4c, 0xca,
	0xcc, 0xfb, 0x09, 0xeb, 0xd0, 0xde, 0x69, 0x2e, 0x75, 0x7c, 0x1d, 0x3a, 0x9d, 0x2a, 0xfe, 0x47,
	0x94, 0x91, 0xfb, 0xdf, 0x20, 0xfe, 0x5e, 0xa7, 0x2a, 0xf8, 0x47, 0xad, 0xcc, 0x13, 0x78, 0x12,
	0x7c, 0xd0, 0xad, 0x1d, 0xac, 0x3d, 0xfa, 0x5b, 0xad, 0xf7, 0x55, 0x14, 0xf0, 0xde, 0x42, 0x61,
	0x83, 0x0b, 0xf8, 0x14, 0xbf, 0x14, 0x6e, 0xbc, 0x7f, 0xc8, 0x3a, 0x0b, 0x3c, 0x94, 0x4d, 0x48,
	0xe0, 0xa9, 0x48, 0x26, 0x65, 0x4d, 0x27, 0x63, 0xff, 0xbf, 0xb5, 0x1b, 0x5a, 0xd8, 0x2b, 0xf8,
	0xf7, 0x97, 0x50, 0x5e, 0xeb, 0x2e, 0x1f, 0xac, 0x3d, 0xfa, 0xfb, 0x57, 0xeb, 0xa2, 0x50, 0x3b,
	0x95, 0x0e, 0xd6, 0x01, 0x1f, 0xe8, 0xe1, 0x99, 0x4c, 0xd1, 0x59, 0x7f, 0x6d, 0xb1, 0xed, 0x2f,
	0x4b, 0xe7, 0x1b, 0x50, 0x6d, 0xb1, 0x01, 0x1d, 0xb1, 0x40, 0x67, 0x71, 0x6e, 0x52, 0xe1, 0x20,
	0x2b, 0x6c, 0x3e, 0x31, 0x51, 0xd9, 0x13, 0x3b, 0x73, 0xcc, 0x3b, 0x24, 0xa0, 0x24, 0x4f, 0x8f,
	0xc3, 0x4c, 0x5d, 0xb9, 0x71, 0x5e, 0xf8, 0x06, 0xd9, 0x9c, 0x1e, 0xbf, 0x25, 0x00, 0xaa, 0xa8,
	0xce, 0x9c, 0x32, 0xb1, 0x88, 0x16, 0x1a, 0x64, 0xbb, 0x42, 0x31, 0xad, 0x66, 0x85, 0x6e, 0xe5,
	0x66, 0xa1, 0x4b, 0x72, 0x21, 0x43, 0x4f, 0xd6, 0x29, 0xa3, 0x00, 0x7a, 0x43, 0x02, 0xce, 0x56,
	0xb1, 0x38, 0x1f, 0x3f, 0xc1, 0x46, 0x79, 0x67, 0x50, 0x9a, 0xb3, 0xaa, 0xde, 0x98, 0xaf, 0xea,
	0xd8, 0x9f, 0xf4, 0x54, 0x38, 0xe5, 0x8b, 0x7a, 0xb3, 0x6c, 0x85, 0x08, 0x52, 0x4d, 0xdf, 0x66,
	0xf5, 0x24, 0xcf, 0x0b, 0x25, 0x39, 0xa3, 0x62, 0x42, 0x56, 0x70, 0xc8, 0x3a, 0x70, 0xd1, 0x70,
	0x9c, 0x17, 0x3e, 0x78, 0x5a, 0xf2, 0x35, 0xdc, 0x60, 0x1d, 0x88, 0x57, 0x79, 0x81, 0x2d, 0xbf,
	0xbf, 0x28, 0xad, 0x46, 0x8e, 0x16, 0xcd, 0x06, 0x5e, 0xfa, 0xc1, 0x4f, 0x1e, 0x47, 0x6c, 0xf3,
	0xc6, 0xae, 0x28, 0x6e, 0xa3, 0x78, 0x63, 0x7e, 0x5f, 0x94, 0x77, 0x59, 0xab, 0x92, 0x8b, 0x58,
	0xf3, 0x75, 0xf2, 0x89, 0xd7, 0x9d, 0xc4, 0x3a, 0xd8, 0x67, 0xed, 0x4a, 0x61, 0x41, 0x72, 0x17,
	0x25, 0x6b, 0x5e, 0xf2, 0x4e, 0xc4, 0xfa, 0x66, 0xa9, 0xda, 0xb8, 0x55, 0xaa, 0xf6, 0x58, 0xd3,
	0x4d, 0xb2, 0x4c, 0x61, 0x9f, 0xef, 0x50, 0xa1, 0x23, 0xa0, 0x2f, 0x71, 0xd0, 0x10, 0x6e, 0xac,
	0x25, 0x0f, 0x28, 0x5c, 0x64, 0x81, 0x77, 0x87, 0x22, 0xfa, 0x38, 0x29, 0x42, 0x4f, 0x6f, 0x92,
	0x77, 0x09, 0xbc, 0x24, 0xd1, 0x21, 0xeb, 0x18, 0x15, 0x87, 0x51, 0xe6, 0xc2, 0x3c, 0x0e, 0x89,
	0xe2, 0x5b, 0xe4, 0x45, 0xa3, 0xe2, 0xd3, 0xcc, 0x5d, 0xc4, 0xcf, 0x10, 0x0d, 0x4e, 0xd9, 0xbd,
	0x6c, 0x92, 0x0e, 0x95, 0x01, 0x65, 0xd5, 0x8d, 0xa3, 0x3c, 0x4d, 0x27, 0x99, 0x76, 0x5a, 0x59,
	0xfe, 0x23, 0x5c, 0xb7, 0x47, 0xaa, 0x8b, 0xf8, 0xcc, 0x6b, 0x4e, 0x67, 0x92, 0xe0, 0x3e, 0x6b,
	0xa5, 0xd3, 0x02, 0xea, 0xbb, 0xb2, 0x2a, 0x73, 0x7c, 0x1b, 0x63, 0xba, 0x06, 0xd8, 0x25, 0x41,
	0xf0, 0x4a, 0xe1, 0xc0, 0xc6, 0x55, 0xa2, 0x1d, 0x14, 0xb5, 0x09, 0x2d, 0x65, 0x0f, 0xd9, 0xe6,
	0xd4, 0xc4, 0x3a, 0x2d, 0x72, 0xe3, 0xe6, 0xb4, 0x1c, 0xb5, 0xc1, 0x1c, 0x55, 0x2e, 0x38, 0x62,
	0x01, 0xe5, 0x8f, 0xb0, 0x73, 0xfa, 0x5d, 0xd4, 0x77, 0x66, 0x4c, 0x29, 0x3f, 0x64, 0x1b, 0x04,
	0x1a, 0x59, 0x89, 0xbf, 0x41, 0xf1, 0xdd, 0x12, 0x2f, 0xa5, 0x4f, 0xd9, 0xae, 0x55, 0xa3, 0x54,
	0x65, 0x4e, 0xc9, 0x32, 0xfb, 0xaa, 0x35, 0x7b, 0xb8, 0x66, 0xa7, 0x12, 0xf8, 0x64, 0x2c, 0xd7,
	0xde, 0x63, 0x6b, 0xd5, 0xfb, 0xd0, 0x92, 0x7f, 0x4b, 0x63, 0x94, 0x7f, 0x1d, 0x7d, 0x19, 0x3c,
	0x64, 0x5b, 0x73, 0x7c, 0x68, 0x54, 0x4c, 0x3d, 0xf7, 0x3b, 0x1a, 0x1f, 0x2a, 0xe1, 0xc0, 0x13,
	0xf0, 0x24, 0x73, 0x5b, 0xc4, 0xa1, 0x30, 0x4a, 0xc0, 0x8e, 0xf7, 0xf0, 0xe9, 0x32, 0xc0, 0x4e,
	0x8c, 0x12, 0x7d, 0x19, 0xfc, 0xab, 0xc6, 0x02, 0xa3, 0xd2, 0xdc, 0x29, 0x1f, 0x70, 0x2c, 0xdc,
	0xfc, 0x7b, 0x2c, 0xbb, 0x5f, 0x67, 0x67, 0x82, 0x0b, 0x0c, 0x36, 0xe8, 0x52, 0xf4, 0x5c, 0x4f,
	0xa0, 0x17, 0xdd, 0x67, 0xad, 0xb1, 0xb0, 0x94, 0x56, 0xd6, 0x7d, 0xe4, 0x5d, 0x7a, 0x6d, 0x63,
	0x61, 0xcf, 0x3d, 0x04, 0x25, 0x33, 0x9b, 0xa4, 0x5e, 0xc2, 0xef, 0x7b, 0xf7, 0x4f, 0x52, 0x12,
	0xc0, 0xfc, 0x58, 0xad, 0xde, 0xef, 0x2e, 0x43, 0xe2, 0x95, 0x36, 0x26, 0x98, 0xce, 0xa4, 0xce,
	0x46, 0x3e, 0x71, 0x7f, 0xf0, 0x09, 0x46, 0x60, 0x95, 0xba, 0xd9, 0x58, 0xcb, 0x30, 0x56, 0x5a,
	0xf2, 0x1f, 0x63, 0x55, 0x6c, 0x00, 0xf0, 0x42, 0x69, 0x09, 0x64, 0x5a, 0x24, 0x96, 0xc8, 0x07,
	0x44, 0x02, 0x00, 0xe4, 0xfe, 0x9f, 0x97, 0xd8, 0x66, 0x79, 0xbf, 0x24, 0x8f, 0x44, 0x42, 0x5f,
	0xb9, 0x3d, 0xd5, 0xd7, 0xbe, 0x30, 0xd5, 0x2f, 0x4c, 0xee, 0x4b, 0x37, 0x26, 0xf7, 0x2d, 0xb6,
	0x62, 0x9d, 0x48, 0xe8, 0x27, 0x54, 0x7b, 0x40, 0x06, 0x5c, 0x35, 0xd5, 0xc6, 0xe4, 0x46, 0x49,
	0xec, 0x0b, 0xed, 0x41, 0x65, 0xc3, 0x37, 0x53, 0x65, 0x46, 0x0a, 0xc6, 0x64, 0x88, 0x98, 0xef,
	0x0c, 0x2d, 0x04, 0x9f, 0x13, 0x86, 0x65, 0x4c, 0xc1, 0x1c, 0x1d, 0xe6, 0x59, 0x72, 0x5d, 0xf6,
	0x07, 0x82, 0x2e, 0xb2, 0xe4, 0x1a, 0xbe, 0x4b, 0x8e, 0xa2, 0x9f, 0x51, 0x64, 0x2c, 0x8c, 0xe8,
	0x8d, 0xc5, 0x11, 0x7d, 0x58, 0xc7, 0x4b, 0x3d, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x20, 0x12, 0xcc, 0x0a, 0x0f, 0x00, 0x00,
}
