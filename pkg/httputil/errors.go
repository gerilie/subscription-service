package httputil

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"go.uber.org/zap"
)

var (
	// ErrNotFound indicates that the requested resource was not found.
	ErrNotFound = errors.New("resource not found")

	// ErrRequestCanceled indicates that the request was canceled by the client.
	ErrRequestCanceled = errors.New("request canceled")

	// ErrInternalServer indicates an unexpected internal server error occurred.
	ErrInternalServer = errors.New("internal server error")
)

// HandleDefaultErrors maps common application errors to HTTP responses
// and logs them using the logger stored in the context.
func HandleDefaultErrors(ctx context.Context, w http.ResponseWriter, err error) {
	var pgErr *pgconn.PgError
	log := logger.FromContext(ctx)

	switch {
	case errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded):
		log.Warn(ErrRequestCanceled.Error(), zap.Error(err))
		http.Error(w, ErrRequestCanceled.Error(), http.StatusGatewayTimeout)

	case errors.Is(err, sql.ErrNoRows):
		log.Warn(ErrNotFound.Error(), zap.Error(err))
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)

	case errors.As(err, &pgErr):
		log.Error("database error",
			zap.String("code", pgErr.Code),
			zap.String("detail", pgErr.Detail),
			zap.String("table", pgErr.TableName),
			zap.Error(err))
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)

	default:
		log.Error("unexpected error", zap.Error(err))
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)
	}
}
