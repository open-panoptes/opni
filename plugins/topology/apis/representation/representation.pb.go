// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/plugins/topology/apis/representation/representation.proto

package representation

import (
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
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

type GraphRepr int32

const (
	GraphRepr_None         GraphRepr = 0
	GraphRepr_KubectlGraph GraphRepr = 1
)

// Enum value maps for GraphRepr.
var (
	GraphRepr_name = map[int32]string{
		0: "None",
		1: "KubectlGraph",
	}
	GraphRepr_value = map[string]int32{
		"None":         0,
		"KubectlGraph": 1,
	}
)

func (x GraphRepr) Enum() *GraphRepr {
	p := new(GraphRepr)
	*p = x
	return p
}

func (x GraphRepr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphRepr) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_enumTypes[0].Descriptor()
}

func (GraphRepr) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_enumTypes[0]
}

func (x GraphRepr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphRepr.Descriptor instead.
func (GraphRepr) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescGZIP(), []int{0}
}

type TopologyGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *v1.Reference `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Repr GraphRepr     `protobuf:"varint,3,opt,name=repr,proto3,enum=representation.GraphRepr" json:"repr,omitempty"`
}

func (x *TopologyGraph) Reset() {
	*x = TopologyGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyGraph) ProtoMessage() {}

func (x *TopologyGraph) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyGraph.ProtoReflect.Descriptor instead.
func (*TopologyGraph) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescGZIP(), []int{0}
}

func (x *TopologyGraph) GetId() *v1.Reference {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TopologyGraph) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TopologyGraph) GetRepr() GraphRepr {
	if x != nil {
		return x.Repr
	}
	return GraphRepr_None
}

type DOTRepresentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawDotFormat string `protobuf:"bytes,1,opt,name=rawDotFormat,proto3" json:"rawDotFormat,omitempty"`
}

func (x *DOTRepresentation) Reset() {
	*x = DOTRepresentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DOTRepresentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DOTRepresentation) ProtoMessage() {}

func (x *DOTRepresentation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DOTRepresentation.ProtoReflect.Descriptor instead.
func (*DOTRepresentation) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescGZIP(), []int{1}
}

func (x *DOTRepresentation) GetRawDotFormat() string {
	if x != nil {
		return x.RawDotFormat
	}
	return ""
}

var File_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDesc = []byte{
	0x0a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x52, 0x65, 0x70, 0x72, 0x52, 0x04, 0x72, 0x65, 0x70, 0x72, 0x22, 0x37, 0x0a, 0x11, 0x44,
	0x4f, 0x54, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x44, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61, 0x77, 0x44, 0x6f, 0x74, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2a, 0x27, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x70,
	0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4b,
	0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x10, 0x01, 0x32, 0xb4, 0x01,
	0x0a, 0x16, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x5e, 0x0a, 0x0b, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x4f, 0x54, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x22, 0x10, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescData = file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_goTypes = []interface{}{
	(GraphRepr)(0),            // 0: representation.GraphRepr
	(*TopologyGraph)(nil),     // 1: representation.TopologyGraph
	(*DOTRepresentation)(nil), // 2: representation.DOTRepresentation
	(*v1.Reference)(nil),      // 3: core.Reference
}
var file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_depIdxs = []int32{
	3, // 0: representation.TopologyGraph.id:type_name -> core.Reference
	0, // 1: representation.TopologyGraph.repr:type_name -> representation.GraphRepr
	3, // 2: representation.TopologyRepresentation.GetGraph:input_type -> core.Reference
	3, // 3: representation.TopologyRepresentation.RenderGraph:input_type -> core.Reference
	1, // 4: representation.TopologyRepresentation.GetGraph:output_type -> representation.TopologyGraph
	2, // 5: representation.TopologyRepresentation.RenderGraph:output_type -> representation.DOTRepresentation
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_init()
}
func file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_init() {
	if File_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologyGraph); i {
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
		file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DOTRepresentation); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto = out.File
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_topology_apis_representation_representation_proto_depIdxs = nil
}
