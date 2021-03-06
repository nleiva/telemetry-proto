// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm_mgr_show_alarm_info_cli.proto

package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_system_history

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

// alarm mgr show alarm info for CLI
type AlarmMgrShowAlarmInfoCli_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmInfoCli_KEYS) Reset()         { *m = AlarmMgrShowAlarmInfoCli_KEYS{} }
func (m *AlarmMgrShowAlarmInfoCli_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli_KEYS) ProtoMessage()    {}
func (*AlarmMgrShowAlarmInfoCli_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b, []int{0}
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.Size(m)
}
func (m *AlarmMgrShowAlarmInfoCli_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmInfoCli_KEYS proto.InternalMessageInfo

type AlarmMgrShowAlarmInfoCli struct {
	// Alarm List
	AlarmInfo            []*AlarmMgrShowAlarmData `protobuf:"bytes,50,rep,name=alarm_info,json=alarmInfo,proto3" json:"alarm_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AlarmMgrShowAlarmInfoCli) Reset()         { *m = AlarmMgrShowAlarmInfoCli{} }
func (m *AlarmMgrShowAlarmInfoCli) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli) ProtoMessage()    {}
func (*AlarmMgrShowAlarmInfoCli) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b, []int{1}
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmInfoCli) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmInfoCli.Size(m)
}
func (m *AlarmMgrShowAlarmInfoCli) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmInfoCli.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmInfoCli proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmInfoCli) GetAlarmInfo() []*AlarmMgrShowAlarmData {
	if m != nil {
		return m.AlarmInfo
	}
	return nil
}

// Alarm transport attributes
type AlarmOtn struct {
	// Alarm direction
	Direction string `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
	// Source of Alarm
	NotificationSource   string   `protobuf:"bytes,2,opt,name=notification_source,json=notificationSource,proto3" json:"notification_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmOtn) Reset()         { *m = AlarmOtn{} }
func (m *AlarmOtn) String() string { return proto.CompactTextString(m) }
func (*AlarmOtn) ProtoMessage()    {}
func (*AlarmOtn) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b, []int{2}
}
func (m *AlarmOtn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOtn.Unmarshal(m, b)
}
func (m *AlarmOtn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOtn.Marshal(b, m, deterministic)
}
func (dst *AlarmOtn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOtn.Merge(dst, src)
}
func (m *AlarmOtn) XXX_Size() int {
	return xxx_messageInfo_AlarmOtn.Size(m)
}
func (m *AlarmOtn) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOtn.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOtn proto.InternalMessageInfo

func (m *AlarmOtn) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *AlarmOtn) GetNotificationSource() string {
	if m != nil {
		return m.NotificationSource
	}
	return ""
}

// Alarm TCA Attributes
type AlarmTca struct {
	// Alarm Threshold
	ThresholdValue string `protobuf:"bytes,1,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	// Alarm Threshold
	CurrentValue string `protobuf:"bytes,2,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	// Timing Bucket
	BucketType           string   `protobuf:"bytes,3,opt,name=bucket_type,json=bucketType,proto3" json:"bucket_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmTca) Reset()         { *m = AlarmTca{} }
func (m *AlarmTca) String() string { return proto.CompactTextString(m) }
func (*AlarmTca) ProtoMessage()    {}
func (*AlarmTca) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b, []int{3}
}
func (m *AlarmTca) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmTca.Unmarshal(m, b)
}
func (m *AlarmTca) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmTca.Marshal(b, m, deterministic)
}
func (dst *AlarmTca) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmTca.Merge(dst, src)
}
func (m *AlarmTca) XXX_Size() int {
	return xxx_messageInfo_AlarmTca.Size(m)
}
func (m *AlarmTca) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmTca.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmTca proto.InternalMessageInfo

func (m *AlarmTca) GetThresholdValue() string {
	if m != nil {
		return m.ThresholdValue
	}
	return ""
}

func (m *AlarmTca) GetCurrentValue() string {
	if m != nil {
		return m.CurrentValue
	}
	return ""
}

func (m *AlarmTca) GetBucketType() string {
	if m != nil {
		return m.BucketType
	}
	return ""
}

