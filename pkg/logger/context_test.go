package logger_test

import (
	"context"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	gomock "go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func Test_Context(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockLogger := logger.NewMockLogger(ctrl)

	ctx := logger.WithLogger(context.Background(), mockLogger)
	extracted := logger.FromContext(ctx)

	require.NotNil(t, extracted)
	require.Equal(t, mockLogger, extracted)
}

func Test_NoopLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	extracted := logger.FromContext(ctx)

	require.NotNil(t, extracted)
}

func Test_RequestID(t *testing.T) {
	t.Parallel()

	mockLogger := logger.NewMockLogger(gomock.NewController(t))
	ctx := logger.WithLogger(context.Background(), mockLogger)

	mockLogger.EXPECT().Zap().Return(zap.NewNop())

	ctx = logger.WithRequestID(ctx, "request-id")

	require.Equal(t, "request-id", logger.RequestIDFromContext(ctx))
}

func Test_RequestID_NoopLogger(t *testing.T) {
	t.Parallel()

	ctx := logger.WithRequestID(context.Background(), "request-id")

	require.Equal(t, "request-id", logger.RequestIDFromContext(ctx))
}

func Test_RequestID_Repeated(t *testing.T) {
	t.Parallel()

	mockLogger := logger.NewMockLogger(gomock.NewController(t))
	ctx := logger.WithLogger(context.Background(), mockLogger)

	mockLogger.EXPECT().Zap().Return(zap.NewNop())

	ctx = logger.WithRequestID(ctx, "request-id")

	require.Equal(t, "request-id", logger.RequestIDFromContext(ctx))

	ctx = logger.WithRequestID(ctx, "request-ID")

	require.Equal(t, "request-ID", logger.RequestIDFromContext(ctx))
}

func Test_RequestID_WithoutID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	require.Equal(t, "", logger.RequestIDFromContext(ctx))
}
