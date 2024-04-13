package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGauge(t *testing.T) {
	gauge := Gauge
	gauge.Add("testm", 1)

	assert.Equal(t, float64(1), gauge.data["testm"])
}
