package test

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// CreatePostgresContainer starts a PostgreSQL 18.4 testcontainer with pre-configured
// credentials (test/test/test) and basic health check. The container is ready when
// PostgreSQL accepts connections on port 5432.
//
// The caller must terminate the container after use to prevent resource leaks:
//
//	container, err := CreatePostgresContainer(ctx)
//	if err != nil {
//	    t.Fatal(err)
//	}
//	defer container.Terminate(ctx)
func CreatePostgresContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	container, err := postgres.Run(
		ctx,
		"postgres:18.4-bookworm",
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, err
	}

	return container, nil
}
