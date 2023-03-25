package yougo

// go function to find duplicate between two sorted lists
func FindDuplicates(a, b []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

//The time complexity of the given function is O(m+n), where m and n are the lengths of the input slices a and b, respectively.
//
//The function uses a two-pointer technique to compare the elements of the two input slices a and b. It initializes two indices i and j to 0, which represent the current positions in slices a and b, respectively. Then, it compares the values at those positions. If the values are equal, it adds the value to the result slice and increments both indices i and j. If the value at a[i] is less than the value at b[j], it increments i, otherwise, it increments j.
//
//In the worst case, each element of both slices must be compared, so the number of iterations of the loop is at most m+n. Therefore, the time complexity of the function is O(m+n).
