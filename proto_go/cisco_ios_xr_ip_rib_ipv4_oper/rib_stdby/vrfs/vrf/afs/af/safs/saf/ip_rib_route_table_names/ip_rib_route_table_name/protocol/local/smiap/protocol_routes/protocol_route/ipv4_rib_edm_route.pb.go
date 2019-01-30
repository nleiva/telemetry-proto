// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_route.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_smiap_protocol_routes_protocol_route

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
	return fileDescriptor_44c1b2f42127243a, []int{0}
}

func (m *Ipv4RibEdmRoute_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmRoute_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute_KEYS.Merge(m, src)
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
	return fileDescriptor_44c1b2f42127243a, []int{1}
}

func (m *Ipv4RibEdmRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmRoute.Unmarshal(m, b)
}
func (m *Ipv4RibEdmRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmRoute.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmRoute.Merge(m, src)
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
	return fileDescriptor_44c1b2f42127243a, []int{2}
}

func (m *Ipv4RibEdmPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPath.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPath.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPath.Merge(m, src)
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
	return fileDescriptor_44c1b2f42127243a, []int{3}
}

func (m *Ipv4RibEdmPathItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Unmarshal(m, b)
}
func (m *Ipv4RibEdmPathItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmPathItem.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmPathItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmPathItem.Merge(m, src)
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
	return fileDescriptor_44c1b2f42127243a, []int{4}
}

func (m *RibEdmLocalLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmLocalLabel.Unmarshal(m, b)
}
func (m *RibEdmLocalLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmLocalLabel.Marshal(b, m, deterministic)
}
func (m *RibEdmLocalLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmLocalLabel.Merge(m, src)
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.smiap.protocol_routes.protocol_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.smiap.protocol_routes.protocol_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.smiap.protocol_routes.protocol_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.smiap.protocol_routes.protocol_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.smiap.protocol_routes.protocol_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor_44c1b2f42127243a) }

