package subscription_test

import (
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_LoadConfig_Success(t *testing.T) {
	type args map[string]string

	tests := []test.Expected[args, subscription.Config]{
		{
			Name: "required envs",
			Args: args{
				"HOST": "localhost",
				"PORT": "8080",
			},
			Expected: subscription.Config{
				Env:                 "prod",
				Host:                "localhost",
				Port:                "8080",
				ReadHTO:             5 * time.Second,
				ReadTO:              15 * time.Second,
				WriteTO:             30 * time.Second,
				IdleTO:              120 * time.Second,
				RLRequestsPerSecond: 10,
				RLBurst:             30,
				RLCleanUpInterval:   5 * time.Minute,
				RLCLeanUpMaxIdle:    30 * time.Minute,
			},
		},
		{
			Name: "override envs",
			Args: args{
				"ENV":                            "env",
				"HOST":                           "localhost",
				"PORT":                           "8080",
				"READ_HEADER_TIMEOUT":            "10s",
				"READ_TIMEOUT":                   "20s",
				"WRITE_TIMEOUT":                  "50s",
				"IDLE_TIMEOUT":                   "200s",
				"RATE_LIMIT_REQUESTS_PER_SECOND": "20",
				"RATE_LIMIT_BURST":               "40",
				"RATE_LIMIT_CLEANUP_INTERVAL":    "10m",
				"RATE_LIMIT_CLEANUP_MAX_IDLE":    "20m",
			},
			Expected: subscription.Config{
				Env:                 "env",
				Host:                "localhost",
				Port:                "8080",
				ReadHTO:             10 * time.Second,
				ReadTO:              20 * time.Second,
				WriteTO:             50 * time.Second,
				IdleTO:              200 * time.Second,
				RLRequestsPerSecond: 20,
				RLBurst:             40,
				RLCleanUpInterval:   10 * time.Minute,
				RLCLeanUpMaxIdle:    20 * time.Minute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			for k, v := range tt.Args {
				t.Setenv(k, v)
			}

			cfg, err := subscription.LoadConfig()

			require.NoError(t, err)

			require.Equal(t, tt.Expected, cfg)
		})
	}
}

func Test_LoadConfig_EmptyRequiredEnvs(t *testing.T) {
	t.Parallel()

	cfg, err := subscription.LoadConfig()

	require.Error(t, err)
	require.Empty(t, cfg)
}
