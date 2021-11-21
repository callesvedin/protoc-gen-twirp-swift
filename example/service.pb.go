// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: example/cmd/haberdasher/service.proto

//package cmd.haberdasher;

package example

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Hat is a piece of headwear made by a Haberdasher.
type Hat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The size of a hat should always be in inches.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// The color of a hat will never be 'invisible', but other than
	// that, anything is fair game.
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// The name of a hat is it's type. Like, 'bowler', or something.
	Name                  string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedOn             *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	RgbColor              *Color                 `protobuf:"bytes,5,opt,name=rgbColor,proto3" json:"rgbColor,omitempty"`
	AvailableSizes        []*Size                `protobuf:"bytes,6,rep,name=availableSizes,proto3" json:"availableSizes,omitempty"`
	Roles                 []int32                `protobuf:"varint,7,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Dictionary            map[string]int64       `protobuf:"bytes,8,rep,name=dictionary,proto3" json:"dictionary,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DictionaryWithMessage map[string]*Size       `protobuf:"bytes,9,rep,name=dictionaryWithMessage,proto3" json:"dictionaryWithMessage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Hat) Reset() {
	*x = Hat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_cmd_haberdasher_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hat) ProtoMessage() {}

func (x *Hat) ProtoReflect() protoreflect.Message {
	mi := &file_example_cmd_haberdasher_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hat.ProtoReflect.Descriptor instead.
func (*Hat) Descriptor() ([]byte, []int) {
	return file_example_cmd_haberdasher_service_proto_rawDescGZIP(), []int{0}
}

func (x *Hat) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Hat) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Hat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hat) GetCreatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *Hat) GetRgbColor() *Color {
	if x != nil {
		return x.RgbColor
	}
	return nil
}

func (x *Hat) GetAvailableSizes() []*Size {
	if x != nil {
		return x.AvailableSizes
	}
	return nil
}

func (x *Hat) GetRoles() []int32 {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Hat) GetDictionary() map[string]int64 {
	if x != nil {
		return x.Dictionary
	}
	return nil
}

func (x *Hat) GetDictionaryWithMessage() map[string]*Size {
	if x != nil {
		return x.DictionaryWithMessage
	}
	return nil
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Red   int32 `protobuf:"varint,1,opt,name=red,proto3" json:"red,omitempty"`
	Green int32 `protobuf:"varint,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue  int32 `protobuf:"varint,3,opt,name=blue,proto3" json:"blue,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_cmd_haberdasher_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_example_cmd_haberdasher_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_example_cmd_haberdasher_service_proto_rawDescGZIP(), []int{1}
}

func (x *Color) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *Color) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *Color) GetBlue() int32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total float64 `protobuf:"fixed64,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_cmd_haberdasher_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_example_cmd_haberdasher_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_example_cmd_haberdasher_service_proto_rawDescGZIP(), []int{2}
}

func (x *Receipt) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// Size is passed when requesting a new hat to be made. It's always
// measured in inches.
type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inches int32 `protobuf:"varint,1,opt,name=inches,proto3" json:"inches,omitempty"`
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_cmd_haberdasher_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_example_cmd_haberdasher_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_example_cmd_haberdasher_service_proto_rawDescGZIP(), []int{3}
}

func (x *Size) GetInches() int32 {
	if x != nil {
		return x.Inches
	}
	return 0
}

var File_example_cmd_haberdasher_service_proto protoreflect.FileDescriptor

var file_example_cmd_haberdasher_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x68, 0x61,
	0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61,
	0x73, 0x68, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x03, 0x48, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x67, 0x62, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72,
	0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x72, 0x67,
	0x62, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x61,
	0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x61, 0x0a, 0x15, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72,
	0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x0a, 0x0f,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b, 0x0a, 0x1a, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x62,
	0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x43, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a,
	0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x1e,
	0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x32, 0x6b,
	0x0a, 0x0b, 0x48, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x07, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72,
	0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x10, 0x2e, 0x68, 0x61,
	0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x12, 0x2c, 0x0a,
	0x06, 0x42, 0x75, 0x79, 0x48, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64,
	0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x1a, 0x10, 0x2e, 0x68, 0x61, 0x62, 0x65,
	0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x42, 0x19, 0x5a, 0x17, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x68, 0x61, 0x62, 0x65, 0x72,
	0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_cmd_haberdasher_service_proto_rawDescOnce sync.Once
	file_example_cmd_haberdasher_service_proto_rawDescData = file_example_cmd_haberdasher_service_proto_rawDesc
)

func file_example_cmd_haberdasher_service_proto_rawDescGZIP() []byte {
	file_example_cmd_haberdasher_service_proto_rawDescOnce.Do(func() {
		file_example_cmd_haberdasher_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_cmd_haberdasher_service_proto_rawDescData)
	})
	return file_example_cmd_haberdasher_service_proto_rawDescData
}

var file_example_cmd_haberdasher_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_example_cmd_haberdasher_service_proto_goTypes = []interface{}{
	(*Hat)(nil),                   // 0: haberdasher.Hat
	(*Color)(nil),                 // 1: haberdasher.Color
	(*Receipt)(nil),               // 2: haberdasher.Receipt
	(*Size)(nil),                  // 3: haberdasher.Size
	nil,                           // 4: haberdasher.Hat.DictionaryEntry
	nil,                           // 5: haberdasher.Hat.DictionaryWithMessageEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_example_cmd_haberdasher_service_proto_depIdxs = []int32{
	6, // 0: haberdasher.Hat.created_on:type_name -> google.protobuf.Timestamp
	1, // 1: haberdasher.Hat.rgbColor:type_name -> haberdasher.Color
	3, // 2: haberdasher.Hat.availableSizes:type_name -> haberdasher.Size
	4, // 3: haberdasher.Hat.dictionary:type_name -> haberdasher.Hat.DictionaryEntry
	5, // 4: haberdasher.Hat.dictionaryWithMessage:type_name -> haberdasher.Hat.DictionaryWithMessageEntry
	3, // 5: haberdasher.Hat.DictionaryWithMessageEntry.value:type_name -> haberdasher.Size
	3, // 6: haberdasher.Haberdasher.MakeHat:input_type -> haberdasher.Size
	0, // 7: haberdasher.Haberdasher.BuyHat:input_type -> haberdasher.Hat
	0, // 8: haberdasher.Haberdasher.MakeHat:output_type -> haberdasher.Hat
	0, // 9: haberdasher.Haberdasher.BuyHat:output_type -> haberdasher.Hat
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_example_cmd_haberdasher_service_proto_init() }
func file_example_cmd_haberdasher_service_proto_init() {
	if File_example_cmd_haberdasher_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_cmd_haberdasher_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hat); i {
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
		file_example_cmd_haberdasher_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_example_cmd_haberdasher_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_example_cmd_haberdasher_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
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
			RawDescriptor: file_example_cmd_haberdasher_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_cmd_haberdasher_service_proto_goTypes,
		DependencyIndexes: file_example_cmd_haberdasher_service_proto_depIdxs,
		MessageInfos:      file_example_cmd_haberdasher_service_proto_msgTypes,
	}.Build()
	File_example_cmd_haberdasher_service_proto = out.File
	file_example_cmd_haberdasher_service_proto_rawDesc = nil
	file_example_cmd_haberdasher_service_proto_goTypes = nil
	file_example_cmd_haberdasher_service_proto_depIdxs = nil
}