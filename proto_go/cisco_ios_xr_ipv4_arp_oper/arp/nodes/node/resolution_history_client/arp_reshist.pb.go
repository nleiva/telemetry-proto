// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arp_reshist.proto

package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_resolution_history_client

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
	return fileDescriptor_957bd29664843221, []int{0}
}

func (m *ArpReshist_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshist_KEYS.Unmarshal(m, b)
}
func (m *ArpReshist_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshist_KEYS.Marshal(b, m, deterministic)
}
func (m *ArpReshist_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshist_KEYS.Merge(m, src)
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
	return fileDescriptor_957bd29664843221, []int{1}
}

func (m *ArpReshist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshist.Unmarshal(m, b)
}
func (m *ArpReshist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshist.Marshal(b, m, deterministic)
}
func (m *ArpReshist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshist.Merge(m, src)
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
	return fileDescriptor_957bd29664843221, []int{2}
}

func (m *ArpReshistEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpReshistEntry.Unmarshal(m, b)
}
func (m *ArpReshistEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpReshistEntry.Marshal(b, m, deterministic)
}
func (m *ArpReshistEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpReshistEntry.Merge(m, src)
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
	proto.RegisterType((*ArpReshist_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_client.arp_reshist_KEYS")
	proto.RegisterType((*ArpReshist)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_client.arp_reshist")
	proto.RegisterType((*ArpReshistEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.resolution_history_client.arp_reshist_entry")
}

func init() { proto.RegisterFile("arp_reshist.proto", fileDescriptor_957bd29664843221) }

var fileDescriptor_957bd29664843221 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x6b, 0xeb, 0x30,
	0x14, 0x85, 0x71, 0x92, 0x97, 0xe7, 0xc8, 0x2f, 0x8f, 0x44, 0x43, 0x10, 0xbc, 0xe1, 0xb9, 0x81,
	0x82, 0x87, 0xe2, 0x42, 0xda, 0xa1, 0x6b, 0x09, 0x19, 0x42, 0xa1, 0x83, 0x53, 0x0a, 0x9d, 0x84,
	0x22, 0xdf, 0x52, 0x41, 0x2c, 0xb9, 0x92, 0x5c, 0x9a, 0xb5, 0xbf, 0xab, 0x3f, 0xae, 0x5c, 0x39,
	0x69, 0xdd, 0xbd, 0x8b, 0xb1, 0xbe, 0x73, 0x74, 0xb8, 0xe7, 0x22, 0x32, 0x15, 0xb6, 0xe6, 0x16,
	0xdc, 0x93, 0x72, 0x3e, 0xaf, 0xad, 0xf1, 0x86, 0x2e, 0xa5, 0x72, 0xd2, 0x70, 0x65, 0x1c, 0x7f,
	0xb5, 0x5c, 0xd5, 0x2f, 0x97, 0x1c, 0x4d, 0xa6, 0x06, 0x9b, 0x0b, 0x5b, 0xe7, 0xda, 0x94, 0xe0,
	0xc2, 0x37, 0xb7, 0xe0, 0xcc, 0xae, 0xf1, 0xca, 0x68, 0x8e, 0x01, 0xc6, 0xee, 0xb9, 0xdc, 0x29,
	0xd0, 0x7e, 0x7e, 0x4e, 0x26, 0x9d, 0x64, 0x7e, 0xb3, 0x7a, 0xd8, 0xd0, 0x7f, 0x64, 0x84, 0xb7,
	0xb8, 0x16, 0x15, 0xb0, 0x28, 0x8d, 0xb2, 0x51, 0x11, 0x23, 0xb8, 0x15, 0x15, 0xcc, 0xdf, 0x22,
	0x92, 0x74, 0x6e, 0x50, 0x47, 0x46, 0x78, 0x04, 0xed, 0xed, 0x9e, 0x2d, 0xd2, 0x7e, 0x96, 0x2c,
	0xee, 0xf3, 0x1f, 0x98, 0x2c, 0xef, 0x8e, 0x15, 0xd2, 0x8b, 0x58, 0xd8, 0x7a, 0x85, 0x7f, 0xf3,
	0xf7, 0xde, 0xb7, 0x85, 0xb4, 0x3a, 0x3d, 0x25, 0x7f, 0xb5, 0x03, 0xc9, 0xbd, 0xaa, 0xc0, 0x79,
	0x51, 0xd5, 0x61, 0xf8, 0x41, 0x31, 0x46, 0x7a, 0x77, 0x84, 0xf4, 0x8c, 0x50, 0x55, 0x6e, 0xb9,
	0xd2, 0x1e, 0xec, 0xa3, 0x90, 0x87, 0x9e, 0xbd, 0xd0, 0x73, 0xa2, 0xca, 0xed, 0xfa, 0x28, 0x60,
	0x5f, 0x7a, 0x42, 0xfe, 0xb4, 0x05, 0xca, 0xd2, 0x82, 0x73, 0xac, 0x1f, 0x7c, 0x09, 0xb2, 0xeb,
	0x16, 0xd1, 0xff, 0x24, 0xa9, 0x84, 0xfc, 0x74, 0x0c, 0x82, 0x83, 0x54, 0x42, 0x1e, 0x0d, 0x33,
	0x32, 0x74, 0x5e, 0xf8, 0xc6, 0xb1, 0x5f, 0x41, 0x3b, 0x9c, 0x70, 0xd1, 0x6d, 0x59, 0xae, 0x4a,
	0x36, 0x4c, 0xa3, 0x6c, 0x5a, 0xc4, 0x2d, 0x58, 0x97, 0x98, 0x1a, 0x6a, 0x71, 0x34, 0x03, 0xfb,
	0x1d, 0x64, 0x12, 0xd0, 0x06, 0x09, 0xbd, 0x22, 0xac, 0xb3, 0x3d, 0x0b, 0xcf, 0x0d, 0x38, 0xcf,
	0xa5, 0x69, 0xb4, 0x67, 0x71, 0x1a, 0x65, 0xe3, 0x62, 0xf6, 0xa5, 0x17, 0xad, 0xbc, 0x44, 0x75,
	0x3b, 0x0c, 0x0f, 0xe8, 0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x26, 0x37, 0x49, 0xae, 0x55, 0x02,
	0x00, 0x00,
}
