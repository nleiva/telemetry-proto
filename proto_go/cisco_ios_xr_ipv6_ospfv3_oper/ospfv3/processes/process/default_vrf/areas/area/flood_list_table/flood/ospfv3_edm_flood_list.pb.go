// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_flood_list.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_flood_list_table_flood

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

// OSPFv3 flood list information
type Ospfv3EdmFloodList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmFloodList_KEYS) Reset()         { *m = Ospfv3EdmFloodList_KEYS{} }
func (m *Ospfv3EdmFloodList_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmFloodList_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmFloodList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_flood_list_7d7849ea8f2bfbc4, []int{0}
}
func (m *Ospfv3EdmFloodList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmFloodList_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmFloodList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmFloodList_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmFloodList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmFloodList_KEYS.Merge(dst, src)
}
func (m *Ospfv3EdmFloodList_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmFloodList_KEYS.Size(m)
}
func (m *Ospfv3EdmFloodList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmFloodList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmFloodList_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmFloodList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmFloodList_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Ospfv3EdmFloodList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmFloodList struct {
	// Time until next LS transmission (ms)
	LsTransmissionTimer uint32 `protobuf:"varint,50,opt,name=ls_transmission_timer,json=lsTransmissionTimer,proto3" json:"ls_transmission_timer,omitempty"`
	// Number of LSAs currently being flooded
	QueueLength uint32 `protobuf:"varint,51,opt,name=queue_length,json=queueLength,proto3" json:"queue_length,omitempty"`
	// Link floodlist
	LinkFloodList []*Ospfv3EdmLsaSum `protobuf:"bytes,52,rep,name=link_flood_list,json=linkFloodList,proto3" json:"link_flood_list,omitempty"`
	// Area scope floodlist
	AreaFloodList []*Ospfv3EdmLsaSum `protobuf:"bytes,53,rep,name=area_flood_list,json=areaFloodList,proto3" json:"area_flood_list,omitempty"`
	// AS scope floodlist
	AsFloodList          []*Ospfv3EdmLsaSum `protobuf:"bytes,54,rep,name=as_flood_list,json=asFloodList,proto3" json:"as_flood_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Ospfv3EdmFloodList) Reset()         { *m = Ospfv3EdmFloodList{} }
func (m *Ospfv3EdmFloodList) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmFloodList) ProtoMessage()    {}
func (*Ospfv3EdmFloodList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_flood_list_7d7849ea8f2bfbc4, []int{1}
}
func (m *Ospfv3EdmFloodList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmFloodList.Unmarshal(m, b)
}
func (m *Ospfv3EdmFloodList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmFloodList.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmFloodList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmFloodList.Merge(dst, src)
}
func (m *Ospfv3EdmFloodList) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmFloodList.Size(m)
}
func (m *Ospfv3EdmFloodList) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmFloodList.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmFloodList proto.InternalMessageInfo

func (m *Ospfv3EdmFloodList) GetLsTransmissionTimer() uint32 {
	if m != nil {
		return m.LsTransmissionTimer
	}
	return 0
}

func (m *Ospfv3EdmFloodList) GetQueueLength() uint32 {
	if m != nil {
		return m.QueueLength
	}
	return 0
}

func (m *Ospfv3EdmFloodList) GetLinkFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.LinkFloodList
	}
	return nil
}

func (m *Ospfv3EdmFloodList) GetAreaFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.AreaFloodList
	}
	return nil
}

func (m *Ospfv3EdmFloodList) GetAsFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.AsFloodList
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
	return fileDescriptor_ospfv3_edm_flood_list_7d7849ea8f2bfbc4, []int{2}
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
	proto.RegisterType((*Ospfv3EdmFloodList_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.flood_list_table.flood.ospfv3_edm_flood_list_KEYS")
	proto.RegisterType((*Ospfv3EdmFloodList)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.flood_list_table.flood.ospfv3_edm_flood_list")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.flood_list_table.flood.ospfv3_edm_lsa_sum")
}

func init() {
	proto.RegisterFile("ospfv3_edm_flood_list.proto", fileDescriptor_ospfv3_edm_flood_list_7d7849ea8f2bfbc4)
}

var fileDescriptor_ospfv3_edm_flood_list_7d7849ea8f2bfbc4 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0x15, 0x5a, 0x8a, 0x70, 0x9a, 0x56, 0x18, 0x4a, 0x03, 0x2c, 0xc7, 0x09, 0xd0, 0x4d,
	0x19, 0xee, 0x4a, 0x07, 0xb6, 0x0e, 0x20, 0x55, 0x9c, 0x3a, 0xa4, 0xb7, 0x30, 0x3d, 0xf9, 0xce,
	0x2f, 0x57, 0x0b, 0x27, 0x4e, 0xfd, 0x9c, 0x53, 0x3b, 0xf1, 0x45, 0x18, 0xf8, 0x4e, 0x7c, 0x19,
	0x46, 0x64, 0x27, 0xcd, 0x45, 0xa2, 0x33, 0x5d, 0x22, 0xe7, 0xff, 0xff, 0xbd, 0xe7, 0xbf, 0xad,
	0x97, 0xb0, 0x37, 0x86, 0xea, 0x62, 0x33, 0x03, 0x94, 0x25, 0x14, 0xda, 0x18, 0x09, 0x5a, 0x91,
	0xcb, 0x6a, 0x6b, 0x9c, 0xe1, 0x72, 0xa5, 0x68, 0x65, 0x40, 0x19, 0x82, 0x1b, 0x0b, 0xaa, 0xde,
	0x9c, 0x42, 0x87, 0x9b, 0x1a, 0x6d, 0xd6, 0xae, 0x3d, 0xbb, 0x42, 0x22, 0xa4, 0xbb, 0x55, 0x26,
	0xb1, 0x10, 0x8d, 0x76, 0xb0, 0xb1, 0x45, 0x26, 0x2c, 0x0a, 0x0a, 0xcf, 0x6c, 0xdb, 0x1f, 0x9c,
	0x58, 0x6a, 0x6c, 0x85, 0xf1, 0x0f, 0xf6, 0xfa, 0xde, 0x10, 0xf0, 0xf5, 0xf3, 0xb7, 0x4b, 0xfe,
	0x96, 0xed, 0x77, 0x6d, 0xa1, 0x12, 0x25, 0xa6, 0xd1, 0x28, 0x9a, 0x3c, 0xcd, 0xe3, 0x4e, 0xbb,
	0x10, 0x25, 0xf2, 0x63, 0xf6, 0xc4, 0xf7, 0x07, 0x25, 0xd3, 0x47, 0xa3, 0x68, 0x92, 0xe4, 0x7b,
	0xfe, 0xf5, 0x5c, 0xf2, 0xf7, 0xec, 0x40, 0x55, 0x0e, 0x6d, 0x21, 0x56, 0xd8, 0x56, 0xef, 0x84,
	0xea, 0xa4, 0x57, 0x7d, 0xfd, 0xf8, 0xf7, 0x2e, 0x3b, 0xba, 0x37, 0x01, 0x9f, 0xb2, 0x23, 0x4d,
	0xe0, 0xac, 0xa8, 0xa8, 0x54, 0x44, 0xca, 0x54, 0xe0, 0x54, 0x89, 0x36, 0x9d, 0x86, 0x7d, 0x9e,
	0x6b, 0x5a, 0x0c, 0xbc, 0x85, 0xb7, 0x7c, 0xe0, 0xeb, 0x06, 0x1b, 0x04, 0x8d, 0xd5, 0xda, 0x5d,
	0xa5, 0xb3, 0x80, 0xc6, 0x41, 0x9b, 0x07, 0x89, 0xff, 0x8a, 0xd8, 0xa1, 0x56, 0xd5, 0xf7, 0xc1,
	0x56, 0xe9, 0xc9, 0x68, 0x67, 0x12, 0x4f, 0x6f, 0xb2, 0xff, 0x71, 0xe5, 0xd9, 0xe0, 0xb4, 0x9a,
	0x04, 0x50, 0x53, 0xe6, 0x89, 0x0f, 0xf4, 0xc5, 0xbb, 0x73, 0x7f, 0x72, 0x1f, 0x31, 0x5c, 0xea,
	0x20, 0xe2, 0xc7, 0x87, 0x8e, 0xe8, 0xf9, 0x6d, 0xc4, 0x9f, 0x11, 0x4b, 0x04, 0x0d, 0x03, 0x9e,
	0x3e, 0x70, 0xc0, 0x58, 0x50, 0x1f, 0x6f, 0xfc, 0x27, 0x62, 0xfc, 0x5f, 0x86, 0x7f, 0x60, 0x87,
	0x57, 0x28, 0x24, 0xda, 0xa0, 0xb8, 0xdb, 0xfa, 0x6e, 0xa4, 0x93, 0x56, 0x9e, 0x93, 0x58, 0xdc,
	0xd6, 0xc8, 0xdf, 0xb1, 0x83, 0x01, 0x27, 0xd6, 0xd8, 0xcd, 0xf6, 0x7e, 0x8f, 0x9d, 0xad, 0x91,
	0x8f, 0x59, 0x32, 0xa0, 0x94, 0xec, 0x06, 0x3c, 0xee, 0xa1, 0x73, 0xc9, 0x3f, 0xb1, 0x57, 0x1d,
	0x23, 0xe4, 0x06, 0xad, 0x53, 0xa4, 0xaa, 0x35, 0x58, 0xd3, 0x38, 0xb4, 0xe9, 0x6e, 0xe0, 0x8f,
	0x5b, 0xe0, 0x6c, 0xeb, 0xe7, 0xc1, 0xe6, 0x27, 0xec, 0x65, 0x57, 0x4b, 0x78, 0xdd, 0x60, 0xe5,
	0xbf, 0xa3, 0xa6, 0x5c, 0xa2, 0x4d, 0x1f, 0x8f, 0xa2, 0xc9, 0xb3, 0xfc, 0x45, 0xeb, 0x5e, 0x76,
	0xe6, 0x45, 0xf0, 0x96, 0x7b, 0xe1, 0xf7, 0x31, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x3a,
	0x8f, 0xaa, 0x5d, 0x04, 0x00, 0x00,
}
