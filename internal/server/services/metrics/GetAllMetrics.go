package metrics

import (
	"net/http"
	"strings"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

func GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	counterKeys := metrics.Counter.GetAll()
	gaugeKeys := metrics.Gauge.GetAll()

	w.Write([]byte(strings.Join(append(counterKeys, gaugeKeys...), "\n")))
}
