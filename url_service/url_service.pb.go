// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: url_service/url_service.proto

package url_service

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

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_service_url_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_url_service_url_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_url_service_url_service_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_url_service_url_service_proto protoreflect.FileDescriptor

var file_url_service_url_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x18, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x45, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x11, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6d, 0x77, 0x6f,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x6c, 0x75, 0x72, 0x6c, 0x2f, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_service_url_service_proto_rawDescOnce sync.Once
	file_url_service_url_service_proto_rawDescData = file_url_service_url_service_proto_rawDesc
)

func file_url_service_url_service_proto_rawDescGZIP() []byte {
	file_url_service_url_service_proto_rawDescOnce.Do(func() {
		file_url_service_url_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_service_url_service_proto_rawDescData)
	})
	return file_url_service_url_service_proto_rawDescData
}

var file_url_service_url_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_url_service_url_service_proto_goTypes = []interface{}{
	(*Link)(nil), // 0: url_service.Link
}
var file_url_service_url_service_proto_depIdxs = []int32{
	0, // 0: url_service.URLService.GetShortenURL:input_type -> url_service.Link
	0, // 1: url_service.URLService.GetShortenURL:output_type -> url_service.Link
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_service_url_service_proto_init() }
func file_url_service_url_service_proto_init() {
	if File_url_service_url_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_service_url_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_url_service_url_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_service_url_service_proto_goTypes,
		DependencyIndexes: file_url_service_url_service_proto_depIdxs,
		MessageInfos:      file_url_service_url_service_proto_msgTypes,
	}.Build()
	File_url_service_url_service_proto = out.File
	file_url_service_url_service_proto_rawDesc = nil
	file_url_service_url_service_proto_goTypes = nil
	file_url_service_url_service_proto_depIdxs = nil
}
