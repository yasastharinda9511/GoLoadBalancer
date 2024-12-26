package pool

// LoadBalancer defines the interface for load balancing strategies.
type LoadBalancer interface {
	// Select chooses a server from the pool based on the load balancing strategy.
	LoadBalance() (*Backend, error)
}
