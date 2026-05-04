package logger

import (
	"context"

	"go.uber.org/zap"
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

func addID(ctx context.Context, fields []zap.Field) []zap.Field {
	id, ok := ctx.Value(requestIDKey).(string)
	if ok {
		fields = append(fields, zap.String("request_id", id))
	}

	return fields
}
