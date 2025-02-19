package sorting

// Time: O(n log n) - worst case
// Space: O(n)

// Merge sort is a divide-and-conquer algorithm that divides a list into equal halves until it has two single elements
// and then merges the sub-lists until the entire list has been reassembled in order.
// It is called out-of-place sorting because it requires extra space for sorting.
// It is a stable sort, meaning that it preserves the order of equal elements.

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	m := len(arr) / 2
	left := MergeSort(arr[:m])
	right := MergeSort(arr[m:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
