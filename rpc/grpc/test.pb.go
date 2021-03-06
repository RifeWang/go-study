// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserOrders_Order_Status int32

const (
	UserOrders_Order_SATAUS_INITIAL UserOrders_Order_Status = 0
	UserOrders_Order_SATAUS_SUCCESS UserOrders_Order_Status = 1
	UserOrders_Order_SATAUS_FAIL    UserOrders_Order_Status = 2
)

var UserOrders_Order_Status_name = map[int32]string{
	0: "SATAUS_INITIAL",
	1: "SATAUS_SUCCESS",
	2: "SATAUS_FAIL",
}

var UserOrders_Order_Status_value = map[string]int32{
	"SATAUS_INITIAL": 0,
	"SATAUS_SUCCESS": 1,
	"SATAUS_FAIL":    2,
}

func (x UserOrders_Order_Status) String() string {
	return proto.EnumName(UserOrders_Order_Status_name, int32(x))
}

func (UserOrders_Order_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1, 0, 0}
}

type ReqBody struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pagesize             int32    `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBody) Reset()         { *m = ReqBody{} }
func (m *ReqBody) String() string { return proto.CompactTextString(m) }
func (*ReqBody) ProtoMessage()    {}
func (*ReqBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *ReqBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBody.Unmarshal(m, b)
}
func (m *ReqBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBody.Marshal(b, m, deterministic)
}
func (m *ReqBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBody.Merge(m, src)
}
func (m *ReqBody) XXX_Size() int {
	return xxx_messageInfo_ReqBody.Size(m)
}
func (m *ReqBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBody.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBody proto.InternalMessageInfo

func (m *ReqBody) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReqBody) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReqBody) GetPagesize() int32 {
	if m != nil {
		return m.Pagesize
	}
	return 0
}

type UserOrders struct {
	Id                   int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string              `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string              `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string              `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Orders               []*UserOrders_Order `protobuf:"bytes,5,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserOrders) Reset()         { *m = UserOrders{} }
func (m *UserOrders) String() string { return proto.CompactTextString(m) }
func (*UserOrders) ProtoMessage()    {}
func (*UserOrders) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *UserOrders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrders.Unmarshal(m, b)
}
func (m *UserOrders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrders.Marshal(b, m, deterministic)
}
func (m *UserOrders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrders.Merge(m, src)
}
func (m *UserOrders) XXX_Size() int {
	return xxx_messageInfo_UserOrders.Size(m)
}
func (m *UserOrders) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrders.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrders proto.InternalMessageInfo

func (m *UserOrders) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserOrders) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserOrders) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserOrders) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserOrders) GetOrders() []*UserOrders_Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type UserOrders_Order struct {
	OrderId              int32                   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Info                 string                  `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Status               UserOrders_Order_Status `protobuf:"varint,3,opt,name=status,proto3,enum=grpc.UserOrders_Order_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UserOrders_Order) Reset()         { *m = UserOrders_Order{} }
func (m *UserOrders_Order) String() string { return proto.CompactTextString(m) }
func (*UserOrders_Order) ProtoMessage()    {}
func (*UserOrders_Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1, 0}
}

func (m *UserOrders_Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrders_Order.Unmarshal(m, b)
}
func (m *UserOrders_Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrders_Order.Marshal(b, m, deterministic)
}
func (m *UserOrders_Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrders_Order.Merge(m, src)
}
func (m *UserOrders_Order) XXX_Size() int {
	return xxx_messageInfo_UserOrders_Order.Size(m)
}
func (m *UserOrders_Order) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrders_Order.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrders_Order proto.InternalMessageInfo

func (m *UserOrders_Order) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *UserOrders_Order) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *UserOrders_Order) GetStatus() UserOrders_Order_Status {
	if m != nil {
		return m.Status
	}
	return UserOrders_Order_SATAUS_INITIAL
}

func init() {
	proto.RegisterEnum("grpc.UserOrders_Order_Status", UserOrders_Order_Status_name, UserOrders_Order_Status_value)
	proto.RegisterType((*ReqBody)(nil), "grpc.ReqBody")
	proto.RegisterType((*UserOrders)(nil), "grpc.UserOrders")
	proto.RegisterType((*UserOrders_Order)(nil), "grpc.UserOrders.Order")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0xda, 0xa4, 0xed, 0x14, 0xdb, 0x30, 0x88, 0xc6, 0x82, 0x50, 0x7a, 0xea, 0x29,
	0x87, 0x88, 0x77, 0x63, 0x51, 0x08, 0x14, 0xff, 0xec, 0xb6, 0xe7, 0x12, 0x9b, 0xb1, 0x06, 0x6c,
	0x13, 0x77, 0x53, 0xa1, 0x3e, 0x95, 0x6f, 0xe1, 0x6b, 0x49, 0x26, 0x4b, 0x2d, 0xe2, 0x69, 0xe7,
	0xfb, 0x31, 0xf9, 0xe6, 0x9b, 0x0c, 0x40, 0x49, 0xba, 0x0c, 0x0a, 0x95, 0x97, 0x39, 0x36, 0x57,
	0xaa, 0x58, 0x8e, 0x04, 0xb4, 0x04, 0xbd, 0xdf, 0xe4, 0xe9, 0x0e, 0xcf, 0xa0, 0xb5, 0xd5, 0xa4,
	0x16, 0x59, 0xea, 0x5b, 0x43, 0x6b, 0xdc, 0x11, 0x6e, 0x25, 0xe3, 0x14, 0x11, 0x9a, 0x45, 0xb2,
	0x22, 0xdf, 0x1e, 0x5a, 0x63, 0x47, 0x70, 0x8d, 0x03, 0x68, 0x57, 0xaf, 0xce, 0x3e, 0xc9, 0x6f,
	0x30, 0xdf, 0xeb, 0xd1, 0xb7, 0x0d, 0x30, 0xd7, 0xa4, 0x1e, 0x54, 0x4a, 0x4a, 0x63, 0x0f, 0x6c,
	0x63, 0xe9, 0x08, 0x3b, 0x4b, 0xab, 0x4f, 0x2b, 0xe3, 0x4d, 0xb2, 0xae, 0x2d, 0x3b, 0x62, 0xaf,
	0xf1, 0x04, 0x1c, 0x5a, 0x27, 0xd9, 0x1b, 0x7b, 0x76, 0x44, 0x2d, 0x2a, 0x5a, 0xbc, 0xe6, 0x1b,
	0xf2, 0x9b, 0x35, 0x65, 0x81, 0x01, 0xb8, 0x39, 0x4f, 0xf0, 0x9d, 0x61, 0x63, 0xdc, 0x0d, 0x4f,
	0x83, 0x6a, 0xa3, 0xe0, 0x77, 0x72, 0xc0, 0x8f, 0x30, 0x5d, 0x83, 0x2f, 0x0b, 0x1c, 0x26, 0x78,
	0x0e, 0x6d, 0x66, 0x8b, 0x7d, 0xae, 0x16, 0xeb, 0x7a, 0xd7, 0x6c, 0xf3, 0x92, 0x9b, 0x60, 0x5c,
	0xe3, 0x15, 0xb8, 0xba, 0x4c, 0xca, 0xad, 0xe6, 0x54, 0xbd, 0xf0, 0xe2, 0xff, 0x41, 0x81, 0xe4,
	0x26, 0x61, 0x9a, 0x47, 0x11, 0xb8, 0x35, 0x41, 0x84, 0x9e, 0x8c, 0x66, 0xd1, 0x5c, 0x2e, 0xe2,
	0xfb, 0x78, 0x16, 0x47, 0x53, 0xef, 0xe8, 0x80, 0xc9, 0xf9, 0x64, 0x72, 0x2b, 0xa5, 0x67, 0x61,
	0x1f, 0xba, 0x86, 0xdd, 0x45, 0xf1, 0xd4, 0xb3, 0xc3, 0x6b, 0x00, 0xf1, 0x38, 0x91, 0xa4, 0x3e,
	0xb2, 0x25, 0x61, 0x08, 0xfd, 0xa7, 0x2d, 0xa9, 0xdd, 0xc1, 0xbf, 0x3d, 0xae, 0xa3, 0x98, 0x13,
	0x0e, 0xbc, 0xbf, 0xc9, 0x9e, 0x5d, 0x3e, 0xf6, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x03, 0x41, 0xe0, 0xfa, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	QueryUserOrders(ctx context.Context, in *ReqBody, opts ...grpc.CallOption) (*UserOrders, error)
}

type rPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewRPCServiceClient(cc *grpc.ClientConn) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) QueryUserOrders(ctx context.Context, in *ReqBody, opts ...grpc.CallOption) (*UserOrders, error) {
	out := new(UserOrders)
	err := c.cc.Invoke(ctx, "/grpc.RPCService/QueryUserOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	QueryUserOrders(context.Context, *ReqBody) (*UserOrders, error)
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_QueryUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).QueryUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.RPCService/QueryUserOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).QueryUserOrders(ctx, req.(*ReqBody))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUserOrders",
			Handler:    _RPCService_QueryUserOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
