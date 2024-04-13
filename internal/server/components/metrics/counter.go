package metrics

import "strconv"

type CounterT struct {
	data map[string]int64
}

var Counter = CounterT{make(map[string]int64)}

func (s *CounterT) Add(metricName string, value int64) int64 {
	s.data[metricName] += value
	return s.data[metricName]
}

func (s *CounterT) Get(metricName string) (int64, bool) {
	val, ok := s.data[metricName]

	return val, ok
}

func (s *CounterT) GetAll() []string {
	var response []string
	for k, v := range s.data {
		response = append(response, k+": "+strconv.FormatInt(v, 10))
	}
	return response
}
