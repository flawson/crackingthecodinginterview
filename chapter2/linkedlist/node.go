package linkedlist

type SingleNode struct {
	Value int
	Next  *SingleNode
}

func NewSingleNode(value int) *SingleNode {
	return &SingleNode{
		Value: value,
	}
}

type DoubleNode struct {
	Value int
	Next  *DoubleNode
	Prev  *DoubleNode
}

func NewDoubleNode(value int) *DoubleNode {
	return &DoubleNode{
		Value: value,
	}
}
