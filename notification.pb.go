// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

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

type Type int32

const (
	Type_UNKNOWN        Type = 0
	Type_INVITE         Type = 1
	Type_ACCEPTED       Type = 2
	Type_DECLINED       Type = 3
	Type_RESPONSE       Type = 4
	Type_RESPONSEVIEWED Type = 5
)

var Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "INVITE",
	2: "ACCEPTED",
	3: "DECLINED",
	4: "RESPONSE",
	5: "RESPONSEVIEWED",
}

var Type_value = map[string]int32{
	"UNKNOWN":        0,
	"INVITE":         1,
	"ACCEPTED":       2,
	"DECLINED":       3,
	"RESPONSE":       4,
	"RESPONSEVIEWED": 5,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

type UserObj struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserObj) Reset()         { *m = UserObj{} }
func (m *UserObj) String() string { return proto.CompactTextString(m) }
func (*UserObj) ProtoMessage()    {}
func (*UserObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *UserObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserObj.Unmarshal(m, b)
}
func (m *UserObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserObj.Marshal(b, m, deterministic)
}
func (m *UserObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserObj.Merge(m, src)
}
func (m *UserObj) XXX_Size() int {
	return xxx_messageInfo_UserObj.Size(m)
}
func (m *UserObj) XXX_DiscardUnknown() {
	xxx_messageInfo_UserObj.DiscardUnknown(m)
}

var xxx_messageInfo_UserObj proto.InternalMessageInfo

func (m *UserObj) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserObj) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserObj) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type NotificationData struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	User                 *UserObj `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Meta                 string   `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationData) Reset()         { *m = NotificationData{} }
func (m *NotificationData) String() string { return proto.CompactTextString(m) }
func (*NotificationData) ProtoMessage()    {}
func (*NotificationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *NotificationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationData.Unmarshal(m, b)
}
func (m *NotificationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationData.Marshal(b, m, deterministic)
}
func (m *NotificationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationData.Merge(m, src)
}
func (m *NotificationData) XXX_Size() int {
	return xxx_messageInfo_NotificationData.Size(m)
}
func (m *NotificationData) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationData.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationData proto.InternalMessageInfo

func (m *NotificationData) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *NotificationData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotificationData) GetUser() *UserObj {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *NotificationData) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type Notification struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubjectId            int64             `protobuf:"varint,2,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
	UserId               int64             `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt            string            `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Type                 Type              `protobuf:"varint,5,opt,name=type,proto3,enum=discussproto.Type" json:"type,omitempty"`
	Viewed               bool              `protobuf:"varint,6,opt,name=viewed,proto3" json:"viewed,omitempty"`
	Notification         *NotificationData `protobuf:"bytes,7,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Notification) GetSubjectId() int64 {
	if m != nil {
		return m.SubjectId
	}
	return 0
}

func (m *Notification) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Notification) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Notification) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_UNKNOWN
}

func (m *Notification) GetViewed() bool {
	if m != nil {
		return m.Viewed
	}
	return false
}

func (m *Notification) GetNotification() *NotificationData {
	if m != nil {
		return m.Notification
	}
	return nil
}

type GetNotificationsRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	SortBy               string   `protobuf:"bytes,3,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	SortByDirection      string   `protobuf:"bytes,4,opt,name=sortByDirection,proto3" json:"sortByDirection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotificationsRequest) Reset()         { *m = GetNotificationsRequest{} }
func (m *GetNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotificationsRequest) ProtoMessage()    {}
func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *GetNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationsRequest.Unmarshal(m, b)
}
func (m *GetNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *GetNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationsRequest.Merge(m, src)
}
func (m *GetNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotificationsRequest.Size(m)
}
func (m *GetNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationsRequest proto.InternalMessageInfo

func (m *GetNotificationsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetNotificationsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetNotificationsRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *GetNotificationsRequest) GetSortByDirection() string {
	if m != nil {
		return m.SortByDirection
	}
	return ""
}

type GetNotificationsResponse struct {
	Data                 []*Notification `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Count                int64           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Limit                int32           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32           `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	SortBy               string          `protobuf:"bytes,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	SortByDirection      string          `protobuf:"bytes,6,opt,name=sortByDirection,proto3" json:"sortByDirection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetNotificationsResponse) Reset()         { *m = GetNotificationsResponse{} }
func (m *GetNotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotificationsResponse) ProtoMessage()    {}
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4}
}

