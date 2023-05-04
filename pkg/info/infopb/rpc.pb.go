// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: info/infopb/rpc.proto

package infopb

import (
	labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelSets     []*labelpb.ZLabelSet `protobuf:"bytes,1,rep,name=label_sets,json=labelSets,proto3" json:"label_sets,omitempty"`
	ComponentType string               `protobuf:"bytes,2,opt,name=ComponentType,proto3" json:"ComponentType,omitempty"`
	// StoreInfo holds the metadata related to Store API if exposed by the component otherwise it will be null.
	Store *StoreInfo `protobuf:"bytes,3,opt,name=store,proto3" json:"store,omitempty"`
	// RulesInfo holds the metadata related to Rules API if exposed by the component otherwise it will be null.
	Rules *RulesInfo `protobuf:"bytes,4,opt,name=rules,proto3" json:"rules,omitempty"`
	// MetricMetadataInfo holds the metadata related to Metadata API if exposed by the component otherwise it will be null.
	MetricMetadata *MetricMetadataInfo `protobuf:"bytes,5,opt,name=metric_metadata,json=metricMetadata,proto3" json:"metric_metadata,omitempty"`
	// TargetsInfo holds the metadata related to Targets API if exposed by the component otherwise it will be null.
	Targets *TargetsInfo `protobuf:"bytes,6,opt,name=targets,proto3" json:"targets,omitempty"`
	// ExemplarsInfo holds the metadata related to Exemplars API if exposed by the component otherwise it will be null.
	Exemplars *ExemplarsInfo `protobuf:"bytes,7,opt,name=exemplars,proto3" json:"exemplars,omitempty"`
	// QueryAPIInfo holds the metadata related to Query API if exposed by the component, otherwise it will be null.
	Query *QueryAPIInfo `protobuf:"bytes,8,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetLabelSets() []*labelpb.ZLabelSet {
	if x != nil {
		return x.LabelSets
	}
	return nil
}

func (x *InfoResponse) GetComponentType() string {
	if x != nil {
		return x.ComponentType
	}
	return ""
}

func (x *InfoResponse) GetStore() *StoreInfo {
	if x != nil {
		return x.Store
	}
	return nil
}

func (x *InfoResponse) GetRules() *RulesInfo {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *InfoResponse) GetMetricMetadata() *MetricMetadataInfo {
	if x != nil {
		return x.MetricMetadata
	}
	return nil
}

func (x *InfoResponse) GetTargets() *TargetsInfo {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *InfoResponse) GetExemplars() *ExemplarsInfo {
	if x != nil {
		return x.Exemplars
	}
	return nil
}

func (x *InfoResponse) GetQuery() *QueryAPIInfo {
	if x != nil {
		return x.Query
	}
	return nil
}

// StoreInfo holds the metadata related to Store API exposed by the component.
type StoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinTime          int64 `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime          int64 `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	SupportsSharding bool  `protobuf:"varint,3,opt,name=supports_sharding,json=supportsSharding,proto3" json:"supports_sharding,omitempty"`
	// replica_aware means this store supports without_replica_labels of StoreAPI.Series.
	SupportsWithoutReplicaLabels bool `protobuf:"varint,5,opt,name=supports_without_replica_labels,json=supportsWithoutReplicaLabels,proto3" json:"supports_without_replica_labels,omitempty"`
	// TSDBInfos holds metadata for all TSDBs exposed by the store.
	TsdbInfos []*TSDBInfo `protobuf:"bytes,6,rep,name=tsdb_infos,json=tsdbInfos,proto3" json:"tsdb_infos,omitempty"`
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *StoreInfo) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *StoreInfo) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

func (x *StoreInfo) GetSupportsSharding() bool {
	if x != nil {
		return x.SupportsSharding
	}
	return false
}

func (x *StoreInfo) GetSupportsWithoutReplicaLabels() bool {
	if x != nil {
		return x.SupportsWithoutReplicaLabels
	}
	return false
}

func (x *StoreInfo) GetTsdbInfos() []*TSDBInfo {
	if x != nil {
		return x.TsdbInfos
	}
	return nil
}

// RulesInfo holds the metadata related to Rules API exposed by the component.
type RulesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RulesInfo) Reset() {
	*x = RulesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulesInfo) ProtoMessage() {}

func (x *RulesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulesInfo.ProtoReflect.Descriptor instead.
func (*RulesInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{3}
}

// MetricMetadataInfo holds the metadata related to Metadata API exposed by the component.
type MetricMetadataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetricMetadataInfo) Reset() {
	*x = MetricMetadataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadataInfo) ProtoMessage() {}

func (x *MetricMetadataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadataInfo.ProtoReflect.Descriptor instead.
func (*MetricMetadataInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{4}
}

// TargetsInfo holds the metadata related to Targets API exposed by the component.
type TargetsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TargetsInfo) Reset() {
	*x = TargetsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetsInfo) ProtoMessage() {}

func (x *TargetsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetsInfo.ProtoReflect.Descriptor instead.
func (*TargetsInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{5}
}

// ExemplarsInfo holds the metadata related to Exemplars API exposed by the component.
type ExemplarsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinTime int64 `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64 `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
}

func (x *ExemplarsInfo) Reset() {
	*x = ExemplarsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExemplarsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExemplarsInfo) ProtoMessage() {}

func (x *ExemplarsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExemplarsInfo.ProtoReflect.Descriptor instead.
func (*ExemplarsInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *ExemplarsInfo) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *ExemplarsInfo) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

// QueryInfo holds the metadata related to Query API exposed by the component.
type QueryAPIInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryAPIInfo) Reset() {
	*x = QueryAPIInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAPIInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAPIInfo) ProtoMessage() {}

func (x *QueryAPIInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAPIInfo.ProtoReflect.Descriptor instead.
func (*QueryAPIInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{7}
}

type TSDBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels  *labelpb.ZLabelSet `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	MinTime int64              `protobuf:"varint,2,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64              `protobuf:"varint,3,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
}

