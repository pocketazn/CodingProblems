package _2021_10_test

import (
	_2021_10 "codingProblems/2021_10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPerfectNumber(t *testing.T) {
	type Case struct {
		Name string
		Given int
		Expected int
	}

	Cases := []Case{
		{
			Name: "1",
			Given: 1,
			Expected: 19,
		},
		{
			Name: "10",
			Given: 10,
			Expected: 109,
		},
		{
			Name: "2",
			Given: 2,
			Expected: 28,
		},
		{
			Name: "109",
			Given: 109,
			Expected: 109,
		},
		{
			Name: "1099999",
			Given: 1099999,
			Expected: -1,
		},
	}

	for _, tc := range Cases {
		t.Run(tc.Name, func(t *testing.T) {
			result := _2021_10.GetPerfectNumber(tc.Given)
			assert.Equal(t, tc.Expected, result)
		})
	}
}