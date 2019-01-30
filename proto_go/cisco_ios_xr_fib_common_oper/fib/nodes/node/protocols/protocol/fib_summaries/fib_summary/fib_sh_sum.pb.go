// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fib_sh_sum.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_fib_summaries_fib_summary

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

// FIB summary statistics
type FibShSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	TableId              uint32   `protobuf:"varint,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShSum_KEYS) Reset()         { *m = FibShSum_KEYS{} }
func (m *FibShSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShSum_KEYS) ProtoMessage()    {}
func (*FibShSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_sum_6a53d1029f573e42, []int{0}
}
func (m *FibShSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShSum_KEYS.Unmarshal(m, b)
}
func (m *FibShSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShSum_KEYS.Marshal(b, m, deterministic)
}
func (dst *FibShSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShSum_KEYS.Merge(dst, src)
}
func (m *FibShSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShSum_KEYS.Size(m)
}
func (m *FibShSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShSum_KEYS proto.InternalMessageInfo

func (m *FibShSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibShSum_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibShSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibShSum_KEYS) GetTableId() uint32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

type FibShSum struct {
	// The router-id
	Prefix []byte `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Table Id
	SsTblId uint32 `protobuf:"varint,51,opt,name=ss_tbl_id,json=ssTblId,proto3" json:"ss_tbl_id,omitempty"`
	// Table Id Ptr
	SsTblIdPtr uint32 `protobuf:"varint,52,opt,name=ss_tbl_id_ptr,json=ssTblIdPtr,proto3" json:"ss_tbl_id_ptr,omitempty"`
	// Virtual routing forwarding instance Id
	SsVrfId uint32 `protobuf:"varint,53,opt,name=ss_vrf_id,json=ssVrfId,proto3" json:"ss_vrf_id,omitempty"`
	// Virtual router instance Id
	SsVrId uint32 `protobuf:"varint,54,opt,name=ss_vr_id,json=ssVrId,proto3" json:"ss_vr_id,omitempty"`
	// LBA configuration state
	LoadBalancing string `protobuf:"bytes,55,opt,name=load_balancing,json=loadBalancing,proto3" json:"load_balancing,omitempty"`
	// Number of forwarding elements linked to the table
	ForwardingElements uint32 `protobuf:"varint,56,opt,name=forwarding_elements,json=forwardingElements,proto3" json:"forwarding_elements,omitempty"`
	// Number of routes
	Routes uint32 `protobuf:"varint,57,opt,name=routes,proto3" json:"routes,omitempty"`
	// Number of inplace modifications
	PrefixInPlaceModifications uint32 `protobuf:"varint,58,opt,name=prefix_in_place_modifications,json=prefixInPlaceModifications,proto3" json:"prefix_in_place_modifications,omitempty"`
	// Number of deleted stale leafs
	StalePrefixDeletes uint32 `protobuf:"varint,59,opt,name=stale_prefix_deletes,json=stalePrefixDeletes,proto3" json:"stale_prefix_deletes,omitempty"`
	// Count of load sharing elements
	LoadSharingElements uint32 `protobuf:"varint,60,opt,name=load_sharing_elements,json=loadSharingElements,proto3" json:"load_sharing_elements,omitempty"`
	// Count of load sharing references
	LoadSharingReferences uint64 `protobuf:"varint,61,opt,name=load_sharing_references,json=loadSharingReferences,proto3" json:"load_sharing_references,omitempty"`
	// Total memory used by load sharing elements
	TotalLoadShareElementBytes uint32 `protobuf:"varint,62,opt,name=total_load_share_element_bytes,json=totalLoadShareElementBytes,proto3" json:"total_load_share_element_bytes,omitempty"`
	// Exclusive load sharing element
	ExclusiveLoadSharingElement *FibPlLdiCount `protobuf:"bytes,63,opt,name=exclusive_load_sharing_element,json=exclusiveLoadSharingElement,proto3" json:"exclusive_load_sharing_element,omitempty"`
	// Shared load sharing element
	SharedLoadSharingElement *FibPlLdiCount `protobuf:"bytes,64,opt,name=shared_load_sharing_element,json=sharedLoadSharingElement,proto3" json:"shared_load_sharing_element,omitempty"`
	// Cross-table shared load sharing element
	CrossSharedLoadSharingElement *FibPlLdiCount `protobuf:"bytes,65,opt,name=cross_shared_load_sharing_element,json=crossSharedLoadSharingElement,proto3" json:"cross_shared_load_sharing_element,omitempty"`
	// Label-shared load sharing element
	LabelSharedLoadSharingElement *FibPlLdiCount `protobuf:"bytes,66,opt,name=label_shared_load_sharing_element,json=labelSharedLoadSharingElement,proto3" json:"label_shared_load_sharing_element,omitempty"`
	// Total memory used by leaves
	LeavesUsedBytes uint32 `protobuf:"varint,67,opt,name=leaves_used_bytes,json=leavesUsedBytes,proto3" json:"leaves_used_bytes,omitempty"`
	// Number of reresolved entries
	ReresolveEntries uint32 `protobuf:"varint,68,opt,name=reresolve_entries,json=reresolveEntries,proto3" json:"reresolve_entries,omitempty"`
	// Number of old unresolved entries
	OldUnresolveEntries uint32 `protobuf:"varint,69,opt,name=old_unresolve_entries,json=oldUnresolveEntries,proto3" json:"old_unresolve_entries,omitempty"`
	// Number of new unresolved entries
	NewUnresolveEntries uint32 `protobuf:"varint,70,opt,name=new_unresolve_entries,json=newUnresolveEntries,proto3" json:"new_unresolve_entries,omitempty"`
	// Number of total unresolved entries
	UnresolveEntries uint32 `protobuf:"varint,71,opt,name=unresolve_entries,json=unresolveEntries,proto3" json:"unresolve_entries,omitempty"`
	// Number of routes dropped by CEF
	CefRouteDrops uint32 `protobuf:"varint,72,opt,name=cef_route_drops,json=cefRouteDrops,proto3" json:"cef_route_drops,omitempty"`
	// the number of routes dropped due to version mismatch
	CefVersionMismatchRouteDrops uint64 `protobuf:"varint,73,opt,name=cef_version_mismatch_route_drops,json=cefVersionMismatchRouteDrops,proto3" json:"cef_version_mismatch_route_drops,omitempty"`
	// Number of entries in the route delete cache
	DeleteCacheNumEntries uint32 `protobuf:"varint,74,opt,name=delete_cache_num_entries,json=deleteCacheNumEntries,proto3" json:"delete_cache_num_entries,omitempty"`
	// Number of entries present on addition
	ExistingLeavesRevisions uint32 `protobuf:"varint,75,opt,name=existing_leaves_revisions,json=existingLeavesRevisions,proto3" json:"existing_leaves_revisions,omitempty"`
	// Default prefix
	FibDefaultPrefix uint32 `protobuf:"varint,76,opt,name=fib_default_prefix,json=fibDefaultPrefix,proto3" json:"fib_default_prefix,omitempty"`
	// Default prefix mask length
	FibDefaultPrefixMaskLength uint32 `protobuf:"varint,77,opt,name=fib_default_prefix_mask_length,json=fibDefaultPrefixMaskLength,proto3" json:"fib_default_prefix_mask_length,omitempty"`
	// Number of NHINFOS
	NextHops uint32 `protobuf:"varint,78,opt,name=next_hops,json=nextHops,proto3" json:"next_hops,omitempty"`
	// Number of incomplete NHINFOS
	IncompleteNextHops uint32 `protobuf:"varint,79,opt,name=incomplete_next_hops,json=incompleteNextHops,proto3" json:"incomplete_next_hops,omitempty"`
	// IP CEF resolution timer in seconds
	ResolutionTimer uint32 `protobuf:"varint,80,opt,name=resolution_timer,json=resolutionTimer,proto3" json:"resolution_timer,omitempty"`
	// IP CEF slow processing time in seconds
	SlowProcessTimer uint32 `protobuf:"varint,81,opt,name=slow_process_timer,json=slowProcessTimer,proto3" json:"slow_process_timer,omitempty"`
	// IP CEF max resolution time in seconds
	MaxResolutionTimer uint32 `protobuf:"varint,82,opt,name=max_resolution_timer,json=maxResolutionTimer,proto3" json:"max_resolution_timer,omitempty"`
	// Number of prefixes with imposition LDI
	ImpositionPrefixes uint32 `protobuf:"varint,83,opt,name=imposition_prefixes,json=impositionPrefixes,proto3" json:"imposition_prefixes,omitempty"`
	// Number of prefixes with extended path-list
	ExtendedPrefixes uint32 `protobuf:"varint,84,opt,name=extended_prefixes,json=extendedPrefixes,proto3" json:"extended_prefixes,omitempty"`
	// Number of routes updates with recycled label handled
	CeflBlRecycledRoutes uint32 `protobuf:"varint,85,opt,name=cefl_bl_recycled_routes,json=ceflBlRecycledRoutes,proto3" json:"cefl_bl_recycled_routes,omitempty"`
	// pd backwalks on LDI modify with backup path
	LdiBackwalks uint32 `protobuf:"varint,86,opt,name=ldi_backwalks,json=ldiBackwalks,proto3" json:"ldi_backwalks,omitempty"`
	// Number of routes with FRR protection
	SsProtRouteCount uint32 `protobuf:"varint,87,opt,name=ss_prot_route_count,json=ssProtRouteCount,proto3" json:"ss_prot_route_count,omitempty"`
	// Number of lisp eid prefixes associated with table
	LispEidPrefixes uint32 `protobuf:"varint,88,opt,name=lisp_eid_prefixes,json=lispEidPrefixes,proto3" json:"lisp_eid_prefixes,omitempty"`
	// Number of lisp eid prefixes eligible for forwarding
	LispEidValidPrefixes uint32 `protobuf:"varint,89,opt,name=lisp_eid_valid_prefixes,json=lispEidValidPrefixes,proto3" json:"lisp_eid_valid_prefixes,omitempty"`
	// Number of lisp rloc objects associated with table
	LispRlocObjects uint32 `protobuf:"varint,90,opt,name=lisp_rloc_objects,json=lispRlocObjects,proto3" json:"lisp_rloc_objects,omitempty"`
	// VXLAN local Interface handle
	SsVxlanLtepIfh string `protobuf:"bytes,91,opt,name=ss_vxlan_ltep_ifh,json=ssVxlanLtepIfh,proto3" json:"ss_vxlan_ltep_ifh,omitempty"`
	// Number of dropped pathlists
	SsDropPlCount uint32 `protobuf:"varint,92,opt,name=ss_drop_pl_count,json=ssDropPlCount,proto3" json:"ss_drop_pl_count,omitempty"`
	// Distribution of prefix mask lengths
	PrefixMasklenDistribution *FibShPfxMasklenDistrib `protobuf:"bytes,93,opt,name=prefix_masklen_distribution,json=prefixMasklenDistribution,proto3" json:"prefix_masklen_distribution,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                `json:"-"`
	XXX_unrecognized          []byte                  `json:"-"`
	XXX_sizecache             int32                   `json:"-"`
}

func (m *FibShSum) Reset()         { *m = FibShSum{} }
func (m *FibShSum) String() string { return proto.CompactTextString(m) }
func (*FibShSum) ProtoMessage()    {}
func (*FibShSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_sum_6a53d1029f573e42, []int{1}
}
func (m *FibShSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShSum.Unmarshal(m, b)
}
func (m *FibShSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShSum.Marshal(b, m, deterministic)
}
func (dst *FibShSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShSum.Merge(dst, src)
}
func (m *FibShSum) XXX_Size() int {
	return xxx_messageInfo_FibShSum.Size(m)
}
func (m *FibShSum) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShSum.DiscardUnknown(m)
}

var xxx_messageInfo_FibShSum proto.InternalMessageInfo

func (m *FibShSum) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *FibShSum) GetSsTblId() uint32 {
	if m != nil {
		return m.SsTblId
	}
	return 0
}

func (m *FibShSum) GetSsTblIdPtr() uint32 {
	if m != nil {
		return m.SsTblIdPtr
	}
	return 0
}

func (m *FibShSum) GetSsVrfId() uint32 {
	if m != nil {
		return m.SsVrfId
	}
	return 0
}

func (m *FibShSum) GetSsVrId() uint32 {
	if m != nil {
		return m.SsVrId
	}
	return 0
}

func (m *FibShSum) GetLoadBalancing() string {
	if m != nil {
		return m.LoadBalancing
	}
	return ""
}

func (m *FibShSum) GetForwardingElements() uint32 {
	if m != nil {
		return m.ForwardingElements
	}
	return 0
}

func (m *FibShSum) GetRoutes() uint32 {
	if m != nil {
		return m.Routes
	}
	return 0
}

func (m *FibShSum) GetPrefixInPlaceModifications() uint32 {
	if m != nil {
		return m.PrefixInPlaceModifications
	}
	return 0
}

func (m *FibShSum) GetStalePrefixDeletes() uint32 {
	if m != nil {
		return m.StalePrefixDeletes
	}
	return 0
}

func (m *FibShSum) GetLoadSharingElements() uint32 {
	if m != nil {
		return m.LoadSharingElements
	}
	return 0
}

func (m *FibShSum) GetLoadSharingReferences() uint64 {
	if m != nil {
		return m.LoadSharingReferences
	}
	return 0
}

func (m *FibShSum) GetTotalLoadShareElementBytes() uint32 {
	if m != nil {
		return m.TotalLoadShareElementBytes
	}
	return 0
}

func (m *FibShSum) GetExclusiveLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.ExclusiveLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.SharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetCrossSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.CrossSharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetLabelSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.LabelSharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetLeavesUsedBytes() uint32 {
	if m != nil {
		return m.LeavesUsedBytes
	}
	return 0
}

func (m *FibShSum) GetReresolveEntries() uint32 {
	if m != nil {
		return m.ReresolveEntries
	}
	return 0
}

func (m *FibShSum) GetOldUnresolveEntries() uint32 {
	if m != nil {
		return m.OldUnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetNewUnresolveEntries() uint32 {
	if m != nil {
		return m.NewUnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetUnresolveEntries() uint32 {
	if m != nil {
		return m.UnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetCefRouteDrops() uint32 {
	if m != nil {
		return m.CefRouteDrops
	}
	return 0
}

func (m *FibShSum) GetCefVersionMismatchRouteDrops() uint64 {
	if m != nil {
		return m.CefVersionMismatchRouteDrops
	}
	return 0
}

func (m *FibShSum) GetDeleteCacheNumEntries() uint32 {
	if m != nil {
		return m.DeleteCacheNumEntries
	}
	return 0
}

func (m *FibShSum) GetExistingLeavesRevisions() uint32 {
	if m != nil {
		return m.ExistingLeavesRevisions
	}
	return 0
}

func (m *FibShSum) GetFibDefaultPrefix() uint32 {
	if m != nil {
		return m.FibDefaultPrefix
	}
	return 0
}

func (m *FibShSum) GetFibDefaultPrefixMaskLength() uint32 {
	if m != nil {
		return m.FibDefaultPrefixMaskLength
	}
	return 0
}

func (m *FibShSum) GetNextHops() uint32 {
	if m != nil {
		return m.NextHops
	}
	return 0
}

func (m *FibShSum) GetIncompleteNextHops() uint32 {
	if m != nil {
		return m.IncompleteNextHops
	}
	return 0
}

func (m *FibShSum) GetResolutionTimer() uint32 {
	if m != nil {
		return m.ResolutionTimer
	}
	return 0
}

func (m *FibShSum) GetSlowProcessTimer() uint32 {
	if m != nil {
		return m.SlowProcessTimer
	}
	return 0
}

func (m *FibShSum) GetMaxResolutionTimer() uint32 {
	if m != nil {
		return m.MaxResolutionTimer
	}
	return 0
}

func (m *FibShSum) GetImpositionPrefixes() uint32 {
	if m != nil {
		return m.ImpositionPrefixes
	}
	return 0
}

func (m *FibShSum) GetExtendedPrefixes() uint32 {
	if m != nil {
		return m.ExtendedPrefixes
	}
	return 0
}

func (m *FibShSum) GetCeflBlRecycledRoutes() uint32 {
	if m != nil {
		return m.CeflBlRecycledRoutes
	}
	return 0
}

func (m *FibShSum) GetLdiBackwalks() uint32 {
	if m != nil {
		return m.LdiBackwalks
	}
	return 0
}

func (m *FibShSum) GetSsProtRouteCount() uint32 {
	if m != nil {
		return m.SsProtRouteCount
	}
	return 0
}

func (m *FibShSum) GetLispEidPrefixes() uint32 {
	if m != nil {
		return m.LispEidPrefixes
	}
	return 0
}

func (m *FibShSum) GetLispEidValidPrefixes() uint32 {
	if m != nil {
		return m.LispEidValidPrefixes
	}
	return 0
}

func (m *FibShSum) GetLispRlocObjects() uint32 {
	if m != nil {
		return m.LispRlocObjects
	}
	return 0
}

func (m *FibShSum) GetSsVxlanLtepIfh() string {
	if m != nil {
		return m.SsVxlanLtepIfh
	}
	return ""
}

func (m *FibShSum) GetSsDropPlCount() uint32 {
	if m != nil {
		return m.SsDropPlCount
	}
	return 0
}

func (m *FibShSum) GetPrefixMasklenDistribution() *FibShPfxMasklenDistrib {
	if m != nil {
		return m.PrefixMasklenDistribution
	}
	return nil
}

// FIB Pathlist and Loadinfo summary
type FibPlLdiCount struct {
	// Total memory used by load sharing elements in bytes
	TotalLoadSharingElementBytes uint32 `protobuf:"varint,1,opt,name=total_load_sharing_element_bytes,json=totalLoadSharingElementBytes,proto3" json:"total_load_sharing_element_bytes,omitempty"`
	// Total count of references to load sharing elements
	TotalLoadSharingElementReferences uint64 `protobuf:"varint,2,opt,name=total_load_sharing_element_references,json=totalLoadSharingElementReferences,proto3" json:"total_load_sharing_element_references,omitempty"`
	// Total count of Pathlist elements
	TotalPathListElements uint32 `protobuf:"varint,3,opt,name=total_path_list_elements,json=totalPathListElements,proto3" json:"total_path_list_elements,omitempty"`
	// Count of recursive Pathlist elements
	RecursivePathListElements uint32 `protobuf:"varint,4,opt,name=recursive_path_list_elements,json=recursivePathListElements,proto3" json:"recursive_path_list_elements,omitempty"`
	// Count of platform shared Pathlist elements
	PlatformSharedPathListElements uint32 `protobuf:"varint,5,opt,name=platform_shared_path_list_elements,json=platformSharedPathListElements,proto3" json:"platform_shared_path_list_elements,omitempty"`
	// Count of Pathlist elements in retry
	RetryPathListElements uint32 `protobuf:"varint,6,opt,name=retry_path_list_elements,json=retryPathListElements,proto3" json:"retry_path_list_elements,omitempty"`
	// Total count of Loadinfo elements
	TotalLoadInfoElements uint32 `protobuf:"varint,7,opt,name=total_load_info_elements,json=totalLoadInfoElements,proto3" json:"total_load_info_elements,omitempty"`
	// Count of recursive Loadinfo elements
	RecursiveLoadInfoElements uint32 `protobuf:"varint,8,opt,name=recursive_load_info_elements,json=recursiveLoadInfoElements,proto3" json:"recursive_load_info_elements,omitempty"`
	// Count of platform shared Loadinfo elements
	PlatformSharedLoadInfoElements uint32   `protobuf:"varint,9,opt,name=platform_shared_load_info_elements,json=platformSharedLoadInfoElements,proto3" json:"platform_shared_load_info_elements,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *FibPlLdiCount) Reset()         { *m = FibPlLdiCount{} }
