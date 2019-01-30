// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_non_as_protocol_routes_protocol_route

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
type Ipv4RibEdmRoute_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()         { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb, []int{0}
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmRoute_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Merge(dst, src)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Size(m)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmRoute_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmRoute_KEYS proto.InternalMessageInfo

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv4RibEdmRoute struct {
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
	RoutePath            *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ipv4RibEdmRoute) Reset()         { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()    {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb, []int{1}
}
func (m *Ipv4RibEdmRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute.Merge(dst, src)
}
func (m *Ipv4RibEdmRoute) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmRoute.Size(m)
}
func (m *Ipv4RibEdmRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmRoute.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmRoute proto.InternalMessageInfo

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath       []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath,proto3" json:"ipv4_rib_edm_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ipv4RibEdmPath) Reset()         { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()    {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb, []int{2}
}
func (m *Ipv4RibEdmPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPath.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPath.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPath.Merge(dst, src)
}
func (m *Ipv4RibEdmPath) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmPath.Size(m)
}
func (m *Ipv4RibEdmPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmPath proto.InternalMessageInfo

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
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
	RemoteBackupAddr [][]byte `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
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

func (m *Ipv4RibEdmPathItem) Reset()         { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()    {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb, []int{3}
}
func (m *Ipv4RibEdmPathItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPathItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmPathItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPathItem.Merge(dst, src)
}
func (m *Ipv4RibEdmPathItem) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Size(m)
}
func (m *Ipv4RibEdmPathItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmPathItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmPathItem proto.InternalMessageInfo

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
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
	return fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb, []int{4}
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.non_as.protocol_routes.protocol_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.non_as.protocol_routes.protocol_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.non_as.protocol_routes.protocol_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.non_as.protocol_routes.protocol_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.non_as.protocol_routes.protocol_route.rib_edm_local_label")
}

func init() {
	proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb)
}

var fileDescriptor_ipv4_rib_edm_route_5e2b2de96113fdeb = []byte{
	// 1497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x52, 0x1c, 0xbb,
	0x15, 0xae, 0x01, 0x03, 0x33, 0x82, 0xc1, 0x4c, 0x43, 0x40, 0x18, 0xff, 0x8c, 0x71, 0x5c, 0x35,
	0xa4, 0xc2, 0x38, 0x65, 0x3b, 0x24, 0x71, 0x7e, 0x31, 0xc6, 0xf6, 0xc4, 0xd8, 0x90, 0xb1, 0xcb,
	0x55, 0x59, 0xa9, 0x34, 0x2d, 0xf5, 0x8c, 0xca, 0xdd, 0xad, 0xb6, 0xa4, 0x19, 0xc3, 0x2e, 0x79,
	0x99, 0xbc, 0x43, 0x36, 0x79, 0x87, 0x2c, 0xef, 0x2b, 0xdc, 0xe5, 0x5d, 0xdd, 0xe5, 0x2d, 0x9d,
	0xa3, 0xee, 0x19, 0xc0, 0xf7, 0x01, 0xbc, 0x01, 0xce, 0xf7, 0x7d, 0x52, 0xeb, 0xfc, 0xe8, 0xe8,
	0x40, 0xa8, 0x2a, 0x26, 0x4f, 0x99, 0x51, 0x03, 0x26, 0x45, 0xc6, 0x8c, 0x1e, 0x3b, 0xd9, 0x2d,
	0x8c, 0x76, 0x3a, 0xfa, 0x4f, 0x2d, 0x56, 0x36, 0xd6, 0x4c, 0x69, 0xcb, 0xce, 0x0d, 0x53, 0x05,
	0xa8, 0x40, 0xae, 0x0b, 0x69, 0xba, 0xde, 0xb2, 0x4e, 0x0c, 0x2e, 0xba, 0x13, 0x93, 0x58, 0xff,
	0xa3, 0xcb, 0x13, 0xdb, 0xe5, 0x49, 0xd7, 0xfa, 0xdf, 0x96, 0x27, 0xdd, 0xb0, 0x06, 0x76, 0x65,
	0x8e, 0x0f, 0x52, 0xc9, 0x72, 0x9e, 0x49, 0xfb, 0x73, 0x04, 0x7e, 0x39, 0xd6, 0x69, 0x37, 0xd5,
	0x31, 0x4f, 0xbb, 0xb9, 0xce, 0x19, 0xb7, 0x15, 0x8a, 0xfa, 0xab, 0xf6, 0xee, 0xff, 0x6b, 0x64,
	0xeb, 0xba, 0x17, 0xec, 0xcd, 0xf1, 0x3f, 0xdf, 0x47, 0xdb, 0xa4, 0x3e, 0x31, 0x09, 0x6c, 0x4e,
	0x6b, 0xed, 0x5a, 0xa7, 0xd1, 0x5f, 0x9a, 0x98, 0xe4, 0x1d, 0xcf, 0x64, 0xb4, 0x45, 0x96, 0x78,
	0x60, 0xe6, 0x80, 0x59, 0xe4, 0x48, 0x6c, 0x93, 0xba, 0x2d, 0x99, 0x79, 0x5c, 0x63, 0x03, 0xd5,
	0x21, 0x6b, 0x57, 0xcf, 0x4c, 0x6f, 0x80, 0x64, 0x15, 0xf0, 0x0f, 0x1e, 0x06, 0x25, 0x25, 0x4b,
	0x5c, 0x08, 0x23, 0xad, 0xa5, 0x0b, 0xb8, 0x47, 0x30, 0xa3, 0x07, 0xa4, 0x59, 0x18, 0x99, 0xa8,
	0x73, 0x96, 0xca, 0x7c, 0xe8, 0x46, 0x74, 0xb1, 0x5d, 0xeb, 0x34, 0xfb, 0x2b, 0x08, 0x9e, 0x00,
	0xb6, 0xfb, 0x63, 0x83, 0x44, 0xd7, 0x7d, 0x8a, 0x36, 0xc9, 0x22, 0xca, 0xe8, 0x63, 0x3c, 0x32,
	0x5a, 0xd7, 0xf7, 0x7c, 0x72, 0x7d, 0x4f, 0x2f, 0xc2, 0xc3, 0x4f, 0xa4, 0xb1, 0x4a, 0xe7, 0xf4,
	0x29, 0x8a, 0x00, 0xfc, 0x88, 0x58, 0x74, 0x8f, 0x2c, 0x57, 0xe1, 0x55, 0x82, 0xfe, 0x16, 0x24,
	0xa4, 0x84, 0x7a, 0x02, 0x3f, 0x15, 0x04, 0xe0, 0xff, 0x01, 0x9c, 0x64, 0xa5, 0x04, 0xc1, 0xfb,
	0x5b, 0xa4, 0xae, 0x72, 0xeb, 0x78, 0x1e, 0x4b, 0xfa, 0x3b, 0xe0, 0x2b, 0x3b, 0xda, 0x21, 0x8d,
	0x38, 0x55, 0x32, 0x77, 0x7e, 0xff, 0xdf, 0xc3, 0xfe, 0x75, 0x04, 0x7a, 0x22, 0xba, 0x43, 0x48,
	0x08, 0xf0, 0x45, 0x21, 0xe9, 0x1f, 0x80, 0x6d, 0x60, 0x68, 0x2f, 0x0a, 0xd8, 0xb7, 0x30, 0x4a,
	0x1b, 0xe5, 0x2e, 0xe8, 0x33, 0x5c, 0x5a, 0xda, 0x90, 0xb6, 0x89, 0xc0, 0x85, 0x7f, 0x04, 0x6e,
	0xc9, 0x4e, 0x04, 0x2c, 0xdb, 0x20, 0x0b, 0x49, 0xca, 0x87, 0x96, 0xfe, 0x09, 0x70, 0x34, 0xa2,
	0x87, 0x64, 0x55, 0x9e, 0x3b, 0x99, 0x0b, 0x29, 0x18, 0xd2, 0x7f, 0x6e, 0xd7, 0x3a, 0x37, 0xfa,
	0xcd, 0x12, 0x7d, 0x09, 0xb2, 0x35, 0x32, 0xef, 0xf8, 0x90, 0xfe, 0x05, 0x96, 0xfa, 0x3f, 0xfd,
	0x29, 0x84, 0x0a, 0xde, 0xfd, 0x15, 0x4f, 0x51, 0xda, 0xd1, 0x3e, 0x89, 0x84, 0x0a, 0x01, 0x66,
	0x95, 0xea, 0x6f, 0xa0, 0x6a, 0x55, 0xcc, 0x8b, 0x52, 0xbe, 0x49, 0x16, 0x33, 0xe9, 0x8c, 0x8a,
	0xe9, 0x21, 0x48, 0x82, 0x05, 0x69, 0xe0, 0x6e, 0x64, 0x59, 0xac, 0xc7, 0xb9, 0xa3, 0xcf, 0x43,
	0x1a, 0x3c, 0x74, 0xe4, 0x11, 0xff, 0x1d, 0xee, 0x9c, 0x51, 0x03, 0x1f, 0x2c, 0x25, 0x64, 0xee,
	0x7c, 0x4c, 0x8e, 0xf0, 0x3b, 0x15, 0xd3, 0x0b, 0x84, 0xcf, 0x9a, 0x33, 0x3c, 0x49, 0x54, 0xcc,
	0x54, 0x2e, 0xe4, 0x39, 0x7d, 0x81, 0xb9, 0x0f, 0x60, 0xcf, 0x63, 0xd1, 0x5e, 0x59, 0xdd, 0x85,
	0x91, 0xb1, 0x14, 0xd2, 0x9f, 0xfc, 0x18, 0x74, 0x37, 0x01, 0x3f, 0xab, 0x60, 0x9f, 0xc4, 0xcf,
	0xda, 0xb2, 0xa1, 0xd1, 0xe3, 0x82, 0xbe, 0xc4, 0x18, 0x7c, 0xd6, 0xf6, 0x95, 0xb7, 0x7d, 0x26,
	0x92, 0x54, 0x7f, 0x61, 0x3e, 0x6c, 0xaf, 0x30, 0x13, 0xde, 0xfe, 0xc0, 0x87, 0x7e, 0x5d, 0xf2,
	0x45, 0xb0, 0x38, 0xe5, 0xd6, 0xd2, 0xd7, 0xb8, 0x2e, 0xf9, 0x22, 0x8e, 0xbc, 0xed, 0xc9, 0x42,
	0xc5, 0xc1, 0xe5, 0x5e, 0x48, 0xaf, 0x8a, 0xd1, 0xe1, 0x4d, 0xb2, 0xc8, 0x63, 0xa7, 0x26, 0x92,
	0xfe, 0xbd, 0x5d, 0xeb, 0xd4, 0xfb, 0xc1, 0x8a, 0x6e, 0x93, 0x46, 0x15, 0x56, 0xfa, 0x06, 0xa8,
	0x29, 0x10, 0xfd, 0x86, 0x6c, 0x4c, 0xd3, 0x01, 0x25, 0x8a, 0x45, 0x7b, 0x02, 0x45, 0x39, 0x4d,
	0xd5, 0x99, 0xa7, 0xa0, 0x74, 0x77, 0x08, 0xd6, 0x1b, 0xe3, 0x43, 0x49, 0xdf, 0xe2, 0x21, 0x00,
	0x38, 0x1c, 0x4a, 0x9f, 0x16, 0x24, 0x53, 0x3e, 0x90, 0x29, 0x7d, 0x87, 0x69, 0x01, 0xe8, 0xc4,
	0x23, 0xfe, 0xda, 0x97, 0x67, 0x39, 0x45, 0xcf, 0x27, 0xd3, 0x8b, 0xe5, 0x06, 0x69, 0x75, 0xf7,
	0xce, 0xa0, 0xd4, 0x88, 0x1b, 0xa4, 0xe5, 0xcd, 0xfb, 0x15, 0x69, 0xe1, 0xde, 0x99, 0x16, 0x2a,
	0xb9, 0x60, 0x4e, 0x65, 0x92, 0xfe, 0x03, 0x64, 0x18, 0xfe, 0xb7, 0x80, 0x7f, 0x50, 0x99, 0x8c,
	0xbe, 0xab, 0x95, 0xf7, 0xc4, 0x97, 0x04, 0xed, 0xb7, 0x6b, 0x9d, 0xe5, 0xc7, 0xff, 0xad, 0x75,
	0xbf, 0x8d, 0x8e, 0xdd, 0xbd, 0xd4, 0xd9, 0xbc, 0x07, 0xe1, 0x8e, 0x9f, 0x71, 0x37, 0xda, 0xfd,
	0xd7, 0x1c, 0x69, 0x5d, 0x13, 0x44, 0xdf, 0xd7, 0xbe, 0x82, 0xd2, 0x5a, 0x7b, 0xbe, 0xb3, 0xfc,
	0xf8, 0x7f, 0xdf, 0xae, 0xe3, 0x4c, 0x39, 0x99, 0xf5, 0x57, 0x3d, 0xde, 0x57, 0x83, 0x63, 0x91,
	0x41, 0x08, 0x7e, 0x20, 0x64, 0xf3, 0xeb, 0xd2, 0xd9, 0x77, 0xa5, 0x76, 0xf9, 0x5d, 0xd9, 0x27,
	0x91, 0xca, 0x13, 0x6d, 0x32, 0xee, 0x7c, 0xb1, 0x5b, 0x3d, 0x36, 0x71, 0xf9, 0xb4, 0xb5, 0x66,
	0x98, 0xf7, 0x40, 0xf8, 0x4e, 0x3b, 0x39, 0x60, 0xb9, 0x3c, 0x77, 0x23, 0x5d, 0x84, 0x77, 0xae,
	0x31, 0x39, 0x78, 0x87, 0x80, 0x6f, 0x8e, 0x2a, 0x77, 0xd2, 0x24, 0x3c, 0xbe, 0xf4, 0xce, 0x35,
	0x2b, 0x14, 0x6e, 0xcb, 0xb4, 0x7f, 0x2d, 0x5c, 0xed, 0x5f, 0xa9, 0xe6, 0x82, 0x05, 0x12, 0x9f,
	0x38, 0xe2, 0xa1, 0xb7, 0x28, 0xa0, 0x64, 0x09, 0x7a, 0xee, 0xc1, 0x53, 0xba, 0x04, 0x35, 0x5e,
	0x9a, 0xd3, 0x66, 0x5d, 0x9f, 0x6d, 0xd6, 0xf0, 0xec, 0xa8, 0x09, 0x77, 0x32, 0xf4, 0xea, 0x46,
	0xf9, 0xc2, 0x01, 0x88, 0xad, 0x7a, 0x93, 0x2c, 0xa6, 0x5a, 0x17, 0x52, 0x50, 0x82, 0x3d, 0x02,
	0xad, 0x68, 0x8f, 0xb4, 0xbc, 0xa3, 0x6c, 0xa4, 0x8b, 0x90, 0x42, 0x25, 0xe8, 0x32, 0x6c, 0xb0,
	0xea, 0x89, 0xd7, 0xba, 0x80, 0x97, 0xbb, 0x77, 0x59, 0x5a, 0x4d, 0x0e, 0x2b, 0xf8, 0xc4, 0x07,
	0xe9, 0xc7, 0x30, 0x40, 0xec, 0x93, 0xf5, 0x2b, 0xbb, 0x82, 0xb8, 0x09, 0xe2, 0xb5, 0xd9, 0x7d,
	0x41, 0xde, 0x26, 0x2b, 0x95, 0x9c, 0x27, 0x8a, 0xae, 0x62, 0x4c, 0x82, 0xee, 0x30, 0x51, 0xd1,
	0x2e, 0x69, 0x56, 0x0a, 0xeb, 0x25, 0x37, 0x41, 0xb2, 0x1c, 0x24, 0xef, 0x79, 0xa2, 0xae, 0x76,
	0xa0, 0xb5, 0x6b, 0x1d, 0x68, 0x87, 0x34, 0xdc, 0x38, 0xcf, 0x25, 0x3c, 0xdf, 0x2d, 0xec, 0x5f,
	0x08, 0xf4, 0x04, 0xcc, 0x0f, 0xdc, 0x8d, 0x94, 0xa0, 0x11, 0xa6, 0x0b, 0x2d, 0x1f, 0xdd, 0x01,
	0x8f, 0x3f, 0x8d, 0x0b, 0x16, 0xe8, 0x75, 0x8c, 0x2e, 0x82, 0x67, 0x28, 0xda, 0x23, 0x2d, 0x23,
	0x13, 0x16, 0xe7, 0x8e, 0xe9, 0x84, 0x21, 0x45, 0x37, 0x30, 0x8a, 0x46, 0x26, 0x47, 0xb9, 0x3b,
	0x4d, 0x9e, 0x03, 0x1a, 0x1d, 0x91, 0xbb, 0xf9, 0x38, 0x1b, 0x48, 0xe3, 0x95, 0xd5, 0x23, 0x1b,
	0xeb, 0x2c, 0x1b, 0xe7, 0xca, 0x29, 0x69, 0xe9, 0x2f, 0x60, 0xdd, 0x0e, 0xaa, 0x4e, 0x93, 0xe3,
	0xa0, 0x39, 0x9a, 0x4a, 0xa2, 0xfb, 0x64, 0x25, 0x9b, 0x14, 0xbe, 0x6d, 0x4b, 0x2b, 0x73, 0x47,
	0x37, 0x21, 0xa7, 0xcb, 0x1e, 0x3b, 0x43, 0xc8, 0x57, 0xa9, 0x3f, 0xb0, 0x71, 0x95, 0x68, 0x0b,
	0x44, 0x4d, 0x44, 0x4b, 0xd9, 0x23, 0xb2, 0x3e, 0x31, 0x89, 0xca, 0x0a, 0x6d, 0xdc, 0x8c, 0x96,
	0x82, 0x36, 0x9a, 0xa1, 0xca, 0x05, 0xfb, 0x24, 0xc2, 0xfb, 0xc3, 0xed, 0x8c, 0x7e, 0x1b, 0xf4,
	0xad, 0x29, 0x53, 0xca, 0xf7, 0xc8, 0x1a, 0x82, 0x46, 0x54, 0xe2, 0x5b, 0x20, 0xbe, 0x59, 0xe2,
	0xa5, 0xf4, 0x19, 0xd9, 0xb6, 0x72, 0x98, 0xc9, 0xdc, 0x49, 0x51, 0xde, 0xbe, 0x6a, 0xcd, 0x0e,
	0xac, 0xd9, 0xaa, 0x04, 0xe1, 0x32, 0x96, 0x6b, 0xef, 0x92, 0xe5, 0xaa, 0x3e, 0x94, 0xa0, 0xb7,
	0x71, 0x3a, 0x0a, 0xd5, 0xd1, 0x13, 0xd1, 0x23, 0xb2, 0x31, 0xc3, 0x33, 0x23, 0x13, 0x7c, 0x4a,
	0xef, 0xe0, 0x54, 0x50, 0x09, 0xfb, 0x81, 0xf0, 0x25, 0xa9, 0x6d, 0x91, 0x30, 0x6e, 0x24, 0xf7,
	0x3b, 0xde, 0x85, 0xd2, 0x25, 0x1e, 0x3b, 0x34, 0x92, 0xf7, 0x44, 0xf4, 0x6b, 0x12, 0x19, 0x99,
	0x69, 0x27, 0x43, 0xbe, 0x99, 0xef, 0x36, 0xf4, 0x5e, 0x7b, 0xbe, 0xb3, 0xd2, 0x5f, 0x43, 0x06,
	0x53, 0x7e, 0x28, 0x84, 0xf1, 0x19, 0x1b, 0x71, 0x8b, 0xa5, 0x69, 0xdd, 0x27, 0xda, 0xc6, 0x8c,
	0x8d, 0xb8, 0x3d, 0x09, 0x90, 0x6f, 0x3b, 0xf9, 0x38, 0x0b, 0x12, 0x7a, 0x3f, 0xb8, 0x30, 0xce,
	0x50, 0xe0, 0x47, 0xab, 0x6a, 0xf5, 0x6e, 0x7b, 0xde, 0x17, 0x6f, 0x69, 0x43, 0x91, 0xaa, 0x5c,
	0xa8, 0x7c, 0x18, 0x8a, 0xff, 0x41, 0x28, 0x52, 0x04, 0xab, 0xf2, 0xcf, 0x47, 0x4a, 0xb0, 0x44,
	0x2a, 0x41, 0x7f, 0x09, 0x9d, 0xa5, 0xee, 0x81, 0x97, 0x52, 0x09, 0x4f, 0x66, 0x45, 0x6a, 0x91,
	0x7c, 0x88, 0xa4, 0x07, 0x3c, 0xb9, 0xfb, 0xef, 0x39, 0xb2, 0x5e, 0xf6, 0x5b, 0xe8, 0xe9, 0xf8,
	0x95, 0xeb, 0x03, 0x6f, 0xed, 0x2b, 0x03, 0xef, 0xa5, 0xa1, 0x76, 0xee, 0xca, 0x50, 0xbb, 0x41,
	0x16, 0xac, 0xe3, 0x29, 0xfe, 0x37, 0xd1, 0xec, 0xa3, 0xe1, 0x5d, 0xcd, 0x94, 0x31, 0xda, 0x48,
	0x01, 0xbd, 0xb5, 0xd9, 0xaf, 0x6c, 0xff, 0xcd, 0x4c, 0x9a, 0xa1, 0xf4, 0x13, 0xa4, 0x6f, 0x20,
	0xa1, 0xbb, 0xae, 0x00, 0xf8, 0x02, 0x31, 0x68, 0x05, 0xd2, 0x8f, 0x98, 0x4c, 0xe7, 0xe9, 0x45,
	0xd9, 0x63, 0x11, 0x3a, 0xcd, 0xd3, 0x0b, 0xff, 0x5d, 0x0c, 0xd4, 0x12, 0x7e, 0x17, 0xfd, 0x99,
	0x9d, 0x5e, 0xeb, 0x97, 0xa7, 0xd7, 0xc1, 0x22, 0x38, 0xf5, 0xe4, 0xa7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x74, 0x9f, 0x8b, 0x16, 0x0e, 0x00, 0x00,
}
