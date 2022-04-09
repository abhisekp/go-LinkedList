package core

import "github.com/abhisekp/go-linkedlist/src/internal"

type ILinkedListIterator interface {
	internal.Iterator
	Reset() bool
}

type LinkedListIterator[T NodeDatatype] struct {
	curr *Node[T]
	list *LinkedList[T]
}

func NewLinkedListIterator[T NodeDatatype](list *LinkedList[T]) ILinkedListIterator {
	return &LinkedListIterator[T]{
		curr: list.head,
		list: list,
	}
}

func (llI *LinkedListIterator[T]) HasNext() bool {
	return llI.curr != nil
}

func (llI *LinkedListIterator[T]) Next() any {
	if llI.curr == nil {
		return llI.curr
	}

	curr := llI.curr
	llI.curr = curr.Next()
	return curr
}

func (llI *LinkedListIterator[T]) Reset() bool {
	llI.curr = llI.list.head
	return true
}
