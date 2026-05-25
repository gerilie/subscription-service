package postgres_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/postgres"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_LoadConfig_Success(t *testing.T) {
	tests := []test.Expected[postgres.Config, postgres.Config]{
		{
			Name: "success",
			Args: postgres.Config{
				Host:     "postgres",
				Port:     "5432",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
			Expected: postgres.Config{
				Host:     "postgres",
				Port:     "5432",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
		},
		{
			Name: "default env",
			Args: postgres.Config{
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
			Expected: postgres.Config{
				Host:     "postgres",
				Port:     "5432",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
		},
		{
			Name: "overriding default env",
			Args: postgres.Config{
				Host:     "Postgres",
				Port:     "5433",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
			Expected: postgres.Config{
				Host:     "Postgres",
				Port:     "5433",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Setenv("HOST", tt.Args.Host)
			t.Setenv("PORT", tt.Args.Port)
			t.Setenv("USER", tt.Args.User)
			t.Setenv("PASSWORD", tt.Args.Password)
			t.Setenv("DB", tt.Args.DB)

			cfg, err := postgres.LoadConfig()
			require.NoError(t, err)

			require.Equal(t, tt.Expected, cfg)
		})
	}
}

func Test_LoadConfig_Fail(t *testing.T) {
	t.Setenv("HOST", "postgres")
	t.Setenv("PORT", "5432")
	t.Setenv("USER", "postgres")
	t.Setenv("DB", "postgres")

	_, err := postgres.LoadConfig()
	require.Error(t, err)
}
