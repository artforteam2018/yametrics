package metrics

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

func HandleGauge(w http.ResponseWriter, r *http.Request) {
	splittedValues := strings.Split(r.URL.Path, "/")
	if len(splittedValues) != 5 {
		http.NotFound(w, r)
		return
	}

	metricName, metricValue := splittedValues[3], splittedValues[4]

	value, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	gauge := metrics.GetGauge()

	gauge.Add(metricName, value)

	fmt.Println(metricName, value)
}
