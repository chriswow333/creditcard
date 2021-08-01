// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: bank.proto

package bank

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

type BankPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *BankPayload) Reset() {
	*x = BankPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankPayload) ProtoMessage() {}

func (x *BankPayload) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankPayload.ProtoReflect.Descriptor instead.
func (*BankPayload) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *BankPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankPayload) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type BankCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *BankPayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *BankCreateRequest) Reset() {
	*x = BankCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCreateRequest) ProtoMessage() {}

func (x *BankCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCreateRequest.ProtoReflect.Descriptor instead.
func (*BankCreateRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *BankCreateRequest) GetPayload() *BankPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type BankCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BankCreateReply) Reset() {
	*x = BankCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCreateReply) ProtoMessage() {}

func (x *BankCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCreateReply.ProtoReflect.Descriptor instead.
func (*BankCreateReply) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *BankCreateReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BankCreateReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BankUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *BankPayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *BankUpdateRequest) Reset() {
	*x = BankUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankUpdateRequest) ProtoMessage() {}

func (x *BankUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankUpdateRequest.ProtoReflect.Descriptor instead.
func (*BankUpdateRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{3}
}

func (x *BankUpdateRequest) GetPayload() *BankPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type BankUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BankUpdateReply) Reset() {
	*x = BankUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankUpdateReply) ProtoMessage() {}

func (x *BankUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankUpdateReply.ProtoReflect.Descriptor instead.
func (*BankUpdateReply) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{4}
}

func (x *BankUpdateReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BankUpdateReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BankGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BankGetAllRequest) Reset() {
	*x = BankGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankGetAllRequest) ProtoMessage() {}

func (x *BankGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankGetAllRequest.ProtoReflect.Descriptor instead.
func (*BankGetAllRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{5}
}

type BankGetAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Banks   []*BankPayload `protobuf:"bytes,3,rep,name=banks,proto3" json:"banks,omitempty"`
}

func (x *BankGetAllReply) Reset() {
	*x = BankGetAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankGetAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankGetAllReply) ProtoMessage() {}

func (x *BankGetAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankGetAllReply.ProtoReflect.Descriptor instead.
func (*BankGetAllReply) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{6}
}

func (x *BankGetAllReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BankGetAllReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BankGetAllReply) GetBanks() []*BankPayload {
	if x != nil {
		return x.Banks
	}
	return nil
}

type BankGetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BankGetByIDRequest) Reset() {
	*x = BankGetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankGetByIDRequest) ProtoMessage() {}

func (x *BankGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankGetByIDRequest.ProtoReflect.Descriptor instead.
func (*BankGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{7}
}

func (x *BankGetByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BankGetByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool         `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Bank    *BankPayload `protobuf:"bytes,3,opt,name=bank,proto3" json:"bank,omitempty"`
}

func (x *BankGetByIDReply) Reset() {
	*x = BankGetByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankGetByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankGetByIDReply) ProtoMessage() {}

func (x *BankGetByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankGetByIDReply.ProtoReflect.Descriptor instead.
func (*BankGetByIDReply) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{8}
}

func (x *BankGetByIDReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BankGetByIDReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BankGetByIDReply) GetBank() *BankPayload {
	if x != nil {
		return x.Bank
	}
	return nil
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x6e, 0x6b, 0x22, 0x45, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x11, 0x42, 0x61, 0x6e,
	0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x40, 0x0a, 0x11, 0x42, 0x61,
	0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3d, 0x0a, 0x0f,
	0x42, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x13, 0x0a, 0x11, 0x42,
	0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x66, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x27, 0x0a, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x42, 0x61, 0x6e, 0x6b,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65,
	0x0a, 0x10, 0x42, 0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x25,
	0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x04, 0x62, 0x61, 0x6e, 0x6b, 0x32, 0xf9, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x3a,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42,
	0x61, 0x6e, 0x6b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x32, 0x5a, 0x30, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bank_proto_goTypes = []interface{}{
	(*BankPayload)(nil),        // 0: bank.BankPayload
	(*BankCreateRequest)(nil),  // 1: bank.BankCreateRequest
	(*BankCreateReply)(nil),    // 2: bank.BankCreateReply
	(*BankUpdateRequest)(nil),  // 3: bank.BankUpdateRequest
	(*BankUpdateReply)(nil),    // 4: bank.BankUpdateReply
	(*BankGetAllRequest)(nil),  // 5: bank.BankGetAllRequest
	(*BankGetAllReply)(nil),    // 6: bank.BankGetAllReply
	(*BankGetByIDRequest)(nil), // 7: bank.BankGetByIDRequest
	(*BankGetByIDReply)(nil),   // 8: bank.BankGetByIDReply
}
var file_bank_proto_depIdxs = []int32{
	0, // 0: bank.BankCreateRequest.payload:type_name -> bank.BankPayload
	0, // 1: bank.BankUpdateRequest.payload:type_name -> bank.BankPayload
	0, // 2: bank.BankGetAllReply.banks:type_name -> bank.BankPayload
	0, // 3: bank.BankGetByIDReply.bank:type_name -> bank.BankPayload
	1, // 4: bank.Bank.Create:input_type -> bank.BankCreateRequest
	3, // 5: bank.Bank.Update:input_type -> bank.BankUpdateRequest
	5, // 6: bank.Bank.GetAll:input_type -> bank.BankGetAllRequest
	7, // 7: bank.Bank.GetByID:input_type -> bank.BankGetByIDRequest
	2, // 8: bank.Bank.Create:output_type -> bank.BankCreateReply
	4, // 9: bank.Bank.Update:output_type -> bank.BankUpdateReply
	6, // 10: bank.Bank.GetAll:output_type -> bank.BankGetAllReply
	8, // 11: bank.Bank.GetByID:output_type -> bank.BankGetByIDReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankPayload); i {
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
		file_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCreateRequest); i {
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
		file_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCreateReply); i {
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
		file_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankUpdateRequest); i {
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
		file_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankUpdateReply); i {
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
		file_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankGetAllRequest); i {
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
		file_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankGetAllReply); i {
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
		file_bank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankGetByIDRequest); i {
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
		file_bank_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankGetByIDReply); i {
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
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}
