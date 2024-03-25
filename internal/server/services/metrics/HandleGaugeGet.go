package metrics

import (
	"net/http"
	"strconv"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
	"github.com/go-chi/chi/v5"
)

func HandleGaugeGet(w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricname")
	if metricName == "" {
		http.NotFound(w, r)
		return
	}

	gauge := metrics.GetGauge()

	value, ok := gauge.Get(metricName)

	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(strconv.FormatFloat(value, 'f', 3, 64)))
}
