package httputil_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func Test_HandleDefaultErrors(t *testing.T) {
	t.Parallel()

	pgErr := &pgconn.PgError{
		Code:      "23505",
		Detail:    "duplicate key value violates unique constraint",
		TableName: "users",
	}
	tests := []test.Expected[error, test.Response[string]]{
		{
			Name: "context canceled",
			Args: context.Canceled,
			Expected: test.Response[string]{
				Body: httputil.ErrRequestCanceled.Error(),
				Code: http.StatusGatewayTimeout,
			},
		},
		{
			Name: "context deadline exceeded",
			Args: context.DeadlineExceeded,
			Expected: test.Response[string]{
				Body: httputil.ErrRequestCanceled.Error(),
				Code: http.StatusGatewayTimeout,
			},
		},
		{
			Name: "sql no rows",
			Args: sql.ErrNoRows,
			Expected: test.Response[string]{
				Body: httputil.ErrNotFound.Error(),
				Code: http.StatusNotFound,
			},
		},
		{
			Name: "pgx error",
			Args: pgErr,
			Expected: test.Response[string]{
				Body: httputil.ErrInternalServer.Error(),
				Code: http.StatusInternalServerError,
			},
		},
		{
			Name: "unexpected error",
			Args: httputil.ErrInternalServer,
			Expected: test.Response[string]{
				Body: httputil.ErrInternalServer.Error(),
				Code: http.StatusInternalServerError,
			},
		},
		{
			Name: "nil error",
			Args: nil,
			Expected: test.Response[string]{
				Body: httputil.ErrInternalServer.Error(),
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			mockLogger := logger.NewMockLogger(gomock.NewController(t))
			switch {
			case errors.Is(tt.Args, context.Canceled), errors.Is(tt.Args, context.DeadlineExceeded):
				mockLogger.EXPECT().Warn("request canceled", zap.Error(tt.Args))
			case errors.Is(tt.Args, sql.ErrNoRows):
				mockLogger.EXPECT().Warn("resource not found", zap.Error(tt.Args))
			case errors.Is(tt.Args, pgErr):
				mockLogger.EXPECT().
					Error("database error", gomock.Any(), gomock.Any(), gomock.Any(), zap.Error(tt.Args))
			default:
				mockLogger.EXPECT().Error("unexpected error", zap.Error(tt.Args))
			}

			ctx := logger.WithLogger(context.Background(), mockLogger)
			w := httptest.NewRecorder()

			httputil.HandleDefaultErrors(ctx, w, tt.Args)

			require.Equal(t, tt.Expected.Code, w.Code)
			require.Contains(t, w.Body.String(), tt.Expected.Body)
		})
	}
}
