package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductsOfSliceWithDivision(t *testing.T) {
	type TestCase struct {
		Name           string
		Input          []int
		ExpectedResult []int
	}

	Cases := []TestCase{
		{
			Name:           "Single Element",
			Input:          []int{3},
			ExpectedResult: []int{},
		},
		{
			Name:           "No Elements",
			Input:          []int{},
			ExpectedResult: []int{},
		},
		{
			Name:           "2 Elements",
			Input:          []int{1, 2},
			ExpectedResult: []int{2, 1},
		},
		{
			Name:           "3 Elements",
			Input:          []int{3, 2, 1},
			ExpectedResult: []int{2, 3, 6},
		},
		{
			Name:           "Example from Problem",
			Input:          []int{1, 2, 3, 4, 5},
			ExpectedResult: []int{120, 60, 40, 30, 24},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := GetProductOfSliceWithDivision(tt.Input)
			assert.Equal(t, tt.ExpectedResult, result)
		})
	}
}

func TestGetProductOfAllElements(t *testing.T) {
	type TestCase struct {
		Name           string
		Input          []int
		ExpectedResult int
	}

	Cases := []TestCase{
		{
			Name:           "No Element",
			Input:          []int{},
			ExpectedResult: -1,
		},
		{
			Name:           "One Element",
			Input:          []int{2},
			ExpectedResult: 2,
		},
		{
			Name:           "3 Element",
			Input:          []int{1, 2, 3},
			ExpectedResult: 6,
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := GetProductOfAllElements(tt.Input)
			assert.Equal(t, tt.ExpectedResult, result)
		})
	}
}

func TestGetProductsOfSliceWithoutDivision(t *testing.T) {
	type TestCase struct {
		Name           string
		Input          []int
		ExpectedResult []int
	}

	Cases := []TestCase{
		{
			Name:           "Single Element",
			Input:          []int{3},
			ExpectedResult: []int{},
		},
		{
			Name:           "No Elements",
			Input:          []int{},
			ExpectedResult: []int{},
		},
		{
			Name:           "2 Elements",
			Input:          []int{1, 2},
			ExpectedResult: []int{2, 1},
		},
		{
			Name:           "3 Elements",
			Input:          []int{3, 2, 1},
			ExpectedResult: []int{2, 3, 6},
		},
		{
			Name:           "Example from Problem",
			Input:          []int{1, 2, 3, 4, 5},
			ExpectedResult: []int{120, 60, 40, 30, 24},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := GetProductOfSliceWithOutDivision(tt.Input)
			assert.Equal(t, tt.ExpectedResult, result)
		})
	}
}

func BenchmarkGetProductOfSliceWithDivision(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetProductOfSliceWithDivision([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

// averages about 20% faster than division
// TIL: Multiplication is way less intensive than division
func BenchmarkGetProductOfSliceWithoutDivision(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetProductOfSliceWithOutDivision([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
