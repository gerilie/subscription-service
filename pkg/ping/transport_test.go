package ping_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/ping"
	"go.uber.org/mock/gomock"
)

func Test_Ping(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockLogger := logger.NewMockLogger(ctrl)
	ctx := logger.WithLogger(context.Background(), mockLogger)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	r = r.WithContext(ctx)

	mockLogger.EXPECT().Info("ping success")

	ping.Ping(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, httputil.JSON, w.Header().Get(httputil.ContentType))

	var resp ping.Resp
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.NotEmpty(t, resp.Timestamp)
	require.WithinDuration(t, time.Now(), resp.Timestamp, time.Second)
}
