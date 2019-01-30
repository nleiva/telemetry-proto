// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_protocol.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_process_information_protocol_summary

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

// OSPF Protocol Information
type OspfShProtocol_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShProtocol_KEYS) Reset()         { *m = OspfShProtocol_KEYS{} }
func (m *OspfShProtocol_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShProtocol_KEYS) ProtoMessage()    {}
func (*OspfShProtocol_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4559d5f25e5c5134, []int{0}
}

func (m *OspfShProtocol_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtocol_KEYS.Unmarshal(m, b)
}
func (m *OspfShProtocol_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtocol_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShProtocol_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtocol_KEYS.Merge(m, src)
}
func (m *OspfShProtocol_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShProtocol_KEYS.Size(m)
}
func (m *OspfShProtocol_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtocol_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtocol_KEYS proto.InternalMessageInfo

func (m *OspfShProtocol_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShProtocol_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShProtocol struct {
	// Router ID
	ProtocolRouterId string `protobuf:"bytes,50,opt,name=protocol_router_id,json=protocolRouterId,proto3" json:"protocol_router_id,omitempty"`
	// Administrative distance
	ProtocolDistance uint32 `protobuf:"varint,51,opt,name=protocol_distance,json=protocolDistance,proto3" json:"protocol_distance,omitempty"`
	// Administrative Distance for Inter Area routes
	AdministrativeDistanceInterArea uint32 `protobuf:"varint,52,opt,name=administrative_distance_inter_area,json=administrativeDistanceInterArea,proto3" json:"administrative_distance_inter_area,omitempty"`
	// Administrative Distance for External routes
	AdministrativeDistanceExternal uint32 `protobuf:"varint,53,opt,name=administrative_distance_external,json=administrativeDistanceExternal,proto3" json:"administrative_distance_external,omitempty"`
	// True if NSF enabled
	ProtocolNsf bool `protobuf:"varint,54,opt,name=protocol_nsf,json=protocolNsf,proto3" json:"protocol_nsf,omitempty"`
	// Distribute List In
	DistListIn           string   `protobuf:"bytes,55,opt,name=dist_list_in,json=distListIn,proto3" json:"dist_list_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShProtocol) Reset()         { *m = OspfShProtocol{} }
func (m *OspfShProtocol) String() string { return proto.CompactTextString(m) }
func (*OspfShProtocol) ProtoMessage()    {}
func (*OspfShProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_4559d5f25e5c5134, []int{1}
}

func (m *OspfShProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtocol.Unmarshal(m, b)
}
func (m *OspfShProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtocol.Marshal(b, m, deterministic)
}
func (m *OspfShProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtocol.Merge(m, src)
}
func (m *OspfShProtocol) XXX_Size() int {
	return xxx_messageInfo_OspfShProtocol.Size(m)
}
func (m *OspfShProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtocol proto.InternalMessageInfo

func (m *OspfShProtocol) GetProtocolRouterId() string {
	if m != nil {
		return m.ProtocolRouterId
	}
	return ""
}

func (m *OspfShProtocol) GetProtocolDistance() uint32 {
	if m != nil {
		return m.ProtocolDistance
	}
	return 0
}

func (m *OspfShProtocol) GetAdministrativeDistanceInterArea() uint32 {
	if m != nil {
		return m.AdministrativeDistanceInterArea
	}
	return 0
}

func (m *OspfShProtocol) GetAdministrativeDistanceExternal() uint32 {
	if m != nil {
		return m.AdministrativeDistanceExternal
	}
	return 0
}

func (m *OspfShProtocol) GetProtocolNsf() bool {
	if m != nil {
		return m.ProtocolNsf
	}
	return false
}

func (m *OspfShProtocol) GetDistListIn() string {
	if m != nil {
		return m.DistListIn
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShProtocol_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.protocol_summary.ospf_sh_protocol_KEYS")
	proto.RegisterType((*OspfShProtocol)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.protocol_summary.ospf_sh_protocol")
}

func init() { proto.RegisterFile("ospf_sh_protocol.proto", fileDescriptor_4559d5f25e5c5134) }

var fileDescriptor_4559d5f25e5c5134 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0x1e, 0xbe, 0xaf, 0xc6, 0x0a, 0x35, 0xa0, 0xc4, 0x8b, 0xae, 0x3d, 0x15, 0x94,
	0x3d, 0xd8, 0xaa, 0x67, 0xc1, 0x82, 0xa5, 0xd2, 0x43, 0xc5, 0x83, 0xa7, 0x18, 0x77, 0x13, 0x1c,
	0xe8, 0x26, 0xcb, 0x4c, 0xba, 0xd4, 0x9f, 0xe8, 0xbf, 0x92, 0xa4, 0xbb, 0x15, 0x8b, 0x5e, 0xc2,
	0x30, 0xef, 0x33, 0xcf, 0x90, 0x61, 0xc7, 0x8e, 0x2a, 0x23, 0xe9, 0x5d, 0x56, 0xe8, 0xbc, 0xcb,
	0xdd, 0x32, 0x8b, 0x05, 0x7f, 0xcd, 0x81, 0x72, 0x27, 0xc1, 0x91, 0x5c, 0xa3, 0x84, 0xaa, 0x1e,
	0xcb, 0x48, 0xba, 0x4a, 0x63, 0x16, 0xaa, 0xc0, 0xe5, 0x9a, 0x48, 0x53, 0x5b, 0x65, 0x35, 0x9a,
	0xf8, 0xb4, 0x0d, 0x09, 0xd6, 0x38, 0x2c, 0x95, 0x07, 0x67, 0xb3, 0xd6, 0x2f, 0x69, 0x55, 0x96,
	0x0a, 0x3f, 0x06, 0xcf, 0xec, 0x68, 0x77, 0xb7, 0x9c, 0x4d, 0x5e, 0x9e, 0xf8, 0x39, 0xeb, 0xb5,
	0x02, 0xab, 0x4a, 0x2d, 0x92, 0x34, 0x19, 0xee, 0x2d, 0xf6, 0x9b, 0xde, 0x5c, 0x95, 0x9a, 0x9f,
	0xb0, 0x6e, 0x8d, 0x66, 0x13, 0x77, 0x62, 0xfc, 0xbf, 0x46, 0x13, 0xa2, 0xc1, 0x67, 0x87, 0xf5,
	0x77, 0xbd, 0xfc, 0x92, 0xf1, 0xed, 0x0e, 0x74, 0x2b, 0xaf, 0x51, 0x42, 0x21, 0xae, 0xe2, 0x64,
	0xbf, 0x4d, 0x16, 0x31, 0x98, 0x16, 0xfc, 0x82, 0x1d, 0x6e, 0xe9, 0x02, 0xc8, 0x2b, 0x9b, 0x6b,
	0x31, 0x4a, 0x93, 0xe1, 0xc1, 0x37, 0x7c, 0xdf, 0xf4, 0xf9, 0x8c, 0x0d, 0x54, 0x51, 0x82, 0x05,
	0xf2, 0xa8, 0x3c, 0xd4, 0x7a, 0x3b, 0x22, 0xc1, 0x86, 0x45, 0x0a, 0xb5, 0x12, 0xe3, 0x38, 0x7d,
	0xf6, 0x93, 0x6c, 0x1d, 0xd3, 0xc0, 0xdd, 0xa1, 0x56, 0xfc, 0x81, 0xa5, 0x7f, 0xc9, 0xf4, 0xda,
	0x6b, 0xb4, 0x6a, 0x29, 0xae, 0xa3, 0xea, 0xf4, 0x77, 0xd5, 0xa4, 0xa1, 0x9a, 0x23, 0x6e, 0xfe,
	0x60, 0xc9, 0x88, 0x9b, 0x34, 0x19, 0x76, 0xe3, 0x11, 0x63, 0x6f, 0x4e, 0x86, 0xa7, 0xac, 0x17,
	0xec, 0x72, 0x19, 0x1e, 0xb0, 0xe2, 0x36, 0x9e, 0x83, 0x85, 0xde, 0x23, 0x90, 0x9f, 0xda, 0xb7,
	0x7f, 0x11, 0x1f, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x9c, 0x96, 0x94, 0x25, 0x02, 0x00,
	0x00,
}
