// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/plugin/sdk/v1alpha1/proto/plugin.proto

package proto

import (
	buildeventstream "aspect.build/cli/bazel/buildeventstream"
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

type BEPEventCallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *buildeventstream.BuildEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *BEPEventCallbackReq) Reset() {
	*x = BEPEventCallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BEPEventCallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BEPEventCallbackReq) ProtoMessage() {}

func (x *BEPEventCallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BEPEventCallbackReq.ProtoReflect.Descriptor instead.
func (*BEPEventCallbackReq) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *BEPEventCallbackReq) GetEvent() *buildeventstream.BuildEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type BEPEventCallbackRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BEPEventCallbackRes) Reset() {
	*x = BEPEventCallbackRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BEPEventCallbackRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BEPEventCallbackRes) ProtoMessage() {}

func (x *BEPEventCallbackRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BEPEventCallbackRes.ProtoReflect.Descriptor instead.
func (*BEPEventCallbackRes) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{1}
}

type PostBuildHookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrokerId          uint32 `protobuf:"varint,1,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	IsInteractiveMode bool   `protobuf:"varint,2,opt,name=is_interactive_mode,json=isInteractiveMode,proto3" json:"is_interactive_mode,omitempty"`
}

func (x *PostBuildHookReq) Reset() {
	*x = PostBuildHookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBuildHookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBuildHookReq) ProtoMessage() {}

func (x *PostBuildHookReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBuildHookReq.ProtoReflect.Descriptor instead.
func (*PostBuildHookReq) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PostBuildHookReq) GetBrokerId() uint32 {
	if x != nil {
		return x.BrokerId
	}
	return 0
}

func (x *PostBuildHookReq) GetIsInteractiveMode() bool {
	if x != nil {
		return x.IsInteractiveMode
	}
	return false
}

type PostBuildHookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostBuildHookRes) Reset() {
	*x = PostBuildHookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBuildHookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBuildHookRes) ProtoMessage() {}

func (x *PostBuildHookRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBuildHookRes.ProtoReflect.Descriptor instead.
func (*PostBuildHookRes) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{3}
}

type PromptRunReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label       string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Default     string `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	AllowEdit   bool   `protobuf:"varint,3,opt,name=allow_edit,json=allowEdit,proto3" json:"allow_edit,omitempty"`
	Mask        string `protobuf:"bytes,5,opt,name=mask,proto3" json:"mask,omitempty"`
	HideEntered bool   `protobuf:"varint,6,opt,name=hide_entered,json=hideEntered,proto3" json:"hide_entered,omitempty"`
	IsConfirm   bool   `protobuf:"varint,8,opt,name=is_confirm,json=isConfirm,proto3" json:"is_confirm,omitempty"`
	IsVimMode   bool   `protobuf:"varint,9,opt,name=is_vim_mode,json=isVimMode,proto3" json:"is_vim_mode,omitempty"`
}

func (x *PromptRunReq) Reset() {
	*x = PromptRunReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRunReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRunReq) ProtoMessage() {}

func (x *PromptRunReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRunReq.ProtoReflect.Descriptor instead.
func (*PromptRunReq) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *PromptRunReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PromptRunReq) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *PromptRunReq) GetAllowEdit() bool {
	if x != nil {
		return x.AllowEdit
	}
	return false
}

func (x *PromptRunReq) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *PromptRunReq) GetHideEntered() bool {
	if x != nil {
		return x.HideEntered
	}
	return false
}

func (x *PromptRunReq) GetIsConfirm() bool {
	if x != nil {
		return x.IsConfirm
	}
	return false
}

func (x *PromptRunReq) GetIsVimMode() bool {
	if x != nil {
		return x.IsVimMode
	}
	return false
}

type PromptRunRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string              `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  *PromptRunRes_Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PromptRunRes) Reset() {
	*x = PromptRunRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRunRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRunRes) ProtoMessage() {}

func (x *PromptRunRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRunRes.ProtoReflect.Descriptor instead.
func (*PromptRunRes) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *PromptRunRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *PromptRunRes) GetError() *PromptRunRes_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type PromptRunRes_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Happened bool   `protobuf:"varint,1,opt,name=happened,proto3" json:"happened,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PromptRunRes_Error) Reset() {
	*x = PromptRunRes_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRunRes_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRunRes_Error) ProtoMessage() {}

func (x *PromptRunRes_Error) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRunRes_Error.ProtoReflect.Descriptor instead.
func (*PromptRunRes_Error) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PromptRunRes_Error) GetHappened() bool {
	if x != nil {
		return x.Happened
	}
	return false
}

func (x *PromptRunRes_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_plugin_sdk_v1alpha1_proto_plugin_proto protoreflect.FileDescriptor

var file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x13, 0x42, 0x45, 0x50, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x45, 0x50, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x22, 0xd3, 0x01,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73,
	0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x69, 0x64, 0x65, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x6d, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x3d, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x97, 0x01, 0x0a,
	0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x4a, 0x0a, 0x10, 0x42, 0x45, 0x50, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x45, 0x50, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x45, 0x50, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x32, 0x3b, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x65, 0x72, 0x12, 0x2f, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescOnce sync.Once
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescData = file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDesc
)

func file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescGZIP() []byte {
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescOnce.Do(func() {
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescData)
	})
	return file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDescData
}

