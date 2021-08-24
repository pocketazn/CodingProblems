package _2021_08

import "testing"

func TestF(t *testing.T) {
	t.Run("Test Waiting", func(t *testing.T) {
		Scheduler(F, 3)
	})
}