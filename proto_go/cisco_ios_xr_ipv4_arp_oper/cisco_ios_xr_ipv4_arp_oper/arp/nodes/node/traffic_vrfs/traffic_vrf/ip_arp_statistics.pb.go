// Code generated by protoc-gen-go.
// source: ip_arp_statistics.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_traffic_vrfs_traffic_vrf is a generated protocol buffer package.

It is generated from these files:
	ip_arp_statistics.proto

It has these top-level messages:
	IpArpStatistics_KEYS
	IpArpStatistics
*/
package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_traffic_vrfs_traffic_vrf

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

// IP ARP Statistics information
type IpArpStatistics_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	VrfName  string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *IpArpStatistics_KEYS) Reset()                    { *m = IpArpStatistics_KEYS{} }
func (m *IpArpStatistics_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IpArpStatistics_KEYS) ProtoMessage()               {}
func (*IpArpStatistics_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IpArpStatistics_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *IpArpStatistics_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IpArpStatistics struct {
	// Total ARP requests received
	RequestsReceived uint32 `protobuf:"varint,50,opt,name=requests_received,json=requestsReceived" json:"requests_received,omitempty"`
	// Total ARP replies received
	RepliesReceived uint32 `protobuf:"varint,51,opt,name=replies_received,json=repliesReceived" json:"replies_received,omitempty"`
	// Total ARP requests sent
	RequestsSent uint32 `protobuf:"varint,52,opt,name=requests_sent,json=requestsSent" json:"requests_sent,omitempty"`
	// Total ARP replies sent
	RepliesSent uint32 `protobuf:"varint,53,opt,name=replies_sent,json=repliesSent" json:"replies_sent,omitempty"`
	// Total Proxy ARP replies sent
	ProxyRepliesSent uint32 `protobuf:"varint,54,opt,name=proxy_replies_sent,json=proxyRepliesSent" json:"proxy_replies_sent,omitempty"`
	// Total ARP requests received over subscriber interface
	SubscrRequestsReceived uint32 `protobuf:"varint,55,opt,name=subscr_requests_received,json=subscrRequestsReceived" json:"subscr_requests_received,omitempty"`
	// Total ARP replies sent over subscriber interface
	SubscrRepliesSent uint32 `protobuf:"varint,56,opt,name=subscr_replies_sent,json=subscrRepliesSent" json:"subscr_replies_sent,omitempty"`
	// Total ARP grat replies sent over subscriber interface
	SubscrRepliesGratgSent uint32 `protobuf:"varint,57,opt,name=subscr_replies_gratg_sent,json=subscrRepliesGratgSent" json:"subscr_replies_gratg_sent,omitempty"`
	// Total Local Proxy ARP replies sent
	LocalProxyRepliesSent uint32 `protobuf:"varint,58,opt,name=local_proxy_replies_sent,json=localProxyRepliesSent" json:"local_proxy_replies_sent,omitempty"`
	// Total Gratuituous ARP replies sent
	GratuitousRepliesSent uint32 `protobuf:"varint,59,opt,name=gratuitous_replies_sent,json=gratuitousRepliesSent" json:"gratuitous_replies_sent,omitempty"`
	// Total ARP resolution requests received
	ResolutionRequestsReceived uint32 `protobuf:"varint,60,opt,name=resolution_requests_received,json=resolutionRequestsReceived" json:"resolution_requests_received,omitempty"`
	// Total ARP resolution replies received
	ResolutionRepliesReceived uint32 `protobuf:"varint,61,opt,name=resolution_replies_received,json=resolutionRepliesReceived" json:"resolution_replies_received,omitempty"`
	// total ARP resolution requests dropped
	ResolutionRequestsDropped uint32 `protobuf:"varint,62,opt,name=resolution_requests_dropped,json=resolutionRequestsDropped" json:"resolution_requests_dropped,omitempty"`
	// Total errors for out of memory
	OutOfMemoryErrors uint32 `protobuf:"varint,63,opt,name=out_of_memory_errors,json=outOfMemoryErrors" json:"out_of_memory_errors,omitempty"`
	// Total errors for no buffer
	NoBufferErrors uint32 `protobuf:"varint,64,opt,name=no_buffer_errors,json=noBufferErrors" json:"no_buffer_errors,omitempty"`
	// Total ARP entries in the cache
	TotalEntries uint32 `protobuf:"varint,65,opt,name=total_entries,json=totalEntries" json:"total_entries,omitempty"`
	// Total dynamic entries in the cache
	DynamicEntries uint32 `protobuf:"varint,66,opt,name=dynamic_entries,json=dynamicEntries" json:"dynamic_entries,omitempty"`
	// Total static entries in the cache
	StaticEntries uint32 `protobuf:"varint,67,opt,name=static_entries,json=staticEntries" json:"static_entries,omitempty"`
	// Total alias entries in the cache
	AliasEntries uint32 `protobuf:"varint,68,opt,name=alias_entries,json=aliasEntries" json:"alias_entries,omitempty"`
	// Total interface entries in the cache
	InterfaceEntries uint32 `protobuf:"varint,69,opt,name=interface_entries,json=interfaceEntries" json:"interface_entries,omitempty"`
	// Total standby entries in the cache
	StandbyEntries uint32 `protobuf:"varint,70,opt,name=standby_entries,json=standbyEntries" json:"standby_entries,omitempty"`
	// Total DHCP entries in the cache
	DhcpEntries uint32 `protobuf:"varint,71,opt,name=dhcp_entries,json=dhcpEntries" json:"dhcp_entries,omitempty"`
	// Total VXLAN entries in the cache
	VxlanEntries uint32 `protobuf:"varint,72,opt,name=vxlan_entries,json=vxlanEntries" json:"vxlan_entries,omitempty"`
	// Total ip packets droped on this node
	IpPacketsDroppedNode uint32 `protobuf:"varint,73,opt,name=ip_packets_dropped_node,json=ipPacketsDroppedNode" json:"ip_packets_dropped_node,omitempty"`
	// Total ARP packets on node due to out of subnet
	ArpPacketNodeOutOfSubnet uint32 `protobuf:"varint,74,opt,name=arp_packet_node_out_of_subnet,json=arpPacketNodeOutOfSubnet" json:"arp_packet_node_out_of_subnet,omitempty"`
	// Total ip packets droped on this interface
	IpPacketsDroppedInterface uint32 `protobuf:"varint,75,opt,name=ip_packets_dropped_interface,json=ipPacketsDroppedInterface" json:"ip_packets_dropped_interface,omitempty"`
	// Total arp packets on interface due to out of subnet
	ArpPacketInterfaceOutOfSubnet uint32 `protobuf:"varint,76,opt,name=arp_packet_interface_out_of_subnet,json=arpPacketInterfaceOutOfSubnet" json:"arp_packet_interface_out_of_subnet,omitempty"`
	// Total unsolicited arp packets dropped
	ArpPacketUnsolicitedPacket uint32 `protobuf:"varint,77,opt,name=arp_packet_unsolicited_packet,json=arpPacketUnsolicitedPacket" json:"arp_packet_unsolicited_packet,omitempty"`
	// Total idb structures on this node
	IdbStructures uint32 `protobuf:"varint,78,opt,name=idb_structures,json=idbStructures" json:"idb_structures,omitempty"`
}

func (m *IpArpStatistics) Reset()                    { *m = IpArpStatistics{} }
func (m *IpArpStatistics) String() string            { return proto.CompactTextString(m) }
func (*IpArpStatistics) ProtoMessage()               {}
func (*IpArpStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IpArpStatistics) GetRequestsReceived() uint32 {
	if m != nil {
		return m.RequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetRepliesReceived() uint32 {
	if m != nil {
		return m.RepliesReceived
	}
	return 0
}

func (m *IpArpStatistics) GetRequestsSent() uint32 {
	if m != nil {
		return m.RequestsSent
	}
	return 0
}

func (m *IpArpStatistics) GetRepliesSent() uint32 {
	if m != nil {
		return m.RepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetProxyRepliesSent() uint32 {
	if m != nil {
		return m.ProxyRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRequestsReceived() uint32 {
	if m != nil {
		return m.SubscrRequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRepliesSent() uint32 {
	if m != nil {
		return m.SubscrRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRepliesGratgSent() uint32 {
	if m != nil {
		return m.SubscrRepliesGratgSent
	}
	return 0
}

func (m *IpArpStatistics) GetLocalProxyRepliesSent() uint32 {
	if m != nil {
		return m.LocalProxyRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetGratuitousRepliesSent() uint32 {
	if m != nil {
		return m.GratuitousRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRequestsReceived() uint32 {
	if m != nil {
		return m.ResolutionRequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRepliesReceived() uint32 {
	if m != nil {
		return m.ResolutionRepliesReceived
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRequestsDropped() uint32 {
	if m != nil {
		return m.ResolutionRequestsDropped
	}
	return 0
}

func (m *IpArpStatistics) GetOutOfMemoryErrors() uint32 {
	if m != nil {
		return m.OutOfMemoryErrors
	}
	return 0
}

func (m *IpArpStatistics) GetNoBufferErrors() uint32 {
	if m != nil {
		return m.NoBufferErrors
	}
	return 0
}

func (m *IpArpStatistics) GetTotalEntries() uint32 {
	if m != nil {
		return m.TotalEntries
	}
	return 0
}

func (m *IpArpStatistics) GetDynamicEntries() uint32 {
	if m != nil {
		return m.DynamicEntries
	}
	return 0
}

func (m *IpArpStatistics) GetStaticEntries() uint32 {
	if m != nil {
		return m.StaticEntries
	}
	return 0
}

func (m *IpArpStatistics) GetAliasEntries() uint32 {
	if m != nil {
		return m.AliasEntries
	}
	return 0
}

func (m *IpArpStatistics) GetInterfaceEntries() uint32 {
	if m != nil {
		return m.InterfaceEntries
	}
	return 0
}

func (m *IpArpStatistics) GetStandbyEntries() uint32 {
	if m != nil {
		return m.StandbyEntries
	}
	return 0
}

func (m *IpArpStatistics) GetDhcpEntries() uint32 {
	if m != nil {
		return m.DhcpEntries
	}
	return 0
}

func (m *IpArpStatistics) GetVxlanEntries() uint32 {
	if m != nil {
		return m.VxlanEntries
	}
	return 0
}

func (m *IpArpStatistics) GetIpPacketsDroppedNode() uint32 {
	if m != nil {
		return m.IpPacketsDroppedNode
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketNodeOutOfSubnet() uint32 {
	if m != nil {
		return m.ArpPacketNodeOutOfSubnet
	}
	return 0
}

func (m *IpArpStatistics) GetIpPacketsDroppedInterface() uint32 {
	if m != nil {
		return m.IpPacketsDroppedInterface
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketInterfaceOutOfSubnet() uint32 {
	if m != nil {
		return m.ArpPacketInterfaceOutOfSubnet
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketUnsolicitedPacket() uint32 {
	if m != nil {
		return m.ArpPacketUnsolicitedPacket
	}
	return 0
}

func (m *IpArpStatistics) GetIdbStructures() uint32 {
	if m != nil {
		return m.IdbStructures
	}
	return 0
}

func init() {
	proto.RegisterType((*IpArpStatistics_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.traffic_vrfs.traffic_vrf.ip_arp_statistics_KEYS")
	proto.RegisterType((*IpArpStatistics)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.traffic_vrfs.traffic_vrf.ip_arp_statistics")
}

func init() { proto.RegisterFile("ip_arp_statistics.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0xdb, 0x4f, 0xdb, 0x30,
	0x14, 0xc6, 0xc5, 0x1e, 0x36, 0x38, 0x50, 0xa0, 0x19, 0x83, 0x30, 0x40, 0xe2, 0xa2, 0x69, 0x4c,
	0x4c, 0x9d, 0x34, 0xee, 0xbb, 0x70, 0x1b, 0x1d, 0x63, 0x8c, 0x8b, 0x5a, 0xed, 0x61, 0x4f, 0x96,
	0x9b, 0x38, 0xcc, 0x5a, 0x6a, 0x67, 0xb6, 0x53, 0xd1, 0x7f, 0x6d, 0x7f, 0xdd, 0x94, 0xe3, 0xc4,
	0x4d, 0x5a, 0x5e, 0x90, 0xf8, 0xbe, 0xdf, 0x77, 0xce, 0x71, 0x7c, 0x92, 0xc2, 0x02, 0x4f, 0x08,
	0x55, 0x09, 0xd1, 0x86, 0x1a, 0xae, 0x0d, 0x0f, 0x74, 0x23, 0x51, 0xd2, 0x48, 0xef, 0x2c, 0xe0,
	0x3a, 0x90, 0x84, 0x4b, 0x4d, 0x1e, 0x14, 0xe1, 0x49, 0x6f, 0x07, 0x39, 0x99, 0x30, 0xd5, 0xa0,
	0x2a, 0x69, 0x08, 0x19, 0x32, 0x8d, 0x7f, 0x1b, 0x46, 0xd1, 0x28, 0xe2, 0x01, 0xe9, 0xa9, 0x48,
	0x97, 0xff, 0x59, 0xbf, 0x83, 0xf9, 0x91, 0xf2, 0xe4, 0xaa, 0xf9, 0xab, 0xed, 0x2d, 0xc1, 0x44,
	0x16, 0x25, 0x82, 0x76, 0x99, 0x3f, 0xb6, 0x3a, 0xb6, 0x39, 0xd1, 0x1a, 0xcf, 0x84, 0x1b, 0xda,
	0x65, 0xde, 0x22, 0x8c, 0xf7, 0x54, 0x64, 0xbd, 0x27, 0xe8, 0x3d, 0xeb, 0xa9, 0x28, 0xb3, 0xd6,
	0xff, 0x4d, 0x42, 0x7d, 0xa4, 0xa4, 0xb7, 0x05, 0x75, 0xc5, 0xfe, 0xa6, 0x4c, 0x1b, 0x4d, 0x14,
	0x0b, 0x18, 0xef, 0xb1, 0xd0, 0x7f, 0xbf, 0x3a, 0xb6, 0x59, 0x6b, 0xcd, 0x16, 0x46, 0x2b, 0xd7,
	0xbd, 0x37, 0x30, 0xab, 0x58, 0x12, 0x73, 0x56, 0x62, 0xb7, 0x91, 0x9d, 0xc9, 0x75, 0x87, 0x6e,
	0x40, 0xcd, 0xd5, 0xd5, 0x4c, 0x18, 0x7f, 0x07, 0xb9, 0xa9, 0x42, 0x6c, 0x33, 0x61, 0xbc, 0x35,
	0x98, 0x2a, 0xea, 0x21, 0xb3, 0x8b, 0xcc, 0x64, 0xae, 0x21, 0xf2, 0x16, 0xbc, 0x44, 0xc9, 0x87,
	0x3e, 0xa9, 0x80, 0x7b, 0x76, 0x40, 0x74, 0x5a, 0x25, 0xfa, 0x00, 0x7c, 0x9d, 0x76, 0x74, 0xa0,
	0xc8, 0xe8, 0xa1, 0xf6, 0x31, 0x33, 0x6f, 0xfd, 0xd6, 0xf0, 0xd1, 0x1a, 0xf0, 0xdc, 0x25, 0x4b,
	0x8d, 0x0e, 0x30, 0x54, 0x2f, 0x42, 0x83, 0x4e, 0x87, 0xb0, 0x38, 0xc4, 0xdf, 0x2b, 0x6a, 0xee,
	0x6d, 0xea, 0xb0, 0xda, 0x0a, 0xfd, 0x8b, 0xcc, 0xc6, 0xe8, 0x3e, 0xf8, 0xb1, 0x0c, 0x68, 0x4c,
	0x1e, 0x39, 0xd8, 0x07, 0x4c, 0xbe, 0x40, 0xff, 0x6e, 0xf8, 0x74, 0x7b, 0xb0, 0x90, 0x35, 0x49,
	0xb9, 0x91, 0xa9, 0xae, 0xe6, 0x3e, 0xda, 0xdc, 0xc0, 0x2e, 0xe7, 0x4e, 0x60, 0x59, 0x31, 0x2d,
	0xe3, 0xd4, 0x70, 0x29, 0x1e, 0x79, 0x32, 0x9f, 0x30, 0xfc, 0x72, 0xc0, 0x8c, 0x3c, 0x9d, 0x23,
	0x58, 0xaa, 0x54, 0x18, 0xda, 0x81, 0xcf, 0x58, 0x60, 0xb1, 0x5c, 0xa0, 0xba, 0x0d, 0xc3, 0xf9,
	0x7c, 0x82, 0x50, 0xc9, 0x24, 0x61, 0xa1, 0x7f, 0x34, 0x9a, 0xb7, 0xc4, 0xb9, 0x05, 0xbc, 0x77,
	0x30, 0x27, 0x53, 0x43, 0x64, 0x44, 0xba, 0xac, 0x2b, 0x55, 0x9f, 0x30, 0xa5, 0xa4, 0xd2, 0xfe,
	0xb1, 0xbd, 0x1e, 0x99, 0x9a, 0xdb, 0xe8, 0x1a, 0x9d, 0x26, 0x1a, 0xde, 0x26, 0xcc, 0x0a, 0x49,
	0x3a, 0x69, 0x14, 0x31, 0x55, 0xc0, 0x27, 0x08, 0x4f, 0x0b, 0x79, 0x86, 0x72, 0x4e, 0x6e, 0x40,
	0xcd, 0x48, 0x43, 0x63, 0xc2, 0x84, 0x51, 0x9c, 0x69, 0xff, 0xd4, 0x2e, 0x2a, 0x8a, 0x4d, 0xab,
	0x79, 0xaf, 0x61, 0x26, 0xec, 0x0b, 0xda, 0xe5, 0x81, 0xc3, 0xce, 0x6c, 0xb5, 0x5c, 0x2e, 0xc0,
	0x57, 0x30, 0x8d, 0x2f, 0xd7, 0x80, 0xfb, 0x82, 0x5c, 0xcd, 0xaa, 0x05, 0xb6, 0x01, 0x35, 0x1a,
	0x73, 0xaa, 0x1d, 0x75, 0x6e, 0x9b, 0xa2, 0x58, 0x40, 0x5b, 0x50, 0xe7, 0xc2, 0x30, 0x15, 0xd1,
	0x80, 0x39, 0xb0, 0x69, 0x37, 0xdf, 0x19, 0xa5, 0x09, 0xb5, 0xa1, 0x22, 0xec, 0xf4, 0x1d, 0xfa,
	0xd5, 0x4e, 0x98, 0xcb, 0x05, 0xb8, 0x06, 0x53, 0xe1, 0xef, 0x20, 0x71, 0xd4, 0x85, 0x7d, 0xe7,
	0x32, 0xad, 0x34, 0x5d, 0xef, 0x21, 0xa6, 0xc2, 0x31, 0xdf, 0xec, 0x74, 0x28, 0x16, 0xd0, 0x2e,
	0x7e, 0xff, 0x12, 0x1a, 0xfc, 0x61, 0x83, 0x9b, 0x24, 0xd9, 0x87, 0xc8, 0xbf, 0x44, 0x7c, 0x8e,
	0x27, 0x77, 0xd6, 0xcd, 0x6f, 0xf1, 0x46, 0x86, 0xcc, 0x3b, 0x86, 0x95, 0xec, 0x0b, 0x64, 0x73,
	0x88, 0x93, 0xfc, 0x66, 0x75, 0xda, 0x11, 0xcc, 0xf8, 0xdf, 0x31, 0xec, 0x53, 0x95, 0xa7, 0xb3,
	0xd4, 0x6d, 0x76, 0xbf, 0x6d, 0xf4, 0xbd, 0x63, 0x58, 0x7e, 0xa4, 0xaf, 0x7b, 0x1e, 0xfe, 0x95,
	0xdd, 0xa5, 0xe1, 0xe6, 0x97, 0x05, 0xe0, 0x5d, 0xc2, 0x7a, 0x69, 0x82, 0xc1, 0x13, 0xae, 0x8e,
	0xf1, 0x03, 0xcb, 0xac, 0xb8, 0x31, 0x5c, 0xbe, 0x3c, 0xcb, 0x69, 0xe5, 0x30, 0xa9, 0xd0, 0x32,
	0xe6, 0x01, 0x37, 0x2c, 0xcc, 0x25, 0xff, 0xda, 0xbe, 0x59, 0xae, 0xca, 0xcf, 0x01, 0x62, 0x85,
	0x6c, 0x61, 0x78, 0xd8, 0x21, 0xda, 0xa8, 0x34, 0x30, 0xa9, 0x62, 0xda, 0xbf, 0xb1, 0x0b, 0xc3,
	0xc3, 0x4e, 0xdb, 0x89, 0x9d, 0xa7, 0xf8, 0xcb, 0xb2, 0xfd, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x84,
	0x37, 0x09, 0x43, 0x74, 0x06, 0x00, 0x00,
}