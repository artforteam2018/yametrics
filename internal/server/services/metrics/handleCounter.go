package metrics

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

func HandleCounter(w http.ResponseWriter, r *http.Request) {
	metricName, metricValue, err := ParseMetricNameValue(r.URL.Path)

	if err != nil {
		http.NotFound(w, r)
	}

	value, err := strconv.ParseInt(metricValue, 0, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	counter := metrics.GetCounter()

	counter.Add(metricName, value)
	fmt.Println(metricName, value)
}
