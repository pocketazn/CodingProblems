package _2021_08

import "testing"

func TestXORLinkedList_Add(t *testing.T) {
	type TestCase struct {
		Name string
	}

	Cases := []TestCase{
		{
			Name: "Testing",
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			myList := XORLinkedList{}
			myList.Add(1)
		})
	}
}
