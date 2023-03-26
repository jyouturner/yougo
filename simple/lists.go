package yougo

import (
	"container/heap"
	"fmt"
	"sort"
)

// MergeTwoLists merge the two sorted integer lists
func MergeTwoLists(list1 []int, list2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] == list2[j] {
			res = append(res, list1[i])
			i++
			j++
		} else if list1[i] < list2[j] {
			res = append(res, list1[i])
			i++
		} else {
			res = append(res, list2[j])
			j++
		}
	}
	res = append(res, list1[i:]...)
	res = append(res, list2[j:]...)
	return res
}

// Merge and Reverse Sort two integer lists
func DescSortToOne(a []int, b []int) []int {
	merged := append(a, b...)
	// quick sort the integers in increasing order
	sort.Ints(merged)
	// reverse
	i, j := 0, len(merged)-1
	for i < j {
		merged[i], merged[j] = merged[j], merged[i]
		i++
		j--
	}
	return merged
}

// use priority queue to implement a task list
type Task struct {
	p int
	k string
	v []byte
}

type PriorityTasks []Task

func (pt PriorityTasks) Len() int {
	return len(pt)
}

func (pt PriorityTasks) Less(i int, j int) bool {
	return pt[i].p < pt[j].p
}

func (pt PriorityTasks) Swap(i int, j int) {
	pt[i], pt[j] = pt[j], pt[i]
}

func (pt *PriorityTasks) Push(t any) {
	*pt = append(*pt, t.(Task))
}

func (pt *PriorityTasks) Pop() any {
	res := &((*pt)[len(*pt)-1])
	*pt = (*pt)[0 : len(*pt)-1]
	return *res
}

func TestTaskQueue() {
	l1 := []int{1, 5, 7, 9, 10, 11, 12}
	l2 := []int{4, 6, 7, 8, 10, 14}
	//res := DescSortToOne(l1, l2)
	//fmt.Println(res)
	pt := make(PriorityTasks, len(l1))

	for i, x := range l1 {
		task := Task{
			p: x,
			k: fmt.Sprintf("task%v", i),
			v: []byte(""),
		}
		pt[i] = task
	}

	heap.Init(&pt)
	for j, x := range l2 {
		task := Task{
			p: x,
			k: fmt.Sprintf("task%v", j),
			v: []byte(""),
		}
		heap.Push(&pt, task)
	}

	fmt.Println(heap.Pop(&pt))
	fmt.Println(heap.Pop(&pt))
	fmt.Println(heap.Pop(&pt))
	fmt.Println(heap.Pop(&pt))
}
