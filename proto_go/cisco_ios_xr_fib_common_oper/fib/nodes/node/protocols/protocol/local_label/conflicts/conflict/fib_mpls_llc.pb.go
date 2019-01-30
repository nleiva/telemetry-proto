// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fib_mpls_llc.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_local_label_conflicts_conflict

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

type FibMplsLlc_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	Label                uint32   `protobuf:"varint,3,opt,name=label,proto3" json:"label,omitempty"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	LlcType              string   `protobuf:"bytes,5,opt,name=llc_type,json=llcType,proto3" json:"llc_type,omitempty"`
	PfxTblId             uint32   `protobuf:"varint,6,opt,name=pfx_tbl_id,json=pfxTblId,proto3" json:"pfx_tbl_id,omitempty"`
	Prefix               string   `protobuf:"bytes,7,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLen            uint32   `protobuf:"varint,8,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibMplsLlc_KEYS) Reset()         { *m = FibMplsLlc_KEYS{} }
func (m *FibMplsLlc_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibMplsLlc_KEYS) ProtoMessage()    {}
func (*FibMplsLlc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_mpls_llc_01c81bf59f53383b, []int{0}
}
func (m *FibMplsLlc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsLlc_KEYS.Unmarshal(m, b)
}
func (m *FibMplsLlc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsLlc_KEYS.Marshal(b, m, deterministic)
}
func (dst *FibMplsLlc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsLlc_KEYS.Merge(dst, src)
}
func (m *FibMplsLlc_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibMplsLlc_KEYS.Size(m)
}
func (m *FibMplsLlc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsLlc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsLlc_KEYS proto.InternalMessageInfo

func (m *FibMplsLlc_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibMplsLlc_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibMplsLlc_KEYS) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *FibMplsLlc_KEYS) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *FibMplsLlc_KEYS) GetLlcType() string {
	if m != nil {
		return m.LlcType
	}
	return ""
}

func (m *FibMplsLlc_KEYS) GetPfxTblId() uint32 {
	if m != nil {
		return m.PfxTblId
	}
	return 0
}

func (m *FibMplsLlc_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *FibMplsLlc_KEYS) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

type FibMplsLlc struct {
	LocalLabel           uint32             `protobuf:"varint,50,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	Source               uint32             `protobuf:"varint,51,opt,name=source,proto3" json:"source,omitempty"`
	UpdateTs             uint64             `protobuf:"varint,52,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	RetryTs              uint64             `protobuf:"varint,53,opt,name=retry_ts,json=retryTs,proto3" json:"retry_ts,omitempty"`
	NumRetries           uint32             `protobuf:"varint,54,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"`
	Ext                  *FibMplsLlcTypeExt `protobuf:"bytes,55,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FibMplsLlc) Reset()         { *m = FibMplsLlc{} }
func (m *FibMplsLlc) String() string { return proto.CompactTextString(m) }
func (*FibMplsLlc) ProtoMessage()    {}
func (*FibMplsLlc) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_mpls_llc_01c81bf59f53383b, []int{1}
}
func (m *FibMplsLlc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsLlc.Unmarshal(m, b)
}
func (m *FibMplsLlc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsLlc.Marshal(b, m, deterministic)
}
func (dst *FibMplsLlc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsLlc.Merge(dst, src)
}
func (m *FibMplsLlc) XXX_Size() int {
	return xxx_messageInfo_FibMplsLlc.Size(m)
}
func (m *FibMplsLlc) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsLlc.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsLlc proto.InternalMessageInfo

func (m *FibMplsLlc) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *FibMplsLlc) GetSource() uint32 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *FibMplsLlc) GetUpdateTs() uint64 {
	if m != nil {
		return m.UpdateTs
	}
	return 0
}

func (m *FibMplsLlc) GetRetryTs() uint64 {
	if m != nil {
		return m.RetryTs
	}
	return 0
}

func (m *FibMplsLlc) GetNumRetries() uint32 {
	if m != nil {
		return m.NumRetries
	}
	return 0
}

func (m *FibMplsLlc) GetExt() *FibMplsLlcTypeExt {
	if m != nil {
		return m.Ext
	}
	return nil
}

type FibMplsLlcPfx struct {
	Pfx                  string   `protobuf:"bytes,1,opt,name=pfx,proto3" json:"pfx,omitempty"`
	TblId                uint32   `protobuf:"varint,2,opt,name=tbl_id,json=tblId,proto3" json:"tbl_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibMplsLlcPfx) Reset()         { *m = FibMplsLlcPfx{} }
func (m *FibMplsLlcPfx) String() string { return proto.CompactTextString(m) }
func (*FibMplsLlcPfx) ProtoMessage()    {}
func (*FibMplsLlcPfx) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_mpls_llc_01c81bf59f53383b, []int{2}
}
func (m *FibMplsLlcPfx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsLlcPfx.Unmarshal(m, b)
}
func (m *FibMplsLlcPfx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsLlcPfx.Marshal(b, m, deterministic)
}
func (dst *FibMplsLlcPfx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsLlcPfx.Merge(dst, src)
}
func (m *FibMplsLlcPfx) XXX_Size() int {
	return xxx_messageInfo_FibMplsLlcPfx.Size(m)
}
func (m *FibMplsLlcPfx) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsLlcPfx.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsLlcPfx proto.InternalMessageInfo

