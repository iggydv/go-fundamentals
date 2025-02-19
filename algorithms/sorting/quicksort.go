package sorting

// Time: O(N log N) - average
// Time: O(N^2) - worst
// Space: O(log N) - average
// Space: O(N) - worst

// Quicksort is a divide-and-conquer algorithm that works by selecting a pivot element from the list and partitioning the
// other elements into two sub-arrays according to whether they are less than or greater than the pivot.
// The sub-arrays are then sorted recursively. This can be done in-place, requiring small additional amounts of memory to
// perform the sorting.

func Quicksort(arr []int, start int, end int) []int {
	if end <= start {
		return arr
	}
	pivot := partition(arr, start, end)
	Quicksort(arr, start, pivot-1)
	Quicksort(arr, pivot+1, end)
	return arr
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
