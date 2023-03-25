package yougo

// go function to combine two sorted iterators
type Iterator struct {
	data []int
	pos  int
}

func (i *Iterator) HasNext() bool {
	return i.pos < len(i.data)
}

func (i *Iterator) Next() int {
	if !i.HasNext() {
		return -1
	}
	data := i.data[i.pos]
	i.pos++
	return data
}

func (i *Iterator) Peek() int {
	if !i.HasNext() {
		return -1
	}
	return i.data[i.pos]
}

func CombineIterators(a, b *Iterator) []int {
	var result []int
	for a.HasNext() && b.HasNext() {
		if a.Peek() == b.Peek() {
			result = append(result, a.Next())
			b.Next()
		} else if a.Peek() < b.Peek() {
			a.Next()
		} else {
			b.Next()
		}
	}
	return result
}

//The time complexity of the given function is O(m+n), where m and n are the lengths of the input iterators a and b, respectively.
//
//The function uses a two-pointer technique to compare the elements of the two input iterators a and b. It initializes two indices i and j to 0, which represent the current positions in iterators a and b, respectively. Then, it compares the values at those positions. If the values are equal, it adds the value to the result slice and increments both indices i and j. If the value at a[i] is less than the value at b[j], it increments i, otherwise, it increments j.
//
//In the worst case, each element of both iterators must be compared, so the number of iterations of the loop is at most m+n. Therefore, the time complexity of the function is O(m+n).
