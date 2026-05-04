package deferfunc

import (
	"context"
	"errors"
	"syscall"

	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"go.uber.org/zap"
)

// Close executes the provided cleanup function and logs any error that occurs,
// except for known benign errors (such as syscall.EINVAL).
//
// It is intended to be used with defer to safely handle resource cleanup
// (e.g. closing files, network connections, request bodies) while ensuring
// errors are properly logged using the logger from context.
func Close(ctx context.Context, c func() error, errMsg string) {
	log := logger.FromContext(ctx)

	if err := c(); err != nil && !errors.Is(err, syscall.EINVAL) {
		log.Error(ctx, errMsg, zap.Error(err))
	}
}
