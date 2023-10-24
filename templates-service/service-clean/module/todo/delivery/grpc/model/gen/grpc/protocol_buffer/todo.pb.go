// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: todo.proto

package protocol_buffer

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

type Status int32

const (
	Status_Todo       Status = 0
	Status_InProgress Status = 1
	Status_Done       Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Todo",
		1: "InProgress",
		2: "Done",
	}
	Status_value = map[string]int32{
		"Todo":       0,
		"InProgress": 1,
		"Done":       2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_todo_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

type TodoCreateRequestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Status      Status `protobuf:"varint,2,opt,name=status,proto3,enum=todo.todo.Status" json:"status,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TodoDate    string `protobuf:"bytes,4,opt,name=todoDate,proto3" json:"todoDate,omitempty"` // time RFC3339 format
}

func (x *TodoCreateRequestModel) Reset() {
	*x = TodoCreateRequestModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreateRequestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreateRequestModel) ProtoMessage() {}

func (x *TodoCreateRequestModel) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreateRequestModel.ProtoReflect.Descriptor instead.
func (*TodoCreateRequestModel) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoCreateRequestModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoCreateRequestModel) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Todo
}

func (x *TodoCreateRequestModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoCreateRequestModel) GetTodoDate() string {
	if x != nil {
		return x.TodoDate
	}
	return ""
}

type TodoCreateResponseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodoCreateResponseModel) Reset() {
	*x = TodoCreateResponseModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreateResponseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreateResponseModel) ProtoMessage() {}

func (x *TodoCreateResponseModel) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreateResponseModel.ProtoReflect.Descriptor instead.
func (*TodoCreateResponseModel) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x54, 0x6f, 0x64, 0x6f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x19, 0x0a, 0x17, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2a, 0x2c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x32, 0x69, 0x0a, 0x0b, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x21,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x22, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_todo_proto_goTypes = []interface{}{
	(Status)(0),                     // 0: todo.todo.status
	(*TodoCreateRequestModel)(nil),  // 1: todo.todo.TodoCreateRequestModel
	(*TodoCreateResponseModel)(nil), // 2: todo.todo.TodoCreateResponseModel
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: todo.todo.TodoCreateRequestModel.status:type_name -> todo.todo.status
	1, // 1: todo.todo.TodoService.CreateTodoHandler:input_type -> todo.todo.TodoCreateRequestModel
	2, // 2: todo.todo.TodoService.CreateTodoHandler:output_type -> todo.todo.TodoCreateResponseModel
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreateRequestModel); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreateResponseModel); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		EnumInfos:         file_todo_proto_enumTypes,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
