package metrics

type Gauge struct {
	data map[string]float64
}

func (s *Gauge) Add(metricName string, value float64) {

	s.data[metricName] = value
}
