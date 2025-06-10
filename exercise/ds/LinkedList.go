package ds

type LinkedListNode[T comparable] struct {
	Value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

type LinkedList[T comparable] interface {
	AddNodeAtTail(T)
	AddNodeAfter(*LinkedListNode[T], T)
	Head() *LinkedListNode[T]
}

type linkedList[T comparable] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

func (l *linkedList[T]) AddNodeAtTail(value T) {
	if l.tail == nil {
		l.head = &LinkedListNode[T]{Value: value}
		l.tail = l.head
	} else {
		l.tail.next = &LinkedListNode[T]{Value: value}
		l.tail = l.tail.next
	}
}

func (l *linkedList[T]) Head() *LinkedListNode[T] {
	return l.head
}

func (l *linkedList[T]) AddNodeAfter(node *LinkedListNode[T], value T) {
	newNode := &LinkedListNode[T]{Value: value}
	tempNode := node.next
	newNode.next = tempNode
	node.next = newNode
}

func NewLinkedList[T comparable]() LinkedList[T] {
	return &linkedList[T]{}
}
