// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.3
// source: system/system.proto

package thirdparty

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_thirdparty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_thirdparty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_thirdparty_thirdparty_proto_rawDescGZIP(), []int{0}
}

type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_thirdparty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_thirdparty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_thirdparty_thirdparty_proto_rawDescGZIP(), []int{1}
}

func (x *SendSMSRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendSMSRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_thirdparty_thirdparty_proto protoreflect.FileDescriptor

var file_thirdparty_thirdparty_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x38, 0x0a, 0x0a,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x4d, 0x53, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thirdparty_thirdparty_proto_rawDescOnce sync.Once
	file_thirdparty_thirdparty_proto_rawDescData = file_thirdparty_thirdparty_proto_rawDesc
)

func file_thirdparty_thirdparty_proto_rawDescGZIP() []byte {
	file_thirdparty_thirdparty_proto_rawDescOnce.Do(func() {
		file_thirdparty_thirdparty_proto_rawDescData = protoimpl.X.CompressGZIP(file_thirdparty_thirdparty_proto_rawDescData)
	})
	return file_thirdparty_thirdparty_proto_rawDescData
}

var file_thirdparty_thirdparty_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thirdparty_thirdparty_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: pb.Empty
	(*SendSMSRequest)(nil), // 1: pb.SendSMSRequest
}
var file_thirdparty_thirdparty_proto_depIdxs = []int32{
	1, // 0: pb.ThirdParty.SendSMS:input_type -> pb.SendSMSRequest
	0, // 1: pb.ThirdParty.SendSMS:output_type -> pb.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thirdparty_thirdparty_proto_init() }
func file_thirdparty_thirdparty_proto_init() {
	if File_thirdparty_thirdparty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thirdparty_thirdparty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_thirdparty_thirdparty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
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
			RawDescriptor: file_thirdparty_thirdparty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thirdparty_thirdparty_proto_goTypes,
		DependencyIndexes: file_thirdparty_thirdparty_proto_depIdxs,
		MessageInfos:      file_thirdparty_thirdparty_proto_msgTypes,
	}.Build()
	File_thirdparty_thirdparty_proto = out.File
	file_thirdparty_thirdparty_proto_rawDesc = nil
	file_thirdparty_thirdparty_proto_goTypes = nil
	file_thirdparty_thirdparty_proto_depIdxs = nil
}
