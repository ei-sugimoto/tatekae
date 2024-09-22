// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/bill.proto

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
	BillService_Create_FullMethodName           = "/protoBill.BillService/Create"
	BillService_ListByProject_FullMethodName    = "/protoBill.BillService/ListByProject"
	BillService_SumrizeByProject_FullMethodName = "/protoBill.BillService/SumrizeByProject"
)

// BillServiceClient is the client API for BillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillServiceClient interface {
	Create(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error)
	ListByProject(ctx context.Context, in *ListByProjectRequest, opts ...grpc.CallOption) (*ListByProjectResponse, error)
	SumrizeByProject(ctx context.Context, in *SumrizeRequest, opts ...grpc.CallOption) (*SumrizeResponse, error)
}

type billServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillServiceClient(cc grpc.ClientConnInterface) BillServiceClient {
	return &billServiceClient{cc}
}

func (c *billServiceClient) Create(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBillResponse)
	err := c.cc.Invoke(ctx, BillService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) ListByProject(ctx context.Context, in *ListByProjectRequest, opts ...grpc.CallOption) (*ListByProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListByProjectResponse)
	err := c.cc.Invoke(ctx, BillService_ListByProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) SumrizeByProject(ctx context.Context, in *SumrizeRequest, opts ...grpc.CallOption) (*SumrizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumrizeResponse)
	err := c.cc.Invoke(ctx, BillService_SumrizeByProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillServiceServer is the server API for BillService service.
// All implementations must embed UnimplementedBillServiceServer
// for forward compatibility.
type BillServiceServer interface {
	Create(context.Context, *CreateBillRequest) (*CreateBillResponse, error)
	ListByProject(context.Context, *ListByProjectRequest) (*ListByProjectResponse, error)
	SumrizeByProject(context.Context, *SumrizeRequest) (*SumrizeResponse, error)
	mustEmbedUnimplementedBillServiceServer()
}

// UnimplementedBillServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBillServiceServer struct{}

func (UnimplementedBillServiceServer) Create(context.Context, *CreateBillRequest) (*CreateBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBillServiceServer) ListByProject(context.Context, *ListByProjectRequest) (*ListByProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByProject not implemented")
}
func (UnimplementedBillServiceServer) SumrizeByProject(context.Context, *SumrizeRequest) (*SumrizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumrizeByProject not implemented")
}
func (UnimplementedBillServiceServer) mustEmbedUnimplementedBillServiceServer() {}
func (UnimplementedBillServiceServer) testEmbeddedByValue()                     {}

// UnsafeBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillServiceServer will
// result in compilation errors.
type UnsafeBillServiceServer interface {
	mustEmbedUnimplementedBillServiceServer()
}

func RegisterBillServiceServer(s grpc.ServiceRegistrar, srv BillServiceServer) {
	// If the following call pancis, it indicates UnimplementedBillServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BillService_ServiceDesc, srv)
}

func _BillService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).Create(ctx, req.(*CreateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_ListByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).ListByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_ListByProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).ListByProject(ctx, req.(*ListByProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_SumrizeByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumrizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).SumrizeByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_SumrizeByProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).SumrizeByProject(ctx, req.(*SumrizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillService_ServiceDesc is the grpc.ServiceDesc for BillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoBill.BillService",
	HandlerType: (*BillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BillService_Create_Handler,
		},
		{
			MethodName: "ListByProject",
			Handler:    _BillService_ListByProject_Handler,
		},
		{
			MethodName: "SumrizeByProject",
			Handler:    _BillService_SumrizeByProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bill.proto",
}
