package pool

import (
	"sync"
)

// PoolSelector maintains a map of pools
type PoolSelector struct {
	mu    sync.RWMutex
	pools map[string]*Pool
}

// NewPoolSelector creates a new PoolSelector
func NewPoolSelector() *PoolSelector {
	return &PoolSelector{
		pools: make(map[string]*Pool),
	}
}

// AddPool adds a new pool to the selector
func (ps *PoolSelector) AddPool(pool *Pool) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.pools[pool.GetID()] = pool
}

// GetPool retrieves a pool by its ID
func (ps *PoolSelector) GetPool(id string) (*Pool, bool) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	pool, exists := ps.pools[id]
	return pool, exists
}

// RemovePool removes a pool by its ID
func (ps *PoolSelector) RemovePool(id string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	delete(ps.pools, id)
}
