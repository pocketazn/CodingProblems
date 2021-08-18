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

func TestDeserialize(t *testing.T) {
	type TestCase struct {
		Name           string
		InsertList     []int
		SerializedTree string
	}

	Cases := []TestCase{
		{
			Name:           "Single Node",
			InsertList:     []int{4},
			SerializedTree: "4",
		},
		{
			Name:           "All Right Node",
			InsertList:     []int{1, 2, 3, 4, 5, 6},
			SerializedTree: "1,2,3,4,5,6",
		},
		{
			Name:           "All Left Node",
			InsertList:     []int{6, 5, 4, 3, 2, 1},
			SerializedTree: "6,5,4,3,2,1",
		},
		{
			Name:           "Mixed Left and Right Node",
			InsertList:     []int{6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 20},
			SerializedTree: "6,5,4,3,2,1,2,3,4,5,20",
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			ExpectedTree := BinaryTree{}
			for _, data := range tt.InsertList {
				ExpectedTree.Insert(data)
			}
			result := Deserialize(tt.SerializedTree)
			assert.Equal(t, ExpectedTree, result)
		})
	}
}

func TestSerializeAndDeserialize(t *testing.T) {
	type TestCase struct {
		Name       string
		InsertList []int
	}

	Cases := []TestCase{
		{
			Name:       "Single Node",
			InsertList: []int{4},
		},
		{
			Name:       "Mixed Value Tree",
			InsertList: []int{4, 6, 7, 8, 10, 1, 3, 5, 7, 2, 3, 8, 9, 22, 45, 100, 56, 77},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			ExpectedTree := BinaryTree{}
			for _, data := range tt.InsertList {
				ExpectedTree.Insert(data)
			}

			serialized := Serialize(ExpectedTree.root)
			ResultTree := Deserialize(serialized)

			assert.Equal(t, ExpectedTree, ResultTree)
		})
	}
}