var file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_goTypes = []interface{}{
	(*BEPEventCallbackReq)(nil),         // 0: proto.BEPEventCallbackReq
	(*BEPEventCallbackRes)(nil),         // 1: proto.BEPEventCallbackRes
	(*PostBuildHookReq)(nil),            // 2: proto.PostBuildHookReq
	(*PostBuildHookRes)(nil),            // 3: proto.PostBuildHookRes
	(*PromptRunReq)(nil),                // 4: proto.PromptRunReq
	(*PromptRunRes)(nil),                // 5: proto.PromptRunRes
	(*PromptRunRes_Error)(nil),          // 6: proto.PromptRunRes.Error
	(*buildeventstream.BuildEvent)(nil), // 7: build_event_stream.BuildEvent
}
var file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_depIdxs = []int32{
	7, // 0: proto.BEPEventCallbackReq.event:type_name -> build_event_stream.BuildEvent
	6, // 1: proto.PromptRunRes.error:type_name -> proto.PromptRunRes.Error
	0, // 2: proto.Plugin.BEPEventCallback:input_type -> proto.BEPEventCallbackReq
	2, // 3: proto.Plugin.PostBuildHook:input_type -> proto.PostBuildHookReq
	4, // 4: proto.Prompter.Run:input_type -> proto.PromptRunReq
	1, // 5: proto.Plugin.BEPEventCallback:output_type -> proto.BEPEventCallbackRes
	3, // 6: proto.Plugin.PostBuildHook:output_type -> proto.PostBuildHookRes
	5, // 7: proto.Prompter.Run:output_type -> proto.PromptRunRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_init() }
func file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_init() {
	if File_pkg_plugin_sdk_v1alpha1_proto_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BEPEventCallbackReq); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BEPEventCallbackRes); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBuildHookReq); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBuildHookRes); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRunReq); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRunRes); i {
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
		file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRunRes_Error); i {
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
			RawDescriptor: file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_goTypes,
		DependencyIndexes: file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_depIdxs,
		MessageInfos:      file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_msgTypes,
	}.Build()
	File_pkg_plugin_sdk_v1alpha1_proto_plugin_proto = out.File
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_rawDesc = nil
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_goTypes = nil
	file_pkg_plugin_sdk_v1alpha1_proto_plugin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	BEPEventCallback(ctx context.Context, in *BEPEventCallbackReq, opts ...grpc.CallOption) (*BEPEventCallbackRes, error)
	PostBuildHook(ctx context.Context, in *PostBuildHookReq, opts ...grpc.CallOption) (*PostBuildHookRes, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) BEPEventCallback(ctx context.Context, in *BEPEventCallbackReq, opts ...grpc.CallOption) (*BEPEventCallbackRes, error) {
	out := new(BEPEventCallbackRes)
	err := c.cc.Invoke(ctx, "/proto.Plugin/BEPEventCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) PostBuildHook(ctx context.Context, in *PostBuildHookReq, opts ...grpc.CallOption) (*PostBuildHookRes, error) {
	out := new(PostBuildHookRes)
	err := c.cc.Invoke(ctx, "/proto.Plugin/PostBuildHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	BEPEventCallback(context.Context, *BEPEventCallbackReq) (*BEPEventCallbackRes, error)
	PostBuildHook(context.Context, *PostBuildHookReq) (*PostBuildHookRes, error)
}

// UnimplementedPluginServer can be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (*UnimplementedPluginServer) BEPEventCallback(context.Context, *BEPEventCallbackReq) (*BEPEventCallbackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BEPEventCallback not implemented")
}
func (*UnimplementedPluginServer) PostBuildHook(context.Context, *PostBuildHookReq) (*PostBuildHookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBuildHook not implemented")
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_BEPEventCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BEPEventCallbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).BEPEventCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/BEPEventCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).BEPEventCallback(ctx, req.(*BEPEventCallbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_PostBuildHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostBuildHookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).PostBuildHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/PostBuildHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).PostBuildHook(ctx, req.(*PostBuildHookReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BEPEventCallback",
			Handler:    _Plugin_BEPEventCallback_Handler,
		},
		{
			MethodName: "PostBuildHook",
			Handler:    _Plugin_PostBuildHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugin/sdk/v1alpha1/proto/plugin.proto",
}

// PrompterClient is the client API for Prompter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrompterClient interface {
	Run(ctx context.Context, in *PromptRunReq, opts ...grpc.CallOption) (*PromptRunRes, error)
}

type prompterClient struct {
	cc grpc.ClientConnInterface
}

func NewPrompterClient(cc grpc.ClientConnInterface) PrompterClient {
	return &prompterClient{cc}
}

func (c *prompterClient) Run(ctx context.Context, in *PromptRunReq, opts ...grpc.CallOption) (*PromptRunRes, error) {
	out := new(PromptRunRes)
	err := c.cc.Invoke(ctx, "/proto.Prompter/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrompterServer is the server API for Prompter service.
type PrompterServer interface {
	Run(context.Context, *PromptRunReq) (*PromptRunRes, error)
}

// UnimplementedPrompterServer can be embedded to have forward compatible implementations.
type UnimplementedPrompterServer struct {
}

func (*UnimplementedPrompterServer) Run(context.Context, *PromptRunReq) (*PromptRunRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterPrompterServer(s *grpc.Server, srv PrompterServer) {
	s.RegisterService(&_Prompter_serviceDesc, srv)
}

func _Prompter_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRunReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrompterServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Prompter/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrompterServer).Run(ctx, req.(*PromptRunReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prompter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Prompter",
	HandlerType: (*PrompterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Prompter_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugin/sdk/v1alpha1/proto/plugin.proto",
}