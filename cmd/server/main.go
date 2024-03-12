package main

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

func (s *MemStorage) AddGauge(metricName string, value float64) {
	if s.gauge[metricName] == nil {
		s.gauge[metricName] = []float64{}
	}
	s.gauge[metricName] = append(s.gauge[metricName], value)
}

func (s *MemStorage) AddCounter(metricName string, value int64) {
	if s.counter[metricName] == nil {
		s.counter[metricName] = []int64{}
	}
	s.counter[metricName] = append(s.counter[metricName], value)
}

func handleGauge(w http.ResponseWriter, r *http.Request) {
	splittedValues := strings.Split(r.URL.Path, "/")
	if len(splittedValues) != 5 {
		http.NotFound(w, r)
		return
	}

	metricType, metricName, metricValue := splittedValues[2], splittedValues[3], splittedValues[4]

	if metricType == "gauge" {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		memStorage.AddGauge(metricName, value)
		fmt.Println(metricName, value)

	} else if metricType == "counter" {
		value, err := strconv.ParseInt(metricValue, 0, 64)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		memStorage.AddCounter(metricName, value)
		fmt.Println(metricName, value)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handleGauge)

	fmt.Println("server is listening on port 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
