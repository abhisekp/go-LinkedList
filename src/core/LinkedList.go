package core

import (
	"fmt"
	"strings"

	"github.com/abhisekp/go-linkedlist/src/internal"
)

type ILinkedList[T NodeDatatype] interface {
	LinkedListIterable
	fmt.Stringer
	Head() *Node[T]
	Add(data T) *LinkedList[T]
	Pop() *T
	Size() uint64
	String() string
}

type LinkedList[T NodeDatatype] struct {
	head *Node[T]
	// tail *Node[T]
	size uint64
}

func NewLinkedList[T NodeDatatype]() ILinkedList[T] {
	ll := LinkedList[T]{}

	return &ll
}

func (ll *LinkedList[T]) Iterator() ILinkedListIterator {
	return NewLinkedListIterator(ll)
}

func (ll *LinkedList[T]) Iterator_() internal.Iterator {
	return ll.Iterator()
}

func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

func (ll *LinkedList[T]) Add(data T) *LinkedList[T] {
	NilNode := (*Node[T])(nil)

	node := NewNode[T](data, NilNode)

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

func (ll *LinkedList[T]) Pop() *T {
	NilNode := (*Node[T])(nil)

	nextNode := ll.Head()
	var prevNode *Node[T]

	if nextNode == nil {
		return nil
	}

	for nextNode.Next() != nil {
		prevNode = nextNode
		nextNode = nextNode.Next()
	}

	if prevNode != nil {
		prevNode.SetNext(NilNode)
	} else {
		ll.head = nil
	}

	ll.size--

	return nextNode.Data()
}

func (ll *LinkedList[T]) Size() uint64 {
	return ll.size
}

func (ll *LinkedList[T]) String() string {

	var sb strings.Builder

	sb.WriteString("[")

	iterator := ll.Iterator()

	for iterator.HasNext() {
		item := iterator.Next()
		sb.WriteString(item.(*Node[T]).String())
		if iterator.HasNext() {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")

	return sb.String()
}
