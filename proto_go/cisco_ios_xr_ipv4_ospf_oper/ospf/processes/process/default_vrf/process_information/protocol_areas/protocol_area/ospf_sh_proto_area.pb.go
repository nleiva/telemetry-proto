// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_proto_area.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_process_information_protocol_areas_protocol_area

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

// OSPF Protocol Area Information
type OspfShProtoArea_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShProtoArea_KEYS) Reset()         { *m = OspfShProtoArea_KEYS{} }
func (m *OspfShProtoArea_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoArea_KEYS) ProtoMessage()    {}
func (*OspfShProtoArea_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_proto_area_b0257756c086e7a0, []int{0}
}
func (m *OspfShProtoArea_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Unmarshal(m, b)
}
func (m *OspfShProtoArea_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShProtoArea_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoArea_KEYS.Merge(dst, src)
}
func (m *OspfShProtoArea_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Size(m)
}
func (m *OspfShProtoArea_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoArea_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoArea_KEYS proto.InternalMessageInfo

func (m *OspfShProtoArea_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShProtoArea_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShProtoArea_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type OspfShProtoArea struct {
	// Area ID string in decimal and dotted-decimal format
	ProtcolArea string `protobuf:"bytes,50,opt,name=protcol_area,json=protcolArea,proto3" json:"protcol_area,omitempty"`
	// MPLS-TE enabled
	ProtocolMpls bool `protobuf:"varint,51,opt,name=protocol_mpls,json=protocolMpls,proto3" json:"protocol_mpls,omitempty"`
	// Distribute List In
	ProtocolAreaDistListIn string `protobuf:"bytes,52,opt,name=protocol_area_dist_list_in,json=protocolAreaDistListIn,proto3" json:"protocol_area_dist_list_in,omitempty"`
	// Interface list
	ProtocolInterfaces   []*OspfShProtoIntf `protobuf:"bytes,53,rep,name=protocol_interfaces,json=protocolInterfaces,proto3" json:"protocol_interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OspfShProtoArea) Reset()         { *m = OspfShProtoArea{} }
func (m *OspfShProtoArea) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoArea) ProtoMessage()    {}
func (*OspfShProtoArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_proto_area_b0257756c086e7a0, []int{1}
}
func (m *OspfShProtoArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoArea.Unmarshal(m, b)
}
func (m *OspfShProtoArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoArea.Marshal(b, m, deterministic)
}
func (dst *OspfShProtoArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoArea.Merge(dst, src)
}
func (m *OspfShProtoArea) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoArea.Size(m)
}
func (m *OspfShProtoArea) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoArea.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoArea proto.InternalMessageInfo

func (m *OspfShProtoArea) GetProtcolArea() string {
	if m != nil {
		return m.ProtcolArea
	}
	return ""
}

func (m *OspfShProtoArea) GetProtocolMpls() bool {
	if m != nil {
		return m.ProtocolMpls
	}
	return false
}

func (m *OspfShProtoArea) GetProtocolAreaDistListIn() string {
	if m != nil {
		return m.ProtocolAreaDistListIn
	}
	return ""
}

func (m *OspfShProtoArea) GetProtocolInterfaces() []*OspfShProtoIntf {
	if m != nil {
		return m.ProtocolInterfaces
	}
	return nil
}

// OSPF Protocol Interface Information
type OspfShProtoIntf struct {
	// Interface
	ProtocolInterfaceName string `protobuf:"bytes,1,opt,name=protocol_interface_name,json=protocolInterfaceName,proto3" json:"protocol_interface_name,omitempty"`
	// Authentication type
	ProtocolAuthenticationType string `protobuf:"bytes,2,opt,name=protocol_authentication_type,json=protocolAuthenticationType,proto3" json:"protocol_authentication_type,omitempty"`
	// Distribute List In
	ProtocolInterfaceDistListIn string   `protobuf:"bytes,3,opt,name=protocol_interface_dist_list_in,json=protocolInterfaceDistListIn,proto3" json:"protocol_interface_dist_list_in,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *OspfShProtoIntf) Reset()         { *m = OspfShProtoIntf{} }
func (m *OspfShProtoIntf) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoIntf) ProtoMessage()    {}
func (*OspfShProtoIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_proto_area_b0257756c086e7a0, []int{2}
}
func (m *OspfShProtoIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoIntf.Unmarshal(m, b)
}
func (m *OspfShProtoIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoIntf.Marshal(b, m, deterministic)
}
func (dst *OspfShProtoIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoIntf.Merge(dst, src)
}
func (m *OspfShProtoIntf) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoIntf.Size(m)
}
func (m *OspfShProtoIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoIntf.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoIntf proto.InternalMessageInfo

func (m *OspfShProtoIntf) GetProtocolInterfaceName() string {
	if m != nil {
		return m.ProtocolInterfaceName
	}
	return ""
}

func (m *OspfShProtoIntf) GetProtocolAuthenticationType() string {
	if m != nil {
		return m.ProtocolAuthenticationType
	}
	return ""
}

func (m *OspfShProtoIntf) GetProtocolInterfaceDistListIn() string {
	if m != nil {
		return m.ProtocolInterfaceDistListIn
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShProtoArea_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_area_KEYS")
	proto.RegisterType((*OspfShProtoArea)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_area")
	proto.RegisterType((*OspfShProtoIntf)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_intf")
}

func init() {
	proto.RegisterFile("ospf_sh_proto_area.proto", fileDescriptor_ospf_sh_proto_area_b0257756c086e7a0)
}

var fileDescriptor_ospf_sh_proto_area_b0257756c086e7a0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x85, 0x6c, 0xc0, 0xae, 0x69, 0x7b, 0x61, 0xd1, 0x9a, 0x68, 0x0b, 0x54, 0x75, 0x17, 0x4d,
	0x1a, 0x6c, 0xb7, 0x43, 0xa7, 0x1a, 0x70, 0x06, 0x23, 0x1f, 0x83, 0x92, 0x25, 0x13, 0xc1, 0x48,
	0x14, 0x4c, 0x40, 0x22, 0x09, 0x1e, 0x6d, 0xc4, 0x73, 0x7e, 0x53, 0xfe, 0x47, 0x86, 0xfc, 0xa0,
	0x80, 0xb4, 0x25, 0x47, 0x51, 0xe6, 0x2c, 0xc2, 0xdd, 0xf1, 0xf1, 0xbd, 0xc7, 0x7b, 0x42, 0x44,
	0x81, 0xce, 0x29, 0x6c, 0xa8, 0x36, 0xca, 0x2a, 0xca, 0x0c, 0x67, 0xb1, 0x2f, 0xb1, 0x4a, 0x05,
	0xa4, 0x8a, 0x0a, 0x05, 0xf4, 0xde, 0x50, 0xa1, 0x77, 0x0b, 0xea, 0xb1, 0x4a, 0x73, 0x13, 0xbb,
	0xca, 0xe1, 0x52, 0x0e, 0xc0, 0xa1, 0xaa, 0xe2, 0x8c, 0xe7, 0x6c, 0x5b, 0x58, 0xba, 0x33, 0xf5,
	0x29, 0x15, 0x32, 0x57, 0xa6, 0x64, 0x56, 0x28, 0x79, 0x60, 0x4e, 0x55, 0xe1, 0x75, 0xa0, 0xd9,
	0x4e, 0x15, 0x9a, 0xb4, 0xcd, 0xd0, 0xf3, 0xb3, 0xdb, 0x6b, 0xfc, 0x0b, 0x8d, 0x2a, 0x3a, 0xc9,
	0x4a, 0x4e, 0x82, 0x30, 0x88, 0x06, 0xc9, 0xf0, 0x38, 0xbb, 0x62, 0x25, 0xc7, 0x13, 0xd4, 0xf7,
	0x78, 0x91, 0x91, 0x4e, 0x18, 0x44, 0xe3, 0xa4, 0xe7, 0xda, 0x75, 0x86, 0x09, 0xea, 0xb3, 0x2c,
	0x33, 0x1c, 0x80, 0x74, 0xfd, 0xb5, 0xaa, 0x9d, 0x3e, 0x75, 0x10, 0x6e, 0x2b, 0x1e, 0xc5, 0x6c,
	0xe5, 0x8b, 0xcc, 0x6a, 0x31, 0x37, 0x5b, 0x3a, 0xc8, 0x6f, 0x34, 0xae, 0xbd, 0x97, 0xba, 0x00,
	0x32, 0x0f, 0x83, 0xe8, 0x53, 0x32, 0xaa, 0x86, 0x97, 0xba, 0x00, 0xfc, 0x0f, 0x7d, 0x6b, 0x3c,
	0x90, 0x66, 0x02, 0x2c, 0x2d, 0xdc, 0x47, 0x48, 0xb2, 0xf0, 0xac, 0x5f, 0x2b, 0x84, 0xa3, 0x5d,
	0x09, 0xb0, 0x17, 0x02, 0xec, 0x5a, 0xe2, 0xc7, 0x00, 0x7d, 0xae, 0x2f, 0x0b, 0x69, 0xb9, 0xc9,
	0x59, 0xca, 0x81, 0xfc, 0x09, 0xbb, 0xd1, 0x70, 0xf6, 0x10, 0xc4, 0x1f, 0x1c, 0x4e, 0xdc, 0xdc,
	0x93, 0x90, 0x36, 0x4f, 0x70, 0x85, 0x58, 0xd7, 0xfe, 0xa6, 0xcf, 0xc1, 0xdb, 0x95, 0x3a, 0x28,
	0xfe, 0x8b, 0x26, 0xed, 0xd7, 0xbc, 0x8e, 0xf2, 0x4b, 0x8b, 0xcb, 0x87, 0xfa, 0x1f, 0xfd, 0x38,
	0xd9, 0xd8, 0xda, 0x0d, 0x97, 0x56, 0xa4, 0xde, 0x2d, 0xb5, 0x7b, 0xcd, 0x7d, 0xd2, 0x83, 0xa4,
	0x5e, 0xf3, 0xb2, 0x01, 0xb9, 0xd9, 0x6b, 0x8e, 0x57, 0xe8, 0xe7, 0x3b, 0xca, 0x8d, 0x24, 0x0e,
	0x7f, 0xc5, 0xf7, 0x96, 0x83, 0x53, 0x1c, 0x77, 0x3d, 0x7f, 0x38, 0x7f, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x30, 0x4b, 0xb7, 0x6c, 0x2e, 0x03, 0x00, 0x00,
}
