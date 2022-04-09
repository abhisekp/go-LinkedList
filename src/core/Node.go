package core

import (
	"fmt"
	"strconv"
)

type Node struct {
	data     interface{}
	datatype string
	next     *Node
	// prev     *Node
}

func NewNode(data interface{}, datatype string, next *Node) *Node {
	return &Node{data, datatype, next}
}

func (n *Node) SetNext(node *Node) *Node {
	n.next = node
	return n
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetData(data interface{}, datatype string) *Node {
	n.data = data
	n.datatype = datatype

	if datatype == "" {
		n.datatype = "string"
	}

	return n
}

func (n *Node) Data() (interface{}, string) {
	return n.data, n.datatype
}

func (n *Node) String() string {
	switch n.datatype {
	case "string":
		dataStr := n.data.(string)
		return fmt.Sprintf("\"%s\"", dataStr)
	case "int":
		return strconv.Itoa(n.data.(int))
	case "float":
		return strconv.FormatFloat(n.data.(float64), 'f', -1, 64)
	case "bool":
		if n.data.(bool) {
			return "true"
		}
		return "false"
	case "uint":
		return strconv.Itoa(int(n.data.(uint)))
	case "uint8":
		return strconv.Itoa(int(n.data.(uint8)))
	case "uint16":
		return strconv.Itoa(int(n.data.(uint16)))
	case "uint32":
		return strconv.Itoa(int(n.data.(uint32)))
	case "uint64":
		return strconv.Itoa(int(n.data.(uint64)))
	case "int8":
		return strconv.Itoa(int(n.data.(int8)))
	case "int16":
		return strconv.Itoa(int(n.data.(int16)))
	case "int32":
		return strconv.Itoa(int(n.data.(int32)))
	case "int64":
		return strconv.Itoa(int(n.data.(int64)))
	case "float32":
		return strconv.FormatFloat(float64(n.data.(float32)), 'f', -1, 32)
	case "float64":
		return strconv.FormatFloat(n.data.(float64), 'f', -1, 64)
	case "rune":
		dataStr := string(n.data.(rune))
		return fmt.Sprintf("'%s'", dataStr)
	case "nil":
		return ""
	}

	return ""
}
