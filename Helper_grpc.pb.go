// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: Helper.proto

package Helper

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

// Report_InscribedClient is the client API for Report_Inscribed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Report_InscribedClient interface {
	Inscribe(ctx context.Context, in *Betas, opts ...grpc.CallOption) (*To_Inscribe, error)
}

type report_InscribedClient struct {
	cc grpc.ClientConnInterface
}

func NewReport_InscribedClient(cc grpc.ClientConnInterface) Report_InscribedClient {
	return &report_InscribedClient{cc}
}

func (c *report_InscribedClient) Inscribe(ctx context.Context, in *Betas, opts ...grpc.CallOption) (*To_Inscribe, error) {
	out := new(To_Inscribe)
	err := c.cc.Invoke(ctx, "/Server_helper.Report_Inscribed/Inscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Report_InscribedServer is the server API for Report_Inscribed service.
// All implementations must embed UnimplementedReport_InscribedServer
// for forward compatibility
type Report_InscribedServer interface {
	Inscribe(context.Context, *Betas) (*To_Inscribe, error)
	mustEmbedUnimplementedReport_InscribedServer()
}

// UnimplementedReport_InscribedServer must be embedded to have forward compatible implementations.
type UnimplementedReport_InscribedServer struct {
}

func (UnimplementedReport_InscribedServer) Inscribe(context.Context, *Betas) (*To_Inscribe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inscribe not implemented")
}
func (UnimplementedReport_InscribedServer) mustEmbedUnimplementedReport_InscribedServer() {}

// UnsafeReport_InscribedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Report_InscribedServer will
// result in compilation errors.
type UnsafeReport_InscribedServer interface {
	mustEmbedUnimplementedReport_InscribedServer()
}

func RegisterReport_InscribedServer(s grpc.ServiceRegistrar, srv Report_InscribedServer) {
	s.RegisterService(&Report_Inscribed_ServiceDesc, srv)
}

func _Report_Inscribed_Inscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Betas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Report_InscribedServer).Inscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server_helper.Report_Inscribed/Inscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Report_InscribedServer).Inscribe(ctx, req.(*Betas))
	}
	return interceptor(ctx, in, info, handler)
}

// Report_Inscribed_ServiceDesc is the grpc.ServiceDesc for Report_Inscribed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Report_Inscribed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server_helper.Report_Inscribed",
	HandlerType: (*Report_InscribedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inscribe",
			Handler:    _Report_Inscribed_Inscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Helper.proto",
}
