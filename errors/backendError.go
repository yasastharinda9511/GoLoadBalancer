package errors

import "fmt"

// BackendError represents an error that occurs during backend operations.
type BackendError struct {
	url     string
	message string
}

// Error implements the error interface for BackendError.
func (e *BackendError) Error() string {
	return fmt.Sprintf("%s from  %s", e.message, e.url)
}

// NewBackendError creates a new BackendError with the given message and code.
func NewBackendError(url string, message string) *BackendError {
	return &BackendError{
		message: message,
		url:     url,
	}
}
