package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

type metricsPostRequest struct {
	ID    string  `json:"id"`              // имя метрики
	MType string  `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func PostMetricJSON(w http.ResponseWriter, r *http.Request) {

	var requestData metricsPostRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.ID == "" {
		http.Error(w, "Invalid ID "+requestData.ID, http.StatusBadRequest)
		return
	}

	if requestData.MType != "counter" && requestData.MType != "gauge" {
		http.Error(w, "Invalid type: "+requestData.MType, http.StatusBadRequest)
		return
	}

	if requestData.MType == "counter" {
		requestData.Delta = metrics.Counter.Add(requestData.ID, requestData.Delta)
	} else {
		requestData.Value = metrics.Gauge.Add(requestData.ID, requestData.Value)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(requestData)
}
