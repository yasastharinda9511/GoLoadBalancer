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

func ParseLoadBalancerType(lbType string) (LoadBalancerType, error) {
	switch lbType {
	case string(WEIGHTEDLOADBALANCER):
		return WEIGHTEDLOADBALANCER, nil
	case string(ROUNDROBINLOADBALANCER):
		return ROUNDROBINLOADBALANCER, nil
	case string(RANDOMLOADBALANCER):
		return RANDOMLOADBALANCER, nil
	default:
		return "", errors.New("invalid load balancer type")
	}
}
