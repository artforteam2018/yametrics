package metrics

import (
	"log"
	"net/http"
	"strconv"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
	"github.com/go-chi/chi/v5"
)

func CounterAdd(w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricname")
	metricValue := chi.URLParam(r, "metricvalue")
	if metricName == "" || metricValue == "" {
		http.NotFound(w, r)
		return
	}

	value, err := strconv.ParseInt(metricValue, 0, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	metrics.Counter.Add(metricName, value)
	log.Println(metricName, value)
}
