package logger_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_LoadConfig_Default(t *testing.T) {
	t.Parallel()

	cfg, err := logger.LoadConfig()

	require.NoError(t, err)
	require.Equal(t, "info", cfg.Level.String())
}

func Test_LoadConfig_Custom(t *testing.T) {
	tests := []test.Expected[string, string]{
		{
			Name:     "debug level",
			Args:     "debug",
			Expected: "debug",
		},
		{
			Name:     "info level",
			Args:     "info",
			Expected: "info",
		},
		{
			Name:     "warn level",
			Args:     "warn",
			Expected: "warn",
		},
		{
			Name:     "error level",
			Args:     "error",
			Expected: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Setenv("LEVEL", tt.Args)

			cfg, err := logger.LoadConfig()

			require.NoError(t, err)
			require.Equal(t, tt.Expected, cfg.Level.String())
		})
	}
}

func TestLoadConfig_EmptyValue(t *testing.T) {
	t.Setenv("LEVEL", "")

	cfg, err := logger.LoadConfig()

	require.NoError(t, err)
	require.NotEmpty(t, cfg.Level.String())
}
