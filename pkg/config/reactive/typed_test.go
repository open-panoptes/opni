package reactive_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/rancher/opni/pkg/config/reactive"
	"github.com/rancher/opni/pkg/test/testdata/plugins/ext"
	"github.com/rancher/opni/pkg/test/testutil"
)

var _ = Describe("Typed Reactive Messages", Label("unit"), func() {
	It("should create typed messages", func(ctx SpecContext) {
		rv := reactive.NewReactiveValue()

		actual := &ext.Sample2FieldMsg{
			Field1: 100,
			Field2: 200,
		}

		val := protoreflect.ValueOf(actual.ProtoReflect())
		rv.Update(1, val, make(chan struct{}), true)
		Expect(rv.Value()).To(testutil.ProtoValueEqual(val))

		typedRv := reactive.Message[*ext.Sample2FieldMsg](rv)
		Expect(typedRv.Value()).To(testutil.ProtoEqual(actual))

		typedW := typedRv.Watch(ctx)
		Eventually(typedW).Should(Receive(testutil.ProtoEqual(actual)))

		called := false
		check := func(m *ext.Sample2FieldMsg) {
			Expect(m).To(testutil.ProtoEqual(actual))
		}
		typedRv.WatchFunc(ctx, func(v *ext.Sample2FieldMsg) {
			check(v)
			called = true
		})
		Expect(called).To(BeTrue(), "watch func was not called")

		actual2 := &ext.Sample2FieldMsg{
			Field1: 300,
			Field2: 400,
		}

		called = false
		check = func(m *ext.Sample2FieldMsg) {
			Expect(m).To(testutil.ProtoEqual(actual2))
		}
		rv.Update(2, protoreflect.ValueOf(actual2.ProtoReflect()), make(chan struct{}), true)

		Eventually(typedW).Should(Receive(testutil.ProtoEqual(actual2)))
		Expect(called).To(BeTrue(), "watch func was not called")
	})

	It("should create typed scalars", func(ctx SpecContext) {
		rv := reactive.NewReactiveValue()

		actual := int32(100)

		rv.Update(1, protoreflect.ValueOf(actual), make(chan struct{}), true)
		Expect(rv.Value().Int()).To(Equal(int64(100))) // note the type conversion by Int()

		typedRv := reactive.Scalar[int32](rv)
		Expect(typedRv.Value()).To(Equal(actual)) // note the lack of type conversion

		called := false
		check := func(m int32) {
			Expect(m).To(Equal(actual))
		}
		typedRv.WatchFunc(ctx, func(v int32) {
			check(v)
			called = true
		})
		Expect(called).To(BeTrue(), "watch func was not called")

		typedW := typedRv.Watch(ctx)
		Eventually(typedW).Should(Receive(Equal(actual)))

		actual2 := int32(200)
		called = false
		check = func(m int32) {
			Expect(m).To(Equal(actual2))
		}
		rv.Update(2, protoreflect.ValueOf(actual2), make(chan struct{}), true)

		Eventually(typedW).Should(Receive(Equal(actual2)))
		Expect(called).To(BeTrue(), "watch func was not called")
	})
})
