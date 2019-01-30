// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_request_list.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_areas_area_requests_request is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_request_list.proto

It has these top-level messages:
	OspfShRequestList_KEYS
	OspfShRequestList
	OspfShLsaSum
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_areas_area_requests_request

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

// OSPF Request List Information
type OspfShRequestList_KEYS struct {
	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	AreaId          uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	InterfaceName   string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *OspfShRequestList_KEYS) Reset()                    { *m = OspfShRequestList_KEYS{} }
func (m *OspfShRequestList_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShRequestList_KEYS) ProtoMessage()               {}
func (*OspfShRequestList_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShRequestList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRequestList_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShRequestList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShRequestList_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShRequestList struct {
	// Neighbor ID
	RequestNeighborId string `protobuf:"bytes,50,opt,name=request_neighbor_id,json=requestNeighborId" json:"request_neighbor_id,omitempty"`
	// Neighbor IP address
	RequestNeighborAddress string `protobuf:"bytes,51,opt,name=request_neighbor_address,json=requestNeighborAddress" json:"request_neighbor_address,omitempty"`
	// Request list interface
	RequestInterfaceName string `protobuf:"bytes,52,opt,name=request_interface_name,json=requestInterfaceName" json:"request_interface_name,omitempty"`
	// List of request list entries
	RequestList []*OspfShLsaSum `protobuf:"bytes,53,rep,name=request_list,json=requestList" json:"request_list,omitempty"`
}

func (m *OspfShRequestList) Reset()                    { *m = OspfShRequestList{} }
func (m *OspfShRequestList) String() string            { return proto.CompactTextString(m) }
func (*OspfShRequestList) ProtoMessage()               {}
func (*OspfShRequestList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShRequestList) GetRequestNeighborId() string {
	if m != nil {
		return m.RequestNeighborId
	}
	return ""
}

func (m *OspfShRequestList) GetRequestNeighborAddress() string {
	if m != nil {
		return m.RequestNeighborAddress
	}
	return ""
}

func (m *OspfShRequestList) GetRequestInterfaceName() string {
	if m != nil {
		return m.RequestInterfaceName
	}
	return ""
}

func (m *OspfShRequestList) GetRequestList() []*OspfShLsaSum {
	if m != nil {
		return m.RequestList
	}
	return nil
}

// LSA Summary Entry
type OspfShLsaSum struct {
	// LSA Type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType" json:"header_lsa_type,omitempty"`
	// Age of the LSA (s)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsId string `protobuf:"bytes,3,opt,name=header_ls_id,json=headerLsId" json:"header_ls_id,omitempty"`
	// Router ID of the Advertising Router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber uint32 `protobuf:"varint,5,opt,name=header_sequence_number,json=headerSequenceNumber" json:"header_sequence_number,omitempty"`
	// Checksum of the LSA
	HeaderLsaChecksum uint32 `protobuf:"varint,6,opt,name=header_lsa_checksum,json=headerLsaChecksum" json:"header_lsa_checksum,omitempty"`
}

func (m *OspfShLsaSum) Reset()                    { *m = OspfShLsaSum{} }
func (m *OspfShLsaSum) String() string            { return proto.CompactTextString(m) }
func (*OspfShLsaSum) ProtoMessage()               {}
func (*OspfShLsaSum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	proto.RegisterType((*OspfShRequestList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.requests.request.ospf_sh_request_list_KEYS")
	proto.RegisterType((*OspfShRequestList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.requests.request.ospf_sh_request_list")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.requests.request.ospf_sh_lsa_sum")
}

func init() { proto.RegisterFile("ospf_sh_request_list.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x8f, 0x12, 0x31,
	0x18, 0xce, 0xb0, 0x8a, 0xf1, 0x05, 0x16, 0xb7, 0x12, 0x77, 0xd6, 0x13, 0x12, 0x35, 0x78, 0x99,
	0xc3, 0x2e, 0x26, 0xc6, 0x1b, 0x31, 0x1e, 0x88, 0x1b, 0x0e, 0xac, 0x17, 0xbd, 0x34, 0x65, 0xfa,
	0x02, 0xcd, 0x32, 0x1f, 0xf6, 0xed, 0x10, 0xf7, 0x4f, 0xf8, 0x27, 0x3c, 0xf9, 0x4b, 0xfc, 0x5b,
	0xa6, 0x9d, 0x76, 0x44, 0xe4, 0xba, 0x17, 0xd2, 0x3e, 0x1f, 0x6f, 0x9f, 0xf2, 0x74, 0xe0, 0x79,
	0x41, 0xe5, 0x8a, 0xd3, 0x86, 0x6b, 0xfc, 0x56, 0x21, 0x19, 0xbe, 0x55, 0x64, 0x92, 0x52, 0x17,
	0xa6, 0x60, 0x5f, 0x53, 0x45, 0x69, 0xc1, 0x55, 0x41, 0xfc, 0xbb, 0xe6, 0xaa, 0xdc, 0x4d, 0xb8,
	0x53, 0x17, 0x25, 0xea, 0xc4, 0xae, 0xac, 0x2e, 0x45, 0x22, 0xa4, 0xb0, 0x4a, 0x24, 0xae, 0x44,
	0xb5, 0x35, 0x7c, 0xa7, 0x57, 0x89, 0xd0, 0x28, 0xc8, 0xfd, 0x26, 0x7e, 0x3a, 0x85, 0xc5, 0xe8,
	0x57, 0x04, 0x17, 0xc7, 0x8e, 0xe6, 0x9f, 0x3e, 0x7e, 0xb9, 0x61, 0x2f, 0xa0, 0xeb, 0x07, 0xf2,
	0x5c, 0x64, 0x18, 0x47, 0xc3, 0x68, 0xfc, 0x78, 0xd1, 0xf1, 0xd8, 0x5c, 0x64, 0xc8, 0xce, 0xe1,
	0x91, 0x9d, 0xcc, 0x95, 0x8c, 0x5b, 0xc3, 0x68, 0xdc, 0x5b, 0xb4, 0xed, 0x76, 0x26, 0xd9, 0x2b,
	0x38, 0x55, 0xb9, 0x41, 0xbd, 0x12, 0x29, 0xd6, 0xee, 0x13, 0xe7, 0xee, 0x35, 0xa8, 0xf3, 0xbf,
	0x81, 0x27, 0x39, 0xaa, 0xf5, 0x66, 0x59, 0x68, 0x2e, 0xa4, 0xd4, 0x48, 0x14, 0x3f, 0x70, 0xc2,
	0x7e, 0xc0, 0xa7, 0x35, 0x3c, 0xfa, 0xdd, 0x82, 0xc1, 0xb1, 0xac, 0x2c, 0x81, 0xa7, 0x61, 0xdf,
	0xcc, 0x52, 0x32, 0xbe, 0x74, 0x63, 0xce, 0x3c, 0x35, 0xf7, 0xcc, 0x4c, 0xb2, 0x77, 0x10, 0xff,
	0xa7, 0x0f, 0x67, 0x5f, 0x39, 0xd3, 0xb3, 0x03, 0x93, 0x8f, 0xc0, 0x26, 0x10, 0x18, 0x7e, 0x70,
	0xb9, 0x89, 0xf3, 0x0d, 0x3c, 0x3b, 0xfb, 0xe7, 0x8e, 0x3f, 0x22, 0xe8, 0xee, 0x07, 0x8e, 0xdf,
	0x0e, 0x4f, 0xc6, 0x9d, 0xcb, 0xdb, 0xe4, 0xfe, 0x8a, 0x4d, 0xc2, 0x1f, 0xb5, 0x25, 0xc1, 0xa9,
	0xca, 0x16, 0x1d, 0x4f, 0x5c, 0x2b, 0x32, 0xa3, 0x9f, 0x2d, 0xe8, 0x1f, 0x08, 0xd8, 0x6b, 0xe8,
	0x6f, 0x50, 0x48, 0xd4, 0x0e, 0x31, 0x77, 0x65, 0xa8, 0xbb, 0x57, 0xc3, 0xd7, 0x24, 0x3e, 0xdf,
	0x95, 0xc8, 0x5e, 0xc2, 0xe9, 0x9e, 0x4e, 0xac, 0xd1, 0xf7, 0xde, 0x6d, 0x64, 0xd3, 0x35, 0xb2,
	0x21, 0x74, 0x1b, 0x95, 0xed, 0xa2, 0xee, 0x1e, 0x82, 0x66, 0x26, 0xd9, 0x7b, 0xb8, 0xf0, 0x0a,
	0x21, 0x77, 0xa8, 0x8d, 0x22, 0x95, 0xaf, 0xb9, 0x2e, 0x2a, 0x83, 0xda, 0xbf, 0x80, 0xf3, 0x5a,
	0x30, 0xfd, 0xcb, 0x2f, 0x1c, 0x6d, 0x6b, 0xf0, 0x5e, 0xb2, 0xb7, 0xca, 0x6d, 0x09, 0x55, 0xb6,
	0x44, 0x1d, 0x3f, 0x74, 0x59, 0x06, 0x35, 0x7b, 0xe3, 0xc9, 0xb9, 0xe3, 0xec, 0x33, 0xd9, 0x4b,
	0x9e, 0x6e, 0x30, 0xbd, 0xa5, 0x2a, 0x8b, 0xdb, 0xce, 0x72, 0xd6, 0xc4, 0xff, 0xe0, 0x89, 0x65,
	0xdb, 0x7d, 0x7e, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xe1, 0xa6, 0xdb, 0x9c, 0x03,
	0x00, 0x00,
}
