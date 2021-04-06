package selection

import (
	"reflect"
	"testing"
)

func TestSelectionSortEmpty(t *testing.T) {
	const success = true
	input := []int{}
	expected := []int{}

	if r, actual := SelectionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("SelectionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestSelectionSortSuccess(t *testing.T) {
	const success = true
	input := []int{6, 1, 7, 8, 9, 3, 5, 4, 2}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if r, actual := SelectionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("SelectionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestSelectionSortDuplexSuccess(t *testing.T) {
	const success = true
	input := []int{6, 1, 7, 8, 9, 5, 3, 5, 4, 2}
	expected := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}

	if r, actual := SelectionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("SelectionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}
