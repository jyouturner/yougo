package yougo

import "fmt"

// go function to implement queue with linked list
type Node struct {
	data int
	next *Node
}

type Queue struct {
	size int
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(data int) {
	node := &Node{data: data}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		return -1
	}
	data := q.head.data
	q.head = q.head.next
	q.size--
	return data
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() int {
	if q.head == nil {
		return -1
	}
	return q.head.data
}

func (q *Queue) Print() {
	if q.head == nil {
		return
	}
	node := q.head
	for node != nil {
		fmt.Printf("%d ", node.data)
		node = node.next
	}
	fmt.Println()
}
