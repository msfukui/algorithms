package Dijkstra_algorithm

import "fmt"

// ある頂点とそこに至る辺の重みを表現する型
type Edge struct {
	vertex string // ある頂点
	weight int    // 辺の重み
}

// グラフの要素の型
type Graph struct {
	vertex string // ある頂点
	edges  []Edge // ある頂点と隣り合わせの頂点とその辺の重みの集合
}

// 始点からある頂点への移動にかかるコストを格納する型
type Cost struct {
	vertex string // ある頂点
	cost   int    // 始点からある頂点に到達するためにかかる最小コスト
	from   string // 最小コストとなる隣り合わせの移動元の頂点
}

const max_weight = 100000 // 初期値に設定する始点以外の各頂点の重み(アルゴリズム上は無限大)

// DijkstraAlgorithm は、
// 指定された始点と終点と検索対象の重み付きグラフから、
// ダイクストラ法で探索した最小経路とそのコストを返す。
func DijkstraAlgorithm(start string, end string, graph []Graph) (route []string, cost int) {
	fmt.Printf("input  = %v, %v, %v\n", start, end, graph)

	// 初期設定
	costs := []Cost{} // 始点から他の各頂点への移動にかかるコストテーブル
	for _, v := range graph {
		c := Cost{"", 0, ""}
		if v.vertex == start {
			c = Cost{v.vertex, 0, v.vertex} // 始点への移動コストは0
		} else {
			c = Cost{v.vertex, max_weight, ""} // 始点以外の移動コストの初期値は無限大
		}
		costs = append(costs, c)
	}

	// 始点から探索を開始
	red := start
	green := []Edge{}
	orange := []string{start}

	// 探索が終点になるまで繰り返し
	for red != end {

		for i, v := range graph {

			fmt.Printf("costs = %v, red = %v, green = %v, orange = %v\n", costs, red, green, orange)

			if v.vertex == red {

				// 探索候補となる頂点を絞り込む
				// 経路が確定している頂点を取り除く
				green = []Edge{}
				merge_flag := true
				for _, s := range v.edges {
					for _, t := range orange {
						if s.vertex == t {
							merge_flag = false
							break
						}
					}
					if merge_flag {
						green = append(green, Edge{s.vertex, s.weight})
					}
					merge_flag = true
				}
				fmt.Printf("green = %v\n", green)

				// 探索候補への最短路を計算する
				f := max_weight
				for _, w := range green {

					// 探索対象のコストテーブルでの添え字を取得する
					redi := 0
					for j, x := range costs {
						if x.vertex == w.vertex {
							redi = j
							break
						}
					}

					// 探索先の頂点のコストが、現在の頂点のコストと辺の重みの合計よりも大きい場合は、
					// 現在の頂点のコストと辺の重みの合計で上書きする
					if costs[redi].cost > costs[i].cost+w.weight {
						costs[redi].cost = costs[i].cost + w.weight
					}

					// 探索候補内で未確定の頂点のうち最も小さいコストの頂点を保存する
					if f > costs[redi].cost {
						f = costs[redi].cost
						red = costs[redi].vertex
					}
				}
				fmt.Printf("red = %v, green = %v, orange = %v\n", red, green, orange)
			}
			// 次の探索対象を確定する
			orange = append(orange, red)
			break
		}
	}

	// 始点から終点の最短コストを設定する
	for _, v := range costs {
		if v.vertex == end {
			cost = v.cost
		}
	}

	fmt.Printf("output = %v, %v\n", route, cost)

	return route, cost
}
