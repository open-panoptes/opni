package crds_test

import (
	"context"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	opnicorev1 "github.com/rancher/opni/apis/core/v1"
	opnicorev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	monitoringv1beta1 "github.com/rancher/opni/apis/monitoring/v1beta1"
	"github.com/rancher/opni/pkg/config/reactive"
	"github.com/rancher/opni/pkg/config/reactive/subtle"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/storage/crds"
	"github.com/rancher/opni/pkg/storage/inmemory"
	. "github.com/rancher/opni/pkg/test/conformance/storage"
	_ "github.com/rancher/opni/pkg/test/setup"
	"github.com/rancher/opni/pkg/test/testk8s"
	"github.com/rancher/opni/pkg/test/testruntime"
	"github.com/rancher/opni/pkg/test/testutil"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/flagutil"
	"github.com/rancher/opni/pkg/util/future"
	"github.com/rancher/opni/pkg/util/k8sutil"
	"github.com/rancher/opni/pkg/util/protorand"
	"github.com/rancher/opni/plugins/metrics/apis/cortexops"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestCrds(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(1 * time.Minute)
	RunSpecs(t, "CRDs Storage Suite")
}

var store = future.New[*crds.CRDStore]()
var metricsKvStore = future.New[*metricsBroker]()
var gatewayKvStore = future.New[*gatewayBroker]()

type methods struct{}

// ControllerReference implements crds.ValueStoreMethods.
func (methods) ControllerReference() (client.Object, bool) {
	return nil, false
}

type metricsBroker struct {
	k8sClient  client.WithWatch
	namespaces map[string]struct{}
}

type gatewayBroker struct {
	k8sClient  client.WithWatch
	namespaces map[string]struct{}
}

func (b *metricsBroker) KeyValueStore(namespace string) storage.KeyValueStoreT[*cortexops.CapabilityBackendConfigSpec] {
	return inmemory.NewCustomKeyValueStore(func(key string) storage.ValueStoreT[*cortexops.CapabilityBackendConfigSpec] {
		if b.namespaces == nil {
			b.namespaces = make(map[string]struct{})
		}
		if _, ok := b.namespaces[namespace]; !ok {
			if err := b.k8sClient.Create(context.Background(), &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}); err != nil {
				panic(err)
			}
			b.namespaces[namespace] = struct{}{}
		}

		key = sanitizeKey(key)
		vs := crds.NewCRDValueStore(client.ObjectKey{
			Namespace: namespace,
			Name:      key,
		}, methods{}, crds.WithClient(b.k8sClient))
		return vs
	})
}

func (b *gatewayBroker) KeyValueStore(namespace string) storage.KeyValueStoreT[*configv1.GatewayConfigSpec] {
	return inmemory.NewCustomKeyValueStore(func(key string) storage.ValueStoreT[*configv1.GatewayConfigSpec] {
		if b.namespaces == nil {
			b.namespaces = make(map[string]struct{})
		}
		if _, ok := b.namespaces[namespace]; !ok {
			if err := b.k8sClient.Create(context.Background(), &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}); err != nil {
				panic(err)
			}
			b.namespaces[namespace] = struct{}{}
		}

		key = sanitizeKey(key)
		vs := crds.NewCRDValueStore(client.ObjectKey{
			Namespace: namespace,
			Name:      key,
		}, opnicorev1.ValueStoreMethods{}, crds.WithClient(b.k8sClient))
		return vs
	})
}

func sanitizeKey(key string) string {
	return strings.ReplaceAll(key, "/", "--")
}

// FillConfigFromObject implements crds.ValueStoreMethods.
func (methods) FillConfigFromObject(obj *opnicorev1beta1.MonitoringCluster, conf *cortexops.CapabilityBackendConfigSpec) {
	conf.Enabled = obj.Spec.Cortex.Enabled
	conf.CortexConfig = obj.Spec.Cortex.CortexConfig
	conf.CortexWorkloads = obj.Spec.Cortex.CortexWorkloads
	conf.Grafana = obj.Spec.Grafana.GrafanaConfig
}

