// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/config/v1/config_server.proto

package configv1

import (
	validate "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/rancher/opni/internal/codegen/cli"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	driverutil "github.com/rancher/opni/pkg/plugins/driverutil"
	_ "github.com/rancher/opni/pkg/validation"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *GatewayConfigSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetSpec() *GatewayConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision *v1.Revision           `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Mask     *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
	Patch    *GatewayConfigSpec     `protobuf:"bytes,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP(), []int{1}
}

func (x *ResetRequest) GetRevision() *v1.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

func (x *ResetRequest) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

func (x *ResetRequest) GetPatch() *GatewayConfigSpec {
	if x != nil {
		return x.Patch
	}
	return nil
}

type DryRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   driverutil.Target      `protobuf:"varint,1,opt,name=target,proto3,enum=driverutil.Target" json:"target,omitempty"`
	Action   driverutil.Action      `protobuf:"varint,2,opt,name=action,proto3,enum=driverutil.Action" json:"action,omitempty"`
	Spec     *GatewayConfigSpec     `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`         // Set
	Revision *v1.Revision           `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"` // Reset
	Patch    *GatewayConfigSpec     `protobuf:"bytes,5,opt,name=patch,proto3" json:"patch,omitempty"`       // Reset
	Mask     *fieldmaskpb.FieldMask `protobuf:"bytes,6,opt,name=mask,proto3" json:"mask,omitempty"`         // Reset
}

func (x *DryRunRequest) Reset() {
	*x = DryRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunRequest) ProtoMessage() {}

func (x *DryRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunRequest.ProtoReflect.Descriptor instead.
func (*DryRunRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP(), []int{2}
}

func (x *DryRunRequest) GetTarget() driverutil.Target {
	if x != nil {
		return x.Target
	}
	return driverutil.Target(0)
}

func (x *DryRunRequest) GetAction() driverutil.Action {
	if x != nil {
		return x.Action
	}
	return driverutil.Action(0)
}

func (x *DryRunRequest) GetSpec() *GatewayConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *DryRunRequest) GetRevision() *v1.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

func (x *DryRunRequest) GetPatch() *GatewayConfigSpec {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *DryRunRequest) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

type DryRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current          *GatewayConfigSpec   `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Modified         *GatewayConfigSpec   `protobuf:"bytes,2,opt,name=modified,proto3" json:"modified,omitempty"`
	ValidationErrors *validate.Violations `protobuf:"bytes,3,opt,name=validationErrors,proto3" json:"validationErrors,omitempty"`
}

func (x *DryRunResponse) Reset() {
	*x = DryRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DryRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DryRunResponse) ProtoMessage() {}

func (x *DryRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DryRunResponse.ProtoReflect.Descriptor instead.
func (*DryRunResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP(), []int{3}
}

func (x *DryRunResponse) GetCurrent() *GatewayConfigSpec {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *DryRunResponse) GetModified() *GatewayConfigSpec {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *DryRunResponse) GetValidationErrors() *validate.Violations {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

type HistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GatewayConfigSpec `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryResponse) GetEntries() []*GatewayConfigSpec {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_github_com_rancher_opni_pkg_config_v1_config_server_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x75, 0x74, 0x69, 0x6c, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3e, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x92, 0xc0, 0x0c, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x06, 0x8a, 0xc0, 0x0c, 0x02, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x3a,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0x8a, 0xc0, 0x0c,
	0x02, 0x28, 0x01, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0xb1, 0x02, 0x0a, 0x0d, 0x44,
	0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x06, 0x8a, 0xc0, 0x0c, 0x02, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0xc8,
	0x01, 0x0a, 0x0e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x32, 0xdc, 0x06, 0x0a, 0x0d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x64, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x22, 0x13, 0x82, 0xc0, 0x0c, 0x0f, 0x8a, 0xc0, 0x0c,
	0x0b, 0x67, 0x65, 0x74, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x61, 0x0a, 0x17,
	0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xc0, 0x0c, 0x13, 0x8a, 0xc0, 0x0c, 0x0b,
	0x73, 0x65, 0x74, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0xb0, 0xc0, 0x0c, 0x01, 0x12,
	0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x75, 0x74, 0x69, 0x6c,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x22, 0x0b, 0x82, 0xc0, 0x0c, 0x07, 0x8a,
	0xc0, 0x0c, 0x03, 0x67, 0x65, 0x74, 0x12, 0x52, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0f, 0x82, 0xc0, 0x0c, 0x0b, 0x8a,
	0xc0, 0x0c, 0x03, 0x73, 0x65, 0x74, 0xb0, 0xc0, 0x0c, 0x01, 0x12, 0x62, 0x0a, 0x19, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xc0, 0x0c, 0x11, 0x8a, 0xc0, 0x0c,
	0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x58,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x11, 0x82, 0xc0, 0x0c, 0x0d, 0x8a, 0xc0, 0x0c, 0x05, 0x72,
	0x65, 0x73, 0x65, 0x74, 0xb0, 0xc0, 0x0c, 0x01, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x72, 0x79, 0x52,
	0x75, 0x6e, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x72, 0x79, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0x82, 0xc0, 0x0c, 0x04, 0xa0, 0xc0, 0x0c,
	0x01, 0x12, 0x6c, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f,
	0x82, 0xc0, 0x0c, 0x0b, 0x8a, 0xc0, 0x0c, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x52, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x0d, 0x82, 0xc0, 0x0c, 0x09, 0x8a, 0xc0, 0x0c, 0x05, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x30, 0x01, 0x1a, 0x0e, 0x82, 0xc0, 0x0c, 0x0a, 0x8a, 0xc0, 0x0c, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x36, 0x82, 0xc0, 0x0c, 0x02, 0x08, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescData = file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_rancher_opni_pkg_config_v1_config_server_proto_goTypes = []interface{}{
	(*SetRequest)(nil),                             // 0: config.v1.SetRequest
	(*ResetRequest)(nil),                           // 1: config.v1.ResetRequest
	(*DryRunRequest)(nil),                          // 2: config.v1.DryRunRequest
	(*DryRunResponse)(nil),                         // 3: config.v1.DryRunResponse
	(*HistoryResponse)(nil),                        // 4: config.v1.HistoryResponse
	(*GatewayConfigSpec)(nil),                      // 5: config.v1.GatewayConfigSpec
	(*v1.Revision)(nil),                            // 6: core.Revision
	(*fieldmaskpb.FieldMask)(nil),                  // 7: google.protobuf.FieldMask
	(driverutil.Target)(0),                         // 8: driverutil.Target
	(driverutil.Action)(0),                         // 9: driverutil.Action
	(*validate.Violations)(nil),                    // 10: buf.validate.Violations
	(*driverutil.GetRequest)(nil),                  // 11: driverutil.GetRequest
	(*emptypb.Empty)(nil),                          // 12: google.protobuf.Empty
	(*driverutil.ConfigurationHistoryRequest)(nil), // 13: driverutil.ConfigurationHistoryRequest
	(*v1.ReactiveWatchRequest)(nil),                // 14: core.ReactiveWatchRequest
	(*v1.ReactiveEvents)(nil),                      // 15: core.ReactiveEvents
}
var file_github_com_rancher_opni_pkg_config_v1_config_server_proto_depIdxs = []int32{
	5,  // 0: config.v1.SetRequest.spec:type_name -> config.v1.GatewayConfigSpec
	6,  // 1: config.v1.ResetRequest.revision:type_name -> core.Revision
	7,  // 2: config.v1.ResetRequest.mask:type_name -> google.protobuf.FieldMask
	5,  // 3: config.v1.ResetRequest.patch:type_name -> config.v1.GatewayConfigSpec
	8,  // 4: config.v1.DryRunRequest.target:type_name -> driverutil.Target
	9,  // 5: config.v1.DryRunRequest.action:type_name -> driverutil.Action
	5,  // 6: config.v1.DryRunRequest.spec:type_name -> config.v1.GatewayConfigSpec
	6,  // 7: config.v1.DryRunRequest.revision:type_name -> core.Revision
	5,  // 8: config.v1.DryRunRequest.patch:type_name -> config.v1.GatewayConfigSpec
	7,  // 9: config.v1.DryRunRequest.mask:type_name -> google.protobuf.FieldMask
	5,  // 10: config.v1.DryRunResponse.current:type_name -> config.v1.GatewayConfigSpec
	5,  // 11: config.v1.DryRunResponse.modified:type_name -> config.v1.GatewayConfigSpec
	10, // 12: config.v1.DryRunResponse.validationErrors:type_name -> buf.validate.Violations
	5,  // 13: config.v1.HistoryResponse.entries:type_name -> config.v1.GatewayConfigSpec
	11, // 14: config.v1.GatewayConfig.GetDefaultConfiguration:input_type -> driverutil.GetRequest
	0,  // 15: config.v1.GatewayConfig.SetDefaultConfiguration:input_type -> config.v1.SetRequest
	11, // 16: config.v1.GatewayConfig.GetConfiguration:input_type -> driverutil.GetRequest
	0,  // 17: config.v1.GatewayConfig.SetConfiguration:input_type -> config.v1.SetRequest
	12, // 18: config.v1.GatewayConfig.ResetDefaultConfiguration:input_type -> google.protobuf.Empty
	1,  // 19: config.v1.GatewayConfig.ResetConfiguration:input_type -> config.v1.ResetRequest
	2,  // 20: config.v1.GatewayConfig.DryRun:input_type -> config.v1.DryRunRequest
	13, // 21: config.v1.GatewayConfig.ConfigurationHistory:input_type -> driverutil.ConfigurationHistoryRequest
	14, // 22: config.v1.GatewayConfig.WatchReactive:input_type -> core.ReactiveWatchRequest
	5,  // 23: config.v1.GatewayConfig.GetDefaultConfiguration:output_type -> config.v1.GatewayConfigSpec
	12, // 24: config.v1.GatewayConfig.SetDefaultConfiguration:output_type -> google.protobuf.Empty
	5,  // 25: config.v1.GatewayConfig.GetConfiguration:output_type -> config.v1.GatewayConfigSpec
	12, // 26: config.v1.GatewayConfig.SetConfiguration:output_type -> google.protobuf.Empty
	12, // 27: config.v1.GatewayConfig.ResetDefaultConfiguration:output_type -> google.protobuf.Empty
	12, // 28: config.v1.GatewayConfig.ResetConfiguration:output_type -> google.protobuf.Empty
	3,  // 29: config.v1.GatewayConfig.DryRun:output_type -> config.v1.DryRunResponse
	4,  // 30: config.v1.GatewayConfig.ConfigurationHistory:output_type -> config.v1.HistoryResponse
	15, // 31: config.v1.GatewayConfig.WatchReactive:output_type -> core.ReactiveEvents
	23, // [23:32] is the sub-list for method output_type
	14, // [14:23] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_config_v1_config_server_proto_init() }
func file_github_com_rancher_opni_pkg_config_v1_config_server_proto_init() {
	if File_github_com_rancher_opni_pkg_config_v1_config_server_proto != nil {
		return
	}
	file_github_com_rancher_opni_pkg_config_v1_gateway_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunRequest); i {
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
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DryRunResponse); i {
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
		file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryResponse); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_config_v1_config_server_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_config_v1_config_server_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_config_v1_config_server_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_config_v1_config_server_proto = out.File
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_config_v1_config_server_proto_depIdxs = nil
}
