package metrics

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

func HandleGauge(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	metricName, metricValue, err := ParseMetricNameValue(r.URL.Path)

	if err != nil {
		http.NotFound(w, r)
	}

	value, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	gauge := metrics.GetGauge()

	gauge.Add(metricName, value)

	fmt.Println(metricName, value)
}
