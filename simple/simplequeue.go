package yougo

type IntQueue []int

func (q *IntQueue) Enqueue(value int) {
	*q = append(*q, value)
}

func (q *IntQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *IntQueue) IsEmpty() bool {
	return len(*q) == 0
}
