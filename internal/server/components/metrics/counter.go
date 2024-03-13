package metrics

type Counter struct {
	data map[string][]int64
}

var counter = Counter{make(map[string][]int64)}

func (s *Counter) Add(metricName string, value int64) {
	if s.data[metricName] == nil {
		s.data[metricName] = []int64{}
	}
	s.data[metricName] = append(s.data[metricName], value)
}

func GetCounter() Counter {
	return counter
}
