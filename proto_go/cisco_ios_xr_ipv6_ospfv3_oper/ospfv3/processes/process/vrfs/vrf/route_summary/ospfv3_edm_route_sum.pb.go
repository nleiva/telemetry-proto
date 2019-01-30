// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_route_sum.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary

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

// OSPFv3 route summary information
type Ospfv3EdmRouteSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRouteSum_KEYS) Reset()         { *m = Ospfv3EdmRouteSum_KEYS{} }
func (m *Ospfv3EdmRouteSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRouteSum_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmRouteSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_route_sum_5dc5ce9f71adef15, []int{0}
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmRouteSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Merge(dst, src)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.Size(m)
}
func (m *Ospfv3EdmRouteSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRouteSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRouteSum_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmRouteSum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRouteSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ospfv3EdmRouteSum struct {
	// Route summary of a route ID
	RouteId string `protobuf:"bytes,50,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	// Intra route summary
	IntraAreaRoute uint32 `protobuf:"varint,51,opt,name=intra_area_route,json=intraAreaRoute,proto3" json:"intra_area_route,omitempty"`
	// Inter route summary
	InterAreaRoute uint32 `protobuf:"varint,52,opt,name=inter_area_route,json=interAreaRoute,proto3" json:"inter_area_route,omitempty"`
	// Extern 1 route summary
	ExternOneRoute uint32 `protobuf:"varint,53,opt,name=extern_one_route,json=externOneRoute,proto3" json:"extern_one_route,omitempty"`
	// Extern 2 route summary
	ExternTwoRoute uint32 `protobuf:"varint,54,opt,name=extern_two_route,json=externTwoRoute,proto3" json:"extern_two_route,omitempty"`
	// NSSA 1 route summary
	NssaOneRoute uint32 `protobuf:"varint,55,opt,name=nssa_one_route,json=nssaOneRoute,proto3" json:"nssa_one_route,omitempty"`
	// NSSA 2 route summary
	NssaTwoRoute uint32 `protobuf:"varint,56,opt,name=nssa_two_route,json=nssaTwoRoute,proto3" json:"nssa_two_route,omitempty"`
	// Total route summary
	TotalSentRoute uint32 `protobuf:"varint,57,opt,name=total_sent_route,json=totalSentRoute,proto3" json:"total_sent_route,omitempty"`
	// Route connected
	RouteConnected uint32 `protobuf:"varint,58,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	// Redistribution route summary
	RedistributionRoute uint32 `protobuf:"varint,59,opt,name=redistribution_route,json=redistributionRoute,proto3" json:"redistribution_route,omitempty"`
	// Total route received summary
	TotalReceivedRoute   uint32   `protobuf:"varint,60,opt,name=total_received_route,json=totalReceivedRoute,proto3" json:"total_received_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRouteSum) Reset()         { *m = Ospfv3EdmRouteSum{} }
func (m *Ospfv3EdmRouteSum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRouteSum) ProtoMessage()    {}
func (*Ospfv3EdmRouteSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_route_sum_5dc5ce9f71adef15, []int{1}
}
func (m *Ospfv3EdmRouteSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Unmarshal(m, b)
}
func (m *Ospfv3EdmRouteSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmRouteSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRouteSum.Merge(dst, src)
}
func (m *Ospfv3EdmRouteSum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRouteSum.Size(m)
}
func (m *Ospfv3EdmRouteSum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRouteSum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRouteSum proto.InternalMessageInfo

func (m *Ospfv3EdmRouteSum) GetRouteId() string {
	if m != nil {
		return m.RouteId
	}
	return ""
}

func (m *Ospfv3EdmRouteSum) GetIntraAreaRoute() uint32 {
	if m != nil {
		return m.IntraAreaRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetInterAreaRoute() uint32 {
	if m != nil {
		return m.InterAreaRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetExternOneRoute() uint32 {
	if m != nil {
		return m.ExternOneRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetExternTwoRoute() uint32 {
	if m != nil {
		return m.ExternTwoRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetNssaOneRoute() uint32 {
	if m != nil {
		return m.NssaOneRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetNssaTwoRoute() uint32 {
	if m != nil {
		return m.NssaTwoRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetTotalSentRoute() uint32 {
	if m != nil {
		return m.TotalSentRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetRouteConnected() uint32 {
	if m != nil {
		return m.RouteConnected
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetRedistributionRoute() uint32 {
	if m != nil {
		return m.RedistributionRoute
	}
	return 0
}

func (m *Ospfv3EdmRouteSum) GetTotalReceivedRoute() uint32 {
	if m != nil {
		return m.TotalReceivedRoute
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmRouteSum_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.route_summary.ospfv3_edm_route_sum_KEYS")
	proto.RegisterType((*Ospfv3EdmRouteSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.route_summary.ospfv3_edm_route_sum")
}

func init() {
	proto.RegisterFile("ospfv3_edm_route_sum.proto", fileDescriptor_ospfv3_edm_route_sum_5dc5ce9f71adef15)
}

var fileDescriptor_ospfv3_edm_route_sum_5dc5ce9f71adef15 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0xe9, 0x7b, 0xf0, 0xfe, 0xcc, 0xeb, 0xeb, 0x2b, 0x79, 0x5d, 0xa4, 0xae, 0x6a, 0x11,
	0xcc, 0x2a, 0xa8, 0xd5, 0xfa, 0x77, 0x23, 0xe2, 0x42, 0x44, 0x85, 0xd4, 0x4d, 0x57, 0xc3, 0x34,
	0xb9, 0x81, 0x01, 0x33, 0x13, 0xee, 0x4c, 0xd3, 0xfa, 0x61, 0xfd, 0x2e, 0x92, 0xb9, 0xd3, 0xb4,
	0x95, 0x6e, 0xc2, 0xbd, 0xe7, 0xfc, 0xce, 0xb9, 0x21, 0x84, 0xed, 0x69, 0x53, 0xe6, 0xd5, 0x88,
	0x43, 0x56, 0x70, 0xd4, 0x73, 0x0b, 0xdc, 0xcc, 0x8b, 0xb8, 0x44, 0x6d, 0x75, 0xf0, 0x94, 0x4a,
	0x93, 0x6a, 0x2e, 0xb5, 0xe1, 0x4b, 0xe4, 0xb2, 0xac, 0xc6, 0xdc, 0xd3, 0xba, 0x04, 0x8c, 0x69,
	0xae, 0xd9, 0x14, 0x8c, 0x01, 0xb3, 0x9a, 0xe2, 0x0a, 0x73, 0xf7, 0x88, 0x9b, 0xc2, 0x42, 0xe0,
	0xfb, 0x70, 0xca, 0xfa, 0xbb, 0x8e, 0xf1, 0xc7, 0xfb, 0xe9, 0x24, 0xd8, 0x67, 0x6d, 0x1f, 0xe7,
	0x4a, 0x14, 0x10, 0xb6, 0x06, 0xad, 0xe8, 0x77, 0xf2, 0xc7, 0x6b, 0xcf, 0xa2, 0x80, 0xa0, 0xcf,
	0x7e, 0x55, 0x98, 0x93, 0xfd, 0xcd, 0xd9, 0x3f, 0x2b, 0xcc, 0x6b, 0x6b, 0xf8, 0xf1, 0x9d, 0xf5,
	0x76, 0x75, 0xd7, 0x19, 0x5a, 0x64, 0x16, 0x9e, 0x50, 0xc6, 0xed, 0x0f, 0x59, 0x10, 0xb1, 0xae,
	0x54, 0x16, 0x05, 0x17, 0x08, 0x82, 0x22, 0xe1, 0x68, 0xd0, 0x8a, 0xfe, 0x26, 0x1d, 0xa7, 0xdf,
	0x22, 0x88, 0xa4, 0x56, 0x3d, 0x09, 0xb8, 0x49, 0x9e, 0x36, 0x24, 0xe0, 0x16, 0x09, 0x4b, 0x0b,
	0xa8, 0xb8, 0x56, 0xe0, 0xc9, 0x33, 0x22, 0x49, 0x7f, 0x51, 0xf0, 0x95, 0xb4, 0x0b, 0xed, 0xc9,
	0xf1, 0x26, 0xf9, 0xba, 0xd0, 0x44, 0x1e, 0xb0, 0x8e, 0x32, 0x46, 0x6c, 0x34, 0x9e, 0x3b, 0xae,
	0x5d, 0xab, 0x4d, 0xdf, 0x8a, 0x5a, 0xb7, 0x5d, 0xac, 0xa9, 0xa6, 0x2b, 0x62, 0x5d, 0xab, 0xad,
	0x78, 0xe3, 0x06, 0x94, 0xf5, 0xdc, 0x25, 0x5d, 0x75, 0xfa, 0x04, 0x94, 0x25, 0xf2, 0x90, 0xfd,
	0xa3, 0x0f, 0x97, 0x6a, 0xa5, 0x20, 0xb5, 0x90, 0x85, 0x57, 0x04, 0x3a, 0xf9, 0x6e, 0xa5, 0x06,
	0xc7, 0xac, 0x87, 0x90, 0x49, 0x63, 0x51, 0xce, 0xe6, 0x56, 0x6a, 0xe5, 0x6b, 0xaf, 0x1d, 0xfd,
	0x7f, 0xdb, 0xa3, 0xee, 0x23, 0xd6, 0xa3, 0xb7, 0x40, 0x48, 0x41, 0x56, 0x90, 0xf9, 0xc8, 0x8d,
	0x8b, 0x04, 0xce, 0x4b, 0xbc, 0xe5, 0x12, 0xb3, 0x1f, 0xee, 0x87, 0x1c, 0x7d, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x18, 0xe2, 0x74, 0x3d, 0xae, 0x02, 0x00, 0x00,
}
