package core

type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
	~float32 | ~float64
}

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Integer interface {
	Int | UInt
}

type Number interface {
	Integer | Float | Complex
}

type Boolean interface {
	~bool
}

type String interface {
	~string | ~[]byte | ~[]rune
}

type NodeDatatype interface {
	String | Boolean | Number | ~rune | any
}
