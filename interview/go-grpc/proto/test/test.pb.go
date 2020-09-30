// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package server is a generated protocol buffer go-self-pakage.

It is generated from these files:
	server.proto

It has these top-level messages:
	Request
	Response
	Test
*/
package test

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
// is compatible with the proto go-self-pakage it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto go-self-pakage needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto go-self-pakage

// Status 枚举状态
type Test_Status int32

const (
	Test_OK   Test_Status = 0
	Test_FAIL Test_Status = 1
)

var Test_Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}
var Test_Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x Test_Status) String() string {
	return proto.EnumName(Test_Status_name, int32(x))
}
func (Test_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Request 请求结构
type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response 响应结构
type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Test 测试
type Test struct {
	Age    int32             `protobuf:"varint,1,opt,name=age" json:"age,omitempty"`
	Count  int64             `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Money  float64           `protobuf:"fixed64,3,opt,name=money" json:"money,omitempty"`
	Score  float32           `protobuf:"fixed32,4,opt,name=score" json:"score,omitempty"`
	Name   string            `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Fat    bool              `protobuf:"varint,6,opt,name=fat" json:"fat,omitempty"`
	Char   []byte            `protobuf:"bytes,7,opt,name=char,proto3" json:"char,omitempty"`
	Status Test_Status       `protobuf:"varint,8,opt,name=status,enum=server.Test_Status" json:"status,omitempty"`
	Child  *Test_Child       `protobuf:"bytes,9,opt,name=child" json:"child,omitempty"`
	Dict   map[string]string `protobuf:"bytes,10,rep,name=dict" json:"dict,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Test) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Test) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Test) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Test) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Test) GetFat() bool {
	if m != nil {
		return m.Fat
	}
	return false
}

func (m *Test) GetChar() []byte {
	if m != nil {
		return m.Char
	}
	return nil
}

func (m *Test) GetStatus() Test_Status {
	if m != nil {
		return m.Status
	}
	return Test_OK
}

func (m *Test) GetChild() *Test_Child {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *Test) GetDict() map[string]string {
	if m != nil {
		return m.Dict
	}
	return nil
}

// Child 子结构
type Test_Child struct {
	Sex string `protobuf:"bytes,1,opt,name=sex" json:"sex,omitempty"`
}

func (m *Test_Child) Reset()                    { *m = Test_Child{} }
func (m *Test_Child) String() string            { return proto.CompactTextString(m) }
func (*Test_Child) ProtoMessage()               {}
func (*Test_Child) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Test_Child) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "server.Request")
	proto.RegisterType((*Response)(nil), "server.Response")
	proto.RegisterType((*Test)(nil), "server.Test")
	proto.RegisterType((*Test_Child)(nil), "server.Test.Child")
	proto.RegisterEnum("server.Test_Status", Test_Status_name, Test_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc go-self-pakage it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	// Test 测试方法
	Test(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Test(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/server.TestService/Test", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// Test 测试方法
	Test(context.Context, *Request) (*Response, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.TestService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Test(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _TestService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x8d, 0xfc, 0x15, 0x7b, 0xb2, 0x1b, 0xbc, 0x22, 0x07, 0xad, 0x61, 0x41, 0x98, 0x65, 0x57,
	0xbd, 0xe4, 0xe0, 0x42, 0x5b, 0x7a, 0xeb, 0x27, 0x94, 0x16, 0x0a, 0x4a, 0x4f, 0xbd, 0xb9, 0x8e,
	0xda, 0x98, 0x24, 0x76, 0x6a, 0xc9, 0xa1, 0x39, 0xf6, 0x9f, 0x17, 0x49, 0x76, 0x9a, 0xdb, 0x7b,
	0xf3, 0x9e, 0xa4, 0xa7, 0x99, 0x01, 0x50, 0x42, 0xaa, 0xe9, 0xa6, 0xa9, 0x55, 0x8d, 0x3d, 0x8d,
	0xd3, 0x3f, 0x30, 0xe4, 0xe2, 0xbd, 0x15, 0x52, 0x61, 0x0c, 0x5e, 0x95, 0xaf, 0x05, 0x41, 0x14,
	0xb1, 0x88, 0x1b, 0x9c, 0xfe, 0x85, 0x90, 0x0b, 0xb9, 0xa9, 0x2b, 0x29, 0x30, 0x81, 0xe1, 0x5a,
	0x48, 0x99, 0xbf, 0xf5, 0x96, 0x9e, 0xa6, 0x9f, 0x2e, 0x78, 0x4f, 0xfa, 0x8a, 0x18, 0xdc, 0x5e,
	0xf6, 0xb9, 0x86, 0x78, 0x02, 0x7e, 0x51, 0xb7, 0x95, 0x22, 0x0e, 0x45, 0xcc, 0xe5, 0x96, 0xe8,
	0xea, 0xba, 0xae, 0xc4, 0x8e, 0xb8, 0x14, 0x31, 0xc4, 0x2d, 0xd1, 0x55, 0x59, 0xd4, 0x8d, 0x20,
	0x1e, 0x45, 0xcc, 0xe1, 0x96, 0xec, 0x63, 0xf9, 0xdf, 0xb1, 0xf4, 0x3b, 0xaf, 0xb9, 0x22, 0x01,
	0x45, 0x2c, 0xe4, 0x1a, 0x6a, 0x57, 0xb1, 0xc8, 0x1b, 0x32, 0xa4, 0x88, 0xfd, 0xe0, 0x06, 0xe3,
	0x23, 0x08, 0xa4, 0xca, 0x55, 0x2b, 0x49, 0x48, 0x11, 0x1b, 0x67, 0xbf, 0xa6, 0xe6, 0xfb, 0x3a,
	0xe9, 0x74, 0x66, 0x04, 0xde, 0x19, 0xf0, 0x3f, 0xf0, 0x8b, 0x45, 0xb9, 0x9a, 0x93, 0x88, 0x22,
	0x36, 0xca, 0xe2, 0x03, 0xe7, 0x95, 0xae, 0x73, 0x2b, 0x63, 0x06, 0xde, 0xbc, 0x2c, 0x14, 0x01,
	0xea, 0xb2, 0x51, 0x36, 0x39, 0xb0, 0x5d, 0x97, 0x85, 0xba, 0xa9, 0x54, 0xb3, 0xe3, 0xc6, 0x91,
	0xfc, 0x06, 0xdf, 0x9c, 0xd4, 0x59, 0xa5, 0xf8, 0xe8, 0x5a, 0xa6, 0x61, 0x72, 0x0a, 0xd1, 0xde,
	0xad, 0xe5, 0xa5, 0xd8, 0xf5, 0xf2, 0xd2, 0xb6, 0x61, 0x9b, 0xaf, 0x5a, 0x61, 0x5a, 0x16, 0x71,
	0x4b, 0xce, 0x9d, 0x33, 0x94, 0x26, 0x10, 0xd8, 0xdc, 0x38, 0x00, 0xe7, 0xf1, 0x3e, 0x1e, 0xe0,
	0x10, 0xbc, 0xdb, 0x8b, 0xbb, 0x87, 0x18, 0x65, 0x27, 0x30, 0xd2, 0x39, 0x66, 0xa2, 0xd9, 0x96,
	0x85, 0xc0, 0xff, 0xbb, 0x89, 0xfc, 0xb4, 0x11, 0xbb, 0x19, 0x27, 0xe3, 0x9e, 0xda, 0x99, 0xa6,
	0x83, 0xcb, 0xe0, 0xd9, 0x2c, 0xc2, 0x4b, 0x60, 0xb6, 0xe2, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xdf, 0x96, 0xe4, 0x45, 0x23, 0x02, 0x00, 0x00,
}
