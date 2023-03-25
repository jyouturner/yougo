package generics

type Queue[T any] []T

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
