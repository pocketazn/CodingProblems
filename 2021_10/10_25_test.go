package _2021_10_test

import (
	_2021_10 "codingProblems/2021_10"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	sample := _2021_10.Node{}
	for i := 1; i < 5; i++ {
		sample.Insert(i)
	}

	fmt.Println(sample.GetArray())
	assert.NotNil(t, sample)
}

func TestReverse(t *testing.T) {
	sample := _2021_10.Node{}
	for i := 1; i < 5; i++ {
		sample.Insert(i)
	}
	reversed := sample.Reverse()
	fmt.Println(reversed.GetArray())
	assert.NotNil(t, sample)
}