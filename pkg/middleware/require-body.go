package middleware

import (
	"fmt"
	"net/http"

	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"go.uber.org/zap"
)

// ErrRequireBody indicates the request body is required (HTTP 400).
var ErrRequireBody = fmt.Errorf("%w: request body must not be empty", ErrMiddleware)

// RequireBody is a middleware function that checks if the request body is not empty.
// It takes an http.Handler as an argument and returns an http.Handler.
//
// The RequireBody middleware returns an error if the request body is empty.
func RequireBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)

		if r.ContentLength == 0 {
			log.Error("request body validation failed", zap.Error(ErrRequireBody))
			http.Error(w, ErrRequireBody.Error(), http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	})
}
