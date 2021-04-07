package heap

import (
	"reflect"
	"testing"
)

func TestHeapAdd(t *testing.T) {
	tests := []struct {
		in_list  []int
		in_value int
		result   bool
		expected []int
	}{
		{[]int{}, 5, true, []int{5}},
		{[]int{1, 3, 6, 4, 8, 7}, 5, true, []int{1, 3, 5, 4, 8, 7, 6}},
		{[]int{1, 3, 5, 4, 8, 7, 6}, 5, true, []int{1, 3, 5, 4, 8, 7, 6, 5}},
	}

	for _, test := range tests {
		if r, actual := HeapAdd(test.in_list, test.in_value); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("HeapAdd(%v, %v) = %v, %v, want %v, %v", test.in_list, test.in_value, r, actual, test.result, test.expected)
		}
	}
}

func TestHeapDel(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, false, []int{}},
		{[]int{6}, true, []int{}},
		{[]int{3, 6}, true, []int{6}},
		{[]int{3, 7, 6}, true, []int{6, 7}},
		{[]int{3, 7, 6, 9}, true, []int{6, 7, 9}},
		{[]int{1, 3, 5, 4, 8, 7, 6}, true, []int{3, 4, 5, 6, 8, 7}},
	}

	for _, test := range tests {
		if r, actual := HeapDel(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("HeapDel(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		in       []int
		result   bool
		expected []int
	}{
		{[]int{}, true, []int{}},
		{[]int{5, 2, 7, 3, 6, 1, 4}, true, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{5, 2, 7, 3, 5, 6, 1, 4}, true, []int{1, 2, 3, 4, 5, 5, 6, 7}},
	}

	for _, test := range tests {
		if r, actual := HeapSort(test.in); r != test.result || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("HeapSort(%v) = %v, %v, want %v, %v", test.in, r, actual, test.result, test.expected)
		}
	}
}
