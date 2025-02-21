package search

// Assumes sorted array

// Time: O(log n)
// Space: O(1)

// Binary search works by comparing the target value to the middle element of the array.
// If the target value is equal to the middle element, the search is successful.
// If the target value is less than the middle element, the search continues on the lower half of the array.
// If the target value is greater than the middle element, the search continues on the upper half of the array.
// This process continues until the target value is found or the remaining array is empty.
// Binary search is an efficient algorithm for finding an item from a sorted list of items.
// It works by repeatedly dividing in half the portion of the list that could contain the item, until you've narrowed down the possible locations to just one.
// It is called in-place sorting because it does not require extra space for sorting.

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
