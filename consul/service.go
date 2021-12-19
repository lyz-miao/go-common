package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"time"
)

type Registrar struct {
	client *api.Client
	id     string
	logger log.Logger
	config *api.AgentServiceRegistration
	stop   chan struct{}
	// todo Added max retry for re-register
	maxRetry int
}

func NewRegistrar(serverAddress string) (r *Registrar, err error) {
	defConfig := api.DefaultConfig()
	if serverAddress != "" {
		defConfig.Address = serverAddress
	}

	client, err := api.NewClient(defConfig)
	if err != nil {
		return nil, err
	}

	r = &Registrar{
		client: client,
		id:     strconv.FormatInt(time.Now().UnixNano(), 10),
		stop:   make(chan struct{}, 0),
	}

	return
}

func (r *Registrar) RegistrarWithConfig(req *api.AgentServiceRegistration) error {
	if req.ID == "" {
		r.id = req.ID
	}
	r.config = req
	err := r.client.Agent().ServiceRegister(r.config)
	if err != nil {
		return err
	}
	log.Printf("%v register successful\n", req.Name)
	go r.Watch()

	return nil
}

func (r *Registrar) RegistrarGRPC(serviceName, addr string, port int) error {
	r.config = &api.AgentServiceRegistration{
		ID:      r.id,
		Name:    serviceName,
		Address: addr,
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%v:%v", addr, port),
			Interval:                       "2s",
			Timeout:                        "1s",
			Notes:                          "Consul check service health status.",
			DeregisterCriticalServiceAfter: "2s",
		},
	}
	err := r.RegistrarWithConfig(r.config)
	if err != nil {
		return err
	}

	return nil
}

func (r *Registrar) RegistrarHTTP(serviceName, addr string, port int) error {
	r.config = &api.AgentServiceRegistration{
		ID:      r.id,
		Name:    serviceName,
		Address: addr,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%v:%v/health/check", addr, port),
			Interval:                       "2s",
			Timeout:                        "1s",
			Notes:                          "Consul check service health status.",
			DeregisterCriticalServiceAfter: "2s",
		},
	}
	err := r.RegistrarWithConfig(r.config)
	if err != nil {
		return err
	}

	return nil
}

func (r *Registrar) Watch() {
	for {
		select {
		case <-r.stop:
			close(r.stop)
			return
		case <-time.After(time.Second * 5):
			if r.config != nil {
				_, _, err := r.client.Agent().Service(r.config.ID, nil)
				if err != nil {
					log.Print("Discovery connect error or service does not exist. recovering...")

					err = r.client.Agent().ServiceRegister(r.config)
					if err != nil {
						log.Printf("Service re-register error: %v\n", err)
					} else {
						log.Print("Service re-register successful")
					}
				}
			}
		}
	}
}

func (r *Registrar) Deregister() error {
	err := r.client.Agent().ServiceDeregister(r.id)
	if err != nil {
		log.Printf("Service deregister error: %v\n", err)
		return err
	}
	defer log.Printf("Service deregister successful\n")

	r.stop <- struct{}{}
	r.config = nil

	return nil
}
