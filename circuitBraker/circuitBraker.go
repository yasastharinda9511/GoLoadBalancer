package circuitBraker

import (
	"fmt"
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
	case Closed:
		fmt.Println("updated to closed state")
	case Open:
		fmt.Println("time since last failure ", time.Since(cb.lastFailureTime).Milliseconds())
		if time.Duration(time.Since(cb.lastFailureTime).Milliseconds()) > cb.resetTimeout {
			fmt.Println("updated to half open state")
			cb.state = HalfOpen
			break
		}

		fmt.Println("updated to open state")
	case HalfOpen:
		fmt.Println("updated to half open state")
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
			fmt.Println("After Failing open state")
			cb.state = Open
			cb.lastFailureTime = time.Now()
			return
		}
		fmt.Println("After failing closed state")
	case Open:
		fmt.Println("After Failing open state  ", cb.state)
		cb.lastFailureTime = time.Now()
	case HalfOpen:
		fmt.Println("After Failing open state")
		cb.state = Open
		cb.lastFailureTime = time.Now()
	}
	fmt.Print("  ")
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
