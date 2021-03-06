// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s_log.proto

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

type MLogMsg struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MLogMsg) Reset()         { *m = MLogMsg{} }
func (m *MLogMsg) String() string { return proto.CompactTextString(m) }
func (*MLogMsg) ProtoMessage()    {}
func (*MLogMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a276b15ac053355, []int{0}
}

func (m *MLogMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MLogMsg.Unmarshal(m, b)
}
func (m *MLogMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MLogMsg.Marshal(b, m, deterministic)
}
func (m *MLogMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MLogMsg.Merge(m, src)
}
func (m *MLogMsg) XXX_Size() int {
	return xxx_messageInfo_MLogMsg.Size(m)
}
func (m *MLogMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MLogMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MLogMsg proto.InternalMessageInfo

func (m *MLogMsg) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *MLogMsg) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MLogMsg) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MLogMsgAck struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MLogMsgAck) Reset()         { *m = MLogMsgAck{} }
func (m *MLogMsgAck) String() string { return proto.CompactTextString(m) }
func (*MLogMsgAck) ProtoMessage()    {}
func (*MLogMsgAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a276b15ac053355, []int{1}
}

func (m *MLogMsgAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MLogMsgAck.Unmarshal(m, b)
}
func (m *MLogMsgAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MLogMsgAck.Marshal(b, m, deterministic)
}
func (m *MLogMsgAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MLogMsgAck.Merge(m, src)
}
func (m *MLogMsgAck) XXX_Size() int {
	return xxx_messageInfo_MLogMsgAck.Size(m)
}
func (m *MLogMsgAck) XXX_DiscardUnknown() {
	xxx_messageInfo_MLogMsgAck.DiscardUnknown(m)
}

var xxx_messageInfo_MLogMsgAck proto.InternalMessageInfo

func (m *MLogMsgAck) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MLogMsgAck) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*MLogMsg)(nil), "rpc.MLogMsg")
	proto.RegisterType((*MLogMsgAck)(nil), "rpc.MLogMsgAck")
}

func init() {
	proto.RegisterFile("s_log.proto", fileDescriptor_0a276b15ac053355)
}

var fileDescriptor_0a276b15ac053355 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0xc6, 0x71, 0x63, 0x4c, 0x8b, 0xa7, 0x20, 0x1c, 0x22, 0xc1, 0xa9, 0x74, 0xca, 0xd4, 0x41,
	0xa1, 0x7b, 0x9d, 0xdb, 0xc1, 0xb8, 0xb9, 0x88, 0x86, 0x23, 0x94, 0xb6, 0x5e, 0x49, 0xda, 0x7e,
	0x7e, 0x69, 0x5a, 0xe1, 0xdd, 0xf2, 0x23, 0xf0, 0x3f, 0x1e, 0xb8, 0x8b, 0x5f, 0x23, 0xfb, 0x6a,
	0x0e, 0xbc, 0x30, 0xca, 0x30, 0xbb, 0xf2, 0x1d, 0xf2, 0xae, 0x65, 0xdf, 0x45, 0x8f, 0x8f, 0xa0,
	0x46, 0xda, 0x68, 0xd4, 0xa2, 0x10, 0x46, 0xd9, 0x03, 0xf8, 0x04, 0x59, 0xe4, 0x35, 0x38, 0xd2,
	0xd7, 0x85, 0x30, 0xb7, 0xf6, 0x14, 0x6a, 0xc8, 0x27, 0x8a, 0xf1, 0xdb, 0x93, 0x96, 0xe9, 0xe3,
	0x9f, 0x65, 0x0d, 0x70, 0x26, 0x1b, 0x37, 0x20, 0xc2, 0xcd, 0xd2, 0x4f, 0x94, 0xa2, 0xd2, 0xa6,
	0xf7, 0x7e, 0xc9, 0xf1, 0xfa, 0xbb, 0xa4, 0xa4, 0xb2, 0x07, 0x5e, 0x6a, 0x80, 0x96, 0xfd, 0x07,
	0x85, 0xad, 0x77, 0x84, 0x06, 0x64, 0xcb, 0x1e, 0xef, 0xab, 0x30, 0xbb, 0xea, 0xec, 0x3d, 0x3f,
	0x5c, 0xaa, 0x71, 0x43, 0x79, 0x65, 0xc4, 0x9b, 0xfa, 0xdc, 0x97, 0xfc, 0x64, 0x69, 0xd5, 0xeb,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x20, 0x60, 0x2f, 0xe4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	Log(ctx context.Context, opts ...grpc.CallOption) (LogService_LogClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Log(ctx context.Context, opts ...grpc.CallOption) (LogService_LogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[0], "/rpc.LogService/Log", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceLogClient{stream}
	return x, nil
}

type LogService_LogClient interface {
	Send(*MLogMsg) error
	CloseAndRecv() (*MLogMsgAck, error)
	grpc.ClientStream
}

type logServiceLogClient struct {
	grpc.ClientStream
}

func (x *logServiceLogClient) Send(m *MLogMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logServiceLogClient) CloseAndRecv() (*MLogMsgAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MLogMsgAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	Log(LogService_LogServer) error
}

// UnimplementedLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (*UnimplementedLogServiceServer) Log(srv LogService_LogServer) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).Log(&logServiceLogServer{stream})
}

type LogService_LogServer interface {
	SendAndClose(*MLogMsgAck) error
	Recv() (*MLogMsg, error)
	grpc.ServerStream
}

type logServiceLogServer struct {
	grpc.ServerStream
}

func (x *logServiceLogServer) SendAndClose(m *MLogMsgAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logServiceLogServer) Recv() (*MLogMsg, error) {
	m := new(MLogMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Log",
			Handler:       _LogService_Log_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "s_log.proto",
}
