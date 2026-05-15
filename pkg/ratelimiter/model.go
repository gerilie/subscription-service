package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// IPRateLimiter provides per-IP rate limiting functionality.
//
// Implementations are responsible for storing and managing
// independent rate limiters for different client IP addresses.
type IPRateLimiter interface {
	// GetLimiter returns a rate limiter associated with the given IP address.
	//
	// If a limiter for the provided IP does not exist, it is created.
	// If ip is empty, a new standalone limiter is returned without storing it.
	GetLimiter(ip string) *rate.Limiter

	// StartCleanUp starts a background cleanup routine that periodically removes
	// inactive IP limiters from internal storage.
	//
	// interval defines how often cleanup runs.
	// maxIdle defines how long an IP limiter may remain inactive before removal.
	StartCleanUp(interval, maxIdle time.Duration)
}

// client stores a rate limiter and metadata associated with a single IP address.
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// ipRateLimiter implements IPRateLimiter using an in-memory map of IP addresses.
type ipRateLimiter struct {
	mu *sync.Mutex
	r  rate.Limit
	b  int

	ips map[string]*client
}

// NewIPRateLimiter creates a new in-memory IP rate limiter.
//
// If r or b are less than or equal to zero, safe default values are used:
//
//	r = 1 request per second
//	b = 1 burst size
func NewIPRateLimiter(r rate.Limit, b int) IPRateLimiter {
	if r <= 0 {
		r = rate.Limit(1)
	}
	if b <= 0 {
		b = 1
	}

	return &ipRateLimiter{
		mu:  &sync.Mutex{},
		r:   r,
		b:   b,
		ips: make(map[string]*client),
	}
}
