package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLowestPositiveInt(t *testing.T) {
	type TestCase struct {
		Name           string
		Sample         []int
		ExpectedResult int
	}

	Cases := []TestCase{
		{
			Name:           "Sample Size is 0",
			Sample:         []int{},
			ExpectedResult: 0,
		},
		{
			Name:           "Only Negative results",
			Sample:         []int{-1, -2},
			ExpectedResult: 0,
		},
		{
			Name:           "Only Negative Up to 0",
			Sample:         []int{-1, -2, 0},
			ExpectedResult: 1,
		},
		{
			Name:           "Negative to Positive Units",
			Sample:         []int{-1, 0, 2},
			ExpectedResult: 1,
		},
		{
			Name:           "Repeating Units",
			Sample:         []int{-1, 0, 1, 1, 1, 1, 2, 4},
			ExpectedResult: 3,
		},
		{
			Name:           "Problem Example A",
			Sample:         []int{3, 4, -1, 1},
			ExpectedResult: 2,
		},
		{
			Name:           "Problem Example B",
			Sample:         []int{1, 2, 0},
			ExpectedResult: 3,
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := FindLowestPositiveInt(tt.Sample)
			assert.Equal(t, tt.ExpectedResult, result)
		})
	}
}
