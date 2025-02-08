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

type VideoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid           uint64 `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	VideoUrl      string `protobuf:"bytes,2,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	CoverUrl      string `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Tags          uint64 `protobuf:"varint,5,opt,name=tags,proto3" json:"tags,omitempty"`
	LikesCount    int32  `protobuf:"varint,6,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	CollectCount  int32  `protobuf:"varint,7,opt,name=collect_count,json=collectCount,proto3" json:"collect_count,omitempty"`
	CommentsCount int32  `protobuf:"varint,8,opt,name=comments_count,json=commentsCount,proto3" json:"comments_count,omitempty"`
	PublishTime   int64  `protobuf:"varint,9,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Liked         bool   `protobuf:"varint,10,opt,name=liked,proto3" json:"liked,omitempty"`
	Collected     bool   `protobuf:"varint,11,opt,name=collected,proto3" json:"collected,omitempty"`
}

func (x *VideoInfo) Reset() {
	*x = VideoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfo) ProtoMessage() {}

func (x *VideoInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VideoInfo.ProtoReflect.Descriptor instead.
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *VideoInfo) GetVid() uint64 {
	if x != nil {
		return x.Vid
	}
	return 0
}

func (x *VideoInfo) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoInfo) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoInfo) GetTags() uint64 {
	if x != nil {
		return x.Tags
	}
	return 0
}

func (x *VideoInfo) GetLikesCount() int32 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

func (x *VideoInfo) GetCollectCount() int32 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

func (x *VideoInfo) GetCommentsCount() int32 {
	if x != nil {
		return x.CommentsCount
	}
	return 0
}

func (x *VideoInfo) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

func (x *VideoInfo) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

func (x *VideoInfo) GetCollected() bool {
	if x != nil {
		return x.Collected
	}
	return false
}

type GetFeedsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeedsReq) Reset() {
	*x = GetFeedsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsReq) ProtoMessage() {}