var fileDescriptor_44c1b2f42127243a = []byte{
	// 1496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x52, 0x1b, 0x39,
	0x16, 0x2e, 0x87, 0x00, 0xb6, 0xc0, 0x04, 0x37, 0x2c, 0x88, 0x90, 0x1f, 0x87, 0x6c, 0xaa, 0xcc,
	0xd6, 0xe2, 0x6c, 0x25, 0x59, 0x76, 0x37, 0xfb, 0x4b, 0x08, 0x49, 0xbc, 0x21, 0x81, 0x75, 0x52,
	0xa9, 0xda, 0x2b, 0x95, 0xdc, 0x52, 0xdb, 0xaa, 0x74, 0xb7, 0x3a, 0x92, 0xec, 0xc0, 0xe5, 0x3e,
	0xcc, 0xbe, 0xc3, 0xdc, 0xcc, 0x43, 0xcc, 0xdd, 0xbc, 0xc2, 0xdc, 0xce, 0xcd, 0x5c, 0xce, 0x94,
	0xce, 0x51, 0xb7, 0x0d, 0x64, 0x1e, 0x20, 0x37, 0xc0, 0xf9, 0xbe, 0x4f, 0x6a, 0x9d, 0x1f, 0x1d,
	0x1d, 0x08, 0x55, 0xc5, 0xe4, 0x09, 0x33, 0x6a, 0xc0, 0xa4, 0xc8, 0x98, 0xd1, 0x63, 0x27, 0xbb,
	0x85, 0xd1, 0x4e, 0x47, 0xff, 0xaf, 0xc5, 0xca, 0xc6, 0x9a, 0x29, 0x6d, 0xd9, 0x99, 0x61, 0xaa,
	0x00, 0x15, 0xc8, 0x75, 0x21, 0x4d, 0xd7, 0x5b, 0xd6, 0x89, 0xc1, 0x79, 0x77, 0x62, 0x12, 0xeb,
	0x7f, 0x74, 0x79, 0x62, 0xbb, 0x3c, 0xe9, 0x5a, 0xff, 0xdb, 0xf2, 0xa4, 0x1b, 0xd6, 0xc0, 0xae,
	0xcc, 0xf1, 0x41, 0x2a, 0x59, 0xce, 0x33, 0x69, 0x7f, 0x8d, 0xc0, 0x2f, 0xc7, 0x3a, 0xed, 0xa6,
	0x3a, 0xe6, 0x69, 0xd7, 0x66, 0x8a, 0x17, 0x15, 0x88, 0x72, 0x7b, 0xc9, 0xde, 0xf9, 0xae, 0x46,
	0x36, 0xaf, 0x3a, 0xc1, 0x5e, 0x1f, 0xfd, 0xf7, 0x5d, 0xb4, 0x45, 0xea, 0x13, 0x93, 0xc0, 0xde,
	0xb4, 0xd6, 0xae, 0x75, 0x1a, 0xfd, 0xc5, 0x89, 0x49, 0xde, 0xf2, 0x4c, 0x46, 0x9b, 0x64, 0x91,
	0x07, 0xe6, 0x1a, 0x30, 0x0b, 0x1c, 0x89, 0x2d, 0x52, 0xb7, 0x25, 0x33, 0x87, 0x6b, 0x6c, 0xa0,
	0x3a, 0x64, 0xf5, 0xf2, 0x91, 0xe9, 0x75, 0x90, 0xac, 0x00, 0xfe, 0xde, 0xc3, 0xa0, 0xa4, 0x64,
	0x91, 0x0b, 0x61, 0xa4, 0xb5, 0x74, 0x1e, 0xf7, 0x08, 0x66, 0x74, 0x9f, 0x34, 0x0b, 0x23, 0x13,
	0x75, 0xc6, 0x52, 0x99, 0x0f, 0xdd, 0x88, 0x2e, 0xb4, 0x6b, 0x9d, 0x66, 0x7f, 0x19, 0xc1, 0x63,
	0xc0, 0x76, 0x7e, 0x6a, 0x90, 0xe8, 0xaa, 0x4f, 0xd1, 0x06, 0x59, 0x40, 0x19, 0x7d, 0x84, 0x47,
	0x46, 0xeb, 0xea, 0x9e, 0x8f, 0xaf, 0xee, 0xe9, 0x45, 0x78, 0xf8, 0x89, 0x34, 0x56, 0xe9, 0x9c,
	0x3e, 0x41, 0x11, 0x80, 0x1f, 0x10, 0x8b, 0xee, 0x92, 0xa5, 0x2a, 0xbc, 0x4a, 0xd0, 0x3f, 0x82,
	0x84, 0x94, 0x50, 0x4f, 0xe0, 0xa7, 0x82, 0x00, 0xfc, 0xdf, 0x87, 0x93, 0x2c, 0x97, 0x20, 0x78,
	0x7f, 0x93, 0xd4, 0x55, 0x6e, 0x1d, 0xcf, 0x63, 0x49, 0xff, 0x04, 0x7c, 0x65, 0x47, 0xdb, 0xa4,
	0x11, 0xa7, 0x4a, 0xe6, 0xce, 0xef, 0xff, 0x67, 0xd8, 0xbf, 0x8e, 0x40, 0x4f, 0x44, 0xb7, 0x09,
	0x09, 0x01, 0x3e, 0x2f, 0x24, 0xfd, 0x0b, 0xb0, 0x0d, 0x0c, 0xed, 0x79, 0x01, 0xfb, 0x16, 0x46,
	0x69, 0xa3, 0xdc, 0x39, 0x7d, 0x8a, 0x4b, 0x4b, 0x1b, 0xd2, 0x36, 0x11, 0xb8, 0xf0, 0xaf, 0xc0,
	0x2d, 0xda, 0x89, 0x80, 0x65, 0xeb, 0x64, 0x3e, 0x49, 0xf9, 0xd0, 0xd2, 0xbf, 0x01, 0x8e, 0x46,
	0xf4, 0x80, 0xac, 0xc8, 0x33, 0x27, 0x73, 0x21, 0x05, 0x43, 0xfa, 0xef, 0xed, 0x5a, 0xe7, 0x7a,
	0xbf, 0x59, 0xa2, 0x2f, 0x40, 0xb6, 0x4a, 0xe6, 0x1c, 0x1f, 0xd2, 0x7f, 0xc0, 0x52, 0xff, 0xa7,
	0x3f, 0x85, 0x50, 0xc1, 0xbb, 0x7f, 0xe2, 0x29, 0x4a, 0x3b, 0xda, 0x23, 0x91, 0x50, 0x21, 0xc0,
	0xac, 0x52, 0xfd, 0x0b, 0x54, 0xad, 0x8a, 0x79, 0x5e, 0xca, 0x37, 0xc8, 0x42, 0x26, 0x9d, 0x51,
	0x31, 0x3d, 0x00, 0x49, 0xb0, 0x20, 0x0d, 0xdc, 0x8d, 0x2c, 0x8b, 0xf5, 0x38, 0x77, 0xf4, 0x59,
	0x48, 0x83, 0x87, 0x0e, 0x3d, 0xe2, 0xbf, 0xc3, 0x9d, 0x33, 0x6a, 0xe0, 0x83, 0xa5, 0x84, 0xcc,
	0x9d, 0x8f, 0xc9, 0x21, 0x7e, 0xa7, 0x62, 0x7a, 0x81, 0xf0, 0x59, 0x73, 0x86, 0x27, 0x89, 0x8a,
	0x99, 0xca, 0x85, 0x3c, 0xa3, 0xcf, 0x31, 0xf7, 0x01, 0xec, 0x79, 0x2c, 0xda, 0x2d, 0xab, 0xbb,
	0x30, 0x32, 0x96, 0x42, 0xfa, 0x93, 0x1f, 0x81, 0xee, 0x06, 0xe0, 0xa7, 0x15, 0xec, 0x93, 0xf8,
	0x49, 0x5b, 0x36, 0x34, 0x7a, 0x5c, 0xd0, 0x17, 0x18, 0x83, 0x4f, 0xda, 0xbe, 0xf4, 0xb6, 0xcf,
	0x44, 0x92, 0xea, 0xcf, 0xcc, 0x87, 0xed, 0x25, 0x66, 0xc2, 0xdb, 0xef, 0xf9, 0xd0, 0xaf, 0x4b,
	0x3e, 0x0b, 0x16, 0xa7, 0xdc, 0x5a, 0xfa, 0x0a, 0xd7, 0x25, 0x9f, 0xc5, 0xa1, 0xb7, 0x3d, 0x59,
	0xa8, 0x38, 0xb8, 0xdc, 0x0b, 0xe9, 0x55, 0x31, 0x3a, 0xbc, 0x41, 0x16, 0x78, 0xec, 0xd4, 0x44,
	0xd2, 0x7f, 0xb7, 0x6b, 0x9d, 0x7a, 0x3f, 0x58, 0xd1, 0x2d, 0xd2, 0xa8, 0xc2, 0x4a, 0x5f, 0x03,
	0x35, 0x05, 0xa2, 0x3f, 0x90, 0xf5, 0x69, 0x3a, 0xa0, 0x44, 0xb1, 0x68, 0x8f, 0xa1, 0x28, 0xa7,
	0xa9, 0x3a, 0xf5, 0x14, 0x94, 0xee, 0x36, 0xc1, 0x7a, 0x63, 0x7c, 0x28, 0xe9, 0x1b, 0x3c, 0x04,
	0x00, 0x07, 0x43, 0xe9, 0xd3, 0x82, 0x64, 0xca, 0x07, 0x32, 0xa5, 0x6f, 0x31, 0x2d, 0x00, 0x1d,
	0x7b, 0xc4, 0x5f, 0xfb, 0xf2, 0x2c, 0x27, 0xe8, 0xf9, 0x64, 0x7a, 0xb1, 0xdc, 0x20, 0xad, 0xee,
	0xde, 0x29, 0x94, 0x1a, 0x71, 0x83, 0xb4, 0xbc, 0x79, 0xbf, 0x23, 0x2d, 0xdc, 0x3b, 0xd3, 0x42,
	0x25, 0xe7, 0xcc, 0xa9, 0x4c, 0xd2, 0xff, 0x80, 0x0c, 0xc3, 0xff, 0x06, 0xf0, 0xf7, 0x2a, 0x93,
	0xd1, 0xf7, 0xb5, 0xf2, 0x9e, 0xf8, 0x92, 0xa0, 0xfd, 0x76, 0xad, 0xb3, 0xf4, 0xe8, 0x9b, 0x5a,
	0xf7, 0xab, 0x68, 0xd8, 0xdd, 0x0b, 0x8d, 0xcd, 0x3b, 0x10, 0xae, 0xf8, 0x29, 0x77, 0xa3, 0x9d,
	0x9f, 0x6b, 0xa4, 0x75, 0x45, 0x10, 0xfd, 0xf0, 0x25, 0x94, 0xd6, 0xda, 0x73, 0x9d, 0xa5, 0x47,
	0xdf, 0x7e, 0xb5, 0x7e, 0x33, 0xe5, 0x64, 0xd6, 0x5f, 0xf1, 0x78, 0x5f, 0x0d, 0x8e, 0x44, 0x06,
	0x11, 0xf8, 0x91, 0x90, 0x8d, 0x2f, 0x4b, 0x67, 0x5f, 0x95, 0xda, 0xc5, 0x57, 0x65, 0x8f, 0x44,
	0x2a, 0x4f, 0xb4, 0xc9, 0xb8, 0xf3, 0xa5, 0x6e, 0xf5, 0xd8, 0xc4, 0xe5, 0xc3, 0xd6, 0x9a, 0x61,
	0xde, 0x01, 0xe1, 0xfb, 0xec, 0x64, 0x9f, 0xe5, 0xf2, 0xcc, 0x8d, 0x74, 0x11, 0x5e, 0xb9, 0xc6,
	0x64, 0xff, 0x2d, 0x02, 0xbe, 0x35, 0xaa, 0xdc, 0x49, 0x93, 0xf0, 0xf8, 0xc2, 0x2b, 0xd7, 0xac,
	0x50, 0xb8, 0x2b, 0xd3, 0xee, 0x35, 0x7f, 0xb9, 0x7b, 0xa5, 0x9a, 0x0b, 0x16, 0x48, 0x7c, 0xe0,
	0x88, 0x87, 0xde, 0xa0, 0x80, 0x92, 0x45, 0xe8, 0xb8, 0xfb, 0x4f, 0xe8, 0x22, 0x54, 0x78, 0x69,
	0x4e, 0x5b, 0x75, 0x7d, 0xb6, 0x55, 0xc3, 0xa3, 0xa3, 0x26, 0xdc, 0xc9, 0xd0, 0xa9, 0x1b, 0xe5,
	0xfb, 0x06, 0x20, 0x36, 0xea, 0x0d, 0xb2, 0x90, 0x6a, 0x5d, 0x48, 0x41, 0x09, 0x76, 0x08, 0xb4,
	0xa2, 0x5d, 0xd2, 0xf2, 0x8e, 0xb2, 0x91, 0x2e, 0x42, 0x06, 0x95, 0xa0, 0x4b, 0xb0, 0xc1, 0x8a,
	0x27, 0x5e, 0xe9, 0x02, 0xde, 0xed, 0xde, 0x45, 0x69, 0x35, 0x37, 0x2c, 0xe3, 0x03, 0x1f, 0xa4,
	0x1f, 0xc2, 0xf8, 0xb0, 0x47, 0xd6, 0x2e, 0xed, 0x0a, 0xe2, 0x26, 0x88, 0x57, 0x67, 0xf7, 0x05,
	0x79, 0x9b, 0x2c, 0x57, 0x72, 0x9e, 0x28, 0xba, 0x82, 0x31, 0x09, 0xba, 0x83, 0x44, 0x45, 0x3b,
	0xa4, 0x59, 0x29, 0xac, 0x97, 0xdc, 0x00, 0xc9, 0x52, 0x90, 0xbc, 0xe3, 0x89, 0xba, 0xdc, 0x7f,
	0x56, 0xaf, 0xf4, 0x9f, 0x6d, 0xd2, 0x70, 0xe3, 0x3c, 0x97, 0xf0, 0x78, 0xb7, 0xb0, 0x7b, 0x21,
	0xd0, 0x13, 0x30, 0x3d, 0x70, 0x37, 0x52, 0x82, 0x46, 0x98, 0x2e, 0xb4, 0x7c, 0x74, 0x07, 0x3c,
	0xfe, 0x38, 0x2e, 0x58, 0xa0, 0xd7, 0x30, 0xba, 0x08, 0x9e, 0xa2, 0x68, 0x97, 0xb4, 0x8c, 0x4c,
	0x58, 0x9c, 0x3b, 0xa6, 0x13, 0x86, 0x14, 0x5d, 0xc7, 0x28, 0x1a, 0x99, 0x1c, 0xe6, 0xee, 0x24,
	0x79, 0x06, 0x68, 0x74, 0x48, 0xee, 0xe4, 0xe3, 0x6c, 0x20, 0x8d, 0x57, 0x56, 0x4f, 0x6c, 0xac,
	0xb3, 0x6c, 0x9c, 0x2b, 0xa7, 0xa4, 0xa5, 0xbf, 0x81, 0x75, 0xdb, 0xa8, 0x3a, 0x49, 0x8e, 0x82,
	0xe6, 0x70, 0x2a, 0x89, 0xee, 0x91, 0xe5, 0x6c, 0x52, 0xf8, 0xa6, 0x2d, 0xad, 0xcc, 0x1d, 0xdd,
	0x80, 0x9c, 0x2e, 0x79, 0xec, 0x14, 0x21, 0x5f, 0xa5, 0xfe, 0xc0, 0xc6, 0x55, 0xa2, 0x4d, 0x10,
	0x35, 0x11, 0x2d, 0x65, 0x0f, 0xc9, 0xda, 0xc4, 0x24, 0x2a, 0x2b, 0xb4, 0x71, 0x33, 0x5a, 0x0a,
	0xda, 0x68, 0x86, 0x2a, 0x17, 0xec, 0x91, 0x08, 0xef, 0x0f, 0xb7, 0x33, 0xfa, 0x2d, 0xd0, 0xb7,
	0xa6, 0x4c, 0x29, 0xdf, 0x25, 0xab, 0x08, 0x1a, 0x51, 0x89, 0x6f, 0x82, 0xf8, 0x46, 0x89, 0x97,
	0xd2, 0xa7, 0x64, 0xcb, 0xca, 0x61, 0x26, 0x73, 0x27, 0x45, 0x79, 0xfb, 0xaa, 0x35, 0xdb, 0xb0,
	0x66, 0xb3, 0x12, 0x84, 0xcb, 0x58, 0xae, 0xbd, 0x43, 0x96, 0xaa, 0xfa, 0x50, 0x82, 0xde, 0xc2,
	0xd9, 0x28, 0x54, 0x47, 0x4f, 0x44, 0x0f, 0xc9, 0xfa, 0x0c, 0xcf, 0x8c, 0x4c, 0xf0, 0x21, 0xbd,
	0x8d, 0x33, 0x41, 0x25, 0xec, 0x07, 0xc2, 0x97, 0xa4, 0xb6, 0x45, 0xc2, 0xb8, 0x91, 0xdc, 0xef,
	0x78, 0x07, 0x4a, 0x97, 0x78, 0xec, 0xc0, 0x48, 0xde, 0x13, 0xd1, 0xef, 0x49, 0x64, 0x64, 0xa6,
	0x9d, 0x0c, 0xf9, 0x66, 0xbe, 0xdb, 0xd0, 0xbb, 0xed, 0xb9, 0xce, 0x72, 0x7f, 0x15, 0x19, 0x4c,
	0xf9, 0x81, 0x10, 0xc6, 0x67, 0x6c, 0xc4, 0x2d, 0x96, 0xa6, 0x75, 0x1f, 0x69, 0x1b, 0x33, 0x36,
	0xe2, 0xf6, 0x38, 0x40, 0xbe, 0xed, 0xe4, 0xe3, 0x2c, 0x48, 0xe8, 0xbd, 0xe0, 0xc2, 0x38, 0x43,
	0x81, 0x1f, 0xac, 0xaa, 0xd5, 0x3b, 0xed, 0x39, 0x5f, 0xbc, 0xa5, 0x0d, 0x45, 0xaa, 0x72, 0xa1,
	0xf2, 0x61, 0x28, 0xfe, 0xfb, 0xa1, 0x48, 0x11, 0xac, 0xca, 0x3f, 0x1f, 0x29, 0xc1, 0x12, 0xa9,
	0x04, 0xfd, 0x2d, 0x74, 0x96, 0xba, 0x07, 0x5e, 0x48, 0x25, 0x3c, 0x99, 0x15, 0xa9, 0x45, 0xf2,
	0x01, 0x92, 0x1e, 0xf0, 0xe4, 0xce, 0xff, 0xae, 0x91, 0xb5, 0xb2, 0xdf, 0x42, 0x4b, 0xc7, 0xaf,
	0x5c, 0x1d, 0x77, 0x6b, 0x5f, 0x18, 0x77, 0x2f, 0x8c, 0xb4, 0xd7, 0x2e, 0x8d, 0xb4, 0xeb, 0x64,
	0xde, 0x3a, 0x9e, 0xe2, 0xff, 0x12, 0xcd, 0x3e, 0x1a, 0xde, 0xd5, 0x4c, 0x19, 0xa3, 0x8d, 0x14,
	0xd0, 0x5b, 0x9b, 0xfd, 0xca, 0xf6, 0xdf, 0xcc, 0xa4, 0x19, 0x4a, 0x3f, 0x3f, 0xfa, 0x06, 0x12,
	0xba, 0xeb, 0x32, 0x80, 0xcf, 0x11, 0x83, 0x56, 0x20, 0xfd, 0x80, 0xc9, 0x74, 0x9e, 0x9e, 0x97,
	0x3d, 0x16, 0xa1, 0x93, 0x3c, 0x3d, 0xf7, 0xdf, 0xc5, 0x40, 0x2d, 0xe2, 0x77, 0xd1, 0x9f, 0xd9,
	0xd9, 0xb5, 0x7e, 0x71, 0x76, 0x1d, 0x2c, 0x80, 0x53, 0x8f, 0x7f, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x71, 0xb9, 0x26, 0x13, 0x0e, 0x00, 0x00,
}
