// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_retrans_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_adjacency_information_retransmissions_retransmission

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

// OSPF Retransmission List
type OspfShRetransList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRetransList_KEYS) Reset()         { *m = OspfShRetransList_KEYS{} }
func (m *OspfShRetransList_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRetransList_KEYS) ProtoMessage()    {}
func (*OspfShRetransList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_513be076a78114a4, []int{0}
}

func (m *OspfShRetransList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList_KEYS.Unmarshal(m, b)
}
func (m *OspfShRetransList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRetransList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList_KEYS.Merge(m, src)
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

func (m *OspfShRetransList_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	return fileDescriptor_513be076a78114a4, []int{1}
}

func (m *OspfShRetransList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList.Unmarshal(m, b)
}
func (m *OspfShRetransList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList.Marshal(b, m, deterministic)
}
func (m *OspfShRetransList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList.Merge(m, src)
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
	return fileDescriptor_513be076a78114a4, []int{2}
}

func (m *OspfShLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShLsaSum.Unmarshal(m, b)
}
func (m *OspfShLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShLsaSum.Marshal(b, m, deterministic)
}
func (m *OspfShLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShLsaSum.Merge(m, src)
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
	proto.RegisterType((*OspfShRetransList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list_KEYS")
	proto.RegisterType((*OspfShRetransList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum")
}

func init() { proto.RegisterFile("ospf_sh_retrans_list.proto", fileDescriptor_513be076a78114a4) }

var fileDescriptor_513be076a78114a4 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0x96, 0x16, 0x98, 0x36, 0x0d, 0x6c, 0xa3, 0xe2, 0x14, 0x21, 0x42, 0x04, 0x28,
	0x5c, 0x2c, 0xd1, 0x86, 0x0f, 0x21, 0x2e, 0xa1, 0x70, 0x88, 0x5a, 0xe5, 0x90, 0xf6, 0xc2, 0x69,
	0xb5, 0xb1, 0xc7, 0xc9, 0x42, 0xbd, 0x6b, 0x76, 0xec, 0x88, 0x1c, 0x79, 0x16, 0x8e, 0xbc, 0x02,
	0x17, 0x5e, 0x80, 0x67, 0x42, 0x5e, 0x7f, 0x50, 0x5b, 0xe9, 0x15, 0x71, 0x89, 0x36, 0xf3, 0xff,
	0xcd, 0xec, 0x7f, 0x76, 0x77, 0x0c, 0x87, 0x9a, 0xe2, 0x90, 0xd3, 0x82, 0x1b, 0x4c, 0x8c, 0x50,
	0xc4, 0x2f, 0x25, 0x25, 0x5e, 0x6c, 0x74, 0xa2, 0x59, 0xec, 0x4b, 0xf2, 0x35, 0x97, 0x9a, 0xf8,
	0x57, 0xc3, 0x65, 0xbc, 0x1c, 0x72, 0x4b, 0xeb, 0x18, 0x8d, 0x97, 0xad, 0x32, 0xce, 0x47, 0x22,
	0xa4, 0x72, 0xe5, 0x2d, 0x4d, 0x68, 0x7f, 0x3c, 0x11, 0x7c, 0x12, 0x3e, 0x2a, 0x7f, 0xc5, 0xa5,
	0x0a, 0xb5, 0x89, 0x44, 0x22, 0xb5, 0xf2, 0x8a, 0x5d, 0x22, 0x49, 0x24, 0xb5, 0xa2, 0xc6, 0xff,
	0xfe, 0x0f, 0x07, 0xba, 0xeb, 0x0c, 0xf1, 0xd3, 0x0f, 0x1f, 0xcf, 0xd9, 0x23, 0xd8, 0x2d, 0xb6,
	0xe1, 0x4a, 0x44, 0xe8, 0x3a, 0x3d, 0x67, 0x70, 0x7b, 0xba, 0x53, 0xc4, 0x26, 0x22, 0x42, 0xd6,
	0x85, 0x5b, 0x4b, 0x13, 0xe6, 0xf2, 0x86, 0x95, 0x6f, 0x2e, 0x4d, 0x68, 0xa5, 0x27, 0xb0, 0x27,
	0x55, 0x82, 0x26, 0x14, 0x3e, 0xe6, 0xc0, 0xa6, 0x05, 0x5a, 0x55, 0xd4, 0x62, 0xcf, 0xe0, 0x8e,
	0x42, 0x39, 0x5f, 0xcc, 0xb4, 0xe1, 0x22, 0x08, 0x0c, 0x12, 0xb9, 0x37, 0x2c, 0xd8, 0x2e, 0xe3,
	0xa3, 0x3c, 0xdc, 0xff, 0xbd, 0x05, 0x9d, 0x75, 0x6e, 0xd9, 0x5b, 0x38, 0xac, 0x37, 0xc6, 0xab,
	0x92, 0x32, 0x70, 0x8f, 0x6c, 0x35, 0xb7, 0x4e, 0x4c, 0x0a, 0x60, 0x1c, 0xb0, 0x53, 0xe8, 0x5f,
	0x9b, 0x1d, 0x57, 0x9e, 0x8e, 0x6d, 0x95, 0x87, 0xd7, 0x54, 0x89, 0x0b, 0x8f, 0xec, 0x1d, 0x3c,
	0x68, 0x14, 0x6b, 0x1c, 0xc2, 0xd0, 0xd6, 0xb9, 0x5f, 0x87, 0xc6, 0xb5, 0x23, 0x79, 0x0e, 0x9d,
	0x46, 0x8d, 0x44, 0x46, 0x68, 0xdc, 0x17, 0x3d, 0x67, 0xd0, 0x9a, 0xee, 0xd7, 0xb5, 0x8b, 0x4c,
	0x5a, 0x93, 0xe2, 0xeb, 0x54, 0x25, 0xee, 0xcb, 0x75, 0x29, 0x27, 0x99, 0xc4, 0x7e, 0x39, 0xd0,
	0x70, 0xc1, 0x85, 0x41, 0xc1, 0x83, 0x99, 0x3d, 0x54, 0xf7, 0x55, 0x6f, 0x73, 0xb0, 0x73, 0xf4,
	0xcd, 0xf1, 0xfe, 0xf5, 0xab, 0xf4, 0xca, 0x3b, 0xbe, 0x24, 0xc1, 0x29, 0x8d, 0x9a, 0x57, 0x37,
	0x32, 0x28, 0xde, 0xcf, 0xce, 0xb2, 0x8b, 0xff, 0xe9, 0x80, 0xdb, 0xec, 0x81, 0xca, 0x06, 0x5e,
	0xff, 0x37, 0x0d, 0x1c, 0x34, 0x1a, 0xa0, 0xc0, 0xda, 0xef, 0x7f, 0xdf, 0x80, 0x76, 0x83, 0x65,
	0x4f, 0xa1, 0xbd, 0x40, 0x11, 0xa0, 0xb1, 0x91, 0x64, 0x15, 0x97, 0x73, 0xd7, 0xca, 0xc3, 0x67,
	0x24, 0x2e, 0x56, 0x31, 0xb2, 0xc7, 0xb0, 0x77, 0x85, 0x13, 0xf3, 0x7c, 0xfe, 0x5a, 0xd3, 0xdd,
	0x0a, 0x1b, 0xcd, 0x91, 0xf5, 0x60, 0xb7, 0xa2, 0xb2, 0x59, 0xc8, 0x47, 0x10, 0x4a, 0x66, 0x1c,
	0xb0, 0x37, 0xd0, 0x2d, 0x08, 0x11, 0x2c, 0xd1, 0x24, 0x92, 0xa4, 0x9a, 0x73, 0xa3, 0xd3, 0x04,
	0x4d, 0x31, 0x88, 0xf7, 0x72, 0x60, 0xf4, 0x57, 0x9f, 0x5a, 0x99, 0x0d, 0xe1, 0xa0, 0xc8, 0x25,
	0xfc, 0x92, 0xa2, 0xca, 0xde, 0x78, 0x1a, 0xcd, 0xd0, 0xb8, 0x5b, 0xd6, 0x4b, 0x27, 0x57, 0xcf,
	0x0b, 0x71, 0x62, 0x35, 0xe6, 0xc1, 0xfe, 0x15, 0xe7, 0xfe, 0x02, 0xfd, 0xcf, 0x94, 0x46, 0xee,
	0xb6, 0x4d, 0xb9, 0x5b, 0xd9, 0x3f, 0x29, 0x84, 0xd9, 0xb6, 0xfd, 0x3a, 0x1e, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x24, 0x0b, 0x72, 0x22, 0x3b, 0x05, 0x00, 0x00,
}
