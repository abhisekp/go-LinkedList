package core

type LinkedListIterable interface {
	// internal.Iterable
	Iterator() ILinkedListIterator
}
