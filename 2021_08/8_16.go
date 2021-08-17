package _2021_08

import (
	"fmt"
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
	Left *Node
	Right *Node
	Data string
}

func Serialize(tree *Node) string {
	left := ""
	right := ""
	if tree.Left != nil {
		left = Serialize(tree.Left)
	}

	if tree.Right != nil {
		right = Serialize(tree.Right)
	}

	return fmt.Sprintf("[%s.%s.%s]", tree.Data, left, right)
}

func Deserialize(treeString string) Node {
	return Node{}
}
