// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: file_sync.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DownloadInfoRequestType int32

const (
	DownloadInfoRequestType_All   DownloadInfoRequestType = 0
	DownloadInfoRequestType_Queue DownloadInfoRequestType = 1
	DownloadInfoRequestType_Temp  DownloadInfoRequestType = 2
	DownloadInfoRequestType_Sync  DownloadInfoRequestType = 3
)

// Enum value maps for DownloadInfoRequestType.
var (
	DownloadInfoRequestType_name = map[int32]string{
		0: "All",
		1: "Queue",
		2: "Temp",
		3: "Sync",
	}
	DownloadInfoRequestType_value = map[string]int32{
		"All":   0,
		"Queue": 1,
		"Temp":  2,
		"Sync":  3,
	}
)

func (x DownloadInfoRequestType) Enum() *DownloadInfoRequestType {
	p := new(DownloadInfoRequestType)
	*p = x
	return p
}

func (x DownloadInfoRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownloadInfoRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_file_sync_proto_enumTypes[0].Descriptor()
}

func (DownloadInfoRequestType) Type() protoreflect.EnumType {
	return &file_file_sync_proto_enumTypes[0]
}

func (x DownloadInfoRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownloadInfoRequestType.Descriptor instead.
func (DownloadInfoRequestType) EnumDescriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{0}
}

type DownloadTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkingUrl   string                  `protobuf:"bytes,1,opt,name=workingUrl,proto3" json:"workingUrl,omitempty"`
	FileID       string                  `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
	SavePath     string                  `protobuf:"bytes,3,opt,name=savePath,proto3" json:"savePath,omitempty"`
	FileName     string                  `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Cookie       string                  `protobuf:"bytes,5,opt,name=cookie,proto3" json:"cookie,omitempty"`
	DownLoadType DownloadInfoRequestType `protobuf:"varint,6,opt,name=downLoadType,proto3,enum=DownloadInfoRequestType" json:"downLoadType,omitempty"`
}

func (x *DownloadTaskRequest) Reset() {
	*x = DownloadTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTaskRequest) ProtoMessage() {}

func (x *DownloadTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTaskRequest.ProtoReflect.Descriptor instead.
func (*DownloadTaskRequest) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadTaskRequest) GetWorkingUrl() string {
	if x != nil {
		return x.WorkingUrl
	}
	return ""
}

func (x *DownloadTaskRequest) GetFileID() string {
	if x != nil {
		return x.FileID
	}
	return ""
}

func (x *DownloadTaskRequest) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

func (x *DownloadTaskRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DownloadTaskRequest) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

func (x *DownloadTaskRequest) GetDownLoadType() DownloadInfoRequestType {
	if x != nil {
		return x.DownLoadType
	}
	return DownloadInfoRequestType_All
}

type DownloadDirTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkingUrl   string                  `protobuf:"bytes,1,opt,name=workingUrl,proto3" json:"workingUrl,omitempty"`
	DirPath      string                  `protobuf:"bytes,2,opt,name=dirPath,proto3" json:"dirPath,omitempty"`
	SavePath     string                  `protobuf:"bytes,3,opt,name=savePath,proto3" json:"savePath,omitempty"`
	Cookie       string                  `protobuf:"bytes,5,opt,name=cookie,proto3" json:"cookie,omitempty"`
	DownLoadType DownloadInfoRequestType `protobuf:"varint,6,opt,name=downLoadType,proto3,enum=DownloadInfoRequestType" json:"downLoadType,omitempty"`
}

func (x *DownloadDirTaskRequest) Reset() {
	*x = DownloadDirTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadDirTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadDirTaskRequest) ProtoMessage() {}

func (x *DownloadDirTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadDirTaskRequest.ProtoReflect.Descriptor instead.
func (*DownloadDirTaskRequest) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadDirTaskRequest) GetWorkingUrl() string {
	if x != nil {
		return x.WorkingUrl
	}
	return ""
}

func (x *DownloadDirTaskRequest) GetDirPath() string {
	if x != nil {
		return x.DirPath
	}
	return ""
}

func (x *DownloadDirTaskRequest) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

func (x *DownloadDirTaskRequest) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

func (x *DownloadDirTaskRequest) GetDownLoadType() DownloadInfoRequestType {
	if x != nil {
		return x.DownLoadType
	}
	return DownloadInfoRequestType_All
}

type DownloadTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DownloadTaskResult) Reset() {
	*x = DownloadTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTaskResult) ProtoMessage() {}

func (x *DownloadTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTaskResult.ProtoReflect.Descriptor instead.
func (*DownloadTaskResult) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadTaskResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DownloadDirTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DownloadDirTaskResult) Reset() {
	*x = DownloadDirTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadDirTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadDirTaskResult) ProtoMessage() {}

func (x *DownloadDirTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadDirTaskResult.ProtoReflect.Descriptor instead.
func (*DownloadDirTaskResult) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadDirTaskResult) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DownloadTaskCancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DownloadTaskCancelRequest) Reset() {
	*x = DownloadTaskCancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTaskCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTaskCancelRequest) ProtoMessage() {}

func (x *DownloadTaskCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTaskCancelRequest.ProtoReflect.Descriptor instead.
func (*DownloadTaskCancelRequest) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadTaskCancelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DownloadTaskCancelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DownloadTaskCancelResult) Reset() {
	*x = DownloadTaskCancelResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTaskCancelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTaskCancelResult) ProtoMessage() {}

func (x *DownloadTaskCancelResult) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTaskCancelResult.ProtoReflect.Descriptor instead.
func (*DownloadTaskCancelResult) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadTaskCancelResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DownloadInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *string                 `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Type DownloadInfoRequestType `protobuf:"varint,2,opt,name=type,proto3,enum=DownloadInfoRequestType" json:"type,omitempty"`
}

func (x *DownloadInfoRequest) Reset() {
	*x = DownloadInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadInfoRequest) ProtoMessage() {}

func (x *DownloadInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadInfoRequest.ProtoReflect.Descriptor instead.
func (*DownloadInfoRequest) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadInfoRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *DownloadInfoRequest) GetType() DownloadInfoRequestType {
	if x != nil {
		return x.Type
	}
	return DownloadInfoRequestType_All
}

type DownLoadInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type DownloadInfoRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=DownloadInfoRequestType" json:"type,omitempty"`
	Data []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownLoadInfoResult) Reset() {
	*x = DownLoadInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_sync_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownLoadInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownLoadInfoResult) ProtoMessage() {}

func (x *DownLoadInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_file_sync_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownLoadInfoResult.ProtoReflect.Descriptor instead.
func (*DownLoadInfoResult) Descriptor() ([]byte, []int) {
	return file_file_sync_proto_rawDescGZIP(), []int{7}
}

func (x *DownLoadInfoResult) GetType() DownloadInfoRequestType {
	if x != nil {
		return x.Type
	}
	return DownloadInfoRequestType_All
}

func (x *DownLoadInfoResult) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_file_sync_proto protoreflect.FileDescriptor

var file_file_sync_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xc4, 0x01, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x72,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e,
	0x4c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x4c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x15,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x18, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5f, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x12, 0x44, 0x6f, 0x77,
	0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x41, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x79,
	0x6e, 0x63, 0x10, 0x03, 0x32, 0xf8, 0x02, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x6e,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x44, 0x69,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x69, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x14, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x6f,
	0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x6f, 0x77, 0x6e,
	0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_sync_proto_rawDescOnce sync.Once
	file_file_sync_proto_rawDescData = file_file_sync_proto_rawDesc
)

func file_file_sync_proto_rawDescGZIP() []byte {
	file_file_sync_proto_rawDescOnce.Do(func() {
		file_file_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_sync_proto_rawDescData)
	})
	return file_file_sync_proto_rawDescData
}

