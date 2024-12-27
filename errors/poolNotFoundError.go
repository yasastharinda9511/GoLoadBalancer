package errors

import "fmt"

// PoolNotFoundError represents an error when a pool is not found.
type PoolNotFoundError struct {
	PoolID string
}

// Error implements the error interface for PoolNotFoundError.
func (e *PoolNotFoundError) Error() string {
	return fmt.Sprintf("pool with ID %s not found", e.PoolID)
}

// NewPoolNotFoundError creates a new PoolNotFoundError.
func NewPoolNotFoundError(poolID string) error {
	return &PoolNotFoundError{PoolID: poolID}
}
