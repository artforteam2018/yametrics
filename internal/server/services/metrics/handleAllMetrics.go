package metrics

import (
	"net/http"
	"strings"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

func HandleAllMetrics(w http.ResponseWriter, r *http.Request) {
	counter := metrics.GetCounter()
	counterKeys := counter.GetAll()

	gauge := metrics.GetGauge()
	gaugeKeys := gauge.GetAll()

	w.Write([]byte(strings.Join(append(counterKeys, gaugeKeys...), "\n")))
}
