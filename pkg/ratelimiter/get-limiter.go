package ratelimiter

import "golang.org/x/time/rate"

// GetLimiter returns a rate limiter for the given IP address.
// If the limiter does not exist, it is created and stored.
//
// If ip is empty, a new standalone limiter is returned.
func (l *ipRateLimiter) GetLimiter(ip string) *rate.Limiter {
	if ip == "" {
		return rate.NewLimiter(l.r, l.b)
	}

	l.mu.RLock()
	limiter := l.ips[ip]
	l.mu.RUnlock()

	if limiter != nil {
		return limiter
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	limiter = l.ips[ip]
	if limiter == nil {
		limiter = rate.NewLimiter(l.r, l.b)
		l.ips[ip] = limiter
	}

	return limiter
}
