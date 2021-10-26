package _2021_10

type Node struct {
	Next *Node
	Key  int
}

func (n *Node) Insert(k int) {
	if n.Next != nil {
		n.Next.Insert(k)
	} else {
		n.Next = &Node{Key: k}
	}
}

func (n *Node) GetArray() []int {
	var list []int
	position := n
	for true {
		list = append(list, position.Key)
		if position.Next != nil {
			position = position.Next
		} else {
			break
		}
	}

	return list
}

func (n *Node) Reverse() *Node {
	current := n
	var prev *Node
	prev = nil

	for current != nil {
		current, prev, current.Next = current.Next, current, prev
	}
	return prev
}
