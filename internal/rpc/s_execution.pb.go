// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s_execution.proto

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

type MRunStageRequest struct {
	StageId              int32                `protobuf:"varint,1,opt,name=stageId,proto3" json:"stageId,omitempty"`
	RunShuffle           bool                 `protobuf:"varint,2,opt,name=runShuffle,proto3" json:"runShuffle,omitempty"`
	PrepCollect          bool                 `protobuf:"varint,3,opt,name=prepCollect,proto3" json:"prepCollect,omitempty"`
	AssignedBucket       uint64               `protobuf:"varint,4,opt,name=assignedBucket,proto3" json:"assignedBucket,omitempty"`
	Buckets              []uint64             `protobuf:"varint,5,rep,packed,name=buckets,proto3" json:"buckets,omitempty"`
	Workers              []*MWorkerDescriptor `protobuf:"bytes,6,rep,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MRunStageRequest) Reset()         { *m = MRunStageRequest{} }
func (m *MRunStageRequest) String() string { return proto.CompactTextString(m) }
func (*MRunStageRequest) ProtoMessage()    {}
func (*MRunStageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1270cbd13ab8b72, []int{0}
}

func (m *MRunStageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MRunStageRequest.Unmarshal(m, b)
}
func (m *MRunStageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MRunStageRequest.Marshal(b, m, deterministic)
}
func (m *MRunStageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MRunStageRequest.Merge(m, src)
}
func (m *MRunStageRequest) XXX_Size() int {
	return xxx_messageInfo_MRunStageRequest.Size(m)
}
func (m *MRunStageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MRunStageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MRunStageRequest proto.InternalMessageInfo

func (m *MRunStageRequest) GetStageId() int32 {
	if m != nil {
		return m.StageId
	}
	return 0
}

func (m *MRunStageRequest) GetRunShuffle() bool {
	if m != nil {
		return m.RunShuffle
	}
	return false
}

func (m *MRunStageRequest) GetPrepCollect() bool {
	if m != nil {
		return m.PrepCollect
	}
	return false
}

func (m *MRunStageRequest) GetAssignedBucket() uint64 {
	if m != nil {
		return m.AssignedBucket
	}
	return 0
}

func (m *MRunStageRequest) GetBuckets() []uint64 {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *MRunStageRequest) GetWorkers() []*MWorkerDescriptor {
	if m != nil {
		return m.Workers
	}
	return nil
}

type MRunStageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MRunStageResponse) Reset()         { *m = MRunStageResponse{} }
func (m *MRunStageResponse) String() string { return proto.CompactTextString(m) }
func (*MRunStageResponse) ProtoMessage()    {}
func (*MRunStageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1270cbd13ab8b72, []int{1}
}

func (m *MRunStageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MRunStageResponse.Unmarshal(m, b)
}
func (m *MRunStageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MRunStageResponse.Marshal(b, m, deterministic)
}
func (m *MRunStageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MRunStageResponse.Merge(m, src)
}
func (m *MRunStageResponse) XXX_Size() int {
	return xxx_messageInfo_MRunStageResponse.Size(m)
}
func (m *MRunStageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MRunStageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MRunStageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MRunStageRequest)(nil), "rpc.MRunStageRequest")
	proto.RegisterType((*MRunStageResponse)(nil), "rpc.MRunStageResponse")
}

func init() {
	proto.RegisterFile("s_execution.proto", fileDescriptor_a1270cbd13ab8b72)
}

var fileDescriptor_a1270cbd13ab8b72 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xdf, 0x69, 0xfa, 0xc5, 0x2d, 0xbc, 0xb6, 0x23, 0x96, 0xa1, 0x88, 0x0c, 0x5d, 0x48,
	0x56, 0x41, 0xea, 0xd2, 0x5d, 0xd4, 0x85, 0x8b, 0x22, 0x4c, 0x16, 0x82, 0x9b, 0xd2, 0x4e, 0x6f,
	0x6b, 0x68, 0xc8, 0x8c, 0x73, 0x67, 0xd4, 0x5f, 0xec, 0xef, 0x90, 0x26, 0x04, 0x42, 0x76, 0xf7,
	0x3c, 0x1c, 0x0e, 0xdc, 0x07, 0x66, 0xb4, 0xc1, 0x1f, 0xd4, 0xc1, 0xe7, 0xa6, 0x4c, 0xac, 0x33,
	0xde, 0xf0, 0xc8, 0x59, 0xbd, 0xb8, 0xa0, 0x8d, 0x2e, 0x02, 0x79, 0x74, 0x35, 0x5d, 0xfe, 0x32,
	0x98, 0xae, 0x55, 0x28, 0x33, 0xbf, 0x3d, 0xa2, 0xc2, 0xcf, 0x80, 0xe4, 0xb9, 0x80, 0x11, 0x9d,
	0xf3, 0xcb, 0x5e, 0x30, 0xc9, 0xe2, 0x81, 0x6a, 0x22, 0xbf, 0x01, 0x70, 0xa1, 0xcc, 0x3e, 0xc2,
	0xe1, 0x50, 0xa0, 0xe8, 0x49, 0x16, 0x8f, 0x55, 0x8b, 0x70, 0x09, 0x13, 0xeb, 0xd0, 0x3e, 0x9a,
	0xa2, 0x40, 0xed, 0x45, 0x54, 0x15, 0xda, 0x88, 0xdf, 0xc2, 0xff, 0x2d, 0x51, 0x7e, 0x2c, 0x71,
	0x9f, 0x06, 0x7d, 0x42, 0x2f, 0xfa, 0x92, 0xc5, 0x7d, 0xd5, 0xa1, 0xfc, 0x1a, 0x46, 0xbb, 0xea,
	0x22, 0x31, 0x90, 0x51, 0xdc, 0x4f, 0x7b, 0x53, 0xa6, 0x1a, 0xc4, 0xef, 0x60, 0xf4, 0x6d, 0xdc,
	0x09, 0x1d, 0x89, 0xa1, 0x8c, 0xe2, 0xc9, 0x6a, 0x9e, 0x38, 0xab, 0x93, 0xf5, 0x5b, 0x05, 0x9f,
	0x90, 0xb4, 0xcb, 0xad, 0x37, 0x4e, 0x35, 0xb5, 0xe5, 0x25, 0xcc, 0x5a, 0x7f, 0x92, 0x35, 0x25,
	0xe1, 0xea, 0x15, 0xa6, 0xcf, 0x8d, 0xa6, 0x0c, 0xdd, 0x57, 0xae, 0x91, 0x3f, 0xc0, 0xb8, 0xe9,
	0xf1, 0xab, 0x7a, 0xb5, 0xe3, 0x67, 0x31, 0xef, 0xe2, 0x7a, 0x6e, 0xf9, 0x2f, 0x1d, 0xbc, 0x9f,
	0x35, 0xef, 0x86, 0x95, 0xdc, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xca, 0x46, 0xca,
	0x87, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExecutionServiceClient is the client API for ExecutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutionServiceClient interface {
	RunStage(ctx context.Context, in *MRunStageRequest, opts ...grpc.CallOption) (*MRunStageResponse, error)
}

type executionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionServiceClient(cc grpc.ClientConnInterface) ExecutionServiceClient {
	return &executionServiceClient{cc}
}

func (c *executionServiceClient) RunStage(ctx context.Context, in *MRunStageRequest, opts ...grpc.CallOption) (*MRunStageResponse, error) {
	out := new(MRunStageResponse)
	err := c.cc.Invoke(ctx, "/rpc.ExecutionService/RunStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionServiceServer is the server API for ExecutionService service.
type ExecutionServiceServer interface {
	RunStage(context.Context, *MRunStageRequest) (*MRunStageResponse, error)
}

// UnimplementedExecutionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExecutionServiceServer struct {
}

func (*UnimplementedExecutionServiceServer) RunStage(ctx context.Context, req *MRunStageRequest) (*MRunStageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunStage not implemented")
}

func RegisterExecutionServiceServer(s *grpc.Server, srv ExecutionServiceServer) {
	s.RegisterService(&_ExecutionService_serviceDesc, srv)
}

func _ExecutionService_RunStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MRunStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServiceServer).RunStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ExecutionService/RunStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServiceServer).RunStage(ctx, req.(*MRunStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecutionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ExecutionService",
	HandlerType: (*ExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunStage",
			Handler:    _ExecutionService_RunStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s_execution.proto",
}
