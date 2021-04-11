package liner_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		in_array []int
		in_value int
		result   bool
		expected int
	}{
		{[]int{}, 0, false, 0},
		{[]int{1}, 0, false, 0},
		{[]int{1}, 1, true, 0},
		{[]int{1, 2}, 1, true, 0},
		{[]int{1, 2}, 2, true, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6, true, 5},
		{[]int{1, 2, 3, 4, 5, 6, 6, 8, 9}, 6, true, 6}, // 同じ値が含まれる場合は先に見つかった方を返す
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, false, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, false, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, true, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, true, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, true, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, true, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, true, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, true, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8, true, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, true, 8},
	}

	for _, test := range tests {
		if r, actual := BinarySearch(test.in_array, test.in_value); r != test.result || actual != test.expected {
			t.Errorf("BinarySearch(%v, %v) = %v, %v, want %v, %v", test.in_array, test.in_value, r, actual, test.result, test.expected)
		}
	}
}
