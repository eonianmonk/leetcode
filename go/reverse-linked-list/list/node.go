package list

import "fmt"

type LNode[T any] struct {
	value T
	next  *LNode[T]
}

func newNode[T any](val T) *LNode[T] {
	return &LNode[T]{value: val, next: nil}
}

func (n *LNode[T]) addNode(val *LNode[T]) error {
	if n.next != nil {
		return fmt.Errorf("node already has next node")
	}
	n.next = val
	return nil
}
