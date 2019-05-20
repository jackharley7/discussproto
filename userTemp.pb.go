// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userTemp.proto

package proto

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

type UserTemp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	VerifyNumber         int32    `protobuf:"varint,5,opt,name=VerifyNumber,proto3" json:"VerifyNumber,omitempty"`
	Verified             bool     `protobuf:"varint,6,opt,name=Verified,proto3" json:"Verified,omitempty"`
	TimeZone             string   `protobuf:"bytes,7,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTemp) Reset()         { *m = UserTemp{} }
func (m *UserTemp) String() string { return proto.CompactTextString(m) }
func (*UserTemp) ProtoMessage()    {}
func (*UserTemp) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{0}
}

func (m *UserTemp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTemp.Unmarshal(m, b)
}
func (m *UserTemp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTemp.Marshal(b, m, deterministic)
}
func (m *UserTemp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTemp.Merge(m, src)
}
func (m *UserTemp) XXX_Size() int {
	return xxx_messageInfo_UserTemp.Size(m)
}
func (m *UserTemp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTemp.DiscardUnknown(m)
}

var xxx_messageInfo_UserTemp proto.InternalMessageInfo

func (m *UserTemp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserTemp) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserTemp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserTemp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserTemp) GetVerifyNumber() int32 {
	if m != nil {
		return m.VerifyNumber
	}
	return 0
}

func (m *UserTemp) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *UserTemp) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

type CreateUserTempRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TimeZone             string   `protobuf:"bytes,3,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserTempRequest) Reset()         { *m = CreateUserTempRequest{} }
func (m *CreateUserTempRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserTempRequest) ProtoMessage()    {}
func (*CreateUserTempRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{1}
}

