package internal

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
