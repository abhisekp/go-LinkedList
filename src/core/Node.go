package core

import (
	"fmt"
	"strconv"
)

type INode[T NodeDatatype] interface {
	SetNext(node *Node[T]) *Node[T]
	Next() *Node[T]
	SetData(data T, datatype string) *Node[T]
	Data() (T, string)
	String() string
}

func NewNode[T NodeDatatype](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data: data, next: next}
}

type Node[T NodeDatatype] struct {
	fmt.Stringer
	data T
	next *Node[T]
	// prev     *Node
}

func (n *Node[T]) SetNext(node *Node[T]) *Node[T] {
	n.next = node
	return n
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) SetData(data T) *Node[T] {
	n.data = data

	return n
}

func (n *Node[T]) Data() *T {
	return &n.data
}

func (n *Node[T]) String() string {
	datatype := fmt.Sprintf("%T", n.data)
	data := fmt.Sprintf("%v", n.data)
	// fmt.Printf("{%v} = [%v]\n", datatype, data)
	if datatype == "string" {
		return fmt.Sprintf("%s", strconv.Quote(data))
	} else if datatype == "int32" {
		return strconv.QuoteRune(rune(fmt.Sprintf("%c", n.data)[0]))
	} else {
		return data
	}
}
