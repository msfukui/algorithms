package insertion

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{5, 3, 4, 7, 2, 8, 6, 9, 1}, true, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{5, 3, 4, 7, 2, 5, 8, 6, 9, 1}, true, []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		if r, actual := InsertionSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("InsertionSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
