package _2021_08

import (
	"fmt"
	"unsafe"
)

//An XOR linked list is a more memory efficient doubly linked list.
//Instead of each node holding next and prev fields, it holds a field named both,
//which is an XOR of the next node and the previous node. Implement an XOR linked list;
//it has an add(element) which adds the element to the end, and a get(index)
//which returns the node at index.
//
//If using a language that has no pointers (such as Python), you can
//assume you have access to get_pointer and dereference_pointer
//functions that converts between nodes and memory addresses.

// incomplete due to not figuriong out how to get binaries to address

type XORNode struct {
	Both *XORNode
	Value interface{}
}

type XORLinkedList struct {
	First *XORNode
	Last *XORNode
}

func (r XORLinkedList) Add(val interface{}) {
	NewNode := XORNode{Value: val}
	if r.First != nil {
		r.First = &NewNode
	}
	if r.Last != nil {
		r.Last = &NewNode
	}

	startAddress := uintptr(unsafe.Pointer(&NewNode))
	fmt.Printf("Start Address of s: %d\n", startAddress)
}