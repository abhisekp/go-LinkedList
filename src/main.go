package main

import (
	"github.com/abhisekp/go-linkedlist/src/core"
)

func main() {
	ll := core.NewLinkedList(nil).
		Add(1, "int").
		Add(2, "int").
		Add(3, "int").
		Add(float32(7), "float32").
		Add(5, "int")

	ll.Pop()

	iterator := ll.Iterator()

	for iterator.HasNext() {
		item := iterator.Next()
		println(item.(*core.Node).String())
	}
	iterator.Reset()
	println(ll.String())

	ll.Pop()

	for iterator.HasNext() {
		item := iterator.Next()
		println(item.(*core.Node).String())
	}
	iterator.Reset()

	println(ll.String())

	ll.Add('A', "rune")
	ll.Add("Hello World", "string")
	ll.Add("Abhisek", "string")

	iterator2 := core.NewLinkedListIterator(ll)
	for iterator2.HasNext() {
		item := iterator2.Next()
		println(item.(*core.Node).String())
	}
	iterator2.Reset()

	println(ll.String())
}
