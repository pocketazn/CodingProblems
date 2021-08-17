package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareSumToK(t *testing.T) {
	type TestCase struct {
		Name          string
		XValue        int
		YValue        int
		ExpectedValue bool
	}

	testCases := []TestCase{
		{
			Name:          "Is Equal",
			XValue:        10,
			YValue:        7,
			ExpectedValue: true,
		},
		{
			Name:          "Is Not Equal",
			XValue:        10,
			YValue:        8,
			ExpectedValue: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := CompareSumToK(tt.XValue, tt.YValue)
			assert.Equal(t, k, 17)
			assert.Equal(t, result, tt.ExpectedValue)
		})
	}

}

func TestDoesListContainSumPairEqualToK(t *testing.T) {
	type TestCase struct {
		Name          string
		List          []int
		ExpectedValue bool
	}

	testCases := []TestCase{
		{
			Name:          "Does Contain Pair",
			List:          []int{10, 15, 3, 7},
			ExpectedValue: true,
		},
		{
			Name:          "Last Result Pair",
			List:          []int{1, 2, 3, 10, 7},
			ExpectedValue: true,
		},
		{
			Name:          "Many Similar Result Pair",
			List:          []int{1, 1, 1, 1, 1, 1, 1, 10, 7},
			ExpectedValue: true,
		},
		{
			Name:          "Does Not Contain Pair",
			List:          []int{10, 2, 3},
			ExpectedValue: false,
		},
		{
			Name:          "Empty List",
			List:          []int{},
			ExpectedValue: false,
		},
		{
			Name:          "Only 1 value in List",
			List:          []int{1},
			ExpectedValue: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := DoesListContainSumPairEqualToK(tt.List)
			assert.Equal(t, k, 17)
			assert.Equal(t, result, tt.ExpectedValue)
		})
	}

}

func BenchmarkDoesListContainSumPairEqualToK(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoesListContainSumPairEqualToK([]int{1, 1, 1, 1, 1, 1, 1, 1, 10})
	}
}
