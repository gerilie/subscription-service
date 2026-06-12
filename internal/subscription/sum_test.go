//go:build integration

package subscription_test

import (
	"net/http"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
)

func Test_SumSubscriptions_Success(t *testing.T) {
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

	for range 5 {
		result := &subscription.SubResp{}
		res, err := client.R().
			SetBody(body).
			SetResult(result).
			Post(baseURL)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode())
	}

	result := subscription.SubSumResp{}
	res, err := client.R().
		SetQueryParams(map[string]string{
			"start_date": "01-2000",
			"end_date":   "02-2001",
		}).
		SetResult(&result).
		Get(baseURL + "/sum")

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.Equal(t, 500, result.TotalPrice)
}
