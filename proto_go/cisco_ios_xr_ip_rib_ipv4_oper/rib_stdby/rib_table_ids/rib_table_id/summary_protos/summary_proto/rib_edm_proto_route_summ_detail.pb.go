// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rib_edm_proto_route_summ_detail.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_rib_table_ids_rib_table_id_summary_protos_summary_proto

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

// Details of Protocol route types
type RibEdmProtoRouteSummDetail_KEYS struct {
	Tableid              uint32   `protobuf:"varint,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Protoid              uint32   `protobuf:"varint,2,opt,name=protoid,proto3" json:"protoid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmProtoRouteSummDetail_KEYS) Reset()         { *m = RibEdmProtoRouteSummDetail_KEYS{} }
func (m *RibEdmProtoRouteSummDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RibEdmProtoRouteSummDetail_KEYS) ProtoMessage()    {}
func (*RibEdmProtoRouteSummDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rib_edm_proto_route_summ_detail_b04b317a0cdc3056, []int{0}
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Unmarshal(m, b)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *RibEdmProtoRouteSummDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Merge(dst, src)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.Size(m)
}
func (m *RibEdmProtoRouteSummDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmProtoRouteSummDetail_KEYS proto.InternalMessageInfo

func (m *RibEdmProtoRouteSummDetail_KEYS) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *RibEdmProtoRouteSummDetail_KEYS) GetProtoid() uint32 {
	if m != nil {
		return m.Protoid
	}
	return 0
}

type RibEdmProtoRouteSummDetail struct {
	// Proto name
	Name string `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance,proto3" json:"instance,omitempty"`
	// Count for proto. instance
	ProtoRouteCount *RibEdmRouteCount `protobuf:"bytes,52,opt,name=proto_route_count,json=protoRouteCount,proto3" json:"proto_route_count,omitempty"`
	// No route type
	RtypeNone *RibEdmRouteCount `protobuf:"bytes,53,opt,name=rtype_none,json=rtypeNone,proto3" json:"rtype_none,omitempty"`
	// Unknown route type
	RtypeOther *RibEdmRouteCount `protobuf:"bytes,54,opt,name=rtype_other,json=rtypeOther,proto3" json:"rtype_other,omitempty"`
	// OSPF route within an area
	RtypeOspfIntra *RibEdmRouteCount `protobuf:"bytes,55,opt,name=rtype_ospf_intra,json=rtypeOspfIntra,proto3" json:"rtype_ospf_intra,omitempty"`
	// OSPF route across diff. areas
	RtypeOspfInter *RibEdmRouteCount `protobuf:"bytes,56,opt,name=rtype_ospf_inter,json=rtypeOspfInter,proto3" json:"rtype_ospf_inter,omitempty"`
	// OSPF external route of type 1
	RtypeOspfExtern1 *RibEdmRouteCount `protobuf:"bytes,57,opt,name=rtype_ospf_extern1,json=rtypeOspfExtern1,proto3" json:"rtype_ospf_extern1,omitempty"`
	// OSPF external route of type 2
	RtypeOspfExtern2 *RibEdmRouteCount `protobuf:"bytes,58,opt,name=rtype_ospf_extern2,json=rtypeOspfExtern2,proto3" json:"rtype_ospf_extern2,omitempty"`
	// IS-IS summary route
	RtypeIsisSum *RibEdmRouteCount `protobuf:"bytes,59,opt,name=rtype_isis_sum,json=rtypeIsisSum,proto3" json:"rtype_isis_sum,omitempty"`
	// IS-IS level 1 route
	RtypeIsisL1 *RibEdmRouteCount `protobuf:"bytes,60,opt,name=rtype_isis_l1,json=rtypeIsisL1,proto3" json:"rtype_isis_l1,omitempty"`
	// IS-IS level 2 route
	RtypeIsisL2 *RibEdmRouteCount `protobuf:"bytes,61,opt,name=rtype_isis_l2,json=rtypeIsisL2,proto3" json:"rtype_isis_l2,omitempty"`
	// IS-IS level1 inter-area route
	RtypeIsisL1Ia *RibEdmRouteCount `protobuf:"bytes,62,opt,name=rtype_isis_l1_ia,json=rtypeIsisL1Ia,proto3" json:"rtype_isis_l1_ia,omitempty"`
	// iBGP route
	RtypeBgpInt *RibEdmRouteCount `protobuf:"bytes,63,opt,name=rtype_bgp_int,json=rtypeBgpInt,proto3" json:"rtype_bgp_int,omitempty"`
	// eBGP route
	RtypeBgpExt *RibEdmRouteCount `protobuf:"bytes,64,opt,name=rtype_bgp_ext,json=rtypeBgpExt,proto3" json:"rtype_bgp_ext,omitempty"`
	// BGP local route
	RtypeBgpLoc *RibEdmRouteCount `protobuf:"bytes,65,opt,name=rtype_bgp_loc,json=rtypeBgpLoc,proto3" json:"rtype_bgp_loc,omitempty"`
	// OSPF NSSA ext. route type 1
	RtypeOspfNssa1 *RibEdmRouteCount `protobuf:"bytes,66,opt,name=rtype_ospf_nssa1,json=rtypeOspfNssa1,proto3" json:"rtype_ospf_nssa1,omitempty"`
	// OSPF NSSA ext. route type 2
	RtypeOspfNssa2 *RibEdmRouteCount `protobuf:"bytes,67,opt,name=rtype_ospf_nssa2,json=rtypeOspfNssa2,proto3" json:"rtype_ospf_nssa2,omitempty"`
	// EIGRP internal route
	RtypeIgrp2Int *RibEdmRouteCount `protobuf:"bytes,68,opt,name=rtype_igrp2_int,json=rtypeIgrp2Int,proto3" json:"rtype_igrp2_int,omitempty"`
	// EIGRP external route
	RtypeIgrp2Ext        *RibEdmRouteCount `protobuf:"bytes,69,opt,name=rtype_igrp2_ext,json=rtypeIgrp2Ext,proto3" json:"rtype_igrp2_ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RibEdmProtoRouteSummDetail) Reset()         { *m = RibEdmProtoRouteSummDetail{} }
func (m *RibEdmProtoRouteSummDetail) String() string { return proto.CompactTextString(m) }
func (*RibEdmProtoRouteSummDetail) ProtoMessage()    {}
func (*RibEdmProtoRouteSummDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rib_edm_proto_route_summ_detail_b04b317a0cdc3056, []int{1}
}
func (m *RibEdmProtoRouteSummDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Unmarshal(m, b)
}
func (m *RibEdmProtoRouteSummDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Marshal(b, m, deterministic)
}
func (dst *RibEdmProtoRouteSummDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmProtoRouteSummDetail.Merge(dst, src)
}
func (m *RibEdmProtoRouteSummDetail) XXX_Size() int {
	return xxx_messageInfo_RibEdmProtoRouteSummDetail.Size(m)
}
func (m *RibEdmProtoRouteSummDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmProtoRouteSummDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmProtoRouteSummDetail proto.InternalMessageInfo

func (m *RibEdmProtoRouteSummDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RibEdmProtoRouteSummDetail) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *RibEdmProtoRouteSummDetail) GetProtoRouteCount() *RibEdmRouteCount {
	if m != nil {
		return m.ProtoRouteCount
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeNone() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeNone
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOther() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOther
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfIntra() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfIntra
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfInter() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfInter
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfExtern1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfExtern1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfExtern2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfExtern2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisSum() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisSum
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIsisL1Ia() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIsisL1Ia
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpInt() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpInt
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpExt() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpExt
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeBgpLoc() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeBgpLoc
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfNssa1() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfNssa1
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeOspfNssa2() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeOspfNssa2
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIgrp2Int() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIgrp2Int
	}
	return nil
}

func (m *RibEdmProtoRouteSummDetail) GetRtypeIgrp2Ext() *RibEdmRouteCount {
	if m != nil {
		return m.RtypeIgrp2Ext
	}
	return nil
}

// Specifics of route count
type RibEdmRouteCount struct {
	// Number of active routes
	ActiveRoutesCount uint32 `protobuf:"varint,1,opt,name=active_routes_count,json=activeRoutesCount,proto3" json:"active_routes_count,omitempty"`
	// Number of backup (inactive) routes
	NumBackupRoutes uint32 `protobuf:"varint,2,opt,name=num_backup_routes,json=numBackupRoutes,proto3" json:"num_backup_routes,omitempty"`
	// Number of paths to active routes
	NumActivePaths uint32 `protobuf:"varint,3,opt,name=num_active_paths,json=numActivePaths,proto3" json:"num_active_paths,omitempty"`
	// Number of paths to backup routes
	NumBackupPaths       uint32   `protobuf:"varint,4,opt,name=num_backup_paths,json=numBackupPaths,proto3" json:"num_backup_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RibEdmRouteCount) Reset()         { *m = RibEdmRouteCount{} }
func (m *RibEdmRouteCount) String() string { return proto.CompactTextString(m) }
func (*RibEdmRouteCount) ProtoMessage()    {}
func (*RibEdmRouteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_rib_edm_proto_route_summ_detail_b04b317a0cdc3056, []int{2}
}
func (m *RibEdmRouteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RibEdmRouteCount.Unmarshal(m, b)
}
func (m *RibEdmRouteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RibEdmRouteCount.Marshal(b, m, deterministic)
}
func (dst *RibEdmRouteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RibEdmRouteCount.Merge(dst, src)
}
func (m *RibEdmRouteCount) XXX_Size() int {
	return xxx_messageInfo_RibEdmRouteCount.Size(m)
}
func (m *RibEdmRouteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_RibEdmRouteCount.DiscardUnknown(m)
}

var xxx_messageInfo_RibEdmRouteCount proto.InternalMessageInfo

func (m *RibEdmRouteCount) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumBackupRoutes() uint32 {
	if m != nil {
		return m.NumBackupRoutes
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumActivePaths() uint32 {
	if m != nil {
		return m.NumActivePaths
	}
	return 0
}

func (m *RibEdmRouteCount) GetNumBackupPaths() uint32 {
	if m != nil {
		return m.NumBackupPaths
	}
	return 0
}

func init() {
	proto.RegisterType((*RibEdmProtoRouteSummDetail_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_proto_route_summ_detail_KEYS")
	proto.RegisterType((*RibEdmProtoRouteSummDetail)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_proto_route_summ_detail")
	proto.RegisterType((*RibEdmRouteCount)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.summary_protos.summary_proto.rib_edm_route_count")
}

func init() {
	proto.RegisterFile("rib_edm_proto_route_summ_detail.proto", fileDescriptor_rib_edm_proto_route_summ_detail_b04b317a0cdc3056)
}

var fileDescriptor_rib_edm_proto_route_summ_detail_b04b317a0cdc3056 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x5d, 0x6b, 0x14, 0x3d,
	0x14, 0xc7, 0xc9, 0xf3, 0x14, 0xb5, 0xa9, 0x7d, 0xd9, 0x14, 0x24, 0x78, 0xe3, 0xb2, 0x28, 0x2c,
	0x5e, 0x0c, 0xec, 0xb4, 0xbe, 0xbf, 0x76, 0xeb, 0x5e, 0x2c, 0x96, 0x2a, 0xdb, 0x2b, 0xbd, 0x09,
	0x99, 0x99, 0x74, 0x1b, 0xdc, 0x49, 0x42, 0x92, 0x29, 0xbb, 0x9f, 0x43, 0x04, 0x11, 0x15, 0xf1,
	0xbb, 0xf8, 0x85, 0xfc, 0x04, 0x92, 0x64, 0x76, 0x99, 0xae, 0x0b, 0xbd, 0x92, 0x99, 0xbb, 0x39,
	0xff, 0xf3, 0x32, 0xbf, 0x1c, 0x38, 0xc9, 0x81, 0x77, 0x34, 0x4f, 0x08, 0xcb, 0x72, 0xa2, 0xb4,
	0xb4, 0x92, 0x68, 0x59, 0x58, 0x46, 0x4c, 0x91, 0xe7, 0x24, 0x63, 0x96, 0xf2, 0x49, 0xe4, 0x75,
	0x44, 0x52, 0x6e, 0x52, 0x49, 0xb8, 0x34, 0x64, 0xaa, 0x09, 0x57, 0xc4, 0xa5, 0x71, 0x75, 0xbe,
	0x4f, 0xa4, 0x62, 0x3a, 0x72, 0x96, 0xb1, 0x59, 0x32, 0xf3, 0x5f, 0x96, 0x26, 0x13, 0x46, 0x78,
	0x66, 0x2e, 0x58, 0x91, 0xab, 0x4a, 0xf5, 0x2c, 0xfc, 0xc9, 0x5c, 0x34, 0x3b, 0xef, 0xe1, 0xed,
	0x4b, 0x48, 0xc8, 0xeb, 0xc1, 0xbb, 0x13, 0x84, 0xe1, 0x55, 0x5f, 0x90, 0x67, 0x18, 0xb4, 0x41,
	0x77, 0x73, 0x34, 0x37, 0x9d, 0xc7, 0x67, 0xf2, 0x0c, 0xff, 0x17, 0x3c, 0xa5, 0xd9, 0xf9, 0x7d,
	0x03, 0xde, 0xba, 0xa4, 0x38, 0x42, 0x70, 0x4d, 0xd0, 0x9c, 0xe1, 0xb8, 0x0d, 0xba, 0xeb, 0x23,
	0xff, 0x8d, 0x6e, 0xc2, 0x6b, 0x5c, 0x18, 0x4b, 0x45, 0xca, 0xf0, 0x9e, 0xd7, 0x17, 0x36, 0xfa,
	0x01, 0x60, 0xab, 0x5a, 0x2b, 0x95, 0x85, 0xb0, 0x78, 0xbf, 0x0d, 0xba, 0x1b, 0xb1, 0x8d, 0xfe,
	0x71, 0xb7, 0xa2, 0xf9, 0x69, 0x2a, 0xff, 0x1e, 0x6d, 0x7b, 0xd7, 0xc8, 0x29, 0x87, 0x4e, 0x40,
	0x1f, 0x01, 0x84, 0xda, 0xce, 0x14, 0x23, 0x42, 0x0a, 0x86, 0xef, 0xd5, 0xc8, 0xb6, 0xee, 0x39,
	0x8e, 0xa5, 0x60, 0xe8, 0x13, 0x80, 0x1b, 0x81, 0x4a, 0xda, 0x33, 0xa6, 0xf1, 0xfd, 0x1a, 0xb1,
	0x42, 0x7b, 0xde, 0x38, 0x0e, 0xf4, 0x1d, 0xc0, 0x9d, 0x92, 0xcb, 0xa8, 0x53, 0xc2, 0x85, 0xd5,
	0x14, 0x3f, 0xa8, 0x11, 0x6e, 0x2b, 0xc0, 0x19, 0x75, 0x3a, 0x74, 0x2c, 0x2b, 0x00, 0x99, 0xc6,
	0x0f, 0x9b, 0x02, 0xc8, 0x34, 0xfa, 0x09, 0x20, 0xaa, 0x00, 0xb2, 0xa9, 0x65, 0x5a, 0xf4, 0xf0,
	0xa3, 0x1a, 0x11, 0x77, 0x16, 0x88, 0x83, 0x40, 0xb3, 0x1a, 0x32, 0xc6, 0x8f, 0x1b, 0x04, 0x19,
	0xa3, 0x2f, 0x00, 0x86, 0xe6, 0x12, 0x6e, 0xb8, 0x71, 0xf7, 0x14, 0x7e, 0x52, 0x23, 0xe0, 0x75,
	0xcf, 0x32, 0x34, 0xdc, 0x9c, 0x14, 0x39, 0xfa, 0x0c, 0xe0, 0x66, 0x05, 0x6e, 0xd2, 0xc3, 0x4f,
	0x6b, 0x64, 0xdb, 0x58, 0xb0, 0x1d, 0xf5, 0xfe, 0x42, 0x8b, 0xf1, 0xb3, 0x66, 0xa0, 0xc5, 0xe8,
	0xdb, 0x62, 0x7a, 0xcb, 0xae, 0x11, 0x4e, 0xf1, 0xf3, 0x1a, 0xe9, 0x36, 0x2b, 0x8d, 0x1b, 0xd2,
	0x4a, 0xeb, 0x92, 0xb1, 0x72, 0x97, 0x0b, 0x7e, 0x51, 0x7b, 0xeb, 0xfa, 0x63, 0x35, 0x14, 0x76,
	0x09, 0x8d, 0x4d, 0x2d, 0x7e, 0xd9, 0x04, 0xb4, 0xc1, 0x74, 0x19, 0x6d, 0x22, 0x53, 0x7c, 0xd0,
	0x04, 0xb4, 0x23, 0x99, 0x2e, 0x3f, 0x17, 0xc2, 0x18, 0xda, 0xc3, 0xfd, 0x46, 0x3c, 0x17, 0xc7,
	0x8e, 0x65, 0x15, 0x60, 0x8c, 0x0f, 0x1b, 0x03, 0x18, 0xa3, 0xaf, 0x00, 0x6e, 0x97, 0x23, 0x3b,
	0xd6, 0x2a, 0xf6, 0x43, 0xf1, 0xaa, 0xfe, 0x89, 0x75, 0x2c, 0x6e, 0x2c, 0x96, 0xf1, 0xdc, 0x60,
	0x0c, 0x9a, 0x81, 0x37, 0x98, 0xda, 0xce, 0x2f, 0x00, 0x77, 0x57, 0x84, 0xa1, 0x08, 0xee, 0xd2,
	0xd4, 0xf2, 0x73, 0x16, 0x54, 0x53, 0x6e, 0xce, 0x61, 0x99, 0x6f, 0x05, 0x97, 0x5f, 0x62, 0x4d,
	0xd8, 0x62, 0xef, 0xc2, 0x96, 0x28, 0x72, 0x92, 0xd0, 0xf4, 0x43, 0xa1, 0xca, 0x9c, 0x72, 0xc1,
	0xdf, 0x16, 0x45, 0xde, 0xf7, 0x7a, 0x48, 0x40, 0x5d, 0xb8, 0xe3, 0x62, 0xcb, 0xfa, 0x8a, 0xda,
	0x33, 0x83, 0xff, 0xf7, 0xa1, 0x5b, 0xa2, 0xc8, 0x0f, 0xbc, 0xfc, 0xd6, 0xa9, 0xf3, 0xc8, 0xb2,
	0x6a, 0x88, 0x5c, 0x5b, 0x44, 0x86, 0xa2, 0x3e, 0x32, 0xb9, 0xe2, 0x4f, 0xbc, 0xf7, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0xc3, 0x5f, 0x7b, 0x29, 0x0d, 0x00, 0x00,
}
