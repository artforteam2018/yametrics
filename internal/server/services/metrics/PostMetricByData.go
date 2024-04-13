package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

type metricsPostRequest struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

type metricsPostResponse struct {
	ID    string  `json:"id"`    // имя метрики
	MType string  `json:"type"`  // параметр, принимающий значение gauge или counter
	Delta int64   `json:"delta"` // значение метрики в случае передачи counter
	Value float64 `json:"value"` // значение метрики в случае передачи gauge
}

func PostMetricByData(w http.ResponseWriter, r *http.Request) {

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

	responseData := metricsPostResponse{requestData.ID, requestData.MType, 0, 0}

	if requestData.MType == "counter" {
		if requestData.Delta == nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
		}
		responseData.Delta = metrics.Counter.Add(requestData.ID, *requestData.Delta)
	} else {
		if requestData.Value == nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
		}
		responseData.Value = metrics.Gauge.Add(requestData.ID, *requestData.Value)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(responseData)
}
