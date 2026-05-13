package ratelimiter

import "golang.org/x/time/rate"

// CleanUp resets all stored IP rate limiters.
func (l *ipRateLimiter) CleanUp() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.ips = make(map[string]*rate.Limiter)
}
