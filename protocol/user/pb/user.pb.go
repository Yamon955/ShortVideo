// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: user.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PROFILE_TYPES int32

const (
	PROFILE_TYPES_UNKNOW             PROFILE_TYPES = 0
	PROFILE_TYPES_MAIN_PAGE_INFO     PROFILE_TYPES = 100 // 个人主页信息
	PROFILE_TYPES_PUBLISH_LIST_COUNT PROFILE_TYPES = 101 // 作品列表总数
	PROFILE_TYPES_LIKED_LIST_COUNT   PROFILE_TYPES = 102 // 赞过列表总数
	PROFILE_TYPES_COLLECT_LIST_COUNT PROFILE_TYPES = 103 // 收藏列表总数
	PROFILE_TYPES_USERNAME           PROFILE_TYPES = 200 // 用户名
	PROFILE_TYPES_AVATOR             PROFILE_TYPES = 201 // 头像
	PROFILE_TYPES_SIGN               PROFILE_TYPES = 202 // 简介
	PROFILE_TYPES_GENDER             PROFILE_TYPES = 203 // 性别
)

// Enum value maps for PROFILE_TYPES.
var (
	PROFILE_TYPES_name = map[int32]string{
		0:   "UNKNOW",
		100: "MAIN_PAGE_INFO",
		101: "PUBLISH_LIST_COUNT",
		102: "LIKED_LIST_COUNT",
		103: "COLLECT_LIST_COUNT",
		200: "USERNAME",
		201: "AVATOR",
		202: "SIGN",
		203: "GENDER",
	}
	PROFILE_TYPES_value = map[string]int32{
		"UNKNOW":             0,
		"MAIN_PAGE_INFO":     100,
		"PUBLISH_LIST_COUNT": 101,
		"LIKED_LIST_COUNT":   102,
		"COLLECT_LIST_COUNT": 103,
		"USERNAME":           200,
		"AVATOR":             201,
		"SIGN":               202,
		"GENDER":             203,
	}
)

func (x PROFILE_TYPES) Enum() *PROFILE_TYPES {
	p := new(PROFILE_TYPES)
	*p = x
	return p
}

func (x PROFILE_TYPES) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PROFILE_TYPES) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (PROFILE_TYPES) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x PROFILE_TYPES) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PROFILE_TYPES.Descriptor instead.
func (PROFILE_TYPES) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainPageInfo     *MainPageInfo `protobuf:"bytes,1,opt,name=main_page_info,json=mainPageInfo,proto3" json:"main_page_info,omitempty"`              // 个人主页信息
	PublishListCount int32         `protobuf:"varint,2,opt,name=publish_list_count,json=publishListCount,proto3" json:"publish_list_count,omitempty"` // 作品列表总数
	LikedListCount   int32         `protobuf:"varint,3,opt,name=liked_list_count,json=likedListCount,proto3" json:"liked_list_count,omitempty"`       // 赞过列表总数
	CollectListCount int32         `protobuf:"varint,4,opt,name=collect_list_count,json=collectListCount,proto3" json:"collect_list_count,omitempty"` // 收藏列表总数
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetMainPageInfo() *MainPageInfo {
	if x != nil {
		return x.MainPageInfo
	}
	return nil
}

func (x *UserInfo) GetPublishListCount() int32 {
	if x != nil {
		return x.PublishListCount
	}
	return 0
}

func (x *UserInfo) GetLikedListCount() int32 {
	if x != nil {
		return x.LikedListCount
	}
	return 0
}

func (x *UserInfo) GetCollectListCount() int32 {
	if x != nil {
		return x.CollectListCount
	}
	return 0
}

// 个人主页信息
type MainPageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`                                         // uid
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`                              // 用户名
	Avator       string `protobuf:"bytes,3,opt,name=avator,proto3" json:"avator,omitempty"`                                  // 头像
	Sign         string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`                                      // 简介
	Gender       int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`                                 // 性别 0-未知 1-男 2-女
	FansCount    int32  `protobuf:"varint,6,opt,name=fans_count,json=fansCount,proto3" json:"fans_count,omitempty"`          // 粉丝数
	FollowsCount int32  `protobuf:"varint,7,opt,name=follows_count,json=followsCount,proto3" json:"follows_count,omitempty"` // 关注数
	RegisterTime int64  `protobuf:"varint,8,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"` // 注册时间
}

func (x *MainPageInfo) Reset() {
	*x = MainPageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainPageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainPageInfo) ProtoMessage() {}

