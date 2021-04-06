package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSortEmpty(t *testing.T) {
	const success = true
	input := []int{}
	expected := []int{}

	if r, actual := BubbleSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("BubbleSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestBubbleSortSuccess(t *testing.T) {
	const success = true
	input := []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if r, actual := BubbleSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("BubbleSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}

func TestBubbleSortDuplexSuccess(t *testing.T) {
	const success = true
	input := []int{5, 9, 3, 1, 2, 5, 8, 4, 7, 6}
	expected := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}

	if r, actual := BubbleSort(input); r != success || !reflect.DeepEqual(actual, expected) {
		t.Errorf("BubbleSort(%v) = %v, %v, want %v, %v", input, r, actual, success, expected)
	}
}
