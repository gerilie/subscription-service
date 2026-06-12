//go:build integration

package subscription_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/internal/subscription"
)

func Test_CreateSubscription_Success(t *testing.T) {
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
	expected := &subscription.SubResp{
		ID:          1,
		ServiceName: "test",
		Price:       100,
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		StartDate:   "01-2001",
	}

	result := &subscription.SubResp{}
	res, err := client.R().
		SetBody(body).
		SetResult(result).
		Post(baseURL)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, res.StatusCode())
	require.Equal(t, expected, result)

	res, err = client.R().
		SetResult(result).
		Get(fmt.Sprintf("%s/%d", baseURL, result.ID))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.Equal(t, expected, result)
}

func Test_CreateSubscription_BadRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	client, baseURL := setupTestServer(t)

	body := subscription.SubReq{
		ServiceName: "test",
		Price:       -100,
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		StartDate:   "01-01-2001",
	}

	result := &subscription.SubResp{}
	res, err := client.R().
		SetBody(body).
		SetResult(result).
		Post(baseURL)

	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, res.StatusCode())
}
