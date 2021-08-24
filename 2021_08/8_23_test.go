package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestF(t *testing.T) {
	t.Run("Test Scheduler", func(t *testing.T) {
		start := time.Now()
		Scheduler(F, 5000)
		elapsed := time.Since(start)
		assert.GreaterOrEqual(t, elapsed, time.Duration(5000)*time.Millisecond)
	})
}