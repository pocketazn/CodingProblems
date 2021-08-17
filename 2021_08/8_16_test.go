package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	type TestCase struct {
		Name                  string
		InsertList            []int
		ExpectedSerialization string
	}

	Cases := []TestCase{
		{
			Name:                  "Single Node",
			InsertList:            []int{4},
			ExpectedSerialization: "4",
		},
		{
			Name:                  "All Right Node",
			InsertList:            []int{1, 2, 3, 4, 5, 6},
			ExpectedSerialization: "1,2,3,4,5,6",
		},
		{
			Name:                  "All Left Node",
			InsertList:            []int{6, 5, 4, 3, 2, 1},
			ExpectedSerialization: "6,5,4,3,2,1",
		},
		{
			Name:                  "Mixed Left and Right Node",
			InsertList:            []int{6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 20},
			ExpectedSerialization: "6,5,4,3,2,1,2,3,4,5,20",
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			initTree := BinaryTree{}
			for _, data := range tt.InsertList {
				initTree.Insert(data)
			}
			result := Serialize(initTree.root)
			assert.Equal(t, tt.ExpectedSerialization, result)
		})
	}
}
