// Code generated by internal/codegen/pathbuilder/generator.go. DO NOT EDIT.
// source: github.com/rancher/opni/pkg/test/testdata/plugins/ext/ext.proto

package ext

import (
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	protopath "google.golang.org/protobuf/reflect/protopath"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type (
	sampleConfigurationPathBuilder protopath.Path
	revisionPathBuilder            protopath.Path
	timestampPathBuilder           protopath.Path
	sampleMessagePathBuilder       protopath.Path
	sample1FieldMsgPathBuilder     protopath.Path
	sample2FieldMsgPathBuilder     protopath.Path
	sample3FieldMsgPathBuilder     protopath.Path
	sample4FieldMsgPathBuilder     protopath.Path
	sample5FieldMsgPathBuilder     protopath.Path
	sample6FieldMsgPathBuilder     protopath.Path
	sampleMessage2PathBuilder      protopath.Path
)

func (*SampleConfiguration) ProtoPath() sampleConfigurationPathBuilder {
	return sampleConfigurationPathBuilder{protopath.Root(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor())}
}

func (p sampleConfigurationPathBuilder) Revision() revisionPathBuilder {
	return revisionPathBuilder(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sampleConfigurationPathBuilder) MessageField() sampleMessagePathBuilder {
	return sampleMessagePathBuilder(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(8))))
}
func (p revisionPathBuilder) Timestamp() timestampPathBuilder {
	return timestampPathBuilder(append(p, protopath.FieldAccess(((*v1.Revision)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sampleMessagePathBuilder) Field1() sample1FieldMsgPathBuilder {
	return sample1FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sampleMessagePathBuilder) Field2() sample2FieldMsgPathBuilder {
	return sample2FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sampleMessagePathBuilder) Field3() sample3FieldMsgPathBuilder {
	return sample3FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sampleMessagePathBuilder) Field4() sample4FieldMsgPathBuilder {
	return sample4FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sampleMessagePathBuilder) Field5() sample5FieldMsgPathBuilder {
	return sample5FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p sampleMessagePathBuilder) Field6() sample6FieldMsgPathBuilder {
	return sample6FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
func (p sampleMessagePathBuilder) Msg() sampleMessage2PathBuilder {
	return sampleMessage2PathBuilder(append(p, protopath.FieldAccess(((*SampleMessage)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(7))))
}
func (p sampleMessage2PathBuilder) Field1() sample1FieldMsgPathBuilder {
	return sample1FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sampleMessage2PathBuilder) Field2() sample2FieldMsgPathBuilder {
	return sample2FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sampleMessage2PathBuilder) Field3() sample3FieldMsgPathBuilder {
	return sample3FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sampleMessage2PathBuilder) Field4() sample4FieldMsgPathBuilder {
	return sample4FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sampleMessage2PathBuilder) Field5() sample5FieldMsgPathBuilder {
	return sample5FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p sampleMessage2PathBuilder) Field6() sample6FieldMsgPathBuilder {
	return sample6FieldMsgPathBuilder(append(p, protopath.FieldAccess(((*SampleMessage2)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}

func (p sampleConfigurationPathBuilder) Enabled() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sampleConfigurationPathBuilder) StringField() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sampleConfigurationPathBuilder) SecretField() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sampleConfigurationPathBuilder) MapField() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p sampleConfigurationPathBuilder) RepeatedField() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
func (p sampleConfigurationPathBuilder) EnumField() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*SampleConfiguration)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(7))))
}
func (p revisionPathBuilder) Revision() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*v1.Revision)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p timestampPathBuilder) Seconds() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*timestamppb.Timestamp)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p timestampPathBuilder) Nanos() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*timestamppb.Timestamp)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample1FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample1FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample2FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample2FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample2FieldMsgPathBuilder) Field2() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample2FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample3FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample3FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample3FieldMsgPathBuilder) Field2() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample3FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample3FieldMsgPathBuilder) Field3() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample3FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sample4FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample4FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample4FieldMsgPathBuilder) Field2() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample4FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample4FieldMsgPathBuilder) Field3() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample4FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sample4FieldMsgPathBuilder) Field4() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample4FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sample5FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample5FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample5FieldMsgPathBuilder) Field2() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample5FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample5FieldMsgPathBuilder) Field3() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample5FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sample5FieldMsgPathBuilder) Field4() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample5FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sample5FieldMsgPathBuilder) Field5() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample5FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p sample6FieldMsgPathBuilder) Field1() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p sample6FieldMsgPathBuilder) Field2() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p sample6FieldMsgPathBuilder) Field3() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p sample6FieldMsgPathBuilder) Field4() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p sample6FieldMsgPathBuilder) Field5() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p sample6FieldMsgPathBuilder) Field6() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*Sample6FieldMsg)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}