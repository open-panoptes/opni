package integration_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/gmeasure"
	"github.com/prometheus/procfs"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/clients"
	"github.com/rancher/opni/pkg/config/adapt"
	"github.com/rancher/opni/pkg/config/meta"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/test"
	"github.com/rancher/opni/pkg/test/freeport"
	"github.com/rancher/opni/pkg/test/testlog"
	"github.com/rancher/opni/pkg/tokens"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"sigs.k8s.io/yaml"
)

func buildPrerequisites() error {
	testlog.Log.Debug("building prerequisite binaries...")
	cmd := exec.Command("mage", "build:plugin", "example", "build:opniminimal", "build:opni")
	cmd.Dir = "../.."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var _ = Describe("Agent Memory Tests", Ordered, Serial, Label("integration", "slow"), func() {
	var environment *test.Environment
	var client managementv1.ManagementClient
	var fingerprint string
	var gatewayConfig *v1beta1.GatewayConfig
	var agentSession *gexec.Session
	var gatewaySession *gexec.Session
	var startGateway func()
	agentListenPort := freeport.GetFreePort()

	waitForGatewayReady := func(timeout time.Duration) {
		// ping /healthz until it returns 200
		testlog.Log.Debug("waiting for gateway to be ready")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		for ctx.Err() == nil {
			resp, err := http.Get(fmt.Sprintf("http://%s/healthz", gatewayConfig.Spec.MetricsListenAddress))
			if err == nil && resp.StatusCode == 200 {
				return
			} else if err == nil {
				testlog.Log.Debug(fmt.Sprintf("gateway not ready yet: %s", resp.Status))
			} else {
				testlog.Log.Debug(fmt.Sprintf("gateway not ready yet: %s", err.Error()))
			}
			time.Sleep(50 * time.Millisecond)
		}
		Fail("timed out waiting for gateway to be ready")
	}

	waitForAgentReady := func(timeout time.Duration) {
		// ping /healthz until it returns 200
		testlog.Log.Debug("waiting for agent to be ready")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		for ctx.Err() == nil {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/healthz", agentListenPort))
			if err == nil && resp.StatusCode == 200 {
				return
			} else if err == nil {
				status, _ := io.ReadAll(resp.Body)
				testlog.Log.Debug(fmt.Sprintf("agent not ready yet: %s", string(status)))
			}
			time.Sleep(200 * time.Millisecond)
		}
		Fail("timed out waiting for agent to be ready")
	}

	waitForAgentUnready := func(timeout time.Duration) {
		testlog.Log.Debug("waiting for agent to become unready")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		for ctx.Err() == nil {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/healthz", agentListenPort))
			if err != nil || resp.StatusCode != 200 {
				return
			}
			testlog.Log.Debug("agent still ready...")
			time.Sleep(200 * time.Millisecond)
		}
		Fail("timed out waiting for agent to become unready")
	}

	BeforeAll(func() {
		if testing.Short() {
			Skip("skipping agent memory tests in short mode")
		}

		Expect(buildPrerequisites()).To(Succeed())

		environment = &test.Environment{}
		Expect(environment.Start(test.WithEnableGateway(false), test.WithStorageBackend(v1beta1.StorageTypeEtcd))).To(Succeed())

		DeferCleanup(environment.Stop)
		tempDir, err := os.MkdirTemp("", "opni-test")
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(func() {
			os.RemoveAll(tempDir)
		})
		os.Mkdir(path.Join(tempDir, "plugins"), 0755)
		os.Mkdir(path.Join(tempDir, "cache"), 0755)
		// copy plugin_example only to the plugins dir
		file, err := os.Open("../../bin/plugins/plugin_example")
		Expect(err).NotTo(HaveOccurred())

		dest, err := os.OpenFile(filepath.Join(tempDir, "plugins/plugin_example"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		Expect(err).NotTo(HaveOccurred())

		_, err = io.Copy(dest, file)
		Expect(err).NotTo(HaveOccurred())
		dest.Close()
		file.Close()

		gatewayConfig = environment.NewGatewayConfig()

		cfgv1 := adapt.V1ConfigOf[*configv1.GatewayConfigSpec](&gatewayConfig.Spec)

		startGateway = func() {
			cmd := exec.Command("bin/opni", "gateway",
				"--storage.etcd.endpoints", strings.Join(cfgv1.Storage.Etcd.GetEndpoints(), ","),
				"--defaults.certs.ca-cert-data", cfgv1.Certs.GetCaCertData(),
				"--defaults.certs.serving-cert-data", cfgv1.Certs.GetServingCertData(),
				"--defaults.certs.serving-key-data", cfgv1.Certs.GetServingKeyData(),
				"--defaults.management.grpc-listen-address", cfgv1.Management.GetGrpcListenAddress(),
				"--defaults.management.http-listen-address", cfgv1.Management.GetHttpListenAddress(),
				"--defaults.server.grpc-listen-address", cfgv1.Server.GetGrpcListenAddress(),
				"--defaults.server.http-listen-address", cfgv1.Server.GetHttpListenAddress(),
				"--defaults.health.http-listen-address", cfgv1.Health.GetHttpListenAddress(),
				"--defaults.dashboard.http-listen-address", cfgv1.Dashboard.GetHttpListenAddress(),
				"--defaults.relay.grpc-listen-address", cfgv1.Relay.GetGrpcListenAddress(),
				"--defaults.plugins.dir", path.Join(tempDir, "plugins"),
				"--defaults.plugins.cache.filesystem.dir", path.Join(tempDir, "cache"),
			)

			cmd.Dir = "../../"
			cmd.SysProcAttr = &syscall.SysProcAttr{
				Setsid: true,
			}
			gatewaySession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			waitForGatewayReady(10 * time.Second)
			DeferCleanup(func() {
				gatewaySession.Interrupt().Wait()
			})
		}

		By("starting an external gateway process", startGateway)

		By("starting an external agent process", func() {
			client, err = clients.NewManagementClient(environment.Context(), clients.WithAddress(
				strings.TrimPrefix(gatewayConfig.Spec.Management.GRPCListenAddress, "tcp://"),
			), clients.WithDialOptions(grpc.WithBlock(), grpc.FailOnNonTempDialError(false)))
			Expect(err).NotTo(HaveOccurred())

			certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
			Expect(err).NotTo(HaveOccurred())
			fingerprint = certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
			Expect(fingerprint).NotTo(BeEmpty())

			token, err := client.CreateBootstrapToken(context.Background(), &managementv1.CreateBootstrapTokenRequest{
				Ttl: durationpb.New(time.Minute),
			})
			Expect(err).NotTo(HaveOccurred())
			t, err := tokens.FromBootstrapToken(token)
			Expect(err).NotTo(HaveOccurred())

			tempDir, err := os.MkdirTemp("", "opni-test")
			Expect(err).NotTo(HaveOccurred())

			DeferCleanup(func() {
				os.RemoveAll(tempDir)
			})
			agentConfig := &v1beta1.AgentConfig{
				TypeMeta: meta.TypeMeta{
					APIVersion: "v1beta1",
					Kind:       "AgentConfig",
				},
				Spec: v1beta1.AgentConfigSpec{
					TrustStrategy:    v1beta1.TrustStrategyPKP,
					ListenAddress:    fmt.Sprintf("localhost:%d", agentListenPort),
					GatewayAddress:   gatewayConfig.Spec.GRPCListenAddress,
					IdentityProvider: "env",
					Storage: v1beta1.StorageSpec{
						Type: v1beta1.StorageTypeEtcd,
						Etcd: &v1beta1.EtcdStorageSpec{
							Endpoints: environment.EtcdConfig().Endpoints,
						},
					},
					Bootstrap: &v1beta1.BootstrapSpec{
						Token: t.EncodeHex(),
						Pins:  []string{fingerprint},
					},
					PluginDir: path.Join(tempDir, "plugins"),
					PluginUpgrade: v1beta1.PluginUpgradeSpec{
						Type:   v1beta1.PluginUpgradeBinary,
						Binary: &v1beta1.BinaryUpgradeSpec{},
					},
				},
			}
			configFile := path.Join(tempDir, "config.yaml")
			configData, err := yaml.Marshal(agentConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(os.WriteFile(configFile, configData, 0644)).To(Succeed())

			cmd := exec.Command("bin/opni-minimal", "agentv2", "--config", configFile)
			cmd.Dir = "../.."
			cmd.SysProcAttr = &syscall.SysProcAttr{
				Setsid: true,
			}
			cmd.Env = append(os.Environ(), "OPNI_UNIQUE_IDENTIFIER=agent1")
			agentSession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			DeferCleanup(func() {
				agentSession.Interrupt().Wait()
			})
		})
	})

	Specify("watching agent memory usage", func() {
		var rssValues []int
		exp := gmeasure.NewExperiment("agent rss")
		for i := 0; i < 10; i++ {
			waitForAgentReady(10 * time.Second)
			time.Sleep(1 * time.Second) // wait for the agent to settle

			pid := agentSession.Command.Process.Pid
			proc, err := procfs.NewProc(pid)
			Expect(err).NotTo(HaveOccurred())

			status, err := proc.Stat()
			Expect(err).NotTo(HaveOccurred())

			rssKB := status.RSS * os.Getpagesize() / 1024
			testlog.Log.Info("agent rss", " rssKB ", rssKB)
			exp.RecordValue("rss", float64(rssKB), gmeasure.Units("KB"))
			rssValues = append(rssValues, rssKB)

			gatewaySession.Interrupt().Wait()
			waitForAgentUnready(10 * time.Second)
			By("restarting the gateway", startGateway)
		}

		AddReportEntry(exp.Name, exp)

		// check that the memory usage is not monotonically increasing
		var changeOverTime []int
		testlog.Log.Debug(fmt.Sprintf("rss 0: %d", rssValues[0]))
		for i := 1; i < len(rssValues); i++ {
			diff := rssValues[i] - rssValues[i-1]
			changeOverTime = append(changeOverTime, diff)
			if diff >= 0 {
				testlog.Log.Debug(fmt.Sprintf("rss %d: %d (+%d)", i, rssValues[i], diff))
			} else {
				testlog.Log.Debug(fmt.Sprintf("rss %d: %d (%d)", i, rssValues[i], diff))
			}
		}
		Expect(changeOverTime).To(ContainElement(BeNumerically("<=", 0)), "memory usage should not be monotonically increasing")
	})
})
