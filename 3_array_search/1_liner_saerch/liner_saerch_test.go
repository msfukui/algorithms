package liner_search

import (
	"testing"
)

func TestLinerSearch(t *testing.T) {
	tests := []struct {
		in_array []int
		in_value int
		result   bool
		expected int
	}{
		{[]int{}, 0, false, 0},
		{[]int{1}, 0, false, 0},
		{[]int{1}, 1, true, 0},
		{[]int{3, 9, 8, 2, 1, 4, 6, 5, 7}, 6, true, 6},
		{[]int{3, 9, 8, 2, 1, 4, 6, 5, 7}, 10, false, 0},
	}

	for _, test := range tests {
		if r, actual := LinerSearch(test.in_array, test.in_value); r != test.result || actual != test.expected {
			t.Errorf("LinerSearch(%v, %v) = %v, %v, want %v, %v", test.in_array, test.in_value, r, actual, test.result, test.expected)
		}
	}
}
