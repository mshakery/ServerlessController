// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: kubelet.proto

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

type PodMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CpuUsage    string `protobuf:"bytes,3,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	MemoryUsage string `protobuf:"bytes,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
}

func (x *PodMetrics) Reset() {
	*x = PodMetrics{}
	mi := &file_kubelet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodMetrics) ProtoMessage() {}

func (x *PodMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_kubelet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodMetrics.ProtoReflect.Descriptor instead.
func (*PodMetrics) Descriptor() ([]byte, []int) {
	return file_kubelet_proto_rawDescGZIP(), []int{0}
}

func (x *PodMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodMetrics) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodMetrics) GetCpuUsage() string {
	if x != nil {
		return x.CpuUsage
	}
	return ""
}

func (x *PodMetrics) GetMemoryUsage() string {
	if x != nil {
		return x.MemoryUsage
	}
	return ""
}

type NodeMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodMetrics []*PodMetrics `protobuf:"bytes,1,rep,name=pod_metrics,json=podMetrics,proto3" json:"pod_metrics,omitempty"`
}

func (x *NodeMetrics) Reset() {
	*x = NodeMetrics{}
	mi := &file_kubelet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetrics) ProtoMessage() {}

func (x *NodeMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_kubelet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetrics.ProtoReflect.Descriptor instead.
func (*NodeMetrics) Descriptor() ([]byte, []int) {
	return file_kubelet_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMetrics) GetPodMetrics() []*PodMetrics {
	if x != nil {
		return x.PodMetrics
	}
	return nil
}

var File_kubelet_proto protoreflect.FileDescriptor

var file_kubelet_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x11, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e,
	0x0a, 0x0a, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x37, 0x0a,
	0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x70, 0x6f, 0x64, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x32, 0x74, 0x0a, 0x07, 0x4b, 0x75, 0x62, 0x65, 0x6c, 0x65,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x75, 0x6e,
	0x5f, 0x61, 0x5f, 0x70, 0x6f, 0x64, 0x12, 0x0f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x1a, 0x11, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubelet_proto_rawDescOnce sync.Once
	file_kubelet_proto_rawDescData = file_kubelet_proto_rawDesc
)

func file_kubelet_proto_rawDescGZIP() []byte {
	file_kubelet_proto_rawDescOnce.Do(func() {
		file_kubelet_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubelet_proto_rawDescData)
	})
	return file_kubelet_proto_rawDescData
}

var file_kubelet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kubelet_proto_goTypes = []any{
	(*PodMetrics)(nil),  // 0: kubernetes.PodMetrics
	(*NodeMetrics)(nil), // 1: kubernetes.NodeMetrics
	(*Empty)(nil),       // 2: kubernetes.Empty
	(*Pod)(nil),         // 3: kubernetes.Pod
}
var file_kubelet_proto_depIdxs = []int32{
	0, // 0: kubernetes.NodeMetrics.pod_metrics:type_name -> kubernetes.PodMetrics
	2, // 1: kubernetes.Kubelet.Metric:input_type -> kubernetes.Empty
	3, // 2: kubernetes.Kubelet.run_a_pod:input_type -> kubernetes.Pod
	1, // 3: kubernetes.Kubelet.Metric:output_type -> kubernetes.NodeMetrics
	2, // 4: kubernetes.Kubelet.run_a_pod:output_type -> kubernetes.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kubelet_proto_init() }
func file_kubelet_proto_init() {
	if File_kubelet_proto != nil {
		return
	}
	file_generalView_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kubelet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubelet_proto_goTypes,
		DependencyIndexes: file_kubelet_proto_depIdxs,
		MessageInfos:      file_kubelet_proto_msgTypes,
	}.Build()
	File_kubelet_proto = out.File
	file_kubelet_proto_rawDesc = nil
	file_kubelet_proto_goTypes = nil
	file_kubelet_proto_depIdxs = nil
}
