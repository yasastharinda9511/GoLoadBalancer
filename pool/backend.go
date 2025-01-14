package pool

import (
	"strings"

	"github.com/yasastharinda9511/go_gateway_api/circuitBraker"
)

type Protocol string

const (
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
	NONE  Protocol = "NONE"
)

type Backend struct {
	url           string
	weight        int
	active        bool
	protocol      Protocol
	circuitBraker *circuitBraker.CircuitBreaker
}

// NewBackend creates a new Backend instance
func NewBackend(url string, weight int) *Backend {

	return &Backend{
		url:           url,
		weight:        weight,
		active:        true,
		protocol:      deriveProtocol(url),
		circuitBraker: circuitBraker.NewCircuitBreaker(5, 5000, 5),
	}
}

func NewBackendWithWeight(url string) *Backend {
	return &Backend{
		url:           url,
		weight:        1,
		active:        true,
		protocol:      deriveProtocol(url),
		circuitBraker: circuitBraker.NewCircuitBreaker(5, 5000, 5),
	}
}

func NewBackendWithActive(url string, weight int, active bool) *Backend {
	return &Backend{
		url:           url,
		weight:        weight,
		active:        active,
		protocol:      deriveProtocol(url),
		circuitBraker: circuitBraker.NewCircuitBreaker(5, 5000, 5),
	}
}

// SetActive sets the backend's active status
func (b *Backend) SetActive(active bool) {
	b.active = active
}

// GetURL returns the backend's URL
func (b *Backend) GetURL() string {
	return b.url
}

// GetWeight returns the backend's weight
func (b *Backend) GetWeight() int {
	return b.weight
}

// IsActive returns the backend's active status
func (b *Backend) IsActive() bool {
	return b.active
}

func (b *Backend) GetProtocol() Protocol {
	return b.protocol
}

func (b *Backend) GetCircuitBreaker() *circuitBraker.CircuitBreaker {
	return b.circuitBraker
}

func deriveProtocol(url string) Protocol {
	parts := strings.Split(url, "://")
	if len(parts) > 1 && parts[0] == "https" {
		return HTTPS
	} else if len(parts) > 1 && parts[0] == "http" {
		return HTTP
	}

	return NONE
}
