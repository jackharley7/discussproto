// Code generated by protoc-gen-go. DO NOT EDIT.
// source: counts.proto

package discussproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Count struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PrimaryCount         int32    `protobuf:"varint,2,opt,name=primaryCount,proto3" json:"primaryCount,omitempty"`
	SecondaryCount       int32    `protobuf:"varint,3,opt,name=secondaryCount,proto3" json:"secondaryCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1dca8b389756398, []int{0}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Count) GetPrimaryCount() int32 {
	if m != nil {
		return m.PrimaryCount
	}
	return 0
}

func (m *Count) GetSecondaryCount() int32 {
	if m != nil {
		return m.SecondaryCount
	}
	return 0
}

type GetUserCountsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserCountsRequest) Reset()         { *m = GetUserCountsRequest{} }
func (m *GetUserCountsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserCountsRequest) ProtoMessage()    {}
func (*GetUserCountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1dca8b389756398, []int{1}
}

func (m *GetUserCountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCountsRequest.Unmarshal(m, b)
}
func (m *GetUserCountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCountsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserCountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCountsRequest.Merge(m, src)
}
func (m *GetUserCountsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserCountsRequest.Size(m)
}
func (m *GetUserCountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCountsRequest proto.InternalMessageInfo

type GetUserCountsResponse struct {
	Counts               *Count   `protobuf:"bytes,1,opt,name=counts,proto3" json:"counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserCountsResponse) Reset()         { *m = GetUserCountsResponse{} }
func (m *GetUserCountsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserCountsResponse) ProtoMessage()    {}
func (*GetUserCountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1dca8b389756398, []int{2}
}

