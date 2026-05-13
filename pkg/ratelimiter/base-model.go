package ratelimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// RateLimiter defines a basic interface for rate limiter implementations
// that support cleanup of internal state.
type RateLimiter interface {
	CleanUp()
}

// rateLimiter contains shared configuration for rate limiting.
type rateLimiter struct {
	mu *sync.RWMutex
	r  rate.Limit
	b  int
}
