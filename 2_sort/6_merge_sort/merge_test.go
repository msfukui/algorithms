package merge

import (
	"reflect"
	"testing"
)

func TestMergeSortEmpty(t *testing.T) {
	const success = true
	input := []int{}
	expected := []int{}

	if r, actual := MergeSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("MergeSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestMergeSortOne(t *testing.T) {
	const success = true
	input := []int{1}
	expected := []int{1}

	if r, actual := MergeSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("MergeSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestMergeSortSuccess(t *testing.T) {
	const success = true
	input := []int{6, 4, 3, 7, 5, 1, 2}
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if r, actual := MergeSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("MergeSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestMergeSortDuplexSuccess(t *testing.T) {
	const success = true
	input := []int{6, 4, 3, 7, 4, 5, 1, 2}
	expected := []int{1, 2, 3, 4, 4, 5, 6, 7}

	if r, actual := MergeSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("MergeSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}
