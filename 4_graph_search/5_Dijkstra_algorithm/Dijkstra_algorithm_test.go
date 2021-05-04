package Dijkstra_algorithm

import (
	"reflect"
	"testing"
)

func TestDijkstraAlgorithm(t *testing.T) {
	tests := []struct {
		in_start  string
		in_end    string
		in_graph  []Graph
		out_route []string
		out_cost  int
	}{
		{"", "",
			[]Graph{
				Graph{"", []Edge{Edge{"", 0}}},
			}, []string{"", ""}, 0},
		{"A", "B",
			[]Graph{
				Graph{"A", []Edge{Edge{"B", 2}}},
				Graph{"B", []Edge{Edge{"A", 2}}},
			}, []string{"A", "B"}, 2},
		{"A", "C",
			[]Graph{
				Graph{"A", []Edge{Edge{"B", 2}, Edge{"C", 4}}},
				Graph{"B", []Edge{Edge{"A", 2}, Edge{"C", 1}}},
				Graph{"C", []Edge{Edge{"A", 4}, Edge{"B", 1}}},
			}, []string{"A", "B", "C"}, 3},
		{"A", "C",
			[]Graph{
				Graph{"A", []Edge{Edge{"B", 2}, Edge{"C", 4}}},
				Graph{"B", []Edge{Edge{"A", 2}, Edge{"C", 2}}},
				Graph{"C", []Edge{Edge{"A", 4}, Edge{"B", 2}}},
			}, []string{"A", "C"}, 4},
		{"A", "G",
			[]Graph{
				Graph{"A", []Edge{Edge{"B", 9}, Edge{"C", 2}}},
				Graph{"B", []Edge{Edge{"A", 9}, Edge{"C", 6}, Edge{"D", 3}, Edge{"E", 1}}},
				Graph{"C", []Edge{Edge{"A", 2}, Edge{"B", 6}, Edge{"D", 2}, Edge{"F", 9}}},
				Graph{"D", []Edge{Edge{"B", 3}, Edge{"C", 2}, Edge{"E", 5}, Edge{"F", 6}}},
				Graph{"E", []Edge{Edge{"B", 1}, Edge{"D", 5}, Edge{"F", 3}, Edge{"G", 7}}},
				Graph{"F", []Edge{Edge{"C", 9}, Edge{"D", 6}, Edge{"E", 3}, Edge{"G", 4}}},
				Graph{"G", []Edge{Edge{"E", 7}, Edge{"F", 4}}},
			}, []string{"A", "C", "D", "F", "G"}, 14},
		{"A", "G",
			[]Graph{
				Graph{"A", []Edge{Edge{"B", 2}, Edge{"C", 5}}},
				Graph{"B", []Edge{Edge{"A", 2}, Edge{"C", 6}, Edge{"D", 1}, Edge{"E", 3}}},
				Graph{"C", []Edge{Edge{"A", 5}, Edge{"B", 6}, Edge{"F", 8}}},
				Graph{"D", []Edge{Edge{"B", 1}, Edge{"E", 4}}},
				Graph{"E", []Edge{Edge{"B", 3}, Edge{"D", 4}, Edge{"G", 9}}},
				Graph{"F", []Edge{Edge{"C", 8}, Edge{"G", 7}}},
				Graph{"G", []Edge{Edge{"E", 9}, Edge{"F", 7}}},
			}, []string{"A", "B", "E", "G"}, 14},
	}

	for _, test := range tests {
		actual_route, actual_cost := DijkstraAlgorithm(test.in_start, test.in_end, test.in_graph)
		if !reflect.DeepEqual(actual_route, test.out_route) || actual_cost != test.out_cost {
			t.Errorf("DijkstraAlgorithm(%v, %v, %v) = %v, %v want %v, %v", test.in_start, test.in_end, test.in_graph, actual_route, actual_cost, test.out_route, test.out_cost)
		}
	}
}