func (x *TSDBInfo) Reset() {
	*x = TSDBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_infopb_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TSDBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TSDBInfo) ProtoMessage() {}

func (x *TSDBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_info_infopb_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TSDBInfo.ProtoReflect.Descriptor instead.
func (*TSDBInfo) Descriptor() ([]byte, []int) {
	return file_info_infopb_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *TSDBInfo) GetLabels() *labelpb.ZLabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TSDBInfo) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *TSDBInfo) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

var File_info_infopb_rpc_proto protoreflect.FileDescriptor

var file_info_infopb_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xab,
	0x03, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x5a, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74,
	0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x50,
	0x49, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0xf1, 0x01, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a,
	0x1f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x73, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f,
	0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x53, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x74, 0x73, 0x64, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x22, 0x0b, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x14, 0x0a,
	0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6b, 0x0a, 0x08, 0x54, 0x53, 0x44,
	0x42, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x5a,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x43, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73,
	0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_info_infopb_rpc_proto_rawDescOnce sync.Once
	file_info_infopb_rpc_proto_rawDescData = file_info_infopb_rpc_proto_rawDesc
)

func file_info_infopb_rpc_proto_rawDescGZIP() []byte {
	file_info_infopb_rpc_proto_rawDescOnce.Do(func() {
		file_info_infopb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_infopb_rpc_proto_rawDescData)
	})
	return file_info_infopb_rpc_proto_rawDescData
}

var file_info_infopb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_info_infopb_rpc_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),        // 0: thanos.info.InfoRequest
	(*InfoResponse)(nil),       // 1: thanos.info.InfoResponse
	(*StoreInfo)(nil),          // 2: thanos.info.StoreInfo
	(*RulesInfo)(nil),          // 3: thanos.info.RulesInfo
	(*MetricMetadataInfo)(nil), // 4: thanos.info.MetricMetadataInfo
	(*TargetsInfo)(nil),        // 5: thanos.info.TargetsInfo
	(*ExemplarsInfo)(nil),      // 6: thanos.info.ExemplarsInfo
	(*QueryAPIInfo)(nil),       // 7: thanos.info.QueryAPIInfo
	(*TSDBInfo)(nil),           // 8: thanos.info.TSDBInfo
	(*labelpb.ZLabelSet)(nil),  // 9: thanos.ZLabelSet
}
var file_info_infopb_rpc_proto_depIdxs = []int32{
	9,  // 0: thanos.info.InfoResponse.label_sets:type_name -> thanos.ZLabelSet
	2,  // 1: thanos.info.InfoResponse.store:type_name -> thanos.info.StoreInfo
	3,  // 2: thanos.info.InfoResponse.rules:type_name -> thanos.info.RulesInfo
	4,  // 3: thanos.info.InfoResponse.metric_metadata:type_name -> thanos.info.MetricMetadataInfo
	5,  // 4: thanos.info.InfoResponse.targets:type_name -> thanos.info.TargetsInfo
	6,  // 5: thanos.info.InfoResponse.exemplars:type_name -> thanos.info.ExemplarsInfo
	7,  // 6: thanos.info.InfoResponse.query:type_name -> thanos.info.QueryAPIInfo
	8,  // 7: thanos.info.StoreInfo.tsdb_infos:type_name -> thanos.info.TSDBInfo
	9,  // 8: thanos.info.TSDBInfo.labels:type_name -> thanos.ZLabelSet
	0,  // 9: thanos.info.Info.Info:input_type -> thanos.info.InfoRequest
	1,  // 10: thanos.info.Info.Info:output_type -> thanos.info.InfoResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_info_infopb_rpc_proto_init() }
func file_info_infopb_rpc_proto_init() {
	if File_info_infopb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_infopb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulesInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadataInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExemplarsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAPIInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_infopb_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TSDBInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_info_infopb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_infopb_rpc_proto_goTypes,
		DependencyIndexes: file_info_infopb_rpc_proto_depIdxs,
		MessageInfos:      file_info_infopb_rpc_proto_msgTypes,
	}.Build()
	File_info_infopb_rpc_proto = out.File
	file_info_infopb_rpc_proto_rawDesc = nil
	file_info_infopb_rpc_proto_goTypes = nil
	file_info_infopb_rpc_proto_depIdxs = nil
}
