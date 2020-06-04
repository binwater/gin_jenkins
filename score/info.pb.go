// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: info.proto

package score

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// @brief base_score_info
type BaseScoreInfoT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WinCount       *int32 `protobuf:"varint,1,opt,name=win_count,json=winCount" json:"win_count,omitempty"`                   // 玩家胜局局数
	LoseCount      *int32 `protobuf:"varint,2,opt,name=lose_count,json=loseCount" json:"lose_count,omitempty"`                // 玩家负局局数
	ExceptionCount *int32 `protobuf:"varint,3,opt,name=exception_count,json=exceptionCount" json:"exception_count,omitempty"` // 玩家异常局局数
	KillCount      *int32 `protobuf:"varint,4,opt,name=kill_count,json=killCount" json:"kill_count,omitempty"`                // 总人头数
	DeathCount     *int32 `protobuf:"varint,5,opt,name=death_count,json=deathCount" json:"death_count,omitempty"`             // 总死亡数
	AssistCount    *int32 `protobuf:"varint,6,opt,name=assist_count,json=assistCount" json:"assist_count,omitempty"`          // 总总助攻数
	Rating         *int64 `protobuf:"varint,7,opt,name=rating" json:"rating,omitempty"`                                       // 评价积分
}

func (x *BaseScoreInfoT) Reset() {
	*x = BaseScoreInfoT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseScoreInfoT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseScoreInfoT) ProtoMessage() {}

func (x *BaseScoreInfoT) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseScoreInfoT.ProtoReflect.Descriptor instead.
func (*BaseScoreInfoT) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{0}
}

func (x *BaseScoreInfoT) GetWinCount() int32 {
	if x != nil && x.WinCount != nil {
		return *x.WinCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetLoseCount() int32 {
	if x != nil && x.LoseCount != nil {
		return *x.LoseCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetExceptionCount() int32 {
	if x != nil && x.ExceptionCount != nil {
		return *x.ExceptionCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetKillCount() int32 {
	if x != nil && x.KillCount != nil {
		return *x.KillCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetDeathCount() int32 {
	if x != nil && x.DeathCount != nil {
		return *x.DeathCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetAssistCount() int32 {
	if x != nil && x.AssistCount != nil {
		return *x.AssistCount
	}
	return 0
}

func (x *BaseScoreInfoT) GetRating() int64 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

var File_info_proto protoreflect.FileDescriptor

var file_info_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x69,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x73, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x61, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x61, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
}

var (
	file_info_proto_rawDescOnce sync.Once
	file_info_proto_rawDescData = file_info_proto_rawDesc
)

func file_info_proto_rawDescGZIP() []byte {
	file_info_proto_rawDescOnce.Do(func() {
		file_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_proto_rawDescData)
	})
	return file_info_proto_rawDescData
}

var file_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_info_proto_goTypes = []interface{}{
	(*BaseScoreInfoT)(nil), // 0: score.base_score_info_t
}
var file_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_info_proto_init() }
func file_info_proto_init() {
	if File_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseScoreInfoT); i {
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
			RawDescriptor: file_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_info_proto_goTypes,
		DependencyIndexes: file_info_proto_depIdxs,
		MessageInfos:      file_info_proto_msgTypes,
	}.Build()
	File_info_proto = out.File
	file_info_proto_rawDesc = nil
	file_info_proto_goTypes = nil
	file_info_proto_depIdxs = nil
}