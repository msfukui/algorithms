package merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{1}, true, []int{1}},
		{[]int{6, 4, 3, 7, 5, 1, 2}, true, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{6, 4, 3, 7, 4, 5, 1, 2}, true, []int{1, 2, 3, 4, 4, 5, 6, 7}},
	}

	for _, test := range tests {
		if r, actual := MergeSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("MergeSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
