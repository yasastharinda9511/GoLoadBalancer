package pool

import (
	"errors"
)

// LoadBalancerType represents the type of load balancer
type LoadBalancerType string

const (
	WEIGHTEDLOADBALANCER   LoadBalancerType = "WEIGHTEDLOADBALANCER"
	ROUNDROBINLOADBALANCER LoadBalancerType = "ROUNDROBINLOADBALANCER"
	RANDOMLOADBALANCER     LoadBalancerType = "RANDOMLOADBALANCER"
)

// LoadBalancerFactory returns the appropriate load balancer based on the type
func LoadBalancerFactory(lbType LoadBalancerType, backends []*Backend) (LoadBalancer, error) {
	switch lbType {
	case WEIGHTEDLOADBALANCER:
		lb, err := NewWeightedLoadBalancer(backends)
		if err != nil {
			return nil, err
		}
		return lb, nil
	default:
		return nil, errors.New("unknown load balancer type")
	}
}
