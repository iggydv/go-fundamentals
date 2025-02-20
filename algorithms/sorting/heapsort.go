package sorting

// Time: O(N log N) - worst
// Time: O(N log N) - average
// Space: O(1)

func HeapSort(arr []int) []int {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyDown(arr, 0, i)
	}
	return arr
}

func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapifyDown(arr, i, len(arr))
	}
}

func heapifyDown(arr []int, i int, length int) {
	l := left(i)
	r := right(i)
	largest := i

	if l < length && arr[l] > arr[largest] {
		largest = l
	}

	if r < length && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyDown(arr, largest, length)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
