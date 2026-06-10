package middleware

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/ratelimiter"
	"go.uber.org/zap"
)

var (
	// ErrRateLimit indicates the client has exceeded their rate limit (HTTP 429).
	ErrRateLimit = fmt.Errorf("%w: too many requests", ErrMiddleware)

	// ErrGetClientIP indicates the client IP could not be determined (HTTP 500).
	ErrGetClientIP = fmt.Errorf("%w: failed to get client ip", ErrMiddleware)

	// ErrGetLimiter indicates the rate limiter could not be retrieved (HTTP 500).
	ErrGetLimiter = fmt.Errorf("%w: failed to get limiter", ErrMiddleware)
)

// RateLimiter is a middleware function that applies rate limiting to HTTP requests.
// It takes an http.Handler and an IP rate limiter as arguments and returns an http.Handler.
//
// The RateLimiter middleware checks if the request is rate limited and returns an error if it is.
func RateLimiter(next http.Handler, l ratelimiter.IPRateLimiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)

		ip := httputil.GetClientIP(r)
		if ip == "" {
			log.Error("internal middleware error", zap.Error(ErrGetClientIP))
			http.Error(w, httputil.ErrInternalServer.Error(), http.StatusInternalServerError)

			return
		}

		limiter := l.GetLimiter(ip)
		if limiter == nil {
			log.Error("internal middleware error", zap.Error(ErrGetLimiter))
			http.Error(w, httputil.ErrInternalServer.Error(), http.StatusInternalServerError)

			return
		}

		limit := strconv.Itoa(limiter.Burst())
		tokens := limiter.Tokens()
		remaining := fmt.Sprintf("%.0f", math.Floor(tokens))

		w.Header().Set("X-Ratelimit-Limit", limit)
		w.Header().Set("X-Ratelimit-Remaining", remaining)

		if !limiter.Allow() {
			log.Error("rate limit exceeded", zap.Error(ErrRateLimit))
			http.Error(w, ErrRateLimit.Error(), http.StatusTooManyRequests)

			return
		}

		next.ServeHTTP(w, r)
	})
}
