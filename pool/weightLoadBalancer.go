package pool

import (
	"errors"
	"sync"
)

// WeightedLoadBalancer is a weighted load balancer
type WeightedLoadBalancer struct {
	backends     []*Backend
	toltalWeight int
	current      int
	mutex        sync.Mutex
}

// NewWeightedLoadBalancer creates a new WeightedLoadBalancer
func NewWeightedLoadBalancer(backends []*Backend) (*WeightedLoadBalancer, error) {
	totalWeight := 0
	for _, backend := range backends {
		totalWeight += backend.GetWeight()
	}

	return &WeightedLoadBalancer{
		backends:     backends,
		toltalWeight: totalWeight,
		current:      -1,
	}, nil
}

// Next returns the next server based on the weights
func (lb *WeightedLoadBalancer) LoadBalance() (*Backend, error) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if lb.toltalWeight == 0 {
		return nil, errors.New("no servers available")
	}

	lb.current = (lb.current + 1) % lb.toltalWeight

	currentWeight := lb.current
	for _, backend := range lb.backends {
		if currentWeight < backend.weight {
			return backend, nil
		}
		currentWeight -= backend.GetWeight()
	}

	return nil, errors.New("no servers available")
}
