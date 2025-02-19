package sorting

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	type fields struct {
		input []int
		start int
		end   int
	}
	tests := []struct {
		fields   fields
		expected []int
	}{
		{fields{[]int{3, 2, 1, 5, 4}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{fields{[]int{1, 2, 3, 4, 5}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{fields{[]int{5, 4, 3, 2, 1}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{fields{[]int{1, 3, 2, 5, 4, -99, 0}, 0, 6}, []int{-99, 0, 1, 2, 3, 4, 5}},
		{fields{[]int{}, 0, 0}, []int{}},
		{fields{[]int{1}, 0, 0}, []int{1}},
	}

	for _, test := range tests {
		result := Quicksort(test.fields.input, test.fields.start, test.fields.end)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MergeSort(%v) = %v; want %v", test.fields.input, result, test.expected)
		}
	}
}
