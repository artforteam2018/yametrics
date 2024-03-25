package metrics

import (
	"net/http"
	"strconv"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
	"github.com/go-chi/chi/v5"
)

func HandleCounterGet(w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricname")
	if metricName == "" {
		http.NotFound(w, r)
		return
	}

	counter := metrics.GetCounter()

	value, ok := counter.Get(metricName)

	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(strconv.FormatInt(value, 10)))
}
