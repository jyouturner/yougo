package datastructures

import (
	fmt "fmt"
	testing "testing"
)

func TestMain(m *testing.M) {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.Iterator()
	for node := iterator(); node != nil; node = iterator() {
		fmt.Println(node.value)
	}

	list.Remove(2)

	iterator = list.Iterator()
	for node := iterator(); node != nil; node = iterator() {
		fmt.Println(node.value)
	}
}
