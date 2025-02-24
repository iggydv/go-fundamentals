package search

// Given a sorted array, find the target element in the array using the two-pointer technique.
// Time: O(N)
// Space: O(1)
// The two-pointer technique is a simple and effective technique for solving problems with arrays.
// The technique involves using two pointers to traverse the array and find the desired element.
// The two pointers start at the beginning and end of the array and move towards each other until they meet.
// The technique is useful for solving problems that involve searching for a target element in a sorted array.

func TwoPointerSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		if arr[left] == target {
			return left
		}

		if arr[right] == target {
			return right
		}

		left++
		right--
	}

	return -1
}

// Given a sorted array, find the target element in the array using the fast and slow pointer technique.
// Time: O(N)
// Space: O(1)
// The fast and slow pointer technique is a two-pointer technique that is used to solve problems with linked lists and arrays.
// The fast pointer moves two steps at a time, while the slow pointer moves one step at a time. The technique is used to find
// the middle of the linked list or to find a cycle in the linked list.
// The technique is useful for solving problems that involve searching for a target element in a sorted array.

func fastAndSlowPointer(arr []int, target int) int {
	slow := 0
	fast := 1

	for fast < len(arr) {
		if arr[fast] == target {
			return fast
		}

		slow++
		fast = fast + 2
	}

	return -1
}
