// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_destination_kw_dest_next_hop_routes_dest_next_hop_route

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
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmRoute_KEYS) Reset()         { *m = Ipv6RibEdmRoute_KEYS{} }
func (m *Ipv6RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{0}
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

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
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
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{1}
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
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{2}
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
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{3}
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
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{4}
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
	return fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426, []int{5}
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
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.ipv6_rib_edm_route")
	proto.RegisterType((*Ipv6RibEdmAddr)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.ipv6_rib_edm_addr")
	proto.RegisterType((*Ipv6RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.ipv6_rib_edm_path")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_next_hop_routes.dest_next_hop_route.rib_edm_local_label")
}

func init() {
	proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426)
}

var fileDescriptor_ipv6_rib_edm_route_6bd76dcde1f01426 = []byte{
	// 1513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x72, 0x13, 0xbd,
	0x19, 0x1e, 0x27, 0x24, 0xb1, 0x15, 0x3b, 0xc4, 0x9b, 0x34, 0x51, 0x08, 0x3f, 0x26, 0x94, 0x99,
	0xa4, 0x33, 0x98, 0x0e, 0xd0, 0xb4, 0xa5, 0xbf, 0x21, 0x04, 0x70, 0x09, 0x24, 0x35, 0x0c, 0x33,
	0x3d, 0xd2, 0xc8, 0x2b, 0xad, 0xad, 0x61, 0x77, 0xb5, 0x48, 0xb2, 0x93, 0x1c, 0xb6, 0x57, 0xd3,
	0x83, 0x9e, 0xf4, 0xb8, 0x27, 0xbd, 0x8f, 0xde, 0x41, 0xa7, 0x57, 0xd0, 0xa3, 0x6f, 0xf4, 0xbe,
	0xda, 0xb5, 0xf3, 0xf3, 0x5d, 0x00, 0x27, 0xe0, 0xf7, 0x79, 0x1e, 0x69, 0xa5, 0xf7, 0x57, 0x21,
	0x54, 0x15, 0x93, 0x7d, 0x66, 0xd4, 0x80, 0x49, 0x91, 0x31, 0xa3, 0xc7, 0x4e, 0x76, 0x0b, 0xa3,
	0x9d, 0x8e, 0xfe, 0x5e, 0x8b, 0x95, 0x8d, 0x35, 0x53, 0xda, 0xb2, 0x73, 0xc3, 0x54, 0x01, 0x2a,
	0x90, 0xeb, 0x42, 0x9a, 0x6e, 0xb9, 0xb0, 0x3b, 0x31, 0x89, 0xf5, 0xff, 0x74, 0x79, 0x62, 0xbb,
	0x3c, 0xe9, 0x5a, 0xff, 0xbf, 0xe5, 0x49, 0x37, 0x2c, 0x81, 0x4d, 0x99, 0xe3, 0x83, 0x54, 0xb2,
	0x9c, 0x67, 0xd2, 0xfe, 0x18, 0xd1, 0x15, 0xd2, 0x3a, 0x95, 0x73, 0xa7, 0x74, 0xce, 0xbe, 0x9e,
	0x81, 0xc9, 0x72, 0x79, 0xee, 0xd8, 0x48, 0x17, 0xa8, 0xb6, 0x37, 0x81, 0x3b, 0xff, 0xa8, 0x91,
	0xcd, 0xeb, 0xf7, 0x60, 0xef, 0x8f, 0xfe, 0xf2, 0x29, 0xda, 0x22, 0xf5, 0x89, 0x49, 0x60, 0x7f,
	0x5a, 0xeb, 0xd4, 0x76, 0x1b, 0xfd, 0xa5, 0x89, 0x49, 0x3e, 0xf2, 0x4c, 0x46, 0x9b, 0x64, 0x89,
	0x07, 0x66, 0x0e, 0x98, 0x45, 0x8e, 0xc4, 0x16, 0xa9, 0xdb, 0x92, 0x99, 0xc7, 0x35, 0x36, 0x50,
	0xbb, 0x64, 0xf5, 0xea, 0xb1, 0xe9, 0x2d, 0x90, 0xac, 0x00, 0xfe, 0xd9, 0xc3, 0xa0, 0xa4, 0x64,
	0x89, 0x0b, 0x61, 0xa4, 0xb5, 0x74, 0x01, 0xf7, 0x08, 0xe6, 0xce, 0xff, 0x1b, 0x24, 0xba, 0x7e,
	0xdc, 0x68, 0x83, 0x2c, 0x16, 0x46, 0x26, 0xea, 0x9c, 0x3e, 0xc3, 0xd3, 0xa0, 0x15, 0x3d, 0x22,
	0x2d, 0xfc, 0xc5, 0x52, 0x99, 0x0f, 0xdd, 0x88, 0x3e, 0xef, 0xd4, 0x76, 0x5b, 0xfd, 0x26, 0x82,
	0xc7, 0x80, 0x79, 0x11, 0x9e, 0x6b, 0x22, 0x8d, 0x55, 0x3a, 0xa7, 0x2f, 0x50, 0x04, 0xe0, 0x17,
	0xc4, 0xa2, 0x07, 0x64, 0x19, 0x62, 0x1b, 0xeb, 0x94, 0x29, 0x41, 0x7f, 0x01, 0x12, 0x52, 0x42,
	0x3d, 0x81, 0x9f, 0x0a, 0x02, 0xb8, 0xda, 0x3e, 0x9c, 0xa4, 0x59, 0x82, 0x70, 0xb1, 0x3b, 0xa4,
	0xae, 0x72, 0xeb, 0x78, 0x1e, 0x4b, 0xfa, 0x4b, 0xe0, 0x2b, 0x3b, 0xda, 0x26, 0x8d, 0x38, 0x55,
	0x32, 0x77, 0x7e, 0xff, 0x5f, 0xc1, 0xfe, 0x75, 0x04, 0x7a, 0x22, 0xba, 0x47, 0x48, 0xf0, 0xdd,
	0x45, 0x21, 0xe9, 0xaf, 0x81, 0x6d, 0xa0, 0xd7, 0x2e, 0x0a, 0xd8, 0xb7, 0x30, 0x4a, 0x1b, 0xe5,
	0x2e, 0xe8, 0x4b, 0x5c, 0x5a, 0xda, 0x10, 0x91, 0x89, 0xc0, 0x85, 0xbf, 0x01, 0x6e, 0xc9, 0x4e,
	0x04, 0x2c, 0x5b, 0x27, 0x0b, 0x49, 0xca, 0x87, 0x96, 0xfe, 0x16, 0x70, 0x34, 0xa2, 0xc7, 0x64,
	0x45, 0x9e, 0x3b, 0x99, 0x0b, 0x29, 0x18, 0xd2, 0xbf, 0xeb, 0xd4, 0x76, 0x6f, 0xf5, 0x5b, 0x25,
	0xfa, 0x06, 0x64, 0xab, 0x64, 0xde, 0xf1, 0x21, 0xfd, 0x3d, 0x2c, 0xf5, 0x3f, 0xfd, 0x29, 0x84,
	0x0a, 0xb7, 0xfb, 0x03, 0x9e, 0xa2, 0xb4, 0xa3, 0x27, 0x24, 0x12, 0x2a, 0x38, 0x98, 0x55, 0xaa,
	0x3f, 0x82, 0xaa, 0x5d, 0x31, 0xaf, 0x4b, 0xf9, 0x06, 0x59, 0xcc, 0xa4, 0x33, 0x2a, 0xa6, 0x07,
	0x20, 0x09, 0x16, 0x84, 0x81, 0xbb, 0x91, 0x65, 0xb1, 0x1e, 0xe7, 0x8e, 0xbe, 0x0a, 0x61, 0xf0,
	0xd0, 0xa1, 0x47, 0xfc, 0x77, 0xb8, 0x73, 0x46, 0x0d, 0xbc, 0xb3, 0x94, 0x90, 0xb9, 0xf3, 0x3e,
	0x39, 0xc4, 0xef, 0x54, 0x4c, 0x2f, 0x10, 0x3e, 0x6a, 0xce, 0xf0, 0x24, 0x51, 0x31, 0x53, 0xb9,
	0x90, 0xe7, 0xf4, 0x35, 0xc6, 0x3e, 0x80, 0x3d, 0x8f, 0x45, 0x7b, 0x65, 0xe2, 0x16, 0x46, 0xc6,
	0x52, 0x48, 0x7f, 0xf2, 0x23, 0xd0, 0xdd, 0x06, 0xfc, 0xb4, 0x82, 0x7d, 0x10, 0xbf, 0x69, 0xcb,
	0x86, 0x46, 0x8f, 0x0b, 0xfa, 0x06, 0x7d, 0xf0, 0x4d, 0xdb, 0xb7, 0xde, 0xf6, 0x91, 0x48, 0x52,
	0x7d, 0xc6, 0xbc, 0xdb, 0xde, 0x62, 0x24, 0xbc, 0xfd, 0x99, 0x0f, 0xfd, 0xba, 0xe4, 0x4c, 0xb0,
	0x38, 0xe5, 0xd6, 0xd2, 0x77, 0xb8, 0x2e, 0x39, 0x13, 0x87, 0xde, 0xf6, 0x64, 0xa1, 0xe2, 0x70,
	0xe5, 0x5e, 0x08, 0xaf, 0x8a, 0xf1, 0xc2, 0x1b, 0x64, 0x91, 0xc7, 0x4e, 0x4d, 0x24, 0xfd, 0x53,
	0xa7, 0xb6, 0x5b, 0xef, 0x07, 0x2b, 0xba, 0x4b, 0x1a, 0x95, 0x5b, 0xe9, 0x7b, 0xa0, 0xa6, 0x40,
	0xf4, 0x73, 0xb2, 0x3e, 0x0d, 0x07, 0xa4, 0x28, 0x26, 0xed, 0x31, 0x24, 0xe5, 0x34, 0x54, 0xa7,
	0x9e, 0x82, 0xd4, 0xdd, 0x26, 0x98, 0x6f, 0x8c, 0x0f, 0x25, 0xfd, 0x80, 0x87, 0x00, 0xe0, 0x60,
	0x28, 0x7d, 0x58, 0x90, 0x4c, 0xf9, 0x40, 0xa6, 0xf4, 0x23, 0x86, 0x05, 0xa0, 0x63, 0x8f, 0xf8,
	0x8a, 0x2e, 0xcf, 0x72, 0x82, 0x37, 0x9f, 0x4c, 0x0b, 0xcb, 0x0d, 0xd2, 0xaa, 0xf6, 0x4e, 0x21,
	0xd5, 0x88, 0x1b, 0xa4, 0x65, 0xe5, 0xfd, 0x8c, 0xb4, 0x71, 0xef, 0x4c, 0x0b, 0x95, 0x5c, 0x30,
	0xa7, 0x32, 0x49, 0xff, 0x0c, 0x32, 0x74, 0xff, 0x07, 0xc0, 0x3f, 0xab, 0x4c, 0x46, 0xff, 0xa9,
	0x95, 0x75, 0xe2, 0x53, 0x82, 0xf6, 0x3b, 0xb5, 0xdd, 0xe5, 0x67, 0xff, 0xaa, 0x75, 0xbf, 0x97,
	0x76, 0xdc, 0xbd, 0xd4, 0xdb, 0xfc, 0x1d, 0x42, 0x95, 0x9f, 0x72, 0x37, 0xda, 0xd9, 0x23, 0xed,
	0x4b, 0xbc, 0x6f, 0x8a, 0xbe, 0x86, 0x27, 0x3c, 0x1d, 0x97, 0x1d, 0x1a, 0x8d, 0x9d, 0xbf, 0xcd,
	0x5d, 0xd1, 0xfa, 0xbd, 0xa2, 0xff, 0xd5, 0x6e, 0x40, 0x69, 0xad, 0x33, 0xbf, 0xbb, 0xfc, 0xec,
	0xdf, 0xdf, 0xb3, 0x97, 0x98, 0x72, 0x32, 0xeb, 0xaf, 0x78, 0xbc, 0xaf, 0x06, 0x47, 0x22, 0x03,
	0x7f, 0xfd, 0xb3, 0x49, 0x36, 0x6e, 0x96, 0xce, 0x4e, 0x98, 0xda, 0xa5, 0x09, 0xe3, 0x1b, 0x88,
	0xca, 0x13, 0x6d, 0x32, 0x3c, 0x96, 0xd5, 0x63, 0x13, 0x97, 0x43, 0xae, 0x3d, 0xc3, 0x7c, 0x02,
	0xc2, 0x37, 0xe6, 0xc9, 0x3e, 0x9c, 0x6f, 0xa4, 0x8b, 0x30, 0xf1, 0x1a, 0x93, 0xfd, 0x8f, 0x08,
	0xf8, 0x5e, 0xaa, 0x72, 0x27, 0x4d, 0xc2, 0xe3, 0x4b, 0x13, 0xaf, 0x55, 0xa1, 0x50, 0x5c, 0xd3,
	0x76, 0xb7, 0x70, 0xb5, 0xdd, 0xa5, 0x9a, 0x0b, 0x16, 0xc8, 0x45, 0xac, 0x2b, 0x0f, 0x7d, 0x40,
	0x01, 0x25, 0x4b, 0xd0, 0xa2, 0xf7, 0x5f, 0xd0, 0x25, 0x28, 0x89, 0xd2, 0x9c, 0xf6, 0xf6, 0xfa,
	0x6c, 0x6f, 0x87, 0x29, 0xa5, 0x26, 0xdc, 0xc9, 0xd0, 0xda, 0x1b, 0xe5, 0x40, 0x04, 0x10, 0x3b,
	0xfb, 0x06, 0x59, 0x4c, 0xb5, 0x2e, 0xa4, 0xa0, 0x04, 0x5b, 0x0a, 0x5a, 0xd1, 0x1e, 0x69, 0x57,
	0x81, 0xc0, 0x28, 0x2a, 0x41, 0x97, 0x61, 0x83, 0x15, 0x4f, 0xbc, 0xd3, 0x05, 0xcc, 0xf0, 0xde,
	0x65, 0x69, 0xf5, 0x86, 0x68, 0xe2, 0xb0, 0x0f, 0xd2, 0x2f, 0xe1, 0x29, 0xf1, 0x84, 0xac, 0x5d,
	0xd9, 0x15, 0xc4, 0x2d, 0x10, 0xaf, 0xce, 0xee, 0x0b, 0xf2, 0x0e, 0x69, 0x56, 0x72, 0x9e, 0x28,
	0xba, 0x82, 0x3e, 0x09, 0xba, 0x83, 0x44, 0x45, 0x3b, 0xa4, 0x55, 0x29, 0xac, 0x97, 0xdc, 0x06,
	0xc9, 0x72, 0x90, 0x7c, 0xe2, 0x89, 0xba, 0xda, 0xb0, 0x56, 0xaf, 0x35, 0xac, 0x6d, 0xd2, 0x70,
	0xe3, 0x3c, 0x97, 0x30, 0xed, 0xdb, 0xd8, 0xee, 0x10, 0xe8, 0x09, 0x78, 0x6e, 0x70, 0x37, 0x52,
	0x82, 0x46, 0x18, 0x2e, 0xb4, 0xbc, 0x77, 0x07, 0x3c, 0xfe, 0x3a, 0x2e, 0x58, 0xa0, 0xd7, 0xd0,
	0xbb, 0x08, 0x9e, 0xa2, 0x68, 0x8f, 0xb4, 0x8d, 0x4c, 0x58, 0x9c, 0x3b, 0xa6, 0x13, 0x86, 0x14,
	0x5d, 0x47, 0x2f, 0x1a, 0x99, 0x1c, 0xe6, 0xee, 0x24, 0x79, 0x05, 0x68, 0x74, 0x48, 0xee, 0xe7,
	0xe3, 0x6c, 0x20, 0x8d, 0x57, 0x56, 0x33, 0x39, 0xd6, 0x59, 0x36, 0xce, 0x95, 0x53, 0xd2, 0xd2,
	0x9f, 0xc0, 0xba, 0x6d, 0x54, 0x9d, 0x24, 0x47, 0x41, 0x73, 0x38, 0x95, 0x44, 0x0f, 0x49, 0x33,
	0x9b, 0x14, 0xbe, 0xcb, 0x4b, 0x2b, 0x73, 0x47, 0x37, 0x20, 0xa6, 0xcb, 0x1e, 0x3b, 0x45, 0xc8,
	0x67, 0xa9, 0x3f, 0xb0, 0x71, 0x95, 0x68, 0x13, 0x44, 0x2d, 0x44, 0x4b, 0xd9, 0x53, 0xb2, 0x36,
	0x31, 0x89, 0xca, 0x0a, 0x6d, 0xdc, 0x8c, 0x96, 0x82, 0x36, 0x9a, 0xa1, 0xca, 0x05, 0x4f, 0x48,
	0x84, 0xf5, 0xc3, 0xed, 0x8c, 0x7e, 0x0b, 0xf4, 0xed, 0x29, 0x53, 0xca, 0xf7, 0xc8, 0x2a, 0x82,
	0x46, 0x54, 0xe2, 0x3b, 0x20, 0xbe, 0x5d, 0xe2, 0xa5, 0xf4, 0x25, 0xd9, 0xb2, 0x72, 0x98, 0xc9,
	0xdc, 0x49, 0x51, 0x56, 0x5f, 0xb5, 0x66, 0x1b, 0xd6, 0x6c, 0x56, 0x82, 0x50, 0x8c, 0xe5, 0xda,
	0xfb, 0x64, 0xb9, 0xca, 0x0f, 0x25, 0xe8, 0x5d, 0x7c, 0x4c, 0x85, 0xec, 0xe8, 0x89, 0xe8, 0x29,
	0x59, 0x9f, 0xe1, 0x99, 0x91, 0x09, 0x4e, 0xde, 0x7b, 0xf8, 0x88, 0xa8, 0x84, 0xfd, 0x40, 0xf8,
	0x94, 0xd4, 0xb6, 0x48, 0x18, 0x37, 0x92, 0xfb, 0x1d, 0xef, 0x43, 0xea, 0x12, 0x8f, 0x1d, 0x18,
	0xc9, 0x7b, 0x22, 0xfa, 0x6f, 0x8d, 0x44, 0x46, 0x66, 0xda, 0xc9, 0x10, 0x70, 0xe8, 0xdd, 0xf4,
	0x01, 0x74, 0xde, 0xef, 0x76, 0x3e, 0xf9, 0x3b, 0xf4, 0x57, 0xf1, 0x5e, 0x98, 0xb1, 0x07, 0x7e,
	0x22, 0x3d, 0x24, 0xcd, 0x11, 0xb7, 0x58, 0x59, 0xd6, 0x7d, 0xa5, 0x1d, 0x4c, 0xb8, 0x11, 0xb7,
	0xc7, 0x01, 0xf2, 0x5d, 0x33, 0x1f, 0x67, 0x41, 0x42, 0x1f, 0x86, 0x08, 0x8c, 0x33, 0x14, 0xf8,
	0x87, 0x64, 0xb5, 0x7a, 0xa7, 0x33, 0xef, 0x6b, 0xaf, 0xb4, 0xa1, 0xc6, 0x54, 0x2e, 0x54, 0x3e,
	0x0c, 0xb5, 0xfb, 0x28, 0xd4, 0x18, 0x82, 0x55, 0xf5, 0xe6, 0x23, 0x25, 0x58, 0x22, 0x95, 0xa0,
	0x3f, 0x85, 0xc6, 0x58, 0xf7, 0xc0, 0x1b, 0xa9, 0x84, 0x27, 0xb3, 0x22, 0xb5, 0x48, 0x3e, 0x46,
	0xd2, 0x03, 0x9e, 0xdc, 0xf9, 0xeb, 0x1c, 0x59, 0x2b, 0xef, 0x97, 0xea, 0x98, 0xa7, 0xf8, 0x95,
	0xeb, 0xcf, 0xfb, 0xda, 0x0d, 0xcf, 0xfb, 0x4b, 0x4f, 0xf8, 0xb9, 0x2b, 0x4f, 0xf8, 0x75, 0xb2,
	0x60, 0x1d, 0x4f, 0xf1, 0xcf, 0xa2, 0x56, 0x1f, 0x0d, 0x7f, 0xd5, 0x4c, 0x19, 0xa3, 0x8d, 0x14,
	0x30, 0x1a, 0x5a, 0xfd, 0xca, 0xf6, 0xdf, 0xcc, 0xa4, 0x19, 0x4a, 0xff, 0x5e, 0xf6, 0xa1, 0x0b,
	0xc3, 0xa1, 0x09, 0xe0, 0x6b, 0xc4, 0xa0, 0x93, 0x49, 0xff, 0xa0, 0x66, 0x3a, 0x4f, 0x2f, 0xca,
	0x11, 0x81, 0xd0, 0x49, 0x9e, 0x5e, 0xf8, 0xef, 0xa2, 0xa3, 0x96, 0xf0, 0xbb, 0x78, 0x9f, 0xd9,
	0xb7, 0x7a, 0xfd, 0xf2, 0x5b, 0x7d, 0xb0, 0x08, 0x97, 0x7a, 0xfe, 0x43, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0x9b, 0x2f, 0xd6, 0xe1, 0x0e, 0x00, 0x00,
}
