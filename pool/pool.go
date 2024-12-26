package pool

type Pool struct {
	backends     []Backend
	loadBalancer LoadBalancer
	id           string
}

func NewPool(id string, loadBalancer LoadBalancerType, backends []*Backend) *Pool {

	lb, _ := LoadBalancerFactory(loadBalancer, backends)
	return &Pool{
		backends:     []Backend{},
		loadBalancer: lb,
		id:           id,
	}
}

func (p *Pool) Next() (*Backend, error) {
	if p.loadBalancer == nil {
		return nil, nil
	}

	backend, err := p.loadBalancer.LoadBalance()
	return backend, err
}

func (p *Pool) AddBackend(backend Backend) {
	p.backends = append(p.backends, backend)
}

func (p *Pool) GetBackends() []Backend {
	return p.backends
}

func (p *Pool) GetID() string {
	return p.id
}
