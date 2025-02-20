package sorting

import "sort"

// Time: O(n+k), O(n^2) - worst case
// Space: O(n+k)

// Bucket sort is a sorting algorithm that works by distributing the elements of an array into a number of buckets.
// Each bucket is then sorted individually, either using a different sorting algorithm, or by recursively applying the
// bucket sorting algorithm. It is called a distribution sort because it distributes elements of an array into a number of buckets.
// It is useful when input is uniformly distributed over a range.
// It is not a comparison sort, so it is not suitable for all types of data.

func BucketSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// find min and max ranges
	minVal, maxVal := nums[0], nums[0]

	for _, num := range nums {
		if num < minVal {
			minVal = num
		} else if num > maxVal {
			maxVal = num
		}
	}

	numBuckets := len(nums)
	buckets := make([][]int, numBuckets)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	for _, num := range nums {
		index := 0
		if num == maxVal {
			index = numBuckets - 1
		} else {
			index = (num - minVal) * (numBuckets - 1) / (maxVal - minVal)
		}
		buckets[index] = append(buckets[index], num)

	}

	// sort individual buckets
	for i := range buckets {
		sort.Ints(buckets[i])
	}

	// gather results - merge sorted buckets
	result := make([]int, 0)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}
	return result
}
