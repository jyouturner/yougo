package generics

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	q := &Queue[string]{}
	q.Enqueue("hello")
	q.Enqueue("world")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
