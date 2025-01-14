package circuitBraker

import (
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state                State
	failureCount         int
	successCount         int
	maxFailures          int
	resetTimeout         time.Duration
	halfOpenMaxSuccesses int
	mutex                sync.Mutex
	lastFailureTime      time.Time
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration, halfOpenMaxSuccesses int) *CircuitBreaker {
	return &CircuitBreaker{
		state:                Closed,
		maxFailures:          maxFailures,
		resetTimeout:         resetTimeout,
		halfOpenMaxSuccesses: halfOpenMaxSuccesses,
	}
}

func (cb *CircuitBreaker) State() State {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Open:
		if time.Duration(time.Since(cb.lastFailureTime).Milliseconds()) > cb.resetTimeout {
			cb.state = HalfOpen
			break
		}
	}
	return cb.state
}

func (cb *CircuitBreaker) HandleFail() {

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Closed:
		cb.failureCount++
		if cb.failureCount > cb.maxFailures {
			cb.state = Open
			cb.lastFailureTime = time.Now()
			return
		}
	case Open:
		cb.lastFailureTime = time.Now()
	case HalfOpen:
		cb.state = Open
		cb.lastFailureTime = time.Now()
	}
}

func (cb *CircuitBreaker) HandleSuccess() {

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Open:
		cb.state = Closed
	case HalfOpen:
		cb.successCount++
		if cb.successCount >= cb.halfOpenMaxSuccesses {
			cb.state = Closed
			cb.successCount = 0
			cb.failureCount = 0
		}
	}
}
