package metrics

import "strconv"

type Counter struct {
	data map[string]int64
}

var counter = Counter{make(map[string]int64)}

func (s *Counter) Add(metricName string, value int64) {
	s.data[metricName] += value
}

func (s *Counter) Get(metricName string) (int64, bool) {
	val, ok := s.data[metricName]

	return val, ok
}

func (s *Counter) GetAll() []string {
	var response []string
	for k, v := range s.data {
		response = append(response, k+": "+strconv.FormatInt(v, 10))
	}
	return response
}

func GetCounter() Counter {
	return counter
}
