package pool

type Backend struct {
	url    string
	weight int
	active bool
}

// NewBackend creates a new Backend instance
func NewBackend(url string, weight int) *Backend {
	return &Backend{
		url:    url,
		weight: weight,
		active: true,
	}
}

func NewBackendWithWeight(url string) *Backend {
	return &Backend{
		url:    url,
		weight: 1,
		active: true,
	}
}

func NewBackendWithActive(url string, weight int, active bool) *Backend {
	return &Backend{
		url:    url,
		weight: weight,
		active: active,
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
