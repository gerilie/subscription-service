//go:build integration

package subscription_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
	"github.com/yushafro/effective-mobile-tz/pkg/env"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
	"resty.dev/v3"
)

func setupTestServer(t *testing.T) (*resty.Client, string) {
	t.Helper()

	ctx := context.Background()
	pc, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err)

	client := resty.New()

	t.Cleanup(func() {
		require.NoError(t, client.Close())

		pc.Pool.Close()
		require.NoError(t, pc.Container.Terminate(ctx))
	})

	t.Setenv("HOST", "localhost")
	t.Setenv("PORT", "8090")
	serverCfg, err := subscription.LoadConfig()
	require.NoError(t, err)

	loggerCfg, err := logger.LoadConfig()
	require.NoError(t, err)
	logger, err := logger.New(loggerCfg, env.Dev)
	require.NoError(t, err)

	repo := subscription.NewPGRepository(pc.Pool)
	service := subscription.NewService(repo)
	s := subscription.NewServer(service, serverCfg, logger)
	ts := httptest.NewServer(s.Handler())

	return client, ts.URL + "/subscriptions"
}