func (m *GetNotificationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationsResponse.Unmarshal(m, b)
}
func (m *GetNotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationsResponse.Marshal(b, m, deterministic)
}
func (m *GetNotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationsResponse.Merge(m, src)
}
func (m *GetNotificationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetNotificationsResponse.Size(m)
}
func (m *GetNotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationsResponse proto.InternalMessageInfo

func (m *GetNotificationsResponse) GetData() []*Notification {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetNotificationsResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetNotificationsResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetNotificationsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetNotificationsResponse) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *GetNotificationsResponse) GetSortByDirection() string {
	if m != nil {
		return m.SortByDirection
	}
	return ""
}

type ViewedNotificationRequest struct {
	NotificationId       string   `protobuf:"bytes,1,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewedNotificationRequest) Reset()         { *m = ViewedNotificationRequest{} }
func (m *ViewedNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*ViewedNotificationRequest) ProtoMessage()    {}
func (*ViewedNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{5}
}

func (m *ViewedNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewedNotificationRequest.Unmarshal(m, b)
}
func (m *ViewedNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewedNotificationRequest.Marshal(b, m, deterministic)
}
func (m *ViewedNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewedNotificationRequest.Merge(m, src)
}
func (m *ViewedNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_ViewedNotificationRequest.Size(m)
}
func (m *ViewedNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewedNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewedNotificationRequest proto.InternalMessageInfo

func (m *ViewedNotificationRequest) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

type ViewedNotificationResponse struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewedNotificationResponse) Reset()         { *m = ViewedNotificationResponse{} }
func (m *ViewedNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*ViewedNotificationResponse) ProtoMessage()    {}
func (*ViewedNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{6}
}

func (m *ViewedNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewedNotificationResponse.Unmarshal(m, b)
}
func (m *ViewedNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewedNotificationResponse.Marshal(b, m, deterministic)
}
func (m *ViewedNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewedNotificationResponse.Merge(m, src)
}
func (m *ViewedNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_ViewedNotificationResponse.Size(m)
}
func (m *ViewedNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewedNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ViewedNotificationResponse proto.InternalMessageInfo

func (m *ViewedNotificationResponse) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterEnum("discussproto.Type", Type_name, Type_value)
	proto.RegisterType((*UserObj)(nil), "discussproto.UserObj")
	proto.RegisterType((*NotificationData)(nil), "discussproto.NotificationData")
	proto.RegisterType((*Notification)(nil), "discussproto.Notification")
	proto.RegisterType((*GetNotificationsRequest)(nil), "discussproto.GetNotificationsRequest")
	proto.RegisterType((*GetNotificationsResponse)(nil), "discussproto.GetNotificationsResponse")
	proto.RegisterType((*ViewedNotificationRequest)(nil), "discussproto.ViewedNotificationRequest")
	proto.RegisterType((*ViewedNotificationResponse)(nil), "discussproto.ViewedNotificationResponse")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x4e, 0x9c, 0xa4, 0x39, 0x89, 0x72, 0xad, 0xb9, 0xbd, 0xbd, 0xbe, 0x51, 0x85, 0x22,
	0x4b, 0x14, 0x17, 0xd4, 0x18, 0x82, 0xc4, 0xbe, 0x4d, 0x2c, 0x64, 0x81, 0xdc, 0x6a, 0xfa, 0x27,
	0x21, 0xb1, 0x98, 0xd8, 0x43, 0x98, 0xd2, 0x78, 0x8c, 0x67, 0x92, 0xaa, 0x42, 0x08, 0x09, 0x89,
	0x25, 0x2b, 0x76, 0xbc, 0x06, 0xaf, 0xc1, 0x8e, 0x57, 0xe0, 0x19, 0x58, 0x23, 0x8f, 0x1d, 0x6a,
	0xa7, 0xa9, 0xba, 0xca, 0xf9, 0x8e, 0xcf, 0x9c, 0xf3, 0x7d, 0xe7, 0x9b, 0x0c, 0xa0, 0x88, 0x4b,
	0xf6, 0x8a, 0x05, 0x44, 0x32, 0x1e, 0xf5, 0xe3, 0x84, 0x4b, 0x8e, 0xda, 0x21, 0x13, 0xc1, 0x4c,
	0x08, 0x85, 0xba, 0x9b, 0x13, 0xce, 0x27, 0xe7, 0xd4, 0x21, 0x31, 0x73, 0x48, 0x14, 0x71, 0xa9,
	0x4a, 0x45, 0x56, 0xdb, 0x7d, 0x32, 0x61, 0xf2, 0xf5, 0x6c, 0xdc, 0x0f, 0xf8, 0xd4, 0x99, 0x5e,
	0x30, 0xf9, 0x86, 0x5f, 0x38, 0x13, 0xbe, 0xa3, 0x3e, 0xee, 0xcc, 0xc9, 0x39, 0x0b, 0x89, 0xe4,
	0x89, 0x70, 0xfe, 0x84, 0xd9, 0x39, 0xcb, 0x85, 0xc6, 0xb1, 0xa0, 0xc9, 0xfe, 0xf8, 0x0c, 0x75,
	0xa0, 0xc2, 0x42, 0x53, 0xeb, 0x69, 0x76, 0x15, 0x57, 0x58, 0x88, 0x10, 0xe8, 0x11, 0x99, 0x52,
	0xb3, 0xd2, 0xd3, 0xec, 0x26, 0x56, 0x31, 0xda, 0x80, 0x3a, 0x99, 0x13, 0x49, 0x12, 0xb3, 0xaa,
	0xb2, 0x39, 0xb2, 0x3e, 0x6b, 0x60, 0xf8, 0x05, 0x05, 0x23, 0x22, 0x09, 0x32, 0xa1, 0x21, 0x66,
	0xe3, 0x33, 0x1a, 0x48, 0xd5, 0xb5, 0x89, 0x17, 0x10, 0xf5, 0xa0, 0x15, 0x52, 0x11, 0x24, 0x2c,
	0x4e, 0x8b, 0xf3, 0x09, 0xc5, 0x14, 0xda, 0x06, 0x7d, 0x26, 0x68, 0x36, 0xa6, 0x35, 0xf8, 0xb7,
	0x5f, 0x5c, 0x45, 0x3f, 0x67, 0x8c, 0x55, 0x49, 0xca, 0x73, 0x4a, 0x25, 0x31, 0xf5, 0x8c, 0x67,
	0x1a, 0x5b, 0xbf, 0x34, 0x68, 0x17, 0xf9, 0x14, 0xc4, 0x35, 0x95, 0xb8, 0x4d, 0x68, 0xe6, 0x64,
	0xbc, 0x50, 0xcd, 0xaf, 0xe2, 0xab, 0x44, 0x2a, 0x33, 0x6d, 0xed, 0x85, 0x6a, 0x7e, 0x15, 0xe7,
	0x28, 0x3d, 0x15, 0x24, 0x94, 0x48, 0x1a, 0xee, 0xca, 0x7c, 0xde, 0x55, 0x02, 0x6d, 0x81, 0x2e,
	0x2f, 0x63, 0x6a, 0xd6, 0x7a, 0x9a, 0xdd, 0x19, 0xa0, 0x32, 0xe7, 0xa3, 0xcb, 0x98, 0x62, 0xf5,
	0x3d, 0xed, 0x3e, 0x67, 0xf4, 0x82, 0x86, 0x66, 0xbd, 0xa7, 0xd9, 0x6b, 0x38, 0x47, 0x68, 0x0f,
	0xda, 0xc5, 0x5b, 0x60, 0x36, 0x94, 0xf6, 0x3b, 0xe5, 0x3e, 0xcb, 0x5b, 0xc6, 0xa5, 0x33, 0xd6,
	0x27, 0x0d, 0xfe, 0x7b, 0x4a, 0x65, 0xb1, 0x4a, 0x60, 0xfa, 0x76, 0x46, 0x85, 0x44, 0xeb, 0x50,
	0x3b, 0x67, 0x53, 0x96, 0xb9, 0x51, 0xc3, 0x19, 0x48, 0xd7, 0x17, 0x93, 0x49, 0x66, 0x73, 0x0d,
	0xab, 0x38, 0x65, 0x28, 0x78, 0x22, 0xf7, 0x2e, 0x17, 0x36, 0x67, 0x08, 0xd9, 0xf0, 0x77, 0x16,
	0x8d, 0x58, 0x42, 0x03, 0x45, 0x32, 0xdb, 0xc2, 0x72, 0xda, 0xfa, 0xae, 0x81, 0x79, 0x9d, 0x87,
	0x88, 0x79, 0x24, 0x28, 0xea, 0x83, 0x1e, 0x12, 0x49, 0x4c, 0xad, 0x57, 0xb5, 0x5b, 0x83, 0xee,
	0xcd, 0x02, 0xb1, 0xaa, 0x4b, 0x89, 0x07, 0x7c, 0x16, 0xc9, 0xdc, 0xa8, 0x0c, 0x5c, 0xc9, 0xa9,
	0xae, 0x92, 0xa3, 0xaf, 0x94, 0x53, 0xbb, 0x4d, 0x4e, 0x7d, 0xb5, 0x9c, 0x21, 0xfc, 0x7f, 0xa2,
	0x4c, 0x2a, 0xb1, 0xcb, 0xf7, 0xba, 0x05, 0x9d, 0xa2, 0x07, 0xde, 0xe2, 0x9e, 0x2d, 0x65, 0xad,
	0x87, 0xd0, 0x5d, 0xd5, 0x24, 0x5f, 0x0a, 0x02, 0x3d, 0xe4, 0x11, 0x55, 0x67, 0xd7, 0xb0, 0x8a,
	0xef, 0xbf, 0x04, 0x3d, 0xbd, 0x37, 0xa8, 0x05, 0x8d, 0x63, 0xff, 0x99, 0xbf, 0x7f, 0xea, 0x1b,
	0x7f, 0x21, 0x80, 0xba, 0xe7, 0x9f, 0x78, 0x47, 0xae, 0xa1, 0xa1, 0x36, 0xac, 0xed, 0x0e, 0x87,
	0xee, 0xc1, 0x91, 0x3b, 0x32, 0x2a, 0x29, 0x1a, 0xb9, 0xc3, 0xe7, 0x9e, 0xef, 0x8e, 0x8c, 0x6a,
	0x8a, 0xb0, 0x7b, 0x78, 0xb0, 0xef, 0x1f, 0xba, 0x86, 0x8e, 0x10, 0x74, 0x16, 0xe8, 0xc4, 0x73,
	0x4f, 0xdd, 0x91, 0x51, 0x1b, 0x7c, 0xab, 0xc0, 0x3f, 0x45, 0x2e, 0x87, 0x34, 0x99, 0xb3, 0x80,
	0xa2, 0x0f, 0x60, 0x2c, 0x7b, 0x87, 0xee, 0x96, 0x5d, 0xba, 0xe1, 0x8e, 0x75, 0xb7, 0x6e, 0x2b,
	0xcb, 0xd4, 0x5a, 0x9b, 0x1f, 0x7f, 0xfc, 0xfc, 0x52, 0xd9, 0x40, 0xeb, 0xea, 0x3d, 0x2b, 0x2e,
	0xca, 0x99, 0x52, 0xf4, 0x55, 0x03, 0x74, 0x7d, 0x55, 0xe8, 0x5e, 0xb9, 0xf9, 0x8d, 0x8e, 0x74,
	0xed, 0xdb, 0x0b, 0x73, 0x1e, 0x8f, 0x14, 0x8f, 0x07, 0x83, 0xed, 0xeb, 0x3c, 0xde, 0x95, 0xed,
	0x7b, 0xef, 0x64, 0x7f, 0xd3, 0xbd, 0xce, 0x8b, 0xd2, 0xc3, 0x3c, 0xae, 0xab, 0x9f, 0xc7, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x5c, 0x56, 0x7f, 0xc3, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	ViewedNotification(ctx context.Context, in *ViewedNotificationRequest, opts ...grpc.CallOption) (*ViewedNotificationResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/discussproto.NotificationService/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ViewedNotification(ctx context.Context, in *ViewedNotificationRequest, opts ...grpc.CallOption) (*ViewedNotificationResponse, error) {
	out := new(ViewedNotificationResponse)
	err := c.cc.Invoke(ctx, "/discussproto.NotificationService/ViewedNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	ViewedNotification(context.Context, *ViewedNotificationRequest) (*ViewedNotificationResponse, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.NotificationService/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ViewedNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewedNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ViewedNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discussproto.NotificationService/ViewedNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ViewedNotification(ctx, req.(*ViewedNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discussproto.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifications",
			Handler:    _NotificationService_GetNotifications_Handler,
		},
		{
			MethodName: "ViewedNotification",
			Handler:    _NotificationService_ViewedNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