var file_file_sync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_file_sync_proto_goTypes = []interface{}{
	(DownloadInfoRequestType)(0),      // 0: DownloadInfoRequestType
	(*DownloadTaskRequest)(nil),       // 1: DownloadTaskRequest
	(*DownloadDirTaskRequest)(nil),    // 2: DownloadDirTaskRequest
	(*DownloadTaskResult)(nil),        // 3: DownloadTaskResult
	(*DownloadDirTaskResult)(nil),     // 4: DownloadDirTaskResult
	(*DownloadTaskCancelRequest)(nil), // 5: DownloadTaskCancelRequest
	(*DownloadTaskCancelResult)(nil),  // 6: DownloadTaskCancelResult
	(*DownloadInfoRequest)(nil),       // 7: DownloadInfoRequest
	(*DownLoadInfoResult)(nil),        // 8: DownLoadInfoResult
}
var file_file_sync_proto_depIdxs = []int32{
	0, // 0: DownloadTaskRequest.downLoadType:type_name -> DownloadInfoRequestType
	0, // 1: DownloadDirTaskRequest.downLoadType:type_name -> DownloadInfoRequestType
	0, // 2: DownloadInfoRequest.type:type_name -> DownloadInfoRequestType
	0, // 3: DownLoadInfoResult.type:type_name -> DownloadInfoRequestType
	1, // 4: FileSyncService.AddDownloadTask:input_type -> DownloadTaskRequest
	2, // 5: FileSyncService.AddDownloadTasksByDirPath:input_type -> DownloadDirTaskRequest
	7, // 6: FileSyncService.GetDownloadInfoStream:input_type -> DownloadInfoRequest
	7, // 7: FileSyncService.GetDownloadInfo:input_type -> DownloadInfoRequest
	5, // 8: FileSyncService.CancelDownloadTask:input_type -> DownloadTaskCancelRequest
	3, // 9: FileSyncService.AddDownloadTask:output_type -> DownloadTaskResult
	4, // 10: FileSyncService.AddDownloadTasksByDirPath:output_type -> DownloadDirTaskResult
	8, // 11: FileSyncService.GetDownloadInfoStream:output_type -> DownLoadInfoResult
	8, // 12: FileSyncService.GetDownloadInfo:output_type -> DownLoadInfoResult
	6, // 13: FileSyncService.CancelDownloadTask:output_type -> DownloadTaskCancelResult
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_file_sync_proto_init() }
func file_file_sync_proto_init() {
	if File_file_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTaskRequest); i {
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
		file_file_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadDirTaskRequest); i {
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
		file_file_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTaskResult); i {
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
		file_file_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadDirTaskResult); i {
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
		file_file_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTaskCancelRequest); i {
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
		file_file_sync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTaskCancelResult); i {
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
		file_file_sync_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadInfoRequest); i {
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
		file_file_sync_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownLoadInfoResult); i {
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
	file_file_sync_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_sync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_sync_proto_goTypes,
		DependencyIndexes: file_file_sync_proto_depIdxs,
		EnumInfos:         file_file_sync_proto_enumTypes,
		MessageInfos:      file_file_sync_proto_msgTypes,
	}.Build()
	File_file_sync_proto = out.File
	file_file_sync_proto_rawDesc = nil
	file_file_sync_proto_goTypes = nil
	file_file_sync_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileSyncServiceClient is the client API for FileSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileSyncServiceClient interface {
	AddDownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error)
	AddDownloadTasksByDirPath(ctx context.Context, in *DownloadDirTaskRequest, opts ...grpc.CallOption) (*DownloadDirTaskResult, error)
	GetDownloadInfoStream(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadInfoStreamClient, error)
	GetDownloadInfo(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (*DownLoadInfoResult, error)
	CancelDownloadTask(ctx context.Context, in *DownloadTaskCancelRequest, opts ...grpc.CallOption) (*DownloadTaskCancelResult, error)
}

type fileSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileSyncServiceClient(cc grpc.ClientConnInterface) FileSyncServiceClient {
	return &fileSyncServiceClient{cc}
}

