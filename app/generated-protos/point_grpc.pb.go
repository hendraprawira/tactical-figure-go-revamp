// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: point.proto

package generated_proto_point

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TodoService_CreatePoint_FullMethodName = "/point.v1.TodoService/CreatePoint"
	TodoService_GetPoint_FullMethodName    = "/point.v1.TodoService/GetPoint"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...grpc.CallOption) (*CreatePointResponse, error)
	GetPoint(ctx context.Context, in *GetPointRequest, opts ...grpc.CallOption) (*GetPointResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...grpc.CallOption) (*CreatePointResponse, error) {
	out := new(CreatePointResponse)
	err := c.cc.Invoke(ctx, TodoService_CreatePoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetPoint(ctx context.Context, in *GetPointRequest, opts ...grpc.CallOption) (*GetPointResponse, error) {
	out := new(GetPointResponse)
	err := c.cc.Invoke(ctx, TodoService_GetPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreatePoint(context.Context, *CreatePointRequest) (*CreatePointResponse, error)
	GetPoint(context.Context, *GetPointRequest) (*GetPointResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreatePoint(context.Context, *CreatePointRequest) (*CreatePointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoint not implemented")
}
func (UnimplementedTodoServiceServer) GetPoint(context.Context, *GetPointRequest) (*GetPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoint not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreatePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreatePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreatePoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreatePoint(ctx, req.(*CreatePointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetPoint(ctx, req.(*GetPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "point.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoint",
			Handler:    _TodoService_CreatePoint_Handler,
		},
		{
			MethodName: "GetPoint",
			Handler:    _TodoService_GetPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "point.proto",
}
