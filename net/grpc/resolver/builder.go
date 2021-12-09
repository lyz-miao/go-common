package resolver

import (
    "context"
    "github.com/hashicorp/consul/api"
    "google.golang.org/grpc/resolver"
    "os"
)

type consulBuilder struct {
}

func (b *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
    config := api.DefaultConfig()
    host := os.Getenv("DISCOVER_SERVER")
    if host != ""{
        config.Address = host
    }
    ctx, cancel := context.WithCancel(context.Background())

    r := &consulResolver{
        ctx:         ctx,
        cancel:      cancel,
        cc:          cc,
        config:      *config,
        serviceName: target.Endpoint,
    }

    go r.watchServices()
    r.ResolveNow(resolver.ResolveNowOptions{})
    return r, nil
}

func (b *consulBuilder) Scheme() string {
    return schemeName
}
