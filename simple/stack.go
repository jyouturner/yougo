package datastructures

type Stack []int

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
