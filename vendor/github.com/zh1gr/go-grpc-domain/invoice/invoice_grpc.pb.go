// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: invoice.proto

package invoice

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

// InvoiceClient is the client API for Invoice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type invoiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceClient(cc grpc.ClientConnInterface) InvoiceClient {
	return &invoiceClient{cc}
}

func (c *invoiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/Invoice/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServer is the server API for Invoice service.
// All implementations must embed UnimplementedInvoiceServer
// for forward compatibility
type InvoiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedInvoiceServer()
}

// UnimplementedInvoiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceServer struct {
}

func (UnimplementedInvoiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInvoiceServer) mustEmbedUnimplementedInvoiceServer() {}

// UnsafeInvoiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServer will
// result in compilation errors.
type UnsafeInvoiceServer interface {
	mustEmbedUnimplementedInvoiceServer()
}

func RegisterInvoiceServer(s grpc.ServiceRegistrar, srv InvoiceServer) {
	s.RegisterService(&Invoice_ServiceDesc, srv)
}

func _Invoice_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoice/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoice_ServiceDesc is the grpc.ServiceDesc for Invoice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoice",
	HandlerType: (*InvoiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoice_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice.proto",
}