func (x *MainPageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainPageInfo.ProtoReflect.Descriptor instead.
func (*MainPageInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *MainPageInfo) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MainPageInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MainPageInfo) GetAvator() string {
	if x != nil {
		return x.Avator
	}
	return ""
}

func (x *MainPageInfo) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *MainPageInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *MainPageInfo) GetFansCount() int32 {
	if x != nil {
		return x.FansCount
	}
	return 0
}

func (x *MainPageInfo) GetFollowsCount() int32 {
	if x != nil {
		return x.FollowsCount
	}
	return 0
}

func (x *MainPageInfo) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RegisterRsp) Reset() {
	*x = RegisterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRsp) ProtoMessage() {}

func (x *RegisterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRsp.ProtoReflect.Descriptor instead.
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Uid   uint64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginRsp) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type BatchGetProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uids         []uint64        `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	ProfileTypes []PROFILE_TYPES `protobuf:"varint,2,rep,packed,name=profileTypes,proto3,enum=trpc.shortvideo.user.PROFILE_TYPES" json:"profileTypes,omitempty"`
	DisableCache bool            `protobuf:"varint,3,opt,name=disableCache,proto3" json:"disableCache,omitempty"`
}

func (x *BatchGetProfileReq) Reset() {
	*x = BatchGetProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetProfileReq) ProtoMessage() {}

func (x *BatchGetProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetProfileReq.ProtoReflect.Descriptor instead.
func (*BatchGetProfileReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *BatchGetProfileReq) GetUids() []uint64 {
	if x != nil {
		return x.Uids
	}
	return nil
}

func (x *BatchGetProfileReq) GetProfileTypes() []PROFILE_TYPES {
	if x != nil {
		return x.ProfileTypes
	}
	return nil
}

func (x *BatchGetProfileReq) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

type BatchGetProfileRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfos  []*UserInfo `protobuf:"bytes,1,rep,name=userInfos,proto3" json:"userInfos,omitempty"`
	FailedUIDs []int64     `protobuf:"varint,2,rep,packed,name=failedUIDs,proto3" json:"failedUIDs,omitempty"`
}

func (x *BatchGetProfileRsp) Reset() {
	*x = BatchGetProfileRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetProfileRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetProfileRsp) ProtoMessage() {}

func (x *BatchGetProfileRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetProfileRsp.ProtoReflect.Descriptor instead.
func (*BatchGetProfileRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *BatchGetProfileRsp) GetUserInfos() []*UserInfo {
	if x != nil {
		return x.UserInfos
	}
	return nil
}

func (x *BatchGetProfileRsp) GetFailedUIDs() []int64 {
	if x != nil {
		return x.FailedUIDs
	}
	return nil
}

type SetProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          uint64          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProfileTypes []PROFILE_TYPES `protobuf:"varint,2,rep,packed,name=profileTypes,proto3,enum=trpc.shortvideo.user.PROFILE_TYPES" json:"profileTypes,omitempty"`
	DisableCache bool            `protobuf:"varint,3,opt,name=disableCache,proto3" json:"disableCache,omitempty"`
	Username     string          `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"` // 用户名
	Avator       string          `protobuf:"bytes,5,opt,name=avator,proto3" json:"avator,omitempty"`     // 头像
	Sign         string          `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`         // 简介
	Gender       int32           `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`    // 性别 0-未知 1-男 2-女
}

func (x *SetProfileReq) Reset() {
	*x = SetProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetProfileReq) ProtoMessage() {}

func (x *SetProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetProfileReq.ProtoReflect.Descriptor instead.
func (*SetProfileReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *SetProfileReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetProfileReq) GetProfileTypes() []PROFILE_TYPES {
	if x != nil {
		return x.ProfileTypes
	}
	return nil
}

func (x *SetProfileReq) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

func (x *SetProfileReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SetProfileReq) GetAvator() string {
	if x != nil {
		return x.Avator
	}
	return ""
}

func (x *SetProfileReq) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *SetProfileReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

type SetProfileRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo    *UserInfo        `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	FailedTypes map[int32]string `protobuf:"bytes,2,rep,name=failedTypes,proto3" json:"failedTypes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetProfileRsp) Reset() {
	*x = SetProfileRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetProfileRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetProfileRsp) ProtoMessage() {}

