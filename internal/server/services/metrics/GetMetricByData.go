package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

type metricsGetRequest struct {
	ID    string `json:"id"`   // имя метрики
	MType string `json:"type"` // параметр, принимающий значение gauge или counter
}

type metricsResponse struct {
	ID    string  `json:"id"`              // имя метрики
	MType string  `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func GetMetricByData(w http.ResponseWriter, r *http.Request) {

	var requestData metricsGetRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.ID == "" {
		http.Error(w, "Invalid ID "+requestData.ID, http.StatusBadRequest)
		return
	}

	if requestData.MType != "counter" && requestData.MType != "gauge" {
		http.Error(w, "Invalid type "+requestData.MType, http.StatusBadRequest)
		return
	}

	responseData := metricsResponse{requestData.ID, requestData.MType, 0, 0}

	if requestData.MType == "counter" {
		responseData.Delta, _ = metrics.Counter.Get(requestData.ID)
	} else {
		responseData.Value, _ = metrics.Gauge.Get(requestData.ID)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(responseData)

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Cannot send response "+err.Error(), http.StatusInternalServerError)
		return
	}
}
