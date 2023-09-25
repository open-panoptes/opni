// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto

package apiextensions

import (
	totem "github.com/kralicky/totem"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CertConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ca       string `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	CaData   []byte `protobuf:"bytes,2,opt,name=caData,proto3" json:"caData,omitempty"`
	Cert     string `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	CertData []byte `protobuf:"bytes,4,opt,name=certData,proto3" json:"certData,omitempty"`
	Key      string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	KeyData  []byte `protobuf:"bytes,6,opt,name=keyData,proto3" json:"keyData,omitempty"`
	Insecure bool   `protobuf:"varint,7,opt,name=insecure,proto3" json:"insecure,omitempty"`
}

func (x *CertConfig) Reset() {
	*x = CertConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertConfig) ProtoMessage() {}

func (x *CertConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertConfig.ProtoReflect.Descriptor instead.
func (*CertConfig) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{0}
}

func (x *CertConfig) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

func (x *CertConfig) GetCaData() []byte {
	if x != nil {
		return x.CaData
	}
	return nil
}

func (x *CertConfig) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *CertConfig) GetCertData() []byte {
	if x != nil {
		return x.CertData
	}
	return nil
}

func (x *CertConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CertConfig) GetKeyData() []byte {
	if x != nil {
		return x.KeyData
	}
	return nil
}

func (x *CertConfig) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

type HTTPAPIExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAddr string       `protobuf:"bytes,1,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	Routes   []*RouteInfo `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *HTTPAPIExtensionConfig) Reset() {
	*x = HTTPAPIExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPAPIExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPAPIExtensionConfig) ProtoMessage() {}

func (x *HTTPAPIExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPAPIExtensionConfig.ProtoReflect.Descriptor instead.
func (*HTTPAPIExtensionConfig) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPAPIExtensionConfig) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

func (x *HTTPAPIExtensionConfig) GetRoutes() []*RouteInfo {
	if x != nil {
		return x.Routes
	}
	return nil
}

type ServiceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceDescriptor *descriptorpb.ServiceDescriptorProto `protobuf:"bytes,1,opt,name=serviceDescriptor,proto3" json:"serviceDescriptor,omitempty"`
	Options           *ServiceOptions                      `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ServiceDescriptor) Reset() {
	*x = ServiceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDescriptor) ProtoMessage() {}

func (x *ServiceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDescriptor.ProtoReflect.Descriptor instead.
func (*ServiceDescriptor) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceDescriptor) GetServiceDescriptor() *descriptorpb.ServiceDescriptorProto {
	if x != nil {
		return x.ServiceDescriptor
	}
	return nil
}

func (x *ServiceDescriptor) GetOptions() *ServiceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the service will only be available to clusters that have this
	// capability.
	RequireCapability string `protobuf:"bytes,1,opt,name=requireCapability,proto3" json:"requireCapability,omitempty"`
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceOptions) GetRequireCapability() string {
	if x != nil {
		return x.RequireCapability
	}
	return ""
}

type ServiceDescriptorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ServiceDescriptor `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ServiceDescriptorList) Reset() {
	*x = ServiceDescriptorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDescriptorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDescriptorList) ProtoMessage() {}

