package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/metrics"
)

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

	var answer []byte
	var err error

	if requestData.MType == "counter" {
		if requestData.Delta == nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
		}
		responseData := metricsResponseCounter{requestData.ID, requestData.MType, 0}
		responseData.Delta = metrics.Counter.Add(requestData.ID, *requestData.Delta)

		answer, err = json.Marshal(responseData)

	} else {
		if requestData.Value == nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
		}
		responseData := metricsResponseGauge{requestData.ID, requestData.MType, 0}
		responseData.Value = metrics.Gauge.Add(requestData.ID, *requestData.Value)

		answer, err = json.Marshal(responseData)
	}

	if err != nil {
		http.Error(w, "Cannot send response "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)

}
