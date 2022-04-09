package core

import (
	"bytes"
)

type ILinkedList interface {
	Add(item interface{}, datatype string) *LinkedList
	Pop() (interface{}, string)
	Size() uint64
	String() string
}

type LinkedList struct {
	// internal.Iterable
	head *Node
	// tail *Node
	size uint64
}

func NewLinkedList(head *Node) ILinkedList {
	ll := LinkedList{}
	if head == nil {
		return &ll
	}

	data, datatype := head.Data()

	return ll.Add(data, datatype)
}

func (ll *LinkedList) Add(data interface{}, datatype string) *LinkedList {
	node := NewNode(data, datatype, nil)

	nextNode := ll.Head()

	if nextNode == nil {
		ll.SetHead(node)
	} else {
		for nextNode.Next() != nil {
			nextNode = nextNode.Next()
		}
		nextNode.SetNext(node)
	}

	ll.size++

	return ll
}

func (ll *LinkedList) Size() uint64 {
	return ll.size
}

func (ll *LinkedList) SetHead(node *Node) *LinkedList {
	ll.head = node
	return ll
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}

func (ll *LinkedList) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	node := ll.Head()

	for node != nil {
		buffer.WriteString(node.String())
		node = node.Next()
		if node != nil {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}

func (ll *LinkedList) Pop() (interface{}, string) {
	nextNode := ll.Head()
	var prevNode *Node

	if nextNode == nil {
		return nil, "nil"
	}

	for nextNode.Next() != nil {
		prevNode = nextNode
		nextNode = nextNode.Next()
	}

	if prevNode != nil {
		prevNode.SetNext(nil)
	} else {
		ll.SetHead(nil)
	}

	ll.size--

	return nextNode.Data()
}

func (ll *LinkedList) Iterator() ILinkedListIterator {
	return NewLinkedListIterator(ll)
}
