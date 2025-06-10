package ds

type DoubleLL[T comparable] interface {
	Head() *LinkedListNode[T]
	Tail() *LinkedListNode[T]
	AddNodeAtTail(T)
	AddNodeAtHead(T)
	AddNodeAfter(*LinkedListNode[T], T)
	RemoveNodeAtTail()
	MoveNodeToHead(*LinkedListNode[T])
	RemoveNode(*LinkedListNode[T])
}

type doubleLL[T comparable] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

func (d *doubleLL[T]) Head() *LinkedListNode[T] {
	return d.head
}

func (d *doubleLL[T]) Tail() *LinkedListNode[T] {
	return d.tail
}

func (d *doubleLL[T]) AddNodeAtTail(val T) {
	newNode := &LinkedListNode[T]{Value: val}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *doubleLL[T]) AddNodeAtHead(t T) {
	newNode := &LinkedListNode[T]{Value: t}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *doubleLL[T]) AddNodeAfter(node *LinkedListNode[T], val T) {
	newNode := &LinkedListNode[T]{Value: val}
	tempNode := node.next
	tempNode.prev = newNode
	newNode.next = tempNode
	node.next = newNode
}

func (d *doubleLL[T]) RemoveNodeAtTail() {
	d.tail = d.tail.prev
	d.tail.next = nil
}

func (d *doubleLL[T]) RemoveNode(node *LinkedListNode[T]) {
	tempNode := node.prev
	tempNode.next = node.next
	node.next.prev = tempNode
}

func (d *doubleLL[T]) MoveNodeToHead(node *LinkedListNode[T]) {
	d.RemoveNode(node)
	d.AddNodeAtHead(node.Value)
}

func NewDoubleLL[T comparable]() DoubleLL[T] {
	return &doubleLL[T]{}
}
