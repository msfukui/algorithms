package insertion

import (
	"reflect"
	"testing"
)

func TestInsertionSortEmpty(t *testing.T) {
	const success = true
	input := []int{}
	expected := []int{}

	if r, actual := InsertionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("InsertionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestInsertionSortSuccess(t *testing.T) {
	const success = true
	input := []int{5, 3, 4, 7, 2, 8, 6, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if r, actual := InsertionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("InsertionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestInsertionSortDuplexSuccess(t *testing.T) {
	const success = true
	input := []int{5, 3, 4, 7, 2, 5, 8, 6, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}

	if r, actual := InsertionSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("InsertionSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}
