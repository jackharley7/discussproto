// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

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

type Profile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Bio                  string   `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Website              string   `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Locale               string   `protobuf:"bytes,5,opt,name=locale,proto3" json:"locale,omitempty"`
	Postcode             string   `protobuf:"bytes,6,opt,name=postcode,proto3" json:"postcode,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Progress             int32    `protobuf:"varint,8,opt,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Profile) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *Profile) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Profile) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Profile) GetPostcode() string {
	if m != nil {
		return m.Postcode
	}
	return ""
}

func (m *Profile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Profile) GetProgress() int32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

type UpdateProfileRequest struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileRequest) Reset()         { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()    {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *UpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileRequest.Unmarshal(m, b)
}
func (m *UpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileRequest.Merge(m, src)
}
func (m *UpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileRequest.Size(m)
}
func (m *UpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileRequest proto.InternalMessageInfo

func (m *UpdateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateProfileResponse struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileResponse) Reset()         { *m = UpdateProfileResponse{} }
func (m *UpdateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileResponse) ProtoMessage()    {}
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{2}
}

func (m *UpdateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileResponse.Unmarshal(m, b)
}
func (m *UpdateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileResponse.Merge(m, src)
}
func (m *UpdateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileResponse.Size(m)
}
func (m *UpdateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileResponse proto.InternalMessageInfo

func (m *UpdateProfileResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GetProfileProgressRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileProgressRequest) Reset()         { *m = GetProfileProgressRequest{} }
func (m *GetProfileProgressRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileProgressRequest) ProtoMessage()    {}
func (*GetProfileProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{3}
}

func (m *GetProfileProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileProgressRequest.Unmarshal(m, b)
}
func (m *GetProfileProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileProgressRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileProgressRequest.Merge(m, src)
}
func (m *GetProfileProgressRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileProgressRequest.Size(m)
}
func (m *GetProfileProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileProgressRequest proto.InternalMessageInfo

type GetProfileProgressResponse struct {
	Progress             int32    `protobuf:"varint,1,opt,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileProgressResponse) Reset()         { *m = GetProfileProgressResponse{} }
func (m *GetProfileProgressResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileProgressResponse) ProtoMessage()    {}
func (*GetProfileProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{4}
}

func (m *GetProfileProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileProgressResponse.Unmarshal(m, b)
}
func (m *GetProfileProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileProgressResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileProgressResponse.Merge(m, src)
}
func (m *GetProfileProgressResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileProgressResponse.Size(m)
}
func (m *GetProfileProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileProgressResponse proto.InternalMessageInfo

func (m *GetProfileProgressResponse) GetProgress() int32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func init() {
	proto.RegisterType((*Profile)(nil), "discussproto.Profile")
	proto.RegisterType((*UpdateProfileRequest)(nil), "discussproto.UpdateProfileRequest")
	proto.RegisterType((*UpdateProfileResponse)(nil), "discussproto.UpdateProfileResponse")
	proto.RegisterType((*GetProfileProgressRequest)(nil), "discussproto.GetProfileProgressRequest")
	proto.RegisterType((*GetProfileProgressResponse)(nil), "discussproto.GetProfileProgressResponse")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x95, 0x37, 0x34, 0x69, 0x07, 0x1a, 0x21, 0xab, 0x45, 0xee, 0x52, 0xd4, 0x68, 0x11, 0x22,
	0x42, 0x6a, 0x2c, 0x05, 0x09, 0x21, 0x6e, 0x70, 0x29, 0xdc, 0xaa, 0x20, 0x2e, 0xdc, 0x9c, 0xf5,
	0xb0, 0x58, 0x6c, 0x32, 0x8b, 0xed, 0x6d, 0xc4, 0x95, 0x0f, 0xe0, 0xc2, 0x95, 0x3f, 0xe2, 0xc8,
	0x2f, 0xf0, 0x21, 0x28, 0x5e, 0x6f, 0xc8, 0x42, 0x00, 0x71, 0x5a, 0xbf, 0x99, 0x37, 0xe3, 0xf7,
	0x9e, 0xbc, 0x70, 0x58, 0x59, 0x7a, 0x63, 0x4a, 0x9c, 0x54, 0x96, 0x3c, 0xf1, 0x1b, 0xda, 0xb8,
	0xbc, 0x76, 0x2e, 0xa0, 0xf4, 0xb4, 0x20, 0x2a, 0x4a, 0x94, 0xaa, 0x32, 0x52, 0x2d, 0x97, 0xe4,
	0x95, 0x37, 0xb4, 0x74, 0x0d, 0x37, 0x7d, 0x54, 0x18, 0xff, 0xb6, 0x9e, 0x4f, 0x72, 0x5a, 0xc8,
	0xc5, 0xca, 0xf8, 0x77, 0xb4, 0x92, 0x05, 0x9d, 0x87, 0xe6, 0xf9, 0x95, 0x2a, 0x8d, 0x56, 0x9e,
	0xac, 0x93, 0x9b, 0x63, 0x33, 0x97, 0x7d, 0x65, 0x30, 0xb8, 0x6c, 0x6e, 0xe5, 0x43, 0x48, 0x8c,
	0x16, 0x6c, 0xc4, 0xc6, 0xbd, 0x59, 0x62, 0x34, 0xbf, 0x05, 0xfd, 0xda, 0xa1, 0x7d, 0xa1, 0x45,
	0x12, 0x6a, 0x11, 0xf1, 0x9b, 0xd0, 0x9b, 0x1b, 0x12, 0xbd, 0x11, 0x1b, 0x1f, 0xcc, 0xd6, 0x47,
	0x2e, 0x60, 0xb0, 0xc2, 0xb9, 0x33, 0x1e, 0xc5, 0xb5, 0x50, 0x6d, 0xe1, 0x7a, 0x47, 0x49, 0xb9,
	0x2a, 0x51, 0xec, 0x85, 0x46, 0x44, 0x3c, 0x85, 0xfd, 0x8a, 0x9c, 0xcf, 0x49, 0xa3, 0xe8, 0x87,
	0xce, 0x06, 0xf3, 0x53, 0x38, 0xc8, 0x2d, 0x2a, 0x8f, 0xfa, 0xa9, 0x17, 0x83, 0xd0, 0xfc, 0x59,
	0x08, 0x93, 0x96, 0x0a, 0x8b, 0xce, 0x89, 0xfd, 0x11, 0x1b, 0xef, 0xcd, 0x36, 0x38, 0xbb, 0x80,
	0xa3, 0x57, 0x95, 0x56, 0x1e, 0xa3, 0xa5, 0x19, 0xbe, 0xaf, 0xd1, 0x79, 0x2e, 0x61, 0x10, 0xa3,
	0x0d, 0xf6, 0xae, 0x4f, 0x8f, 0x27, 0xdb, 0xd9, 0x4e, 0x5a, 0x7a, 0xcb, 0xca, 0x9e, 0xc3, 0xf1,
	0x2f, 0x8b, 0x5c, 0x45, 0x4b, 0x87, 0xff, 0xbf, 0xe9, 0x36, 0x9c, 0x5c, 0xa0, 0x8f, 0xe5, 0xcb,
	0x28, 0x34, 0xea, 0xca, 0x1e, 0x43, 0xba, 0xab, 0x19, 0xef, 0xda, 0x76, 0xca, 0xba, 0x4e, 0xa7,
	0x5f, 0x12, 0x18, 0xc6, 0xb9, 0x97, 0x68, 0xaf, 0x4c, 0x8e, 0xfc, 0x03, 0x1c, 0x76, 0x34, 0xf3,
	0xac, 0x2b, 0x6d, 0x57, 0x32, 0xe9, 0xdd, 0xbf, 0x72, 0x1a, 0x21, 0xd9, 0xd9, 0xc7, 0x6f, 0xdf,
	0x3f, 0x27, 0x27, 0xd3, 0xa3, 0xf0, 0xf8, 0xd6, 0xaf, 0x40, 0x2e, 0x50, 0x46, 0x87, 0x4f, 0xd8,
	0x03, 0xfe, 0x89, 0x01, 0xff, 0xdd, 0x08, 0xbf, 0xdf, 0x5d, 0xfe, 0xc7, 0x1c, 0xd2, 0xf1, 0xbf,
	0x89, 0x51, 0xca, 0xbd, 0x20, 0xe5, 0x8c, 0xdf, 0xd9, 0x25, 0x45, 0xb6, 0xf1, 0x3c, 0x1b, 0xbe,
	0xee, 0xfc, 0x3c, 0xf3, 0x7e, 0xf8, 0x3c, 0xfc, 0x11, 0x00, 0x00, 0xff, 0xff, 0x17, 0x99, 0x81,
	0xa9, 0x62, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
	GetProfileProgress(ctx context.Context, in *GetProfileProgressRequest, opts ...grpc.CallOption) (*GetProfileProgressResponse, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, "/discussproto.ProfileService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileProgress(ctx context.Context, in *GetProfileProgressRequest, opts ...grpc.CallOption) (*GetProfileProgressResponse, error) {
	out := new(GetProfileProgressResponse)
	err := c.cc.Invoke(ctx, "/discussproto.ProfileService/GetProfileProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	GetProfileProgress(context.Context, *GetProfileProgressRequest) (*GetProfileProgressResponse, error)
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.ProfileService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.ProfileService/GetProfileProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileProgress(ctx, req.(*GetProfileProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discussproto.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfileService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetProfileProgress",
			Handler:    _ProfileService_GetProfileProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
