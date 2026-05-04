package logger

import (
	"context"

	"go.uber.org/zap"
)

func (l *logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}