func (m *FibMplsLlcPfx) GetPfx() string {
	if m != nil {
		return m.Pfx
	}
	return ""
}

func (m *FibMplsLlcPfx) GetTblId() uint32 {
	if m != nil {
		return m.TblId
	}
	return 0
}

type FibMplsLlcLsm struct {
	Nh                   string   `protobuf:"bytes,1,opt,name=nh,proto3" json:"nh,omitempty"`
	McastId              uint32   `protobuf:"varint,2,opt,name=mcast_id,json=mcastId,proto3" json:"mcast_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibMplsLlcLsm) Reset()         { *m = FibMplsLlcLsm{} }
func (m *FibMplsLlcLsm) String() string { return proto.CompactTextString(m) }
func (*FibMplsLlcLsm) ProtoMessage()    {}
func (*FibMplsLlcLsm) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_mpls_llc_01c81bf59f53383b, []int{3}
}
func (m *FibMplsLlcLsm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsLlcLsm.Unmarshal(m, b)
}
func (m *FibMplsLlcLsm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsLlcLsm.Marshal(b, m, deterministic)
}
func (dst *FibMplsLlcLsm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsLlcLsm.Merge(dst, src)
}
func (m *FibMplsLlcLsm) XXX_Size() int {
	return xxx_messageInfo_FibMplsLlcLsm.Size(m)
}
func (m *FibMplsLlcLsm) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsLlcLsm.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsLlcLsm proto.InternalMessageInfo

func (m *FibMplsLlcLsm) GetNh() string {
	if m != nil {
		return m.Nh
	}
	return ""
}

func (m *FibMplsLlcLsm) GetMcastId() uint32 {
	if m != nil {
		return m.McastId
	}
	return 0
}

type FibMplsLlcTypeExt struct {
	Type                 string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Pfx                  *FibMplsLlcPfx `protobuf:"bytes,2,opt,name=pfx,proto3" json:"pfx,omitempty"`
	Lsm                  *FibMplsLlcLsm `protobuf:"bytes,3,opt,name=lsm,proto3" json:"lsm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FibMplsLlcTypeExt) Reset()         { *m = FibMplsLlcTypeExt{} }
func (m *FibMplsLlcTypeExt) String() string { return proto.CompactTextString(m) }
func (*FibMplsLlcTypeExt) ProtoMessage()    {}
func (*FibMplsLlcTypeExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_mpls_llc_01c81bf59f53383b, []int{4}
}
func (m *FibMplsLlcTypeExt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsLlcTypeExt.Unmarshal(m, b)
}
func (m *FibMplsLlcTypeExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsLlcTypeExt.Marshal(b, m, deterministic)
}
func (dst *FibMplsLlcTypeExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsLlcTypeExt.Merge(dst, src)
}
func (m *FibMplsLlcTypeExt) XXX_Size() int {
	return xxx_messageInfo_FibMplsLlcTypeExt.Size(m)
}
func (m *FibMplsLlcTypeExt) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsLlcTypeExt.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsLlcTypeExt proto.InternalMessageInfo

func (m *FibMplsLlcTypeExt) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FibMplsLlcTypeExt) GetPfx() *FibMplsLlcPfx {
	if m != nil {
		return m.Pfx
	}
	return nil
}

