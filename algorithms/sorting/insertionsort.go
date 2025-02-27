package sorting

// Time: O(n^2) - worst case
// Space: O(1)

// Insertion sort works by taking elements from the list one by one and inserting them in their correct position in a
// new sorted list. The choice of the element to be inserted is made by a key. The key is compared with the elements of
// the new list, and the elements are shifted accordingly.
// It is called in-place sorting because it does not require extra space for sorting.
// Useful

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
