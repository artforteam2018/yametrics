package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/artforteam2018/yametrics/internal/server/components/logger"
	"go.uber.org/zap"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timer := time.Now()

		next.ServeHTTP(w, r)

		logger.Log.Info("HTTP Request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("time", strconv.FormatInt(time.Since(timer).Milliseconds(), 10)),
		)
	})

}
