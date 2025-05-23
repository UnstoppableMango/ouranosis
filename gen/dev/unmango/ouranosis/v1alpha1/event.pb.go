// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: dev/unmango/ouranosis/v1alpha1/event.proto

package ouranosisv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientEvent) Reset() {
	*x = ClientEvent{}
	mi := &file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent) ProtoMessage() {}

func (x *ClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent.ProtoReflect.Descriptor instead.
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescGZIP(), []int{0}
}

type ServerEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerEvent) Reset() {
	*x = ServerEvent{}
	mi := &file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEvent) ProtoMessage() {}

func (x *ServerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEvent.ProtoReflect.Descriptor instead.
func (*ServerEvent) Descriptor() ([]byte, []int) {
	return file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescGZIP(), []int{1}
}

var File_dev_unmango_ouranosis_v1alpha1_event_proto protoreflect.FileDescriptor

const file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDesc = "" +
	"\n" +
	"*dev/unmango/ouranosis/v1alpha1/event.proto\x12\x1edev.unmango.ouranosis.v1alpha1\"\r\n" +
	"\vClientEvent\"\r\n" +
	"\vServerEventB\xa3\x02\n" +
	"\"com.dev.unmango.ouranosis.v1alpha1B\n" +
	"EventProtoP\x01ZVgithub.com/unstoppablemango/ouranosis/dev/unmango/ouranosis/v1alpha1;ouranosisv1alpha1\xa2\x02\x03DUO\xaa\x02\x1eDev.Unmango.Ouranosis.V1alpha1\xca\x02\x1eDev\\Unmango\\Ouranosis\\V1alpha1\xe2\x02*Dev\\Unmango\\Ouranosis\\V1alpha1\\GPBMetadata\xea\x02!Dev::Unmango::Ouranosis::V1alpha1b\x06proto3"

var (
	file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescOnce sync.Once
	file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescData []byte
)

func file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescGZIP() []byte {
	file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescOnce.Do(func() {
		file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDesc), len(file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDesc)))
	})
	return file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDescData
}

var file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dev_unmango_ouranosis_v1alpha1_event_proto_goTypes = []any{
	(*ClientEvent)(nil), // 0: dev.unmango.ouranosis.v1alpha1.ClientEvent
	(*ServerEvent)(nil), // 1: dev.unmango.ouranosis.v1alpha1.ServerEvent
}
var file_dev_unmango_ouranosis_v1alpha1_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dev_unmango_ouranosis_v1alpha1_event_proto_init() }
func file_dev_unmango_ouranosis_v1alpha1_event_proto_init() {
	if File_dev_unmango_ouranosis_v1alpha1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDesc), len(file_dev_unmango_ouranosis_v1alpha1_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dev_unmango_ouranosis_v1alpha1_event_proto_goTypes,
		DependencyIndexes: file_dev_unmango_ouranosis_v1alpha1_event_proto_depIdxs,
		MessageInfos:      file_dev_unmango_ouranosis_v1alpha1_event_proto_msgTypes,
	}.Build()
	File_dev_unmango_ouranosis_v1alpha1_event_proto = out.File
	file_dev_unmango_ouranosis_v1alpha1_event_proto_goTypes = nil
	file_dev_unmango_ouranosis_v1alpha1_event_proto_depIdxs = nil
}
