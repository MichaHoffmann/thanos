// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: exemplars/exemplarspb/rpc.proto

package exemplarspb

import (
	labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
	storepb "github.com/thanos-io/thanos/pkg/store/storepb"
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

type ExemplarsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query                   string                          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Start                   int64                           `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                     int64                           `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	PartialResponseStrategy storepb.PartialResponseStrategy `protobuf:"varint,4,opt,name=partial_response_strategy,json=partialResponseStrategy,proto3,enum=thanos.PartialResponseStrategy" json:"partial_response_strategy,omitempty"`
}

func (x *ExemplarsRequest) Reset() {
	*x = ExemplarsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExemplarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExemplarsRequest) ProtoMessage() {}

func (x *ExemplarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExemplarsRequest.ProtoReflect.Descriptor instead.
func (*ExemplarsRequest) Descriptor() ([]byte, []int) {
	return file_exemplars_exemplarspb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ExemplarsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ExemplarsRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ExemplarsRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ExemplarsRequest) GetPartialResponseStrategy() storepb.PartialResponseStrategy {
	if x != nil {
		return x.PartialResponseStrategy
	}
	return storepb.PartialResponseStrategy(0)
}

type ExemplarsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*ExemplarsResponse_Data
	//	*ExemplarsResponse_Warning
	Result isExemplarsResponse_Result `protobuf_oneof:"result"`
}

func (x *ExemplarsResponse) Reset() {
	*x = ExemplarsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExemplarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExemplarsResponse) ProtoMessage() {}

func (x *ExemplarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExemplarsResponse.ProtoReflect.Descriptor instead.
func (*ExemplarsResponse) Descriptor() ([]byte, []int) {
	return file_exemplars_exemplarspb_rpc_proto_rawDescGZIP(), []int{1}
}

func (m *ExemplarsResponse) GetResult() isExemplarsResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ExemplarsResponse) GetData() *ExemplarData {
	if x, ok := x.GetResult().(*ExemplarsResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *ExemplarsResponse) GetWarning() string {
	if x, ok := x.GetResult().(*ExemplarsResponse_Warning); ok {
		return x.Warning
	}
	return ""
}

type isExemplarsResponse_Result interface {
	isExemplarsResponse_Result()
}

type ExemplarsResponse_Data struct {
	Data *ExemplarData `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type ExemplarsResponse_Warning struct {
	Warning string `protobuf:"bytes,2,opt,name=warning,proto3,oneof"`
}

func (*ExemplarsResponse_Data) isExemplarsResponse_Result() {}

func (*ExemplarsResponse_Warning) isExemplarsResponse_Result() {}

type ExemplarData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"seriesLabels"
	SeriesLabels *labelpb.ZLabelSet `protobuf:"bytes,1,opt,name=seriesLabels,proto3" json:"seriesLabels"`
	// @gotags: json:"exemplars"
	Exemplars []*Exemplar `protobuf:"bytes,2,rep,name=exemplars,proto3" json:"exemplars"`
}

func (x *ExemplarData) Reset() {
	*x = ExemplarData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExemplarData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExemplarData) ProtoMessage() {}

func (x *ExemplarData) ProtoReflect() protoreflect.Message {
	mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExemplarData.ProtoReflect.Descriptor instead.
func (*ExemplarData) Descriptor() ([]byte, []int) {
	return file_exemplars_exemplarspb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ExemplarData) GetSeriesLabels() *labelpb.ZLabelSet {
	if x != nil {
		return x.SeriesLabels
	}
	return nil
}

func (x *ExemplarData) GetExemplars() []*Exemplar {
	if x != nil {
		return x.Exemplars
	}
	return nil
}

type Exemplar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"labels"
	Labels *labelpb.ZLabelSet `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels"`
	// @gotags: json:"value"
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value"`
	// @gotags: json:"timestamp"
	Ts int64 `protobuf:"varint,3,opt,name=ts,proto3" json:"timestamp"`
}

func (x *Exemplar) Reset() {
	*x = Exemplar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exemplar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exemplar) ProtoMessage() {}

func (x *Exemplar) ProtoReflect() protoreflect.Message {
	mi := &file_exemplars_exemplarspb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exemplar.ProtoReflect.Descriptor instead.
func (*Exemplar) Descriptor() ([]byte, []int) {
	return file_exemplars_exemplarspb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Exemplar) GetLabels() *labelpb.ZLabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Exemplar) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Exemplar) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

