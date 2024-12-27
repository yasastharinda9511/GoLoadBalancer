package errors

import "fmt"

// RuleNotFoundError represents an error when a rule is not found.
type RuleNotFoundError struct {
	uid string
}

// Error implements the error interface for RuleNotFoundError.
func (e *RuleNotFoundError) Error() string {
	return fmt.Sprintf("rule is not found for %s message", e.uid)
}

// NewRuleNotFoundError creates a new RuleNotFoundError.
func NewRuleNotFoundError(uid string) *RuleNotFoundError {
	return &RuleNotFoundError{uid: uid}
}
