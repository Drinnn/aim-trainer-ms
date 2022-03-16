// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: set_highscore_request.proto

package protos

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

type SetHighscoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Highscore float64 `protobuf:"fixed64,1,opt,name=highscore,proto3" json:"highscore,omitempty"`
}

func (x *SetHighscoreRequest) Reset() {
	*x = SetHighscoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_highscore_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHighscoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHighscoreRequest) ProtoMessage() {}

func (x *SetHighscoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_highscore_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHighscoreRequest.ProtoReflect.Descriptor instead.
func (*SetHighscoreRequest) Descriptor() ([]byte, []int) {
	return file_set_highscore_request_proto_rawDescGZIP(), []int{0}
}

func (x *SetHighscoreRequest) GetHighscore() float64 {
	if x != nil {
		return x.Highscore
	}
	return 0
}

var File_set_highscore_request_proto protoreflect.FileDescriptor

var file_set_highscore_request_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x33, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x72, 0x69, 0x6e, 0x6e, 0x6e, 0x2f,
	0x61, 0x69, 0x6d, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x6d, 0x73, 0x2f, 0x68,
	0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_set_highscore_request_proto_rawDescOnce sync.Once
	file_set_highscore_request_proto_rawDescData = file_set_highscore_request_proto_rawDesc
)

func file_set_highscore_request_proto_rawDescGZIP() []byte {
	file_set_highscore_request_proto_rawDescOnce.Do(func() {
		file_set_highscore_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_set_highscore_request_proto_rawDescData)
	})
	return file_set_highscore_request_proto_rawDescData
}

var file_set_highscore_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_set_highscore_request_proto_goTypes = []interface{}{
	(*SetHighscoreRequest)(nil), // 0: protos.SetHighscoreRequest
}
var file_set_highscore_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_set_highscore_request_proto_init() }
func file_set_highscore_request_proto_init() {
	if File_set_highscore_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_set_highscore_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHighscoreRequest); i {
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
			RawDescriptor: file_set_highscore_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_set_highscore_request_proto_goTypes,
		DependencyIndexes: file_set_highscore_request_proto_depIdxs,
		MessageInfos:      file_set_highscore_request_proto_msgTypes,
	}.Build()
	File_set_highscore_request_proto = out.File
	file_set_highscore_request_proto_rawDesc = nil
	file_set_highscore_request_proto_goTypes = nil
	file_set_highscore_request_proto_depIdxs = nil
}
