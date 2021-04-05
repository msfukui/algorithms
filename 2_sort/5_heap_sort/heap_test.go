package main

import (
	"reflect"
	"testing"
)

func TestHeapAddEmpty(t *testing.T) {
	const success = true
	input_list := []int{}
	input_value := 5
	expected := []int{5}

	if r, actual := HeapAdd(input_list, input_value); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapAdd(%v, %v) = %v, %v want %v, %v", input_list, input_value, r, actual, success, expected)
	}
}

func TestHeapAddSuccess(t *testing.T) {
	const success = true
	input_list := []int{1, 3, 6, 4, 8, 7}
	input_value := 5
	expected := []int{1, 3, 5, 4, 8, 7, 6}

	if r, actual := HeapAdd(input_list, input_value); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapAdd(%v, %v) = %v, %v want %v, %v", input_list, input_value, r, actual, success, expected)
	}
}

func TestHeapAddDuplexSuccess(t *testing.T) {
	const success = true
	input_list := []int{1, 3, 5, 4, 8, 7, 6}
	input_value := 5
	expected := []int{1, 3, 5, 4, 8, 7, 6, 5}

	if r, actual := HeapAdd(input_list, input_value); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapAdd(%v, %v) = %v, %v want %v, %v", input_list, input_value, r, actual, success, expected)
	}
}

func TestHeapDelEmpty(t *testing.T) {
	const failure = false
	input_list := []int{}

	if r, _ := HeapDel(input_list); r != failure {
		t.Errorf("HeapDel(%v) = %v want %v", input_list, r, failure)
	}
}

func TestHeapDelOne(t *testing.T) {
	const success = true
	input_list := []int{6}
	expected := []int{}

	if r, actual := HeapDel(input_list); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapDel(%v) = %v, %v want %v, %v", input_list, r, actual, success, expected)
	}
}

func TestHeapDelTwo(t *testing.T) {
	const success = true
	input_list := []int{3, 6}
	expected := []int{6}

	if r, actual := HeapDel(input_list); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapDel(%v) = %v, %v want %v, %v", input_list, r, actual, success, expected)
	}
}

func TestHeapDelThree(t *testing.T) {
	const success = true
	input_list := []int{3, 7, 6}
	expected := []int{6, 7}

	if r, actual := HeapDel(input_list); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapDel(%v) = %v, %v want %v, %v", input_list, r, actual, success, expected)
	}
}

func TestHeapDelFour(t *testing.T) {
	const success = true
	input_list := []int{3, 7, 6, 9}
	expected := []int{6, 7, 9}

	if r, actual := HeapDel(input_list); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapDel(%v) = %v, %v want %v, %v", input_list, r, actual, success, expected)
	}
}

func TestHeapDelSuccess(t *testing.T) {
	const success = true
	input_list := []int{1, 3, 5, 4, 8, 7, 6}
	expected := []int{3, 4, 5, 6, 8, 7}

	if r, actual := HeapDel(input_list); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapDel(%v) = %v, %v want %v, %v", input_list, r, actual, success, expected)
	}
}

func TestHeapSortEmpty(t *testing.T) {
	const success = true
	input := []int{}
	expected := []int{}

	if r, actual := HeapSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestHeapSortSuccess(t *testing.T) {
	const success = true
	input := []int{5, 2, 7, 3, 6, 1, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if r, actual := HeapSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestHeapSortDuplexSuccess(t *testing.T) {
	const success = true
	input := []int{5, 2, 7, 3, 5, 6, 1, 4}
	expected := []int{1, 2, 3, 4, 5, 5, 6, 7}

	if r, actual := HeapSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("HeapSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}
