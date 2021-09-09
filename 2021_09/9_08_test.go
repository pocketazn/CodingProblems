package _2021_09

import (
	"testing"
)

func TestList_DeleteElement(t *testing.T) {
	t.Run("happyPath", func(t *testing.T) {
		SampleLinkedList := List{}
		for i := 1; i < 5; i++ {
			SampleLinkedList.Insert(i)
		}

		SampleLinkedList.DeleteElementAtIndex(2)
	})
}