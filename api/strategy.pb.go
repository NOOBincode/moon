// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: strategy.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 持续类型定义
//
//	m时间内出现n次
//	m时间内最多出现n次
//	m时间内最少出现n次
type MetricSustainType int32

const (
	// 持续类型定义
	MetricSustainType_MetricSustainType_UNKNOWN MetricSustainType = 0
	// m时间内出现n次
	MetricSustainType_MetricSustainType_FOR MetricSustainType = 1
	// m时间内最多出现n次
	MetricSustainType_MetricSustainType_MAX MetricSustainType = 2
	// m时间内最少出现n次
	MetricSustainType_MetricSustainType_MIN MetricSustainType = 3
)

// Enum value maps for MetricSustainType.
var (
	MetricSustainType_name = map[int32]string{
		0: "MetricSustainType_UNKNOWN",
		1: "MetricSustainType_FOR",
		2: "MetricSustainType_MAX",
		3: "MetricSustainType_MIN",
	}
	MetricSustainType_value = map[string]int32{
		"MetricSustainType_UNKNOWN": 0,
		"MetricSustainType_FOR":     1,
		"MetricSustainType_MAX":     2,
		"MetricSustainType_MIN":     3,
	}
)

func (x MetricSustainType) Enum() *MetricSustainType {
	p := new(MetricSustainType)
	*p = x
	return p
}

func (x MetricSustainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricSustainType) Descriptor() protoreflect.EnumDescriptor {
	return file_strategy_proto_enumTypes[0].Descriptor()
}

func (MetricSustainType) Type() protoreflect.EnumType {
	return &file_strategy_proto_enumTypes[0]
}

func (x MetricSustainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricSustainType.Descriptor instead.
func (MetricSustainType) EnumDescriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{0}
}

type MetricStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 策略名称
	Alert string `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	// 策略语句
	Expr string `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	// 策略持续时间
	For *durationpb.Duration `protobuf:"bytes,3,opt,name=for,proto3" json:"for,omitempty"`
	// 持续次数
	Count uint32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	// 持续的类型
	SustainType MetricSustainType `protobuf:"varint,5,opt,name=sustainType,proto3,enum=api.MetricSustainType" json:"sustainType,omitempty"`
	// 策略标签
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 策略注解
	Annotations map[string]string `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 执行频率
	Interval *durationpb.Duration `protobuf:"bytes,8,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *MetricStrategy) Reset() {
	*x = MetricStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStrategy) ProtoMessage() {}

func (x *MetricStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStrategy.ProtoReflect.Descriptor instead.
func (*MetricStrategy) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *MetricStrategy) GetAlert() string {
	if x != nil {
		return x.Alert
	}
	return ""
}

func (x *MetricStrategy) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *MetricStrategy) GetFor() *durationpb.Duration {
	if x != nil {
		return x.For
	}
	return nil
}

func (x *MetricStrategy) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MetricStrategy) GetSustainType() MetricSustainType {
	if x != nil {
		return x.SustainType
	}
	return MetricSustainType_MetricSustainType_UNKNOWN
}

func (x *MetricStrategy) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricStrategy) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *MetricStrategy) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

var File_strategy_proto protoreflect.FileDescriptor

var file_strategy_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x2b, 0x0a,
	0x03, 0x66, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x66, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x38, 0x0a, 0x0b, 0x73, 0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73,
	0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x83, 0x01,
	0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x75, 0x73,
	0x74, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x75, 0x73, 0x74,
	0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4d, 0x49,
	0x4e, 0x10, 0x03, 0x42, 0x2c, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x2f, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strategy_proto_rawDescOnce sync.Once
	file_strategy_proto_rawDescData = file_strategy_proto_rawDesc
)

func file_strategy_proto_rawDescGZIP() []byte {
	file_strategy_proto_rawDescOnce.Do(func() {
		file_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_proto_rawDescData)
	})
	return file_strategy_proto_rawDescData
}

var file_strategy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_strategy_proto_goTypes = []interface{}{
	(MetricSustainType)(0),      // 0: api.MetricSustainType
	(*MetricStrategy)(nil),      // 1: api.MetricStrategy
	nil,                         // 2: api.MetricStrategy.LabelsEntry
	nil,                         // 3: api.MetricStrategy.AnnotationsEntry
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_strategy_proto_depIdxs = []int32{
	4, // 0: api.MetricStrategy.for:type_name -> google.protobuf.Duration
	0, // 1: api.MetricStrategy.sustainType:type_name -> api.MetricSustainType
	2, // 2: api.MetricStrategy.labels:type_name -> api.MetricStrategy.LabelsEntry
	3, // 3: api.MetricStrategy.annotations:type_name -> api.MetricStrategy.AnnotationsEntry
	4, // 4: api.MetricStrategy.interval:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_strategy_proto_init() }
func file_strategy_proto_init() {
	if File_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricStrategy); i {
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
			RawDescriptor: file_strategy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_strategy_proto_goTypes,
		DependencyIndexes: file_strategy_proto_depIdxs,
		EnumInfos:         file_strategy_proto_enumTypes,
		MessageInfos:      file_strategy_proto_msgTypes,
	}.Build()
	File_strategy_proto = out.File
	file_strategy_proto_rawDesc = nil
	file_strategy_proto_goTypes = nil
	file_strategy_proto_depIdxs = nil
}
