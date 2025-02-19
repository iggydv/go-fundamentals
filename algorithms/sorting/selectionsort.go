package sorting

// Mostly sorted
// Time: O(N^2) - worst
// Space: O(1)

// Selection sort works by selecting the smallest element from the list and swapping it with the first element.
// The second-smallest element is selected and swapped with the second element, and so on until the list is sorted.
// It is called in-place sorting because it does not require extra space for sorting.

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
