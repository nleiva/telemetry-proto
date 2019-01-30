// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_retrans.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_areas_area_retransmission_list_table_retransmission

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

// OSPFv3 retransmission list information
type Ospfv3EdmRetrans_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,5,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRetrans_KEYS) Reset()         { *m = Ospfv3EdmRetrans_KEYS{} }
func (m *Ospfv3EdmRetrans_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmRetrans_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_retrans_b0f3269124081bce, []int{0}
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmRetrans_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Merge(dst, src)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Size(m)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRetrans_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRetrans_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmRetrans_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Ospfv3EdmRetrans_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmRetrans struct {
	// Neighbor IP address
	RetransmissionNeighborAddress string `protobuf:"bytes,50,opt,name=retransmission_neighbor_address,json=retransmissionNeighborAddress,proto3" json:"retransmission_neighbor_address,omitempty"`
	// If true, virtual link is retransmitted
	IsRetransmissionVirtualLink bool `protobuf:"varint,51,opt,name=is_retransmission_virtual_link,json=isRetransmissionVirtualLink,proto3" json:"is_retransmission_virtual_link,omitempty"`
	// Retransmission virtual link ID
	RetransmissionVirtualLinkId uint32 `protobuf:"varint,52,opt,name=retransmission_virtual_link_id,json=retransmissionVirtualLinkId,proto3" json:"retransmission_virtual_link_id,omitempty"`
	// If true, sham link is retransmitted
	IsRetransmissionShamLink bool `protobuf:"varint,53,opt,name=is_retransmission_sham_link,json=isRetransmissionShamLink,proto3" json:"is_retransmission_sham_link,omitempty"`
	// Retransmission sham link ID
	RetransmissionShamLinkId uint32 `protobuf:"varint,54,opt,name=retransmission_sham_link_id,json=retransmissionShamLinkId,proto3" json:"retransmission_sham_link_id,omitempty"`
	// Amount of time remaining on retransmission timer (ms)
	RetransmissionTimer uint32 `protobuf:"varint,55,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	// Retransmission queue length
	RetransmissionLength uint32 `protobuf:"varint,56,opt,name=retransmission_length,json=retransmissionLength,proto3" json:"retransmission_length,omitempty"`
	// List of virtual link scope entries
	RetransmissionVirtualLinkDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,57,rep,name=retransmission_virtual_link_db_list,json=retransmissionVirtualLinkDbList,proto3" json:"retransmission_virtual_link_db_list,omitempty"`
	// List of area scope entries
	RetransmissionAreaDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,58,rep,name=retransmission_area_db_list,json=retransmissionAreaDbList,proto3" json:"retransmission_area_db_list,omitempty"`
	// List of AS scope entries
	RetransmissionAsdbList []*Ospfv3EdmLsaSum `protobuf:"bytes,59,rep,name=retransmission_asdb_list,json=retransmissionAsdbList,proto3" json:"retransmission_asdb_list,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *Ospfv3EdmRetrans) Reset()         { *m = Ospfv3EdmRetrans{} }
func (m *Ospfv3EdmRetrans) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans) ProtoMessage()    {}
func (*Ospfv3EdmRetrans) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_retrans_b0f3269124081bce, []int{1}
}
func (m *Ospfv3EdmRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRetrans.Unmarshal(m, b)
}
func (m *Ospfv3EdmRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRetrans.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRetrans.Merge(dst, src)
}
func (m *Ospfv3EdmRetrans) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRetrans.Size(m)
}
func (m *Ospfv3EdmRetrans) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRetrans.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRetrans proto.InternalMessageInfo

func (m *Ospfv3EdmRetrans) GetRetransmissionNeighborAddress() string {
	if m != nil {
		return m.RetransmissionNeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionVirtualLink() bool {
	if m != nil {
		return m.IsRetransmissionVirtualLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkId() uint32 {
	if m != nil {
		return m.RetransmissionVirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionShamLink() bool {
	if m != nil {
		return m.IsRetransmissionShamLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionShamLinkId() uint32 {
	if m != nil {
		return m.RetransmissionShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionTimer() uint32 {
	if m != nil {
		return m.RetransmissionTimer
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionLength() uint32 {
	if m != nil {
		return m.RetransmissionLength
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkDbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionVirtualLinkDbList
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAreaDbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAreaDbList
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAsdbList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAsdbList
	}
	return nil
}

// LSA summary entry
type Ospfv3EdmLsaSum struct {
	// LSA type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	// Age of the LSA (seconds)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsaId string `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId,proto3" json:"header_lsa_id,omitempty"`
	// Router ID of the advertising router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber int32    `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmLsaSum) Reset()         { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()    {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_retrans_b0f3269124081bce, []int{2}
}
func (m *Ospfv3EdmLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Unmarshal(m, b)
}
func (m *Ospfv3EdmLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmLsaSum.Merge(dst, src)
}
func (m *Ospfv3EdmLsaSum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Size(m)
}
func (m *Ospfv3EdmLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmLsaSum proto.InternalMessageInfo

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmRetrans_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans_KEYS")
	proto.RegisterType((*Ospfv3EdmRetrans)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_lsa_sum")
}

func init() {
	proto.RegisterFile("ospfv3_edm_retrans.proto", fileDescriptor_ospfv3_edm_retrans_b0f3269124081bce)
}

var fileDescriptor_ospfv3_edm_retrans_b0f3269124081bce = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xbf, 0x7e, 0xf4, 0x67, 0xda, 0xb4, 0x30, 0x94, 0xd6, 0x55, 0x04, 0x0d, 0xe1, 0x47,
	0x61, 0x63, 0x89, 0xa6, 0x14, 0x28, 0xea, 0x22, 0xe2, 0x47, 0x8a, 0x88, 0xba, 0x70, 0x2b, 0x24,
	0x56, 0xa3, 0x71, 0xe6, 0x26, 0x19, 0x35, 0xfe, 0x61, 0xee, 0xc4, 0xa2, 0x4f, 0xc0, 0x0b, 0xf1,
	0x04, 0xdd, 0xf0, 0x06, 0x3c, 0x0b, 0x4b, 0xe4, 0x6b, 0x27, 0x24, 0x0e, 0xe9, 0x96, 0x6e, 0x2c,
	0xe7, 0x9e, 0x73, 0x4f, 0xce, 0x19, 0x5f, 0x5f, 0x33, 0x37, 0xc6, 0xa4, 0x97, 0x36, 0x05, 0xa8,
	0x50, 0x18, 0xb0, 0x46, 0x46, 0xe8, 0x25, 0x26, 0xb6, 0x31, 0xc7, 0xae, 0xc6, 0x6e, 0x2c, 0x74,
	0x8c, 0xe2, 0xab, 0x11, 0x3a, 0x49, 0x8f, 0x44, 0xc1, 0x8d, 0x13, 0x30, 0x5e, 0x7e, 0x9f, 0x71,
	0xbb, 0x80, 0x08, 0x38, 0xbe, 0xf3, 0x52, 0xd3, 0xa3, 0x8b, 0x27, 0x0d, 0x48, 0xa4, 0xab, 0x57,
	0x28, 0x87, 0x1a, 0x51, 0xc7, 0x91, 0x18, 0x6a, 0xb4, 0xc2, 0xca, 0x60, 0x08, 0x25, 0xa4, 0x7e,
	0xe5, 0xb0, 0xdd, 0x79, 0x47, 0xe2, 0xe3, 0xfb, 0xcf, 0x67, 0xfc, 0x21, 0xdb, 0x28, 0xfe, 0x43,
	0x44, 0x32, 0x04, 0xd7, 0xa9, 0x39, 0x8d, 0x35, 0x7f, 0xbd, 0xa8, 0x9d, 0xca, 0x10, 0xf8, 0x1e,
	0x5b, 0x4d, 0x4d, 0x2f, 0x87, 0xff, 0x23, 0x78, 0x25, 0x35, 0x3d, 0x82, 0x76, 0xd9, 0x4a, 0x66,
	0x45, 0x68, 0xe5, 0x2e, 0xd5, 0x9c, 0x46, 0xc5, 0x5f, 0xce, 0x7e, 0xb6, 0x15, 0x7f, 0xc2, 0x36,
	0x75, 0x64, 0xc1, 0xf4, 0x64, 0x17, 0xf2, 0xce, 0xff, 0xa9, 0xb3, 0x32, 0xa9, 0x52, 0xff, 0x33,
	0x76, 0x3b, 0x02, 0xdd, 0x1f, 0x04, 0xb1, 0x11, 0x52, 0x29, 0x03, 0x88, 0xee, 0x2d, 0x22, 0x6e,
	0x8d, 0xeb, 0xad, 0xbc, 0x5c, 0xff, 0xbe, 0xca, 0xf8, 0x7c, 0x08, 0xfe, 0x81, 0xed, 0x97, 0xce,
	0x61, 0x4e, 0xf0, 0x80, 0x04, 0xef, 0xcf, 0xd2, 0x4e, 0x67, 0xe5, 0xf9, 0x5b, 0xf6, 0x40, 0xa3,
	0x28, 0x49, 0xa5, 0xda, 0xd8, 0x91, 0x1c, 0x8a, 0xa1, 0x8e, 0x2e, 0xdc, 0x66, 0xcd, 0x69, 0xac,
	0xfa, 0x55, 0x8d, 0xfe, 0x0c, 0xe9, 0x53, 0xce, 0xe9, 0xe8, 0xe8, 0x22, 0x13, 0xb9, 0x46, 0x21,
	0x3b, 0xa5, 0x43, 0x3a, 0xa5, 0xaa, 0x59, 0x24, 0xd1, 0x56, 0xfc, 0x84, 0x55, 0xe7, 0x9d, 0xe0,
	0x40, 0x86, 0xb9, 0x8d, 0x17, 0x64, 0xc3, 0x2d, 0xdb, 0x38, 0x1b, 0xc8, 0x90, 0x3c, 0x9c, 0xb0,
	0xea, 0xa2, 0xde, 0xcc, 0xc0, 0x11, 0x19, 0x70, 0xcd, 0x5f, 0x9b, 0xdb, 0x8a, 0x3f, 0x67, 0xdb,
	0xa5, 0x76, 0xab, 0x43, 0x30, 0xee, 0x4b, 0xea, 0xbb, 0x3b, 0x8b, 0x9d, 0x67, 0x10, 0x6f, 0xb2,
	0x7b, 0xe5, 0x51, 0x84, 0xa8, 0x6f, 0x07, 0xee, 0x2b, 0xea, 0x29, 0xe9, 0x75, 0x08, 0xe3, 0x3f,
	0x1d, 0xf6, 0xe8, 0xba, 0xb3, 0x52, 0x01, 0x0d, 0xb4, 0xfb, 0xba, 0xb6, 0xd4, 0x58, 0x3f, 0xf8,
	0xe6, 0x78, 0xff, 0xe0, 0xc5, 0xf1, 0xa6, 0xe6, 0x6d, 0x88, 0x52, 0xe0, 0x28, 0xf4, 0xf7, 0x17,
	0x3e, 0xba, 0x77, 0x41, 0x47, 0xa3, 0xe5, 0x3f, 0x9c, 0xb9, 0x07, 0x40, 0xaf, 0xc8, 0x38, 0xd0,
	0xf1, 0x0d, 0x0b, 0x54, 0x1a, 0x85, 0x96, 0x01, 0x59, 0x24, 0xb9, 0x72, 0x98, 0x5b, 0x4e, 0x82,
	0xe3, 0x18, 0x6f, 0x6e, 0x58, 0x8c, 0x9d, 0x52, 0x0c, 0x54, 0x14, 0xa2, 0xfe, 0xcb, 0x99, 0x59,
	0x1b, 0x05, 0x9d, 0x3f, 0x65, 0x5b, 0x03, 0x90, 0x0a, 0x0c, 0x55, 0xec, 0x65, 0x32, 0xde, 0x7c,
	0x95, 0xbc, 0xdc, 0x41, 0x79, 0x7e, 0x99, 0x00, 0x7f, 0xcc, 0x36, 0xa7, 0x78, 0xb2, 0x9f, 0x6f,
	0xc0, 0x8a, 0xbf, 0x31, 0xa1, 0xb5, 0xfa, 0xc0, 0xeb, 0xac, 0x32, 0xc5, 0x2a, 0x96, 0xe1, 0x9a,
	0xbf, 0x3e, 0x21, 0xb5, 0x15, 0x3f, 0x66, 0x7b, 0x05, 0x47, 0xaa, 0x14, 0x8c, 0xd5, 0xa8, 0xa3,
	0xbe, 0x30, 0xf1, 0xc8, 0x82, 0x29, 0x96, 0xe3, 0x6e, 0x4e, 0x68, 0xfd, 0xc1, 0x7d, 0x82, 0xf9,
	0x21, 0xdb, 0x29, 0x7a, 0x11, 0xbe, 0x8c, 0x20, 0xca, 0x76, 0xea, 0x28, 0x0c, 0xc0, 0xd0, 0xb2,
	0xbc, 0xe3, 0x6f, 0xe7, 0xe8, 0x59, 0x01, 0x9e, 0x12, 0x16, 0x2c, 0xd3, 0x27, 0xa7, 0xf9, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x83, 0x44, 0x1c, 0xf2, 0x8e, 0x06, 0x00, 0x00,
}
