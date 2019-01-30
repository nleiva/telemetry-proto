// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ip_arp_table_entry.proto

package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_entries_entry

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

// IP ARP Table entry
type IpArpTableEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArpTableEntry_KEYS) Reset()         { *m = IpArpTableEntry_KEYS{} }
func (m *IpArpTableEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpArpTableEntry_KEYS) ProtoMessage()    {}
func (*IpArpTableEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c9e11a0588998, []int{0}
}

func (m *IpArpTableEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArpTableEntry_KEYS.Unmarshal(m, b)
}
func (m *IpArpTableEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArpTableEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *IpArpTableEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArpTableEntry_KEYS.Merge(m, src)
}
func (m *IpArpTableEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpArpTableEntry_KEYS.Size(m)
}
func (m *IpArpTableEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArpTableEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpArpTableEntry_KEYS proto.InternalMessageInfo

func (m *IpArpTableEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *IpArpTableEntry_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IpArpTableEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IpArpTableEntry struct {
	// Media type for this entry
	MediaType string `protobuf:"bytes,50,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// State of this entry
	State string `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	// Flags of this entry
	Flag string `protobuf:"bytes,52,opt,name=flag,proto3" json:"flag,omitempty"`
	// Age of this entry
	Age uint64 `protobuf:"varint,53,opt,name=age,proto3" json:"age,omitempty"`
	// Source encapsulation type
	EncapsulationType string `protobuf:"bytes,54,opt,name=encapsulation_type,json=encapsulationType,proto3" json:"encapsulation_type,omitempty"`
	// Source hardware length
	HardwareLength uint32 `protobuf:"varint,55,opt,name=hardware_length,json=hardwareLength,proto3" json:"hardware_length,omitempty"`
	// Hardware address
	HardwareAddress      string   `protobuf:"bytes,56,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArpTableEntry) Reset()         { *m = IpArpTableEntry{} }
func (m *IpArpTableEntry) String() string { return proto.CompactTextString(m) }
func (*IpArpTableEntry) ProtoMessage()    {}
func (*IpArpTableEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c9e11a0588998, []int{1}
}

func (m *IpArpTableEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArpTableEntry.Unmarshal(m, b)
}
func (m *IpArpTableEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArpTableEntry.Marshal(b, m, deterministic)
}
func (m *IpArpTableEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArpTableEntry.Merge(m, src)
}
func (m *IpArpTableEntry) XXX_Size() int {
	return xxx_messageInfo_IpArpTableEntry.Size(m)
}
func (m *IpArpTableEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArpTableEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IpArpTableEntry proto.InternalMessageInfo

func (m *IpArpTableEntry) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *IpArpTableEntry) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *IpArpTableEntry) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *IpArpTableEntry) GetAge() uint64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *IpArpTableEntry) GetEncapsulationType() string {
	if m != nil {
		return m.EncapsulationType
	}
	return ""
}

func (m *IpArpTableEntry) GetHardwareLength() uint32 {
	if m != nil {
		return m.HardwareLength
	}
	return 0
}

func (m *IpArpTableEntry) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*IpArpTableEntry_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.entries.entry.ip_arp_table_entry_KEYS")
	proto.RegisterType((*IpArpTableEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.entries.entry.ip_arp_table_entry")
}

func init() { proto.RegisterFile("ip_arp_table_entry.proto", fileDescriptor_619c9e11a0588998) }

var fileDescriptor_619c9e11a0588998 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0x5b, 0xff, 0x74, 0xa0, 0xb5, 0x0e, 0x82, 0x01, 0x11, 0x4a, 0x41, 0xac, 0x07,
	0xf7, 0x60, 0xab, 0xf5, 0xea, 0xc1, 0x93, 0xe2, 0xa1, 0x7a, 0xf1, 0x14, 0xa6, 0xbb, 0xd3, 0x36,
	0xb0, 0xdd, 0x84, 0x24, 0x5a, 0xf7, 0x6b, 0xfb, 0x09, 0x64, 0xa7, 0xb6, 0x20, 0xbd, 0x24, 0x33,
	0xbf, 0x99, 0xf7, 0x1e, 0x24, 0xa0, 0x8c, 0xd3, 0xe4, 0x9d, 0x8e, 0x34, 0x2d, 0x58, 0x73, 0x19,
	0x7d, 0x95, 0x3a, 0x6f, 0xa3, 0xc5, 0x71, 0x66, 0x42, 0x66, 0xb5, 0xb1, 0x41, 0x7f, 0x7b, 0x6d,
	0xdc, 0xd7, 0x48, 0x16, 0xad, 0x63, 0x9f, 0x92, 0x77, 0x69, 0x69, 0x73, 0x0e, 0x72, 0xa6, 0xb5,
	0xcc, 0x70, 0x90, 0xbb, 0xea, 0xaf, 0xe0, 0x6c, 0xd7, 0x54, 0x3f, 0x3f, 0x7d, 0xbc, 0xe1, 0x39,
	0xb4, 0x6a, 0x81, 0x2e, 0x69, 0xc9, 0x2a, 0xe9, 0x25, 0x83, 0xd6, 0xe4, 0xa8, 0x06, 0xaf, 0xb4,
	0x64, 0x54, 0x70, 0x48, 0x79, 0xee, 0x39, 0x04, 0xb5, 0x27, 0xa3, 0x4d, 0x8b, 0x97, 0xd0, 0x31,
	0x65, 0x64, 0x3f, 0xa3, 0xec, 0x4f, 0xdb, 0x90, 0x85, 0xf6, 0x96, 0xd6, 0x06, 0xfd, 0x9f, 0x04,
	0x70, 0x37, 0x19, 0x2f, 0x00, 0x96, 0x9c, 0x1b, 0xd2, 0xb1, 0x72, 0xac, 0x6e, 0x45, 0xd9, 0x12,
	0xf2, 0x5e, 0x39, 0xc6, 0x53, 0xd8, 0x0f, 0x91, 0x22, 0xab, 0xa1, 0x4c, 0xd6, 0x0d, 0x22, 0x34,
	0x67, 0x05, 0xcd, 0xd5, 0x48, 0xa0, 0xd4, 0xd8, 0x85, 0x06, 0xcd, 0x59, 0xdd, 0xf5, 0x92, 0x41,
	0x73, 0x52, 0x97, 0x78, 0x03, 0xc8, 0x65, 0x46, 0x2e, 0x7c, 0x16, 0x14, 0x8d, 0x2d, 0xd7, 0x11,
	0xf7, 0xa2, 0x39, 0xf9, 0x37, 0x91, 0xa8, 0x2b, 0x38, 0x5e, 0x90, 0xcf, 0x57, 0xe4, 0x59, 0x17,
	0x5c, 0xce, 0xe3, 0x42, 0x8d, 0x7b, 0xc9, 0xa0, 0x3d, 0xe9, 0x6c, 0xf0, 0x8b, 0x50, 0xbc, 0x86,
	0xee, 0x76, 0x71, 0xf3, 0x26, 0x0f, 0xe2, 0xba, 0x35, 0x78, 0x5c, 0xe3, 0xe9, 0x81, 0xfc, 0xd6,
	0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xd9, 0xbe, 0x7b, 0xc9, 0x01, 0x00, 0x00,
}
