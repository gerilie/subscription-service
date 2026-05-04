package ping

import (
	"net/http"
	"time"

	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"go.uber.org/zap"
)

// Ping is an HTTP handler that returns a ping response with a timestamp.
//
// It returns a JSON response containing a timestamp field.
// The response is in JSON format and includes a timestamp field.
//
// The handler logs an error if the JSON response cannot be written to the response writer.
// It also logs a success message if the ping response is successfully written.
func Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	resp := pingResp{
		Timestamp: time.Now().String(),
	}

	if err := httputil.WriteJSON(ctx, w, http.StatusOK, resp); err != nil {
		log.Error(ctx, "ping failed", zap.Error(err))

		return
	}

	log.Info(ctx, "ping success")
}
