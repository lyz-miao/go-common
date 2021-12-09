package consul

import (
    "errors"
    "github.com/hashicorp/consul/api"
    "sync/atomic"
)

type Balancer interface {
    Get(in []*api.ServiceEntry) (*api.ServiceEntry, error)
}

type roundRobin struct {
    c uint64
}

func NewRoundRobin() Balancer {
    return &roundRobin{
        c: 0,
    }
}

func (rr *roundRobin) Get(in []*api.ServiceEntry) (*api.ServiceEntry, error) {
    index := rr.c
    if int(index) != len(in){
        index = uint64(len(in))
    }
    if index <= 0 {
        return nil, errors.New("no entry available")
    }
    old := atomic.AddUint64(&rr.c, 1) - 1
    idx := old % index
    return in[idx], nil
}
