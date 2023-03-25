package yougo

import (
	"fmt"
	"testing"
)

func TestSortedMap(t *testing.T) {
	m := &SortedMap{data: make(map[string]int)}
	m.Set("apple", 1)
	m.Set("banana", 2)
	m.Set("orange", 3)

	fmt.Println(m.Get("apple"))
	fmt.Println(m.Get("banana"))
	fmt.Println(m.Get("orange"))
}
