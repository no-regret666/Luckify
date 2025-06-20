// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0
// source: app/upload/cmd/rpc/pb/upload.proto

package pb

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
	Upload_Upload_FullMethodName = "/pb.upload/upload"
)

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadClient interface {
	Upload(ctx context.Context, in *FileUploadReq, opts ...grpc.CallOption) (*FileUploadResp, error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) Upload(ctx context.Context, in *FileUploadReq, opts ...grpc.CallOption) (*FileUploadResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileUploadResp)
	err := c.cc.Invoke(ctx, Upload_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServer is the server API for Upload service.
// All implementations must embed UnimplementedUploadServer
// for forward compatibility.
type UploadServer interface {
	Upload(context.Context, *FileUploadReq) (*FileUploadResp, error)
	mustEmbedUnimplementedUploadServer()
}

// UnimplementedUploadServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadServer struct{}

func (UnimplementedUploadServer) Upload(context.Context, *FileUploadReq) (*FileUploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadServer) mustEmbedUnimplementedUploadServer() {}
func (UnimplementedUploadServer) testEmbeddedByValue()                {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	// If the following call pancis, it indicates UnimplementedUploadServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upload_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).Upload(ctx, req.(*FileUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.upload",
	HandlerType: (*UploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upload",
			Handler:    _Upload_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/upload/cmd/rpc/pb/upload.proto",
}
