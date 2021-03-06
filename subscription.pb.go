// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subscription.proto

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

type Subscription struct {
	CId                  int64    `protobuf:"varint,1,opt,name=cId,proto3" json:"cId,omitempty"`
	SeenAt               string   `protobuf:"bytes,2,opt,name=seenAt,proto3" json:"seenAt,omitempty"`
	Unseen               int32    `protobuf:"varint,3,opt,name=unseen,proto3" json:"unseen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{0}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetCId() int64 {
	if m != nil {
		return m.CId
	}
	return 0
}

func (m *Subscription) GetSeenAt() string {
	if m != nil {
		return m.SeenAt
	}
	return ""
}

func (m *Subscription) GetUnseen() int32 {
	if m != nil {
		return m.Unseen
	}
	return 0
}

type SubscriptionCheck struct {
	CId                  int64    `protobuf:"varint,1,opt,name=cId,proto3" json:"cId,omitempty"`
	Subscribed           bool     `protobuf:"varint,2,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionCheck) Reset()         { *m = SubscriptionCheck{} }
func (m *SubscriptionCheck) String() string { return proto.CompactTextString(m) }
func (*SubscriptionCheck) ProtoMessage()    {}
func (*SubscriptionCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{1}
}

func (m *SubscriptionCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionCheck.Unmarshal(m, b)
}
func (m *SubscriptionCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionCheck.Marshal(b, m, deterministic)
}
func (m *SubscriptionCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionCheck.Merge(m, src)
}
func (m *SubscriptionCheck) XXX_Size() int {
	return xxx_messageInfo_SubscriptionCheck.Size(m)
}
func (m *SubscriptionCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionCheck.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionCheck proto.InternalMessageInfo

func (m *SubscriptionCheck) GetCId() int64 {
	if m != nil {
		return m.CId
	}
	return 0
}

func (m *SubscriptionCheck) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

type CreateSubscriptionRequest struct {
	CId                  int64    `protobuf:"varint,1,opt,name=cId,proto3" json:"cId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSubscriptionRequest) Reset()         { *m = CreateSubscriptionRequest{} }
func (m *CreateSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSubscriptionRequest) ProtoMessage()    {}
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{2}
}

func (m *CreateSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSubscriptionRequest.Unmarshal(m, b)
}
func (m *CreateSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *CreateSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSubscriptionRequest.Merge(m, src)
}
func (m *CreateSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSubscriptionRequest.Size(m)
}
func (m *CreateSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSubscriptionRequest proto.InternalMessageInfo

func (m *CreateSubscriptionRequest) GetCId() int64 {
	if m != nil {
		return m.CId
	}
	return 0
}

type CreateSubscriptionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSubscriptionResponse) Reset()         { *m = CreateSubscriptionResponse{} }
func (m *CreateSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSubscriptionResponse) ProtoMessage()    {}
func (*CreateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{3}
}

func (m *CreateSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSubscriptionResponse.Unmarshal(m, b)
}
func (m *CreateSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *CreateSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSubscriptionResponse.Merge(m, src)
}
func (m *CreateSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSubscriptionResponse.Size(m)
}
func (m *CreateSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSubscriptionResponse proto.InternalMessageInfo

type RemoveSubscriptionRequest struct {
	CId                  int64    `protobuf:"varint,1,opt,name=cId,proto3" json:"cId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveSubscriptionRequest) Reset()         { *m = RemoveSubscriptionRequest{} }
func (m *RemoveSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveSubscriptionRequest) ProtoMessage()    {}
func (*RemoveSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{4}
}

func (m *RemoveSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveSubscriptionRequest.Unmarshal(m, b)
}
func (m *RemoveSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *RemoveSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSubscriptionRequest.Merge(m, src)
}
func (m *RemoveSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveSubscriptionRequest.Size(m)
}
func (m *RemoveSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSubscriptionRequest proto.InternalMessageInfo

func (m *RemoveSubscriptionRequest) GetCId() int64 {
	if m != nil {
		return m.CId
	}
	return 0
}

type RemoveSubscriptionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveSubscriptionResponse) Reset()         { *m = RemoveSubscriptionResponse{} }
func (m *RemoveSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveSubscriptionResponse) ProtoMessage()    {}
func (*RemoveSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{5}
}

func (m *RemoveSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveSubscriptionResponse.Unmarshal(m, b)
}
func (m *RemoveSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *RemoveSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSubscriptionResponse.Merge(m, src)
}
func (m *RemoveSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveSubscriptionResponse.Size(m)
}
func (m *RemoveSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSubscriptionResponse proto.InternalMessageInfo

type GetSubscriptionRequest struct {
	CId                  int64    `protobuf:"varint,1,opt,name=cId,proto3" json:"cId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubscriptionRequest) Reset()         { *m = GetSubscriptionRequest{} }
func (m *GetSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionRequest) ProtoMessage()    {}
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{6}
}

func (m *GetSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionRequest.Unmarshal(m, b)
}
func (m *GetSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionRequest.Merge(m, src)
}
func (m *GetSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionRequest.Size(m)
}
func (m *GetSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionRequest proto.InternalMessageInfo

func (m *GetSubscriptionRequest) GetCId() int64 {
	if m != nil {
		return m.CId
	}
	return 0
}

type GetSubscriptionResponse struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetSubscriptionResponse) Reset()         { *m = GetSubscriptionResponse{} }
func (m *GetSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionResponse) ProtoMessage()    {}
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{7}
}

func (m *GetSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionResponse.Unmarshal(m, b)
}
func (m *GetSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionResponse.Merge(m, src)
}
func (m *GetSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionResponse.Size(m)
}
func (m *GetSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionResponse proto.InternalMessageInfo

func (m *GetSubscriptionResponse) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type GetSubscriptionsRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubscriptionsRequest) Reset()         { *m = GetSubscriptionsRequest{} }
func (m *GetSubscriptionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionsRequest) ProtoMessage()    {}
func (*GetSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{8}
}

func (m *GetSubscriptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionsRequest.Unmarshal(m, b)
}
func (m *GetSubscriptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionsRequest.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionsRequest.Merge(m, src)
}
func (m *GetSubscriptionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionsRequest.Size(m)
}
func (m *GetSubscriptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionsRequest proto.InternalMessageInfo

func (m *GetSubscriptionsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetSubscriptionsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetSubscriptionsResponse struct {
	Data                 []*Subscription `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page                 int32           `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Count                int32           `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetSubscriptionsResponse) Reset()         { *m = GetSubscriptionsResponse{} }
func (m *GetSubscriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionsResponse) ProtoMessage()    {}
func (*GetSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{9}
}

func (m *GetSubscriptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubscriptionsResponse.Unmarshal(m, b)
}
func (m *GetSubscriptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubscriptionsResponse.Marshal(b, m, deterministic)
}
func (m *GetSubscriptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubscriptionsResponse.Merge(m, src)
}
func (m *GetSubscriptionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSubscriptionsResponse.Size(m)
}
func (m *GetSubscriptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubscriptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubscriptionsResponse proto.InternalMessageInfo

func (m *GetSubscriptionsResponse) GetData() []*Subscription {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetSubscriptionsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetSubscriptionsResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetSubscriptionsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CheckUserSubscriptionsRequest struct {
	Ids                  string   `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUserSubscriptionsRequest) Reset()         { *m = CheckUserSubscriptionsRequest{} }
func (m *CheckUserSubscriptionsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckUserSubscriptionsRequest) ProtoMessage()    {}
func (*CheckUserSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{10}
}

func (m *CheckUserSubscriptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserSubscriptionsRequest.Unmarshal(m, b)
}
func (m *CheckUserSubscriptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserSubscriptionsRequest.Marshal(b, m, deterministic)
}
func (m *CheckUserSubscriptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserSubscriptionsRequest.Merge(m, src)
}
func (m *CheckUserSubscriptionsRequest) XXX_Size() int {
	return xxx_messageInfo_CheckUserSubscriptionsRequest.Size(m)
}
func (m *CheckUserSubscriptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserSubscriptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserSubscriptionsRequest proto.InternalMessageInfo

func (m *CheckUserSubscriptionsRequest) GetIds() string {
	if m != nil {
		return m.Ids
	}
	return ""
}

type CheckUserSubscriptionsResponse struct {
	Ids                  []int64              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Subscriptions        []*SubscriptionCheck `protobuf:"bytes,2,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CheckUserSubscriptionsResponse) Reset()         { *m = CheckUserSubscriptionsResponse{} }
func (m *CheckUserSubscriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckUserSubscriptionsResponse) ProtoMessage()    {}
func (*CheckUserSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{11}
}

func (m *CheckUserSubscriptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserSubscriptionsResponse.Unmarshal(m, b)
}
func (m *CheckUserSubscriptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserSubscriptionsResponse.Marshal(b, m, deterministic)
}
func (m *CheckUserSubscriptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserSubscriptionsResponse.Merge(m, src)
}
func (m *CheckUserSubscriptionsResponse) XXX_Size() int {
	return xxx_messageInfo_CheckUserSubscriptionsResponse.Size(m)
}
func (m *CheckUserSubscriptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserSubscriptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserSubscriptionsResponse proto.InternalMessageInfo

func (m *CheckUserSubscriptionsResponse) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CheckUserSubscriptionsResponse) GetSubscriptions() []*SubscriptionCheck {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func init() {
	proto.RegisterType((*Subscription)(nil), "discussproto.Subscription")
	proto.RegisterType((*SubscriptionCheck)(nil), "discussproto.SubscriptionCheck")
	proto.RegisterType((*CreateSubscriptionRequest)(nil), "discussproto.CreateSubscriptionRequest")
	proto.RegisterType((*CreateSubscriptionResponse)(nil), "discussproto.CreateSubscriptionResponse")
	proto.RegisterType((*RemoveSubscriptionRequest)(nil), "discussproto.RemoveSubscriptionRequest")
	proto.RegisterType((*RemoveSubscriptionResponse)(nil), "discussproto.RemoveSubscriptionResponse")
	proto.RegisterType((*GetSubscriptionRequest)(nil), "discussproto.GetSubscriptionRequest")
	proto.RegisterType((*GetSubscriptionResponse)(nil), "discussproto.GetSubscriptionResponse")
	proto.RegisterType((*GetSubscriptionsRequest)(nil), "discussproto.GetSubscriptionsRequest")
	proto.RegisterType((*GetSubscriptionsResponse)(nil), "discussproto.GetSubscriptionsResponse")
	proto.RegisterType((*CheckUserSubscriptionsRequest)(nil), "discussproto.CheckUserSubscriptionsRequest")
	proto.RegisterType((*CheckUserSubscriptionsResponse)(nil), "discussproto.CheckUserSubscriptionsResponse")
}

func init() { proto.RegisterFile("subscription.proto", fileDescriptor_c4f8ad1a64b2bad6) }

var fileDescriptor_c4f8ad1a64b2bad6 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0xa4, 0xa2, 0x43, 0x80, 0x74, 0x8a, 0x52, 0xc7, 0x0a, 0x6d, 0x64, 0xf1, 0x13,
	0x05, 0x12, 0x8b, 0x22, 0x71, 0xe0, 0x80, 0x04, 0x51, 0x85, 0x7a, 0x43, 0xae, 0x38, 0xc0, 0xcd,
	0xb1, 0x57, 0xe9, 0xaa, 0x89, 0x37, 0x78, 0xd7, 0xa9, 0x10, 0x20, 0x24, 0x4e, 0x08, 0x8e, 0x88,
	0x47, 0xe2, 0x09, 0x78, 0x05, 0x1e, 0xa4, 0xf2, 0xda, 0x8d, 0xd6, 0x7f, 0x49, 0x4e, 0x99, 0xd9,
	0x7c, 0xf3, 0xcd, 0x37, 0x9e, 0x6f, 0x6d, 0x40, 0x1e, 0x4d, 0xb8, 0x17, 0xd2, 0x85, 0xa0, 0x2c,
	0x18, 0x2d, 0x42, 0x26, 0x18, 0x36, 0x7d, 0xca, 0xbd, 0x88, 0x73, 0x99, 0x99, 0xdd, 0x29, 0x63,
	0xd3, 0x19, 0xb1, 0xdd, 0x05, 0xb5, 0xdd, 0x20, 0x60, 0xc2, 0x8d, 0xa1, 0x3c, 0xc1, 0x9a, 0xcf,
	0xa7, 0x54, 0x9c, 0x47, 0x93, 0x91, 0xc7, 0xe6, 0xf6, 0xfc, 0x92, 0x8a, 0x0b, 0x76, 0x69, 0x4f,
	0xd9, 0x50, 0xfe, 0x39, 0x5c, 0xba, 0x33, 0xea, 0xbb, 0x82, 0x85, 0xdc, 0x5e, 0x85, 0x49, 0x9d,
	0xf5, 0x16, 0x9a, 0x67, 0x4a, 0x67, 0x6c, 0x81, 0xee, 0x9d, 0xfa, 0x86, 0xd6, 0xd3, 0xfa, 0xba,
	0x13, 0x87, 0xd8, 0x86, 0x1d, 0x4e, 0x48, 0xf0, 0x4a, 0x18, 0xb5, 0x9e, 0xd6, 0xdf, 0x75, 0xd2,
	0x2c, 0x3e, 0x8f, 0x82, 0x38, 0x36, 0xf4, 0x9e, 0xd6, 0x6f, 0x38, 0x69, 0x66, 0x9d, 0xc0, 0x9e,
	0xca, 0x38, 0x3e, 0x27, 0xde, 0x45, 0x09, 0xed, 0x21, 0x40, 0x3a, 0xf2, 0x84, 0xf8, 0x92, 0xfa,
	0x86, 0xa3, 0x9c, 0x58, 0x43, 0xe8, 0x8c, 0x43, 0xe2, 0x0a, 0xa2, 0x92, 0x39, 0xe4, 0x63, 0x44,
	0xb8, 0x28, 0xd2, 0x59, 0x5d, 0x30, 0xcb, 0xe0, 0x7c, 0xc1, 0x02, 0x4e, 0x62, 0x32, 0x87, 0xcc,
	0xd9, 0x72, 0x7b, 0xb2, 0x32, 0x78, 0x4a, 0x36, 0x80, 0xf6, 0x1b, 0x22, 0xb6, 0x63, 0x7a, 0x0f,
	0x07, 0x05, 0x6c, 0x42, 0x83, 0x2f, 0xa1, 0xa9, 0xee, 0x5c, 0x56, 0xdd, 0x3c, 0x36, 0x47, 0xea,
	0xd2, 0x47, 0x99, 0xca, 0x0c, 0xde, 0x1a, 0x17, 0xa8, 0xf9, 0xb5, 0x0e, 0x84, 0xfa, 0xc2, 0x9d,
	0x12, 0x49, 0xd9, 0x70, 0x64, 0x8c, 0x77, 0xa1, 0x31, 0xa3, 0x73, 0x9a, 0x6c, 0xb1, 0xe1, 0x24,
	0x89, 0xf5, 0x53, 0x03, 0xa3, 0xc8, 0x92, 0x2a, 0x1c, 0x41, 0xdd, 0x77, 0x85, 0x6b, 0x68, 0x3d,
	0x7d, 0x83, 0x32, 0x89, 0x5b, 0xb5, 0xad, 0x95, 0xb5, 0xd5, 0x95, 0xb6, 0xf1, 0xa9, 0xc7, 0xa2,
	0x40, 0x18, 0xf5, 0xe4, 0x54, 0x26, 0xd6, 0x53, 0xb8, 0x27, 0xdd, 0xf2, 0x8e, 0x93, 0xb0, 0x74,
	0xae, 0x16, 0xe8, 0xd4, 0xe7, 0x72, 0xac, 0x5d, 0x27, 0x0e, 0xad, 0x4f, 0x70, 0x58, 0x55, 0x92,
	0x0e, 0xb1, 0xaa, 0xd1, 0xe3, 0x9d, 0x50, 0x9f, 0xe3, 0x09, 0xdc, 0x52, 0x1f, 0x24, 0x37, 0x6a,
	0x72, 0xbe, 0xa3, 0xea, 0xf9, 0x64, 0x0b, 0x27, 0x5b, 0x75, 0xfc, 0xb7, 0x01, 0xfb, 0x2a, 0xe8,
	0x8c, 0x84, 0x4b, 0xea, 0x11, 0xfc, 0xa5, 0x01, 0x16, 0xad, 0x88, 0x8f, 0xb2, 0xf4, 0x95, 0xde,
	0x36, 0xfb, 0x9b, 0x81, 0xa9, 0x11, 0xad, 0xef, 0xff, 0xfe, 0xff, 0xae, 0x75, 0xad, 0x03, 0xf9,
	0x4e, 0x50, 0xd5, 0xd9, 0x9f, 0xbd, 0x53, 0xff, 0xeb, 0x0b, 0x6d, 0x80, 0x3f, 0x34, 0xc0, 0xa2,
	0x97, 0xf3, 0x6a, 0x2a, 0x2f, 0x47, 0x5e, 0xcd, 0x9a, 0x6b, 0x71, 0x24, 0xd5, 0x74, 0x06, 0x55,
	0x6a, 0xf0, 0x1b, 0xdc, 0xc9, 0x59, 0x0d, 0xef, 0x67, 0xd9, 0xcb, 0xaf, 0x95, 0xf9, 0x60, 0x03,
	0x2a, 0x2b, 0x00, 0x2b, 0x05, 0x7c, 0x81, 0x56, 0xde, 0xeb, 0xb8, 0x9e, 0xfb, 0xda, 0x79, 0xe6,
	0xc3, 0x4d, 0xb0, 0x54, 0x43, 0x47, 0x6a, 0xd8, 0xc7, 0xbd, 0x82, 0x06, 0xfc, 0xa3, 0x41, 0xbb,
	0xdc, 0xab, 0xf8, 0x38, 0xb7, 0xf2, 0x75, 0x97, 0xc0, 0x7c, 0xb2, 0x1d, 0x38, 0x15, 0xd4, 0x93,
	0x82, 0x4c, 0x34, 0x0a, 0x82, 0xb8, 0xed, 0xc5, 0xa5, 0xaf, 0x6f, 0x7f, 0xc8, 0x7c, 0x67, 0x26,
	0x3b, 0xf2, 0xe7, 0xd9, 0x55, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xb3, 0x6f, 0xd9, 0x92, 0x06,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error)
	RemoveSubscription(ctx context.Context, in *RemoveSubscriptionRequest, opts ...grpc.CallOption) (*RemoveSubscriptionResponse, error)
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error)
	CheckUserSubscriptions(ctx context.Context, in *CheckUserSubscriptionsRequest, opts ...grpc.CallOption) (*CheckUserSubscriptionsResponse, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error) {
	out := new(CreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/discussproto.SubscriptionService/CreateSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) RemoveSubscription(ctx context.Context, in *RemoveSubscriptionRequest, opts ...grpc.CallOption) (*RemoveSubscriptionResponse, error) {
	out := new(RemoveSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/discussproto.SubscriptionService/RemoveSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/discussproto.SubscriptionService/GetSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error) {
	out := new(GetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/discussproto.SubscriptionService/GetSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CheckUserSubscriptions(ctx context.Context, in *CheckUserSubscriptionsRequest, opts ...grpc.CallOption) (*CheckUserSubscriptionsResponse, error) {
	out := new(CheckUserSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/discussproto.SubscriptionService/CheckUserSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
type SubscriptionServiceServer interface {
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error)
	RemoveSubscription(context.Context, *RemoveSubscriptionRequest) (*RemoveSubscriptionResponse, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error)
	CheckUserSubscriptions(context.Context, *CheckUserSubscriptionsRequest) (*CheckUserSubscriptionsResponse, error)
}

// UnimplementedSubscriptionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (*UnimplementedSubscriptionServiceServer) CreateSubscription(ctx context.Context, req *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) RemoveSubscription(ctx context.Context, req *RemoveSubscriptionRequest) (*RemoveSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) GetSubscription(ctx context.Context, req *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) GetSubscriptions(ctx context.Context, req *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (*UnimplementedSubscriptionServiceServer) CheckUserSubscriptions(ctx context.Context, req *CheckUserSubscriptionsRequest) (*CheckUserSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserSubscriptions not implemented")
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.SubscriptionService/CreateSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_RemoveSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).RemoveSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.SubscriptionService/RemoveSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).RemoveSubscription(ctx, req.(*RemoveSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.SubscriptionService/GetSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.SubscriptionService/GetSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, req.(*GetSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CheckUserSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CheckUserSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.SubscriptionService/CheckUserSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CheckUserSubscriptions(ctx, req.(*CheckUserSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discussproto.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscription",
			Handler:    _SubscriptionService_CreateSubscription_Handler,
		},
		{
			MethodName: "RemoveSubscription",
			Handler:    _SubscriptionService_RemoveSubscription_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _SubscriptionService_GetSubscription_Handler,
		},
		{
			MethodName: "GetSubscriptions",
			Handler:    _SubscriptionService_GetSubscriptions_Handler,
		},
		{
			MethodName: "CheckUserSubscriptions",
			Handler:    _SubscriptionService_CheckUserSubscriptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
