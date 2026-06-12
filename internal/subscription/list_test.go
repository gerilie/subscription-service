//go:build integration

package subscription_test

import (
	"net/http"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
)

func Test_ListSubscriptions_Success(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	client, baseURL := setupTestServer(t)

	body := subscription.SubReq{
		ServiceName: "test",
		Price:       100,
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		StartDate:   "01-2001",
	}

	for range 20 {
		result := &subscription.SubResp{}
		res, err := client.R().
			SetBody(body).
			SetResult(result).
			Post(baseURL)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode())
	}

	result := subscription.SubListResp{}
	res, err := client.R().
		SetQueryParams(map[string]string{
			"page":  "1",
			"limit": "10",
		}).
		SetResult(&result).
		Get(baseURL)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.Len(t, result, 10)
}
