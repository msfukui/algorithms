package Bellman_Ford_algorithm

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

// BellmanFordAlgorithm は、
// 指定された始点と終点と検索対象の重み付きグラフから、
// ベルマン・フォード法で探索した最小経路とそのコストを返す。
func BellmanFordAlgorithm(start string, end string, graph []Graph) (route []string, cost int) {
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

	// 頂点数分を繰り返す
	for range graph {

		// 各頂点のコストが一度でも変更されたかどうかを示すフラグ
		fixFlag := false

		// 頂点ごとの、
		for _, v := range graph {

			// 起点の重みを取得し、
			s_weight := 0
			for _, c := range costs {
				if c.vertex == v.vertex {
					s_weight = c.cost
				}
			}

			// 各辺に対して、
			for _, green := range v.edges {

				for j, k := range costs {

					// 終点の重みが、起点の重み + 辺の重みより大きい場合は、
					if k.vertex == green.vertex && k.cost > s_weight+green.weight {
						// 終点の重みを、起点の重み + 辺の重みで上書きする
						costs[j].cost = s_weight + green.weight
						// 移動元の頂点を合わせて保存する
						costs[j].from = v.vertex
						// 変更されたことを示すフラグを立てる
						fixFlag = true

						fmt.Printf("v = %v, costs = %v, green = %v\n", v, costs, green)
					}
				}
			}
		}

		// 各頂点のコストが一度も変更されていなければ、
		// それ以上の変更は不要なため、終了する
		if fixFlag == false {
			break
		}
	}

	// コストテーブルから始点から終点の最短路を求める
	// 終点から逆順に始点までたどる
	route = append(route, end)
	buf := end
	for range graph {
		for _, v := range costs {
			if v.vertex == buf {
				route = append([]string{v.from}, route...)
				buf = v.from
				break
			}
		}
		// 始点までたどり着いたら終了
		if route[0] == start {
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
