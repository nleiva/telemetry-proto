// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_statistics.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_duplicate_drop

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

// SNMP Statistics extension
type SnmpStatistics_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpStatistics_KEYS) Reset()         { *m = SnmpStatistics_KEYS{} }
func (m *SnmpStatistics_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpStatistics_KEYS) ProtoMessage()    {}
func (*SnmpStatistics_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_statistics_b9d493ef5dec4040, []int{0}
}
func (m *SnmpStatistics_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpStatistics_KEYS.Unmarshal(m, b)
}
func (m *SnmpStatistics_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpStatistics_KEYS.Marshal(b, m, deterministic)
}
func (dst *SnmpStatistics_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpStatistics_KEYS.Merge(dst, src)
}
func (m *SnmpStatistics_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpStatistics_KEYS.Size(m)
}
func (m *SnmpStatistics_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpStatistics_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpStatistics_KEYS proto.InternalMessageInfo

type SnmpStatistics struct {
	// Duplicate requests drop feature status.
	DuplicateRequestStatus string `protobuf:"bytes,50,opt,name=duplicate_request_status,json=duplicateRequestStatus,proto3" json:"duplicate_request_status,omitempty"`
	// Duplicate request drop feature last enable disable time (Day Mon Date HH:MM:SS)
	LastStatusChangeTime string `protobuf:"bytes,51,opt,name=last_status_change_time,json=lastStatusChangeTime,proto3" json:"last_status_change_time,omitempty"`
	// Configured Duplicate Drop feature Timeout.
	DuplicateDropConfiguredTimeout uint32 `protobuf:"varint,52,opt,name=duplicate_drop_configured_timeout,json=duplicateDropConfiguredTimeout,proto3" json:"duplicate_drop_configured_timeout,omitempty"`
	// Number of duplicate SNMP requests are dropped.
	DuplicateDroppedRequests uint32 `protobuf:"varint,53,opt,name=duplicate_dropped_requests,json=duplicateDroppedRequests,proto3" json:"duplicate_dropped_requests,omitempty"`
	// Number of Retry SNMP requests are Processed.
	RetryProcessedRequests uint32 `protobuf:"varint,54,opt,name=retry_processed_requests,json=retryProcessedRequests,proto3" json:"retry_processed_requests,omitempty"`
	// Duplicate request drop feature first  enable time (Day Mon Date HH:MM:SS)
	FirstEnableTime string `protobuf:"bytes,55,opt,name=first_enable_time,json=firstEnableTime,proto3" json:"first_enable_time,omitempty"`
	// Number of duplicate SNMP requests dropped, from the last enable time.
	LatestDuplicateDroppedRequests uint32 `protobuf:"varint,56,opt,name=latest_duplicate_dropped_requests,json=latestDuplicateDroppedRequests,proto3" json:"latest_duplicate_dropped_requests,omitempty"`
	// Number of retry SNMP requests processed, from the last enable time.
	LatestRetryProcessedRequests uint32 `protobuf:"varint,57,opt,name=latest_retry_processed_requests,json=latestRetryProcessedRequests,proto3" json:"latest_retry_processed_requests,omitempty"`
	// Duplicate request drop feature last enable time(Day Mon Date HH:MM:SS)
	DuplicateRequestLatestEnableTime string `protobuf:"bytes,58,opt,name=duplicate_request_latest_enable_time,json=duplicateRequestLatestEnableTime,proto3" json:"duplicate_request_latest_enable_time,omitempty"`
	//  Number of times duplicate request drop feature is enabled.
	DuplicateDropEnableCount uint32 `protobuf:"varint,59,opt,name=duplicate_drop_enable_count,json=duplicateDropEnableCount,proto3" json:"duplicate_drop_enable_count,omitempty"`
	//  Number of times duplicate request drop feature is disabled.
	DuplicateDropDisableCount uint32   `protobuf:"varint,60,opt,name=duplicate_drop_disable_count,json=duplicateDropDisableCount,proto3" json:"duplicate_drop_disable_count,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *SnmpStatistics) Reset()         { *m = SnmpStatistics{} }
func (m *SnmpStatistics) String() string { return proto.CompactTextString(m) }
func (*SnmpStatistics) ProtoMessage()    {}
func (*SnmpStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_statistics_b9d493ef5dec4040, []int{1}
}
func (m *SnmpStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpStatistics.Unmarshal(m, b)
}
func (m *SnmpStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpStatistics.Marshal(b, m, deterministic)
}
func (dst *SnmpStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpStatistics.Merge(dst, src)
}
func (m *SnmpStatistics) XXX_Size() int {
	return xxx_messageInfo_SnmpStatistics.Size(m)
}
func (m *SnmpStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpStatistics proto.InternalMessageInfo

func (m *SnmpStatistics) GetDuplicateRequestStatus() string {
	if m != nil {
		return m.DuplicateRequestStatus
	}
	return ""
}

func (m *SnmpStatistics) GetLastStatusChangeTime() string {
	if m != nil {
		return m.LastStatusChangeTime
	}
	return ""
}

func (m *SnmpStatistics) GetDuplicateDropConfiguredTimeout() uint32 {
	if m != nil {
		return m.DuplicateDropConfiguredTimeout
	}
	return 0
}

func (m *SnmpStatistics) GetDuplicateDroppedRequests() uint32 {
	if m != nil {
		return m.DuplicateDroppedRequests
	}
	return 0
}

func (m *SnmpStatistics) GetRetryProcessedRequests() uint32 {
	if m != nil {
		return m.RetryProcessedRequests
	}
	return 0
}

func (m *SnmpStatistics) GetFirstEnableTime() string {
	if m != nil {
		return m.FirstEnableTime
	}
	return ""
}

func (m *SnmpStatistics) GetLatestDuplicateDroppedRequests() uint32 {
	if m != nil {
		return m.LatestDuplicateDroppedRequests
	}
	return 0
}

func (m *SnmpStatistics) GetLatestRetryProcessedRequests() uint32 {
	if m != nil {
		return m.LatestRetryProcessedRequests
	}
	return 0
}

func (m *SnmpStatistics) GetDuplicateRequestLatestEnableTime() string {
	if m != nil {
		return m.DuplicateRequestLatestEnableTime
	}
	return ""
}

func (m *SnmpStatistics) GetDuplicateDropEnableCount() uint32 {
	if m != nil {
		return m.DuplicateDropEnableCount
	}
	return 0
}

func (m *SnmpStatistics) GetDuplicateDropDisableCount() uint32 {
	if m != nil {
		return m.DuplicateDropDisableCount
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpStatistics_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.duplicate_drop.snmp_statistics_KEYS")
	proto.RegisterType((*SnmpStatistics)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.duplicate_drop.snmp_statistics")
}

func init() {
	proto.RegisterFile("snmp_statistics.proto", fileDescriptor_snmp_statistics_b9d493ef5dec4040)
}

var fileDescriptor_snmp_statistics_b9d493ef5dec4040 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x4b, 0xe3, 0x40,
	0x1c, 0xc5, 0xe9, 0x61, 0x17, 0x76, 0x60, 0x29, 0x1b, 0xba, 0xdd, 0xec, 0x6e, 0xd1, 0x5a, 0x3c,
	0x14, 0x0f, 0x39, 0x58, 0xab, 0x55, 0x2b, 0x1e, 0xda, 0x1e, 0x44, 0x11, 0x49, 0x7b, 0xf1, 0x34,
	0x4c, 0x27, 0xd3, 0x3a, 0x90, 0xcc, 0x8c, 0x33, 0x13, 0xd0, 0xcf, 0xe9, 0x17, 0x92, 0xfc, 0x27,
	0x8d, 0x49, 0x6c, 0x3d, 0x26, 0xef, 0xfd, 0x1e, 0xef, 0xe5, 0x4f, 0xd0, 0x6f, 0x23, 0x12, 0x85,
	0x8d, 0x25, 0x96, 0x1b, 0xcb, 0xa9, 0x09, 0x94, 0x96, 0x56, 0x7a, 0x63, 0xca, 0x0d, 0x95, 0x98,
	0x4b, 0x83, 0x5f, 0x34, 0x06, 0x0f, 0x59, 0x33, 0x61, 0xb1, 0x54, 0x4c, 0x07, 0xd9, 0x73, 0xc0,
	0xc5, 0x4a, 0xea, 0x84, 0x58, 0x2e, 0x45, 0x10, 0xa5, 0x2a, 0xe6, 0x94, 0x58, 0x86, 0x23, 0x2d,
	0x55, 0xaf, 0x8d, 0x5a, 0xb5, 0x58, 0x7c, 0x3b, 0x7b, 0x9c, 0xf7, 0xde, 0xbe, 0xa1, 0x66, 0x4d,
	0xf0, 0x46, 0xc8, 0xff, 0xa0, 0x35, 0x7b, 0x4e, 0x99, 0xb1, 0xa0, 0xa7, 0xc6, 0x3f, 0xee, 0x36,
	0xfa, 0x3f, 0xc2, 0x76, 0xa1, 0x87, 0x4e, 0x9e, 0x83, 0xea, 0x0d, 0xd1, 0x9f, 0x98, 0x14, 0x66,
	0x4c, 0x9f, 0x88, 0x58, 0x33, 0x6c, 0x79, 0xc2, 0xfc, 0x01, 0x80, 0xad, 0x4c, 0x76, 0xe6, 0x09,
	0x88, 0x0b, 0x9e, 0x30, 0xef, 0x06, 0x1d, 0x54, 0xeb, 0x62, 0x2a, 0xc5, 0x8a, 0xaf, 0x53, 0xcd,
	0x22, 0xa0, 0x65, 0x6a, 0xfd, 0x93, 0x6e, 0xa3, 0xff, 0x33, 0xdc, 0x2b, 0x8c, 0x53, 0x2d, 0xd5,
	0xa4, 0xb0, 0x2d, 0x9c, 0xcb, 0x1b, 0xa3, 0x7f, 0xd5, 0x28, 0xc5, 0xa2, 0xcd, 0x06, 0xe3, 0x0f,
	0x21, 0xc3, 0xaf, 0x64, 0x28, 0x16, 0xe5, 0x23, 0x60, 0xb9, 0x66, 0x56, 0xbf, 0x62, 0xa5, 0x25,
	0x65, 0xc6, 0x94, 0xd9, 0x53, 0x60, 0xdb, 0xa0, 0x3f, 0x6c, 0xe4, 0x82, 0x3c, 0x42, 0xbf, 0x56,
	0x5c, 0x1b, 0x8b, 0x99, 0x20, 0xcb, 0x38, 0xdf, 0x7c, 0x06, 0x9b, 0x9b, 0x20, 0xcc, 0xe0, 0xfd,
	0x66, 0x6e, 0x4c, 0x6c, 0xf6, 0x51, 0xbf, 0xa8, 0x3a, 0x72, 0x73, 0x9d, 0x71, 0xba, 0xab, 0xf0,
	0x0c, 0xed, 0xe7, 0x51, 0x3b, 0x7b, 0x9f, 0x43, 0x50, 0xc7, 0xd9, 0xc2, 0xed, 0xed, 0xef, 0xd1,
	0xe1, 0xe7, 0x8b, 0xe7, 0xc1, 0xe5, 0x41, 0x17, 0x30, 0xa8, 0x5b, 0xbf, 0xfe, 0x1d, 0x38, 0x4b,
	0x0b, 0xaf, 0xd0, 0xff, 0xda, 0x41, 0xf3, 0x14, 0x2a, 0x53, 0x61, 0xfd, 0xcb, 0x2d, 0x67, 0x70,
	0xf4, 0x24, 0xd3, 0xbd, 0x6b, 0xd4, 0xa9, 0xe1, 0x11, 0x37, 0x25, 0x7e, 0x0c, 0xfc, 0xdf, 0x0a,
	0x3f, 0x75, 0x0e, 0x08, 0x58, 0x7e, 0x87, 0x5f, 0x66, 0xf0, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x6a, 0x9f, 0xb8, 0x4b, 0x03, 0x00, 0x00,
}
