// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: nodeChecker.proto

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
	NodeChecker_CheckNode_FullMethodName = "/kubernetes.NodeChecker/CheckNode"
)

// NodeCheckerClient is the client API for NodeChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeCheckerClient interface {
	CheckNode(ctx context.Context, in *NodeNamee, opts ...grpc.CallOption) (*Empty, error)
}

type nodeCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeCheckerClient(cc grpc.ClientConnInterface) NodeCheckerClient {
	return &nodeCheckerClient{cc}
}

func (c *nodeCheckerClient) CheckNode(ctx context.Context, in *NodeNamee, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, NodeChecker_CheckNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeCheckerServer is the server API for NodeChecker service.
// All implementations must embed UnimplementedNodeCheckerServer
// for forward compatibility.
type NodeCheckerServer interface {
	CheckNode(context.Context, *NodeNamee) (*Empty, error)
	mustEmbedUnimplementedNodeCheckerServer()
}

// UnimplementedNodeCheckerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeCheckerServer struct{}

func (UnimplementedNodeCheckerServer) CheckNode(context.Context, *NodeNamee) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNode not implemented")
}
func (UnimplementedNodeCheckerServer) mustEmbedUnimplementedNodeCheckerServer() {}
func (UnimplementedNodeCheckerServer) testEmbeddedByValue()                     {}

// UnsafeNodeCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeCheckerServer will
// result in compilation errors.
type UnsafeNodeCheckerServer interface {
	mustEmbedUnimplementedNodeCheckerServer()
}

func RegisterNodeCheckerServer(s grpc.ServiceRegistrar, srv NodeCheckerServer) {
	// If the following call pancis, it indicates UnimplementedNodeCheckerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NodeChecker_ServiceDesc, srv)
}

func _NodeChecker_CheckNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeNamee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCheckerServer).CheckNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeChecker_CheckNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCheckerServer).CheckNode(ctx, req.(*NodeNamee))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeChecker_ServiceDesc is the grpc.ServiceDesc for NodeChecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeChecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.NodeChecker",
	HandlerType: (*NodeCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckNode",
			Handler:    _NodeChecker_CheckNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeChecker.proto",
}