// Alarm mgr show alarm data
type AlarmMgrShowAlarmData struct {
	// Alarm description
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Alarm location
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Alarm aid
	Aid string `protobuf:"bytes,3,opt,name=aid,proto3" json:"aid,omitempty"`
	// Alarm tag description
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	// Alarm module description
	Module string `protobuf:"bytes,5,opt,name=module,proto3" json:"module,omitempty"`
	// Alarm eid
	Eid string `protobuf:"bytes,6,opt,name=eid,proto3" json:"eid,omitempty"`
	// Reporting agent id
	ReportingAgentId uint32 `protobuf:"varint,7,opt,name=reporting_agent_id,json=reportingAgentId,proto3" json:"reporting_agent_id,omitempty"`
	// Pending async flag
	PendingSync bool `protobuf:"varint,8,opt,name=pending_sync,json=pendingSync,proto3" json:"pending_sync,omitempty"`
	// Alarm severity
	Severity string `protobuf:"bytes,9,opt,name=severity,proto3" json:"severity,omitempty"`
	// Alarm status
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// Alarm group
	Group string `protobuf:"bytes,11,opt,name=group,proto3" json:"group,omitempty"`
	// Alarm set time
	SetTime string `protobuf:"bytes,12,opt,name=set_time,json=setTime,proto3" json:"set_time,omitempty"`
	// Alarm set time(timestamp format)
	SetTimestamp uint64 `protobuf:"varint,13,opt,name=set_timestamp,json=setTimestamp,proto3" json:"set_timestamp,omitempty"`
	// Alarm clear time
	ClearTime string `protobuf:"bytes,14,opt,name=clear_time,json=clearTime,proto3" json:"clear_time,omitempty"`
	// Alarm clear time(timestamp format)
	ClearTimestamp uint64 `protobuf:"varint,15,opt,name=clear_timestamp,json=clearTimestamp,proto3" json:"clear_timestamp,omitempty"`
	// Alarm service affecting
	ServiceAffecting string `protobuf:"bytes,16,opt,name=service_affecting,json=serviceAffecting,proto3" json:"service_affecting,omitempty"`
	// OTN feature specific alarm attributes
	Otn *AlarmOtn `protobuf:"bytes,17,opt,name=otn,proto3" json:"otn,omitempty"`
	// TCA feature specific alarm attributes
	Tca *AlarmTca `protobuf:"bytes,18,opt,name=tca,proto3" json:"tca,omitempty"`
	// alarm_event_type
	Type string `protobuf:"bytes,19,opt,name=type,proto3" json:"type,omitempty"`
	// Alarm interface name
	Interface string `protobuf:"bytes,20,opt,name=interface,proto3" json:"interface,omitempty"`
	// Alarm name
	AlarmName            string   `protobuf:"bytes,21,opt,name=alarm_name,json=alarmName,proto3" json:"alarm_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmData) Reset()         { *m = AlarmMgrShowAlarmData{} }
func (m *AlarmMgrShowAlarmData) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmData) ProtoMessage()    {}
func (*AlarmMgrShowAlarmData) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b, []int{4}
}
func (m *AlarmMgrShowAlarmData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmData.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmData) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmData.Size(m)
}
func (m *AlarmMgrShowAlarmData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmData proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetReportingAgentId() uint32 {
	if m != nil {
		return m.ReportingAgentId
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetPendingSync() bool {
	if m != nil {
		return m.PendingSync
	}
	return false
}

func (m *AlarmMgrShowAlarmData) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTime() string {
	if m != nil {
		return m.SetTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTimestamp() uint64 {
	if m != nil {
		return m.SetTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetClearTime() string {
	if m != nil {
		return m.ClearTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetClearTimestamp() uint64 {
	if m != nil {
		return m.ClearTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetServiceAffecting() string {
	if m != nil {
		return m.ServiceAffecting
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetOtn() *AlarmOtn {
	if m != nil {
		return m.Otn
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetTca() *AlarmTca {
	if m != nil {
		return m.Tca
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAlarmName() string {
	if m != nil {
		return m.AlarmName
	}
	return ""
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.history.alarm_mgr_show_alarm_info_cli_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.history.alarm_mgr_show_alarm_info_cli")
	proto.RegisterType((*AlarmOtn)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.history.alarm_otn")
	proto.RegisterType((*AlarmTca)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.history.alarm_tca")
	proto.RegisterType((*AlarmMgrShowAlarmData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.history.alarm_mgr_show_alarm_data")
}

func init() {
	proto.RegisterFile("alarm_mgr_show_alarm_info_cli.proto", fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b)
}

var fileDescriptor_alarm_mgr_show_alarm_info_cli_3bd41c7f5d328f8b = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x15, 0xd2, 0xa6, 0xcd, 0x24, 0x6d, 0x53, 0xb7, 0x20, 0x17, 0x51, 0x11, 0x52, 0xa4,
	0x46, 0x02, 0x05, 0xa9, 0x3c, 0x41, 0x25, 0x38, 0x54, 0x48, 0x20, 0xa5, 0x15, 0x12, 0x5c, 0x2c,
	0xd7, 0x3b, 0xd9, 0x58, 0xec, 0xda, 0x2b, 0x7b, 0xb6, 0xb0, 0x1c, 0x78, 0x04, 0x5e, 0x81, 0x57,
	0x45, 0xb6, 0x37, 0x69, 0x0f, 0xd0, 0x53, 0x4f, 0xf1, 0x7c, 0xff, 0xaf, 0x7f, 0xc6, 0x8e, 0xbd,
	0x70, 0x22, 0x0b, 0xe9, 0x4a, 0x51, 0xe6, 0x4e, 0xf8, 0xa5, 0xfd, 0x2e, 0x52, 0xa9, 0xcd, 0xc2,
	0x0a, 0x55, 0xe8, 0x59, 0xe5, 0x2c, 0x59, 0xf6, 0x4e, 0x69, 0xaf, 0xac, 0xd0, 0xd6, 0x8b, 0x1f,
	0x2e, 0x59, 0x82, 0x1f, 0xdd, 0x0d, 0x3a, 0x61, 0x2b, 0x74, 0xb3, 0xc8, 0xfc, 0x2c, 0x43, 0x92,
	0xba, 0x68, 0x7f, 0x84, 0x6f, 0x3c, 0x61, 0x39, 0x5b, 0x6a, 0x4f, 0xd6, 0x35, 0x93, 0x97, 0x30,
	0xb9, 0xb7, 0x99, 0xf8, 0xf0, 0xfe, 0xcb, 0xe5, 0xe4, 0x4f, 0x07, 0x8e, 0xef, 0xb5, 0xb1, 0x5f,
	0x00, 0xb7, 0x84, 0x9f, 0x8d, 0xbb, 0xd3, 0xc1, 0x99, 0x98, 0x3d, 0xc4, 0x88, 0xb3, 0x7f, 0x36,
	0xce, 0x24, 0xc9, 0x79, 0x3f, 0xae, 0x2f, 0xcc, 0xc2, 0x4e, 0xbe, 0x42, 0x2a, 0x84, 0x25, 0xc3,
	0x9e, 0x41, 0x3f, 0xd3, 0x0e, 0x15, 0x69, 0x6b, 0x78, 0x67, 0xdc, 0x99, 0xf6, 0xe7, 0xb7, 0x80,
	0xbd, 0x81, 0x03, 0x63, 0x49, 0x2f, 0xb4, 0x92, 0xa1, 0x16, 0xde, 0xd6, 0x4e, 0x21, 0x7f, 0x14,
	0x7d, 0xec, 0xae, 0x74, 0x19, 0x95, 0xc9, 0xcf, 0x55, 0x36, 0x29, 0xc9, 0x4e, 0x61, 0x8f, 0x96,
	0x0e, 0xfd, 0xd2, 0x16, 0x99, 0xb8, 0x91, 0x45, 0x8d, 0x6d, 0x87, 0xdd, 0x35, 0xfe, 0x1c, 0x28,
	0x3b, 0x81, 0x1d, 0x55, 0x3b, 0x87, 0x86, 0x5a, 0x5b, 0x6a, 0x30, 0x6c, 0x61, 0x32, 0x3d, 0x87,
	0xc1, 0x75, 0xad, 0xbe, 0x21, 0x09, 0x6a, 0x2a, 0xe4, 0xdd, 0x68, 0x81, 0x84, 0xae, 0x9a, 0x0a,
	0x27, 0xbf, 0x7b, 0x70, 0xf4, 0xdf, 0x03, 0x60, 0x63, 0x18, 0x64, 0xe8, 0x95, 0xd3, 0xd5, 0x9d,
	0xad, 0xde, 0x45, 0xec, 0x29, 0x6c, 0x17, 0x36, 0xed, 0xa6, 0x1d, 0x60, 0x5d, 0xb3, 0x11, 0x74,
	0xa5, 0xce, 0xda, 0xa6, 0x61, 0x19, 0x08, 0xc9, 0x9c, 0x6f, 0x24, 0x42, 0x32, 0x67, 0x4f, 0xa0,
	0x57, 0xda, 0xac, 0x2e, 0x90, 0x6f, 0x46, 0xd8, 0x56, 0xc1, 0x89, 0x3a, 0xe3, 0xbd, 0xe4, 0x44,
	0x9d, 0xb1, 0xd7, 0xc0, 0x1c, 0x56, 0xd6, 0x91, 0x36, 0xb9, 0x90, 0x79, 0xd8, 0xb7, 0xce, 0xf8,
	0xd6, 0xb8, 0x33, 0xdd, 0x99, 0x8f, 0xd6, 0xca, 0x79, 0x10, 0x2e, 0x32, 0xf6, 0x02, 0x86, 0x15,
	0x9a, 0x2c, 0x78, 0x7d, 0x63, 0x14, 0xdf, 0x1e, 0x77, 0xa6, 0xdb, 0xf3, 0x41, 0xcb, 0x2e, 0x1b,
	0xa3, 0xc2, 0xe8, 0x1e, 0x6f, 0xd0, 0x69, 0x6a, 0x78, 0x3f, 0x8d, 0xbe, 0xaa, 0xc3, 0x58, 0x9e,
	0x24, 0xd5, 0x9e, 0x43, 0x1a, 0x2b, 0x55, 0xec, 0x10, 0x36, 0x73, 0x67, 0xeb, 0x8a, 0x0f, 0x22,
	0x4e, 0x05, 0x3b, 0x0a, 0x49, 0x24, 0x48, 0x97, 0xc8, 0x87, 0x51, 0xd8, 0xf2, 0x48, 0x57, 0xba,
	0x8c, 0xff, 0xd2, 0x4a, 0xf2, 0x24, 0xcb, 0x8a, 0xef, 0x8c, 0x3b, 0xd3, 0x8d, 0xf9, 0xb0, 0xd5,
	0x23, 0x63, 0xc7, 0x00, 0xaa, 0x40, 0xe9, 0x52, 0xc2, 0x6e, 0xba, 0x50, 0x91, 0xc4, 0x8c, 0x53,
	0xd8, 0xbb, 0x95, 0x53, 0xca, 0x5e, 0x4c, 0xd9, 0x5d, 0x7b, 0x52, 0xce, 0x2b, 0xd8, 0x0f, 0x97,
	0x5f, 0x2b, 0x14, 0x72, 0xb1, 0x08, 0xd7, 0xd1, 0xe4, 0x7c, 0x14, 0xe3, 0x46, 0xad, 0x70, 0xbe,
	0xe2, 0x4c, 0x42, 0xd7, 0x92, 0xe1, 0xfb, 0xe3, 0xce, 0x74, 0x70, 0xf6, 0xe9, 0x21, 0x9f, 0x92,
	0x25, 0x33, 0x0f, 0xd9, 0xa1, 0x05, 0x29, 0xc9, 0xd9, 0xc3, 0xb7, 0x20, 0x25, 0xe7, 0x21, 0x9b,
	0x31, 0xd8, 0x88, 0x37, 0xfb, 0x20, 0xee, 0x32, 0xae, 0xc3, 0xf3, 0xd4, 0x86, 0xd0, 0x2d, 0xa4,
	0x42, 0x7e, 0x98, 0x4e, 0x73, 0x0d, 0xc2, 0x61, 0xa7, 0x0c, 0x23, 0x4b, 0xe4, 0x8f, 0x93, 0x1c,
	0xc9, 0x47, 0x59, 0xe2, 0x75, 0x2f, 0x7e, 0xfd, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x7b, 0x23, 0x8c, 0x24, 0x05, 0x00, 0x00,
}
