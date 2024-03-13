package metrics

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (s *MemStorage) addCounter(metricName string, value int64) {
	if s.counter[metricName] == nil {
		s.counter[metricName] = []int64{}
	}
	s.counter[metricName] = append(s.counter[metricName], value)
}

func HandleCounter(w http.ResponseWriter, r *http.Request) {
	splittedValues := strings.Split(r.URL.Path, "/")
	if len(splittedValues) != 5 {
		http.NotFound(w, r)
		return
	}

	metricName, metricValue := splittedValues[3], splittedValues[4]

	value, err := strconv.ParseInt(metricValue, 0, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	memStorage.addCounter(metricName, value)
	fmt.Println(metricName, value)
}
