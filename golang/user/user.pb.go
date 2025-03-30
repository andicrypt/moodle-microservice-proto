// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: user/user.proto

package user

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

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fullname      string                 `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	ProfileUrl    string                 `protobuf:"bytes,2,opt,name=profile_url,json=profileUrl,proto3" json:"profile_url,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateUserRequest) GetProfileUrl() string {
	if x != nil {
		return x.ProfileUrl
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fullname      string                 `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	ProfileUrl    string                 `protobuf:"bytes,2,opt,name=profile_url,json=profileUrl,proto3" json:"profile_url,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserResponse) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *GetUserResponse) GetProfileUrl() string {
	if x != nil {
		return x.ProfileUrl
	}
	return ""
}

func (x *GetUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_user_user_proto protoreflect.FileDescriptor

const file_user_user_proto_rawDesc = "" +
	"\n" +
	"\x0fuser/user.proto\"f\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\bfullname\x18\x01 \x01(\tR\bfullname\x12\x1f\n" +
	"\vprofile_url\x18\x02 \x01(\tR\n" +
	"profileUrl\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"-\n" +
	"\x12CreateUserResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\")\n" +
	"\x0eGetUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"d\n" +
	"\x0fGetUserResponse\x12\x1a\n" +
	"\bfullname\x18\x01 \x01(\tR\bfullname\x12\x1f\n" +
	"\vprofile_url\x18\x02 \x01(\tR\n" +
	"profileUrl\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email2h\n" +
	"\x04User\x123\n" +
	"\x06Create\x12\x12.CreateUserRequest\x1a\x13.CreateUserResponse\"\x00\x12+\n" +
	"\x03Get\x12\x10.GetUserResponse\x1a\x10.GetUserResponse\"\x00B5Z3github.com/andicrypt/moodle-microservice-proto/userb\x06proto3"

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData []byte
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)))
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: CreateUserRequest
	(*CreateUserResponse)(nil), // 1: CreateUserResponse
	(*GetUserRequest)(nil),     // 2: GetUserRequest
	(*GetUserResponse)(nil),    // 3: GetUserResponse
}
var file_user_user_proto_depIdxs = []int32{
	0, // 0: User.Create:input_type -> CreateUserRequest
	3, // 1: User.Get:input_type -> GetUserResponse
	1, // 2: User.Create:output_type -> CreateUserResponse
	3, // 3: User.Get:output_type -> GetUserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
