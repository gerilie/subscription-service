package httputil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_GetClientIP_XRealIP_Success(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Real-IP", "1.2.3.4:1234")

	require.Equal(t, "1.2.3.4", httputil.GetClientIP(r))
}

func Test_GetClientIP_XRealIP_Empty(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Real-IP", "")

	require.NotZero(t, httputil.GetClientIP(r))
}

func Test_GetClientIP_RemoteAddr(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.RemoteAddr = "5.6.7.8:1234"

	require.Equal(t, "5.6.7.8", httputil.GetClientIP(r))

	tests := []test.Expected[string, string]{
		{
			Name:     "IPv4 with port",
			Args:     "5.6.7.8:1234",
			Expected: "5.6.7.8",
		},
		{
			Name:     "IPv4 without port",
			Args:     "5.6.7.8",
			Expected: "5.6.7.8",
		},
		{
			Name:     "IPv6",
			Args:     "[5be8:dde9:7f0b:d5a7:bd01:b3be:9c69:573b]:1234",
			Expected: "5be8:dde9:7f0b:d5a7:bd01:b3be:9c69:573b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.RemoteAddr = tt.Args

			require.Equal(t, tt.Expected, httputil.GetClientIP(r))
		})
	}
}
