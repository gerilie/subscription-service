package middleware

import (
	"net/http"

	"github.com/yushafro/effective-mobile-tz/pkg/logger"
)

// NoBody is a middleware function that checks if the request body is empty.
// It takes an http.Handler as an argument and returns an http.Handler.
//
// The NoBody middleware returns an error if the request body is not empty.
func NoBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)

		if r.ContentLength != 0 {
			log.Error(ctx, "request body must be empty")
			http.Error(w, "request body must be empty", http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	})
}
