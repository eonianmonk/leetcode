package list

import "fmt"

type LList[T any] struct {
	head *LNode[T]
	tail *LNode[T]
}

func NewEmptyLList[T any]() *LList[T] {
	return &LList[T]{}
}

func NewLList[T any](n *LNode[T]) (*LList[T], error) {
	if n == nil {
		return nil, fmt.Errorf("node should not be nil")
	}
	list := LList[T]{head: n}
	ntry := n
	for ntry.next != nil {
		ntry = ntry.next
	}
	list.tail = ntry
	return &list, nil
}

func (l *LList[T]) Add(val T) error {
	if l.head == nil {
		node := newNode[T](val)
		l.head = node
		l.tail = node
		return nil
	}
	err := l.tail.addNode(newNode[T](val))
	if err != nil {
		return fmt.Errorf("failed to add to tail node: %s", err.Error())
	}
	l.tail = l.tail.next
	return nil
}

func (l *LList[T]) Print(separator string) {
	ln := l.head
	for {
		fmt.Printf("%v%s", ln.value, separator)
		if ln == l.tail {
			break
		}
		ln = ln.next
	}
}

func ReverseLList[T any](src *LList[T]) (res *LList[T]) {
	if src == nil {
		return
	}
	res = reverseLList[T](src)
	return
}

// func inverseLListWRecursion[T any](head *LNode[T], l2 *LList[T]) {
// 	if head.next != nil {
// 		inverseLListWRecursion[T](head.next, l2)
// 	}
// 	l2.Add(head.value)
// }

func reverseLList[T any](l1 *LList[T]) (l2 *LList[T]) {
	var cur, prev *LNode[T]
	iter := l1.head
	prev = nil
	for iter != nil {
		cur = newNode[T](iter.value)
		cur.next = prev
		prev = cur
		iter = iter.next
	}
	l2, _ = NewLList[T](cur)
	return

}
