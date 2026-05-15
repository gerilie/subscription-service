package ratelimiter

import (
	"context"
	"time"
)

// StartCleanUp starts a background cleanup routine that periodically removes
// inactive IP limiters from internal storage.
//
// Cleanup behavior is configured through the limiter Config:
//
//   - CleanUpInterval defines how often cleanup runs.
//   - CleanUpMaxIdle defines how long a limiter may remain inactive
//     before removal.
//
// The cleanup routine runs until the provided context is canceled.
//
// Multiple calls to StartCleanUp are safe, but the cleanup routine
// will only be started once.
func (l *ipRateLimiter) StartCleanUp(ctx context.Context) {
	l.cleanUpOnce.Do(func() {
		ticker := time.NewTicker(l.CleanUpInterval)

		go func() {
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					l.cleanUp(l.CleanUpMaxIdle)
				case <-ctx.Done():
					return
				}
			}
		}()
	})
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
