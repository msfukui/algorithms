package main

import "fmt"

// SelectionSort は
// 与えられた整数のスライスから、
// ソートの結果とバブルソートで並び替えた結果の整数のスライスを返します。
// c.f. No.2-3 選択ソート
func SelectionSort(list []int) (success bool, sortedList []int) {
	// 返却値用のバッファを作ってコピー
	sortedList = make([]int, len(list))
	copy(sortedList, list)

	min_buf := 0 // 一番小さい数を覚えておく変数
	min_idx := 0 // 一番小さい数の位置を覚えておく変数

	fmt.Printf("input  = %v\n", list)

	// ソート済みの場所を先頭から一つずつ進める
	for i := 0; i < len(list)-1; i++ {
		// 状況把握のため、並び替え途中の情報を出力する
		fmt.Printf("  list(%v) = %v\n", i+1, sortedList)

		min_buf = sortedList[i]
		min_idx = i
		for j := i; j < len(list); j++ {
			if min_buf > sortedList[j] {
				min_buf = sortedList[j]
				min_idx = j
			}
		}
		sortedList[min_idx] = sortedList[i]
		sortedList[i] = min_buf
	}

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}
