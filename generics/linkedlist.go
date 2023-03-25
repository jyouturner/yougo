package generics

type ListNode[T comparable] struct {
	value T
	next  *ListNode[T]
}

type LinkedList[T comparable] struct {
	head *ListNode[T]
}

func (l *LinkedList[T]) Add(value T) {
	node := &ListNode[T]{value: value}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (l *LinkedList[T]) Remove(value T) bool {
	if l.head == nil {
		return false
	}
	if l.head.value == value {
		l.head = l.head.next
		return true
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList[T]) Iterator() func() *ListNode[T] {
	current := l.head
	return func() *ListNode[T] {
		if current == nil {
			return nil
		}
		node := current
		current = current.next
		return node
	}
}
