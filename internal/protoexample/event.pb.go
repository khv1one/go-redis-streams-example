// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: internal/protoexample/event.proto

package protoexample

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Foo      int64     `protobuf:"varint,3,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar      int64     `protobuf:"varint,4,opt,name=bar,proto3" json:"bar,omitempty"`
	SubEvent *SubEvent `protobuf:"bytes,5,opt,name=sub_event,json=subEvent,proto3" json:"sub_event,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protoexample_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protoexample_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internal_protoexample_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetFoo() int64 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *Event) GetBar() int64 {
	if x != nil {
		return x.Bar
	}
	return 0
}

func (x *Event) GetSubEvent() *SubEvent {
	if x != nil {
		return x.SubEvent
	}
	return nil
}

type SubEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barbar string       `protobuf:"bytes,1,opt,name=barbar,proto3" json:"barbar,omitempty"`
	Foofoo *SubSubEvent `protobuf:"bytes,2,opt,name=foofoo,proto3" json:"foofoo,omitempty"`
}

func (x *SubEvent) Reset() {
	*x = SubEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protoexample_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubEvent) ProtoMessage() {}

func (x *SubEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protoexample_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubEvent.ProtoReflect.Descriptor instead.
func (*SubEvent) Descriptor() ([]byte, []int) {
	return file_internal_protoexample_event_proto_rawDescGZIP(), []int{1}
}

func (x *SubEvent) GetBarbar() string {
	if x != nil {
		return x.Barbar
	}
	return ""
}

func (x *SubEvent) GetFoofoo() *SubSubEvent {
	if x != nil {
		return x.Foofoo
	}
	return nil
}

type SubSubEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foofoofoo int64 `protobuf:"varint,1,opt,name=foofoofoo,proto3" json:"foofoofoo,omitempty"`
}

func (x *SubSubEvent) Reset() {
	*x = SubSubEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protoexample_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubSubEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubSubEvent) ProtoMessage() {}

func (x *SubSubEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protoexample_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubSubEvent.ProtoReflect.Descriptor instead.
func (*SubSubEvent) Descriptor() ([]byte, []int) {
	return file_internal_protoexample_event_proto_rawDescGZIP(), []int{2}
}

func (x *SubSubEvent) GetFoofoofoo() int64 {
	if x != nil {
		return x.Foofoofoo
	}
	return 0
}

var File_internal_protoexample_event_proto protoreflect.FileDescriptor

var file_internal_protoexample_event_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x6f, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x62, 0x61, 0x72, 0x12,
	0x26, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x75, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x72, 0x62, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x72, 0x62, 0x61, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x66,
	0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x75,
	0x62, 0x53, 0x75, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x66, 0x6f,
	0x6f, 0x22, 0x2b, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x53, 0x75, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x42, 0x17,
	0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protoexample_event_proto_rawDescOnce sync.Once
	file_internal_protoexample_event_proto_rawDescData = file_internal_protoexample_event_proto_rawDesc
)

func file_internal_protoexample_event_proto_rawDescGZIP() []byte {
	file_internal_protoexample_event_proto_rawDescOnce.Do(func() {
		file_internal_protoexample_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protoexample_event_proto_rawDescData)
	})
	return file_internal_protoexample_event_proto_rawDescData
}

var file_internal_protoexample_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_protoexample_event_proto_goTypes = []interface{}{
	(*Event)(nil),       // 0: Event
	(*SubEvent)(nil),    // 1: SubEvent
	(*SubSubEvent)(nil), // 2: SubSubEvent
}
var file_internal_protoexample_event_proto_depIdxs = []int32{
	1, // 0: Event.sub_event:type_name -> SubEvent
	2, // 1: SubEvent.foofoo:type_name -> SubSubEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_protoexample_event_proto_init() }
func file_internal_protoexample_event_proto_init() {
	if File_internal_protoexample_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protoexample_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_internal_protoexample_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubEvent); i {
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
		file_internal_protoexample_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubSubEvent); i {
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
			RawDescriptor: file_internal_protoexample_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protoexample_event_proto_goTypes,
		DependencyIndexes: file_internal_protoexample_event_proto_depIdxs,
		MessageInfos:      file_internal_protoexample_event_proto_msgTypes,
	}.Build()
	File_internal_protoexample_event_proto = out.File
	file_internal_protoexample_event_proto_rawDesc = nil
	file_internal_protoexample_event_proto_goTypes = nil
	file_internal_protoexample_event_proto_depIdxs = nil
}