func (x *ServiceDescriptorList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDescriptorList.ProtoReflect.Descriptor instead.
func (*ServiceDescriptorList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceDescriptorList) GetItems() []*ServiceDescriptor {
	if x != nil {
		return x.Items
	}
	return nil
}

type ServiceDescriptorProtoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*descriptorpb.ServiceDescriptorProto `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ServiceDescriptorProtoList) Reset() {
	*x = ServiceDescriptorProtoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDescriptorProtoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDescriptorProtoList) ProtoMessage() {}

func (x *ServiceDescriptorProtoList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDescriptorProtoList.ProtoReflect.Descriptor instead.
func (*ServiceDescriptorProtoList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceDescriptorProtoList) GetItems() []*descriptorpb.ServiceDescriptorProto {
	if x != nil {
		return x.Items
	}
	return nil
}

type RouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *RouteInfo) Reset() {
	*x = RouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteInfo) ProtoMessage() {}

func (x *RouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteInfo.ProtoReflect.Descriptor instead.
func (*RouteInfo) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{6}
}

func (x *RouteInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RouteInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79,
	0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x43, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x63, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x63, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x22, 0x66, 0x0a, 0x16, 0x48, 0x54, 0x54, 0x50, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74,
	0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74,
	0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x55,
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3e,
	0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x4f,
	0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x5b, 0x0a, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x09,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0x6a, 0x0a, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x50, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0x61, 0x0a, 0x10, 0x48, 0x54, 0x54, 0x50, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x32, 0x43, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x50,
	0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x0f, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0a, 0x2e,
	0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x50, 0x43, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x74, 0x65,
	0x6d, 0x2e, 0x52, 0x50, 0x43, 0x28, 0x01, 0x30, 0x01, 0x32, 0x67, 0x0a, 0x11, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x52,
	0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = []interface{}{
	(*CertConfig)(nil),                          // 0: apiextensions.CertConfig
	(*HTTPAPIExtensionConfig)(nil),              // 1: apiextensions.HTTPAPIExtensionConfig
	(*ServiceDescriptor)(nil),                   // 2: apiextensions.ServiceDescriptor
	(*ServiceOptions)(nil),                      // 3: apiextensions.ServiceOptions
	(*ServiceDescriptorList)(nil),               // 4: apiextensions.ServiceDescriptorList
	(*ServiceDescriptorProtoList)(nil),          // 5: apiextensions.ServiceDescriptorProtoList
	(*RouteInfo)(nil),                           // 6: apiextensions.RouteInfo
	(*descriptorpb.ServiceDescriptorProto)(nil), // 7: google.protobuf.ServiceDescriptorProto
	(*emptypb.Empty)(nil),                       // 8: google.protobuf.Empty
	(*totem.RPC)(nil),                           // 9: totem.RPC
}
var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = []int32{
	6, // 0: apiextensions.HTTPAPIExtensionConfig.routes:type_name -> apiextensions.RouteInfo
	7, // 1: apiextensions.ServiceDescriptor.serviceDescriptor:type_name -> google.protobuf.ServiceDescriptorProto
	3, // 2: apiextensions.ServiceDescriptor.options:type_name -> apiextensions.ServiceOptions
	2, // 3: apiextensions.ServiceDescriptorList.items:type_name -> apiextensions.ServiceDescriptor
	7, // 4: apiextensions.ServiceDescriptorProtoList.items:type_name -> google.protobuf.ServiceDescriptorProto
	8, // 5: apiextensions.ManagementAPIExtension.Descriptors:input_type -> google.protobuf.Empty
	0, // 6: apiextensions.HTTPAPIExtension.Configure:input_type -> apiextensions.CertConfig
	9, // 7: apiextensions.StreamAPIExtension.ConnectInternal:input_type -> totem.RPC
	8, // 8: apiextensions.UnaryAPIExtension.UnaryDescriptor:input_type -> google.protobuf.Empty
	5, // 9: apiextensions.ManagementAPIExtension.Descriptors:output_type -> apiextensions.ServiceDescriptorProtoList
	1, // 10: apiextensions.HTTPAPIExtension.Configure:output_type -> apiextensions.HTTPAPIExtensionConfig
	9, // 11: apiextensions.StreamAPIExtension.ConnectInternal:output_type -> totem.RPC
	7, // 12: apiextensions.UnaryAPIExtension.UnaryDescriptor:output_type -> google.protobuf.ServiceDescriptorProto
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_init() }
func file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_init() {
	if File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertConfig); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPAPIExtensionConfig); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDescriptor); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDescriptorList); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDescriptorProtoList); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteInfo); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto = out.File
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = nil
}
