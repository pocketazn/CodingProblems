package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarConTest(t *testing.T) {
	t.Run("Example Problem A", func(t *testing.T) {
		result := Car(Cons(3, 4))
		assert.Equal(t, 3, result)
	})
}

func TestCdrConTest(t *testing.T) {
	t.Run("Example Problem B", func(t *testing.T) {
		result := Cdr(Cons(3, 4))
		assert.Equal(t, 4, result)
	})
}
