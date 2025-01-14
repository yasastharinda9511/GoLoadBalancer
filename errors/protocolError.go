package errors

import "fmt"

// RuleNotFoundError represents an error when a rule is not found.
type ProtocolError struct {
	backend string
}

// Error implements the error interface for RuleNotFoundError.
func (e *ProtocolError) Error() string {
	return fmt.Sprintf("Undefined protocol in %s backend", e.backend)
}

// NewRuleNotFoundError creates a new RuleNotFoundError.
func NewProtocolError(backend string) *ProtocolError {
	return &ProtocolError{backend: backend}
}
