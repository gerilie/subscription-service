package logger

import "go.uber.org/zap"

func (l *logger) With(fields ...zap.Field) Logger {
	return &logger{
		l: l.l.With(fields...),
	}
}