var File_exemplars_exemplarspb_rpc_proto protoreflect.FileDescriptor

var file_exemplars_exemplarspb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x5b, 0x0a, 0x19, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22,
	0x65, 0x0a, 0x11, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x75, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x5a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2e, 0x0a,
	0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x72, 0x52, 0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x22, 0x5b, 0x0a,
	0x08, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e,
	0x6f, 0x73, 0x2e, 0x5a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x32, 0x4f, 0x0a, 0x09, 0x45, 0x78,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x78,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73,
	0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65,
	0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exemplars_exemplarspb_rpc_proto_rawDescOnce sync.Once
	file_exemplars_exemplarspb_rpc_proto_rawDescData = file_exemplars_exemplarspb_rpc_proto_rawDesc
)

func file_exemplars_exemplarspb_rpc_proto_rawDescGZIP() []byte {
	file_exemplars_exemplarspb_rpc_proto_rawDescOnce.Do(func() {
		file_exemplars_exemplarspb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_exemplars_exemplarspb_rpc_proto_rawDescData)
	})
	return file_exemplars_exemplarspb_rpc_proto_rawDescData
}

var file_exemplars_exemplarspb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_exemplars_exemplarspb_rpc_proto_goTypes = []interface{}{
	(*ExemplarsRequest)(nil),             // 0: thanos.ExemplarsRequest
	(*ExemplarsResponse)(nil),            // 1: thanos.ExemplarsResponse
	(*ExemplarData)(nil),                 // 2: thanos.ExemplarData
	(*Exemplar)(nil),                     // 3: thanos.Exemplar
	(storepb.PartialResponseStrategy)(0), // 4: thanos.PartialResponseStrategy
	(*labelpb.ZLabelSet)(nil),            // 5: thanos.ZLabelSet
}
var file_exemplars_exemplarspb_rpc_proto_depIdxs = []int32{
	4, // 0: thanos.ExemplarsRequest.partial_response_strategy:type_name -> thanos.PartialResponseStrategy
	2, // 1: thanos.ExemplarsResponse.data:type_name -> thanos.ExemplarData
	5, // 2: thanos.ExemplarData.seriesLabels:type_name -> thanos.ZLabelSet
	3, // 3: thanos.ExemplarData.exemplars:type_name -> thanos.Exemplar
	5, // 4: thanos.Exemplar.labels:type_name -> thanos.ZLabelSet
	0, // 5: thanos.Exemplars.Exemplars:input_type -> thanos.ExemplarsRequest
	1, // 6: thanos.Exemplars.Exemplars:output_type -> thanos.ExemplarsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_exemplars_exemplarspb_rpc_proto_init() }
func file_exemplars_exemplarspb_rpc_proto_init() {
	if File_exemplars_exemplarspb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exemplars_exemplarspb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExemplarsRequest); i {
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
		file_exemplars_exemplarspb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExemplarsResponse); i {
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
		file_exemplars_exemplarspb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExemplarData); i {
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
		file_exemplars_exemplarspb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exemplar); i {
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
	file_exemplars_exemplarspb_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExemplarsResponse_Data)(nil),
		(*ExemplarsResponse_Warning)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exemplars_exemplarspb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exemplars_exemplarspb_rpc_proto_goTypes,
		DependencyIndexes: file_exemplars_exemplarspb_rpc_proto_depIdxs,
		MessageInfos:      file_exemplars_exemplarspb_rpc_proto_msgTypes,
	}.Build()
	File_exemplars_exemplarspb_rpc_proto = out.File
	file_exemplars_exemplarspb_rpc_proto_rawDesc = nil
	file_exemplars_exemplarspb_rpc_proto_goTypes = nil
	file_exemplars_exemplarspb_rpc_proto_depIdxs = nil
}
