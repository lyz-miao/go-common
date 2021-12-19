package grpc

import (
	"context"
	"fmt"
	_ "github.com/lyz-miao/go-common/net/grpc/resolver"
	"google.golang.org/grpc"
)

func NewConn(ctx context.Context, name string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))

	return grpc.DialContext(ctx, fmt.Sprintf("consul://%v", name), opts...)
}
