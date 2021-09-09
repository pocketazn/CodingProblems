package _2021_09

type Node struct {
	Next *Node
	Key  int
}

type List struct {
	Head *Node
	Len int
}

func (L *List) Insert(key int) {
	N := &Node{
		Key:  key,
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
	element := L.Head
	if index == 0 {
		L.Head = element.Next
		return
	}

	i := 0
	var previousElement *Node
	for element.Next != nil  {
		if i == index {
			previousElement.Next = element.Next
		} else {
			previousElement = element
			element = element.Next
		}
		i++
	}
}