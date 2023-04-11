package datastructures

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	fmt.Println(q.Dequeue())
	q.Print()
	fmt.Println(q.Size())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())
}
