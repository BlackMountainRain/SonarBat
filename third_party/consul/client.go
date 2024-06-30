package consul

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/hashicorp/consul/api"
)

type Datacenter string

const (
	SingleDatacenter Datacenter = "SINGLE"
	MultiDatacenter  Datacenter = "MULTI"
)

// Client is consul client config
type Client struct {
	dc     Datacenter
	cli    *api.Client
	ctx    context.Context
	cancel context.CancelFunc

	// resolve service entry endpoints
	resolver ServiceResolver
	// healthcheck time interval in seconds
	healthcheckInterval int
	// heartbeat enable heartbeat
	heartbeat bool
	// deregisterCriticalServiceAfter time interval in seconds
	deregisterCriticalServiceAfter int
	// serviceChecks  user custom checks
	serviceChecks api.AgentServiceChecks
}

func defaultResolver(_ context.Context, entries []*api.ServiceEntry) []*registry.ServiceInstance {
	services := make([]*registry.ServiceInstance, 0, len(entries))
	for _, entry := range entries {
		var version string
		for _, tag := range entry.Service.Tags {
			ss := strings.SplitN(tag, "=", 2)
			if len(ss) == 2 && ss[0] == "version" {
				version = ss[1]
			}
		}
		endpoints := make([]string, 0)
		for scheme, addr := range entry.Service.TaggedAddresses {
			if scheme == "lan_ipv4" || scheme == "wan_ipv4" || scheme == "lan_ipv6" || scheme == "wan_ipv6" {
				continue
			}
			endpoints = append(endpoints, addr.Address)
		}
		if len(endpoints) == 0 && entry.Service.Address != "" && entry.Service.Port != 0 {
			endpoints = append(endpoints, fmt.Sprintf("http://%s:%d", entry.Service.Address, entry.Service.Port))
		}
		services = append(services, &registry.ServiceInstance{
			ID:        entry.Service.ID,
			Name:      entry.Service.Service,
			Metadata:  entry.Service.Meta,
			Version:   version,
			Endpoints: endpoints,
		})
	}

	return services
}

// ServiceResolver is used to resolve service endpoints
type ServiceResolver func(ctx context.Context, entries []*api.ServiceEntry) []*registry.ServiceInstance

// Service get services from consul
func (c *Client) Service(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	if c.dc == MultiDatacenter {
		return c.multiDCService(ctx, service, index, passingOnly)
	}

	opts := &api.QueryOptions{
		WaitIndex:  index,
		WaitTime:   time.Second * 55,
		Datacenter: string(c.dc),
	}
	opts = opts.WithContext(ctx)

	if c.dc == SingleDatacenter {
		opts.Datacenter = ""
	}

	entries, meta, err := c.singleDCEntries(service, "", passingOnly, opts)
	if err != nil {
		return nil, 0, err
	}
	return c.resolver(ctx, entries), meta.LastIndex, nil
}

func (c *Client) multiDCService(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	opts := &api.QueryOptions{
		WaitIndex: index,
		WaitTime:  time.Second * 55,
	}
	opts = opts.WithContext(ctx)

	var instances []*registry.ServiceInstance

	dcs, err := c.cli.Catalog().Datacenters()
	if err != nil {
		return nil, 0, err
	}

	for _, dc := range dcs {
		opts.Datacenter = dc
		e, m, err := c.singleDCEntries(service, "", passingOnly, opts)
		if err != nil {
			return nil, 0, err
		}

		ins := c.resolver(ctx, e)
		for _, in := range ins {
			if in.Metadata == nil {
				in.Metadata = make(map[string]string, 1)
			}
			in.Metadata["dc"] = dc
		}

		instances = append(instances, ins...)
		opts.WaitIndex = m.LastIndex
	}

	return instances, opts.WaitIndex, nil
}

func (c *Client) singleDCEntries(service, tag string, passingOnly bool, opts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.cli.Health().Service(service, tag, passingOnly, opts)
}

