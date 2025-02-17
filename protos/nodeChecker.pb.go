// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: nodeChecker.proto

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

type NodeNamee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NodeNamee) Reset() {
	*x = NodeNamee{}
	mi := &file_nodeChecker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeNamee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNamee) ProtoMessage() {}

func (x *NodeNamee) ProtoReflect() protoreflect.Message {
	mi := &file_nodeChecker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNamee.ProtoReflect.Descriptor instead.
func (*NodeNamee) Descriptor() ([]byte, []int) {
	return file_nodeChecker_proto_rawDescGZIP(), []int{0}
}

func (x *NodeNamee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_nodeChecker_proto protoreflect.FileDescriptor

var file_nodeChecker_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a,
	0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x46, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x15, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x65, 0x1a, 0x11, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeChecker_proto_rawDescOnce sync.Once
	file_nodeChecker_proto_rawDescData = file_nodeChecker_proto_rawDesc
)

func file_nodeChecker_proto_rawDescGZIP() []byte {
	file_nodeChecker_proto_rawDescOnce.Do(func() {
		file_nodeChecker_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeChecker_proto_rawDescData)
	})
	return file_nodeChecker_proto_rawDescData
}

var file_nodeChecker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nodeChecker_proto_goTypes = []any{
	(*NodeNamee)(nil), // 0: kubernetes.NodeNamee
	(*Empty)(nil),     // 1: kubernetes.Empty
}
var file_nodeChecker_proto_depIdxs = []int32{
	0, // 0: kubernetes.NodeChecker.CheckNode:input_type -> kubernetes.NodeNamee
	1, // 1: kubernetes.NodeChecker.CheckNode:output_type -> kubernetes.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nodeChecker_proto_init() }
func file_nodeChecker_proto_init() {
	if File_nodeChecker_proto != nil {
		return
	}
	file_generalView_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodeChecker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodeChecker_proto_goTypes,
		DependencyIndexes: file_nodeChecker_proto_depIdxs,
		MessageInfos:      file_nodeChecker_proto_msgTypes,
	}.Build()
	File_nodeChecker_proto = out.File
	file_nodeChecker_proto_rawDesc = nil
	file_nodeChecker_proto_goTypes = nil
	file_nodeChecker_proto_depIdxs = nil
}