func (m *FibPlLdiCount) String() string { return proto.CompactTextString(m) }
func (*FibPlLdiCount) ProtoMessage()    {}
func (*FibPlLdiCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_sum_6a53d1029f573e42, []int{2}
}
func (m *FibPlLdiCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibPlLdiCount.Unmarshal(m, b)
}
func (m *FibPlLdiCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibPlLdiCount.Marshal(b, m, deterministic)
}
func (dst *FibPlLdiCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibPlLdiCount.Merge(dst, src)
}
func (m *FibPlLdiCount) XXX_Size() int {
	return xxx_messageInfo_FibPlLdiCount.Size(m)
}
func (m *FibPlLdiCount) XXX_DiscardUnknown() {
	xxx_messageInfo_FibPlLdiCount.DiscardUnknown(m)
}

var xxx_messageInfo_FibPlLdiCount proto.InternalMessageInfo

func (m *FibPlLdiCount) GetTotalLoadSharingElementBytes() uint32 {
	if m != nil {
		return m.TotalLoadSharingElementBytes
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalLoadSharingElementReferences() uint64 {
	if m != nil {
		return m.TotalLoadSharingElementReferences
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalPathListElements() uint32 {
	if m != nil {
		return m.TotalPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRecursivePathListElements() uint32 {
	if m != nil {
		return m.RecursivePathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetPlatformSharedPathListElements() uint32 {
	if m != nil {
		return m.PlatformSharedPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRetryPathListElements() uint32 {
	if m != nil {
		return m.RetryPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalLoadInfoElements() uint32 {
	if m != nil {
		return m.TotalLoadInfoElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRecursiveLoadInfoElements() uint32 {
	if m != nil {
		return m.RecursiveLoadInfoElements
	}
	return 0
}

func (m *FibPlLdiCount) GetPlatformSharedLoadInfoElements() uint32 {
	if m != nil {
		return m.PlatformSharedLoadInfoElements
	}
	return 0
}

// FIB Prefix Masklength counts
type FibShPfxMasklenCnt struct {
	// Mask length
	MaskLength uint32 `protobuf:"varint,1,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	// Number of prefixes with given mask length
	NumberOfPrefixes     uint32   `protobuf:"varint,2,opt,name=number_of_prefixes,json=numberOfPrefixes,proto3" json:"number_of_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShPfxMasklenCnt) Reset()         { *m = FibShPfxMasklenCnt{} }
func (m *FibShPfxMasklenCnt) String() string { return proto.CompactTextString(m) }
func (*FibShPfxMasklenCnt) ProtoMessage()    {}
func (*FibShPfxMasklenCnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_sum_6a53d1029f573e42, []int{3}
}
func (m *FibShPfxMasklenCnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShPfxMasklenCnt.Unmarshal(m, b)
}
func (m *FibShPfxMasklenCnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShPfxMasklenCnt.Marshal(b, m, deterministic)
}
func (dst *FibShPfxMasklenCnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShPfxMasklenCnt.Merge(dst, src)
}
func (m *FibShPfxMasklenCnt) XXX_Size() int {
	return xxx_messageInfo_FibShPfxMasklenCnt.Size(m)
}
func (m *FibShPfxMasklenCnt) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShPfxMasklenCnt.DiscardUnknown(m)
}

var xxx_messageInfo_FibShPfxMasklenCnt proto.InternalMessageInfo

func (m *FibShPfxMasklenCnt) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

func (m *FibShPfxMasklenCnt) GetNumberOfPrefixes() uint32 {
	if m != nil {
		return m.NumberOfPrefixes
	}
	return 0
}

// FIB Prefix Masklengths distribution
type FibShPfxMasklenDistrib struct {
	// Masklength counts for unicast prefixes
	UnicastPrefixes []*FibShPfxMasklenCnt `protobuf:"bytes,1,rep,name=unicast_prefixes,json=unicastPrefixes,proto3" json:"unicast_prefixes,omitempty"`
	// Masklength counts for broadcast prefixes
	BroadcastPrefixes []*FibShPfxMasklenCnt `protobuf:"bytes,2,rep,name=broadcast_prefixes,json=broadcastPrefixes,proto3" json:"broadcast_prefixes,omitempty"`
	// Masklength counts for multicast prefixes
	MulticastPrefix      []*FibShPfxMasklenCnt `protobuf:"bytes,3,rep,name=multicast_prefix,json=multicastPrefix,proto3" json:"multicast_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FibShPfxMasklenDistrib) Reset()         { *m = FibShPfxMasklenDistrib{} }
func (m *FibShPfxMasklenDistrib) String() string { return proto.CompactTextString(m) }
func (*FibShPfxMasklenDistrib) ProtoMessage()    {}
func (*FibShPfxMasklenDistrib) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_sum_6a53d1029f573e42, []int{4}
}
func (m *FibShPfxMasklenDistrib) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Unmarshal(m, b)
}
func (m *FibShPfxMasklenDistrib) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Marshal(b, m, deterministic)
}
func (dst *FibShPfxMasklenDistrib) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShPfxMasklenDistrib.Merge(dst, src)
}
func (m *FibShPfxMasklenDistrib) XXX_Size() int {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Size(m)
}
func (m *FibShPfxMasklenDistrib) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShPfxMasklenDistrib.DiscardUnknown(m)
}

var xxx_messageInfo_FibShPfxMasklenDistrib proto.InternalMessageInfo

func (m *FibShPfxMasklenDistrib) GetUnicastPrefixes() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.UnicastPrefixes
	}
	return nil
}

func (m *FibShPfxMasklenDistrib) GetBroadcastPrefixes() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.BroadcastPrefixes
	}
	return nil
}

func (m *FibShPfxMasklenDistrib) GetMulticastPrefix() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.MulticastPrefix
	}
	return nil
}

func init() {
	proto.RegisterType((*FibShSum_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_sum_KEYS")
	proto.RegisterType((*FibShSum)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_sum")
	proto.RegisterType((*FibPlLdiCount)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_pl_ldi_count")
	proto.RegisterType((*FibShPfxMasklenCnt)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_pfx_masklen_cnt")
	proto.RegisterType((*FibShPfxMasklenDistrib)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_pfx_masklen_distrib")
}

func init() { proto.RegisterFile("fib_sh_sum.proto", fileDescriptor_fib_sh_sum_6a53d1029f573e42) }

var fileDescriptor_fib_sh_sum_6a53d1029f573e42 = []byte{
	// 1456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdf, 0x57, 0x14, 0x39,
	0x16, 0x3e, 0x0d, 0x2e, 0x42, 0x14, 0x81, 0x02, 0xb4, 0x10, 0x75, 0x11, 0x8f, 0xbb, 0xb8, 0xba,
	0xe8, 0xc1, 0x55, 0x76, 0xdd, 0xdd, 0x71, 0x44, 0x70, 0x6c, 0x05, 0xec, 0x29, 0x94, 0xd1, 0xf9,
	0x71, 0x72, 0x52, 0x55, 0xb7, 0xe8, 0x48, 0x2a, 0xa9, 0x93, 0xa4, 0x9a, 0xe6, 0x4f, 0x98, 0xbf,
	0x61, 0xe6, 0x9c, 0x79, 0x9e, 0xa7, 0x79, 0xf1, 0x1f, 0x9b, 0xd7, 0x79, 0x9a, 0x93, 0xa4, 0xaa,
	0xba, 0xa0, 0x5b, 0xde, 0xc4, 0x17, 0x0e, 0x75, 0xbf, 0xef, 0xde, 0x7c, 0xf7, 0x76, 0xf2, 0x25,
	0x68, 0x32, 0xa1, 0x21, 0x56, 0x6d, 0xac, 0xf2, 0x74, 0x39, 0x93, 0x42, 0x0b, 0xef, 0x6d, 0x44,
	0x55, 0x24, 0x30, 0x15, 0x0a, 0x77, 0x25, 0x36, 0x70, 0x24, 0xd2, 0x54, 0x70, 0x2c, 0x32, 0x90,
	0xcb, 0x09, 0x0d, 0x97, 0xb9, 0x88, 0x41, 0xd9, 0xbf, 0x2e, 0x25, 0x12, 0x4c, 0x55, 0xff, 0x2d,
	0xdb, 0x82, 0x79, 0x9a, 0x12, 0x49, 0x41, 0xd5, 0xbe, 0x0e, 0x17, 0x7f, 0x6c, 0xa0, 0x89, 0xde,
	0x72, 0xf8, 0xe5, 0xc6, 0xbb, 0x1d, 0x6f, 0x1e, 0x8d, 0x99, 0x42, 0x98, 0x93, 0x14, 0xfc, 0xc6,
	0x42, 0x63, 0x69, 0x2c, 0x18, 0x35, 0x81, 0x6d, 0x92, 0x82, 0x77, 0x03, 0x8d, 0x97, 0x65, 0x1d,
	0x61, 0xc8, 0x12, 0xce, 0x97, 0x41, 0x4b, 0x9a, 0x43, 0xa3, 0x1d, 0x99, 0x38, 0x7c, 0xd8, 0xe2,
	0x67, 0x3b, 0x32, 0x29, 0x21, 0x4d, 0x42, 0x06, 0x98, 0xc6, 0xfe, 0x99, 0x85, 0xc6, 0xd2, 0x78,
	0x70, 0xd6, 0x7e, 0x37, 0xe3, 0xc5, 0x9f, 0x67, 0x11, 0xea, 0x69, 0xf1, 0x2e, 0xa2, 0x91, 0x4c,
	0x42, 0x42, 0xbb, 0xfe, 0xca, 0x42, 0x63, 0xe9, 0x7c, 0x50, 0x7c, 0x79, 0x97, 0xd1, 0x98, 0x52,
	0x58, 0x87, 0xcc, 0x94, 0xb8, 0xef, 0x4a, 0x28, 0xf5, 0x3a, 0x64, 0xcd, 0xd8, 0xbb, 0x8e, 0xc6,
	0x2b, 0x0c, 0x67, 0x5a, 0xfa, 0xff, 0xb2, 0x38, 0x2a, 0xf0, 0x96, 0x96, 0x45, 0xba, 0x91, 0x47,
	0x63, 0xff, 0x41, 0x99, 0xbe, 0x2b, 0x93, 0x66, 0xec, 0xf9, 0x68, 0xd4, 0x62, 0x06, 0x7a, 0x68,
	0xa1, 0x11, 0x03, 0x35, 0x63, 0xef, 0x26, 0xba, 0xc0, 0x04, 0x89, 0x71, 0x48, 0x18, 0xe1, 0x11,
	0xe5, 0x7b, 0xfe, 0xaa, 0xed, 0x6b, 0xdc, 0x44, 0xd7, 0xca, 0xa0, 0x77, 0x17, 0x4d, 0x27, 0x42,
	0x1e, 0x10, 0x19, 0x53, 0xbe, 0x87, 0x81, 0x41, 0x0a, 0x5c, 0x2b, 0xff, 0xdf, 0xb6, 0x96, 0xd7,
	0x83, 0x36, 0x0a, 0xc4, 0x34, 0x29, 0x45, 0xae, 0x41, 0xf9, 0xff, 0x71, 0xeb, 0xb9, 0x2f, 0xef,
	0x09, 0xba, 0xea, 0xda, 0xc5, 0x94, 0xe3, 0x8c, 0x91, 0x08, 0x70, 0x2a, 0x62, 0x9a, 0xd0, 0x88,
	0x68, 0x2a, 0xb8, 0xf2, 0x1f, 0x59, 0xfa, 0x65, 0x47, 0x6a, 0xf2, 0x96, 0xa1, 0x6c, 0xd5, 0x19,
	0xde, 0x3d, 0x34, 0xa3, 0x34, 0x61, 0x80, 0x8b, 0x42, 0x31, 0x30, 0x30, 0x0b, 0xfd, 0xd7, 0x89,
	0xb1, 0x58, 0xcb, 0x42, 0xeb, 0x0e, 0xf1, 0x56, 0xd0, 0xac, 0x6d, 0x52, 0xb5, 0x89, 0x3c, 0xa2,
	0xff, 0x7f, 0x36, 0x65, 0xda, 0x80, 0x3b, 0x0e, 0xab, 0x1a, 0x78, 0x88, 0x2e, 0x1d, 0xc9, 0x91,
	0x90, 0x80, 0x04, 0x1e, 0x81, 0xf2, 0xff, 0xbf, 0xd0, 0x58, 0x3a, 0x13, 0xcc, 0xd6, 0xb2, 0x82,
	0x0a, 0xf4, 0xd6, 0xd0, 0x35, 0x2d, 0x34, 0x61, 0xb8, 0xca, 0x86, 0x72, 0x3d, 0x1c, 0x1e, 0x1a,
	0x9d, 0x5f, 0xb8, 0x0e, 0x2d, 0x6b, 0xb3, 0xa8, 0x01, 0xc5, 0xba, 0x6b, 0x86, 0xe1, 0xfd, 0xd6,
	0x40, 0xd7, 0xa0, 0x1b, 0xb1, 0x5c, 0xd1, 0x0e, 0xe0, 0x41, 0xd2, 0xfd, 0xc7, 0x0b, 0x8d, 0xa5,
	0x73, 0x2b, 0xef, 0x97, 0x3f, 0xd5, 0x01, 0xb2, 0xff, 0x67, 0x0c, 0xb3, 0x98, 0xe2, 0x48, 0xe4,
	0x5c, 0x07, 0xf3, 0x95, 0xa2, 0xcd, 0xbe, 0x71, 0x79, 0xbf, 0x36, 0xd0, 0xbc, 0xed, 0x35, 0x1e,
	0x2c, 0xf7, 0xcb, 0x53, 0x97, 0xeb, 0x3b, 0x39, 0x03, 0xb4, 0x7e, 0x68, 0xa0, 0xeb, 0x91, 0x14,
	0x4a, 0xe1, 0x93, 0x14, 0x3f, 0x39, 0x75, 0xc5, 0x57, 0xad, 0xa8, 0x9d, 0x93, 0x64, 0x33, 0x12,
	0x02, 0x3b, 0x51, 0xf6, 0xda, 0xe9, 0xcb, 0xb6, 0xa2, 0x3e, 0x2a, 0xfb, 0x1f, 0x68, 0x8a, 0x01,
	0xe9, 0x80, 0xc2, 0xb9, 0x82, 0xb8, 0x38, 0x02, 0x4f, 0xed, 0x11, 0x98, 0x70, 0xc0, 0x1b, 0x05,
	0xb1, 0xdb, 0xf7, 0xb7, 0xd1, 0x94, 0x04, 0x09, 0x4a, 0xb0, 0x0e, 0x60, 0xe0, 0xda, 0x2c, 0xee,
	0xaf, 0x5b, 0xee, 0x64, 0x05, 0x6c, 0xb8, 0xb8, 0x39, 0xd4, 0x82, 0xc5, 0x38, 0xe7, 0xc7, 0x13,
	0x36, 0xdc, 0xa1, 0x16, 0x2c, 0x7e, 0xc3, 0xfb, 0x73, 0x38, 0x1c, 0x0c, 0xc8, 0x79, 0xe6, 0x72,
	0x38, 0x1c, 0xf4, 0xe5, 0xdc, 0x46, 0x53, 0xfd, 0xfc, 0xaf, 0x9c, 0xa8, 0xfc, 0x38, 0xf9, 0x6f,
	0x68, 0x22, 0x82, 0x04, 0x5b, 0xb3, 0xc3, 0xb1, 0x14, 0x99, 0xf2, 0x9f, 0x5b, 0xea, 0x78, 0x04,
	0x49, 0x60, 0xa2, 0xeb, 0x26, 0xe8, 0x3d, 0x43, 0x0b, 0x86, 0xd7, 0x01, 0xa9, 0xa8, 0xe0, 0x38,
	0xa5, 0x2a, 0x25, 0x3a, 0x6a, 0x1f, 0x49, 0x6c, 0x5a, 0x9b, 0xb9, 0x12, 0x41, 0xb2, 0xeb, 0x68,
	0x5b, 0x05, 0xab, 0x56, 0x67, 0x15, 0xf9, 0xce, 0xfe, 0x70, 0x44, 0xa2, 0x36, 0x60, 0x9e, 0xa7,
	0x95, 0xc6, 0x17, 0x76, 0xe1, 0x59, 0x87, 0x3f, 0x35, 0xf0, 0x76, 0x9e, 0x96, 0x42, 0x1f, 0xa1,
	0x39, 0xe8, 0x52, 0xa5, 0xcd, 0xde, 0x29, 0x7e, 0x1f, 0x09, 0x1d, 0xaa, 0xac, 0x07, 0xbf, 0xb4,
	0x99, 0x97, 0x4a, 0xc2, 0xa6, 0xc5, 0x83, 0x12, 0xf6, 0xee, 0x20, 0xcf, 0xec, 0x82, 0x18, 0x12,
	0x92, 0x33, 0x5d, 0xd8, 0xb0, 0xbf, 0xe9, 0x46, 0x92, 0xd0, 0x70, 0xdd, 0x01, 0xce, 0x83, 0x8d,
	0x21, 0xf6, 0xb3, 0x71, 0x4a, 0xd4, 0x3e, 0x66, 0xc0, 0xf7, 0x74, 0xdb, 0xdf, 0x72, 0x86, 0x78,
	0x3c, 0x73, 0x8b, 0xa8, 0xfd, 0x4d, 0xcb, 0xb0, 0x37, 0x37, 0x74, 0x35, 0x6e, 0x9b, 0xb9, 0x6c,
	0x5b, 0xfa, 0xa8, 0x09, 0x3c, 0x37, 0x33, 0xb8, 0x87, 0x66, 0x28, 0x8f, 0x44, 0x9a, 0xd9, 0x39,
	0xf4, 0x78, 0xaf, 0xdc, 0x7d, 0xd0, 0xc3, 0xb6, 0xcb, 0x8c, 0x5b, 0x68, 0xd2, 0xfe, 0x6e, 0xb9,
	0xb9, 0x50, 0xb0, 0xa6, 0x29, 0x48, 0xbf, 0xe5, 0xb6, 0x64, 0x2f, 0xfe, 0xda, 0x84, 0x4d, 0xaf,
	0x8a, 0x89, 0x03, 0x9c, 0x49, 0x11, 0x81, 0xb9, 0x82, 0x2d, 0xf9, 0x6b, 0xd7, 0xab, 0x41, 0x5a,
	0x0e, 0x70, 0xec, 0x7b, 0x68, 0x26, 0x25, 0x5d, 0xdc, 0x57, 0x3c, 0x70, 0x52, 0x52, 0xd2, 0x0d,
	0x8e, 0xd5, 0xbf, 0x8b, 0xa6, 0x69, 0x9a, 0x09, 0x45, 0x2d, 0xdb, 0x0d, 0x07, 0x94, 0xbf, 0x53,
	0x68, 0xaf, 0xa0, 0x56, 0x81, 0x98, 0xed, 0x08, 0x5d, 0x0d, 0x3c, 0x86, 0xb8, 0x47, 0x7f, 0xed,
	0xf4, 0x94, 0x40, 0x45, 0x7e, 0x80, 0x2e, 0x45, 0x90, 0x30, 0x1c, 0x32, 0x2c, 0x21, 0x3a, 0x8c,
	0x18, 0xc4, 0xb8, 0xb8, 0x96, 0xdf, 0xd8, 0x94, 0x19, 0x03, 0xaf, 0xb1, 0xa0, 0x00, 0x03, 0x77,
	0x49, 0xdf, 0x40, 0xe3, 0xe6, 0x7c, 0x87, 0x24, 0xda, 0x3f, 0x20, 0x6c, 0x5f, 0xf9, 0xbb, 0x96,
	0x7c, 0x9e, 0xc5, 0x74, 0xad, 0x8c, 0x79, 0xff, 0x44, 0xd3, 0x4a, 0x99, 0xb9, 0xe8, 0x62, 0xd7,
	0x5a, 0x3b, 0xf0, 0xbf, 0x29, 0x46, 0xa3, 0x5a, 0x52, 0x68, 0x5b, 0xef, 0xa9, 0x89, 0x5b, 0x1f,
	0xa0, 0x2a, 0xc3, 0x40, 0x6b, 0xba, 0xdf, 0x16, 0x3e, 0x40, 0x55, 0xb6, 0x41, 0x8f, 0xc8, 0xae,
	0xb8, 0x1d, 0xc2, 0xea, 0x19, 0xef, 0x9c, 0xec, 0x22, 0x63, 0xd7, 0x80, 0x55, 0x5a, 0xb9, 0x84,
	0x64, 0x22, 0xc2, 0x22, 0x7c, 0x0f, 0x91, 0x56, 0xfe, 0xb7, 0xbd, 0x25, 0x02, 0x26, 0xa2, 0x57,
	0x2e, 0xec, 0xdd, 0x42, 0x53, 0xe6, 0x45, 0xd4, 0x65, 0x84, 0x63, 0xa6, 0x21, 0xc3, 0x34, 0x69,
	0xfb, 0xdf, 0xd9, 0xa7, 0xcf, 0x05, 0xa5, 0x76, 0x4d, 0x7c, 0x53, 0x43, 0xd6, 0x4c, 0xda, 0xde,
	0xdf, 0xd1, 0xa4, 0x52, 0xf6, 0x4c, 0x1a, 0xe3, 0x73, 0x5d, 0x7e, 0xef, 0x0e, 0xb5, 0x52, 0xe6,
	0x18, 0xb6, 0x98, 0x6b, 0xf1, 0x43, 0x03, 0xcd, 0xd7, 0xb6, 0x37, 0x03, 0x8e, 0x63, 0xaa, 0xb4,
	0xa4, 0xa1, 0xfd, 0xc1, 0xfd, 0x1f, 0xac, 0x37, 0xeb, 0x4f, 0xeb, 0xcd, 0xaa, 0x8d, 0xb3, 0xa4,
	0x4f, 0x40, 0x30, 0x97, 0x55, 0x67, 0x8a, 0x01, 0x5f, 0xaf, 0xc9, 0x5a, 0xfc, 0xfd, 0x8c, 0x7b,
	0x99, 0xd7, 0x5d, 0xdd, 0x18, 0xd4, 0xb1, 0x67, 0x4c, 0xed, 0x96, 0x29, 0x5c, 0xbc, 0x61, 0x87,
	0x70, 0xe5, 0xc8, 0x43, 0xa6, 0xe7, 0xfc, 0xce, 0xd2, 0x5b, 0xe8, 0xe6, 0x09, 0x75, 0x6a, 0x8f,
	0xaa, 0x21, 0xeb, 0x76, 0xd7, 0x3f, 0x52, 0xac, 0xf6, 0xc0, 0x5a, 0x45, 0xbe, 0xab, 0x98, 0x11,
	0xdd, 0xc6, 0x8c, 0x2a, 0xdd, 0x7b, 0xcf, 0x0d, 0x3b, 0xcb, 0xb3, 0x78, 0x8b, 0xe8, 0xf6, 0x26,
	0x55, 0xba, 0x7a, 0xd1, 0x3d, 0x46, 0x57, 0x24, 0x44, 0xb9, 0xb4, 0x8f, 0xaa, 0x01, 0xc9, 0xee,
	0xd5, 0x3e, 0x57, 0x71, 0xfa, 0x0a, 0xbc, 0x40, 0x8b, 0x19, 0x23, 0x3a, 0x11, 0x32, 0x2d, 0xef,
	0xe0, 0x01, 0x65, 0xfe, 0x62, 0xcb, 0x5c, 0x2b, 0x99, 0xee, 0x62, 0xec, 0xab, 0xb5, 0x8a, 0x7c,
	0x09, 0x5a, 0x1e, 0x0e, 0xaa, 0x30, 0xe2, 0xba, 0xb0, 0xf8, 0xa0, 0xc4, 0xda, 0x40, 0x29, 0x4f,
	0x44, 0x2f, 0xf1, 0x6c, 0xad, 0x7d, 0x33, 0xc3, 0x26, 0x4f, 0xc4, 0xe0, 0xf6, 0x07, 0x24, 0x8f,
	0x1e, 0x6b, 0xbf, 0xaf, 0xc0, 0x80, 0xf6, 0x07, 0x94, 0x19, 0x1b, 0xd4, 0xfe, 0xf1, 0x5a, 0x8b,
	0x7b, 0xe8, 0xe2, 0x80, 0xcd, 0x1a, 0x71, 0xed, 0xfd, 0x15, 0x9d, 0xab, 0xdf, 0x0d, 0x6e, 0x8f,
	0xa1, 0xb4, 0x77, 0x17, 0xdc, 0x41, 0x1e, 0xcf, 0xd3, 0x10, 0x24, 0x16, 0x49, 0xcf, 0x17, 0x86,
	0x9c, 0xed, 0x38, 0xe4, 0x55, 0x52, 0x7a, 0xc2, 0xe2, 0x1f, 0xc3, 0xe8, 0xf2, 0xc7, 0x8f, 0x85,
	0xf7, 0x53, 0x03, 0x4d, 0xe6, 0x9c, 0x46, 0x44, 0xe9, 0x5e, 0xad, 0xc6, 0xc2, 0xf0, 0xd2, 0xb9,
	0x95, 0xec, 0x54, 0xcf, 0x69, 0xc4, 0x75, 0x30, 0x51, 0x28, 0xa9, 0x0c, 0xed, 0x97, 0x06, 0xf2,
	0x42, 0x29, 0x48, 0x7c, 0x54, 0xdf, 0xd0, 0x67, 0xd2, 0x37, 0x55, 0x69, 0xa9, 0x14, 0x9a, 0xf9,
	0xa5, 0x39, 0xd3, 0xf5, 0x09, 0xfa, 0xc3, 0x9f, 0x6b, 0x7e, 0x95, 0x12, 0xa7, 0x2f, 0x1c, 0xb1,
	0xb5, 0xee, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x89, 0xb3, 0x30, 0x79, 0x10, 0x00, 0x00,
}
