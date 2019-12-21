// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

package frontend

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

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Type                 bool     `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetRequest) GetType() bool {
	if m != nil {
		return m.Type
	}
	return false
}

type GetResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetResponse) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type PutRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type                 bool     `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{2}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PutRequest) GetType() bool {
	if m != nil {
		return m.Type
	}
	return false
}

type PutResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{3}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{4}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "frontend.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "frontend.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "frontend.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "frontend.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "frontend.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "frontend.DeleteResponse")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe) }

var fileDescriptor_eca3873955a29cfe = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0xc6, 0x94, 0x74, 0x4a, 0xab, 0x2c, 0x55, 0x43, 0x4f, 0x75, 0x41, 0xc8, 0xa9,
	0x87, 0xaa, 0x17, 0x8f, 0xa5, 0x58, 0x8f, 0x61, 0xdf, 0x20, 0x92, 0x51, 0xa4, 0x21, 0x89, 0xdd,
	0xd9, 0x42, 0x7d, 0x2c, 0x9f, 0x50, 0x76, 0x93, 0x4d, 0x52, 0xc5, 0xdc, 0xfe, 0xd9, 0x9d, 0xef,
	0xdf, 0x99, 0x9f, 0x85, 0xe9, 0xdb, 0xbe, 0xc8, 0x09, 0xf3, 0x74, 0x59, 0xee, 0x0b, 0x2a, 0x78,
	0xe0, 0x6a, 0xb1, 0x02, 0xd8, 0x22, 0x49, 0xfc, 0xd4, 0xa8, 0x88, 0x5f, 0x82, 0xb7, 0xc3, 0x63,
	0xc8, 0x16, 0x2c, 0x1a, 0x49, 0x23, 0x39, 0x87, 0x73, 0x3a, 0x96, 0x18, 0x9e, 0x2d, 0x58, 0x14,
	0x48, 0xab, 0xc5, 0x23, 0x8c, 0x2d, 0xa3, 0xca, 0x22, 0x57, 0x68, 0x5a, 0xd2, 0x84, 0x92, 0x9a,
	0xb2, 0xda, 0x9c, 0xa9, 0x8f, 0xaf, 0x0a, 0xf3, 0xa5, 0xd5, 0xe2, 0x05, 0x20, 0xd6, 0x3d, 0x4f,
	0xcd, 0xc0, 0x3f, 0x24, 0x99, 0xae, 0xa0, 0x91, 0xac, 0x8a, 0x66, 0x00, 0xaf, 0x33, 0xc0, 0x1d,
	0x8c, 0xad, 0x53, 0x3d, 0xc0, 0x35, 0x0c, 0x15, 0x25, 0xa4, 0x95, 0x75, 0xf3, 0x65, 0x5d, 0x89,
	0x5b, 0x98, 0x6c, 0x30, 0x43, 0xc2, 0x7f, 0xdf, 0x14, 0x11, 0x4c, 0x5d, 0x4b, 0xbf, 0xd9, 0xea,
	0x9b, 0x41, 0xf0, 0x5c, 0xa7, 0xc6, 0x1f, 0xc0, 0xdb, 0x22, 0xf1, 0xd9, 0xb2, 0xc9, 0xb5, 0x0d,
	0x71, 0x7e, 0xf5, 0xeb, 0xb4, 0x32, 0x16, 0x03, 0x43, 0xc5, 0xfa, 0x84, 0x6a, 0xf3, 0xe8, 0x52,
	0x9d, 0xdd, 0xc4, 0x80, 0x3f, 0x81, 0xb7, 0xc1, 0x8c, 0xdf, 0xb4, 0xf7, 0x27, 0x4b, 0xcd, 0xc3,
	0xbf, 0x17, 0x8e, 0x5d, 0x0b, 0xb8, 0xc0, 0xf4, 0x1d, 0x77, 0x87, 0xa6, 0x67, 0x3d, 0x71, 0x4b,
	0xc4, 0xe6, 0x27, 0xc4, 0xec, 0x75, 0x68, 0xbf, 0xc4, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0x1f, 0x93, 0xda, 0x24, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Del(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/frontend.Frontend/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/frontend.Frontend/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) Del(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/frontend.Frontend/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Del(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedFrontendServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (*UnimplementedFrontendServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedFrontendServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedFrontendServer) Del(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.Frontend/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.Frontend/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.Frontend/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Del(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontend.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Frontend_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Frontend_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Frontend_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frontend.proto",
}
