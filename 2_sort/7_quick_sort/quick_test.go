package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{1}, true, []int{1}},
		{[]int{2, 1}, true, []int{1, 2}},
		{[]int{3, 2, 1}, true, []int{1, 2, 3}},
		{[]int{3, 5, 8, 1, 2, 9, 4, 7, 6}, true, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{3, 5, 8, 1, 2, 5, 9, 4, 7, 6}, true, []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		if r, actual := QuickSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("QuickSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
