package main

import (
	"github.com/abhisekp/go-linkedlist/src/core"
)

type N interface {
	~int32
}

func main() {
	ll := core.NewLinkedList[any]().
		Add(1).
		Add(2).
		Add(3).
		Add(float32(7)).
		Add(5)

	ll.Pop()

	iterator := ll.Iterator()

	for iterator.HasNext() {
		item := iterator.Next()
		println(item.(*core.Node[any]).String())
	}
	iterator.Reset()
	println(ll.String())

	ll.Pop()

	for iterator.HasNext() {
		item := iterator.Next()
		println(item.(*core.Node[any]).String())
	}
	iterator.Reset()

	println(ll.String())

	ll.Add('A')
	ll.Add(int32(65))
	ll.Add(65)
	ll.Add("Hello World")
	ll.Add("Abhisek")

	iterator2 := core.NewLinkedListIterator(ll)
	for iterator2.HasNext() {
		item := iterator2.Next()
		println(item.(*core.Node[any]).String())
	}
	iterator2.Reset()

	println(ll.String())
}
