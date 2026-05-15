package ratelimiter

import "time"

// StartCleanUp starts a background cleanup routine that periodically removes
// inactive IP limiters from internal storage.
//
// interval defines how often cleanup runs.
// maxIdle defines how long a limiter may remain inactive before removal.
func (l *ipRateLimiter) StartCleanUp(interval, maxIdle time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			l.cleanUp(maxIdle)
		}
	}()
}

// cleanUp removes IP limiters that have been inactive longer than maxIdle.
func (l *ipRateLimiter) cleanUp(maxIdle time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for ip, c := range l.ips {
		if time.Since(c.lastSeen) > maxIdle {
			delete(l.ips, ip)
		}
	}
}
