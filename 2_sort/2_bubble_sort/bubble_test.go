package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{5, 9, 3, 1, 2, 8, 4, 7, 6}, true, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{5, 9, 3, 1, 2, 5, 8, 4, 7, 6}, true, []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		if r, actual := BubbleSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("BubbleSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
