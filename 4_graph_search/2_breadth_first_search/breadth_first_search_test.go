package breadth_first_search

import (
	"reflect"
	"testing"
)

func TestQueueZero(t *testing.T) {
	in := ""
	result := -1

	var q Queue
	q.init()

	if actual, r := q.pop(); actual != in || r != result {
		t.Errorf("Queue#pop() = %v, %v want %v, %v", actual, r, in, result)
	}
}

func TestQueue(t *testing.T) {
	tests := []struct {
		in     []string
		result []int
	}{
		{[]string{"sample"}, []int{0}},
		{[]string{"sample1", "sample2", "sample3"}, []int{2, 1, 0}},
	}

	for _, test := range tests {
		var q Queue
		q.init()
		for _, in := range test.in {
			q.push(in)
		}
		if actual := q.el; !reflect.DeepEqual(actual, test.in) {
			t.Errorf("Queue#push(%v) = %v, want %v", test.in, actual, test.in)
		}
		for i, in := range test.in {
			if actual, r := q.pop(); actual != in || r != test.result[i] {
				t.Errorf("Queue#pop() = %v, %v want %v, %v", actual, r, in, test.result[i])
			}
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		in_start string
		in_end   string
		in_graph []Graph
		result   bool
	}{
		{"", "", []Graph{Graph{"", []string{""}}}, true},
		{"A", "G",
			[]Graph{
				Graph{"A", []string{"B", "C", "D"}},
				Graph{"B", []string{"E", "F"}},
				Graph{"C", []string{"H"}},
				Graph{"D", []string{"I", "J"}},
				Graph{"E", []string{"K"}},
				Graph{"F", []string{""}},
				Graph{"G", []string{""}},
				Graph{"H", []string{"G"}},
				Graph{"I", []string{""}},
				Graph{"J", []string{"L"}},
				Graph{"K", []string{""}},
				Graph{"L", []string{""}}}, true},
		{"B", "L",
			[]Graph{
				Graph{"A", []string{"B", "C", "D"}},
				Graph{"B", []string{"E", "F"}},
				Graph{"C", []string{"H"}},
				Graph{"D", []string{"I", "J"}},
				Graph{"E", []string{"K"}},
				Graph{"F", []string{""}},
				Graph{"G", []string{""}},
				Graph{"H", []string{"G"}},
				Graph{"I", []string{""}},
				Graph{"J", []string{"L"}},
				Graph{"K", []string{""}},
				Graph{"L", []string{""}}}, false},
	}

	for _, test := range tests {
		if actual := BreadthFirstSearch(test.in_start, test.in_end, test.in_graph); actual != test.result {
			t.Errorf("BreadthFirstSearch(%v, %v, %v) = %v, want %v", test.in_start, test.in_end, test.in_graph, actual, test.result)
		}
	}
}
