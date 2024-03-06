package linkedlist

type SingleLinkedList struct {
	Head *SingleNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Insert(node *SingleNode) *SingleLinkedList {
	if l.Head == nil {
		l.Head = node
		return l
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	return l
}

type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func NewDoubleleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) Insert(node *DoubleNode) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = l.Tail.Next
}
