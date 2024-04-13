package metrics

import (
	"net/http"
)

func NoMetric(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
