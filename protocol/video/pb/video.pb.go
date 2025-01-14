// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: video.proto

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

type VIDEO_TAGS int32

const (
	VIDEO_TAGS_ENTERTAINMENT VIDEO_TAGS = 0  // 娱乐
	VIDEO_TAGS_BUSINESS      VIDEO_TAGS = 1  // 商业
	VIDEO_TAGS_MUSIC         VIDEO_TAGS = 2  // 音乐
	VIDEO_TAGS_LIFESTYLE     VIDEO_TAGS = 3  // 生活
	VIDEO_TAGS_EDUCATION     VIDEO_TAGS = 4  // 教育
	VIDEO_TAGS_TECHNOLOGY    VIDEO_TAGS = 5  // 科技
	VIDEO_TAGS_Fashion       VIDEO_TAGS = 6  // 时尚
	VIDEO_TAGS_SPORTS        VIDEO_TAGS = 7  // 运动
	VIDEO_TAGS_AUTOMOTIVE    VIDEO_TAGS = 8  // 汽车
	VIDEO_TAGS_ART           VIDEO_TAGS = 9  // 艺术
	VIDEO_TAGS_NEWS          VIDEO_TAGS = 10 // 新闻
	VIDEO_TAGS_GAMING        VIDEO_TAGS = 11 // 游戏
	VIDEO_TAGS_HEALTH        VIDEO_TAGS = 12 // 健康
	VIDEO_TAGS_DIY           VIDEO_TAGS = 13 // DIY
	VIDEO_TAGS_CULTURE       VIDEO_TAGS = 14 // 文化
)

// Enum value maps for VIDEO_TAGS.
var (
	VIDEO_TAGS_name = map[int32]string{
		0:  "ENTERTAINMENT",
		1:  "BUSINESS",
		2:  "MUSIC",
		3:  "LIFESTYLE",
		4:  "EDUCATION",
		5:  "TECHNOLOGY",
		6:  "Fashion",
		7:  "SPORTS",
		8:  "AUTOMOTIVE",
		9:  "ART",
		10: "NEWS",
		11: "GAMING",
		12: "HEALTH",
		13: "DIY",
		14: "CULTURE",
	}
	VIDEO_TAGS_value = map[string]int32{
		"ENTERTAINMENT": 0,
		"BUSINESS":      1,
		"MUSIC":         2,
		"LIFESTYLE":     3,
		"EDUCATION":     4,
		"TECHNOLOGY":    5,
		"Fashion":       6,
		"SPORTS":        7,
		"AUTOMOTIVE":    8,
		"ART":           9,
		"NEWS":          10,
		"GAMING":        11,
		"HEALTH":        12,
		"DIY":           13,
		"CULTURE":       14,
	}
)

func (x VIDEO_TAGS) Enum() *VIDEO_TAGS {
	p := new(VIDEO_TAGS)
	*p = x
	return p
}

func (x VIDEO_TAGS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VIDEO_TAGS) Descriptor() protoreflect.EnumDescriptor {
	return file_video_proto_enumTypes[0].Descriptor()
}

func (VIDEO_TAGS) Type() protoreflect.EnumType {
	return &file_video_proto_enumTypes[0]
}

func (x VIDEO_TAGS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VIDEO_TAGS.Descriptor instead.
func (VIDEO_TAGS) EnumDescriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

type GetFeedsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeedsReq) Reset() {
	*x = GetFeedsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsReq) ProtoMessage() {}

func (x *GetFeedsReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsReq.ProtoReflect.Descriptor instead.
func (*GetFeedsReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

type GetFeedsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeedsRsp) Reset() {
	*x = GetFeedsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsRsp) ProtoMessage() {}

func (x *GetFeedsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsRsp.ProtoReflect.Descriptor instead.
func (*GetFeedsRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

type PublishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *PublishRsp) Reset() {
	*x = PublishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRsp) ProtoMessage() {}

func (x *PublishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRsp.ProtoReflect.Descriptor instead.
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *PublishRsp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PublishRsp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type GetPublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPublishListReq) Reset() {
	*x = GetPublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishListReq) ProtoMessage() {}

func (x *GetPublishListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublishListReq.ProtoReflect.Descriptor instead.
func (*GetPublishListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

type GetPublishListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPublishListRsp) Reset() {
	*x = GetPublishListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishListRsp) ProtoMessage() {}

func (x *GetPublishListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublishListRsp.ProtoReflect.Descriptor instead.
func (*GetPublishListRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52,
	0x73, 0x70, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x22, 0x4c, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2a, 0xd0, 0x01, 0x0a, 0x0a, 0x56, 0x49, 0x44,
	0x45, 0x4f, 0x5f, 0x54, 0x41, 0x47, 0x53, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x54, 0x45, 0x52,
	0x54, 0x41, 0x49, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55,
	0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x55, 0x53, 0x49,
	0x43, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x46, 0x45, 0x53, 0x54, 0x59, 0x4c, 0x45,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x73, 0x68, 0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x53, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55,
	0x54, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x56, 0x45, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x52,
	0x54, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x0a, 0x12, 0x0a, 0x0a,
	0x06, 0x47, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x49, 0x59, 0x10, 0x0d, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x45, 0x10, 0x0e, 0x32, 0x92, 0x02, 0x0a, 0x05,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x52, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64,
	0x73, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x73, 0x70, 0x12, 0x4f, 0x0a, 0x07, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59,
	0x61, 0x6d, 0x6f, 0x6e, 0x39, 0x35, 0x35, 0x2f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_video_proto_goTypes = []interface{}{
	(VIDEO_TAGS)(0),           // 0: trpc.shortvideo.video.VIDEO_TAGS
	(*GetFeedsReq)(nil),       // 1: trpc.shortvideo.video.GetFeedsReq
	(*GetFeedsRsp)(nil),       // 2: trpc.shortvideo.video.GetFeedsRsp
	(*PublishReq)(nil),        // 3: trpc.shortvideo.video.PublishReq
	(*PublishRsp)(nil),        // 4: trpc.shortvideo.video.PublishRsp
	(*GetPublishListReq)(nil), // 5: trpc.shortvideo.video.GetPublishListReq
	(*GetPublishListRsp)(nil), // 6: trpc.shortvideo.video.GetPublishListRsp
}
var file_video_proto_depIdxs = []int32{
	1, // 0: trpc.shortvideo.video.Video.GetFeeds:input_type -> trpc.shortvideo.video.GetFeedsReq
	3, // 1: trpc.shortvideo.video.Video.Publish:input_type -> trpc.shortvideo.video.PublishReq
	5, // 2: trpc.shortvideo.video.Video.GetPublishList:input_type -> trpc.shortvideo.video.GetPublishListReq
	2, // 3: trpc.shortvideo.video.Video.GetFeeds:output_type -> trpc.shortvideo.video.GetFeedsRsp
	4, // 4: trpc.shortvideo.video.Video.Publish:output_type -> trpc.shortvideo.video.PublishRsp
	6, // 5: trpc.shortvideo.video.Video.GetPublishList:output_type -> trpc.shortvideo.video.GetPublishListRsp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsReq); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsRsp); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRsp); i {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublishListReq); i {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublishListRsp); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		EnumInfos:         file_video_proto_enumTypes,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}