package sorting

// Time: O(n^2) - worst case
// Space: O(1)

// Insertion sort works by taking elements from the list one by one and inserting them in their correct position in a
// new sorted list. The choice of the element to be inserted is made by a key. The key is compared with the elements of
// the new list, and the elements are shifted accordingly.
// It is called in-place sorting because it does not require extra space for sorting.
// Useful

func InsertionSort(arr []int) []int {
	for unsortedIndex := 1; unsortedIndex < len(arr); unsortedIndex++ {
		key := arr[unsortedIndex]
		sortedIndex := unsortedIndex - 1

		for sortedIndex >= 0 && arr[sortedIndex] > key {
			arr[sortedIndex+1] = arr[sortedIndex]
			sortedIndex--
		}
		arr[sortedIndex+1] = key
	}
	return arr
}
