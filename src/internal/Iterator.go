package internal

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Iterable interface {
	Iterator() *Iterator
}
