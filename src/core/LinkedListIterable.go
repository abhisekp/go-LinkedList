package core

import "github.com/abhisekp/go-linkedlist/src/internal"

type LinkedListIterable interface {
	internal.Iterable
	Iterator() ILinkedListIterator
}
