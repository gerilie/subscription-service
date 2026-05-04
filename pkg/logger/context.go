package logger

import (
	"context"
)

type contextKey struct{}

var (
	loggerKey    = contextKey{}
	requestIDKey = contextKey{}
)

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func FromContext(ctx context.Context) Logger {
	l, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		return nil
	}

	return l
}
