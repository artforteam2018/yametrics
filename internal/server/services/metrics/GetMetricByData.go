package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

type metricsGetRequest struct {
	ID    string `json:"id"`   // имя метрики
	MType string `json:"type"` // параметр, принимающий значение gauge или counter
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

	var answer []byte
	var err error

	if requestData.MType == "counter" {
		responseData := metricsResponseCounter{requestData.ID, requestData.MType, 0}
		responseData.Delta, _ = metrics.Counter.Get(requestData.ID)
		answer, err = json.Marshal(responseData)
	} else {
		responseData := metricsResponseGauge{requestData.ID, requestData.MType, 0}
		responseData.Value, _ = metrics.Gauge.Get(requestData.ID)
		answer, err = json.Marshal(responseData)
	}

	if err != nil {
		http.Error(w, "Cannot send response "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}
