package pool

import (
	"math/rand"
	"sync"
)

type RandomLoadBalancer struct {
	backends []*Backend
	mu       sync.Mutex
}

// UpdateBackends implements LoadBalancer.
func (r *RandomLoadBalancer) UpdateBackends(backends []*Backend) {
	r.backends = backends
}

func NewRandomLoadBalancer(backends []*Backend) (*RandomLoadBalancer, error) {

	return &RandomLoadBalancer{
		backends: backends,
	}, nil

}

func (r *RandomLoadBalancer) LoadBalance() (*Backend, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	size := len(r.backends)
	index := rand.Intn(size)
	return r.backends[index], nil
}
