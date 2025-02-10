// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: recommend.proto

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

type FeedsRecommendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *FeedsRecommendReq) Reset() {
	*x = FeedsRecommendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedsRecommendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedsRecommendReq) ProtoMessage() {}

func (x *FeedsRecommendReq) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedsRecommendReq.ProtoReflect.Descriptor instead.
func (*FeedsRecommendReq) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *FeedsRecommendReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type FeedsRecommendRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vids []uint64 `protobuf:"varint,1,rep,packed,name=vids,proto3" json:"vids,omitempty"`
}

func (x *FeedsRecommendRsp) Reset() {
	*x = FeedsRecommendRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedsRecommendRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedsRecommendRsp) ProtoMessage() {}

func (x *FeedsRecommendRsp) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedsRecommendRsp.ProtoReflect.Descriptor instead.
func (*FeedsRecommendRsp) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *FeedsRecommendRsp) GetVids() []uint64 {
	if x != nil {
		return x.Vids
	}
	return nil
}

var File_recommend_proto protoreflect.FileDescriptor

var file_recommend_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x22, 0x25, 0x0a, 0x11,
	0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x76, 0x69, 0x64, 0x73, 0x32, 0x79, 0x0a, 0x09,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x6c, 0x0a, 0x0e, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x2c, 0x2e, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x61, 0x6d, 0x6f, 0x6e, 0x39, 0x35, 0x35, 0x2f, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommend_proto_rawDescOnce sync.Once
	file_recommend_proto_rawDescData = file_recommend_proto_rawDesc
)

func file_recommend_proto_rawDescGZIP() []byte {
	file_recommend_proto_rawDescOnce.Do(func() {
		file_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommend_proto_rawDescData)
	})
	return file_recommend_proto_rawDescData
}

var file_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_recommend_proto_goTypes = []interface{}{
	(*FeedsRecommendReq)(nil), // 0: trpc.shortvideo.recommend.FeedsRecommendReq
	(*FeedsRecommendRsp)(nil), // 1: trpc.shortvideo.recommend.FeedsRecommendRsp
}
var file_recommend_proto_depIdxs = []int32{
	0, // 0: trpc.shortvideo.recommend.Recommend.FeedsRecommend:input_type -> trpc.shortvideo.recommend.FeedsRecommendReq
	1, // 1: trpc.shortvideo.recommend.Recommend.FeedsRecommend:output_type -> trpc.shortvideo.recommend.FeedsRecommendRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recommend_proto_init() }
func file_recommend_proto_init() {
	if File_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedsRecommendReq); i {
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
		file_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedsRecommendRsp); i {
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
			RawDescriptor: file_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommend_proto_goTypes,
		DependencyIndexes: file_recommend_proto_depIdxs,
		MessageInfos:      file_recommend_proto_msgTypes,
	}.Build()
	File_recommend_proto = out.File
	file_recommend_proto_rawDesc = nil
	file_recommend_proto_goTypes = nil
	file_recommend_proto_depIdxs = nil
}
