package liner_search

import "fmt"

// BinarySearch は
// ソート済みの検索対象のスライスと検索したい値から
// 一致する値があれば、一致したことと、そのスライスの位置（添え字）を返します。
// 一致する値がなければ、一致しなかったことを返します。
// c.f. No.3-2 2分探索
func BinarySearch(array []int, value int) (result bool, index int) {
	// スライスが空の場合は見つからなかったとして終了する
	if len(array) == 0 {
		return false, 0
	}

	fmt.Printf("input  = %v, %v\n", array, value)

	// スライスの中央値を比較対象として、最初に検索する添え字を決める
	m := 0              // 検索空間の下限
	n := len(array) - 1 // 検索空間の上限
	k := n - m + 1      // 検索数
	i := k / 2          // 検索対象の添え字の初期値

	// 検索数がなくなるまで半分ずつ検索を進める
	for m <= n {
		fmt.Printf("  i = %v, m = %v, n = %v, k = %v\n", i, m, n, k)

		if array[i] == value {
			// 検索したい値と一致したら終了する
			fmt.Printf("output = %v\n", i)
			return true, i
		} else if array[i] < value {
			// 検索したい値の方が大きければ、
			// 大きい方の半分を次の検索対象にする
			m = i + 1
			k = n - m + 1
			if k == 1 {
				i = i + 1
			} else {
				i = i + (k / 2)
			}
		} else {
			// 検索したい値の方が小さければ、
			// 小さい方の半分を次の検索対象にする
			n = i - 1
			k = n - m + 1
			if k == 1 {
				i = i - 1
			} else {
				i = i - (k / 2)
			}
		}
	}

	fmt.Printf("output = %v\n", 0)

	// ここまできたら見つからなかったとして終了する
	return false, 0
}
