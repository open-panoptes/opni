package reactive

import (
	"fmt"

	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/plugins/driverutil"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type WatchReactiveServerStream interface {
	Send(*corev1.ReactiveEvents) error
	grpc.ServerStream
}

type ControllerServer[S WatchReactiveServerStream, T driverutil.ConfigType[T]] struct {
	*Controller[T]
}

func (c *ControllerServer[S, T]) Build(ctrl *Controller[T]) {
	c.Controller = ctrl
}

func (s *ControllerServer[S, T]) WatchReactive(in *corev1.ReactiveWatchRequest, stream S) error {
	rvs := make([]Value, 0, len(in.Paths))
	var err error
	s.usePathTrie(func(p *pathTrie[*reactiveValue]) {
		for _, path := range in.Paths {
			if path == "." {
				rvs = append(rvs, p.root.value)
				continue
			}
			if v := p.FindString(path); v != nil {
				rvs = append(rvs, v.value)
			} else {
				err = fmt.Errorf("no such path %s", path)
				return
			}
		}
	})
	if err != nil {
		return err
	}
	if in.Bind {
		Bind(stream.Context(), func(v []protoreflect.Value) {
			items := make([]*corev1.ReactiveEvent, 0, len(v))
			for i, value := range v {
				items = append(items, &corev1.ReactiveEvent{
					Index: int32(i),
					Value: corev1.NewValue(value),
				})
			}
			stream.Send(&corev1.ReactiveEvents{Items: items})
		}, rvs...)
	} else {
		for i, rv := range rvs {
			i := i
			rv.WatchFunc(stream.Context(), func(value protoreflect.Value) {
				stream.Send(&corev1.ReactiveEvents{
					Items: []*corev1.ReactiveEvent{
						{
							Index: int32(i),
							Value: corev1.NewValue(value),
						},
					},
				})
			})
		}
	}

	return nil
}
