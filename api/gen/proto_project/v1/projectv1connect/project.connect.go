// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto_project/v1/project.proto

package projectv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ei-sugimoto/tatekae/api/gen/proto_project/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProjectServiceName is the fully-qualified name of the ProjectService service.
	ProjectServiceName = "proto_project.v1.ProjectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProjectServiceCreateProcedure is the fully-qualified name of the ProjectService's Create RPC.
	ProjectServiceCreateProcedure = "/proto_project.v1.ProjectService/Create"
	// ProjectServiceListProcedure is the fully-qualified name of the ProjectService's List RPC.
	ProjectServiceListProcedure = "/proto_project.v1.ProjectService/List"
	// ProjectServiceJoinProcedure is the fully-qualified name of the ProjectService's Join RPC.
	ProjectServiceJoinProcedure = "/proto_project.v1.ProjectService/Join"
	// ProjectServiceGetProcedure is the fully-qualified name of the ProjectService's Get RPC.
	ProjectServiceGetProcedure = "/proto_project.v1.ProjectService/Get"
	// ProjectServiceJoinListProcedure is the fully-qualified name of the ProjectService's JoinList RPC.
	ProjectServiceJoinListProcedure = "/proto_project.v1.ProjectService/JoinList"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	projectServiceServiceDescriptor        = v1.File_proto_project_v1_project_proto.Services().ByName("ProjectService")
	projectServiceCreateMethodDescriptor   = projectServiceServiceDescriptor.Methods().ByName("Create")
	projectServiceListMethodDescriptor     = projectServiceServiceDescriptor.Methods().ByName("List")
	projectServiceJoinMethodDescriptor     = projectServiceServiceDescriptor.Methods().ByName("Join")
	projectServiceGetMethodDescriptor      = projectServiceServiceDescriptor.Methods().ByName("Get")
	projectServiceJoinListMethodDescriptor = projectServiceServiceDescriptor.Methods().ByName("JoinList")
)

// ProjectServiceClient is a client for the proto_project.v1.ProjectService service.
type ProjectServiceClient interface {
	Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error)
	List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error)
	Join(context.Context, *connect.Request[v1.ProjectServiceJoinRequest]) (*connect.Response[v1.ProjectServiceJoinResponse], error)
	Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error)
	JoinList(context.Context, *connect.Request[v1.ProjectServiceJoinListRequest]) (*connect.Response[v1.ProjectServiceJoinListResponse], error)
}

// NewProjectServiceClient constructs a client for the proto_project.v1.ProjectService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &projectServiceClient{
		create: connect.NewClient[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse](
			httpClient,
			baseURL+ProjectServiceCreateProcedure,
			connect.WithSchema(projectServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse](
			httpClient,
			baseURL+ProjectServiceListProcedure,
			connect.WithSchema(projectServiceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		join: connect.NewClient[v1.ProjectServiceJoinRequest, v1.ProjectServiceJoinResponse](
			httpClient,
			baseURL+ProjectServiceJoinProcedure,
			connect.WithSchema(projectServiceJoinMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse](
			httpClient,
			baseURL+ProjectServiceGetProcedure,
			connect.WithSchema(projectServiceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		joinList: connect.NewClient[v1.ProjectServiceJoinListRequest, v1.ProjectServiceJoinListResponse](
			httpClient,
			baseURL+ProjectServiceJoinListProcedure,
			connect.WithSchema(projectServiceJoinListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// projectServiceClient implements ProjectServiceClient.
type projectServiceClient struct {
	create   *connect.Client[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse]
	list     *connect.Client[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse]
	join     *connect.Client[v1.ProjectServiceJoinRequest, v1.ProjectServiceJoinResponse]
	get      *connect.Client[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse]
	joinList *connect.Client[v1.ProjectServiceJoinListRequest, v1.ProjectServiceJoinListResponse]
}

// Create calls proto_project.v1.ProjectService.Create.
func (c *projectServiceClient) Create(ctx context.Context, req *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// List calls proto_project.v1.ProjectService.List.
func (c *projectServiceClient) List(ctx context.Context, req *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Join calls proto_project.v1.ProjectService.Join.
func (c *projectServiceClient) Join(ctx context.Context, req *connect.Request[v1.ProjectServiceJoinRequest]) (*connect.Response[v1.ProjectServiceJoinResponse], error) {
	return c.join.CallUnary(ctx, req)
}

// Get calls proto_project.v1.ProjectService.Get.
func (c *projectServiceClient) Get(ctx context.Context, req *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// JoinList calls proto_project.v1.ProjectService.JoinList.
func (c *projectServiceClient) JoinList(ctx context.Context, req *connect.Request[v1.ProjectServiceJoinListRequest]) (*connect.Response[v1.ProjectServiceJoinListResponse], error) {
	return c.joinList.CallUnary(ctx, req)
}

// ProjectServiceHandler is an implementation of the proto_project.v1.ProjectService service.
type ProjectServiceHandler interface {
	Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error)
	List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error)
	Join(context.Context, *connect.Request[v1.ProjectServiceJoinRequest]) (*connect.Response[v1.ProjectServiceJoinResponse], error)
	Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error)
	JoinList(context.Context, *connect.Request[v1.ProjectServiceJoinListRequest]) (*connect.Response[v1.ProjectServiceJoinListResponse], error)
}

// NewProjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectServiceHandler(svc ProjectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	projectServiceCreateHandler := connect.NewUnaryHandler(
		ProjectServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(projectServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceListHandler := connect.NewUnaryHandler(
		ProjectServiceListProcedure,
		svc.List,
		connect.WithSchema(projectServiceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceJoinHandler := connect.NewUnaryHandler(
		ProjectServiceJoinProcedure,
		svc.Join,
		connect.WithSchema(projectServiceJoinMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceGetHandler := connect.NewUnaryHandler(
		ProjectServiceGetProcedure,
		svc.Get,
		connect.WithSchema(projectServiceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceJoinListHandler := connect.NewUnaryHandler(
		ProjectServiceJoinListProcedure,
		svc.JoinList,
		connect.WithSchema(projectServiceJoinListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto_project.v1.ProjectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProjectServiceCreateProcedure:
			projectServiceCreateHandler.ServeHTTP(w, r)
		case ProjectServiceListProcedure:
			projectServiceListHandler.ServeHTTP(w, r)
		case ProjectServiceJoinProcedure:
			projectServiceJoinHandler.ServeHTTP(w, r)
		case ProjectServiceGetProcedure:
			projectServiceGetHandler.ServeHTTP(w, r)
		case ProjectServiceJoinListProcedure:
			projectServiceJoinListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectServiceHandler struct{}

func (UnimplementedProjectServiceHandler) Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto_project.v1.ProjectService.Create is not implemented"))
}

func (UnimplementedProjectServiceHandler) List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto_project.v1.ProjectService.List is not implemented"))
}

func (UnimplementedProjectServiceHandler) Join(context.Context, *connect.Request[v1.ProjectServiceJoinRequest]) (*connect.Response[v1.ProjectServiceJoinResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto_project.v1.ProjectService.Join is not implemented"))
}

func (UnimplementedProjectServiceHandler) Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto_project.v1.ProjectService.Get is not implemented"))
}

func (UnimplementedProjectServiceHandler) JoinList(context.Context, *connect.Request[v1.ProjectServiceJoinListRequest]) (*connect.Response[v1.ProjectServiceJoinListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto_project.v1.ProjectService.JoinList is not implemented"))
}
