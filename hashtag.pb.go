// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hashtag.proto

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

type GetHashTagRequest_Sort int32

const (
	GetHashTagRequest_new      GetHashTagRequest_Sort = 0
	GetHashTagRequest_hot      GetHashTagRequest_Sort = 1
	GetHashTagRequest_engaging GetHashTagRequest_Sort = 2
)

var GetHashTagRequest_Sort_name = map[int32]string{
	0: "new",
	1: "hot",
	2: "engaging",
}

var GetHashTagRequest_Sort_value = map[string]int32{
	"new":      0,
	"hot":      1,
	"engaging": 2,
}

func (x GetHashTagRequest_Sort) String() string {
	return proto.EnumName(GetHashTagRequest_Sort_name, int32(x))
}

func (GetHashTagRequest_Sort) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61450c0ee072706f, []int{2, 0}
}

type HashTag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashTag) Reset()         { *m = HashTag{} }
func (m *HashTag) String() string { return proto.CompactTextString(m) }
func (*HashTag) ProtoMessage()    {}
func (*HashTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_61450c0ee072706f, []int{0}
}

func (m *HashTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashTag.Unmarshal(m, b)
}
func (m *HashTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashTag.Marshal(b, m, deterministic)
}
func (m *HashTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashTag.Merge(m, src)
}
func (m *HashTag) XXX_Size() int {
	return xxx_messageInfo_HashTag.Size(m)
}
func (m *HashTag) XXX_DiscardUnknown() {
	xxx_messageInfo_HashTag.DiscardUnknown(m)
}

var xxx_messageInfo_HashTag proto.InternalMessageInfo

func (m *HashTag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HashTag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HashTagItem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashTagItem) Reset()         { *m = HashTagItem{} }
func (m *HashTagItem) String() string { return proto.CompactTextString(m) }
func (*HashTagItem) ProtoMessage()    {}
func (*HashTagItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_61450c0ee072706f, []int{1}
}

func (m *HashTagItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashTagItem.Unmarshal(m, b)
}
func (m *HashTagItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashTagItem.Marshal(b, m, deterministic)
}
func (m *HashTagItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashTagItem.Merge(m, src)
}
func (m *HashTagItem) XXX_Size() int {
	return xxx_messageInfo_HashTagItem.Size(m)
}
func (m *HashTagItem) XXX_DiscardUnknown() {
	xxx_messageInfo_HashTagItem.DiscardUnknown(m)
}

var xxx_messageInfo_HashTagItem proto.InternalMessageInfo

