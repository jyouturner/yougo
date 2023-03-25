package yougo

type ListNode struct {
	value int
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) Add(value int) {
	node := &ListNode{value: value}
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

func (l *LinkedList) Remove(value int) bool {
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

func (l *LinkedList) Iterator() func() *ListNode {
	current := l.head
	return func() *ListNode {
		if current == nil {
			return nil
		}
		node := current
		current = current.next
		return node
	}
}