func (c *fileSyncServiceClient) AddDownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error) {
	out := new(DownloadTaskResult)
	err := c.cc.Invoke(ctx, "/FileSyncService/AddDownloadTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) AddDownloadTasksByDirPath(ctx context.Context, in *DownloadDirTaskRequest, opts ...grpc.CallOption) (*DownloadDirTaskResult, error) {
	out := new(DownloadDirTaskResult)
	err := c.cc.Invoke(ctx, "/FileSyncService/AddDownloadTasksByDirPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) GetDownloadInfoStream(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadInfoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSyncService_serviceDesc.Streams[0], "/FileSyncService/GetDownloadInfoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceGetDownloadInfoStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_GetDownloadInfoStreamClient interface {
	Recv() (*DownLoadInfoResult, error)
	grpc.ClientStream
}

type fileSyncServiceGetDownloadInfoStreamClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceGetDownloadInfoStreamClient) Recv() (*DownLoadInfoResult, error) {
	m := new(DownLoadInfoResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncServiceClient) GetDownloadInfo(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (*DownLoadInfoResult, error) {
	out := new(DownLoadInfoResult)
	err := c.cc.Invoke(ctx, "/FileSyncService/GetDownloadInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) CancelDownloadTask(ctx context.Context, in *DownloadTaskCancelRequest, opts ...grpc.CallOption) (*DownloadTaskCancelResult, error) {
	out := new(DownloadTaskCancelResult)
	err := c.cc.Invoke(ctx, "/FileSyncService/CancelDownloadTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileSyncServiceServer is the server API for FileSyncService service.
type FileSyncServiceServer interface {
	AddDownloadTask(context.Context, *DownloadTaskRequest) (*DownloadTaskResult, error)
	AddDownloadTasksByDirPath(context.Context, *DownloadDirTaskRequest) (*DownloadDirTaskResult, error)
	GetDownloadInfoStream(*DownloadInfoRequest, FileSyncService_GetDownloadInfoStreamServer) error
	GetDownloadInfo(context.Context, *DownloadInfoRequest) (*DownLoadInfoResult, error)
	CancelDownloadTask(context.Context, *DownloadTaskCancelRequest) (*DownloadTaskCancelResult, error)
}

// UnimplementedFileSyncServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileSyncServiceServer struct {
}

func (*UnimplementedFileSyncServiceServer) AddDownloadTask(context.Context, *DownloadTaskRequest) (*DownloadTaskResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownloadTask not implemented")
}
func (*UnimplementedFileSyncServiceServer) AddDownloadTasksByDirPath(context.Context, *DownloadDirTaskRequest) (*DownloadDirTaskResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownloadTasksByDirPath not implemented")
}
func (*UnimplementedFileSyncServiceServer) GetDownloadInfoStream(*DownloadInfoRequest, FileSyncService_GetDownloadInfoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDownloadInfoStream not implemented")
}
func (*UnimplementedFileSyncServiceServer) GetDownloadInfo(context.Context, *DownloadInfoRequest) (*DownLoadInfoResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadInfo not implemented")
}
func (*UnimplementedFileSyncServiceServer) CancelDownloadTask(context.Context, *DownloadTaskCancelRequest) (*DownloadTaskCancelResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDownloadTask not implemented")
}

func RegisterFileSyncServiceServer(s *grpc.Server, srv FileSyncServiceServer) {
	s.RegisterService(&_FileSyncService_serviceDesc, srv)
}

func _FileSyncService_AddDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).AddDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileSyncService/AddDownloadTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).AddDownloadTask(ctx, req.(*DownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_AddDownloadTasksByDirPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadDirTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).AddDownloadTasksByDirPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileSyncService/AddDownloadTasksByDirPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).AddDownloadTasksByDirPath(ctx, req.(*DownloadDirTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_GetDownloadInfoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).GetDownloadInfoStream(m, &fileSyncServiceGetDownloadInfoStreamServer{stream})
}

type FileSyncService_GetDownloadInfoStreamServer interface {
	Send(*DownLoadInfoResult) error
	grpc.ServerStream
}

type fileSyncServiceGetDownloadInfoStreamServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceGetDownloadInfoStreamServer) Send(m *DownLoadInfoResult) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSyncService_GetDownloadInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).GetDownloadInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileSyncService/GetDownloadInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).GetDownloadInfo(ctx, req.(*DownloadInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_CancelDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTaskCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).CancelDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileSyncService/CancelDownloadTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).CancelDownloadTask(ctx, req.(*DownloadTaskCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileSyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileSyncService",
	HandlerType: (*FileSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDownloadTask",
			Handler:    _FileSyncService_AddDownloadTask_Handler,
		},
		{
			MethodName: "AddDownloadTasksByDirPath",
			Handler:    _FileSyncService_AddDownloadTasksByDirPath_Handler,
		},
		{
			MethodName: "GetDownloadInfo",
			Handler:    _FileSyncService_GetDownloadInfo_Handler,
		},
		{
			MethodName: "CancelDownloadTask",
			Handler:    _FileSyncService_CancelDownloadTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDownloadInfoStream",
			Handler:       _FileSyncService_GetDownloadInfoStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_sync.proto",
}
