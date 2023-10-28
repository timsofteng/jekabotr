// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: youtube.proto

package __protot

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
	Youtube_GetRandomVideo_FullMethodName = "/Youtube/GetRandomVideo"
)

// YoutubeClient is the client API for Youtube service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YoutubeClient interface {
	GetRandomVideo(ctx context.Context, in *RandomVideoRequest, opts ...grpc.CallOption) (*RandomVideoResponse, error)
}

type youtubeClient struct {
	cc grpc.ClientConnInterface
}

func NewYoutubeClient(cc grpc.ClientConnInterface) YoutubeClient {
	return &youtubeClient{cc}
}

func (c *youtubeClient) GetRandomVideo(ctx context.Context, in *RandomVideoRequest, opts ...grpc.CallOption) (*RandomVideoResponse, error) {
	out := new(RandomVideoResponse)
	err := c.cc.Invoke(ctx, Youtube_GetRandomVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeServer is the server API for Youtube service.
// All implementations must embed UnimplementedYoutubeServer
// for forward compatibility
type YoutubeServer interface {
	GetRandomVideo(context.Context, *RandomVideoRequest) (*RandomVideoResponse, error)
	mustEmbedUnimplementedYoutubeServer()
}

// UnimplementedYoutubeServer must be embedded to have forward compatible implementations.
type UnimplementedYoutubeServer struct {
}

func (UnimplementedYoutubeServer) GetRandomVideo(context.Context, *RandomVideoRequest) (*RandomVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomVideo not implemented")
}
func (UnimplementedYoutubeServer) mustEmbedUnimplementedYoutubeServer() {}

// UnsafeYoutubeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YoutubeServer will
// result in compilation errors.
type UnsafeYoutubeServer interface {
	mustEmbedUnimplementedYoutubeServer()
}

func RegisterYoutubeServer(s grpc.ServiceRegistrar, srv YoutubeServer) {
	s.RegisterService(&Youtube_ServiceDesc, srv)
}

func _Youtube_GetRandomVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServer).GetRandomVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Youtube_GetRandomVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServer).GetRandomVideo(ctx, req.(*RandomVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Youtube_ServiceDesc is the grpc.ServiceDesc for Youtube service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Youtube_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Youtube",
	HandlerType: (*YoutubeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandomVideo",
			Handler:    _Youtube_GetRandomVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "youtube.proto",
}
