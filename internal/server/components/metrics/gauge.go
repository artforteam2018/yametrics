package metrics

import "strconv"

type GaugeT struct {
	data map[string]float64
}

var Gauge = GaugeT{make(map[string]float64)}

func (s *GaugeT) Add(metricName string, value float64) float64 {
	s.data[metricName] = value
	return s.data[metricName]
}

func (s *GaugeT) Get(metricName string) (float64, bool) {
	val, ok := s.data[metricName]

	return val, ok
}

func (s *GaugeT) GetAll() []string {
	var response []string
	for k, v := range s.data {
		response = append(response, k+": "+strconv.FormatFloat(v, 'f', -1, 64))

	}
	return response
}
