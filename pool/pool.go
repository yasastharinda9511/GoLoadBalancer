package pool

import (
	"io"
	"time"

	"github.com/yasastharinda9511/go_gateway_api/dispatcher"
	"github.com/yasastharinda9511/go_gateway_api/errors"
	"github.com/yasastharinda9511/go_gateway_api/message"
)

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

func (p *Pool) HandleBackendCall(requestMessage *message.HttpRequestMessage) (int, []byte, error) {

	backend, err := p.Next()
	if err != nil {
		return -1, nil, err
	}

	dispatch := dispatcher.NewHTTPDispatcher(10 * time.Second)
	resp, err := dispatch.CallBackend(dispatcher.GET, backend.GetURL(), requestMessage.GetHeaders(), requestMessage.GetQueryParams())

	if err != nil {
		// Write an error response
		err = errors.NewBackendError(backend.GetURL(), err.Error())
		return -1, nil, err
	}
	defer resp.Body.Close()

	// Read backend response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.NewBackendError(backend.GetURL(), err.Error())
		return -1, nil, err
	}

	statusCode := resp.StatusCode
	return statusCode, body, nil
}
