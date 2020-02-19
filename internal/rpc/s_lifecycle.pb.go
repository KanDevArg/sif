// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s_lifecycle.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MStopResponse struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MStopResponse) Reset()         { *m = MStopResponse{} }
func (m *MStopResponse) String() string { return proto.CompactTextString(m) }
func (*MStopResponse) ProtoMessage()    {}
func (*MStopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed52758a02481eb, []int{0}
}

func (m *MStopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MStopResponse.Unmarshal(m, b)
}
func (m *MStopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MStopResponse.Marshal(b, m, deterministic)
}
func (m *MStopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MStopResponse.Merge(m, src)
}
func (m *MStopResponse) XXX_Size() int {
	return xxx_messageInfo_MStopResponse.Size(m)
}
func (m *MStopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MStopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MStopResponse proto.InternalMessageInfo

func (m *MStopResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*MStopResponse)(nil), "rpc.MStopResponse")
}

func init() { proto.RegisterFile("s_lifecycle.proto", fileDescriptor_5ed52758a02481eb) }

var fileDescriptor_5ed52758a02481eb = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x8e, 0xcf, 0xc9,
	0x4c, 0x4b, 0x4d, 0xae, 0x4c, 0xce, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e,
	0x2a, 0x48, 0x96, 0xe2, 0x2f, 0x8e, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x82, 0x88, 0x2a,
	0x29, 0x73, 0xf1, 0xfa, 0x06, 0x97, 0xe4, 0x17, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x09, 0x71, 0xb1, 0x94, 0x64, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81,
	0xd9, 0x46, 0x6d, 0x8c, 0x5c, 0x02, 0x3e, 0x30, 0xe3, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x6c, 0xb8, 0x78, 0xdc, 0x8b, 0x12, 0x93, 0x53, 0xd3, 0x4a, 0x73, 0x40, 0x06, 0x08, 0x89,
	0xe9, 0x15, 0x15, 0x24, 0xeb, 0xf9, 0x86, 0xe7, 0x17, 0x65, 0xa7, 0x16, 0xb9, 0xa4, 0x16, 0x27,
	0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0x49, 0x09, 0x41, 0xc4, 0x91, 0x2d, 0x51, 0x62, 0x10, 0x32,
	0xe1, 0x62, 0x21, 0x5d, 0x57, 0x12, 0x1b, 0xd8, 0xd1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x83, 0xbf, 0x48, 0x0e, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LifecycleServiceClient is the client API for LifecycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LifecycleServiceClient interface {
	GracefulStop(ctx context.Context, in *MWorkerDescriptor, opts ...grpc.CallOption) (*MStopResponse, error)
	Stop(ctx context.Context, in *MWorkerDescriptor, opts ...grpc.CallOption) (*MStopResponse, error)
}

type lifecycleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLifecycleServiceClient(cc grpc.ClientConnInterface) LifecycleServiceClient {
	return &lifecycleServiceClient{cc}
}

func (c *lifecycleServiceClient) GracefulStop(ctx context.Context, in *MWorkerDescriptor, opts ...grpc.CallOption) (*MStopResponse, error) {
	out := new(MStopResponse)
	err := c.cc.Invoke(ctx, "/rpc.LifecycleService/GracefulStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecycleServiceClient) Stop(ctx context.Context, in *MWorkerDescriptor, opts ...grpc.CallOption) (*MStopResponse, error) {
	out := new(MStopResponse)
	err := c.cc.Invoke(ctx, "/rpc.LifecycleService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifecycleServiceServer is the server API for LifecycleService service.
type LifecycleServiceServer interface {
	GracefulStop(context.Context, *MWorkerDescriptor) (*MStopResponse, error)
	Stop(context.Context, *MWorkerDescriptor) (*MStopResponse, error)
}

// UnimplementedLifecycleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLifecycleServiceServer struct {
}

func (*UnimplementedLifecycleServiceServer) GracefulStop(ctx context.Context, req *MWorkerDescriptor) (*MStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GracefulStop not implemented")
}
func (*UnimplementedLifecycleServiceServer) Stop(ctx context.Context, req *MWorkerDescriptor) (*MStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterLifecycleServiceServer(s *grpc.Server, srv LifecycleServiceServer) {
	s.RegisterService(&_LifecycleService_serviceDesc, srv)
}

func _LifecycleService_GracefulStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MWorkerDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecycleServiceServer).GracefulStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LifecycleService/GracefulStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecycleServiceServer).GracefulStop(ctx, req.(*MWorkerDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecycleService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MWorkerDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecycleServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LifecycleService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecycleServiceServer).Stop(ctx, req.(*MWorkerDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

var _LifecycleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LifecycleService",
	HandlerType: (*LifecycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GracefulStop",
			Handler:    _LifecycleService_GracefulStop_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LifecycleService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s_lifecycle.proto",
}