package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func TimerPrint(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timer := time.Now()
		next.ServeHTTP(w, r)

		fmt.Println("URL ", r.URL, " lasted -> ", time.Since(timer).Milliseconds(), " ms.")
	}
	return http.HandlerFunc(fn)
}
