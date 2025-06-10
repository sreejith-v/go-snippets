package ds

import (
	"testing"
)

func TestLinkedList_AddNodeAtTail(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.AddNodeAtTail(10)
	ll.AddNodeAtTail(20)

	head := ll.Head()
	if head == nil || head.value != 10 {
		t.Fatalf("expected head value to be 10, got %v", head)
	}

	second := head.next
	if second == nil || second.value != 20 {
		t.Fatalf("expected second node value to be 20, got %v", second)
	}

	if second.next != nil {
		t.Fatalf("expected second node to be the tail, but found a next node")
	}
}

func TestLinkedList_AddNodeAfter(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.AddNodeAtTail(10)
	head := ll.Head()

	ll.AddNodeAfter(head, 15)

	if head.next == nil || head.next.value != 15 {
		t.Fatalf("expected next node value after head to be 15, got %v", head.next)
	}

	second := head.next
	ll.AddNodeAfter(second, 20)

	if second.next == nil || second.next.value != 20 {
		t.Fatalf("expected next node value after second to be 20, got %v", second.next)
	}
}
