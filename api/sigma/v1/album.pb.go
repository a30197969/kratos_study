// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/sigma/v1/album.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{0}
}

type CreateAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAlbumReply) Reset() {
	*x = CreateAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumReply) ProtoMessage() {}

func (x *CreateAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumReply.ProtoReflect.Descriptor instead.
func (*CreateAlbumReply) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{1}
}

type UpdateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAlbumRequest) Reset() {
	*x = UpdateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlbumRequest) ProtoMessage() {}

func (x *UpdateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlbumRequest.ProtoReflect.Descriptor instead.
func (*UpdateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{2}
}

type UpdateAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAlbumReply) Reset() {
	*x = UpdateAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlbumReply) ProtoMessage() {}

func (x *UpdateAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlbumReply.ProtoReflect.Descriptor instead.
func (*UpdateAlbumReply) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{3}
}

type DeleteAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlbumRequest) Reset() {
	*x = DeleteAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumRequest) ProtoMessage() {}

func (x *DeleteAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlbumRequest) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{4}
}

type DeleteAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlbumReply) Reset() {
	*x = DeleteAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumReply) ProtoMessage() {}

func (x *DeleteAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumReply.ProtoReflect.Descriptor instead.
func (*DeleteAlbumReply) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{5}
}

type GetAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAlbumRequest) Reset() {
	*x = GetAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumRequest) ProtoMessage() {}

func (x *GetAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumRequest) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{6}
}

func (x *GetAlbumRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *GetAlbumReply) Reset() {
	*x = GetAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumReply) ProtoMessage() {}

func (x *GetAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumReply.ProtoReflect.Descriptor instead.
func (*GetAlbumReply) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{7}
}

func (x *GetAlbumReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAlbumReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAlbumReply) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type ListAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAlbumRequest) Reset() {
	*x = ListAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumRequest) ProtoMessage() {}

func (x *ListAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumRequest.ProtoReflect.Descriptor instead.
func (*ListAlbumRequest) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{8}
}

type ListAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAlbumReply) Reset() {
	*x = ListAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sigma_v1_album_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumReply) ProtoMessage() {}

func (x *ListAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sigma_v1_album_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumReply.ProtoReflect.Descriptor instead.
func (*ListAlbumReply) Descriptor() ([]byte, []int) {
	return file_api_sigma_v1_album_proto_rawDescGZIP(), []int{9}
}

var File_api_sigma_v1_album_proto protoreflect.FileDescriptor

var file_api_sigma_v1_album_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa2, 0x03, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x12, 0x4f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x49, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2e, 0x0a, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1c, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x69, 0x67, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_sigma_v1_album_proto_rawDescOnce sync.Once
	file_api_sigma_v1_album_proto_rawDescData = file_api_sigma_v1_album_proto_rawDesc
)

func file_api_sigma_v1_album_proto_rawDescGZIP() []byte {
	file_api_sigma_v1_album_proto_rawDescOnce.Do(func() {
		file_api_sigma_v1_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sigma_v1_album_proto_rawDescData)
	})
	return file_api_sigma_v1_album_proto_rawDescData
}

var file_api_sigma_v1_album_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_sigma_v1_album_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil), // 0: api.sigma.v1.CreateAlbumRequest
	(*CreateAlbumReply)(nil),   // 1: api.sigma.v1.CreateAlbumReply
	(*UpdateAlbumRequest)(nil), // 2: api.sigma.v1.UpdateAlbumRequest
	(*UpdateAlbumReply)(nil),   // 3: api.sigma.v1.UpdateAlbumReply
	(*DeleteAlbumRequest)(nil), // 4: api.sigma.v1.DeleteAlbumRequest
	(*DeleteAlbumReply)(nil),   // 5: api.sigma.v1.DeleteAlbumReply
	(*GetAlbumRequest)(nil),    // 6: api.sigma.v1.GetAlbumRequest
	(*GetAlbumReply)(nil),      // 7: api.sigma.v1.GetAlbumReply
	(*ListAlbumRequest)(nil),   // 8: api.sigma.v1.ListAlbumRequest
	(*ListAlbumReply)(nil),     // 9: api.sigma.v1.ListAlbumReply
}
var file_api_sigma_v1_album_proto_depIdxs = []int32{
	0, // 0: api.sigma.v1.Album.CreateAlbum:input_type -> api.sigma.v1.CreateAlbumRequest
	2, // 1: api.sigma.v1.Album.UpdateAlbum:input_type -> api.sigma.v1.UpdateAlbumRequest
	4, // 2: api.sigma.v1.Album.DeleteAlbum:input_type -> api.sigma.v1.DeleteAlbumRequest
	6, // 3: api.sigma.v1.Album.GetAlbum:input_type -> api.sigma.v1.GetAlbumRequest
	8, // 4: api.sigma.v1.Album.ListAlbum:input_type -> api.sigma.v1.ListAlbumRequest
	1, // 5: api.sigma.v1.Album.CreateAlbum:output_type -> api.sigma.v1.CreateAlbumReply
	3, // 6: api.sigma.v1.Album.UpdateAlbum:output_type -> api.sigma.v1.UpdateAlbumReply
	5, // 7: api.sigma.v1.Album.DeleteAlbum:output_type -> api.sigma.v1.DeleteAlbumReply
	7, // 8: api.sigma.v1.Album.GetAlbum:output_type -> api.sigma.v1.GetAlbumReply
	9, // 9: api.sigma.v1.Album.ListAlbum:output_type -> api.sigma.v1.ListAlbumReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_sigma_v1_album_proto_init() }
func file_api_sigma_v1_album_proto_init() {
	if File_api_sigma_v1_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sigma_v1_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_api_sigma_v1_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumReply); i {
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
		file_api_sigma_v1_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlbumRequest); i {
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
		file_api_sigma_v1_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlbumReply); i {
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
		file_api_sigma_v1_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumRequest); i {
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
		file_api_sigma_v1_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumReply); i {
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
		file_api_sigma_v1_album_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumRequest); i {
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
		file_api_sigma_v1_album_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumReply); i {
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
		file_api_sigma_v1_album_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumRequest); i {
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
		file_api_sigma_v1_album_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumReply); i {
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
			RawDescriptor: file_api_sigma_v1_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sigma_v1_album_proto_goTypes,
		DependencyIndexes: file_api_sigma_v1_album_proto_depIdxs,
		MessageInfos:      file_api_sigma_v1_album_proto_msgTypes,
	}.Build()
	File_api_sigma_v1_album_proto = out.File
	file_api_sigma_v1_album_proto_rawDesc = nil
	file_api_sigma_v1_album_proto_goTypes = nil
	file_api_sigma_v1_album_proto_depIdxs = nil
}
