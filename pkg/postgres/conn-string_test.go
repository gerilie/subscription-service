package postgres_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/postgres"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_BuildConnString(t *testing.T) {
	t.Parallel()

	//nolint:gosec
	tests := []test.Expected[postgres.Config, string]{
		{
			Name: "success",
			Args: postgres.Config{
				Host:     "postgres",
				Port:     "5432",
				User:     "postgres",
				Password: "postgres",
				DB:       "postgres",
			},
			Expected: "postgres://postgres:postgres@postgres:5432/postgres",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			connString := postgres.BuildConnString(tt.Args)

			require.Equal(t, tt.Expected, connString)
		})
	}
}
