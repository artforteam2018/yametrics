package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGauge(t *testing.T) {
	counter := GetGauge()
	counter.Add("testm", 1)

	assert.Equal(t, float64(1), counter.data["testm"])
}
