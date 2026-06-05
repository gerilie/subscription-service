package logger_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/env"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
	"go.uber.org/zap"
)

func Test_New(t *testing.T) {
	t.Parallel()

	type args struct {
		level zap.AtomicLevel
		env   string
	}

	tests := []test.Expected[args, string]{
		{
			Name: "success",
			Args: args{
				level: zap.NewAtomicLevelAt(zap.DebugLevel),
				env:   env.Dev,
			},
			Expected: "debug",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			l, err := logger.New(logger.Config{Level: tt.Args.level}, tt.Args.env)

			require.NoError(t, err)
			require.Equal(t, tt.Expected, l.Zap().Level().String())
		})
	}
}
