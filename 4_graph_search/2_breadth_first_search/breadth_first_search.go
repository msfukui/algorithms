package breadth_first_search

import "fmt"

// グラフの要素の型
type Graph struct {
	vertex string   // 起点
	edges  []string // 隣り合わせの点の集合
}

// 文字列のキュー構造
type Queue struct {
	el []string
}

// キューのインターフェース
type Queueable interface {
	init()              // 初期化
	push(string)        // 積む
	pop() (string, int) // 取り出す
}

// キューの初期化
func (q *Queue) init() {
	q.el = make([]string, 0)
	return
}

// キューに積む
func (q *Queue) push(s string) {
	q.el = append(q.el, s)
	return
}

// キューから取り出す
func (q *Queue) pop() (string, int) {
	if len(q.el) == 0 {
		return "", -1
	}
	v := q.el[0]
	q.el = q.el[1:]
	return v, len(q.el)
}

// BreadthFirstSerch は、
// 指定された始点と終点と検索対象のグラフから、
// 幅優先探索(breadth first search) で検索した到達可否を返す
func BreadthFirstSearch(start string, end string, graph []Graph) (result bool) {
	result = true

	// キューを使う
	var q Queue
	q.init()

	fmt.Printf("input  = %v, %v, %v\n", start, end, graph)

	for i, v := range graph {
		fmt.Printf("graph[%v] = {%v -> %v}\n", i, v.vertex, v.edges)
	}

	fmt.Printf("output = %v\n", result)

	return result
}
