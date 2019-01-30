// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_incomingQ_stats_entry.proto

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_incoming_queue is a generated protocol buffer package.

It is generated from these files:
	snmp_incomingQ_stats_entry.proto

It has these top-level messages:
	SnmpIncomingQStatsEntry_KEYS
	SnmpIncomingQStatsEntry
	SnmpInqDetail
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_incoming_queue

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

// SNMP IncomingQ Stats Entry
type SnmpIncomingQStatsEntry_KEYS struct {
}

func (m *SnmpIncomingQStatsEntry_KEYS) Reset()                    { *m = SnmpIncomingQStatsEntry_KEYS{} }
func (m *SnmpIncomingQStatsEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpIncomingQStatsEntry_KEYS) ProtoMessage()               {}
func (*SnmpIncomingQStatsEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SnmpIncomingQStatsEntry struct {
	// Number of NMS Queues Exist.
	QueueCount uint32 `protobuf:"varint,50,opt,name=queue_count,json=queueCount" json:"queue_count,omitempty"`
	// Each Entry Details.
	InqEntry []*SnmpInqDetail `protobuf:"bytes,51,rep,name=inq_entry,json=inqEntry" json:"inq_entry,omitempty"`
}

func (m *SnmpIncomingQStatsEntry) Reset()                    { *m = SnmpIncomingQStatsEntry{} }
func (m *SnmpIncomingQStatsEntry) String() string            { return proto.CompactTextString(m) }
func (*SnmpIncomingQStatsEntry) ProtoMessage()               {}
func (*SnmpIncomingQStatsEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpIncomingQStatsEntry) GetQueueCount() uint32 {
	if m != nil {
		return m.QueueCount
	}
	return 0
}

func (m *SnmpIncomingQStatsEntry) GetInqEntry() []*SnmpInqDetail {
	if m != nil {
		return m.InqEntry
	}
	return nil
}

// SNMP Address Information
type SnmpInqDetail struct {
	// Address of NMS Q
	AddressOfQueue string `protobuf:"bytes,1,opt,name=address_of_queue,json=addressOfQueue" json:"address_of_queue,omitempty"`
	// Request Count of Each Queue.
	RequestCount uint32 `protobuf:"varint,2,opt,name=request_count,json=requestCount" json:"request_count,omitempty"`
	// Processed request Count.
	ProcessedRequestCount uint32 `protobuf:"varint,3,opt,name=processed_request_count,json=processedRequestCount" json:"processed_request_count,omitempty"`
	// Last Access time of Each Queue.
	LastAccessTime string `protobuf:"bytes,4,opt,name=last_access_time,json=lastAccessTime" json:"last_access_time,omitempty"`
	// Priority of Each Queue.
	Priority uint32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
}

func (m *SnmpInqDetail) Reset()                    { *m = SnmpInqDetail{} }
func (m *SnmpInqDetail) String() string            { return proto.CompactTextString(m) }
func (*SnmpInqDetail) ProtoMessage()               {}
func (*SnmpInqDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SnmpInqDetail) GetAddressOfQueue() string {
	if m != nil {
		return m.AddressOfQueue
	}
	return ""
}

func (m *SnmpInqDetail) GetRequestCount() uint32 {
	if m != nil {
		return m.RequestCount
	}
	return 0
}

func (m *SnmpInqDetail) GetProcessedRequestCount() uint32 {
	if m != nil {
		return m.ProcessedRequestCount
	}
	return 0
}

func (m *SnmpInqDetail) GetLastAccessTime() string {
	if m != nil {
		return m.LastAccessTime
	}
	return ""
}

func (m *SnmpInqDetail) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpIncomingQStatsEntry_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.incoming_queue.snmp_incomingQ_stats_entry_KEYS")
	proto.RegisterType((*SnmpIncomingQStatsEntry)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.incoming_queue.snmp_incomingQ_stats_entry")
	proto.RegisterType((*SnmpInqDetail)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.incoming_queue.snmp_inq_detail")
}

func init() { proto.RegisterFile("snmp_incomingQ_stats_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x55, 0x69, 0xb7, 0x56, 0x4b, 0x40, 0x0c, 0xbd, 0x34, 0xd6, 0x4b, 0x4e, 0x39,
	0xb4, 0xe0, 0xc9, 0x8b, 0x48, 0x4f, 0x22, 0xd2, 0xe8, 0xc5, 0xd3, 0xb0, 0x26, 0x93, 0x32, 0xd2,
	0xec, 0x26, 0xbb, 0x13, 0xb0, 0x3f, 0xc9, 0x3f, 0xe5, 0x6f, 0x91, 0x6c, 0x62, 0x51, 0xa1, 0x27,
	0x8f, 0xf3, 0xbd, 0x97, 0x97, 0xf7, 0x58, 0x11, 0x5a, 0x55, 0x94, 0x40, 0x2a, 0xd5, 0x05, 0xa9,
	0xf5, 0x0a, 0x2c, 0x4b, 0xb6, 0x80, 0x8a, 0xcd, 0x36, 0x2e, 0x8d, 0x66, 0xed, 0xdf, 0xa4, 0x64,
	0x53, 0x0d, 0xa4, 0x2d, 0xbc, 0x1b, 0x70, 0x76, 0xb9, 0x46, 0xc5, 0xa0, 0x4b, 0x34, 0x71, 0x73,
	0xc7, 0xa4, 0x72, 0x6d, 0x0a, 0xc9, 0xa4, 0x55, 0xfc, 0x1d, 0x05, 0x55, 0x8d, 0x35, 0xce, 0x2e,
	0xc5, 0x74, 0xff, 0x1f, 0xe0, 0x7e, 0xf9, 0xf2, 0x34, 0xfb, 0xf0, 0xc4, 0x64, 0xbf, 0xc7, 0x9f,
	0x8a, 0xa1, 0x8b, 0x82, 0x54, 0xd7, 0x8a, 0x83, 0x79, 0xe8, 0x45, 0xa3, 0x44, 0x38, 0x74, 0xd7,
	0x10, 0xff, 0x4d, 0x0c, 0x48, 0x55, 0xad, 0x3b, 0x58, 0x84, 0xbd, 0x68, 0x38, 0x7f, 0x88, 0xff,
	0x53, 0x3a, 0xee, 0xda, 0x54, 0x90, 0x21, 0x4b, 0xda, 0x24, 0x7d, 0x52, 0xd5, 0xb2, 0x89, 0x9f,
	0x7d, 0x7a, 0xe2, 0xec, 0x8f, 0xea, 0x47, 0x62, 0x2c, 0xb3, 0xcc, 0xa0, 0xb5, 0xa0, 0xf3, 0x36,
	0x21, 0xf0, 0x42, 0x2f, 0x1a, 0x24, 0xa7, 0x1d, 0x7f, 0xcc, 0x57, 0x0d, 0xf5, 0xaf, 0xc4, 0xc8,
	0x60, 0x55, 0xa3, 0xe5, 0x6e, 0xcc, 0x81, 0x1b, 0x73, 0xd2, 0xc1, 0x76, 0xce, 0xb5, 0xb8, 0x28,
	0x8d, 0x4e, 0xd1, 0x5a, 0xcc, 0xe0, 0xb7, 0xbd, 0xe7, 0xec, 0xe7, 0x3b, 0x39, 0xf9, 0xf9, 0x5d,
	0x24, 0xc6, 0x1b, 0x69, 0x19, 0x64, 0xda, 0x88, 0xc0, 0x54, 0x60, 0x70, 0xd8, 0xd6, 0x68, 0xf8,
	0xad, 0xc3, 0xcf, 0x54, 0xa0, 0x3f, 0x11, 0xfd, 0xd2, 0x90, 0x36, 0xc4, 0xdb, 0xe0, 0xc8, 0x45,
	0xee, 0xee, 0xd7, 0x63, 0xf7, 0xe8, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xca, 0xde,
	0xc3, 0x18, 0x02, 0x00, 0x00,
}
