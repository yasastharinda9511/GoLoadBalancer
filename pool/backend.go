package pool

import "strings"

type Protocol string

const (
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
	NONE  Protocol = "NONE"
)

type Backend struct {
	url      string
	weight   int
	active   bool
	protocol Protocol
}

// NewBackend creates a new Backend instance
func NewBackend(url string, weight int) *Backend {

	return &Backend{
		url:      url,
		weight:   weight,
		active:   true,
		protocol: deriveProtocol(url),
	}
}

func NewBackendWithWeight(url string) *Backend {
	return &Backend{
		url:      url,
		weight:   1,
		active:   true,
		protocol: deriveProtocol(url),
	}
}

func NewBackendWithActive(url string, weight int, active bool) *Backend {
	return &Backend{
		url:      url,
		weight:   weight,
		active:   active,
		protocol: deriveProtocol(url),
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

func deriveProtocol(url string) Protocol {
	parts := strings.Split(url, "://")
	if len(parts) > 1 && parts[0] == "https" {
		return HTTPS
	} else if len(parts) > 1 && parts[0] == "http" {
		return HTTP
	}

	return NONE
}
