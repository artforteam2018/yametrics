package metrics

import (
	"net/http"
)

func HandleNoMetric(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
