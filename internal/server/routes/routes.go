package routes

import (
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/services/metrics"
)

func Init(mux *http.ServeMux) {

	mux.HandleFunc("/update/", metrics.HandleNoMetric)
	mux.HandleFunc("/update/gauge/", metrics.HandleGauge)
	mux.HandleFunc("/update/counter/", metrics.HandleCounter)

}
