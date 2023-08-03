// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/raw/v1/gconv.proto

package raw_v1

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
	ConvertService_Convert_FullMethodName = "/raw.v1.ConvertService/Convert"
)

// ConvertServiceClient is the client API for ConvertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConvertServiceClient interface {
	Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
}

type convertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConvertServiceClient(cc grpc.ClientConnInterface) ConvertServiceClient {
	return &convertServiceClient{cc}
}

func (c *convertServiceClient) Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, ConvertService_Convert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConvertServiceServer is the server API for ConvertService service.
// All implementations must embed UnimplementedConvertServiceServer
// for forward compatibility
type ConvertServiceServer interface {
	Convert(context.Context, *ConvertRequest) (*ConvertResponse, error)
	mustEmbedUnimplementedConvertServiceServer()
}

// UnimplementedConvertServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConvertServiceServer struct {
}

func (UnimplementedConvertServiceServer) Convert(context.Context, *ConvertRequest) (*ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedConvertServiceServer) mustEmbedUnimplementedConvertServiceServer() {}

// UnsafeConvertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConvertServiceServer will
// result in compilation errors.
type UnsafeConvertServiceServer interface {
	mustEmbedUnimplementedConvertServiceServer()
}

func RegisterConvertServiceServer(s grpc.ServiceRegistrar, srv ConvertServiceServer) {
	s.RegisterService(&ConvertService_ServiceDesc, srv)
}

func _ConvertService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConvertServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConvertService_Convert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConvertServiceServer).Convert(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConvertService_ServiceDesc is the grpc.ServiceDesc for ConvertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConvertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raw.v1.ConvertService",
	HandlerType: (*ConvertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _ConvertService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/raw/v1/gconv.proto",
}
