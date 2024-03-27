package metrics

import "strconv"

type Gauge struct {
	data map[string]float64
}

var gauge = Gauge{make(map[string]float64)}

func (s *Gauge) Add(metricName string, value float64) {

	s.data[metricName] = value
}

func (s *Gauge) Get(metricName string) (float64, bool) {
	val, ok := s.data[metricName]

	return val, ok
}

func GetGauge() Gauge {
	return gauge
}

func (s *Gauge) GetAll() []string {
	var response []string
	for k, v := range s.data {
		response = append(response, k+": "+strconv.FormatFloat(v, 'f', -1, 64))

	}
	return response
}
