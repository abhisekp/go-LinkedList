package core

import "github.com/abhisekp/go-linkedlist/src/internal"

type ILinkedListIterator interface {
	internal.Iterator
	Reset() bool
}

type LinkedListIterator struct {
	curr *Node
	list *LinkedList
}

func NewLinkedListIterator(list *LinkedList) ILinkedListIterator {
	return &LinkedListIterator{
		curr: list.head,
		list: list,
	}
}

func (llI *LinkedListIterator) HasNext() bool {
	return llI.curr != nil
}

func (llI *LinkedListIterator) Next() interface{} {
	if llI.curr == nil {
		return llI.curr
	}

	curr := llI.curr
	llI.curr = curr.Next()
	return curr
}

func (llI *LinkedListIterator) Reset() bool {
	llI.curr = llI.list.head
	return true
}
