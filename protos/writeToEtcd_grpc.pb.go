//  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative authentication.proto generalView.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: writeToEtcd.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WriteToEtcd_Apply_FullMethodName = "/kubernetes.WriteToEtcd/Apply"
)

// WriteToEtcdClient is the client API for WriteToEtcd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteToEtcdClient interface {
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*Response, error)
}

type writeToEtcdClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteToEtcdClient(cc grpc.ClientConnInterface) WriteToEtcdClient {
	return &writeToEtcdClient{cc}
}

func (c *writeToEtcdClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, WriteToEtcd_Apply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteToEtcdServer is the server API for WriteToEtcd service.
// All implementations must embed UnimplementedWriteToEtcdServer
// for forward compatibility.
type WriteToEtcdServer interface {
	Apply(context.Context, *ApplyRequest) (*Response, error)
	mustEmbedUnimplementedWriteToEtcdServer()
}

// UnimplementedWriteToEtcdServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWriteToEtcdServer struct{}

func (UnimplementedWriteToEtcdServer) Apply(context.Context, *ApplyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedWriteToEtcdServer) mustEmbedUnimplementedWriteToEtcdServer() {}
func (UnimplementedWriteToEtcdServer) testEmbeddedByValue()                     {}

// UnsafeWriteToEtcdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteToEtcdServer will
// result in compilation errors.
type UnsafeWriteToEtcdServer interface {
	mustEmbedUnimplementedWriteToEtcdServer()
}

func RegisterWriteToEtcdServer(s grpc.ServiceRegistrar, srv WriteToEtcdServer) {
	// If the following call pancis, it indicates UnimplementedWriteToEtcdServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WriteToEtcd_ServiceDesc, srv)
}

func _WriteToEtcd_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteToEtcdServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WriteToEtcd_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteToEtcdServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteToEtcd_ServiceDesc is the grpc.ServiceDesc for WriteToEtcd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteToEtcd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.WriteToEtcd",
	HandlerType: (*WriteToEtcdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _WriteToEtcd_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "writeToEtcd.proto",
}
