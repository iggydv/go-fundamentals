package sorting

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 2, 1, 5, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 2, 5, 4, -99, 0}, []int{-99, 0, 1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, test := range tests {
		result := BucketSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("BucketSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
