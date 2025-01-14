package errors

import "fmt"

// BackendError represents an error that occurs during backend operations.
type BackendsNotFoundError struct {
	poolId string
}

// Error implements the error interface for BackendError.
func (e *BackendsNotFoundError) Error() string {
	return fmt.Sprintf("No backends are foun in %s", e.poolId)
}

// NewBackendError creates a new BackendError with the given message and code.
func NewBackendsNotFoundError(poolId string) *BackendsNotFoundError {
	return &BackendsNotFoundError{
		poolId: poolId,
	}
}