// FillObjectFromConfig implements crds.ValueStoreMethods.
func (methods) FillObjectFromConfig(obj *opnicorev1beta1.MonitoringCluster, conf *cortexops.CapabilityBackendConfigSpec) {
	if conf == nil {
		obj.Spec.Cortex = opnicorev1beta1.CortexSpec{}
		obj.Spec.Grafana = opnicorev1beta1.GrafanaSpec{}
		return
	}
	obj.Spec.Cortex.Enabled = conf.Enabled
	obj.Spec.Cortex.CortexConfig = conf.CortexConfig
	obj.Spec.Cortex.CortexWorkloads = conf.CortexWorkloads
	obj.Spec.Grafana.GrafanaConfig = conf.Grafana

}

func newMetricsObject(seed ...int64) *cortexops.CapabilityBackendConfigSpec {
	if len(seed) == 0 {
		return nil
	}
	rand := protorand.New[*cortexops.CapabilityBackendConfigSpec]()
	rand.MaxCollectionElements = 1
	rand.Seed(seed[0])
	out := rand.MustGen()
	out.Revision = nil
	return out
}

func newGatewayObject(seed ...int64) *configv1.GatewayConfigSpec {
	if len(seed) == 0 {
		return nil
	}
	rand := protorand.New[*configv1.GatewayConfigSpec]()
	rand.MaxCollectionElements = 1
	rand.Seed(seed[0])
	out := rand.MustGen()
	out.Revision = nil
	return out
}

var _ crds.ValueStoreMethods[*opnicorev1beta1.MonitoringCluster, *cortexops.CapabilityBackendConfigSpec] = methods{}

var k8sClient client.WithWatch

var _ = BeforeSuite(func() {
	// sanity-check the valuestoremethods impl
	conf := protorand.New[*cortexops.CapabilityBackendConfigSpec]().MustGen()
	conf.Revision = nil

	var obj opnicorev1beta1.MonitoringCluster
	methods{}.FillObjectFromConfig(&obj, conf.DeepCopy())

	conf2 := util.NewMessage[*cortexops.CapabilityBackendConfigSpec]()
	methods{}.FillConfigFromObject(obj.DeepCopy(), conf2)

	Expect(conf2).To(testutil.ProtoEqual(conf))

	testruntime.IfLabelFilterMatches(Label("integration", "slow"), func() {
		ctx, ca := context.WithCancel(context.Background())
		s := scheme.Scheme
		opnicorev1.AddToScheme(s)
		opnicorev1beta1.AddToScheme(s)
		monitoringv1beta1.AddToScheme(s)
		config, _, err := testk8s.StartK8s(ctx, []string{"../../../config/crd/bases"}, s)
		Expect(err).NotTo(HaveOccurred())

		store.Set(crds.NewCRDStore(crds.WithRestConfig(config)))

		k8sClient, err = k8sutil.NewK8sClient(k8sutil.ClientOptions{
			RestConfig: config,
			Scheme:     s,
		})
		Expect(err).NotTo(HaveOccurred())

		metricsKvStore.Set(&metricsBroker{k8sClient: k8sClient})
		gatewayKvStore.Set(&gatewayBroker{k8sClient: k8sClient})

		DeferCleanup(ca)
	})
})

