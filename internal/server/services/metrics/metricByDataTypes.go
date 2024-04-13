package metrics

type metricsPostRequest struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

type metricsResponseCounter struct {
	ID    string `json:"id"`    // имя метрики
	MType string `json:"type"`  // параметр, принимающий значение gauge или counter
	Delta int64  `json:"delta"` // значение метрики в случае передачи counter
}

type metricsResponseGauge struct {
	ID    string  `json:"id"`    // имя метрики
	MType string  `json:"type"`  // параметр, принимающий значение gauge или counter
	Value float64 `json:"value"` // значение метрики в случае передачи gauge
}