// Register register service instance to consul
func (c *Client) Register(_ context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	addresses := make(map[string]api.ServiceAddress, len(svc.Endpoints))
	asrs := make([]*api.AgentServiceRegistration, 0, len(svc.Endpoints))
	for _, endpoint := range svc.Endpoints {
		// will separate the GRPC and HTTP services
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		addr := raw.Hostname()
		port, _ := strconv.ParseUint(raw.Port(), 10, 16)

		switch raw.Scheme {
		case "grpc":
			// Register gRPC service
			asr := &api.AgentServiceRegistration{
				ID:      svc.ID + "-grpc",
				Name:    svc.Name + "-grpc",
				Meta:    svc.Metadata,
				Tags:    []string{fmt.Sprintf("version=%s", svc.Version)},
				Address: addr,
				Port:    int(port),
			}
			if enableHealthCheck {
				asr.Checks = []*api.AgentServiceCheck{
					{
						GRPC:                           net.JoinHostPort(addr, strconv.FormatUint(port, 10)),
						Interval:                       fmt.Sprintf("%ds", c.healthcheckInterval),
						DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", c.deregisterCriticalServiceAfter),
						Timeout:                        "5s",
					},
				}
				asr.Checks = append(asr.Checks, c.serviceChecks...)
			}
			if c.heartbeat {
				asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
					CheckID:                        "service:" + svc.ID + "-grpc",
					TTL:                            fmt.Sprintf("%ds", c.healthcheckInterval*2),
					DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", c.deregisterCriticalServiceAfter),
				})
			}
			asrs = append(asrs, asr)
		case "http", "https":
			// Register HTTP service
			asr := &api.AgentServiceRegistration{
				ID:      svc.ID + "-http",
				Name:    svc.Name + "-http",
				Meta:    svc.Metadata,
				Tags:    []string{fmt.Sprintf("version=%s", svc.Version)},
				Address: addr,
				Port:    int(port),
			}

			if enableHealthCheck {
				asr.Checks = []*api.AgentServiceCheck{
					{
						HTTP:                           fmt.Sprintf("%s/api/v1/%s/health", endpoint, svc.Name),
						Interval:                       fmt.Sprintf("%ds", c.healthcheckInterval),
						DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", c.deregisterCriticalServiceAfter),
						Timeout:                        "5s",
						Method:                         "GET",
					},
				}
				asr.Checks = append(asr.Checks, c.serviceChecks...)
			}
			if c.heartbeat {
				asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
					CheckID:                        "service:" + svc.ID + "-http",
					TTL:                            fmt.Sprintf("%ds", c.healthcheckInterval*2),
					DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", c.deregisterCriticalServiceAfter),
				})
			}
			asrs = append(asrs, asr)
		}

		addresses[raw.Scheme] = api.ServiceAddress{Address: endpoint, Port: int(port)}
	}

	var registered []*api.AgentServiceRegistration
	var lastErr error
	for _, asr := range asrs {
		err := c.cli.Agent().ServiceRegister(asr)
		if err != nil {
			lastErr = err
			continue
		}
		registered = append(registered, asr)
	}
	if lastErr != nil {
		for _, asr := range registered {
			_ = c.cli.Agent().ServiceDeregister(asr.ID)
		}
		return lastErr
	}

	if c.heartbeat {
		for _, asr := range asrs {
			go func(svcID string, asr *api.AgentServiceRegistration) {
				time.Sleep(time.Second)
				err := c.cli.Agent().UpdateTTL("service:"+svcID, "pass", "pass")
				if err != nil {
					log.Errorf("[Consul]update ttl heartbeat to consul failed!err:=%v", err)
				}
				ticker := time.NewTicker(time.Second * time.Duration(c.healthcheckInterval))
				defer ticker.Stop()
				for {
					select {
					case <-c.ctx.Done():
						_ = c.cli.Agent().ServiceDeregister(svcID)
						return
					default:
					}
					select {
					case <-c.ctx.Done():
						_ = c.cli.Agent().ServiceDeregister(svcID)
						return
					case <-ticker.C:
						// ensure that unregistered services will not be re-registered by mistake
						if errors.Is(c.ctx.Err(), context.Canceled) || errors.Is(c.ctx.Err(), context.DeadlineExceeded) {
							_ = c.cli.Agent().ServiceDeregister(svcID)
							return
						}
						err = c.cli.Agent().UpdateTTLOpts("service:"+svcID, "pass", "pass", new(api.QueryOptions).WithContext(c.ctx))
						if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
							_ = c.cli.Agent().ServiceDeregister(svcID)
							return
						}
						if err != nil {
							log.Errorf("[Consul] update ttl heartbeat to consul failed! err=%v", err)
							// when the previous report fails, try to re register the service
							time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
							if err = c.cli.Agent().ServiceRegister(asr); err != nil {
								log.Errorf("[Consul] re registry service failed!, err=%v", err)
							} else {
								log.Warn("[Consul] re registry of service occurred success")
							}
						}
					}
				}
			}(asr.ID, asr)
		}
	}
	return nil
}

// Deregister service by service ID
func (c *Client) Deregister(_ context.Context, serviceID string) error {
	defer c.cancel()
	err := c.cli.Agent().ServiceDeregister(serviceID)
	if err != nil && !strings.HasPrefix(err.Error(), "Unexpected response code: 404") {
		// just mitigate the 404 error
		return err
	}
	return nil
}