func (x *GetFeedsReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetFeedsReq.ProtoReflect.Descriptor instead.
func (*GetFeedsReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

type GetFeedsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeedsRsp) Reset() {
	*x = GetFeedsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsRsp) ProtoMessage() {}

func (x *GetFeedsRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetFeedsRsp.ProtoReflect.Descriptor instead.
func (*GetFeedsRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

type PublishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Vid        uint64 `protobuf:"varint,3,opt,name=vid,proto3" json:"vid,omitempty"`
}

func (x *PublishRsp) Reset() {
	*x = PublishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRsp) ProtoMessage() {}

func (x *PublishRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublishRsp.ProtoReflect.Descriptor instead.
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
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

func (x *PublishRsp) GetVid() uint64 {
	if x != nil {
		return x.Vid
	}
	return 0
}

type GetPublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId uint64 `protobuf:"varint,1,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
}

func (x *GetPublishListReq) Reset() {
	*x = GetPublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishListReq) ProtoMessage() {}

func (x *GetPublishListReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPublishListReq.ProtoReflect.Descriptor instead.
func (*GetPublishListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *GetPublishListReq) GetStartId() uint64 {
	if x != nil {
		return x.StartId
	}
	return 0
}

type GetPublishListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextId     uint64       `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	IsFinash   bool         `protobuf:"varint,2,opt,name=is_finash,json=isFinash,proto3" json:"is_finash,omitempty"`
	VideoInfos []*VideoInfo `protobuf:"bytes,3,rep,name=video_infos,json=videoInfos,proto3" json:"video_infos,omitempty"`
}

func (x *GetPublishListRsp) Reset() {
	*x = GetPublishListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishListRsp) ProtoMessage() {}

func (x *GetPublishListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[6]
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
	return file_video_proto_rawDescGZIP(), []int{6}
}

func (x *GetPublishListRsp) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

func (x *GetPublishListRsp) GetIsFinash() bool {
	if x != nil {
		return x.IsFinash
	}
	return false
}

func (x *GetPublishListRsp) GetVideoInfos() []*VideoInfo {
	if x != nil {
		return x.VideoInfos
	}
	return nil
}

type GetLikeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId uint64 `protobuf:"varint,1,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
}

func (x *GetLikeListReq) Reset() {
	*x = GetLikeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikeListReq) ProtoMessage() {}

func (x *GetLikeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikeListReq.ProtoReflect.Descriptor instead.
func (*GetLikeListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{7}
}

func (x *GetLikeListReq) GetStartId() uint64 {
	if x != nil {
		return x.StartId
	}
	return 0
}

type GetLikeListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextId     uint64       `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	IsFinash   bool         `protobuf:"varint,2,opt,name=is_finash,json=isFinash,proto3" json:"is_finash,omitempty"`
	VideoInfos []*VideoInfo `protobuf:"bytes,3,rep,name=video_infos,json=videoInfos,proto3" json:"video_infos,omitempty"`
}

func (x *GetLikeListRsp) Reset() {
	*x = GetLikeListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikeListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikeListRsp) ProtoMessage() {}

func (x *GetLikeListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikeListRsp.ProtoReflect.Descriptor instead.
func (*GetLikeListRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{8}
}

func (x *GetLikeListRsp) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

func (x *GetLikeListRsp) GetIsFinash() bool {
	if x != nil {
		return x.IsFinash
	}
	return false
}

func (x *GetLikeListRsp) GetVideoInfos() []*VideoInfo {
	if x != nil {
		return x.VideoInfos
	}
	return nil
}

type GetCollectListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId uint64 `protobuf:"varint,1,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
}

func (x *GetCollectListReq) Reset() {
	*x = GetCollectListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectListReq) ProtoMessage() {}

func (x *GetCollectListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectListReq.ProtoReflect.Descriptor instead.
func (*GetCollectListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{9}
}

func (x *GetCollectListReq) GetStartId() uint64 {
	if x != nil {
		return x.StartId
	}
	return 0
}

type GetCollectListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextId     uint64       `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	IsFinash   bool         `protobuf:"varint,2,opt,name=is_finash,json=isFinash,proto3" json:"is_finash,omitempty"`
	VideoInfos []*VideoInfo `protobuf:"bytes,3,rep,name=video_infos,json=videoInfos,proto3" json:"video_infos,omitempty"`
}

func (x *GetCollectListRsp) Reset() {
	*x = GetCollectListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectListRsp) ProtoMessage() {}

func (x *GetCollectListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectListRsp.ProtoReflect.Descriptor instead.
func (*GetCollectListRsp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{10}
}

func (x *GetCollectListRsp) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

func (x *GetCollectListRsp) GetIsFinash() bool {
	if x != nil {
		return x.IsFinash
	}
	return false
}

func (x *GetCollectListRsp) GetVideoInfos() []*VideoInfo {
	if x != nil {
		return x.VideoInfos
	}
	return nil
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x76, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x0d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x22, 0x0d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x73, 0x70, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x22, 0x5e, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x76, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x6e, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a,
	0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64,
	0x22, 0x8c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x0b,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x2a,
	0xd0, 0x01, 0x0a, 0x0a, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x54, 0x41, 0x47, 0x53, 0x12, 0x11,
	0x0a, 0x0d, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x54, 0x41, 0x49, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49,
	0x46, 0x45, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x44, 0x55,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x43, 0x48,
	0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x73, 0x68,
	0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x53, 0x10,
	0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x08, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x52, 0x54, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45,
	0x57, 0x53, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x0b,
	0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x49, 0x59, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x45,
	0x10, 0x0e, 0x32, 0x84, 0x03, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x52, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x73, 0x70,
	0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x5b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x73, 0x70, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x28, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x32, 0x5a, 0x0a, 0x07, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x4f, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12,
	0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x73, 0x70, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x61, 0x6d, 0x6f, 0x6e, 0x39, 0x35, 0x35, 0x2f, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_video_proto_goTypes = []interface{}{
	(VIDEO_TAGS)(0),           // 0: trpc.shortvideo.video.VIDEO_TAGS
	(*VideoInfo)(nil),         // 1: trpc.shortvideo.video.VideoInfo
	(*GetFeedsReq)(nil),       // 2: trpc.shortvideo.video.GetFeedsReq
	(*GetFeedsRsp)(nil),       // 3: trpc.shortvideo.video.GetFeedsRsp
	(*PublishReq)(nil),        // 4: trpc.shortvideo.video.PublishReq
	(*PublishRsp)(nil),        // 5: trpc.shortvideo.video.PublishRsp
	(*GetPublishListReq)(nil), // 6: trpc.shortvideo.video.GetPublishListReq
	(*GetPublishListRsp)(nil), // 7: trpc.shortvideo.video.GetPublishListRsp
	(*GetLikeListReq)(nil),    // 8: trpc.shortvideo.video.GetLikeListReq
	(*GetLikeListRsp)(nil),    // 9: trpc.shortvideo.video.GetLikeListRsp
	(*GetCollectListReq)(nil), // 10: trpc.shortvideo.video.GetCollectListReq
	(*GetCollectListRsp)(nil), // 11: trpc.shortvideo.video.GetCollectListRsp
}
var file_video_proto_depIdxs = []int32{
	1,  // 0: trpc.shortvideo.video.GetPublishListRsp.video_infos:type_name -> trpc.shortvideo.video.VideoInfo
	1,  // 1: trpc.shortvideo.video.GetLikeListRsp.video_infos:type_name -> trpc.shortvideo.video.VideoInfo
	1,  // 2: trpc.shortvideo.video.GetCollectListRsp.video_infos:type_name -> trpc.shortvideo.video.VideoInfo
	2,  // 3: trpc.shortvideo.video.Video.GetFeeds:input_type -> trpc.shortvideo.video.GetFeedsReq
	6,  // 4: trpc.shortvideo.video.Video.GetPublishList:input_type -> trpc.shortvideo.video.GetPublishListReq
	8,  // 5: trpc.shortvideo.video.Video.GetLikeList:input_type -> trpc.shortvideo.video.GetLikeListReq
	10, // 6: trpc.shortvideo.video.Video.GetCollectList:input_type -> trpc.shortvideo.video.GetCollectListReq
	4,  // 7: trpc.shortvideo.video.Publish.Publish:input_type -> trpc.shortvideo.video.PublishReq
	3,  // 8: trpc.shortvideo.video.Video.GetFeeds:output_type -> trpc.shortvideo.video.GetFeedsRsp
	7,  // 9: trpc.shortvideo.video.Video.GetPublishList:output_type -> trpc.shortvideo.video.GetPublishListRsp
	9,  // 10: trpc.shortvideo.video.Video.GetLikeList:output_type -> trpc.shortvideo.video.GetLikeListRsp
	11, // 11: trpc.shortvideo.video.Video.GetCollectList:output_type -> trpc.shortvideo.video.GetCollectListRsp
	5,  // 12: trpc.shortvideo.video.Publish.Publish:output_type -> trpc.shortvideo.video.PublishRsp
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoInfo); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikeListReq); i {
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
		file_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikeListRsp); i {
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
		file_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectListReq); i {
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
		file_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectListRsp); i {
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
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
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
