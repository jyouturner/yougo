package generics

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_Priority(t *testing.T) {
	pq := make(PriorityQueue[int], 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item[int]{value: 1, priority: 3})
	heap.Push(&pq, &Item[int]{value: 2, priority: 1})
	heap.Push(&pq, &Item[int]{value: 3, priority: 2})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[int])
		fmt.Println(item.value, item.priority)
	}
}