func (m *CreateUserTempRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserTempRequest.Unmarshal(m, b)
}
func (m *CreateUserTempRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserTempRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserTempRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserTempRequest.Merge(m, src)
}
func (m *CreateUserTempRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserTempRequest.Size(m)
}
func (m *CreateUserTempRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserTempRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserTempRequest proto.InternalMessageInfo

func (m *CreateUserTempRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserTempRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserTempRequest) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

type CreateUserTempResponse struct {
	UserTemp             *UserTemp `protobuf:"bytes,1,opt,name=userTemp,proto3" json:"userTemp,omitempty"`
	Token                string    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateUserTempResponse) Reset()         { *m = CreateUserTempResponse{} }
func (m *CreateUserTempResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserTempResponse) ProtoMessage()    {}
func (*CreateUserTempResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{2}
}

func (m *CreateUserTempResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserTempResponse.Unmarshal(m, b)
}
func (m *CreateUserTempResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserTempResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserTempResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserTempResponse.Merge(m, src)
}
func (m *CreateUserTempResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserTempResponse.Size(m)
}
func (m *CreateUserTempResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserTempResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserTempResponse proto.InternalMessageInfo

func (m *CreateUserTempResponse) GetUserTemp() *UserTemp {
	if m != nil {
		return m.UserTemp
	}
	return nil
}

func (m *CreateUserTempResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyUserTempRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	VerifyNumber         int32    `protobuf:"varint,2,opt,name=verifyNumber,proto3" json:"verifyNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserTempRequest) Reset()         { *m = VerifyUserTempRequest{} }
func (m *VerifyUserTempRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyUserTempRequest) ProtoMessage()    {}
func (*VerifyUserTempRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{3}
}

func (m *VerifyUserTempRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserTempRequest.Unmarshal(m, b)
}
func (m *VerifyUserTempRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserTempRequest.Marshal(b, m, deterministic)
}
func (m *VerifyUserTempRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserTempRequest.Merge(m, src)
}
func (m *VerifyUserTempRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyUserTempRequest.Size(m)
}
func (m *VerifyUserTempRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserTempRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserTempRequest proto.InternalMessageInfo

func (m *VerifyUserTempRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyUserTempRequest) GetVerifyNumber() int32 {
	if m != nil {
		return m.VerifyNumber
	}
	return 0
}

type VerifyUserTempResponse struct {
	UserTemp             *UserTemp `protobuf:"bytes,1,opt,name=userTemp,proto3" json:"userTemp,omitempty"`
	Token                string    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Verified             bool      `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VerifyUserTempResponse) Reset()         { *m = VerifyUserTempResponse{} }
func (m *VerifyUserTempResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyUserTempResponse) ProtoMessage()    {}
func (*VerifyUserTempResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{4}
}

func (m *VerifyUserTempResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserTempResponse.Unmarshal(m, b)
}
func (m *VerifyUserTempResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserTempResponse.Marshal(b, m, deterministic)
}
func (m *VerifyUserTempResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserTempResponse.Merge(m, src)
}
func (m *VerifyUserTempResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyUserTempResponse.Size(m)
}
func (m *VerifyUserTempResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserTempResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserTempResponse proto.InternalMessageInfo

func (m *VerifyUserTempResponse) GetUserTemp() *UserTemp {
	if m != nil {
		return m.UserTemp
	}
	return nil
}

func (m *VerifyUserTempResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyUserTempResponse) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

type GetTwitterRequestTokenRequest struct {
	CallbackUrl          string   `protobuf:"bytes,1,opt,name=callbackUrl,proto3" json:"callbackUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTwitterRequestTokenRequest) Reset()         { *m = GetTwitterRequestTokenRequest{} }
func (m *GetTwitterRequestTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTwitterRequestTokenRequest) ProtoMessage()    {}
func (*GetTwitterRequestTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{5}
}

func (m *GetTwitterRequestTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTwitterRequestTokenRequest.Unmarshal(m, b)
}
func (m *GetTwitterRequestTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTwitterRequestTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTwitterRequestTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTwitterRequestTokenRequest.Merge(m, src)
}
func (m *GetTwitterRequestTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTwitterRequestTokenRequest.Size(m)
}
func (m *GetTwitterRequestTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTwitterRequestTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTwitterRequestTokenRequest proto.InternalMessageInfo

func (m *GetTwitterRequestTokenRequest) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

type GetTwitterRequestTokenResponse struct {
	OauthToken             string   `protobuf:"bytes,1,opt,name=oauth_token,json=oauthToken,proto3" json:"oauth_token,omitempty"`
	OauthTokenSecret       string   `protobuf:"bytes,2,opt,name=oauth_token_secret,json=oauthTokenSecret,proto3" json:"oauth_token_secret,omitempty"`
	OauthCallbackConfirmed bool     `protobuf:"varint,3,opt,name=oauth_callback_confirmed,json=oauthCallbackConfirmed,proto3" json:"oauth_callback_confirmed,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetTwitterRequestTokenResponse) Reset()         { *m = GetTwitterRequestTokenResponse{} }
func (m *GetTwitterRequestTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTwitterRequestTokenResponse) ProtoMessage()    {}
func (*GetTwitterRequestTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{6}
}

func (m *GetTwitterRequestTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTwitterRequestTokenResponse.Unmarshal(m, b)
}
func (m *GetTwitterRequestTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTwitterRequestTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetTwitterRequestTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTwitterRequestTokenResponse.Merge(m, src)
}
func (m *GetTwitterRequestTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTwitterRequestTokenResponse.Size(m)
}
func (m *GetTwitterRequestTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTwitterRequestTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTwitterRequestTokenResponse proto.InternalMessageInfo

func (m *GetTwitterRequestTokenResponse) GetOauthToken() string {
	if m != nil {
		return m.OauthToken
	}
	return ""
}

func (m *GetTwitterRequestTokenResponse) GetOauthTokenSecret() string {
	if m != nil {
		return m.OauthTokenSecret
	}
	return ""
}

func (m *GetTwitterRequestTokenResponse) GetOauthCallbackConfirmed() bool {
	if m != nil {
		return m.OauthCallbackConfirmed
	}
	return false
}

type TwitterSignupRequest struct {
	OauthToken           string   `protobuf:"bytes,1,opt,name=oauth_token,json=oauthToken,proto3" json:"oauth_token,omitempty"`
	OauthVerifier        string   `protobuf:"bytes,2,opt,name=oauth_verifier,json=oauthVerifier,proto3" json:"oauth_verifier,omitempty"`
	TimeZone             string   `protobuf:"bytes,3,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwitterSignupRequest) Reset()         { *m = TwitterSignupRequest{} }
func (m *TwitterSignupRequest) String() string { return proto.CompactTextString(m) }
func (*TwitterSignupRequest) ProtoMessage()    {}
func (*TwitterSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{7}
}

func (m *TwitterSignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwitterSignupRequest.Unmarshal(m, b)
}
func (m *TwitterSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwitterSignupRequest.Marshal(b, m, deterministic)
}
func (m *TwitterSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwitterSignupRequest.Merge(m, src)
}
func (m *TwitterSignupRequest) XXX_Size() int {
	return xxx_messageInfo_TwitterSignupRequest.Size(m)
}
func (m *TwitterSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TwitterSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TwitterSignupRequest proto.InternalMessageInfo

func (m *TwitterSignupRequest) GetOauthToken() string {
	if m != nil {
		return m.OauthToken
	}
	return ""
}

func (m *TwitterSignupRequest) GetOauthVerifier() string {
	if m != nil {
		return m.OauthVerifier
	}
	return ""
}

func (m *TwitterSignupRequest) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

type TwitterSignupResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwitterSignupResponse) Reset()         { *m = TwitterSignupResponse{} }
func (m *TwitterSignupResponse) String() string { return proto.CompactTextString(m) }
func (*TwitterSignupResponse) ProtoMessage()    {}
func (*TwitterSignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684109c25ee20132, []int{8}
}

func (m *TwitterSignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwitterSignupResponse.Unmarshal(m, b)
}
func (m *TwitterSignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwitterSignupResponse.Marshal(b, m, deterministic)
}
func (m *TwitterSignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwitterSignupResponse.Merge(m, src)
}
func (m *TwitterSignupResponse) XXX_Size() int {
	return xxx_messageInfo_TwitterSignupResponse.Size(m)
}
func (m *TwitterSignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TwitterSignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TwitterSignupResponse proto.InternalMessageInfo

func (m *TwitterSignupResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *TwitterSignupResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*UserTemp)(nil), "proto.UserTemp")
	proto.RegisterType((*CreateUserTempRequest)(nil), "proto.CreateUserTempRequest")
	proto.RegisterType((*CreateUserTempResponse)(nil), "proto.CreateUserTempResponse")
	proto.RegisterType((*VerifyUserTempRequest)(nil), "proto.VerifyUserTempRequest")
	proto.RegisterType((*VerifyUserTempResponse)(nil), "proto.VerifyUserTempResponse")
	proto.RegisterType((*GetTwitterRequestTokenRequest)(nil), "proto.GetTwitterRequestTokenRequest")
	proto.RegisterType((*GetTwitterRequestTokenResponse)(nil), "proto.GetTwitterRequestTokenResponse")
	proto.RegisterType((*TwitterSignupRequest)(nil), "proto.TwitterSignupRequest")
	proto.RegisterType((*TwitterSignupResponse)(nil), "proto.TwitterSignupResponse")
}

func init() { proto.RegisterFile("userTemp.proto", fileDescriptor_684109c25ee20132) }

var fileDescriptor_684109c25ee20132 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xfe, 0x9d, 0x4b, 0x9b, 0x9c, 0xb4, 0x69, 0x35, 0x6a, 0x22, 0xff, 0x6e, 0x4a, 0xa3, 0x01,
	0x44, 0x54, 0x68, 0x2d, 0x15, 0x09, 0xa1, 0xee, 0x68, 0x85, 0x10, 0x9b, 0x2e, 0xdc, 0xb4, 0x42,
	0x65, 0x11, 0x39, 0xce, 0x34, 0x1d, 0x35, 0xf6, 0x84, 0xf1, 0x24, 0x51, 0x59, 0xb2, 0x84, 0x25,
	0x4f, 0xc1, 0x4b, 0xf0, 0x12, 0x3c, 0x00, 0x12, 0x82, 0xf7, 0x40, 0x73, 0xb1, 0x1b, 0x07, 0x17,
	0x84, 0xc4, 0xca, 0x3e, 0xe7, 0x3b, 0x3e, 0xdf, 0x77, 0x6e, 0x09, 0xd4, 0x27, 0x31, 0xe1, 0x5d,
	0x12, 0x8e, 0xf7, 0xc6, 0x9c, 0x09, 0x86, 0xca, 0xea, 0xe1, 0x80, 0x74, 0x6b, 0x97, 0xd3, 0x1a,
	0x32, 0x36, 0x1c, 0x11, 0xd7, 0x1f, 0x53, 0xd7, 0x8f, 0x22, 0x26, 0x7c, 0x41, 0x59, 0x14, 0x1b,
	0xf4, 0xc9, 0x90, 0x8a, 0xcb, 0x49, 0x7f, 0x2f, 0x60, 0xa1, 0x1b, 0xce, 0xa8, 0xb8, 0x62, 0x33,
	0x77, 0xc8, 0x76, 0x15, 0xb8, 0x3b, 0xf5, 0x47, 0x74, 0xe0, 0x0b, 0xc6, 0x63, 0x37, 0x7d, 0xd5,
	0xdf, 0xe1, 0xcf, 0x16, 0x54, 0x4e, 0x0d, 0x37, 0xaa, 0x43, 0x81, 0x0e, 0x6c, 0xab, 0x6d, 0x75,
	0x8a, 0x5e, 0x81, 0x0e, 0x50, 0x0b, 0xaa, 0x01, 0x27, 0xbe, 0x20, 0x83, 0x67, 0xc2, 0x2e, 0xb4,
	0xad, 0x4e, 0xd5, 0xbb, 0x71, 0x20, 0x04, 0xa5, 0x63, 0x3f, 0x24, 0x76, 0x51, 0x01, 0xea, 0x1d,
	0x6d, 0x40, 0xf9, 0x79, 0xe8, 0xd3, 0x91, 0x5d, 0x52, 0x4e, 0x6d, 0x20, 0x0c, 0x2b, 0x67, 0x84,
	0xd3, 0x8b, 0xeb, 0xe3, 0x49, 0xd8, 0x27, 0xdc, 0x2e, 0xb7, 0xad, 0x4e, 0xd9, 0xcb, 0xf8, 0x90,
	0x03, 0x15, 0x65, 0x53, 0x32, 0xb0, 0x97, 0xda, 0x56, 0xa7, 0xe2, 0xa5, 0xb6, 0xc4, 0xba, 0x34,
	0x24, 0xe7, 0x2c, 0x22, 0xf6, 0xb2, 0x4a, 0x9c, 0xda, 0x38, 0x84, 0xc6, 0x91, 0x92, 0x94, 0x54,
	0xe1, 0x91, 0x37, 0x13, 0x12, 0x0b, 0xd4, 0x82, 0x32, 0x51, 0x52, 0x64, 0x3d, 0xd5, 0xc3, 0xa5,
	0x6f, 0x5f, 0xb7, 0x0b, 0xaf, 0x2c, 0x4f, 0x3b, 0x91, 0x03, 0xa5, 0x48, 0x8a, 0x2f, 0x64, 0x40,
	0xe5, 0x93, 0x74, 0x22, 0xa1, 0xd3, 0xc5, 0xa5, 0x36, 0x7e, 0x0d, 0xcd, 0x45, 0xba, 0x78, 0xcc,
	0xa2, 0x98, 0xa0, 0x87, 0x50, 0x49, 0x86, 0xa8, 0x28, 0x6b, 0xfb, 0x6b, 0xba, 0xc7, 0x7b, 0x69,
	0x68, 0x1a, 0x20, 0xfb, 0x24, 0xd8, 0x15, 0x89, 0x4c, 0x57, 0xb5, 0x81, 0x7d, 0x68, 0xe8, 0x9e,
	0xfc, 0x5d, 0x2d, 0x3b, 0xb0, 0x32, 0x9d, 0x6f, 0xaf, 0xcc, 0x59, 0xd6, 0x41, 0xeb, 0xff, 0x79,
	0x19, 0x0c, 0xcf, 0xa0, 0xb9, 0x48, 0xf1, 0xcf, 0xf4, 0xcb, 0xc6, 0x4d, 0x93, 0x19, 0x16, 0xf5,
	0x0c, 0x13, 0x1b, 0xbf, 0x84, 0xad, 0x17, 0x44, 0x74, 0x67, 0x54, 0x08, 0xc2, 0x4d, 0x5d, 0x5d,
	0xf9, 0x55, 0x52, 0x63, 0x07, 0x6a, 0x81, 0x3f, 0x1a, 0xf5, 0xfd, 0xe0, 0xea, 0x94, 0x2f, 0x56,
	0x3a, 0x0f, 0xe1, 0x4f, 0x16, 0xdc, 0xb9, 0x2d, 0x97, 0x29, 0x66, 0x1b, 0x6a, 0xcc, 0x9f, 0x88,
	0xcb, 0x9e, 0x56, 0xa9, 0x92, 0x79, 0xa0, 0x5c, 0x2a, 0x10, 0x3d, 0x02, 0x34, 0x17, 0xd0, 0x8b,
	0x49, 0xc0, 0x49, 0xb2, 0xe3, 0xeb, 0x37, 0x71, 0x27, 0xca, 0x8f, 0x9e, 0x82, 0xad, 0xa3, 0x13,
	0x19, 0xbd, 0x80, 0x45, 0x17, 0x94, 0x87, 0x69, 0xa1, 0x4d, 0x85, 0x1f, 0x19, 0xf8, 0x28, 0x41,
	0xf1, 0x07, 0x0b, 0x36, 0x8c, 0xd0, 0x13, 0x3a, 0x8c, 0x26, 0xe9, 0x48, 0x1f, 0xe4, 0x28, 0x4c,
	0xcb, 0x9d, 0x57, 0xba, 0x0b, 0x75, 0x1d, 0x68, 0x5a, 0xc9, 0x17, 0x76, 0x76, 0x55, 0xa1, 0xe6,
	0x56, 0x38, 0xda, 0x84, 0xaa, 0x5c, 0xd6, 0xde, 0xdb, 0xbc, 0xed, 0x3d, 0x86, 0xc6, 0x82, 0x98,
	0xb4, 0x5f, 0x25, 0x39, 0x5b, 0x33, 0xf8, 0xda, 0xdc, 0xe0, 0x3d, 0x05, 0xe4, 0x0f, 0x7c, 0xff,
	0x47, 0x11, 0xd6, 0x92, 0xed, 0x38, 0x21, 0x7c, 0x4a, 0x03, 0x82, 0x86, 0x50, 0xcf, 0x5e, 0x08,
	0x6a, 0x99, 0x74, 0xb9, 0x77, 0xea, 0x6c, 0xdd, 0x82, 0x6a, 0x65, 0xd8, 0x7e, 0xf7, 0xe5, 0xfb,
	0xc7, 0x02, 0xc2, 0xab, 0xea, 0x87, 0x6f, 0xec, 0xc6, 0x4a, 0xf8, 0x81, 0xb5, 0x83, 0x22, 0xa8,
	0x9f, 0x11, 0x7e, 0x41, 0xaf, 0x7f, 0x21, 0xca, 0x3d, 0xa2, 0x94, 0x28, 0x7f, 0xff, 0xf1, 0xb6,
	0x22, 0xfa, 0x1f, 0x6f, 0x64, 0x88, 0x5c, 0x7d, 0x3d, 0x92, 0xef, 0xbd, 0x05, 0xcd, 0xfc, 0xb5,
	0x43, 0xf7, 0x4c, 0xea, 0xdf, 0x6e, 0xb8, 0x73, 0xff, 0x0f, 0x51, 0x46, 0xc8, 0x5d, 0x25, 0x64,
	0x0b, 0x6d, 0x66, 0x85, 0x08, 0xfd, 0x89, 0xab, 0x4f, 0x6d, 0x04, 0xab, 0x99, 0x49, 0xa2, 0x4d,
	0x93, 0x3c, 0x6f, 0xd9, 0x9c, 0x56, 0x3e, 0x68, 0x08, 0xdb, 0x8a, 0xd0, 0xc1, 0x8d, 0x5c, 0xc2,
	0x03, 0x6b, 0xe7, 0x70, 0xf9, 0x5c, 0xff, 0x21, 0xf5, 0x97, 0xd4, 0xe3, 0xf1, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x35, 0xf9, 0xd7, 0xe5, 0xb0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserTempServiceClient is the client API for UserTempService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserTempServiceClient interface {
	CreateUserTemp(ctx context.Context, in *CreateUserTempRequest, opts ...grpc.CallOption) (*CreateUserTempResponse, error)
	VerfiyUserTemp(ctx context.Context, in *VerifyUserTempRequest, opts ...grpc.CallOption) (*VerifyUserTempResponse, error)
	GetTwitterRequestToken(ctx context.Context, in *GetTwitterRequestTokenRequest, opts ...grpc.CallOption) (*GetTwitterRequestTokenResponse, error)
	TwitterSignup(ctx context.Context, in *TwitterSignupRequest, opts ...grpc.CallOption) (*TwitterSignupResponse, error)
}

type userTempServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserTempServiceClient(cc *grpc.ClientConn) UserTempServiceClient {
	return &userTempServiceClient{cc}
}

func (c *userTempServiceClient) CreateUserTemp(ctx context.Context, in *CreateUserTempRequest, opts ...grpc.CallOption) (*CreateUserTempResponse, error) {
	out := new(CreateUserTempResponse)
	err := c.cc.Invoke(ctx, "/proto.UserTempService/CreateUserTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTempServiceClient) VerfiyUserTemp(ctx context.Context, in *VerifyUserTempRequest, opts ...grpc.CallOption) (*VerifyUserTempResponse, error) {
	out := new(VerifyUserTempResponse)
	err := c.cc.Invoke(ctx, "/proto.UserTempService/VerfiyUserTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTempServiceClient) GetTwitterRequestToken(ctx context.Context, in *GetTwitterRequestTokenRequest, opts ...grpc.CallOption) (*GetTwitterRequestTokenResponse, error) {
	out := new(GetTwitterRequestTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.UserTempService/GetTwitterRequestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTempServiceClient) TwitterSignup(ctx context.Context, in *TwitterSignupRequest, opts ...grpc.CallOption) (*TwitterSignupResponse, error) {
	out := new(TwitterSignupResponse)
	err := c.cc.Invoke(ctx, "/proto.UserTempService/TwitterSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTempServiceServer is the server API for UserTempService service.
type UserTempServiceServer interface {
	CreateUserTemp(context.Context, *CreateUserTempRequest) (*CreateUserTempResponse, error)
	VerfiyUserTemp(context.Context, *VerifyUserTempRequest) (*VerifyUserTempResponse, error)
	GetTwitterRequestToken(context.Context, *GetTwitterRequestTokenRequest) (*GetTwitterRequestTokenResponse, error)
	TwitterSignup(context.Context, *TwitterSignupRequest) (*TwitterSignupResponse, error)
}

func RegisterUserTempServiceServer(s *grpc.Server, srv UserTempServiceServer) {
	s.RegisterService(&_UserTempService_serviceDesc, srv)
}

func _UserTempService_CreateUserTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserTempRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTempServiceServer).CreateUserTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserTempService/CreateUserTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTempServiceServer).CreateUserTemp(ctx, req.(*CreateUserTempRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTempService_VerfiyUserTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserTempRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTempServiceServer).VerfiyUserTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserTempService/VerfiyUserTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTempServiceServer).VerfiyUserTemp(ctx, req.(*VerifyUserTempRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTempService_GetTwitterRequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTwitterRequestTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTempServiceServer).GetTwitterRequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserTempService/GetTwitterRequestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTempServiceServer).GetTwitterRequestToken(ctx, req.(*GetTwitterRequestTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTempService_TwitterSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitterSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTempServiceServer).TwitterSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserTempService/TwitterSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTempServiceServer).TwitterSignup(ctx, req.(*TwitterSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserTempService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserTempService",
	HandlerType: (*UserTempServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserTemp",
			Handler:    _UserTempService_CreateUserTemp_Handler,
		},
		{
			MethodName: "VerfiyUserTemp",
			Handler:    _UserTempService_VerfiyUserTemp_Handler,
		},
		{
			MethodName: "GetTwitterRequestToken",
			Handler:    _UserTempService_GetTwitterRequestToken_Handler,
		},
		{
			MethodName: "TwitterSignup",
			Handler:    _UserTempService_TwitterSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userTemp.proto",
}
