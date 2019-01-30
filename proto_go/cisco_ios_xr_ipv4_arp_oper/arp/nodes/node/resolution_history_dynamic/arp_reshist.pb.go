// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arp_reshist.proto

package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_resolution_history_dynamic

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

// ARP Resolution History
type ArpReshist_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArpReshist_KEYS) Reset()         { *m = ArpReshist_KEYS{} }
func (m *ArpReshist_KEYS) String() string { return proto.CompactTextString(m) }
func (*ArpReshist_KEYS) ProtoMessage()    {}
func (*ArpReshist_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_arp_reshist_9dbd724845497750, []int{0}
}
func (m *ArpReshist_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshist_KEYS.Unmarshal(m, b)
}
func (m *ArpReshist_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshist_KEYS.Marshal(b, m, deterministic)
}
func (dst *ArpReshist_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshist_KEYS.Merge(dst, src)
}
func (m *ArpReshist_KEYS) XXX_Size() int {
	return xxx_messageInfo_ArpReshist_KEYS.Size(m)
}
func (m *ArpReshist_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpReshist_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ArpReshist_KEYS proto.InternalMessageInfo

func (m *ArpReshist_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type ArpReshist struct {
	// Resolution history array
	ArpEntry             []*ArpReshistEntry `protobuf:"bytes,50,rep,name=arp_entry,json=arpEntry,proto3" json:"arp_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ArpReshist) Reset()         { *m = ArpReshist{} }
func (m *ArpReshist) String() string { return proto.CompactTextString(m) }
func (*ArpReshist) ProtoMessage()    {}
func (*ArpReshist) Descriptor() ([]byte, []int) {
	return fileDescriptor_arp_reshist_9dbd724845497750, []int{1}
}
func (m *ArpReshist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshist.Unmarshal(m, b)
}
func (m *ArpReshist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshist.Marshal(b, m, deterministic)
}
func (dst *ArpReshist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshist.Merge(dst, src)
}
func (m *ArpReshist) XXX_Size() int {
	return xxx_messageInfo_ArpReshist.Size(m)
}
func (m *ArpReshist) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpReshist.DiscardUnknown(m)
}

var xxx_messageInfo_ArpReshist proto.InternalMessageInfo

func (m *ArpReshist) GetArpEntry() []*ArpReshistEntry {
	if m != nil {
		return m.ArpEntry
	}
	return nil
}

// ARP resolution history entry
type ArpReshistEntry struct {
	// Timestamp for entry in nanoseconds since Epoch, i.e. since 00:00:00 UTC, January 1, 1970
	NsecTimestamp uint64 `protobuf:"varint,1,opt,name=nsec_timestamp,json=nsecTimestamp,proto3" json:"nsec_timestamp,omitempty"`
	// Interface
	IdbInterfaceName string `protobuf:"bytes,2,opt,name=idb_interface_name,json=idbInterfaceName,proto3" json:"idb_interface_name,omitempty"`
	// IPv4 address
	Ipv4Address string `protobuf:"bytes,3,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	// MAC address
	MacAddress string `protobuf:"bytes,4,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Resolution status
	Status string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Resolving Client ID
	ClientId int32 `protobuf:"zigzag32,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// ARP entry state
	EntryState int32 `protobuf:"zigzag32,7,opt,name=entry_state,json=entryState,proto3" json:"entry_state,omitempty"`
	// Resolution Request count
	ResolutionRequestCount uint32   `protobuf:"varint,8,opt,name=resolution_request_count,json=resolutionRequestCount,proto3" json:"resolution_request_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ArpReshistEntry) Reset()         { *m = ArpReshistEntry{} }
func (m *ArpReshistEntry) String() string { return proto.CompactTextString(m) }
func (*ArpReshistEntry) ProtoMessage()    {}
func (*ArpReshistEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_arp_reshist_9dbd724845497750, []int{2}
}
func (m *ArpReshistEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshistEntry.Unmarshal(m, b)
}
func (m *ArpReshistEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshistEntry.Marshal(b, m, deterministic)
}
func (dst *ArpReshistEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshistEntry.Merge(dst, src)
}
func (m *ArpReshistEntry) XXX_Size() int {
	return xxx_messageInfo_ArpReshistEntry.Size(m)
}
func (m *ArpReshistEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpReshistEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ArpReshistEntry proto.InternalMessageInfo

func (m *ArpReshistEntry) GetNsecTimestamp() uint64 {
	if m != nil {
		return m.NsecTimestamp
	}
	return 0
}

func (m *ArpReshistEntry) GetIdbInterfaceName() string {
	if m != nil {
		return m.IdbInterfaceName
	}
	return ""
}

func (m *ArpReshistEntry) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *ArpReshistEntry) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *ArpReshistEntry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ArpReshistEntry) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *ArpReshistEntry) GetEntryState() int32 {
	if m != nil {
		return m.EntryState
	}
	return 0
}

func (m *ArpReshistEntry) GetResolutionRequestCount() uint32 {
	if m != nil {
		return m.ResolutionRequestCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ArpReshist_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_dynamic.arp_reshist_KEYS")
	proto.RegisterType((*ArpReshist)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_dynamic.arp_reshist")
	proto.RegisterType((*ArpReshistEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_dynamic.arp_reshist_entry")
}

func init() { proto.RegisterFile("arp_reshist.proto", fileDescriptor_arp_reshist_9dbd724845497750) }

var fileDescriptor_arp_reshist_9dbd724845497750 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4b, 0x23, 0x31,
	0x18, 0xc5, 0x99, 0xb6, 0xdb, 0x9d, 0x66, 0xb6, 0x4b, 0x9b, 0x43, 0x09, 0xec, 0x61, 0xc7, 0x82,
	0x30, 0x07, 0x19, 0xa1, 0x7a, 0xf0, 0x2a, 0xda, 0x43, 0x11, 0x3c, 0x4c, 0x05, 0xf1, 0x14, 0xd2,
	0xe4, 0x13, 0x03, 0x9d, 0x24, 0x26, 0x19, 0xb1, 0x67, 0xff, 0x2d, 0xff, 0x38, 0x49, 0xa6, 0xd5,
	0xf1, 0xee, 0x65, 0x98, 0xfc, 0xde, 0xcb, 0xe3, 0x7b, 0x1f, 0x41, 0x53, 0x66, 0x0d, 0xb5, 0xe0,
	0x9e, 0xa4, 0xf3, 0xa5, 0xb1, 0xda, 0x6b, 0x7c, 0xcd, 0xa5, 0xe3, 0x9a, 0x4a, 0xed, 0xe8, 0xab,
	0xa5, 0xd2, 0xbc, 0x9c, 0xd3, 0x60, 0xd2, 0x06, 0x6c, 0xc9, 0xac, 0x29, 0x95, 0x16, 0xe0, 0xe2,
	0xb7, 0xb4, 0xe0, 0xf4, 0xb6, 0xf1, 0x52, 0x2b, 0x1a, 0x02, 0xb4, 0xdd, 0x51, 0xb1, 0x53, 0xac,
	0x96, 0x7c, 0x7e, 0x8a, 0x26, 0x9d, 0x68, 0x7a, 0xb3, 0x7c, 0x58, 0xe3, 0x7f, 0x68, 0x14, 0xae,
	0x51, 0xc5, 0x6a, 0x20, 0x49, 0x9e, 0x14, 0xa3, 0x2a, 0x0d, 0xe0, 0x96, 0xd5, 0x30, 0x7f, 0x4b,
	0x50, 0xd6, 0xb9, 0x81, 0x3d, 0x1a, 0x85, 0x23, 0x28, 0x6f, 0x77, 0x64, 0x91, 0xf7, 0x8b, 0x6c,
	0x71, 0x5f, 0xfe, 0xc4, 0x68, 0x65, 0x77, 0xae, 0x18, 0x5f, 0xa5, 0xcc, 0x9a, 0x65, 0xf8, 0x9b,
	0xbf, 0xf7, 0xbe, 0xad, 0xa4, 0xd5, 0xf1, 0x31, 0xfa, 0xab, 0x1c, 0x70, 0xea, 0x65, 0x0d, 0xce,
	0xb3, 0xda, 0xc4, 0xe9, 0x07, 0xd5, 0x38, 0xd0, 0xbb, 0x03, 0xc4, 0x27, 0x08, 0x4b, 0xb1, 0xa1,
	0x52, 0x79, 0xb0, 0x8f, 0x8c, 0xef, 0x8b, 0xf6, 0x62, 0xd1, 0x89, 0x14, 0x9b, 0xd5, 0x41, 0x08,
	0x85, 0xf1, 0x11, 0xfa, 0xd3, 0x36, 0x10, 0xc2, 0x82, 0x73, 0xa4, 0x1f, 0x7d, 0x59, 0x60, 0x97,
	0x2d, 0xc2, 0xff, 0x51, 0x56, 0x33, 0xfe, 0xe9, 0x18, 0x44, 0x07, 0xaa, 0x19, 0x3f, 0x18, 0x66,
	0x68, 0xe8, 0x3c, 0xf3, 0x8d, 0x23, 0xbf, 0xa2, 0xb6, 0x3f, 0x85, 0x4d, 0xf3, 0xad, 0x04, 0xe5,
	0xa9, 0x14, 0x64, 0x98, 0x27, 0xc5, 0xb4, 0x4a, 0x5b, 0xb0, 0x12, 0x21, 0x35, 0xd6, 0xa2, 0xc1,
	0x0c, 0xe4, 0x77, 0x94, 0x51, 0x44, 0xeb, 0x40, 0xf0, 0x05, 0x22, 0x9d, 0xf5, 0x59, 0x78, 0x6e,
	0xc0, 0x79, 0xca, 0x75, 0xa3, 0x3c, 0x49, 0xf3, 0xa4, 0x18, 0x57, 0xb3, 0x2f, 0xbd, 0x6a, 0xe5,
	0xab, 0xa0, 0x6e, 0x86, 0xf1, 0x09, 0x9d, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xff, 0xbc,
	0xd9, 0x57, 0x02, 0x00, 0x00,
}