func (m *GetUserCountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCountsResponse.Unmarshal(m, b)
}
func (m *GetUserCountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCountsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserCountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCountsResponse.Merge(m, src)
}
func (m *GetUserCountsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserCountsResponse.Size(m)
}
func (m *GetUserCountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCountsResponse proto.InternalMessageInfo

func (m *GetUserCountsResponse) GetCounts() *Count {
	if m != nil {
		return m.Counts
	}
	return nil
}

func init() {
	proto.RegisterType((*Count)(nil), "discussproto.Count")
	proto.RegisterType((*GetUserCountsRequest)(nil), "discussproto.GetUserCountsRequest")
	proto.RegisterType((*GetUserCountsResponse)(nil), "discussproto.GetUserCountsResponse")
}

func init() { proto.RegisterFile("counts.proto", fileDescriptor_b1dca8b389756398) }

var fileDescriptor_b1dca8b389756398 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xc6, 0x7a, 0x88, 0x75, 0x87, 0xa8, 0x63, 0x54, 0x85, 0x91, 0x81, 0x0c, 0x64,
	0x0d, 0x4c, 0xf0, 0x07, 0xa8, 0x20, 0x5e, 0x2b, 0x5e, 0xbc, 0x65, 0x69, 0xac, 0x61, 0x6b, 0xbe,
	0x9a, 0x2f, 0xdd, 0xf0, 0xe0, 0xc5, 0xbb, 0x27, 0x7f, 0x9a, 0x7f, 0xc1, 0x1f, 0x22, 0xa6, 0x65,
	0xd8, 0x21, 0x3b, 0xb5, 0xef, 0x9b, 0xe7, 0x4b, 0x5e, 0xbe, 0x97, 0x44, 0x12, 0x2a, 0xe3, 0x30,
	0x29, 0x2d, 0x38, 0xa0, 0x51, 0xa6, 0x51, 0x56, 0x88, 0x5e, 0xc5, 0x27, 0x39, 0x40, 0xbe, 0x54,
	0x5c, 0x94, 0x9a, 0x0b, 0x63, 0xc0, 0x09, 0xa7, 0xc1, 0x34, 0x6c, 0x7c, 0x99, 0x6b, 0xf7, 0x5c,
	0xcd, 0x13, 0x09, 0x05, 0x2f, 0xd6, 0xda, 0x2d, 0x60, 0xcd, 0x73, 0x98, 0xfa, 0xc3, 0xe9, 0x4a,
	0x2c, 0x75, 0x26, 0x1c, 0x58, 0xe4, 0x9b, 0xdf, 0x7a, 0x8e, 0x2d, 0x48, 0xef, 0xfa, 0xf7, 0x4d,
	0x3a, 0x20, 0x61, 0x85, 0xca, 0xde, 0x65, 0xc3, 0x60, 0x14, 0x4c, 0xba, 0x69, 0xa3, 0x28, 0x23,
	0x51, 0x69, 0x75, 0x21, 0xec, 0xab, 0xe7, 0x86, 0x9d, 0x51, 0x30, 0xe9, 0xa5, 0x2d, 0x8f, 0x9e,
	0x91, 0x3e, 0x2a, 0x09, 0x26, 0xdb, 0x50, 0x5d, 0x4f, 0x6d, 0xb9, 0x6c, 0x40, 0x0e, 0x6f, 0x95,
	0x7b, 0x40, 0x65, 0xbd, 0xc6, 0x54, 0xbd, 0x54, 0x0a, 0x1d, 0xbb, 0x21, 0x47, 0x5b, 0x3e, 0x96,
	0x60, 0x50, 0xd1, 0x73, 0x12, 0xd6, 0x1b, 0xf1, 0xa1, 0xf6, 0x66, 0x07, 0xc9, 0xdf, 0x95, 0x24,
	0x9e, 0x4e, 0x1b, 0x64, 0xf6, 0x11, 0x90, 0xc8, 0x3b, 0xf7, 0xca, 0xae, 0xb4, 0x54, 0xf4, 0x8d,
	0xec, 0xb7, 0xae, 0xa5, 0xac, 0x3d, 0xfe, 0x5f, 0x96, 0x78, 0xbc, 0x93, 0xa9, 0x73, 0xb1, 0xf1,
	0xfb, 0xd7, 0xf7, 0x67, 0xe7, 0x94, 0x1e, 0xfb, 0x36, 0x0c, 0x38, 0xfd, 0xa4, 0xa5, 0xef, 0x83,
	0xd7, 0x61, 0x78, 0xa1, 0xae, 0xfa, 0x8f, 0xad, 0x02, 0xe7, 0xa1, 0xff, 0x5c, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0x29, 0x88, 0x4b, 0xe5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CountServiceClient is the client API for CountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountServiceClient interface {
	GetUserCounts(ctx context.Context, in *GetUserCountsRequest, opts ...grpc.CallOption) (*GetUserCountsResponse, error)
}

type countServiceClient struct {
	cc *grpc.ClientConn
}

func NewCountServiceClient(cc *grpc.ClientConn) CountServiceClient {
	return &countServiceClient{cc}
}

func (c *countServiceClient) GetUserCounts(ctx context.Context, in *GetUserCountsRequest, opts ...grpc.CallOption) (*GetUserCountsResponse, error) {
	out := new(GetUserCountsResponse)
	err := c.cc.Invoke(ctx, "/discussproto.CountService/GetUserCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountServiceServer is the server API for CountService service.
type CountServiceServer interface {
	GetUserCounts(context.Context, *GetUserCountsRequest) (*GetUserCountsResponse, error)
}

// UnimplementedCountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCountServiceServer struct {
}

func (*UnimplementedCountServiceServer) GetUserCounts(ctx context.Context, req *GetUserCountsRequest) (*GetUserCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCounts not implemented")
}

func RegisterCountServiceServer(s *grpc.Server, srv CountServiceServer) {
	s.RegisterService(&_CountService_serviceDesc, srv)
}

func _CountService_GetUserCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountServiceServer).GetUserCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.CountService/GetUserCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountServiceServer).GetUserCounts(ctx, req.(*GetUserCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discussproto.CountService",
	HandlerType: (*CountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCounts",
			Handler:    _CountService_GetUserCounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counts.proto",
}
