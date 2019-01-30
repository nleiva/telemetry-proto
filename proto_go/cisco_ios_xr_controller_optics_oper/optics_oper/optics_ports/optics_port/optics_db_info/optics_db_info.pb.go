// Code generated by protoc-gen-go. DO NOT EDIT.
// source: optics_db_info.proto

package cisco_ios_xr_controller_optics_oper_optics_oper_optics_ports_optics_port_optics_db_info

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

//  Optics DB Info
type OpticsDbInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpticsDbInfo_KEYS) Reset()         { *m = OpticsDbInfo_KEYS{} }
func (m *OpticsDbInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OpticsDbInfo_KEYS) ProtoMessage()    {}
func (*OpticsDbInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_optics_db_info_32eccb062a78f912, []int{0}
}
func (m *OpticsDbInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsDbInfo_KEYS.Unmarshal(m, b)
}
func (m *OpticsDbInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsDbInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *OpticsDbInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsDbInfo_KEYS.Merge(dst, src)
}
func (m *OpticsDbInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OpticsDbInfo_KEYS.Size(m)
}
func (m *OpticsDbInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsDbInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsDbInfo_KEYS proto.InternalMessageInfo

func (m *OpticsDbInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type OpticsDbInfo struct {
	// Transport Admin State
	TransportAdminState string `protobuf:"bytes,50,opt,name=transport_admin_state,json=transportAdminState,proto3" json:"transport_admin_state,omitempty"`
	// Optics controller state: Up, Down or Administratively Down
	ControllerState string `protobuf:"bytes,51,opt,name=controller_state,json=controllerState,proto3" json:"controller_state,omitempty"`
	// Network SRLG information
	NetworkSrlgInfo      *OpticsEdmNetworkSrlgInfo `protobuf:"bytes,52,opt,name=network_srlg_info,json=networkSrlgInfo,proto3" json:"network_srlg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OpticsDbInfo) Reset()         { *m = OpticsDbInfo{} }
func (m *OpticsDbInfo) String() string { return proto.CompactTextString(m) }
func (*OpticsDbInfo) ProtoMessage()    {}
func (*OpticsDbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_optics_db_info_32eccb062a78f912, []int{1}
}
func (m *OpticsDbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsDbInfo.Unmarshal(m, b)
}
func (m *OpticsDbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsDbInfo.Marshal(b, m, deterministic)
}
func (dst *OpticsDbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsDbInfo.Merge(dst, src)
}
func (m *OpticsDbInfo) XXX_Size() int {
	return xxx_messageInfo_OpticsDbInfo.Size(m)
}
func (m *OpticsDbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsDbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsDbInfo proto.InternalMessageInfo

func (m *OpticsDbInfo) GetTransportAdminState() string {
	if m != nil {
		return m.TransportAdminState
	}
	return ""
}

func (m *OpticsDbInfo) GetControllerState() string {
	if m != nil {
		return m.ControllerState
	}
	return ""
}

func (m *OpticsDbInfo) GetNetworkSrlgInfo() *OpticsEdmNetworkSrlgInfo {
	if m != nil {
		return m.NetworkSrlgInfo
	}
	return nil
}

// Network SRLG Information
type OpticsEdmNetworkSrlgStructure struct {
	// Array to maintain set number
	SetNumber uint32 `protobuf:"varint,1,opt,name=set_number,json=setNumber,proto3" json:"set_number,omitempty"`
	// Network Srlg
	NetworkSrlg          []uint32 `protobuf:"varint,2,rep,packed,name=network_srlg,json=networkSrlg,proto3" json:"network_srlg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpticsEdmNetworkSrlgStructure) Reset()         { *m = OpticsEdmNetworkSrlgStructure{} }
func (m *OpticsEdmNetworkSrlgStructure) String() string { return proto.CompactTextString(m) }
func (*OpticsEdmNetworkSrlgStructure) ProtoMessage()    {}
func (*OpticsEdmNetworkSrlgStructure) Descriptor() ([]byte, []int) {
	return fileDescriptor_optics_db_info_32eccb062a78f912, []int{2}
}
func (m *OpticsEdmNetworkSrlgStructure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsEdmNetworkSrlgStructure.Unmarshal(m, b)
}
func (m *OpticsEdmNetworkSrlgStructure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsEdmNetworkSrlgStructure.Marshal(b, m, deterministic)
}
func (dst *OpticsEdmNetworkSrlgStructure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsEdmNetworkSrlgStructure.Merge(dst, src)
}
func (m *OpticsEdmNetworkSrlgStructure) XXX_Size() int {
	return xxx_messageInfo_OpticsEdmNetworkSrlgStructure.Size(m)
}
func (m *OpticsEdmNetworkSrlgStructure) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsEdmNetworkSrlgStructure.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsEdmNetworkSrlgStructure proto.InternalMessageInfo

func (m *OpticsEdmNetworkSrlgStructure) GetSetNumber() uint32 {
	if m != nil {
		return m.SetNumber
	}
	return 0
}

func (m *OpticsEdmNetworkSrlgStructure) GetNetworkSrlg() []uint32 {
	if m != nil {
		return m.NetworkSrlg
	}
	return nil
}

// Network SRLG Information Array
type OpticsEdmNetworkSrlgInfo struct {
	// Network Srlg Array
	NetworkSrlgArray     []*OpticsEdmNetworkSrlgStructure `protobuf:"bytes,1,rep,name=network_srlg_array,json=networkSrlgArray,proto3" json:"network_srlg_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *OpticsEdmNetworkSrlgInfo) Reset()         { *m = OpticsEdmNetworkSrlgInfo{} }
func (m *OpticsEdmNetworkSrlgInfo) String() string { return proto.CompactTextString(m) }
func (*OpticsEdmNetworkSrlgInfo) ProtoMessage()    {}
func (*OpticsEdmNetworkSrlgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_optics_db_info_32eccb062a78f912, []int{3}
}
func (m *OpticsEdmNetworkSrlgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpticsEdmNetworkSrlgInfo.Unmarshal(m, b)
}
func (m *OpticsEdmNetworkSrlgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpticsEdmNetworkSrlgInfo.Marshal(b, m, deterministic)
}
func (dst *OpticsEdmNetworkSrlgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpticsEdmNetworkSrlgInfo.Merge(dst, src)
}
func (m *OpticsEdmNetworkSrlgInfo) XXX_Size() int {
	return xxx_messageInfo_OpticsEdmNetworkSrlgInfo.Size(m)
}
func (m *OpticsEdmNetworkSrlgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpticsEdmNetworkSrlgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpticsEdmNetworkSrlgInfo proto.InternalMessageInfo

func (m *OpticsEdmNetworkSrlgInfo) GetNetworkSrlgArray() []*OpticsEdmNetworkSrlgStructure {
	if m != nil {
		return m.NetworkSrlgArray
	}
	return nil
}

func init() {
	proto.RegisterType((*OpticsDbInfo_KEYS)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_db_info.optics_db_info_KEYS")
	proto.RegisterType((*OpticsDbInfo)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_db_info.optics_db_info")
	proto.RegisterType((*OpticsEdmNetworkSrlgStructure)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_db_info.optics_edm_network_srlg_structure")
	proto.RegisterType((*OpticsEdmNetworkSrlgInfo)(nil), "cisco_ios_xr_controller_optics_oper.optics_oper.optics_ports.optics_port.optics_db_info.optics_edm_network_srlg_info")
}

func init() {
	proto.RegisterFile("optics_db_info.proto", fileDescriptor_optics_db_info_32eccb062a78f912)
}

var fileDescriptor_optics_db_info_32eccb062a78f912 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x3d, 0x4e, 0xf3, 0x40,
	0x10, 0xd5, 0x26, 0x9f, 0x3e, 0x29, 0x13, 0x42, 0xc2, 0x06, 0x24, 0x17, 0x20, 0x39, 0xa9, 0x9c,
	0xc6, 0x85, 0xc3, 0x05, 0x52, 0x50, 0x20, 0x24, 0x0a, 0xa7, 0x40, 0x54, 0xab, 0x8d, 0xbd, 0x44,
	0x16, 0xf6, 0xae, 0x35, 0x3b, 0x16, 0x3f, 0xd7, 0xe0, 0x04, 0x74, 0x5c, 0x85, 0x5b, 0x21, 0x6f,
	0x42, 0xb0, 0xc5, 0x4f, 0x07, 0xdd, 0xce, 0x9b, 0x37, 0xf3, 0xde, 0x3c, 0x2d, 0x1c, 0x9a, 0x92,
	0xb2, 0xc4, 0x8a, 0x74, 0x25, 0x32, 0x7d, 0x63, 0xc2, 0x12, 0x0d, 0x19, 0x7e, 0x95, 0x64, 0x36,
	0x31, 0x22, 0x33, 0x56, 0xdc, 0xa3, 0x48, 0x8c, 0x26, 0x34, 0x79, 0xae, 0x50, 0x6c, 0xd9, 0xa6,
	0x54, 0x18, 0x7e, 0xf1, 0x2e, 0x0d, 0x92, 0x6d, 0x16, 0x61, 0x7b, 0xfd, 0x74, 0x06, 0xe3, 0x36,
	0x22, 0x2e, 0xce, 0xae, 0x97, 0x9c, 0xc3, 0x3f, 0x2d, 0x0b, 0xe5, 0x31, 0x9f, 0x05, 0xbd, 0xd8,
	0xbd, 0xa7, 0x4f, 0x1d, 0xd8, 0x6f, 0x73, 0x79, 0x04, 0x47, 0x84, 0x52, 0xdb, 0x7a, 0xb3, 0x90,
	0x69, 0x91, 0x69, 0x61, 0x49, 0x92, 0xf2, 0x22, 0x37, 0x37, 0xde, 0x35, 0x17, 0x75, 0x6f, 0x59,
	0xb7, 0xf8, 0x0c, 0x46, 0x0d, 0xff, 0x1b, 0xfa, 0xdc, 0xd1, 0x87, 0x1f, 0xf8, 0x86, 0xfa, 0xcc,
	0xe0, 0x40, 0x2b, 0xba, 0x33, 0x78, 0x2b, 0x2c, 0xe6, 0x6b, 0x27, 0xea, 0x9d, 0xfa, 0x2c, 0xe8,
	0x47, 0x55, 0xf8, 0x4b, 0x91, 0xbc, 0x97, 0x2a, 0x2d, 0xc4, 0x27, 0xf1, 0x78, 0xb8, 0x85, 0x96,
	0x98, 0xaf, 0xcf, 0xeb, 0x00, 0x15, 0x4c, 0xbe, 0x1b, 0xb0, 0x84, 0x55, 0x42, 0x15, 0x2a, 0x7e,
	0x02, 0x60, 0x15, 0x09, 0x5d, 0x15, 0x2b, 0x85, 0x2e, 0xd4, 0x41, 0xdc, 0xb3, 0x8a, 0x2e, 0x1d,
	0xc0, 0x27, 0xb0, 0xd7, 0x1c, 0xf4, 0x3a, 0x7e, 0x37, 0x18, 0xc4, 0xfd, 0x86, 0xd4, 0xf4, 0x95,
	0xc1, 0xf1, 0x4f, 0xc6, 0xf8, 0x0b, 0x03, 0xde, 0x42, 0x25, 0xa2, 0x7c, 0xf0, 0x98, 0xdf, 0x0d,
	0xfa, 0xd1, 0xe3, 0x9f, 0x87, 0xb5, 0xbb, 0x3d, 0x1e, 0x35, 0xce, 0x58, 0xd4, 0x9e, 0x56, 0xff,
	0xdd, 0x9f, 0x9e, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x18, 0x67, 0xf5, 0x59, 0xeb, 0x02, 0x00,
	0x00,
}
