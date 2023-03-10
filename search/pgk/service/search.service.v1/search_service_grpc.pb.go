// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: search_service.proto

package search_service_v1

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	SearchNews(ctx context.Context, in *ArticleMessage, opts ...grpc.CallOption) (*ArticleResponse, error)
	NewsQuery(ctx context.Context, in *NewsMessage, opts ...grpc.CallOption) (*NewsResponse, error)
	StarQuery(ctx context.Context, in *StarMessage, opts ...grpc.CallOption) (*StarResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchNews(ctx context.Context, in *ArticleMessage, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/search.service.v1.SearchService/SearchNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) NewsQuery(ctx context.Context, in *NewsMessage, opts ...grpc.CallOption) (*NewsResponse, error) {
	out := new(NewsResponse)
	err := c.cc.Invoke(ctx, "/search.service.v1.SearchService/NewsQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) StarQuery(ctx context.Context, in *StarMessage, opts ...grpc.CallOption) (*StarResponse, error) {
	out := new(StarResponse)
	err := c.cc.Invoke(ctx, "/search.service.v1.SearchService/StarQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	SearchNews(context.Context, *ArticleMessage) (*ArticleResponse, error)
	NewsQuery(context.Context, *NewsMessage) (*NewsResponse, error)
	StarQuery(context.Context, *StarMessage) (*StarResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchNews(context.Context, *ArticleMessage) (*ArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNews not implemented")
}
func (UnimplementedSearchServiceServer) NewsQuery(context.Context, *NewsMessage) (*NewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsQuery not implemented")
}
func (UnimplementedSearchServiceServer) StarQuery(context.Context, *StarMessage) (*StarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarQuery not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.service.v1.SearchService/SearchNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchNews(ctx, req.(*ArticleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_NewsQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).NewsQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.service.v1.SearchService/NewsQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).NewsQuery(ctx, req.(*NewsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_StarQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StarMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).StarQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.service.v1.SearchService/StarQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).StarQuery(ctx, req.(*StarMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.service.v1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchNews",
			Handler:    _SearchService_SearchNews_Handler,
		},
		{
			MethodName: "NewsQuery",
			Handler:    _SearchService_NewsQuery_Handler,
		},
		{
			MethodName: "StarQuery",
			Handler:    _SearchService_StarQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_service.proto",
}
