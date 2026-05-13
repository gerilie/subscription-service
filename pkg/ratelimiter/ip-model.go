package ratelimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter extends RateLimiter with IP-based rate limiting capabilities.
type IPRateLimiter interface {
	RateLimiter

	// GetLimiter returns a rate limiter associated with the given IP address.
	// If the IP is empty, a new independent limiter is created.
	GetLimiter(ip string) *rate.Limiter
}

// ipRateLimiter implements IPRateLimiter and stores per-IP limiters.
type ipRateLimiter struct {
	rateLimiter

	ips map[string]*rate.Limiter
}

// NewIPRateLimiter creates a new IP-based rate limiter.
//
// If r or b are less than or equal to zero, they are replaced with safe defaults.
func NewIPRateLimiter(r rate.Limit, b int) IPRateLimiter {
	if r <= 0 {
		r = rate.Limit(1)
	}
	if b <= 0 {
		b = 1
	}

	return &ipRateLimiter{
		rateLimiter: rateLimiter{
			mu: &sync.RWMutex{},
			r:  r,
			b:  b,
		},
		ips: make(map[string]*rate.Limiter),
	}
}
