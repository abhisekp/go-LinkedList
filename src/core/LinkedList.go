package core

import (
	"strings"

	"github.com/abhisekp/go-linkedlist/src/internal"
)

type ILinkedList interface {
	LinkedListIterable
	Head() *Node
	Add(item interface{}, datatype string) *LinkedList
	Pop() (interface{}, string)
	Size() uint64
	String() string
}

type LinkedList struct {
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

func (ll *LinkedList) Iterator() ILinkedListIterator {
	return NewLinkedListIterator(ll)
}

func (ll *LinkedList) Iterator_() internal.Iterator {
	return ll.Iterator()
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}

func (ll *LinkedList) Add(data interface{}, datatype string) *LinkedList {
	node := NewNode(data, datatype, nil)

	nextNode := ll.Head()

	if nextNode == nil {
		ll.head = node
	} else {
		for nextNode.Next() != nil {
			nextNode = nextNode.Next()
		}
		nextNode.SetNext(node)
	}

	ll.size++

	return ll
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
		ll.head = nil
	}

	ll.size--

	return nextNode.Data()
}

func (ll *LinkedList) Size() uint64 {
	return ll.size
}

func (ll *LinkedList) String() string {

	var sb strings.Builder

	sb.WriteString("[")

	iterator := ll.Iterator()

	for iterator.HasNext() {
		item := iterator.Next()
		sb.WriteString(item.(*Node).String())
		if iterator.HasNext() {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")

	return sb.String()
}
