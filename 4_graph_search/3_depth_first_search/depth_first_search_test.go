package depth_first_search

import (
	"reflect"
	"testing"
)

func TestStackZeroPop(t *testing.T) {
	in := ""
	result := -1

	var s Stack
	s.init()

	if actual, r := s.pop(); actual != in || r != result {
		t.Errorf("Stack#pop() = %v, %v want %v, %v", actual, r, in, result)
	}
}

func TestStackZeroEmpty(t *testing.T) {
	result := true

	var s Stack
	s.init()

	if r := s.empty(); r != result {
		t.Errorf("Stack#empty() = %v want %v", r, result)
	}
}

func TestStackPushAndPop(t *testing.T) {
	tests := []struct {
		in     []string
		result []int
		out    []string
	}{
		{[]string{"sample"}, []int{0}, []string{"sample"}},
		{[]string{"sample1", "sample2", "sample3"}, []int{2, 1, 0}, []string{"sample3", "sample2", "sample1"}},
	}

	for _, test := range tests {
		var s Stack
		s.init()
		for _, in := range test.in {
			s.push(in)
		}
		if actual := s.el; !reflect.DeepEqual(actual, test.in) {
			t.Errorf("Stack#push(%v) = %v, want %v", test.in, actual, test.in)
		}
		for i, _ := range test.in {
			if actual, r := s.pop(); actual != test.out[i] || r != test.result[i] {
				t.Errorf("Stack#pop() = %v, %v want %v, %v", actual, r, test.out[i], test.result[i])
			}
		}
	}
}

func TestStackEmpty(t *testing.T) {
	tests := []struct {
		in     []string
		result bool
	}{
		{[]string{"sample"}, false},
		{[]string{"sample1", "sample2", "sample3"}, false},
	}

	for _, test := range tests {
		var s Stack
		s.init()
		for _, in := range test.in {
			s.push(in)
		}
		if r := s.empty(); r != test.result {
			t.Errorf("Stack#empty() = %v want %v", r, test.result)
		}
	}
}

func TestDepthFirstSearch(t *testing.T) {
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
		if actual := DepthFirstSearch(test.in_start, test.in_end, test.in_graph); actual != test.result {
			t.Errorf("DepthFirstSearch(%v, %v, %v) = %v, want %v", test.in_start, test.in_end, test.in_graph, actual, test.result)
		}
	}
}
