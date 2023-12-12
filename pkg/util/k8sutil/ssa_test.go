package k8sutil_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"

	"github.com/rancher/opni/pkg/util/k8sutil"
)

var _ = Describe("Server-Side Apply", func() {
	It("should decode managedFields", func() {
		entry := k8sutil.DecodeManagedFieldsEntry(metav1.ManagedFieldsEntry{
			APIVersion: "core.opni.io/v1",
			FieldsType: "FieldsV1",
			Manager:    "test-manager",
			Operation:  metav1.ManagedFieldsOperationApply,
			Time:       lo.ToPtr(metav1.Now()),
			FieldsV1: &metav1.FieldsV1{
				Raw: []byte(`{"f:spec":{"f:config":{"f:server":{".":{}, "f:grpcListenAddress": {}}}}}`),
			},
		})
		Expect(entry.Has(fieldpath.MakePathOrDie("spec", "config"))).To(BeFalse()) // intermediate messages without fields are not included
		Expect(entry.Has(fieldpath.MakePathOrDie("spec", "config", "server"))).To(BeTrue())
		Expect(entry.Has(fieldpath.MakePathOrDie("spec", "config", "server", "grpcListenAddress"))).To(BeTrue())
	})

	It("should return the field manager for a given path", func() {
		obj := corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				ManagedFields: []metav1.ManagedFieldsEntry{
					{
						APIVersion: "core.opni.io/v1",
						FieldsType: "FieldsV1",
						Manager:    "test-manager-1",
						Operation:  metav1.ManagedFieldsOperationApply,
						Time:       lo.ToPtr(metav1.Now()),
						FieldsV1: &metav1.FieldsV1{
							Raw: []byte(`{"f:spec":{"f:config":{"f:server":{".":{}, "f:grpcListenAddress": {}}}}}`),
						},
					},
					{
						APIVersion: "core.opni.io/v1",
						FieldsType: "FieldsV1",
						Manager:    "test-manager-2",
						Operation:  metav1.ManagedFieldsOperationApply,
						Time:       lo.ToPtr(metav1.Now()),
						FieldsV1: &metav1.FieldsV1{
							Raw: []byte(`{"f:spec":{"f:config":{"f:management":{".":{}, "f:grpcListenAddress": {}}}}}`),
						},
					},
				},
			},
		}
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec"))).To(Equal(""))
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec", "config"))).To(Equal(""))
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec", "config", "server"))).To(Equal("test-manager-1"))
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec", "config", "server", "grpcListenAddress"))).To(Equal("test-manager-1"))
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec", "config", "management"))).To(Equal("test-manager-2"))
		Expect(k8sutil.FieldManagerForPath(&obj, fieldpath.MakePathOrDie("spec", "config", "management", "grpcListenAddress"))).To(Equal("test-manager-2"))
	})
})
