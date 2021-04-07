package selection

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{6, 1, 7, 8, 9, 3, 5, 4, 2}, true, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{6, 1, 7, 8, 9, 5, 3, 5, 4, 2}, true, []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		if r, actual := SelectionSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("SelectionSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
