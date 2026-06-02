package middleware

import (
	"fmt"
	"net/http"

	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"go.uber.org/zap"
)

// ErrNoBody indicates the request body must be empty (HTTP 400).
var ErrNoBody = fmt.Errorf("%w: request body must be empty", ErrMiddleware)

// NoBody is a middleware function that checks if the request body is empty.
// It takes an http.Handler as an argument and returns an http.Handler.
//
// The NoBody middleware returns an error if the request body is not empty.
func NoBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)

		if r.ContentLength > 0 {
			log.Error("request body validation failed", zap.Error(ErrNoBody))
			http.Error(w, ErrNoBody.Error(), http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	})
}
