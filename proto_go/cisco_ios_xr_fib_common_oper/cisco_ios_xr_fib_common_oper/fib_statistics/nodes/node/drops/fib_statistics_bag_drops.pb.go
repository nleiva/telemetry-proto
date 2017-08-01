// Code generated by protoc-gen-go.
// source: fib_statistics_bag_drops.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_fib_common_oper_fib_statistics_nodes_node_drops is a generated protocol buffer package.

It is generated from these files:
	fib_statistics_bag_drops.proto

It has these top-level messages:
	FibStatisticsBagDrops_KEYS
	FibStatisticsBagDrops
*/
package cisco_ios_xr_fib_common_oper_fib_statistics_nodes_node_drops

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

// FIB drop statistics
type FibStatisticsBagDrops_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *FibStatisticsBagDrops_KEYS) Reset()                    { *m = FibStatisticsBagDrops_KEYS{} }
func (m *FibStatisticsBagDrops_KEYS) String() string            { return proto.CompactTextString(m) }
func (*FibStatisticsBagDrops_KEYS) ProtoMessage()               {}
func (*FibStatisticsBagDrops_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FibStatisticsBagDrops_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type FibStatisticsBagDrops struct {
	// no route pkt
	NoRoutePackets uint64 `protobuf:"varint,50,opt,name=no_route_packets,json=noRoutePackets" json:"no_route_packets,omitempty"`
	// Punt generate unreach pkt
	PuntUnreachablePackets uint64 `protobuf:"varint,51,opt,name=punt_unreachable_packets,json=puntUnreachablePackets" json:"punt_unreachable_packets,omitempty"`
	// DF unreachable pkt
	DfUnreachablePackets uint64 `protobuf:"varint,52,opt,name=df_unreachable_packets,json=dfUnreachablePackets" json:"df_unreachable_packets,omitempty"`
	// encapsulation failure pkt
	EncapsulationFailurePackets uint64 `protobuf:"varint,53,opt,name=encapsulation_failure_packets,json=encapsulationFailurePackets" json:"encapsulation_failure_packets,omitempty"`
	// incomplete adjacency pkt
	IncompleteAdjacencyPackets uint64 `protobuf:"varint,54,opt,name=incomplete_adjacency_packets,json=incompleteAdjacencyPackets" json:"incomplete_adjacency_packets,omitempty"`
	// unresolved prefix pkt
	UnresolvedPrefixPackets uint64 `protobuf:"varint,55,opt,name=unresolved_prefix_packets,json=unresolvedPrefixPackets" json:"unresolved_prefix_packets,omitempty"`
	// unsupported feature pkt
	UnsupportedFeaturePackets uint64 `protobuf:"varint,56,opt,name=unsupported_feature_packets,json=unsupportedFeaturePackets" json:"unsupported_feature_packets,omitempty"`
	// discard pkt
	DiscardPackets uint64 `protobuf:"varint,57,opt,name=discard_packets,json=discardPackets" json:"discard_packets,omitempty"`
	// checksum error pkt
	ChecksumErrorPackets uint64 `protobuf:"varint,58,opt,name=checksum_error_packets,json=checksumErrorPackets" json:"checksum_error_packets,omitempty"`
	// frag consumed packet pkt
	FragmenationConsumedPackets uint64 `protobuf:"varint,59,opt,name=fragmenation_consumed_packets,json=fragmenationConsumedPackets" json:"fragmenation_consumed_packets,omitempty"`
	// fragmenation failure pkt
	FragmenationFailurePackets uint64 `protobuf:"varint,60,opt,name=fragmenation_failure_packets,json=fragmenationFailurePackets" json:"fragmenation_failure_packets,omitempty"`
	// null0 pkt
	NullPackets uint64 `protobuf:"varint,61,opt,name=null_packets,json=nullPackets" json:"null_packets,omitempty"`
	// RPF check failures pkt
	RpfCheckFailurePackets uint64 `protobuf:"varint,62,opt,name=rpf_check_failure_packets,json=rpfCheckFailurePackets" json:"rpf_check_failure_packets,omitempty"`
	// ACL in RPF pkt
	AclInRpfPackets uint64 `protobuf:"varint,63,opt,name=acl_in_rpf_packets,json=aclInRpfPackets" json:"acl_in_rpf_packets,omitempty"`
	// rp dest drop pkt
	RpDestinationDropPackets uint64 `protobuf:"varint,64,opt,name=rp_destination_drop_packets,json=rpDestinationDropPackets" json:"rp_destination_drop_packets,omitempty"`
	// the total number of drop pkt
	TotalNumberOfDropPackets uint64 `protobuf:"varint,65,opt,name=total_number_of_drop_packets,json=totalNumberOfDropPackets" json:"total_number_of_drop_packets,omitempty"`
	// mpls disabled in interface
	MplsDisabledInterface uint64 `protobuf:"varint,66,opt,name=mpls_disabled_interface,json=mplsDisabledInterface" json:"mpls_disabled_interface,omitempty"`
	// GRE tunnel lookup failed drop pkt
	GreLookupFailedDrop uint64 `protobuf:"varint,67,opt,name=gre_lookup_failed_drop,json=greLookupFailedDrop" json:"gre_lookup_failed_drop,omitempty"`
	// GRE processing errors
	GreErrorDrop uint64 `protobuf:"varint,68,opt,name=gre_error_drop,json=greErrorDrop" json:"gre_error_drop,omitempty"`
	// LISP Punt drops
	LispPuntDrops uint64 `protobuf:"varint,69,opt,name=lisp_punt_drops,json=lispPuntDrops" json:"lisp_punt_drops,omitempty"`
	// Lisp encap error drops
	LispEncapErrorDrops uint64 `protobuf:"varint,70,opt,name=lisp_encap_error_drops,json=lispEncapErrorDrops" json:"lisp_encap_error_drops,omitempty"`
	// Lisp decap error drops
	LispDecapErrorDrops uint64 `protobuf:"varint,71,opt,name=lisp_decap_error_drops,json=lispDecapErrorDrops" json:"lisp_decap_error_drops,omitempty"`
	// Drops for the packets with multi[le labels
	MultiLabelDrops uint64 `protobuf:"varint,72,opt,name=multi_label_drops,json=multiLabelDrops" json:"multi_label_drops,omitempty"`
}

func (m *FibStatisticsBagDrops) Reset()                    { *m = FibStatisticsBagDrops{} }
func (m *FibStatisticsBagDrops) String() string            { return proto.CompactTextString(m) }
func (*FibStatisticsBagDrops) ProtoMessage()               {}
func (*FibStatisticsBagDrops) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FibStatisticsBagDrops) GetNoRoutePackets() uint64 {
	if m != nil {
		return m.NoRoutePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetPuntUnreachablePackets() uint64 {
	if m != nil {
		return m.PuntUnreachablePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetDfUnreachablePackets() uint64 {
	if m != nil {
		return m.DfUnreachablePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetEncapsulationFailurePackets() uint64 {
	if m != nil {
		return m.EncapsulationFailurePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetIncompleteAdjacencyPackets() uint64 {
	if m != nil {
		return m.IncompleteAdjacencyPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetUnresolvedPrefixPackets() uint64 {
	if m != nil {
		return m.UnresolvedPrefixPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetUnsupportedFeaturePackets() uint64 {
	if m != nil {
		return m.UnsupportedFeaturePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetDiscardPackets() uint64 {
	if m != nil {
		return m.DiscardPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetChecksumErrorPackets() uint64 {
	if m != nil {
		return m.ChecksumErrorPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetFragmenationConsumedPackets() uint64 {
	if m != nil {
		return m.FragmenationConsumedPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetFragmenationFailurePackets() uint64 {
	if m != nil {
		return m.FragmenationFailurePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetNullPackets() uint64 {
	if m != nil {
		return m.NullPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetRpfCheckFailurePackets() uint64 {
	if m != nil {
		return m.RpfCheckFailurePackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetAclInRpfPackets() uint64 {
	if m != nil {
		return m.AclInRpfPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetRpDestinationDropPackets() uint64 {
	if m != nil {
		return m.RpDestinationDropPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetTotalNumberOfDropPackets() uint64 {
	if m != nil {
		return m.TotalNumberOfDropPackets
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetMplsDisabledInterface() uint64 {
	if m != nil {
		return m.MplsDisabledInterface
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetGreLookupFailedDrop() uint64 {
	if m != nil {
		return m.GreLookupFailedDrop
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetGreErrorDrop() uint64 {
	if m != nil {
		return m.GreErrorDrop
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetLispPuntDrops() uint64 {
	if m != nil {
		return m.LispPuntDrops
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetLispEncapErrorDrops() uint64 {
	if m != nil {
		return m.LispEncapErrorDrops
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetLispDecapErrorDrops() uint64 {
	if m != nil {
		return m.LispDecapErrorDrops
	}
	return 0
}

func (m *FibStatisticsBagDrops) GetMultiLabelDrops() uint64 {
	if m != nil {
		return m.MultiLabelDrops
	}
	return 0
}

func init() {
	proto.RegisterType((*FibStatisticsBagDrops_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib_statistics.nodes.node.drops.fib_statistics_bag_drops_KEYS")
	proto.RegisterType((*FibStatisticsBagDrops)(nil), "cisco_ios_xr_fib_common_oper.fib_statistics.nodes.node.drops.fib_statistics_bag_drops")
}

func init() { proto.RegisterFile("fib_statistics_bag_drops.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x4f, 0x13, 0x4b,
	0x14, 0xc7, 0x43, 0x72, 0x73, 0x73, 0x99, 0x8b, 0xa0, 0x55, 0xcb, 0x62, 0xc1, 0x20, 0x31, 0x4a,
	0x34, 0xe9, 0x83, 0x45, 0x04, 0x04, 0xe4, 0x47, 0x5b, 0x25, 0x12, 0x24, 0x35, 0x3e, 0xf8, 0x34,
	0x99, 0xce, 0x9e, 0x2d, 0x23, 0xb3, 0x33, 0x93, 0x99, 0x59, 0x83, 0xff, 0xb8, 0xcf, 0x66, 0xce,
	0x74, 0xbb, 0xdb, 0x0a, 0x2f, 0x7d, 0x38, 0xdf, 0xcf, 0xe7, 0xcc, 0xf6, 0xcc, 0x0f, 0xf2, 0x34,
	0x13, 0x43, 0xea, 0x3c, 0xf3, 0xc2, 0x79, 0xc1, 0x1d, 0x1d, 0xb2, 0x11, 0x4d, 0xad, 0x36, 0xae,
	0x6d, 0xac, 0xf6, 0xba, 0xb1, 0xcf, 0x85, 0xe3, 0x9a, 0x0a, 0xed, 0xe8, 0x8d, 0xa5, 0x01, 0xe6,
	0x3a, 0xcf, 0xb5, 0xa2, 0xda, 0x80, 0x6d, 0x4f, 0xcb, 0x6d, 0xa5, 0x53, 0x88, 0xbf, 0x6d, 0xec,
	0xb1, 0xb1, 0x4f, 0xd6, 0xee, 0xea, 0x4f, 0x3f, 0xf7, 0xbe, 0x7f, 0x6d, 0xb4, 0xc8, 0x7c, 0xc0,
	0xa9, 0x62, 0x39, 0x24, 0x73, 0xeb, 0x73, 0x9b, 0xf3, 0x83, 0xff, 0x42, 0xe1, 0x82, 0xe5, 0xb0,
	0xf1, 0x7b, 0x9e, 0x24, 0x77, 0xe9, 0x8d, 0x4d, 0x72, 0x5f, 0x69, 0x6a, 0x75, 0xe1, 0x81, 0x1a,
	0xc6, 0xaf, 0xc1, 0xbb, 0xe4, 0xcd, 0xfa, 0xdc, 0xe6, 0x3f, 0x83, 0x45, 0xa5, 0x07, 0xa1, 0x7c,
	0x19, 0xab, 0x8d, 0x1d, 0x92, 0x98, 0x42, 0x79, 0x5a, 0x28, 0x0b, 0x8c, 0x5f, 0xb1, 0xa1, 0xac,
	0x8c, 0x0e, 0x1a, 0xcd, 0x90, 0x7f, 0xab, 0xe2, 0xd2, 0xdc, 0x22, 0xcd, 0x34, 0xbb, 0xd5, 0xdb,
	0x42, 0xef, 0x51, 0x9a, 0xdd, 0x62, 0x9d, 0x90, 0x35, 0x50, 0x9c, 0x19, 0x57, 0x48, 0xe6, 0x85,
	0x56, 0x34, 0x63, 0x42, 0x16, 0xb6, 0x92, 0xdf, 0xa2, 0xdc, 0x9a, 0x82, 0xfa, 0x91, 0x29, 0x7b,
	0x1c, 0x91, 0x55, 0xa1, 0xb8, 0xce, 0x8d, 0x04, 0x0f, 0x94, 0xa5, 0x3f, 0x18, 0x07, 0xc5, 0x7f,
	0x4d, 0x5a, 0x6c, 0x63, 0x8b, 0x27, 0x15, 0x73, 0x5c, 0x22, 0x65, 0x87, 0x3d, 0xb2, 0x12, 0x3e,
	0xdc, 0x69, 0xf9, 0x13, 0x52, 0x6a, 0x2c, 0x64, 0xe2, 0x66, 0xa2, 0xbf, 0x43, 0x7d, 0xb9, 0x02,
	0x2e, 0x31, 0x2f, 0xdd, 0x43, 0xd2, 0x2a, 0x94, 0x2b, 0x8c, 0xd1, 0xd6, 0x43, 0x4a, 0x33, 0x60,
	0xbe, 0xfe, 0xfd, 0x3b, 0x68, 0xaf, 0xd4, 0x90, 0x7e, 0x24, 0x4a, 0xff, 0x25, 0x59, 0x4a, 0x85,
	0xe3, 0xcc, 0xa6, 0x13, 0x67, 0x37, 0x6e, 0xcd, 0xb8, 0x5c, 0x1b, 0x30, 0xbf, 0x02, 0x7e, 0xed,
	0x8a, 0x9c, 0x82, 0xb5, 0xda, 0x4e, 0xf8, 0xbd, 0x38, 0xe0, 0x32, 0xed, 0x85, 0xb0, 0x36, 0xe0,
	0xcc, 0xb2, 0x51, 0x0e, 0x2a, 0xce, 0x97, 0x6b, 0xe5, 0x8a, 0x1c, 0xaa, 0xc5, 0xde, 0xc7, 0x01,
	0xd7, 0xa1, 0xd3, 0x31, 0x53, 0x1b, 0xf0, 0x54, 0x8f, 0xd9, 0x3d, 0xda, 0x8f, 0x03, 0xae, 0x33,
	0x33, 0x5b, 0xf4, 0x8c, 0x2c, 0xa8, 0x42, 0xca, 0x89, 0x71, 0x80, 0xc6, 0xff, 0xa1, 0x56, 0x22,
	0xbb, 0x64, 0xc5, 0x9a, 0x8c, 0xe2, 0x9f, 0xf8, 0x6b, 0x85, 0xc3, 0x78, 0xf4, 0xac, 0xc9, 0x4e,
	0x43, 0x3e, 0xd3, 0xfd, 0x35, 0x69, 0x30, 0x2e, 0xa9, 0x50, 0x34, 0x74, 0x28, 0x9d, 0x0f, 0xe8,
	0x2c, 0x31, 0x2e, 0xcf, 0xd4, 0xc0, 0x64, 0x25, 0x7c, 0x40, 0x5a, 0xd6, 0xd0, 0x14, 0x9c, 0x17,
	0xe3, 0xbf, 0x13, 0xee, 0xc8, 0xc4, 0x3a, 0x42, 0x2b, 0xb1, 0xa6, 0x5b, 0x11, 0x5d, 0xab, 0x4d,
	0xb5, 0xdd, 0xab, 0x5e, 0x7b, 0x26, 0xa9, 0x2a, 0xf2, 0x21, 0x58, 0xaa, 0xb3, 0x69, 0xff, 0x38,
	0xfa, 0xc8, 0x5c, 0x20, 0xf2, 0x25, 0xab, 0xfb, 0xdb, 0x64, 0x39, 0x37, 0xd2, 0xd1, 0x54, 0xb8,
	0x70, 0x0f, 0x52, 0x2a, 0x94, 0x07, 0x9b, 0x31, 0x0e, 0xc9, 0x09, 0xaa, 0x8f, 0x43, 0xdc, 0x1d,
	0xa7, 0x67, 0x65, 0xd8, 0xe8, 0x90, 0xe6, 0xc8, 0x02, 0x95, 0x5a, 0x5f, 0x17, 0x06, 0xe7, 0x03,
	0x29, 0xae, 0x9c, 0x9c, 0xa2, 0xf6, 0x70, 0x64, 0xe1, 0x1c, 0xc3, 0x3e, 0x66, 0x61, 0xcd, 0xc6,
	0x73, 0xb2, 0x18, 0xa4, 0x78, 0x5a, 0x10, 0xee, 0x22, 0xbc, 0x30, 0xb2, 0x80, 0xa7, 0x04, 0xa9,
	0x17, 0x64, 0x49, 0x0a, 0x67, 0x28, 0x5e, 0x7c, 0x7c, 0x30, 0x92, 0x1e, 0x62, 0xf7, 0x42, 0xf9,
	0xb2, 0x50, 0xbe, 0x8b, 0xaf, 0x48, 0x87, 0x34, 0x91, 0xc3, 0xbb, 0x58, 0x6b, 0xea, 0x92, 0x7e,
	0xfc, 0x84, 0x90, 0xf6, 0x42, 0x38, 0xe9, 0x5d, 0x49, 0x29, 0xcc, 0x4a, 0x1f, 0x2b, 0xa9, 0x0b,
	0xd3, 0xd2, 0x2b, 0xf2, 0x20, 0x2f, 0xa4, 0x17, 0x54, 0xb2, 0x21, 0xc8, 0x31, 0xff, 0x29, 0xee,
	0x27, 0x06, 0xe7, 0xa1, 0x8e, 0xec, 0xf0, 0x5f, 0x7c, 0x7b, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x81, 0xe6, 0x6f, 0x1e, 0x9d, 0x05, 0x00, 0x00,
}