var _ = Describe("CRD Token Store", Ordered, Label("integration", "slow"), TokenStoreTestSuite(store))
var _ = Describe("CRD RBAC Store", Ordered, Label("integration", "slow"), RBACStoreTestSuite(store))
var _ = Describe("CRD Keyring Store", Ordered, Label("integration", "slow"), KeyringStoreTestSuite(store))
var _ = Describe("CRD Value Store (Metrics Config)", Ordered, Label("integration", "slow"), KeyValueStoreTestSuite(metricsKvStore, newMetricsObject, func(a any) types.GomegaMatcher {
	return testutil.ProtoEqual(a.(proto.Message))
}))
var _ = Describe("CRD Value Store (Gateway Config)", Ordered, Label("integration", "slow"), KeyValueStoreTestSuite(gatewayKvStore, newGatewayObject, func(a any) types.GomegaMatcher {
	return testutil.ProtoEqual(a.(proto.Message))
}))
var _ = Describe("Gateway Config Manager", Ordered, Label("integration", "slow"), func() {
	It("should set up a gateway config manager using a CRD-backed active store", func(ctx SpecContext) {
		ns := "crds-conformance-test"
		if err := k8sClient.Create(context.Background(), &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: ns,
			},
		}); err != nil {
			panic(err)
		}
		defaultStore := inmemory.NewValueStore[*configv1.GatewayConfigSpec](util.ProtoClone)
		activeStore := crds.NewCRDValueStore[*opnicorev1.Gateway, *configv1.GatewayConfigSpec](k8stypes.NamespacedName{
			Namespace: ns,
			Name:      "test",
		}, opnicorev1.ValueStoreMethods{}, crds.WithClient(k8sClient))
		mgr := configv1.NewGatewayConfigManager(
			defaultStore, activeStore,
			flagutil.LoadDefaults,
			configv1.WithControllerOptions(
				reactive.WithDiffMode(reactive.DiffFull),
			),
		)

		conf := &configv1.GatewayConfigSpec{}
		flagutil.LoadDefaults(conf)

		// ensure fields which do not have flags in GatewayConfigSpec are initialized
		conf.Storage = &configv1.StorageSpec{
			Backend: configv1.StorageBackend_Etcd.Enum(),
		}

		conf.Server.AdvertiseAddress = lo.ToPtr("${POD_IP}:9090")
		conf.Management.AdvertiseAddress = lo.ToPtr("${POD_IP}:11090")
		conf.Relay.AdvertiseAddress = lo.ToPtr("${POD_IP}:11190")
		conf.Dashboard.AdvertiseAddress = lo.ToPtr("${POD_IP}:12080")
		conf.Keyring.RuntimeKeyDirs = []string{"/run/opni/keyring"}
		conf.Plugins.Dir = lo.ToPtr("/var/lib/opni/plugins")
		conf.Certs = &configv1.CertsSpec{}

		conf.Storage.Etcd = &configv1.EtcdSpec{
			Endpoints: []string{"etcd:2379"},
			Certs: &configv1.MTLSSpec{
				ServerCA:   lo.ToPtr("/run/etcd/certs/server/ca.crt"),
				ClientCA:   lo.ToPtr("/run/etcd/certs/client/ca.crt"),
				ClientCert: lo.ToPtr("/run/etcd/certs/client/tls.crt"),
				ClientKey:  lo.ToPtr("/run/etcd/certs/client/tls.key"),
			},
		}

		gw := &opnicorev1.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: ns,
			},
			Spec: opnicorev1.GatewaySpec{
				Config: conf,
			},
		}
		Expect(k8sClient.Create(ctx, gw)).To(Succeed())

		time.Sleep(1 * time.Second)
		go mgr.Start(ctx)

		done := make(chan struct{})
		mgr.Reactive(protopath.Path(configv1.ProtoPath().Plugins())).WatchFunc(ctx, func(val protoreflect.Value) {
			defer GinkgoRecover()
			defer close(done)
			msg := val.Message().Interface().(*configv1.PluginsSpec)
			Expect(msg.GetDir()).NotTo(BeEmpty())
		})
		Eventually(done).Should(BeClosed())

		w := mgr.Reactive(protopath.Path(configv1.ProtoPath().Plugins())).Watch(ctx)
		msg := (<-w).Message().Interface().(*configv1.PluginsSpec)
		Expect(msg.GetDir()).NotTo(BeEmpty())

		Expect(subtle.WaitOneMessage[*configv1.PluginsSpec](ctx, mgr.Reactive(protopath.Path(configv1.ProtoPath().Plugins())))).To(testutil.ProtoEqual(conf.Plugins))
	})
})
