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
	empty() bool        // 空かどうかを判定する
}

// キューの初期化
func (q *Queue) init() {
	q.el = make([]string, 0)
}

// キューに積む
func (q *Queue) push(str string) {
	q.el = append(q.el, str)
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

// キューが空かどうかを判定する(空なら真)
func (q *Queue) empty() bool {
	return len(q.el) == 0
}

// BreadthFirstSerch は、
// 指定された始点と終点と検索対象のグラフから、
// 幅優先探索(breadth first search) で検索した到達可否を返す
func BreadthFirstSearch(start string, end string, graph []Graph) (result bool) {
	fmt.Printf("input  = %v, %v, %v\n", start, end, graph)

	// 始点と終点が同じなら到達したものと判定して終了する
	if start == end {
		result = true
		fmt.Printf("output = %v\n", result)
		return result
	}

	// 初期設定
	red := ""       // 現在の探索対象の点
	var green Queue // 探索対象の候補の点を格納するキュー
	green.init()
	green.push(start)
	orange := []string{} // 探索済みの点
	result = false

	// 探索候補がなくなるまで一つずつ探索する
	for !green.empty() {
		fmt.Printf("red = %v, ", red)
		fmt.Printf("green = %v, ", green)
		fmt.Printf("orange = %v\n", orange)

		// 探索候補から一つ取り出す
		red, _ = green.pop()

		// 既に探索済みの点の場合、探索をスキップして次に進む
		searched := false
		for _, v := range orange {
			if v == red {
				searched = true
				break
			}
		}
		if searched {
			continue
		}

		// 終点と一致していたら到達したので終了する
		if red == end {
			result = true
			break
		}

		// 次の探索候補を探してキューに入れる
		for _, v := range graph {
			if v.vertex == red {
				for _, s := range v.edges {
					green.push(s)
				}
				break
			}
		}

		// 探索が終わったら探索済みリストに追加する
		orange = append(orange, red)
	}

	fmt.Printf("output = %v\n", result)

	return result
}
