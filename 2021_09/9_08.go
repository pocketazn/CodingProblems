package _2021_09

// Given a singly linked list and an integer k, remove the kth last
// element from the list. k is guaranteed to be smaller than the
// length of the list.
//
// The list is very long, so making more than one pass is prohibitively
// expensive.
//
// Do this in constant space and in one pass.

type Node struct {
	Next *Node
	Key  int
}

type List struct {
	Head *Node
	Len  int
}

func (L *List) Insert(key int) {
	N := &Node{
		Key: key,
	}

	if L.Head != nil {
		N.Next = L.Head
		L.Head = N
	} else {
		L.Head = N
	}
	L.Len++
}

func (L *List) DeleteElementAtIndex(index int) {
	if L.Len < index || index < 0 {
		return
	}

	element := L.Head
	if index == 0 {
		L.Head = element.Next
		return
	}

	i := 0
	var previousElement *Node
	for i < index+1 {
		if i == index {
			previousElement.Next = element.Next
			L.Len--
		} else {
			previousElement = element
			element = element.Next
		}
		i++
	}
}