func (m *HashTagItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HashTagItem) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type GetHashTagRequest struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category             string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Page                 int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy               GetHashTagRequest_Sort `protobuf:"varint,5,opt,name=sortBy,proto3,enum=discussproto.GetHashTagRequest_Sort" json:"sortBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetHashTagRequest) Reset()         { *m = GetHashTagRequest{} }
func (m *GetHashTagRequest) String() string { return proto.CompactTextString(m) }
func (*GetHashTagRequest) ProtoMessage()    {}
func (*GetHashTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61450c0ee072706f, []int{2}
}

func (m *GetHashTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHashTagRequest.Unmarshal(m, b)
}
func (m *GetHashTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHashTagRequest.Marshal(b, m, deterministic)
}
func (m *GetHashTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHashTagRequest.Merge(m, src)
}
func (m *GetHashTagRequest) XXX_Size() int {
	return xxx_messageInfo_GetHashTagRequest.Size(m)
}
func (m *GetHashTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHashTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHashTagRequest proto.InternalMessageInfo

func (m *GetHashTagRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetHashTagRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetHashTagRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetHashTagRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetHashTagRequest) GetSortBy() GetHashTagRequest_Sort {
	if m != nil {
		return m.SortBy
	}
	return GetHashTagRequest_new
}

type GetHashTagResponse struct {
	Data                 []int64  `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Page                 int32    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy               string   `protobuf:"bytes,6,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHashTagResponse) Reset()         { *m = GetHashTagResponse{} }
func (m *GetHashTagResponse) String() string { return proto.CompactTextString(m) }
func (*GetHashTagResponse) ProtoMessage()    {}
func (*GetHashTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61450c0ee072706f, []int{3}
}

func (m *GetHashTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHashTagResponse.Unmarshal(m, b)
}
func (m *GetHashTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHashTagResponse.Marshal(b, m, deterministic)
}
func (m *GetHashTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHashTagResponse.Merge(m, src)
}
func (m *GetHashTagResponse) XXX_Size() int {
	return xxx_messageInfo_GetHashTagResponse.Size(m)
}
func (m *GetHashTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHashTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHashTagResponse proto.InternalMessageInfo

func (m *GetHashTagResponse) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetHashTagResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetHashTagResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetHashTagResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetHashTagResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetHashTagResponse) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func init() {
	proto.RegisterEnum("discussproto.GetHashTagRequest_Sort", GetHashTagRequest_Sort_name, GetHashTagRequest_Sort_value)
	proto.RegisterType((*HashTag)(nil), "discussproto.HashTag")
	proto.RegisterType((*HashTagItem)(nil), "discussproto.HashTagItem")
	proto.RegisterType((*GetHashTagRequest)(nil), "discussproto.GetHashTagRequest")
	proto.RegisterType((*GetHashTagResponse)(nil), "discussproto.GetHashTagResponse")
}

func init() { proto.RegisterFile("hashtag.proto", fileDescriptor_61450c0ee072706f) }

var fileDescriptor_61450c0ee072706f = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x14, 0xdc, 0xce, 0xd7, 0xae, 0xcf, 0x35, 0x8c, 0x8d, 0x48, 0x08, 0x82, 0x21, 0x88, 0xe4, 0x32,
	0x09, 0xac, 0x20, 0x08, 0x9e, 0xf6, 0xa2, 0x5e, 0xb3, 0x82, 0xe0, 0xad, 0x37, 0x69, 0x3a, 0x8d,
	0x93, 0xee, 0x98, 0x7e, 0xb3, 0xc3, 0x5e, 0xfd, 0x0b, 0xde, 0xfc, 0x5b, 0xfb, 0x17, 0xfc, 0x21,
	0x92, 0x4e, 0xc8, 0xcc, 0x30, 0x33, 0xee, 0x29, 0x55, 0x49, 0xd5, 0x7b, 0x55, 0x8f, 0xc0, 0xb3,
	0x86, 0x99, 0x06, 0x99, 0xc8, 0xbb, 0x5e, 0xa3, 0xa6, 0x97, 0xb5, 0x34, 0xd5, 0xda, 0x18, 0xcb,
	0xe2, 0x57, 0x42, 0x6b, 0xb1, 0xe2, 0x05, 0xeb, 0x64, 0xc1, 0x94, 0xd2, 0xc8, 0x50, 0x6a, 0x65,
	0x46, 0x6d, 0xfc, 0x5e, 0x48, 0x6c, 0xd6, 0xb7, 0x79, 0xa5, 0xdb, 0xa2, 0xdd, 0x48, 0xfc, 0xa1,
	0x37, 0x85, 0xd0, 0x4b, 0xfb, 0x71, 0x79, 0xc7, 0x56, 0xb2, 0x66, 0xa8, 0x7b, 0x53, 0xcc, 0x70,
	0xf4, 0xa5, 0x4b, 0x38, 0xff, 0xcc, 0x4c, 0xf3, 0x95, 0x09, 0x1a, 0x82, 0x23, 0xeb, 0x88, 0x24,
	0x24, 0x73, 0x4b, 0x47, 0xd6, 0x94, 0x82, 0xa7, 0x58, 0xcb, 0x23, 0x27, 0x21, 0xd9, 0x93, 0xd2,
	0xe2, 0xf4, 0x03, 0x3c, 0x9d, 0xe4, 0x5f, 0x90, 0xb7, 0x07, 0x96, 0x18, 0x2e, 0x2a, 0x86, 0x5c,
	0xe8, 0xfe, 0x7e, 0xb2, 0xcd, 0x3c, 0x7d, 0x20, 0xf0, 0xfc, 0x13, 0xc7, 0xc9, 0x5e, 0xf2, 0x9f,
	0x6b, 0x6e, 0x70, 0x5e, 0x42, 0xb6, 0x4b, 0xfe, 0x37, 0x65, 0xd0, 0x77, 0x4c, 0xf0, 0xc8, 0x4d,
	0x48, 0xe6, 0x97, 0x16, 0xd3, 0x17, 0xe0, 0xaf, 0x64, 0x2b, 0x31, 0xf2, 0xec, 0xcb, 0x91, 0xd0,
	0x8f, 0x10, 0x18, 0xdd, 0xe3, 0xf5, 0x7d, 0xe4, 0x27, 0x24, 0x0b, 0xaf, 0xde, 0xe4, 0xbb, 0xe7,
	0xcc, 0x0f, 0xa2, 0xe4, 0x37, 0xba, 0xc7, 0x72, 0xf2, 0xa4, 0x6f, 0xc1, 0x1b, 0x38, 0x3d, 0x07,
	0x57, 0xf1, 0xcd, 0xe2, 0x6c, 0x00, 0x8d, 0xc6, 0x05, 0xa1, 0x97, 0x70, 0xc1, 0x95, 0x60, 0x42,
	0x2a, 0xb1, 0x70, 0xd2, 0x3f, 0x04, 0xe8, 0xee, 0x28, 0xd3, 0x69, 0x65, 0xf8, 0x10, 0xb3, 0x66,
	0xc8, 0x22, 0x92, 0xb8, 0x99, 0x5b, 0x5a, 0x7c, 0xec, 0x9e, 0x7b, 0x55, 0xdd, 0x13, 0x55, 0xbd,
	0x63, 0x55, 0xfd, 0xdd, 0xaa, 0x2f, 0xe7, 0xaa, 0x81, 0x9d, 0x31, 0xb1, 0x2b, 0x09, 0xe1, 0x14,
	0xec, 0x86, 0xf7, 0x77, 0xb2, 0xe2, 0xf4, 0x1b, 0xc0, 0x36, 0x2d, 0x7d, 0xfd, 0xc8, 0x49, 0xe2,
	0xe4, 0xb4, 0x60, 0x2c, 0x9a, 0x06, 0xbf, 0x1e, 0xfe, 0xfe, 0x76, 0xce, 0xae, 0xc3, 0xef, 0x7b,
	0x7f, 0xeb, 0x6d, 0x60, 0x1f, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x4d, 0xbf, 0x28,
	0xd3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HashTagServiceClient is the client API for HashTagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HashTagServiceClient interface {
	// rpc GetHashTags(GetHashTagsRequest) returns (GetHashTagsResponse) {
	//   option (google.api.http) = {
	//     get: "/api/hashtag"
	//   };
	// }
	GetHashTag(ctx context.Context, in *GetHashTagRequest, opts ...grpc.CallOption) (*GetHashTagResponse, error)
}

type hashTagServiceClient struct {
	cc *grpc.ClientConn
}

func NewHashTagServiceClient(cc *grpc.ClientConn) HashTagServiceClient {
	return &hashTagServiceClient{cc}
}

func (c *hashTagServiceClient) GetHashTag(ctx context.Context, in *GetHashTagRequest, opts ...grpc.CallOption) (*GetHashTagResponse, error) {
	out := new(GetHashTagResponse)
	err := c.cc.Invoke(ctx, "/discussproto.HashTagService/GetHashTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashTagServiceServer is the server API for HashTagService service.
type HashTagServiceServer interface {
	// rpc GetHashTags(GetHashTagsRequest) returns (GetHashTagsResponse) {
	//   option (google.api.http) = {
	//     get: "/api/hashtag"
	//   };
	// }
	GetHashTag(context.Context, *GetHashTagRequest) (*GetHashTagResponse, error)
}

func RegisterHashTagServiceServer(s *grpc.Server, srv HashTagServiceServer) {
	s.RegisterService(&_HashTagService_serviceDesc, srv)
}

func _HashTagService_GetHashTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTagServiceServer).GetHashTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.HashTagService/GetHashTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTagServiceServer).GetHashTag(ctx, req.(*GetHashTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HashTagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discussproto.HashTagService",
	HandlerType: (*HashTagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHashTag",
			Handler:    _HashTagService_GetHashTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hashtag.proto",
}
