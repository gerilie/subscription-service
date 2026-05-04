package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey struct{}

var loggerKey = contextKey{}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func WithRequestID(ctx context.Context, id string) context.Context {
	l := FromContext(ctx)
	if l == nil {
		return ctx
	}

	l = l.With(zap.String("request_id", id))

	return WithLogger(ctx, l)
}

func FromContext(ctx context.Context) Logger {
	l, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		return nil
	}

	return l
}
