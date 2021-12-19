package resolver

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"google.golang.org/grpc/resolver"
	"log"
)

const (
	schemeName = "consul"
)

func init() {
	resolver.Register(&consulBuilder{})
}

type consulResolver struct {
	ctx         context.Context
	cancel      context.CancelFunc
	cc          resolver.ClientConn
	config      api.Config
	serviceName string
}

func (r *consulResolver) ResolveNow(resolver.ResolveNowOptions) {
}

func (r *consulResolver) Close() {
	r.cancel()
}

func (r *consulResolver) watchServices() {
	errChan := make(chan error)
	p, err := watch.Parse(map[string]interface{}{
		"type":        "service",
		"service":     r.serviceName,
		"passingonly": true,
	})
	if err != nil {
		panic(err)
	}

	p.Handler = func(index uint64, result interface{}) {
		if entries, ok := result.([]*api.ServiceEntry); ok {
			conns := make([]resolver.Address, 0)
			for _, e := range entries {
				if e != nil {
					conns = append(conns, resolver.Address{
						Addr: fmt.Sprintf("%v:%v", e.Service.Address, e.Service.Port),
					})
				}
			}
			_ = r.cc.UpdateState(resolver.State{Addresses: conns})
		}
	}

	go func() {
		if err = p.Run(r.config.Address); err != nil {
			r.ctx.Done()
			errChan <- err
		}
	}()

	for {
		select {
		case <-errChan:
			log.Printf("run watch error: %v\n", err)
			return
		case <-r.ctx.Done():
			return
		}
	}
}
