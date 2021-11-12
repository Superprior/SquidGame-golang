// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: squidgame.proto

package squidgame

import (
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

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gamenumber  string `protobuf:"bytes,1,opt,name=gamenumber,proto3" json:"gamenumber,omitempty"`
	Gamename    string `protobuf:"bytes,2,opt,name=gamename,proto3" json:"gamename,omitempty"`
	Players     int64  `protobuf:"varint,3,opt,name=players,proto3" json:"players,omitempty"`
	Rungames    int64  `protobuf:"varint,4,opt,name=rungames,proto3" json:"rungames,omitempty"`
	Concurrence int64  `protobuf:"varint,5,opt,name=concurrence,proto3" json:"concurrence,omitempty"`
	Timeout     int64  `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squidgame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_squidgame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_squidgame_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetGamenumber() string {
	if x != nil {
		return x.Gamenumber
	}
	return ""
}

func (x *Game) GetGamename() string {
	if x != nil {
		return x.Gamename
	}
	return ""
}

func (x *Game) GetPlayers() int64 {
	if x != nil {
		return x.Players
	}
	return 0
}

func (x *Game) GetRungames() int64 {
	if x != nil {
		return x.Rungames
	}
	return 0
}

func (x *Game) GetConcurrence() int64 {
	if x != nil {
		return x.Concurrence
	}
	return 0
}

func (x *Game) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type PlayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squidgame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squidgame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_squidgame_proto_rawDescGZIP(), []int{1}
}

func (x *PlayRequest) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squidgame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squidgame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_squidgame_proto_rawDescGZIP(), []int{2}
}

func (x *PlayResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_squidgame_proto protoreflect.FileDescriptor

var file_squidgame_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61,
	0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61,
	0x6d, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61,
	0x6d, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x28, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x39, 0x0a, 0x10,
	0x53, 0x71, 0x75, 0x69, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x73, 0x63, 0x61, 0x72, 0x4c, 0x6c, 0x61, 0x6d, 0x61,
	0x73, 0x36, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_squidgame_proto_rawDescOnce sync.Once
	file_squidgame_proto_rawDescData = file_squidgame_proto_rawDesc
)

func file_squidgame_proto_rawDescGZIP() []byte {
	file_squidgame_proto_rawDescOnce.Do(func() {
		file_squidgame_proto_rawDescData = protoimpl.X.CompressGZIP(file_squidgame_proto_rawDescData)
	})
	return file_squidgame_proto_rawDescData
}

var file_squidgame_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_squidgame_proto_goTypes = []interface{}{
	(*Game)(nil),         // 0: Game
	(*PlayRequest)(nil),  // 1: PlayRequest
	(*PlayResponse)(nil), // 2: PlayResponse
}
var file_squidgame_proto_depIdxs = []int32{
	0, // 0: PlayRequest.game:type_name -> Game
	1, // 1: SquidGameService.Play:input_type -> PlayRequest
	2, // 2: SquidGameService.Play:output_type -> PlayResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_squidgame_proto_init() }
func file_squidgame_proto_init() {
	if File_squidgame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_squidgame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_squidgame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest); i {
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
		file_squidgame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
			RawDescriptor: file_squidgame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_squidgame_proto_goTypes,
		DependencyIndexes: file_squidgame_proto_depIdxs,
		MessageInfos:      file_squidgame_proto_msgTypes,
	}.Build()
	File_squidgame_proto = out.File
	file_squidgame_proto_rawDesc = nil
	file_squidgame_proto_goTypes = nil
	file_squidgame_proto_depIdxs = nil
}