func (x *SetProfileRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetProfileRsp.ProtoReflect.Descriptor instead.
func (*SetProfileRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *SetProfileRsp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *SetProfileRsp) GetFailedTypes() map[int32]string {
	if x != nil {
		return x.FailedTypes
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xda, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x48, 0x0a, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d,
	0x61, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x6d, 0x61, 0x69,
	0x6e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x69, 0x6b, 0x65, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xe7, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x33, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x58, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x12, 0x47,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x52, 0x4f, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x72, 0x0a, 0x12, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x3c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x55, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x55, 0x49, 0x44, 0x73, 0x22,
	0xee, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x3a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x56,
	0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0xa9, 0x01, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x50, 0x41, 0x47,
	0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x64, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x53, 0x48, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x4b, 0x45, 0x44, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x10, 0x66, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43,
	0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x67, 0x12, 0x0d,
	0x0a, 0x08, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0xc8, 0x01, 0x12, 0x0b, 0x0a,
	0x06, 0x41, 0x56, 0x41, 0x54, 0x4f, 0x52, 0x10, 0xc9, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x53, 0x49,
	0x47, 0x4e, 0x10, 0xca, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10,
	0xcb, 0x01, 0x32, 0xe0, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x47, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x65, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x73, 0x70, 0x12, 0x56, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x73, 0x70, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x61, 0x6d, 0x6f, 0x6e, 0x39, 0x35, 0x35, 0x2f, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_proto_goTypes = []interface{}{
	(PROFILE_TYPES)(0),         // 0: trpc.shortvideo.user.PROFILE_TYPES
	(*UserInfo)(nil),           // 1: trpc.shortvideo.user.UserInfo
	(*MainPageInfo)(nil),       // 2: trpc.shortvideo.user.MainPageInfo
	(*RegisterReq)(nil),        // 3: trpc.shortvideo.user.RegisterReq
	(*RegisterRsp)(nil),        // 4: trpc.shortvideo.user.RegisterRsp
	(*LoginReq)(nil),           // 5: trpc.shortvideo.user.LoginReq
	(*LoginRsp)(nil),           // 6: trpc.shortvideo.user.LoginRsp
	(*BatchGetProfileReq)(nil), // 7: trpc.shortvideo.user.BatchGetProfileReq
	(*BatchGetProfileRsp)(nil), // 8: trpc.shortvideo.user.BatchGetProfileRsp
	(*SetProfileReq)(nil),      // 9: trpc.shortvideo.user.SetProfileReq
	(*SetProfileRsp)(nil),      // 10: trpc.shortvideo.user.SetProfileRsp
	nil,                        // 11: trpc.shortvideo.user.SetProfileRsp.FailedTypesEntry
}
var file_user_proto_depIdxs = []int32{
	2,  // 0: trpc.shortvideo.user.UserInfo.main_page_info:type_name -> trpc.shortvideo.user.MainPageInfo
	0,  // 1: trpc.shortvideo.user.BatchGetProfileReq.profileTypes:type_name -> trpc.shortvideo.user.PROFILE_TYPES
	1,  // 2: trpc.shortvideo.user.BatchGetProfileRsp.userInfos:type_name -> trpc.shortvideo.user.UserInfo
	0,  // 3: trpc.shortvideo.user.SetProfileReq.profileTypes:type_name -> trpc.shortvideo.user.PROFILE_TYPES
	1,  // 4: trpc.shortvideo.user.SetProfileRsp.userInfo:type_name -> trpc.shortvideo.user.UserInfo
	11, // 5: trpc.shortvideo.user.SetProfileRsp.failedTypes:type_name -> trpc.shortvideo.user.SetProfileRsp.FailedTypesEntry
	3,  // 6: trpc.shortvideo.user.User.Register:input_type -> trpc.shortvideo.user.RegisterReq
	5,  // 7: trpc.shortvideo.user.User.Login:input_type -> trpc.shortvideo.user.LoginReq
	7,  // 8: trpc.shortvideo.user.User.BatchGetProfile:input_type -> trpc.shortvideo.user.BatchGetProfileReq
	9,  // 9: trpc.shortvideo.user.User.SetProfile:input_type -> trpc.shortvideo.user.SetProfileReq
	4,  // 10: trpc.shortvideo.user.User.Register:output_type -> trpc.shortvideo.user.RegisterRsp
	6,  // 11: trpc.shortvideo.user.User.Login:output_type -> trpc.shortvideo.user.LoginRsp
	8,  // 12: trpc.shortvideo.user.User.BatchGetProfile:output_type -> trpc.shortvideo.user.BatchGetProfileRsp
	10, // 13: trpc.shortvideo.user.User.SetProfile:output_type -> trpc.shortvideo.user.SetProfileRsp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainPageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetProfileReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetProfileRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetProfileReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetProfileRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
