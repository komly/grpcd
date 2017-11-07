// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package fixtures is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	TestMessage
	TestMessageWithRepeated
	InnerMessage
	TestMessageWithInner
	TestRepeatedInner
*/
package fixtures

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestMessage struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestMessage) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *TestMessage) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type TestMessageWithRepeated struct {
	A []int32 `protobuf:"varint,1,rep,packed,name=a" json:"a,omitempty"`
}

func (m *TestMessageWithRepeated) Reset()                    { *m = TestMessageWithRepeated{} }
func (m *TestMessageWithRepeated) String() string            { return proto.CompactTextString(m) }
func (*TestMessageWithRepeated) ProtoMessage()               {}
func (*TestMessageWithRepeated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestMessageWithRepeated) GetA() []int32 {
	if m != nil {
		return m.A
	}
	return nil
}

type InnerMessage struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
}

func (m *InnerMessage) Reset()                    { *m = InnerMessage{} }
func (m *InnerMessage) String() string            { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()               {}
func (*InnerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InnerMessage) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

type TestMessageWithInner struct {
	Inner *InnerMessage `protobuf:"bytes,1,opt,name=inner" json:"inner,omitempty"`
}

func (m *TestMessageWithInner) Reset()                    { *m = TestMessageWithInner{} }
func (m *TestMessageWithInner) String() string            { return proto.CompactTextString(m) }
func (*TestMessageWithInner) ProtoMessage()               {}
func (*TestMessageWithInner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TestMessageWithInner) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

type TestRepeatedInner struct {
	Inner []*InnerMessage `protobuf:"bytes,1,rep,name=inner" json:"inner,omitempty"`
}

func (m *TestRepeatedInner) Reset()                    { *m = TestRepeatedInner{} }
func (m *TestRepeatedInner) String() string            { return proto.CompactTextString(m) }
func (*TestRepeatedInner) ProtoMessage()               {}
func (*TestRepeatedInner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TestRepeatedInner) GetInner() []*InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "fixtures.TestMessage")
	proto.RegisterType((*TestMessageWithRepeated)(nil), "fixtures.TestMessageWithRepeated")
	proto.RegisterType((*InnerMessage)(nil), "fixtures.InnerMessage")
	proto.RegisterType((*TestMessageWithInner)(nil), "fixtures.TestMessageWithInner")
	proto.RegisterType((*TestRepeatedInner)(nil), "fixtures.TestRepeatedInner")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	TestMethod(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestMethod(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error) {
	out := new(TestMessage)
	err := grpc.Invoke(ctx, "/fixtures.TestService/TestMethod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	TestMethod(context.Context, *TestMessage) (*TestMessage, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fixtures.TestService/TestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestMethod(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fixtures.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _TestService_TestMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xcb, 0xac, 0x28, 0x29,
	0x2d, 0x4a, 0x2d, 0x56, 0xd2, 0xe4, 0xe2, 0x0e, 0x49, 0x2d, 0x2e, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0x15, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c,
	0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x75, 0x2e, 0x71, 0x24, 0xa5, 0xe1, 0x99,
	0x25, 0x19, 0x41, 0xa9, 0x05, 0xa9, 0x89, 0x25, 0xa9, 0x29, 0x30, 0x6d, 0xcc, 0x60, 0x6d, 0x4a,
	0x32, 0x5c, 0x3c, 0x9e, 0x79, 0x79, 0xa9, 0x45, 0x58, 0x0d, 0x55, 0x72, 0xe1, 0x12, 0x41, 0x33,
	0x06, 0xac, 0x58, 0x48, 0x87, 0x8b, 0x35, 0x13, 0xc4, 0x00, 0xab, 0xe4, 0x36, 0x12, 0xd3, 0x83,
	0xb9, 0x51, 0x0f, 0xd9, 0xb0, 0x20, 0x88, 0x22, 0x25, 0x47, 0x2e, 0x41, 0x90, 0x29, 0x30, 0x17,
	0x60, 0x18, 0xc1, 0x4c, 0xd0, 0x08, 0x23, 0x4f, 0x88, 0xd7, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0x85, 0xac, 0xb8, 0xb8, 0x20, 0xee, 0x2a, 0xc9, 0xc8, 0x4f, 0x11, 0x12, 0x45, 0xe8, 0x45,
	0x72, 0xad, 0x14, 0x76, 0xe1, 0x24, 0x36, 0x70, 0xb0, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x08, 0xf8, 0xd4, 0x3c, 0x67, 0x01, 0x00, 0x00,
}
