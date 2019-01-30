// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_backup_routes_backup_route

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
	Protoid              uint32   `protobuf:"varint,7,opt,name=protoid,proto3" json:"protoid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()         { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1, []int{0}
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

func (m *Ipv4RibEdmRoute_KEYS) GetProtoid() uint32 {
	if m != nil {
		return m.Protoid
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
	return fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1, []int{1}
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
	return fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1, []int{2}
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
	return fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1, []int{3}
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
	return fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1, []int{4}
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.rib_edm_local_label")
}

func init() {
	proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1)
}

var fileDescriptor_ipv4_rib_edm_route_2c44c11eef363bc1 = []byte{
	// 1497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcd, 0x72, 0x53, 0x37,
	0x14, 0x1e, 0x13, 0x92, 0xd8, 0x4a, 0x1c, 0xe2, 0x9b, 0x34, 0x51, 0x08, 0x3f, 0x26, 0x94, 0x19,
	0xa7, 0xd3, 0x98, 0x0e, 0xd0, 0xb4, 0xa5, 0xbf, 0x21, 0x04, 0x70, 0x09, 0x24, 0x35, 0x0c, 0x33,
	0x5d, 0x69, 0xe4, 0x2b, 0x5d, 0x5b, 0xc3, 0xfd, 0x43, 0x92, 0x4d, 0xb2, 0xec, 0xbe, 0xef, 0xd0,
	0x4d, 0xdf, 0xa3, 0xaf, 0xd0, 0x27, 0xe8, 0xb2, 0x9b, 0xee, 0xfa, 0x04, 0x1d, 0x9d, 0xa3, 0x7b,
	0x6d, 0x27, 0xe9, 0x9e, 0x0d, 0xe4, 0x7c, 0xdf, 0x77, 0x74, 0xa5, 0x73, 0x8e, 0xce, 0x91, 0x09,
	0x55, 0xf9, 0xe8, 0x01, 0xd3, 0xaa, 0xc7, 0xa4, 0x48, 0x98, 0xce, 0x86, 0x56, 0xb6, 0x73, 0x9d,
	0xd9, 0x2c, 0xf8, 0xb5, 0x12, 0x2a, 0x13, 0x66, 0x4c, 0x65, 0x86, 0x9d, 0x68, 0xa6, 0x72, 0x50,
	0x81, 0x3c, 0xcb, 0xa5, 0x6e, 0x3b, 0xcb, 0x58, 0xd1, 0x3b, 0x6d, 0x8f, 0x74, 0x64, 0xdc, 0x3f,
	0x6d, 0x1e, 0x99, 0x36, 0x8f, 0xda, 0xc6, 0xfd, 0x6f, 0x78, 0xd4, 0xf6, 0x3e, 0xb0, 0x2a, 0xb3,
	0xbc, 0x17, 0x4b, 0x96, 0xf2, 0x44, 0x9a, 0xff, 0x23, 0xda, 0x3d, 0x1e, 0xbe, 0x1d, 0xe6, 0x88,
	0x9b, 0x29, 0x6b, 0xeb, 0x9f, 0x0a, 0x59, 0x3f, 0xbf, 0x57, 0xf6, 0xfc, 0xe0, 0xe7, 0x57, 0xc1,
	0x06, 0xa9, 0x8e, 0x74, 0x04, 0x4b, 0xd0, 0x4a, 0xb3, 0xd2, 0xaa, 0x75, 0xe7, 0x47, 0x3a, 0x7a,
	0xc9, 0x13, 0x19, 0xac, 0x93, 0x79, 0xee, 0x99, 0x4b, 0xc0, 0xcc, 0x71, 0x24, 0x36, 0x48, 0xd5,
	0x14, 0xcc, 0x0c, 0xfa, 0x18, 0x4f, 0xb5, 0xc8, 0xf2, 0xd9, 0x9d, 0xd1, 0xcb, 0x20, 0x59, 0x02,
	0xfc, 0xb5, 0x83, 0x41, 0x49, 0xc9, 0x3c, 0x17, 0x42, 0x4b, 0x63, 0xe8, 0x2c, 0xae, 0xe1, 0xcd,
	0xe0, 0x36, 0xa9, 0xe7, 0x5a, 0x46, 0xea, 0x84, 0xc5, 0x32, 0xed, 0xdb, 0x01, 0x9d, 0x6b, 0x56,
	0x5a, 0xf5, 0xee, 0x22, 0x82, 0x87, 0x80, 0x39, 0x77, 0x88, 0xb5, 0x12, 0x74, 0x1e, 0xe8, 0xc2,
	0xdc, 0xfa, 0xab, 0x46, 0x82, 0xf3, 0xa7, 0x0d, 0xd6, 0xc8, 0x1c, 0x2e, 0x40, 0xef, 0xe1, 0x61,
	0xd0, 0x3a, 0xff, 0xb5, 0xfb, 0x17, 0x7c, 0xed, 0x36, 0xa9, 0xe3, 0xb1, 0x46, 0x52, 0x1b, 0x95,
	0xa5, 0xf4, 0x01, 0x8a, 0x00, 0x7c, 0x83, 0x58, 0x70, 0x93, 0x2c, 0xc0, 0x1e, 0xc2, 0x2c, 0x66,
	0x4a, 0xd0, 0xcf, 0x41, 0x42, 0x0a, 0xa8, 0x23, 0xf0, 0x53, 0x5e, 0x00, 0x91, 0xd9, 0x85, 0x9d,
	0x2c, 0x16, 0x20, 0xc4, 0xe5, 0x2a, 0xa9, 0xaa, 0xd4, 0x58, 0x9e, 0x86, 0x92, 0x7e, 0x01, 0x7c,
	0x69, 0x07, 0x9b, 0xa4, 0x16, 0xc6, 0x4a, 0xa6, 0xd6, 0xad, 0xff, 0x25, 0xac, 0x5f, 0x45, 0xa0,
	0x23, 0x82, 0xeb, 0x84, 0xf8, 0xd0, 0x9f, 0xe6, 0x92, 0x7e, 0x05, 0x6c, 0x0d, 0x83, 0x7e, 0x9a,
	0xc3, 0xba, 0xb9, 0x56, 0x99, 0x56, 0xf6, 0x94, 0x3e, 0x44, 0xd7, 0xc2, 0x86, 0x84, 0x8e, 0x04,
	0x3a, 0x7e, 0x8d, 0xd1, 0x34, 0x23, 0x01, 0x6e, 0xab, 0x64, 0x36, 0x8a, 0x79, 0xdf, 0xd0, 0x6f,
	0x00, 0x47, 0x23, 0xb8, 0x43, 0x96, 0xe4, 0x89, 0x95, 0xa9, 0x90, 0x82, 0x21, 0xfd, 0x6d, 0xb3,
	0xd2, 0xba, 0xdc, 0xad, 0x17, 0xe8, 0x13, 0x90, 0x2d, 0x93, 0x19, 0xcb, 0xfb, 0xf4, 0x3b, 0x70,
	0x75, 0x7f, 0xba, 0x5d, 0x08, 0xe5, 0x4f, 0xf7, 0x3d, 0xee, 0xa2, 0xb0, 0x83, 0x1d, 0x12, 0x08,
	0xe5, 0x03, 0xcc, 0x4a, 0xd5, 0x0f, 0xa0, 0x6a, 0x94, 0xcc, 0xe3, 0x42, 0xbe, 0x46, 0xe6, 0x12,
	0x69, 0xb5, 0x0a, 0xe9, 0x1e, 0x48, 0xbc, 0x05, 0x69, 0xe0, 0x76, 0x60, 0x58, 0x98, 0x0d, 0x53,
	0x4b, 0x1f, 0xf9, 0x34, 0x38, 0x68, 0xdf, 0x21, 0xee, 0x3b, 0xdc, 0x5a, 0xad, 0x7a, 0x2e, 0x58,
	0x4a, 0xc8, 0xd4, 0xba, 0x98, 0xec, 0xe3, 0x77, 0x4a, 0xa6, 0xe3, 0x09, 0x97, 0x35, 0xab, 0x79,
	0x14, 0xa9, 0x90, 0xa9, 0x54, 0xc8, 0x13, 0xfa, 0x18, 0x73, 0xef, 0xc1, 0x8e, 0xc3, 0x82, 0xed,
	0xa2, 0xee, 0x73, 0x2d, 0x43, 0x29, 0xa4, 0xdb, 0xf9, 0x01, 0xe8, 0xae, 0x00, 0x7e, 0x5c, 0xc2,
	0x2e, 0x89, 0xef, 0x32, 0xc3, 0xfa, 0x3a, 0x1b, 0xe6, 0xf4, 0x09, 0xc6, 0xe0, 0x5d, 0x66, 0x9e,
	0x3a, 0xdb, 0x65, 0x22, 0x8a, 0xb3, 0xf7, 0xcc, 0x85, 0xed, 0x29, 0x66, 0xc2, 0xd9, 0xaf, 0x79,
	0xdf, 0xf9, 0x45, 0xef, 0x05, 0x0b, 0x63, 0x6e, 0x0c, 0x7d, 0x86, 0x7e, 0xd1, 0x7b, 0xb1, 0xef,
	0x6c, 0x47, 0xe6, 0x2a, 0xf4, 0x47, 0xee, 0xf8, 0xf4, 0xaa, 0x10, 0x0f, 0xbc, 0x46, 0xe6, 0x78,
	0x68, 0xd5, 0x48, 0xd2, 0x1f, 0x9b, 0x95, 0x56, 0xb5, 0xeb, 0xad, 0xe0, 0x1a, 0xa9, 0x95, 0x61,
	0xa5, 0xcf, 0x81, 0x1a, 0x03, 0xc1, 0x67, 0x64, 0x75, 0x9c, 0x0e, 0x28, 0x51, 0x2c, 0xda, 0x43,
	0x28, 0xca, 0x71, 0xaa, 0x8e, 0x1d, 0x05, 0xa5, 0xbb, 0x49, 0xb0, 0xde, 0x18, 0xef, 0x4b, 0xfa,
	0x02, 0x37, 0x01, 0xc0, 0x5e, 0x5f, 0xba, 0xb4, 0x20, 0x19, 0xf3, 0x9e, 0x8c, 0xe9, 0x4b, 0x4c,
	0x0b, 0x40, 0x87, 0x0e, 0x71, 0x37, 0xba, 0xd8, 0xcb, 0x11, 0x9e, 0x7c, 0x34, 0xbe, 0x58, 0xb6,
	0x17, 0x97, 0x77, 0xef, 0x18, 0x4a, 0x8d, 0xd8, 0x5e, 0x5c, 0xdc, 0xbc, 0x4f, 0x48, 0x03, 0xd7,
	0x4e, 0x32, 0xa1, 0xa2, 0x53, 0x66, 0x55, 0x22, 0xe9, 0x4f, 0x20, 0xc3, 0xf0, 0xbf, 0x00, 0xfc,
	0xb5, 0x4a, 0x64, 0xf0, 0x47, 0xa5, 0xb8, 0x27, 0xae, 0x24, 0x68, 0xb7, 0x59, 0x69, 0x2d, 0xdc,
	0xfb, 0xad, 0xd2, 0xfe, 0x90, 0x3a, 0x76, 0x7b, 0xaa, 0x7f, 0xb9, 0x7d, 0xfa, 0x9b, 0x7c, 0xcc,
	0xed, 0x60, 0xeb, 0xef, 0x0a, 0x69, 0x9c, 0x13, 0x04, 0x7f, 0x5e, 0x84, 0xd2, 0x4a, 0x73, 0xa6,
	0xb5, 0x70, 0xef, 0xf7, 0x0f, 0xfd, 0x78, 0x4c, 0x59, 0x99, 0x74, 0x97, 0x1c, 0xde, 0x55, 0xbd,
	0x03, 0x91, 0xc0, 0x41, 0xff, 0x25, 0x64, 0xed, 0x62, 0xe9, 0xe4, 0xf4, 0xa8, 0x4c, 0x4f, 0x8f,
	0x1d, 0x12, 0xa8, 0x34, 0xca, 0x74, 0xc2, 0xad, 0x2b, 0x5c, 0x93, 0x0d, 0x75, 0x58, 0x0c, 0xb0,
	0xc6, 0x04, 0xf3, 0x0a, 0x08, 0xd7, 0x35, 0x47, 0xbb, 0x2c, 0x95, 0x27, 0x76, 0x90, 0xe5, 0x7e,
	0x9a, 0xd5, 0x46, 0xbb, 0x2f, 0x11, 0x70, 0x8d, 0x4e, 0xa5, 0x56, 0xea, 0x88, 0x87, 0x53, 0xd3,
	0xac, 0x5e, 0xa2, 0x50, 0xf9, 0xe3, 0x5e, 0x34, 0x7b, 0xb6, 0x17, 0xc5, 0x19, 0x17, 0xcc, 0x93,
	0x38, 0xc8, 0x88, 0x83, 0x5e, 0xa0, 0x80, 0x92, 0x79, 0xe8, 0x9f, 0xbb, 0x0f, 0x60, 0x8c, 0x5d,
	0xee, 0x16, 0xe6, 0xb8, 0xf1, 0x56, 0x27, 0x1b, 0x2f, 0x8c, 0x10, 0x35, 0xe2, 0x56, 0xfa, 0xbe,
	0x5b, 0x2b, 0xa6, 0x15, 0x80, 0xd8, 0x76, 0xd7, 0xc8, 0x5c, 0x9c, 0x65, 0xb9, 0x14, 0x94, 0xe0,
	0x7d, 0x47, 0x2b, 0xd8, 0x26, 0x0d, 0x77, 0x50, 0x36, 0xc8, 0x72, 0x9f, 0x28, 0x25, 0xe8, 0x02,
	0x2c, 0xb0, 0xe4, 0x88, 0x67, 0x59, 0x0e, 0xf3, 0xb9, 0x33, 0x2d, 0x2d, 0xdf, 0x07, 0x8b, 0x38,
	0xc8, 0xbd, 0xf4, 0x8d, 0x7f, 0x26, 0xec, 0x90, 0x95, 0x33, 0xab, 0x82, 0xb8, 0x0e, 0xe2, 0xe5,
	0xc9, 0x75, 0x41, 0xde, 0x24, 0x8b, 0xa5, 0x9c, 0x47, 0x8a, 0x2e, 0x61, 0x4c, 0xbc, 0x6e, 0x2f,
	0x52, 0xc1, 0x16, 0xa9, 0x97, 0x0a, 0xe3, 0x24, 0x57, 0x40, 0xb2, 0xe0, 0x25, 0xaf, 0x78, 0xa4,
	0xce, 0x76, 0x93, 0xe5, 0x73, 0xdd, 0x64, 0x93, 0xd4, 0xec, 0x30, 0x4d, 0x25, 0x8c, 0xe2, 0x06,
	0xf6, 0x22, 0x04, 0x3a, 0x02, 0xde, 0x02, 0xdc, 0x0e, 0x94, 0xa0, 0x01, 0xa6, 0x0b, 0x2d, 0x17,
	0x5d, 0x5f, 0xa7, 0x9e, 0x5e, 0xc1, 0xe8, 0x22, 0x78, 0x8c, 0xa2, 0x6d, 0xd2, 0xd0, 0x32, 0x62,
	0x61, 0x6a, 0x59, 0x16, 0x31, 0xa4, 0xe8, 0x2a, 0x46, 0x51, 0xcb, 0x68, 0x3f, 0xb5, 0x47, 0xd1,
	0x23, 0x40, 0x83, 0x7d, 0x72, 0x23, 0x1d, 0x26, 0x3d, 0xa9, 0x9d, 0xb2, 0x1c, 0x98, 0x61, 0x96,
	0x24, 0xc3, 0x54, 0x59, 0x25, 0x0d, 0xfd, 0x08, 0xfc, 0x36, 0x51, 0x75, 0x14, 0x1d, 0x78, 0xcd,
	0xfe, 0x58, 0x12, 0xdc, 0x22, 0x8b, 0xc9, 0x28, 0x77, 0x2d, 0x58, 0x1a, 0x99, 0x5a, 0xba, 0x06,
	0x39, 0x5d, 0x70, 0xd8, 0x31, 0x42, 0xae, 0x4a, 0xdd, 0x86, 0xb5, 0x2d, 0x45, 0xeb, 0x20, 0xaa,
	0x23, 0x5a, 0xc8, 0xee, 0x92, 0x95, 0x91, 0x8e, 0x54, 0x92, 0x67, 0xda, 0x4e, 0x68, 0x29, 0x68,
	0x83, 0x09, 0xaa, 0x70, 0xd8, 0x21, 0x01, 0xde, 0x1f, 0x6e, 0x26, 0xf4, 0x1b, 0xa0, 0x6f, 0x8c,
	0x99, 0x42, 0xbe, 0x4d, 0x96, 0x11, 0xd4, 0xa2, 0x14, 0x5f, 0x05, 0xf1, 0x95, 0x02, 0x2f, 0xa4,
	0x0f, 0xc9, 0x86, 0x91, 0xfd, 0x44, 0xa6, 0x56, 0x8a, 0xe2, 0xf6, 0x95, 0x3e, 0x9b, 0xe0, 0xb3,
	0x5e, 0x0a, 0xfc, 0x65, 0x2c, 0x7c, 0x6f, 0x90, 0x85, 0xb2, 0x3e, 0x94, 0xa0, 0xd7, 0xf0, 0xa5,
	0xe3, 0xab, 0xa3, 0x23, 0x82, 0xbb, 0x64, 0x75, 0x82, 0x67, 0x5a, 0x46, 0x38, 0x16, 0xaf, 0xe3,
	0x84, 0x2f, 0x85, 0x5d, 0x4f, 0xb8, 0x92, 0xcc, 0x4c, 0x1e, 0x31, 0xae, 0x25, 0x77, 0x2b, 0xde,
	0x80, 0xd2, 0x25, 0x0e, 0xdb, 0xd3, 0x92, 0x77, 0x44, 0xf0, 0x29, 0x09, 0xb4, 0x4c, 0x32, 0x2b,
	0x7d, 0xbe, 0x99, 0xeb, 0x36, 0xf4, 0x66, 0x73, 0xa6, 0xb5, 0xd8, 0x5d, 0x46, 0x06, 0x53, 0xbe,
	0x27, 0x84, 0x76, 0x19, 0x1b, 0x70, 0x83, 0xa5, 0x69, 0xec, 0x5b, 0xda, 0xc4, 0x8c, 0x0d, 0xb8,
	0x39, 0xf4, 0x90, 0x6b, 0x3b, 0xe9, 0x30, 0xf1, 0x12, 0x7a, 0xcb, 0x1f, 0x61, 0x98, 0xa0, 0xc0,
	0x3d, 0x93, 0x4a, 0xef, 0xad, 0xe6, 0x8c, 0x2b, 0xde, 0xc2, 0x86, 0x22, 0x55, 0xa9, 0x50, 0x69,
	0xdf, 0x17, 0xff, 0x6d, 0x5f, 0xa4, 0x08, 0x96, 0xe5, 0x9f, 0x0e, 0x94, 0x60, 0x91, 0x54, 0x82,
	0x7e, 0x0c, 0x9d, 0xa5, 0xea, 0x80, 0x27, 0x52, 0x09, 0x47, 0x26, 0x79, 0x6c, 0x90, 0xbc, 0x83,
	0xa4, 0x03, 0x1c, 0xb9, 0xf5, 0xcb, 0x25, 0xb2, 0x52, 0xf4, 0xdb, 0x38, 0x0b, 0x79, 0x8c, 0x5f,
	0x39, 0xff, 0x78, 0xad, 0x5c, 0xf0, 0x78, 0x9d, 0x7a, 0xa0, 0x5e, 0x3a, 0xf3, 0x40, 0x5d, 0x25,
	0xb3, 0xc6, 0xf2, 0x18, 0x7f, 0x33, 0xd4, 0xbb, 0x68, 0xb8, 0xa3, 0x26, 0x4a, 0xeb, 0x4c, 0x4b,
	0x01, 0xbd, 0xb5, 0xde, 0x2d, 0x6d, 0xf7, 0xcd, 0x44, 0xea, 0xbe, 0x74, 0xaf, 0x41, 0xd7, 0x40,
	0x7c, 0x77, 0x5d, 0x04, 0xf0, 0x31, 0x62, 0xd0, 0x0a, 0xa4, 0x7b, 0x2e, 0xb2, 0x2c, 0x8d, 0x4f,
	0x8b, 0x1e, 0x8b, 0xd0, 0x51, 0x1a, 0x9f, 0xba, 0xef, 0x62, 0xa0, 0xf0, 0x87, 0x02, 0x1a, 0x53,
	0x2f, 0xd1, 0xea, 0xf4, 0x4b, 0xb4, 0x37, 0x07, 0x87, 0xba, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0xe4, 0xe8, 0x2a, 0xe2, 0x0d, 0x00, 0x00,
}
