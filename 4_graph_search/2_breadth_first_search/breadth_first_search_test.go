package breadth_first_search

import (
	"reflect"
	"testing"
)

func TestQueueZeroPop(t *testing.T) {
	in := ""
	result := -1

	var q Queue
	q.init()

	if actual, r := q.pop(); actual != in || r != result {
		t.Errorf("Queue#pop() = %v, %v want %v, %v", actual, r, in, result)
	}
}

func TestQueueZeroEmpty(t *testing.T) {
	result := true

	var q Queue
	q.init()

	if r := q.empty(); r != result {
		t.Errorf("Queue#empty() = %v want %v", r, result)
	}
}

func TestQueuePushAndPop(t *testing.T) {
	tests := []struct {
		in     []string
		result []int
		out    []string
	}{
		{[]string{"sample"}, []int{0}, []string{"sample"}},
		{[]string{"sample1", "sample2", "sample3"}, []int{2, 1, 0}, []string{"sample1", "sample2", "sample3"}},
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
		for i := range test.in {
			if actual, r := q.pop(); actual != test.out[i] || r != test.result[i] {
				t.Errorf("Queue#pop() = %v, %v want %v, %v", actual, r, test.out[i], test.result[i])
			}
		}
	}
}

func TestQueueEmpty(t *testing.T) {
	tests := []struct {
		in     []string
		result bool
	}{
		{[]string{"sample"}, false},
		{[]string{"sample1", "sample2", "sample3"}, false},
	}

	for _, test := range tests {
		var q Queue
		q.init()
		for _, in := range test.in {
			q.push(in)
		}
		if r := q.empty(); r != test.result {
			t.Errorf("Queue#empty() = %v want %v", r, test.result)
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
				Graph{"B", []string{"E", "F", "A"}},
				Graph{"C", []string{"H", "A"}},
				Graph{"D", []string{"I", "J", "A"}},
				Graph{"E", []string{"K", "B"}},
				Graph{"F", []string{"B"}},
				Graph{"G", []string{"H"}},
				Graph{"H", []string{"G", "C"}},
				Graph{"I", []string{"D"}},
				Graph{"J", []string{"L", "D"}},
				Graph{"K", []string{"E"}},
				Graph{"L", []string{"J"}}}, true},
		{"B", "L",
			[]Graph{
				Graph{"A", []string{"B", "C", "D"}},
				Graph{"B", []string{"E", "F", "A"}},
				Graph{"C", []string{"H", "A"}},
				Graph{"D", []string{"I", "J", "A"}},
				Graph{"E", []string{"K", "B"}},
				Graph{"F", []string{"B"}},
				Graph{"G", []string{"H"}},
				Graph{"H", []string{"G", "C"}},
				Graph{"I", []string{"D"}},
				Graph{"J", []string{"L", "D"}},
				Graph{"K", []string{"E"}},
				Graph{"L", []string{"J"}}}, true},
		{"B", "I",
			[]Graph{
				Graph{"A", []string{"B", "C"}},
				Graph{"B", []string{"E", "F", "A"}},
				Graph{"C", []string{"H", "A"}},
				Graph{"D", []string{"I", "J"}},
				Graph{"E", []string{"K", "B"}},
				Graph{"F", []string{"B"}},
				Graph{"G", []string{"H"}},
				Graph{"H", []string{"G", "C"}},
				Graph{"I", []string{"D"}},
				Graph{"J", []string{"L", "D"}},
				Graph{"K", []string{"E"}},
				Graph{"L", []string{"J"}}}, false},
	}

	for _, test := range tests {
		if actual := BreadthFirstSearch(test.in_start, test.in_end, test.in_graph); actual != test.result {
			t.Errorf("BreadthFirstSearch(%v, %v, %v) = %v, want %v", test.in_start, test.in_end, test.in_graph, actual, test.result)
		}
	}
}
