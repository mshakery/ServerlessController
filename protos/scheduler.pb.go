// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: scheduler.proto

package protos

import (
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

type PodDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *PodDetail) Reset() {
	*x = PodDetail{}
	mi := &file_scheduler_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodDetail) ProtoMessage() {}

func (x *PodDetail) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodDetail.ProtoReflect.Descriptor instead.
func (*PodDetail) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *PodDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodDetail) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_scheduler_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{1}
}

var File_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x3d, 0x0a,
	0x09, 0x50, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x43, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x15,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x11, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_rawDescData = file_scheduler_proto_rawDesc
)

func file_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_rawDescData
}

var file_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scheduler_proto_goTypes = []any{
	(*PodDetail)(nil), // 0: kubernetes.PodDetail
	(*Empty)(nil),     // 1: kubernetes.Empty
}
var file_scheduler_proto_depIdxs = []int32{
	0, // 0: kubernetes.Scheduler.Schedule:input_type -> kubernetes.PodDetail
	1, // 1: kubernetes.Scheduler.Schedule:output_type -> kubernetes.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scheduler_proto_init() }
func file_scheduler_proto_init() {
	if File_scheduler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_proto = out.File
	file_scheduler_proto_rawDesc = nil
	file_scheduler_proto_goTypes = nil
	file_scheduler_proto_depIdxs = nil
}
