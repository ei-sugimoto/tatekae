// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/project.proto

package proto

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
	ProjectService_Create_FullMethodName = "/protoProject.ProjectService/Create"
	ProjectService_List_FullMethodName   = "/protoProject.ProjectService/List"
	ProjectService_Join_FullMethodName   = "/protoProject.ProjectService/Join"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ProjectService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ProjectService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, ProjectService_Join_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility.
type ProjectServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectServiceServer struct{}

func (UnimplementedProjectServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProjectServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}
func (UnimplementedProjectServiceServer) testEmbeddedByValue()                        {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	// If the following call pancis, it indicates UnimplementedProjectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoProject.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProjectService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProjectService_List_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _ProjectService_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/project.proto",
}
