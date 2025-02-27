package sorting

// Time: O(n^2) - worst case
// Space: O(1)

// Bubble sort works by comparing each element in the list with the element next to it and swapping them if required.
// The algorithm repeats this process until it makes a pass all the way through the list without swapping any elements.
// This causes larger values to "bubble" to the end of the list while smaller values "sink" towards the beginning of the list.
// It is called in-place sorting because it does not require extra space for sorting.

func BubbleSort(arr []int) []int {
	n := len(arr) - 1
	swapped := false
	for i := 0; i < n; i++ {
		swapped = false
		for j := 0; j < n-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
