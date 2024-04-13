package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCounter(t *testing.T) {
	counter := Counter
	counter.Add("testm", 1)

	assert.Equal(t, int64(1), counter.data["testm"])
}
