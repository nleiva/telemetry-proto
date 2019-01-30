// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_traffic_intf.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_interface_statistics_interface_statistic

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

// IS-IS interface traffic data
type IsisShTrafficIntf_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTrafficIntf_KEYS) Reset()         { *m = IsisShTrafficIntf_KEYS{} }
func (m *IsisShTrafficIntf_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShTrafficIntf_KEYS) ProtoMessage()    {}
func (*IsisShTrafficIntf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{0}
}
func (m *IsisShTrafficIntf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTrafficIntf_KEYS.Unmarshal(m, b)
}
func (m *IsisShTrafficIntf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTrafficIntf_KEYS.Marshal(b, m, deterministic)
}
func (dst *IsisShTrafficIntf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTrafficIntf_KEYS.Merge(dst, src)
}
func (m *IsisShTrafficIntf_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShTrafficIntf_KEYS.Size(m)
}
func (m *IsisShTrafficIntf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTrafficIntf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTrafficIntf_KEYS proto.InternalMessageInfo

func (m *IsisShTrafficIntf_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTrafficIntf_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShTrafficIntf struct {
	// Interface to which traffic statistics apply
	TrafficInterface string `protobuf:"bytes,50,opt,name=traffic_interface,json=trafficInterface,proto3" json:"traffic_interface,omitempty"`
	// Interface media class
	InterfaceMediaType string `protobuf:"bytes,51,opt,name=interface_media_type,json=interfaceMediaType,proto3" json:"interface_media_type,omitempty"`
	// P2P interface statistics. NULL for non-P2P interfaces
	P2PStatistics *IsisTrafficIntfP2PType `protobuf:"bytes,52,opt,name=p2_p_statistics,json=p2PStatistics,proto3" json:"p2_p_statistics,omitempty"`
	// Per-area data
	PerAreaData          []*IsisShTrafficIntfArea `protobuf:"bytes,53,rep,name=per_area_data,json=perAreaData,proto3" json:"per_area_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IsisShTrafficIntf) Reset()         { *m = IsisShTrafficIntf{} }
func (m *IsisShTrafficIntf) String() string { return proto.CompactTextString(m) }
func (*IsisShTrafficIntf) ProtoMessage()    {}
func (*IsisShTrafficIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{1}
}
func (m *IsisShTrafficIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTrafficIntf.Unmarshal(m, b)
}
func (m *IsisShTrafficIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTrafficIntf.Marshal(b, m, deterministic)
}
func (dst *IsisShTrafficIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTrafficIntf.Merge(dst, src)
}
func (m *IsisShTrafficIntf) XXX_Size() int {
	return xxx_messageInfo_IsisShTrafficIntf.Size(m)
}
func (m *IsisShTrafficIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTrafficIntf.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTrafficIntf proto.InternalMessageInfo

func (m *IsisShTrafficIntf) GetTrafficInterface() string {
	if m != nil {
		return m.TrafficInterface
	}
	return ""
}

func (m *IsisShTrafficIntf) GetInterfaceMediaType() string {
	if m != nil {
		return m.InterfaceMediaType
	}
	return ""
}

func (m *IsisShTrafficIntf) GetP2PStatistics() *IsisTrafficIntfP2PType {
	if m != nil {
		return m.P2PStatistics
	}
	return nil
}

func (m *IsisShTrafficIntf) GetPerAreaData() []*IsisShTrafficIntfArea {
	if m != nil {
		return m.PerAreaData
	}
	return nil
}

// Per-interface, per-PDU statistics
type IsisTrafficPduCountType struct {
	// PDUs received
	PduReceiveCount uint32 `protobuf:"varint,1,opt,name=pdu_receive_count,json=pduReceiveCount,proto3" json:"pdu_receive_count,omitempty"`
	// PDUs sent
	PduSendCount         uint32   `protobuf:"varint,2,opt,name=pdu_send_count,json=pduSendCount,proto3" json:"pdu_send_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTrafficPduCountType) Reset()         { *m = IsisTrafficPduCountType{} }
func (m *IsisTrafficPduCountType) String() string { return proto.CompactTextString(m) }
func (*IsisTrafficPduCountType) ProtoMessage()    {}
func (*IsisTrafficPduCountType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{2}
}
func (m *IsisTrafficPduCountType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTrafficPduCountType.Unmarshal(m, b)
}
func (m *IsisTrafficPduCountType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTrafficPduCountType.Marshal(b, m, deterministic)
}
func (dst *IsisTrafficPduCountType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTrafficPduCountType.Merge(dst, src)
}
func (m *IsisTrafficPduCountType) XXX_Size() int {
	return xxx_messageInfo_IsisTrafficPduCountType.Size(m)
}
func (m *IsisTrafficPduCountType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTrafficPduCountType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTrafficPduCountType proto.InternalMessageInfo

func (m *IsisTrafficPduCountType) GetPduReceiveCount() uint32 {
	if m != nil {
		return m.PduReceiveCount
	}
	return 0
}

func (m *IsisTrafficPduCountType) GetPduSendCount() uint32 {
	if m != nil {
		return m.PduSendCount
	}
	return 0
}

// Per-interface point-to-point statistics
type IsisTrafficIntfP2PType struct {
	// IIH statistics
	IihCount *IsisTrafficPduCountType `protobuf:"bytes,1,opt,name=iih_count,json=iihCount,proto3" json:"iih_count,omitempty"`
	// IIHs not sent due to memory exhaustion
	MemoryExhaustedIihCount uint32 `protobuf:"varint,2,opt,name=memory_exhausted_iih_count,json=memoryExhaustedIihCount,proto3" json:"memory_exhausted_iih_count,omitempty"`
	// LSP retransmissions
	LspRetransmitCount   uint32   `protobuf:"varint,3,opt,name=lsp_retransmit_count,json=lspRetransmitCount,proto3" json:"lsp_retransmit_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTrafficIntfP2PType) Reset()         { *m = IsisTrafficIntfP2PType{} }
func (m *IsisTrafficIntfP2PType) String() string { return proto.CompactTextString(m) }
func (*IsisTrafficIntfP2PType) ProtoMessage()    {}
func (*IsisTrafficIntfP2PType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{3}
}
func (m *IsisTrafficIntfP2PType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTrafficIntfP2PType.Unmarshal(m, b)
}
func (m *IsisTrafficIntfP2PType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTrafficIntfP2PType.Marshal(b, m, deterministic)
}
func (dst *IsisTrafficIntfP2PType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTrafficIntfP2PType.Merge(dst, src)
}
func (m *IsisTrafficIntfP2PType) XXX_Size() int {
	return xxx_messageInfo_IsisTrafficIntfP2PType.Size(m)
}
func (m *IsisTrafficIntfP2PType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTrafficIntfP2PType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTrafficIntfP2PType proto.InternalMessageInfo

func (m *IsisTrafficIntfP2PType) GetIihCount() *IsisTrafficPduCountType {
	if m != nil {
		return m.IihCount
	}
	return nil
}

func (m *IsisTrafficIntfP2PType) GetMemoryExhaustedIihCount() uint32 {
	if m != nil {
		return m.MemoryExhaustedIihCount
	}
	return 0
}

func (m *IsisTrafficIntfP2PType) GetLspRetransmitCount() uint32 {
	if m != nil {
		return m.LspRetransmitCount
	}
	return 0
}

// Per-interface, per-area LAN-only statistics
type IsisTrafficIntfAreaLanType struct {
	// IIH statistics
	IihCount *IsisTrafficPduCountType `protobuf:"bytes,1,opt,name=iih_count,json=iihCount,proto3" json:"iih_count,omitempty"`
	// IIHs not sent due to memory exhaustion
	MemoryExhaustedIihCount uint32 `protobuf:"varint,2,opt,name=memory_exhausted_iih_count,json=memoryExhaustedIihCount,proto3" json:"memory_exhausted_iih_count,omitempty"`
	// DIS elections
	DisElectionCount     uint32   `protobuf:"varint,3,opt,name=dis_election_count,json=disElectionCount,proto3" json:"dis_election_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTrafficIntfAreaLanType) Reset()         { *m = IsisTrafficIntfAreaLanType{} }
func (m *IsisTrafficIntfAreaLanType) String() string { return proto.CompactTextString(m) }
func (*IsisTrafficIntfAreaLanType) ProtoMessage()    {}
func (*IsisTrafficIntfAreaLanType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{4}
}
func (m *IsisTrafficIntfAreaLanType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTrafficIntfAreaLanType.Unmarshal(m, b)
}
func (m *IsisTrafficIntfAreaLanType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTrafficIntfAreaLanType.Marshal(b, m, deterministic)
}
func (dst *IsisTrafficIntfAreaLanType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTrafficIntfAreaLanType.Merge(dst, src)
}
func (m *IsisTrafficIntfAreaLanType) XXX_Size() int {
	return xxx_messageInfo_IsisTrafficIntfAreaLanType.Size(m)
}
func (m *IsisTrafficIntfAreaLanType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTrafficIntfAreaLanType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTrafficIntfAreaLanType proto.InternalMessageInfo

func (m *IsisTrafficIntfAreaLanType) GetIihCount() *IsisTrafficPduCountType {
	if m != nil {
		return m.IihCount
	}
	return nil
}

func (m *IsisTrafficIntfAreaLanType) GetMemoryExhaustedIihCount() uint32 {
	if m != nil {
		return m.MemoryExhaustedIihCount
	}
	return 0
}

func (m *IsisTrafficIntfAreaLanType) GetDisElectionCount() uint32 {
	if m != nil {
		return m.DisElectionCount
	}
	return 0
}

// Per-interface, per-area statistics
type IsisTrafficIntfAreaType struct {
	// LSP statistics
	LspCount *IsisTrafficPduCountType `protobuf:"bytes,1,opt,name=lsp_count,json=lspCount,proto3" json:"lsp_count,omitempty"`
	// Count of LSPs dropped due to minimum arrival time config
	LspDropCount uint32 `protobuf:"varint,2,opt,name=lsp_drop_count,json=lspDropCount,proto3" json:"lsp_drop_count,omitempty"`
	// CSNP statistics
	CsnpCount *IsisTrafficPduCountType `protobuf:"bytes,3,opt,name=csnp_count,json=csnpCount,proto3" json:"csnp_count,omitempty"`
	// PSNP statistics
	PsnpCount *IsisTrafficPduCountType `protobuf:"bytes,4,opt,name=psnp_count,json=psnpCount,proto3" json:"psnp_count,omitempty"`
	// Count of LSPs already received by neighbors and not flooded
	LspFloodingDupCount  uint32   `protobuf:"varint,5,opt,name=lsp_flooding_dup_count,json=lspFloodingDupCount,proto3" json:"lsp_flooding_dup_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTrafficIntfAreaType) Reset()         { *m = IsisTrafficIntfAreaType{} }
func (m *IsisTrafficIntfAreaType) String() string { return proto.CompactTextString(m) }
func (*IsisTrafficIntfAreaType) ProtoMessage()    {}
func (*IsisTrafficIntfAreaType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{5}
}
func (m *IsisTrafficIntfAreaType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTrafficIntfAreaType.Unmarshal(m, b)
}
func (m *IsisTrafficIntfAreaType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTrafficIntfAreaType.Marshal(b, m, deterministic)
}
func (dst *IsisTrafficIntfAreaType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTrafficIntfAreaType.Merge(dst, src)
}
func (m *IsisTrafficIntfAreaType) XXX_Size() int {
	return xxx_messageInfo_IsisTrafficIntfAreaType.Size(m)
}
func (m *IsisTrafficIntfAreaType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTrafficIntfAreaType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTrafficIntfAreaType proto.InternalMessageInfo

func (m *IsisTrafficIntfAreaType) GetLspCount() *IsisTrafficPduCountType {
	if m != nil {
		return m.LspCount
	}
	return nil
}

func (m *IsisTrafficIntfAreaType) GetLspDropCount() uint32 {
	if m != nil {
		return m.LspDropCount
	}
	return 0
}

func (m *IsisTrafficIntfAreaType) GetCsnpCount() *IsisTrafficPduCountType {
	if m != nil {
		return m.CsnpCount
	}
	return nil
}

func (m *IsisTrafficIntfAreaType) GetPsnpCount() *IsisTrafficPduCountType {
	if m != nil {
		return m.PsnpCount
	}
	return nil
}

func (m *IsisTrafficIntfAreaType) GetLspFloodingDupCount() uint32 {
	if m != nil {
		return m.LspFloodingDupCount
	}
	return 0
}

// Per-interface, per-area data
type IsisShTrafficIntfArea struct {
	// Level of the area this data relates to
	Level string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	// Statistics
	Statistics *IsisTrafficIntfAreaType `protobuf:"bytes,2,opt,name=statistics,proto3" json:"statistics,omitempty"`
	// LAN interface statistics. NULL for non-LAN interfaces
	LanData              *IsisTrafficIntfAreaLanType `protobuf:"bytes,3,opt,name=lan_data,json=lanData,proto3" json:"lan_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *IsisShTrafficIntfArea) Reset()         { *m = IsisShTrafficIntfArea{} }
func (m *IsisShTrafficIntfArea) String() string { return proto.CompactTextString(m) }
func (*IsisShTrafficIntfArea) ProtoMessage()    {}
func (*IsisShTrafficIntfArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3, []int{6}
}
func (m *IsisShTrafficIntfArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTrafficIntfArea.Unmarshal(m, b)
}
func (m *IsisShTrafficIntfArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTrafficIntfArea.Marshal(b, m, deterministic)
}
func (dst *IsisShTrafficIntfArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTrafficIntfArea.Merge(dst, src)
}
func (m *IsisShTrafficIntfArea) XXX_Size() int {
	return xxx_messageInfo_IsisShTrafficIntfArea.Size(m)
}
func (m *IsisShTrafficIntfArea) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTrafficIntfArea.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTrafficIntfArea proto.InternalMessageInfo

func (m *IsisShTrafficIntfArea) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShTrafficIntfArea) GetStatistics() *IsisTrafficIntfAreaType {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *IsisShTrafficIntfArea) GetLanData() *IsisTrafficIntfAreaLanType {
	if m != nil {
		return m.LanData
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShTrafficIntf_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_sh_traffic_intf_KEYS")
	proto.RegisterType((*IsisShTrafficIntf)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_sh_traffic_intf")
	proto.RegisterType((*IsisTrafficPduCountType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_traffic_pdu_count_type")
	proto.RegisterType((*IsisTrafficIntfP2PType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_traffic_intf_p2p_type")
	proto.RegisterType((*IsisTrafficIntfAreaLanType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_traffic_intf_area_lan_type")
	proto.RegisterType((*IsisTrafficIntfAreaType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_traffic_intf_area_type")
	proto.RegisterType((*IsisShTrafficIntfArea)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.interface_statistics.interface_statistic.isis_sh_traffic_intf_area")
}

func init() {
	proto.RegisterFile("isis_sh_traffic_intf.proto", fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3)
}

var fileDescriptor_isis_sh_traffic_intf_66190df178a53cf3 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xa6, 0x2d, 0xb4, 0x4e, 0xd3, 0x1f, 0x13, 0x41, 0x28, 0x07, 0xaa, 0x00, 0x52, 0x05,
	0x28, 0xaa, 0xb6, 0x70, 0xe2, 0x84, 0x68, 0x91, 0x2a, 0x04, 0x42, 0x5b, 0x2e, 0x48, 0x48, 0x96,
	0x59, 0x4f, 0x1a, 0x4b, 0x8e, 0x6d, 0xd9, 0xde, 0xd2, 0x1c, 0x78, 0x0d, 0x04, 0xe2, 0xc2, 0x01,
	0x5e, 0x81, 0x57, 0xe0, 0x55, 0x78, 0x0c, 0x64, 0xef, 0x7f, 0x09, 0x47, 0x5a, 0xc4, 0x6d, 0x77,
	0xe6, 0x9b, 0x99, 0x6f, 0x3e, 0x8f, 0x77, 0x16, 0x6d, 0x71, 0xcb, 0x2d, 0xb1, 0x13, 0xe2, 0x0c,
	0x1d, 0x8f, 0x79, 0x4a, 0xb8, 0x74, 0xe3, 0x91, 0x36, 0xca, 0x29, 0xfc, 0x26, 0xe5, 0x36, 0x55,
	0x84, 0x2b, 0x4b, 0x4e, 0x0d, 0x49, 0x85, 0xb4, 0x24, 0xa0, 0x95, 0x06, 0x33, 0xf2, 0x4f, 0x23,
	0x2e, 0xad, 0xa3, 0x32, 0x85, 0xfa, 0x69, 0xc4, 0xa5, 0x03, 0x33, 0xa6, 0x29, 0x10, 0xeb, 0xa8,
	0xe3, 0xd6, 0xf1, 0xd4, 0xce, 0x33, 0x0e, 0x8f, 0xd1, 0xf5, 0x79, 0xb5, 0xc9, 0xb3, 0x83, 0xd7,
	0x47, 0xf8, 0x16, 0xea, 0x95, 0x19, 0x89, 0xa4, 0x53, 0x18, 0x44, 0xdb, 0xd1, 0xce, 0x4a, 0xb2,
	0x5a, 0x1a, 0x5f, 0xd0, 0x29, 0xe0, 0x3b, 0x68, 0xad, 0x4e, 0x1c, 0x50, 0x9d, 0x80, 0xea, 0x55,
	0x56, 0x0f, 0x1b, 0x7e, 0x5f, 0x40, 0xfd, 0x79, 0x95, 0xf0, 0x3d, 0xb4, 0xd9, 0x78, 0xcf, 0x23,
	0x06, 0x71, 0x48, 0xb1, 0x51, 0x38, 0x0e, 0x4b, 0x3b, 0xde, 0x45, 0xfd, 0xba, 0xd8, 0x14, 0x18,
	0xa7, 0xc4, 0xcd, 0x34, 0x0c, 0xf6, 0x02, 0x1e, 0x57, 0xbe, 0xe7, 0xde, 0xf5, 0x6a, 0xa6, 0x01,
	0x7f, 0x89, 0xd0, 0xba, 0x8e, 0x89, 0x6e, 0x08, 0x31, 0x78, 0xb0, 0x1d, 0xed, 0x74, 0xe3, 0xd3,
	0xd1, 0xdf, 0x54, 0x36, 0x24, 0x68, 0x6b, 0xaa, 0x63, 0x1d, 0xd8, 0x26, 0x3d, 0x1d, 0xbf, 0x3c,
	0xaa, 0xa2, 0xf1, 0xe7, 0x08, 0xf5, 0x34, 0x18, 0x42, 0x0d, 0x50, 0xc2, 0xa8, 0xa3, 0x83, 0x87,
	0xdb, 0x0b, 0x3b, 0xdd, 0xf8, 0xdd, 0x39, 0x10, 0x3c, 0x7b, 0xee, 0x9e, 0x43, 0xd2, 0xd5, 0x60,
	0x1e, 0x1b, 0xa0, 0xfb, 0xd4, 0xd1, 0xa1, 0x42, 0x37, 0x5a, 0xad, 0x68, 0x96, 0x91, 0x54, 0x65,
	0xd2, 0x85, 0x5e, 0xf0, 0x5d, 0xb4, 0xe9, 0x2d, 0x06, 0x52, 0xe0, 0x27, 0x90, 0x7b, 0xc2, 0x9c,
	0xf4, 0x92, 0x75, 0xcd, 0xb2, 0x24, 0xb7, 0x3f, 0xf1, 0x66, 0x7c, 0x1b, 0xad, 0x79, 0xac, 0x05,
	0xc9, 0x0a, 0x60, 0x27, 0x00, 0x57, 0x35, 0xcb, 0x8e, 0x40, 0xb2, 0x80, 0x1a, 0x7e, 0xed, 0x14,
	0xf7, 0x61, 0xae, 0x78, 0xf8, 0x43, 0x84, 0x56, 0x38, 0x9f, 0x34, 0x2a, 0x75, 0xe3, 0xd9, 0x39,
	0x1e, 0x65, 0xbb, 0xff, 0x64, 0x99, 0xf3, 0x49, 0xde, 0xdd, 0x23, 0xb4, 0x35, 0x85, 0xa9, 0x32,
	0x33, 0x02, 0xa7, 0x13, 0x9a, 0x59, 0x07, 0x8c, 0xd4, 0x44, 0xf3, 0x4e, 0xaf, 0xe5, 0x88, 0x83,
	0x12, 0x70, 0x58, 0x06, 0xef, 0xa2, 0xbe, 0xb0, 0x9a, 0x18, 0x70, 0x86, 0x4a, 0x3b, 0xe5, 0xae,
	0x08, 0x5b, 0x08, 0x61, 0x58, 0x58, 0x9d, 0x54, 0xae, 0x5c, 0xa6, 0x6f, 0x1d, 0x74, 0xf3, 0x77,
	0x99, 0xc2, 0x0c, 0x09, 0x2a, 0xff, 0x67, 0xad, 0xee, 0x23, 0xcc, 0xb8, 0x25, 0x20, 0x20, 0x75,
	0x5c, 0xc9, 0x96, 0x52, 0x1b, 0x8c, 0xdb, 0x83, 0xc2, 0x91, 0xeb, 0xf4, 0x63, 0xf1, 0xcc, 0x00,
	0xd7, 0x3a, 0x55, 0x1a, 0x79, 0xe9, 0xff, 0x15, 0x8d, 0x84, 0xd5, 0xd5, 0x6d, 0xf1, 0xbc, 0x98,
	0x51, 0xba, 0x7d, 0x5b, 0x84, 0xd5, 0xfb, 0x46, 0x15, 0xa8, 0x8f, 0x11, 0x42, 0xa9, 0x95, 0xba,
	0xa1, 0xc2, 0x85, 0xf2, 0x5f, 0xf1, 0x64, 0x6a, 0x6a, 0xba, 0xa6, 0xb6, 0x78, 0xe1, 0xd4, 0x74,
	0x45, 0x6d, 0x0f, 0x5d, 0xf5, 0xda, 0x8e, 0x85, 0x52, 0x8c, 0xcb, 0x63, 0xc2, 0xb2, 0x92, 0xe5,
	0x52, 0xd0, 0xf8, 0x8a, 0xb0, 0xfa, 0x69, 0xe1, 0xdc, 0xcf, 0xf2, 0xa0, 0xe1, 0xcf, 0xce, 0x1f,
	0x96, 0xa5, 0x1f, 0x26, 0xdc, 0x47, 0x4b, 0x02, 0x4e, 0x40, 0x14, 0x4b, 0x32, 0x7f, 0xc1, 0x9f,
	0x22, 0x84, 0x1a, 0x9b, 0xa7, 0x73, 0xee, 0x1a, 0xb4, 0xa7, 0x3d, 0x69, 0x90, 0xf1, 0xe7, 0xb3,
	0xec, 0x3f, 0x15, 0x61, 0xe5, 0xe4, 0x83, 0xf3, 0xfe, 0x42, 0x98, 0x95, 0xdf, 0xab, 0xe4, 0xb2,
	0xa0, 0xd2, 0x2f, 0x9d, 0xb7, 0x97, 0xc2, 0xbf, 0xcf, 0xde, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5a, 0x33, 0xe5, 0x4b, 0x19, 0x09, 0x00, 0x00,
}
