// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_retrans_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission

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

// OSPF Retransmission List
type OspfShRetransList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRetransList_KEYS) Reset()         { *m = OspfShRetransList_KEYS{} }
func (m *OspfShRetransList_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRetransList_KEYS) ProtoMessage()    {}
func (*OspfShRetransList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_retrans_list_f6f4a34591d09d7d, []int{0}
}
func (m *OspfShRetransList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList_KEYS.Unmarshal(m, b)
}
func (m *OspfShRetransList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShRetransList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList_KEYS.Merge(dst, src)
}
func (m *OspfShRetransList_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRetransList_KEYS.Size(m)
}
func (m *OspfShRetransList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRetransList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRetransList_KEYS proto.InternalMessageInfo

func (m *OspfShRetransList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRetransList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShRetransList_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShRetransList struct {
	// Neighbor ID
	RetransmissionNeighborId string `protobuf:"bytes,50,opt,name=retransmission_neighbor_id,json=retransmissionNeighborId,proto3" json:"retransmission_neighbor_id,omitempty"`
	// Neighbor IP Address
	RetransmissionNeighborIpAddress string `protobuf:"bytes,51,opt,name=retransmission_neighbor_ip_address,json=retransmissionNeighborIpAddress,proto3" json:"retransmission_neighbor_ip_address,omitempty"`
	// Retransmission list interface
	RetransmissionInterfaceName string `protobuf:"bytes,52,opt,name=retransmission_interface_name,json=retransmissionInterfaceName,proto3" json:"retransmission_interface_name,omitempty"`
	// Amount of time remaining on retransmission timer (ms)
	RetransmissionTimer uint32 `protobuf:"varint,53,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	// Retransmission queue length
	RetransmissionCount uint32 `protobuf:"varint,54,opt,name=retransmission_count,json=retransmissionCount,proto3" json:"retransmission_count,omitempty"`
	// List of Area scope entries
	RetransmissionAreaDbList []*OspfShLsaSum `protobuf:"bytes,55,rep,name=retransmission_area_db_list,json=retransmissionAreaDbList,proto3" json:"retransmission_area_db_list,omitempty"`
	// List of AS Scope entries
	RetransmissionAsdbList []*OspfShLsaSum `protobuf:"bytes,56,rep,name=retransmission_asdb_list,json=retransmissionAsdbList,proto3" json:"retransmission_asdb_list,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *OspfShRetransList) Reset()         { *m = OspfShRetransList{} }
func (m *OspfShRetransList) String() string { return proto.CompactTextString(m) }
func (*OspfShRetransList) ProtoMessage()    {}
func (*OspfShRetransList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_retrans_list_f6f4a34591d09d7d, []int{1}
}
func (m *OspfShRetransList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList.Unmarshal(m, b)
}
func (m *OspfShRetransList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList.Marshal(b, m, deterministic)
}
func (dst *OspfShRetransList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList.Merge(dst, src)
}
func (m *OspfShRetransList) XXX_Size() int {
	return xxx_messageInfo_OspfShRetransList.Size(m)
}
func (m *OspfShRetransList) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRetransList.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRetransList proto.InternalMessageInfo

func (m *OspfShRetransList) GetRetransmissionNeighborId() string {
	if m != nil {
		return m.RetransmissionNeighborId
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionNeighborIpAddress() string {
	if m != nil {
		return m.RetransmissionNeighborIpAddress
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionInterfaceName() string {
	if m != nil {
		return m.RetransmissionInterfaceName
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionTimer() uint32 {
	if m != nil {
		return m.RetransmissionTimer
	}
	return 0
}

func (m *OspfShRetransList) GetRetransmissionCount() uint32 {
	if m != nil {
		return m.RetransmissionCount
	}
	return 0
}

func (m *OspfShRetransList) GetRetransmissionAreaDbList() []*OspfShLsaSum {
	if m != nil {
		return m.RetransmissionAreaDbList
	}
	return nil
}

func (m *OspfShRetransList) GetRetransmissionAsdbList() []*OspfShLsaSum {
	if m != nil {
		return m.RetransmissionAsdbList
	}
	return nil
}

// LSA Summary Entry
type OspfShLsaSum struct {
	// LSA Type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	// Age of the LSA (s)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsId string `protobuf:"bytes,3,opt,name=header_ls_id,json=headerLsId,proto3" json:"header_ls_id,omitempty"`
	// Router ID of the Advertising Router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber uint32 `protobuf:"varint,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	// Checksum of the LSA
	HeaderLsaChecksum    uint32   `protobuf:"varint,6,opt,name=header_lsa_checksum,json=headerLsaChecksum,proto3" json:"header_lsa_checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShLsaSum) Reset()         { *m = OspfShLsaSum{} }
func (m *OspfShLsaSum) String() string { return proto.CompactTextString(m) }
func (*OspfShLsaSum) ProtoMessage()    {}
func (*OspfShLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_retrans_list_f6f4a34591d09d7d, []int{2}
}
func (m *OspfShLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShLsaSum.Unmarshal(m, b)
}
func (m *OspfShLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShLsaSum.Marshal(b, m, deterministic)
}
func (dst *OspfShLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShLsaSum.Merge(dst, src)
}
func (m *OspfShLsaSum) XXX_Size() int {
	return xxx_messageInfo_OspfShLsaSum.Size(m)
}
func (m *OspfShLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShLsaSum proto.InternalMessageInfo

func (m *OspfShLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsId() string {
	if m != nil {
		return m.HeaderLsId
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderSequenceNumber() uint32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsaChecksum() uint32 {
	if m != nil {
		return m.HeaderLsaChecksum
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShRetransList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list_KEYS")
	proto.RegisterType((*OspfShRetransList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum")
}

func init() {
	proto.RegisterFile("ospf_sh_retrans_list.proto", fileDescriptor_ospf_sh_retrans_list_f6f4a34591d09d7d)
}

var fileDescriptor_ospf_sh_retrans_list_f6f4a34591d09d7d = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xb5, 0x2d, 0xad, 0x84, 0x9b, 0x34, 0xe0, 0x46, 0xc5, 0x2d, 0x42, 0x84, 0x08, 0x50,
	0xb8, 0xac, 0x44, 0x1b, 0x3e, 0x84, 0xb8, 0x84, 0xc2, 0x21, 0x6a, 0x95, 0x43, 0xda, 0x0b, 0x27,
	0xcb, 0x59, 0x4f, 0x12, 0x43, 0xd6, 0x5e, 0x3c, 0xde, 0x88, 0xdc, 0x79, 0x00, 0xde, 0x81, 0x07,
	0xe1, 0xc2, 0x23, 0xf0, 0x40, 0x68, 0xbd, 0x1f, 0x34, 0xab, 0xf4, 0x0c, 0x37, 0xef, 0xfc, 0x7f,
	0x33, 0xfe, 0x8f, 0xed, 0x59, 0x72, 0x6c, 0x30, 0x99, 0x72, 0x9c, 0x73, 0x0b, 0xce, 0x0a, 0x8d,
	0x7c, 0xa1, 0xd0, 0x85, 0x89, 0x35, 0xce, 0x50, 0x8c, 0x14, 0x46, 0x86, 0x2b, 0x83, 0xfc, 0xab,
	0xe5, 0x2a, 0x59, 0xf6, 0xb9, 0xa7, 0x4d, 0x02, 0x36, 0xcc, 0x56, 0x19, 0x17, 0x01, 0x22, 0x60,
	0xb9, 0x0a, 0x25, 0x4c, 0x45, 0xba, 0x70, 0x7c, 0x69, 0xa7, 0xa1, 0x90, 0x9f, 0x44, 0x04, 0x3a,
	0x5a, 0x71, 0xa5, 0xa7, 0xc6, 0xc6, 0xc2, 0x29, 0xa3, 0xc3, 0x62, 0xa3, 0x58, 0x21, 0x2a, 0xa3,
	0xb1, 0xf6, 0xdd, 0xfd, 0x1e, 0x90, 0xa3, 0x4d, 0x9e, 0xf8, 0xf9, 0x87, 0x8f, 0x97, 0xf4, 0x11,
	0x69, 0x14, 0x3b, 0x71, 0x2d, 0x62, 0x60, 0x41, 0x27, 0xe8, 0xdd, 0x1e, 0xef, 0x15, 0xb1, 0x91,
	0x88, 0x81, 0x3e, 0x21, 0xfb, 0x4a, 0x3b, 0xb0, 0x53, 0x11, 0x41, 0x0e, 0x6d, 0x79, 0xa8, 0x59,
	0x45, 0x3d, 0xf6, 0x8c, 0xdc, 0xd1, 0xa0, 0x66, 0xf3, 0x89, 0xb1, 0x5c, 0x48, 0x69, 0x01, 0x91,
	0x6d, 0x7b, 0xb0, 0x55, 0xc6, 0x07, 0x79, 0xb8, 0xfb, 0x7b, 0x87, 0xb4, 0x37, 0x59, 0xa2, 0x6f,
	0xc9, 0xf1, 0xba, 0x7b, 0x5e, 0x95, 0x54, 0x92, 0x9d, 0xf8, 0x6a, 0x6c, 0x9d, 0x18, 0x15, 0xc0,
	0x50, 0xd2, 0x73, 0xd2, 0xbd, 0x31, 0x3b, 0xa9, 0x3c, 0x9d, 0xfa, 0x2a, 0x0f, 0x6f, 0xa8, 0x92,
	0x14, 0x1e, 0xe9, 0x3b, 0xf2, 0xa0, 0x56, 0xac, 0x76, 0x08, 0x7d, 0x5f, 0xe7, 0xfe, 0x3a, 0x34,
	0x5c, 0x3b, 0x92, 0xe7, 0xa4, 0x5d, 0xab, 0xe1, 0x54, 0x0c, 0x96, 0xbd, 0xe8, 0x04, 0xbd, 0xe6,
	0xf8, 0x60, 0x5d, 0xbb, 0xca, 0xa4, 0x0d, 0x29, 0x91, 0x49, 0xb5, 0x63, 0x2f, 0x37, 0xa5, 0x9c,
	0x65, 0x12, 0xfd, 0x15, 0x90, 0x9a, 0x0b, 0x2e, 0x2c, 0x08, 0x2e, 0x27, 0xfe, 0x50, 0xd9, 0xab,
	0xce, 0x76, 0x6f, 0xef, 0xe4, 0x5b, 0x10, 0xfe, 0x83, 0xd7, 0x17, 0x96, 0xd7, 0xbc, 0x40, 0xc1,
	0x31, 0x8d, 0xeb, 0xb7, 0x37, 0xb0, 0x20, 0xde, 0x4f, 0x2e, 0xb2, 0xbb, 0xff, 0x19, 0x10, 0x56,
	0x6f, 0x03, 0xcb, 0x1e, 0x5e, 0xff, 0x4f, 0x3d, 0x1c, 0xd6, 0x7a, 0x40, 0xe9, 0x3b, 0xe8, 0xfe,
	0xd8, 0x22, 0xad, 0x1a, 0x4b, 0x9f, 0x92, 0xd6, 0x1c, 0x84, 0x04, 0xeb, 0x23, 0x6e, 0x95, 0x94,
	0x23, 0xd6, 0xcc, 0xc3, 0x17, 0x28, 0xae, 0x56, 0x09, 0xd0, 0xc7, 0x64, 0xff, 0x1a, 0x27, 0x66,
	0xf9, 0x90, 0x35, 0xc7, 0x8d, 0x0a, 0x1b, 0xcc, 0x80, 0x76, 0x48, 0xa3, 0xa2, 0xb2, 0x89, 0xc8,
	0xe7, 0x8b, 0x94, 0xcc, 0x50, 0xd2, 0x37, 0xe4, 0xa8, 0x20, 0x84, 0x5c, 0x82, 0x75, 0x0a, 0x95,
	0x9e, 0x71, 0x6b, 0x52, 0x07, 0x96, 0xdd, 0xf2, 0xf8, 0xbd, 0x1c, 0x18, 0xfc, 0xd5, 0xc7, 0x5e,
	0xa6, 0x7d, 0x72, 0x58, 0xe4, 0x22, 0x7c, 0x49, 0x41, 0x67, 0x2f, 0x3d, 0x8d, 0x27, 0x60, 0xd9,
	0x8e, 0xf7, 0xd2, 0xce, 0xd5, 0xcb, 0x42, 0x1c, 0x79, 0x8d, 0x86, 0xe4, 0xe0, 0x9a, 0xf3, 0x68,
	0x0e, 0xd1, 0x67, 0x4c, 0x63, 0xb6, 0xeb, 0x53, 0xee, 0x56, 0xf6, 0xcf, 0x0a, 0x61, 0xb2, 0xeb,
	0xff, 0x85, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xec, 0x5a, 0xe4, 0x97, 0x29, 0x05, 0x00,
	0x00,
}
