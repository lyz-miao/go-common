package consul

import (
	"github.com/hashicorp/consul/api"
)

func NewKV(addr *string) (*api.KV, error) {
	config := api.DefaultConfig()
	if addr != nil {
		config.Address = *addr
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client.KV(), nil
}