func (m *FibMplsLlcTypeExt) GetLsm() *FibMplsLlcLsm {
	if m != nil {
		return m.Lsm
	}
	return nil
}

func init() {
	proto.RegisterType((*FibMplsLlc_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.local_label.conflicts.conflict.fib_mpls_llc_KEYS")
	proto.RegisterType((*FibMplsLlc)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.local_label.conflicts.conflict.fib_mpls_llc")
	proto.RegisterType((*FibMplsLlcPfx)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.local_label.conflicts.conflict.fib_mpls_llc_pfx")
	proto.RegisterType((*FibMplsLlcLsm)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.local_label.conflicts.conflict.fib_mpls_llc_lsm")
	proto.RegisterType((*FibMplsLlcTypeExt)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.local_label.conflicts.conflict.fib_mpls_llc_type_ext")
}

func init() { proto.RegisterFile("fib_mpls_llc.proto", fileDescriptor_fib_mpls_llc_01c81bf59f53383b) }

var fileDescriptor_fib_mpls_llc_01c81bf59f53383b = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xbd, 0x6e, 0x14, 0x31,
	0x10, 0xd6, 0x6e, 0x2e, 0x77, 0x7b, 0x93, 0x04, 0x05, 0x8b, 0x20, 0xa3, 0x80, 0x38, 0x2d, 0xcd,
	0x55, 0x5b, 0x24, 0xfc, 0x14, 0x88, 0x92, 0x22, 0x22, 0xa2, 0x58, 0xae, 0xa1, 0x40, 0xd6, 0xae,
	0xd7, 0xab, 0xac, 0xe4, 0x3f, 0xad, 0x7d, 0x92, 0xaf, 0xa4, 0xe1, 0x5d, 0xa8, 0x79, 0x34, 0x5e,
	0x00, 0x79, 0x7c, 0x17, 0x2e, 0x11, 0x2d, 0x69, 0xac, 0xf9, 0xfb, 0xe6, 0xfb, 0x3c, 0x33, 0x40,
	0xfa, 0xa1, 0x65, 0xca, 0x4a, 0xc7, 0xa4, 0xe4, 0x95, 0x1d, 0x8d, 0x37, 0xe4, 0x1b, 0x1f, 0x1c,
	0x37, 0x6c, 0x30, 0x8e, 0x85, 0x91, 0xc5, 0x02, 0x6e, 0x94, 0x32, 0x9a, 0x19, 0x2b, 0xc6, 0xaa,
	0x1f, 0xda, 0x4a, 0x9b, 0x4e, 0x38, 0x7c, 0x13, 0x84, 0x1b, 0xe9, 0x6e, 0xad, 0x4a, 0x1a, 0xde,
	0x48, 0x26, 0x9b, 0x56, 0xc8, 0x8a, 0x1b, 0xdd, 0xcb, 0x81, 0x7b, 0x77, 0x6b, 0x95, 0xbf, 0x33,
	0x78, 0xbc, 0xcf, 0xca, 0x3e, 0x7d, 0xfc, 0xfa, 0x85, 0x9c, 0xc3, 0x3c, 0xf6, 0x63, 0xba, 0x51,
	0x82, 0x66, 0x8b, 0x6c, 0x39, 0xaf, 0x8b, 0x18, 0xf8, 0xdc, 0x28, 0x41, 0x5e, 0xc1, 0xc9, 0xae,
	0x7b, 0x2a, 0xc8, 0xb1, 0xe0, 0x78, 0x17, 0xc4, 0xa2, 0x27, 0x70, 0x88, 0x9c, 0xf4, 0x60, 0x91,
	0x2d, 0x4f, 0xea, 0xe4, 0x90, 0xa7, 0x30, 0x75, 0x66, 0x3d, 0x72, 0x41, 0x27, 0x88, 0xd9, 0x7a,
	0xe4, 0x19, 0x14, 0x91, 0xdb, 0x6f, 0xac, 0xa0, 0x87, 0x98, 0x99, 0x49, 0xc9, 0x57, 0x1b, 0x2b,
	0xc8, 0x73, 0x00, 0xdb, 0x07, 0xe6, 0x5b, 0xc9, 0x86, 0x8e, 0x4e, 0xb1, 0x5b, 0x61, 0xfb, 0xb0,
	0x6a, 0xe5, 0x55, 0x17, 0x1b, 0xda, 0x51, 0xf4, 0x43, 0xa0, 0xb3, 0xd4, 0x30, 0x79, 0xe4, 0x05,
	0x40, 0xb2, 0x98, 0x14, 0x9a, 0x16, 0x88, 0x9a, 0xa7, 0xc8, 0xb5, 0xd0, 0xe5, 0xcf, 0x1c, 0x8e,
	0xf7, 0x7f, 0x4d, 0x5e, 0xc2, 0xd1, 0xde, 0xa0, 0xe8, 0x05, 0x02, 0x00, 0x43, 0xd7, 0xf7, 0x94,
	0x5f, 0x62, 0x6e, 0xa7, 0xfc, 0x1c, 0xe6, 0x6b, 0xdb, 0x35, 0x5e, 0x30, 0xef, 0xe8, 0xeb, 0x45,
	0xb6, 0x9c, 0xd4, 0x45, 0x0a, 0xac, 0x5c, 0xfc, 0xd6, 0x28, 0xfc, 0xb8, 0x89, 0xb9, 0x37, 0x98,
	0x9b, 0xa1, 0xbf, 0x72, 0x91, 0x50, 0xaf, 0x15, 0x8b, 0xee, 0x20, 0x1c, 0x7d, 0x9b, 0x08, 0xf5,
	0x5a, 0xd5, 0x29, 0x42, 0x7e, 0x64, 0x70, 0x20, 0x82, 0xa7, 0xef, 0x16, 0xd9, 0xf2, 0xe8, 0xc2,
	0x57, 0xff, 0xf5, 0x0c, 0xaa, 0x3b, 0x27, 0x10, 0xd7, 0xc0, 0x44, 0xf0, 0x75, 0x14, 0x50, 0xbe,
	0x87, 0xd3, 0x3b, 0x59, 0xdb, 0x07, 0x72, 0x0a, 0x07, 0xb6, 0x0f, 0xdb, 0xcb, 0x88, 0x26, 0x39,
	0x83, 0xe9, 0x76, 0x45, 0x79, 0x5a, 0xb8, 0x8f, 0xfb, 0x29, 0x3f, 0xdc, 0x03, 0x4b, 0xa7, 0xc8,
	0x23, 0xc8, 0xf5, 0xcd, 0x16, 0x9b, 0xeb, 0x9b, 0x38, 0x25, 0xc5, 0x1b, 0xe7, 0xff, 0x82, 0x67,
	0xe8, 0x5f, 0x75, 0xe5, 0xaf, 0x1c, 0xce, 0xfe, 0x29, 0x8d, 0x10, 0x98, 0xe0, 0xb5, 0xa4, 0x36,
	0x68, 0x93, 0xef, 0x59, 0x92, 0x95, 0xe3, 0xc8, 0xcc, 0x43, 0x8e, 0xcc, 0xf6, 0x21, 0xcd, 0x21,
	0x6a, 0x90, 0x4e, 0xe1, 0xd9, 0x3f, 0xb0, 0x06, 0xe9, 0x54, 0x1d, 0xb9, 0xdb, 0x29, 0xc2, 0x2f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe8, 0xbc, 0xc2, 0x4f, 0x04, 0x00, 0x00,
}
