// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/color/hsv/v1/conv.proto

package v1

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

type Rgb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Red   float32 `protobuf:"fixed32,1,opt,name=red,proto3" json:"red,omitempty"`
	Green float32 `protobuf:"fixed32,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue  float32 `protobuf:"fixed32,3,opt,name=blue,proto3" json:"blue,omitempty"`
}

func (x *Rgb) Reset() {
	*x = Rgb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rgb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rgb) ProtoMessage() {}

func (x *Rgb) ProtoReflect() protoreflect.Message {
	mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rgb.ProtoReflect.Descriptor instead.
func (*Rgb) Descriptor() ([]byte, []int) {
	return file_proto_color_hsv_v1_conv_proto_rawDescGZIP(), []int{0}
}

func (x *Rgb) GetRed() float32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *Rgb) GetGreen() float32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *Rgb) GetBlue() float32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

type Hsv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hue        float32 `protobuf:"fixed32,1,opt,name=hue,proto3" json:"hue,omitempty"`
	Saturation float32 `protobuf:"fixed32,2,opt,name=saturation,proto3" json:"saturation,omitempty"`
	Value      float32 `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Hsv) Reset() {
	*x = Hsv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hsv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hsv) ProtoMessage() {}

func (x *Hsv) ProtoReflect() protoreflect.Message {
	mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hsv.ProtoReflect.Descriptor instead.
func (*Hsv) Descriptor() ([]byte, []int) {
	return file_proto_color_hsv_v1_conv_proto_rawDescGZIP(), []int{1}
}

func (x *Hsv) GetHue() float32 {
	if x != nil {
		return x.Hue
	}
	return 0
}

func (x *Hsv) GetSaturation() float32 {
	if x != nil {
		return x.Saturation
	}
	return 0
}

func (x *Hsv) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FromHsvRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hsv *Hsv `protobuf:"bytes,1,opt,name=hsv,proto3" json:"hsv,omitempty"`
}

func (x *FromHsvRequest) Reset() {
	*x = FromHsvRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromHsvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromHsvRequest) ProtoMessage() {}

func (x *FromHsvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromHsvRequest.ProtoReflect.Descriptor instead.
func (*FromHsvRequest) Descriptor() ([]byte, []int) {
	return file_proto_color_hsv_v1_conv_proto_rawDescGZIP(), []int{2}
}

func (x *FromHsvRequest) GetHsv() *Hsv {
	if x != nil {
		return x.Hsv
	}
	return nil
}

type FromHsvResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rgb *Rgb `protobuf:"bytes,1,opt,name=rgb,proto3" json:"rgb,omitempty"`
}

func (x *FromHsvResponse) Reset() {
	*x = FromHsvResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromHsvResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromHsvResponse) ProtoMessage() {}

func (x *FromHsvResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_color_hsv_v1_conv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromHsvResponse.ProtoReflect.Descriptor instead.
func (*FromHsvResponse) Descriptor() ([]byte, []int) {
	return file_proto_color_hsv_v1_conv_proto_rawDescGZIP(), []int{3}
}

func (x *FromHsvResponse) GetRgb() *Rgb {
	if x != nil {
		return x.Rgb
	}
	return nil
}

var File_proto_color_hsv_v1_conv_proto protoreflect.FileDescriptor

var file_proto_color_hsv_v1_conv_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2f, 0x68, 0x73,
	0x76, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x68, 0x73, 0x76, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a,
	0x03, 0x52, 0x67, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x62, 0x6c, 0x75, 0x65,
	0x22, 0x4d, 0x0a, 0x03, 0x48, 0x73, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x68, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x61, 0x74,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73,
	0x61, 0x74, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x35, 0x0a, 0x0e, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x73, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x03, 0x68, 0x73, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x68, 0x73, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x73,
	0x76, 0x52, 0x03, 0x68, 0x73, 0x76, 0x22, 0x36, 0x0a, 0x0f, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x73,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x72, 0x67, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x68,
	0x73, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x67, 0x62, 0x52, 0x03, 0x72, 0x67, 0x62, 0x32, 0x5a,
	0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x73, 0x76, 0x12, 0x1c, 0x2e,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x68, 0x73, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x6f,
	0x6d, 0x48, 0x73, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x2e, 0x68, 0x73, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x48,
	0x73, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x68, 0x73,
	0x76, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_color_hsv_v1_conv_proto_rawDescOnce sync.Once
	file_proto_color_hsv_v1_conv_proto_rawDescData = file_proto_color_hsv_v1_conv_proto_rawDesc
)

func file_proto_color_hsv_v1_conv_proto_rawDescGZIP() []byte {
	file_proto_color_hsv_v1_conv_proto_rawDescOnce.Do(func() {
		file_proto_color_hsv_v1_conv_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_color_hsv_v1_conv_proto_rawDescData)
	})
	return file_proto_color_hsv_v1_conv_proto_rawDescData
}

var file_proto_color_hsv_v1_conv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_color_hsv_v1_conv_proto_goTypes = []interface{}{
	(*Rgb)(nil),             // 0: color.hsv.v1.Rgb
	(*Hsv)(nil),             // 1: color.hsv.v1.Hsv
	(*FromHsvRequest)(nil),  // 2: color.hsv.v1.FromHsvRequest
	(*FromHsvResponse)(nil), // 3: color.hsv.v1.FromHsvResponse
}
var file_proto_color_hsv_v1_conv_proto_depIdxs = []int32{
	1, // 0: color.hsv.v1.FromHsvRequest.hsv:type_name -> color.hsv.v1.Hsv
	0, // 1: color.hsv.v1.FromHsvResponse.rgb:type_name -> color.hsv.v1.Rgb
	2, // 2: color.hsv.v1.ColorConvService.FromHsv:input_type -> color.hsv.v1.FromHsvRequest
	3, // 3: color.hsv.v1.ColorConvService.FromHsv:output_type -> color.hsv.v1.FromHsvResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_color_hsv_v1_conv_proto_init() }
func file_proto_color_hsv_v1_conv_proto_init() {
	if File_proto_color_hsv_v1_conv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_color_hsv_v1_conv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rgb); i {
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
		file_proto_color_hsv_v1_conv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hsv); i {
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
		file_proto_color_hsv_v1_conv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromHsvRequest); i {
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
		file_proto_color_hsv_v1_conv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromHsvResponse); i {
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
			RawDescriptor: file_proto_color_hsv_v1_conv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_color_hsv_v1_conv_proto_goTypes,
		DependencyIndexes: file_proto_color_hsv_v1_conv_proto_depIdxs,
		MessageInfos:      file_proto_color_hsv_v1_conv_proto_msgTypes,
	}.Build()
	File_proto_color_hsv_v1_conv_proto = out.File
	file_proto_color_hsv_v1_conv_proto_rawDesc = nil
	file_proto_color_hsv_v1_conv_proto_goTypes = nil
	file_proto_color_hsv_v1_conv_proto_depIdxs = nil
}
