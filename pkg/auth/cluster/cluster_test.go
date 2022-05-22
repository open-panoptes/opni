package cluster_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"io"
	"math"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	"github.com/rancher/opni/pkg/auth/cluster"
	"github.com/rancher/opni/pkg/core"
	"github.com/rancher/opni/pkg/keyring"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/test"
	"github.com/rancher/opni/pkg/test/testutil"
	"github.com/rancher/opni/pkg/util"
	"github.com/valyala/fasthttp/fasthttputil"
)

func bodyStr(body io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}

var _ = Describe("Cluster Auth", Ordered, test.EnableIfCI[FlakeAttempts](5), Label(test.Slow, test.TimeSensitive), func() {
	var app *fiber.App
	var ctrl *gomock.Controller
	var client *http.Client
	var addr string
	newRequest := func(method string, target string, body io.Reader) *http.Request {
		return util.Must(http.NewRequest(method, "http://"+addr+target, body))
	}
	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
	})
	Context("invalid auth headers", func() {
		BeforeAll(func() {
			app = fiber.New(fiber.Config{
				DisableStartupMessage: true,
			})
			broker := test.NewTestKeyringStoreBroker(ctrl)
			cm, err := cluster.New(context.Background(), broker, "X-Test")
			Expect(err).NotTo(HaveOccurred())
			app.Use(cm.Handle)
			app.Post("/", func(c *fiber.Ctx) error {
				return c.SendStatus(http.StatusOK)
			})
			listener := fasthttputil.NewInmemoryListener()
			go app.Listener(listener)
			addr = listener.Addr().String()
			client = &http.Client{
				Transport: &http.Transport{
					Dial: func(network, addr string) (net.Conn, error) {
						return listener.Dial()
					},
				},
			}
			for {
				_, err := client.Get("http://" + addr)
				if err == nil {
					break
				}
				time.Sleep(10 * time.Millisecond)
			}
			DeferCleanup(app.Shutdown)
		})
		When("no auth header is provided", func() {
			It("should return http 400", func() {
				resp, err := client.Do(newRequest(http.MethodPost, "/", nil))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
		When("the auth header has the wrong type", func() {
			It("should return http 400", func() {
				req := newRequest(http.MethodPost, "/", nil)
				req.Header.Set("Authorization", "foo")
				resp, err := client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
		When("the auth mac is malformed", func() {
			It("should return http 400", func() {
				req := newRequest(http.MethodPost, "/", nil)
				req.Header.Set("Authorization", "MAC foo")
				resp, err := client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})
	When("a well-formed auth header is provided", func() {
		var handler test.KeyringStoreHandler
		Context("invalid requests", func() {
			JustBeforeEach(func() {
				app = fiber.New(fiber.Config{
					DisableStartupMessage: true,
				})
				broker := test.NewTestKeyringStoreBroker(ctrl, handler)
				cm, err := cluster.New(context.Background(), broker, "X-Test")
				Expect(err).NotTo(HaveOccurred())
				app.Use(cm.Handle)
				app.Post("/", func(c *fiber.Ctx) error {
					return c.SendStatus(http.StatusOK)
				})
				listener := fasthttputil.NewInmemoryListener()
				go app.Listener(listener)
				addr = listener.Addr().String()
				client = &http.Client{
					Transport: &http.Transport{
						Dial: func(network, addr string) (net.Conn, error) {
							return listener.Dial()
						},
					},
				}
				for {
					_, err := client.Get("http://" + addr)
					if err == nil {
						break
					}
					time.Sleep(10 * time.Millisecond)
				}
				DeferCleanup(func() {
					app.Shutdown()
					handler = nil
				})
			})

			When("the keyring store does not exist for the requested cluster", func() {
				BeforeEach(func() {
					handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
						return nil, errors.New("not found")
					}
				})
				It("should return http 401", func() {
					req := newRequest(http.MethodPost, "/", nil)
					req.Header.Set("Authorization", validAuthHeader("cluster-1", ""))
					resp, err := client.Do(req)
					Expect(err).NotTo(HaveOccurred())
					defer resp.Body.Close()
					Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(bodyStr(resp.Body)).To(Equal(http.StatusText(http.StatusUnauthorized)))
				})
			})

			When("the keyring does not exist in the cluster's keyring store", func() {
				BeforeEach(func() {
					store := test.NewTestKeyringStore(ctrl, "", &core.Reference{
						Id: "does-not-exist",
					})
					handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
						return store, nil
					}
				})
				It("should return http 401", func() {
					req := newRequest(http.MethodPost, "/", nil)
					req.Header.Set("Authorization", validAuthHeader("cluster-1", ""))
					resp, err := client.Do(req)
					Expect(err).NotTo(HaveOccurred())
					defer resp.Body.Close()
					Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(bodyStr(resp.Body)).To(Equal(http.StatusText(http.StatusUnauthorized)))
				})
			})

			When("the keyring exists in the cluster's keyring store", func() {
				When("the keyring is missing the required key", func() {
					BeforeEach(func() {
						store := test.NewTestKeyringStore(ctrl, "", &core.Reference{
							Id: "cluster-1",
						})
						store.Put(context.Background(), keyring.New())
						handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
							return store, nil
						}
					})
					It("should return http 401", func() {
						req := newRequest(http.MethodPost, "/", nil)
						req.Header.Set("Authorization", validAuthHeader("cluster-1", ""))
						resp, err := client.Do(req)
						Expect(err).NotTo(HaveOccurred())
						defer resp.Body.Close()
						Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
						Expect(bodyStr(resp.Body)).To(Equal("invalid or corrupted keyring"))
					})
				})
				When("the request MAC does not match the request body", func() {
					BeforeEach(func() {
						store := test.NewTestKeyringStore(ctrl, "", &core.Reference{
							Id: "cluster-1",
						})
						store.Put(context.Background(), keyring.New(keyring.NewSharedKeys(testSharedSecret)))
						handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
							return store, nil
						}
					})
					It("should return http 401", func() {
						req := newRequest(http.MethodPost, "/", strings.NewReader("not-matching"))
						req.Header.Set("Authorization", validAuthHeader("cluster-1", "Not_Matching"))
						resp, err := client.Do(req)
						Expect(err).NotTo(HaveOccurred())
						defer resp.Body.Close()
						Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
						Expect(bodyStr(resp.Body)).To(Equal(http.StatusText(http.StatusUnauthorized)))
					})
				})
			})

			Context("request timing", func() {
				BeforeAll(func() {
					// temporarily pause garbage collection to avoid interfering with timing
					gcPercent := debug.SetGCPercent(-1)
					DeferCleanup(func() {
						debug.SetGCPercent(gcPercent)
					})
				})
				BeforeEach(func() {
					store := test.NewTestKeyringStore(ctrl, "", &core.Reference{
						Id: "cluster-1",
					})
					store.Put(context.Background(), keyring.New(keyring.NewSharedKeys(testSharedSecret)))
					handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
						if ref.Id == "cluster-1" {
							return store, nil
						}
						return nil, errors.New("not found")
					}
				})
				Specify("different unauthorized requests should take the same amount of time", func() {
					exp := gmeasure.NewExperiment("request-timing")

					largeBody := make([]byte, 2*1024*1024)
					rand.Read(largeBody)
					largeBody2 := make([]byte, 2*1024*1024)
					rand.Read(largeBody2)
					invalidExists := invalidAuthHeader("cluster-1", largeBody2)
					validDoesNotExist := validAuthHeader("cluster-2", largeBody)
					invalidDoesNotExist := invalidAuthHeader("cluster-2", largeBody2)

					wg := &sync.WaitGroup{}
					wg.Add(3)

					go func() {
						defer wg.Done()
						exp.SampleDuration("valid mac, cluster does not exist", func(int) {
							req := newRequest(http.MethodPost, "/", bytes.NewReader(largeBody))
							req.Header.Set("Authorization", validDoesNotExist)
							resp, err := client.Do(req)
							Expect(err).NotTo(HaveOccurred())
							defer resp.Body.Close()
							Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
						}, gmeasure.SamplingConfig{
							N: testutil.IfCI(200).Else(1000),
						}, gmeasure.Precision(time.Microsecond))
					}()

					go func() {
						defer wg.Done()
						exp.SampleDuration("valid mac, cluster exists", func(int) {
							req := newRequest(http.MethodPost, "/", bytes.NewReader(largeBody))
							req.Header.Set("Authorization", invalidDoesNotExist)
							resp, err := client.Do(req)
							Expect(err).NotTo(HaveOccurred())
							defer resp.Body.Close()
							Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
						}, gmeasure.SamplingConfig{
							N: testutil.IfCI(200).Else(1000),
						}, gmeasure.Precision(time.Microsecond))
					}()

					go func() {
						defer wg.Done()
						exp.SampleDuration("invalid mac, cluster exists", func(int) {
							req := newRequest(http.MethodPost, "/", bytes.NewReader(largeBody))
							req.Header.Set("Authorization", invalidExists)
							resp, err := client.Do(req)
							Expect(err).NotTo(HaveOccurred())
							defer resp.Body.Close()
							Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
						}, gmeasure.SamplingConfig{
							N: testutil.IfCI(200).Else(1000),
						}, gmeasure.Precision(time.Microsecond))
					}()

					wg.Wait()

					AddReportEntry(exp.Name, exp)

					s1 := exp.Get("valid mac, cluster does not exist").Stats()
					s2 := exp.Get("valid mac, cluster exists").Stats()
					s3 := exp.Get("invalid mac, cluster exists").Stats()

					m1 := s1.DurationFor(gmeasure.StatMean).Round(s1.PrecisionBundle.Duration)
					m2 := s2.DurationFor(gmeasure.StatMean).Round(s2.PrecisionBundle.Duration)
					m3 := s3.DurationFor(gmeasure.StatMean).Round(s3.PrecisionBundle.Duration)

					// all three requests should take the same amount of time within 1 standard deviation

					// use the smallest standard deviation as a threshold
					d1 := s1.DurationFor(gmeasure.StatStdDev).Microseconds()
					d2 := s2.DurationFor(gmeasure.StatStdDev).Microseconds()
					d3 := s3.DurationFor(gmeasure.StatStdDev).Microseconds()
					stdDev := int64(math.Min(math.Min(float64(d1), float64(d2)), float64(d3)))
					Expect(m1.Microseconds()).To(BeNumerically("~", m2.Microseconds(), stdDev))
					Expect(m2.Microseconds()).To(BeNumerically("~", m3.Microseconds(), stdDev))
					Expect(m3.Microseconds()).To(BeNumerically("~", m1.Microseconds(), stdDev))
				})
			})
		})
		Context("valid requests", func() {
			JustBeforeEach(func() {
				app = fiber.New(fiber.Config{
					DisableStartupMessage: true,
				})
				broker := test.NewTestKeyringStoreBroker(ctrl, handler)
				cm, err := cluster.New(context.Background(), broker, "X-Test")
				Expect(err).NotTo(HaveOccurred())
				app.Use(cm.Handle)
				app.Post("/", func(c *fiber.Ctx) error {
					defer GinkgoRecover()
					Expect(cluster.AuthorizedID(c)).To(Equal("cluster-1"))
					Expect(cluster.AuthorizedKeys(c)).NotTo(BeNil())
					return c.SendStatus(http.StatusOK)
				})
				listener := fasthttputil.NewInmemoryListener()
				go app.Listener(listener)
				addr = listener.Addr().String()
				client = &http.Client{
					Transport: &http.Transport{
						Dial: func(network, addr string) (net.Conn, error) {
							return listener.Dial()
						},
					},
				}
				for {
					_, err := client.Get("http://" + addr)
					if err == nil {
						break
					}
					time.Sleep(10 * time.Millisecond)
				}
				DeferCleanup(app.Shutdown)
			})

			When("the request MAC matches the request body", func() {
				BeforeEach(func() {
					store := test.NewTestKeyringStore(ctrl, "", &core.Reference{
						Id: "cluster-1",
					})
					store.Put(context.Background(), keyring.New(keyring.NewSharedKeys(testSharedSecret)))
					handler = func(_ context.Context, prefix string, ref *core.Reference) (storage.KeyringStore, error) {
						if ref.Id == "cluster-1" {
							return store, nil
						}
						return nil, errors.New("not found")
					}
				})
				It("should return http 200", func() {
					req := newRequest(http.MethodPost, "/", strings.NewReader("payload"))
					req.Header.Set("Authorization", validAuthHeader("cluster-1", "payload"))
					resp, err := client.Do(req)
					Expect(err).NotTo(HaveOccurred())
					defer resp.Body.Close()
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
				})
			})
		})
	})
})
