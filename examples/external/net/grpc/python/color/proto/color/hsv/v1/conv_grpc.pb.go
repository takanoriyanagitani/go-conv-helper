// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/color/hsv/v1/conv.proto

package v1

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
	ColorConvService_FromHsv_FullMethodName = "/color.hsv.v1.ColorConvService/FromHsv"
)

// ColorConvServiceClient is the client API for ColorConvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColorConvServiceClient interface {
	FromHsv(ctx context.Context, in *FromHsvRequest, opts ...grpc.CallOption) (*FromHsvResponse, error)
}

type colorConvServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewColorConvServiceClient(cc grpc.ClientConnInterface) ColorConvServiceClient {
	return &colorConvServiceClient{cc}
}

func (c *colorConvServiceClient) FromHsv(ctx context.Context, in *FromHsvRequest, opts ...grpc.CallOption) (*FromHsvResponse, error) {
	out := new(FromHsvResponse)
	err := c.cc.Invoke(ctx, ColorConvService_FromHsv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ColorConvServiceServer is the server API for ColorConvService service.
// All implementations must embed UnimplementedColorConvServiceServer
// for forward compatibility
type ColorConvServiceServer interface {
	FromHsv(context.Context, *FromHsvRequest) (*FromHsvResponse, error)
	mustEmbedUnimplementedColorConvServiceServer()
}

// UnimplementedColorConvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedColorConvServiceServer struct {
}

func (UnimplementedColorConvServiceServer) FromHsv(context.Context, *FromHsvRequest) (*FromHsvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FromHsv not implemented")
}
func (UnimplementedColorConvServiceServer) mustEmbedUnimplementedColorConvServiceServer() {}

// UnsafeColorConvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColorConvServiceServer will
// result in compilation errors.
type UnsafeColorConvServiceServer interface {
	mustEmbedUnimplementedColorConvServiceServer()
}

func RegisterColorConvServiceServer(s grpc.ServiceRegistrar, srv ColorConvServiceServer) {
	s.RegisterService(&ColorConvService_ServiceDesc, srv)
}

func _ColorConvService_FromHsv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FromHsvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorConvServiceServer).FromHsv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorConvService_FromHsv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorConvServiceServer).FromHsv(ctx, req.(*FromHsvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ColorConvService_ServiceDesc is the grpc.ServiceDesc for ColorConvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColorConvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "color.hsv.v1.ColorConvService",
	HandlerType: (*ColorConvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FromHsv",
			Handler:    _ColorConvService_FromHsv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/color/hsv/v1/conv.proto",
}

const (
	HsvEvtService_Converted_FullMethodName = "/color.hsv.v1.HsvEvtService/Converted"
)

// HsvEvtServiceClient is the client API for HsvEvtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HsvEvtServiceClient interface {
	Converted(ctx context.Context, in *HsvEvt_ConvertedRequest, opts ...grpc.CallOption) (*HsvEvt_ConvertedResponse, error)
}

type hsvEvtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHsvEvtServiceClient(cc grpc.ClientConnInterface) HsvEvtServiceClient {
	return &hsvEvtServiceClient{cc}
}

func (c *hsvEvtServiceClient) Converted(ctx context.Context, in *HsvEvt_ConvertedRequest, opts ...grpc.CallOption) (*HsvEvt_ConvertedResponse, error) {
	out := new(HsvEvt_ConvertedResponse)
	err := c.cc.Invoke(ctx, HsvEvtService_Converted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HsvEvtServiceServer is the server API for HsvEvtService service.
// All implementations must embed UnimplementedHsvEvtServiceServer
// for forward compatibility
type HsvEvtServiceServer interface {
	Converted(context.Context, *HsvEvt_ConvertedRequest) (*HsvEvt_ConvertedResponse, error)
	mustEmbedUnimplementedHsvEvtServiceServer()
}

// UnimplementedHsvEvtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHsvEvtServiceServer struct {
}

func (UnimplementedHsvEvtServiceServer) Converted(context.Context, *HsvEvt_ConvertedRequest) (*HsvEvt_ConvertedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Converted not implemented")
}
func (UnimplementedHsvEvtServiceServer) mustEmbedUnimplementedHsvEvtServiceServer() {}

// UnsafeHsvEvtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HsvEvtServiceServer will
// result in compilation errors.
type UnsafeHsvEvtServiceServer interface {
	mustEmbedUnimplementedHsvEvtServiceServer()
}

func RegisterHsvEvtServiceServer(s grpc.ServiceRegistrar, srv HsvEvtServiceServer) {
	s.RegisterService(&HsvEvtService_ServiceDesc, srv)
}

func _HsvEvtService_Converted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HsvEvt_ConvertedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HsvEvtServiceServer).Converted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HsvEvtService_Converted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HsvEvtServiceServer).Converted(ctx, req.(*HsvEvt_ConvertedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HsvEvtService_ServiceDesc is the grpc.ServiceDesc for HsvEvtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HsvEvtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "color.hsv.v1.HsvEvtService",
	HandlerType: (*HsvEvtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Converted",
			Handler:    _HsvEvtService_Converted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/color/hsv/v1/conv.proto",
}

const (
	HsvCmdService_Convert_FullMethodName = "/color.hsv.v1.HsvCmdService/Convert"
)

// HsvCmdServiceClient is the client API for HsvCmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HsvCmdServiceClient interface {
	Convert(ctx context.Context, in *HsvCmd_ConvertRequest, opts ...grpc.CallOption) (*HsvCmd_ConvertResponse, error)
}

type hsvCmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHsvCmdServiceClient(cc grpc.ClientConnInterface) HsvCmdServiceClient {
	return &hsvCmdServiceClient{cc}
}

func (c *hsvCmdServiceClient) Convert(ctx context.Context, in *HsvCmd_ConvertRequest, opts ...grpc.CallOption) (*HsvCmd_ConvertResponse, error) {
	out := new(HsvCmd_ConvertResponse)
	err := c.cc.Invoke(ctx, HsvCmdService_Convert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HsvCmdServiceServer is the server API for HsvCmdService service.
// All implementations must embed UnimplementedHsvCmdServiceServer
// for forward compatibility
type HsvCmdServiceServer interface {
	Convert(context.Context, *HsvCmd_ConvertRequest) (*HsvCmd_ConvertResponse, error)
	mustEmbedUnimplementedHsvCmdServiceServer()
}

// UnimplementedHsvCmdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHsvCmdServiceServer struct {
}

func (UnimplementedHsvCmdServiceServer) Convert(context.Context, *HsvCmd_ConvertRequest) (*HsvCmd_ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedHsvCmdServiceServer) mustEmbedUnimplementedHsvCmdServiceServer() {}

// UnsafeHsvCmdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HsvCmdServiceServer will
// result in compilation errors.
type UnsafeHsvCmdServiceServer interface {
	mustEmbedUnimplementedHsvCmdServiceServer()
}

func RegisterHsvCmdServiceServer(s grpc.ServiceRegistrar, srv HsvCmdServiceServer) {
	s.RegisterService(&HsvCmdService_ServiceDesc, srv)
}

func _HsvCmdService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HsvCmd_ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HsvCmdServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HsvCmdService_Convert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HsvCmdServiceServer).Convert(ctx, req.(*HsvCmd_ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HsvCmdService_ServiceDesc is the grpc.ServiceDesc for HsvCmdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HsvCmdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "color.hsv.v1.HsvCmdService",
	HandlerType: (*HsvCmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _HsvCmdService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/color/hsv/v1/conv.proto",
}

const (
	HsvCmdRecvService_GetCmd_FullMethodName = "/color.hsv.v1.HsvCmdRecvService/GetCmd"
)

// HsvCmdRecvServiceClient is the client API for HsvCmdRecvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HsvCmdRecvServiceClient interface {
	GetCmd(ctx context.Context, in *HsvOrder_GetCmdRequest, opts ...grpc.CallOption) (*HsvOrder_GetCmdResponse, error)
}

type hsvCmdRecvServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHsvCmdRecvServiceClient(cc grpc.ClientConnInterface) HsvCmdRecvServiceClient {
	return &hsvCmdRecvServiceClient{cc}
}

func (c *hsvCmdRecvServiceClient) GetCmd(ctx context.Context, in *HsvOrder_GetCmdRequest, opts ...grpc.CallOption) (*HsvOrder_GetCmdResponse, error) {
	out := new(HsvOrder_GetCmdResponse)
	err := c.cc.Invoke(ctx, HsvCmdRecvService_GetCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HsvCmdRecvServiceServer is the server API for HsvCmdRecvService service.
// All implementations must embed UnimplementedHsvCmdRecvServiceServer
// for forward compatibility
type HsvCmdRecvServiceServer interface {
	GetCmd(context.Context, *HsvOrder_GetCmdRequest) (*HsvOrder_GetCmdResponse, error)
	mustEmbedUnimplementedHsvCmdRecvServiceServer()
}

// UnimplementedHsvCmdRecvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHsvCmdRecvServiceServer struct {
}

func (UnimplementedHsvCmdRecvServiceServer) GetCmd(context.Context, *HsvOrder_GetCmdRequest) (*HsvOrder_GetCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCmd not implemented")
}
func (UnimplementedHsvCmdRecvServiceServer) mustEmbedUnimplementedHsvCmdRecvServiceServer() {}

// UnsafeHsvCmdRecvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HsvCmdRecvServiceServer will
// result in compilation errors.
type UnsafeHsvCmdRecvServiceServer interface {
	mustEmbedUnimplementedHsvCmdRecvServiceServer()
}

func RegisterHsvCmdRecvServiceServer(s grpc.ServiceRegistrar, srv HsvCmdRecvServiceServer) {
	s.RegisterService(&HsvCmdRecvService_ServiceDesc, srv)
}

func _HsvCmdRecvService_GetCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HsvOrder_GetCmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HsvCmdRecvServiceServer).GetCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HsvCmdRecvService_GetCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HsvCmdRecvServiceServer).GetCmd(ctx, req.(*HsvOrder_GetCmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HsvCmdRecvService_ServiceDesc is the grpc.ServiceDesc for HsvCmdRecvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HsvCmdRecvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "color.hsv.v1.HsvCmdRecvService",
	HandlerType: (*HsvCmdRecvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCmd",
			Handler:    _HsvCmdRecvService_GetCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/color/hsv/v1/conv.proto",
}
