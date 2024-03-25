package interval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartInterval(t *testing.T) {
	timeDuration := time.Millisecond
	checker := false
	stubFunc := func() {
		checker = true
	}

	StartInterval(timeDuration, stubFunc)

	assert.Equal(t, false, checker)

	time.Sleep(timeDuration * 2)

	assert.Equal(t, true, checker)
}

func TestStartDisableInterval(t *testing.T) {
	timeDuration := time.Millisecond
	checker := false
	stubFunc := func() {
		checker = true
	}

	disabler := StartInterval(timeDuration, stubFunc)

	disabler <- true

	time.Sleep(timeDuration)

	assert.Equal(t, false, checker)
}
