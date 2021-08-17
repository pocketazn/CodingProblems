package _2021_08

import (
	"fmt"
	"strings"
)

//Given the root to a binary tree, implement serialize(root), which
//serializes the tree into a string, and deserialize(s), which
//deserializes the string back into the tree.
//
//For example, given the following Node class
//
//class Node:
//    def __init__(self, val, left=None, right=None):
//        self.val = val
//        self.left = left
//        self.right = right
//The following test should pass:
//
//node = Node('root', Node('left', Node('left.left')), Node('right'))
//assert deserialize(serialize(node)).left.left.val == 'left.left'

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func (n *Node) Insert(data int) {
	if n == nil {
		return
	} else if data <= n.Data {
		// make left node or insert into left node
		if n.Left == nil {
			n.Left = &Node{Data: data}
		} else {
			n.Left.Insert(data)
		}
	} else {
		// make right node or insert into right node
		if n.Right == nil {
			n.Right = &Node{Data: data}
		} else {
			n.Right.Insert(data)
		}
	}
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) Insert(Value int) *BinaryTree {
	if b.root == nil {
		b.root = &Node{Data: Value}
	} else {
		b.root.Insert(Value)
	}
	return b
}

func Serialize(root *Node) string {
	var values []string
	values = append(values, fmt.Sprint(root.Data))
	if root.Left != nil {
		values = append(values, Serialize(root.Left))
	}
	if root.Right != nil {
		values = append(values, Serialize(root.Right))
	}
	return strings.Join(values, ",")
}

func Deserialize(treeString string) *Node {
	return &Node{}
}
