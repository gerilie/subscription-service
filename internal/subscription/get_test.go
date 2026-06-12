//go:build integration

package subscription_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
)

func Test_GetSubscription_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	client, baseURL := setupTestServer(t)

	result := &subscription.SubResp{}
	res, err := client.R().
		SetResult(result).
		Get(fmt.Sprintf("%s/%d", baseURL, result.ID))

	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, res.StatusCode())
}
