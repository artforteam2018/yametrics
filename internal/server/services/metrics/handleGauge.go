package metrics

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type MemStorage struct {
	gauge   map[string][]float64
	counter map[string][]int64
}

var memStorage = MemStorage{
	make(map[string][]float64),
	make(map[string][]int64)}

func (s *MemStorage) addGauge(metricName string, value float64) {
	if s.gauge[metricName] == nil {
		s.gauge[metricName] = []float64{}
	}
	s.gauge[metricName] = append(s.gauge[metricName], value)
}

func HandleGauge(w http.ResponseWriter, r *http.Request) {
	splittedValues := strings.Split(r.URL.Path, "/")
	if len(splittedValues) != 5 {
		http.NotFound(w, r)
		return
	}

	metricName, metricValue := splittedValues[3], splittedValues[4]

	value, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	memStorage.addGauge(metricName, value)
	fmt.Println(metricName, value)

}
