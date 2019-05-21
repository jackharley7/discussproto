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
	Unseen               int32    `protobuf:"varint,2,opt,name=unseen,proto3" json:"unseen,omitempty"`
	UnseenInvite         int32    `protobuf:"varint,3,opt,name=unseenInvite,proto3" json:"unseenInvite,omitempty"`
	TotalInvite          int32    `protobuf:"varint,4,opt,name=totalInvite,proto3" json:"totalInvite,omitempty"`
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

func (m *Count) GetUnseen() int32 {
	if m != nil {
		return m.Unseen
	}
	return 0
}

func (m *Count) GetUnseenInvite() int32 {
	if m != nil {
		return m.UnseenInvite
	}
	return 0
}

func (m *Count) GetTotalInvite() int32 {
	if m != nil {
		return m.TotalInvite
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
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6b, 0x7b, 0xd8, 0x46, 0x0f, 0xab, 0x96, 0x52, 0x15, 0x4a, 0x7a, 0x29, 0x48,
	0xb3, 0x50, 0xc1, 0x07, 0x50, 0x41, 0x7a, 0x8d, 0x78, 0xf1, 0xb6, 0x4d, 0xc7, 0xb8, 0xd8, 0xee,
	0xd4, 0xcc, 0xa4, 0x3d, 0xf5, 0xe2, 0xdd, 0x93, 0x8f, 0xe6, 0x2b, 0xf8, 0x20, 0xe2, 0x24, 0x48,
	0x23, 0xe2, 0x29, 0x33, 0xff, 0x7c, 0x93, 0xfd, 0x99, 0x5f, 0x85, 0x29, 0x16, 0x9e, 0x29, 0x5e,
	0xe5, 0xc8, 0xa8, 0xc3, 0xb9, 0xa3, 0xb4, 0x20, 0x92, 0xae, 0x7f, 0x9a, 0x21, 0x66, 0x0b, 0x30,
	0x76, 0xe5, 0x8c, 0xf5, 0x1e, 0xd9, 0xb2, 0x43, 0x5f, 0xb1, 0xfd, 0xcb, 0xcc, 0xf1, 0x53, 0x31,
	0x8b, 0x53, 0x5c, 0x9a, 0xe5, 0xc6, 0xf1, 0x33, 0x6e, 0x4c, 0x86, 0x63, 0x19, 0x8e, 0xd7, 0x76,
	0xe1, 0xe6, 0x96, 0x31, 0x27, 0xf3, 0x53, 0x96, 0x7b, 0xd1, 0x56, 0xb5, 0xae, 0xbf, 0xdf, 0xd4,
	0x5d, 0xd5, 0x2e, 0x08, 0xf2, 0xe9, 0xbc, 0x17, 0x0c, 0x82, 0x51, 0x33, 0xa9, 0x3a, 0xd1, 0x3d,
	0x01, 0xf8, 0x5e, 0x63, 0x10, 0x8c, 0x5a, 0x49, 0xd5, 0xe9, 0x48, 0x85, 0x65, 0x35, 0xf5, 0x6b,
	0xc7, 0xd0, 0x6b, 0xca, 0xb4, 0xa6, 0xe9, 0x81, 0xea, 0x30, 0xb2, 0x5d, 0x54, 0xc8, 0x9e, 0x20,
	0xbb, 0x52, 0xd4, 0x55, 0x47, 0xb7, 0xc0, 0xf7, 0x04, 0xb9, 0xb8, 0xa0, 0x04, 0x5e, 0x0a, 0x20,
	0x8e, 0x6e, 0xd4, 0xf1, 0x2f, 0x9d, 0x56, 0xe8, 0x09, 0xf4, 0xb9, 0x6a, 0x97, 0x37, 0x12, 0x9b,
	0x9d, 0xc9, 0x61, 0xbc, 0x7b, 0xa4, 0x58, 0xe8, 0xa4, 0x42, 0x26, 0x6f, 0x81, 0x0a, 0x45, 0xb9,
	0x83, 0x7c, 0xed, 0x52, 0xd0, 0x5b, 0xb5, 0x5f, 0xfb, 0xad, 0x8e, 0xea, 0xeb, 0x7f, 0x79, 0xe9,
	0x0f, 0xff, 0x65, 0x4a, 0x5f, 0xd1, 0xf0, 0xf5, 0xe3, 0xf3, 0xbd, 0x71, 0xa6, 0x4f, 0x24, 0x1f,
	0x8f, 0xec, 0x1e, 0x5d, 0x2a, 0x09, 0x99, 0xd2, 0x8c, 0x59, 0xc2, 0xd5, 0xc1, 0x43, 0x2d, 0xd2,
	0x59, 0x5b, 0x3e, 0x17, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x63, 0x9d, 0xdd, 0xf7, 0x01,
	0x00, 0x00,
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
