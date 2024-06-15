// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: api/task/v1/task.proto

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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreateTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTaskRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTaskReply) Reset() {
	*x = CreateTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskReply) ProtoMessage() {}

func (x *CreateTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskReply.ProtoReflect.Descriptor instead.
func (*CreateTaskReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  *bool   `protobuf:"varint,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Name    *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Comment *string `protobuf:"bytes,4,opt,name=comment,proto3,oneof" json:"comment,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *UpdateTaskRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateTaskRequest) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

type UpdateTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected uint32 `protobuf:"varint,1,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
}

func (x *UpdateTaskReply) Reset() {
	*x = UpdateTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskReply) ProtoMessage() {}

func (x *UpdateTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskReply.ProtoReflect.Descriptor instead.
func (*UpdateTaskReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskReply) GetRowsAffected() uint32 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type OverwriteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *OverwriteTaskRequest) Reset() {
	*x = OverwriteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverwriteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteTaskRequest) ProtoMessage() {}

func (x *OverwriteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteTaskRequest.ProtoReflect.Descriptor instead.
func (*OverwriteTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{4}
}

func (x *OverwriteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OverwriteTaskRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *OverwriteTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OverwriteTaskRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type OverwriteTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected uint32 `protobuf:"varint,1,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
}

func (x *OverwriteTaskReply) Reset() {
	*x = OverwriteTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverwriteTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteTaskReply) ProtoMessage() {}

func (x *OverwriteTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteTaskReply.ProtoReflect.Descriptor instead.
func (*OverwriteTaskReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{5}
}

func (x *OverwriteTaskReply) GetRowsAffected() uint32 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected uint32 `protobuf:"varint,1,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
}

func (x *DeleteTaskReply) Reset() {
	*x = DeleteTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskReply) ProtoMessage() {}

func (x *DeleteTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskReply.ProtoReflect.Descriptor instead.
func (*DeleteTaskReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskReply) GetRowsAffected() uint32 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{8}
}

func (x *GetTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Comment   string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	UpdatedBy string `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedBy string `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetTaskReply) Reset() {
	*x = GetTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskReply) ProtoMessage() {}

func (x *GetTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskReply.ProtoReflect.Descriptor instead.
func (*GetTaskReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{9}
}

func (x *GetTaskReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTaskReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetTaskReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTaskReply) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *GetTaskReply) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *GetTaskReply) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *GetTaskReply) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetTaskReply) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SingleTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Comment   string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	UpdatedBy string `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedBy string `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SingleTask) Reset() {
	*x = SingleTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleTask) ProtoMessage() {}

func (x *SingleTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleTask.ProtoReflect.Descriptor instead.
func (*SingleTask) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{10}
}

func (x *SingleTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SingleTask) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SingleTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SingleTask) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SingleTask) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *SingleTask) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *SingleTask) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SingleTask) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTasksRequest) Reset() {
	*x = GetTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequest) ProtoMessage() {}

func (x *GetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{11}
}

type GetTasksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*SingleTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTasksReply) Reset() {
	*x = GetTasksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_task_v1_task_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksReply) ProtoMessage() {}

func (x *GetTasksReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_task_v1_task_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksReply.ProtoReflect.Descriptor instead.
func (*GetTasksReply) Descriptor() ([]byte, []int) {
	return file_api_task_v1_task_proto_rawDescGZIP(), []int{12}
}

func (x *GetTasksReply) GetTasks() []*SingleTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_api_task_v1_task_proto protoreflect.FileDescriptor

var file_api_task_v1_task_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x21,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x14, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x39, 0x0a, 0x12, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73,
	0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x36, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe0, 0x01, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xde, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x32, 0xef, 0x04, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x64, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x12, 0x69, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x32, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a,
	0x0d, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65,
	0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x29, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x73, 0x6f, 0x6e, 0x61, 0x72, 0x2d, 0x62, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_task_v1_task_proto_rawDescOnce sync.Once
	file_api_task_v1_task_proto_rawDescData = file_api_task_v1_task_proto_rawDesc
)

func file_api_task_v1_task_proto_rawDescGZIP() []byte {
	file_api_task_v1_task_proto_rawDescOnce.Do(func() {
		file_api_task_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_task_v1_task_proto_rawDescData)
	})
	return file_api_task_v1_task_proto_rawDescData
}

var file_api_task_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_task_v1_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),    // 0: api.task.v1.CreateTaskRequest
	(*CreateTaskReply)(nil),      // 1: api.task.v1.CreateTaskReply
	(*UpdateTaskRequest)(nil),    // 2: api.task.v1.UpdateTaskRequest
	(*UpdateTaskReply)(nil),      // 3: api.task.v1.UpdateTaskReply
	(*OverwriteTaskRequest)(nil), // 4: api.task.v1.OverwriteTaskRequest
	(*OverwriteTaskReply)(nil),   // 5: api.task.v1.OverwriteTaskReply
	(*DeleteTaskRequest)(nil),    // 6: api.task.v1.DeleteTaskRequest
	(*DeleteTaskReply)(nil),      // 7: api.task.v1.DeleteTaskReply
	(*GetTaskRequest)(nil),       // 8: api.task.v1.GetTaskRequest
	(*GetTaskReply)(nil),         // 9: api.task.v1.GetTaskReply
	(*SingleTask)(nil),           // 10: api.task.v1.SingleTask
	(*GetTasksRequest)(nil),      // 11: api.task.v1.GetTasksRequest
	(*GetTasksReply)(nil),        // 12: api.task.v1.GetTasksReply
}
var file_api_task_v1_task_proto_depIdxs = []int32{
	10, // 0: api.task.v1.GetTasksReply.tasks:type_name -> api.task.v1.SingleTask
	0,  // 1: api.task.v1.Task.CreateTask:input_type -> api.task.v1.CreateTaskRequest
	2,  // 2: api.task.v1.Task.UpdateTask:input_type -> api.task.v1.UpdateTaskRequest
	4,  // 3: api.task.v1.Task.OverwriteTask:input_type -> api.task.v1.OverwriteTaskRequest
	6,  // 4: api.task.v1.Task.DeleteTask:input_type -> api.task.v1.DeleteTaskRequest
	8,  // 5: api.task.v1.Task.GetTask:input_type -> api.task.v1.GetTaskRequest
	11, // 6: api.task.v1.Task.GetTasks:input_type -> api.task.v1.GetTasksRequest
	1,  // 7: api.task.v1.Task.CreateTask:output_type -> api.task.v1.CreateTaskReply
	3,  // 8: api.task.v1.Task.UpdateTask:output_type -> api.task.v1.UpdateTaskReply
	5,  // 9: api.task.v1.Task.OverwriteTask:output_type -> api.task.v1.OverwriteTaskReply
	7,  // 10: api.task.v1.Task.DeleteTask:output_type -> api.task.v1.DeleteTaskReply
	9,  // 11: api.task.v1.Task.GetTask:output_type -> api.task.v1.GetTaskReply
	12, // 12: api.task.v1.Task.GetTasks:output_type -> api.task.v1.GetTasksReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_task_v1_task_proto_init() }
func file_api_task_v1_task_proto_init() {
	if File_api_task_v1_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_task_v1_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTaskRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTaskReply); i {
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
		file_api_task_v1_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTaskReply); i {
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
		file_api_task_v1_task_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OverwriteTaskRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OverwriteTaskReply); i {
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
		file_api_task_v1_task_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskReply); i {
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
		file_api_task_v1_task_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskReply); i {
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
		file_api_task_v1_task_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SingleTask); i {
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
		file_api_task_v1_task_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetTasksRequest); i {
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
		file_api_task_v1_task_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetTasksReply); i {
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
	file_api_task_v1_task_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_task_v1_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_task_v1_task_proto_goTypes,
		DependencyIndexes: file_api_task_v1_task_proto_depIdxs,
		MessageInfos:      file_api_task_v1_task_proto_msgTypes,
	}.Build()
	File_api_task_v1_task_proto = out.File
	file_api_task_v1_task_proto_rawDesc = nil
	file_api_task_v1_task_proto_goTypes = nil
	file_api_task_v1_task_proto_depIdxs = nil
}