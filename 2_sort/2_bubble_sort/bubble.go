package main

import "fmt"

// BubbleSort は
// 与えられた整数のスライスから、
// ソートの結果とバブルソートで並び替えた結果の整数のスライスを返します。
// c.f. No.2-2 バブルソート
func BubbleSort(list []int) (success bool, sortedList []int) {
	// 返却値用のバッファを作ってコピー
	sortedList = make([]int, len(list))
	copy(sortedList, list)

	balance := len(list) - 1 // 天秤の場所
	buf := 0                 // 入れ替え用のバッファ

	fmt.Printf("input  = %v\n", list)

	// ソート済みの場所を先頭から一つずつ進める
	for i := 0; i < len(list)-1; i++ {
		// 状況把握のため、並び替え途中の情報を出力する
		fmt.Printf("  list(%v) = %v\n", i+1, sortedList)
		// 末尾からソート済みの手前の場所まで逆順で検索を実施する
		for j := balance; j > i; j-- {
			// 右より左が大きければ左右を入れ替え
			if sortedList[j-1] > sortedList[j] {
				buf = sortedList[j-1]
				sortedList[j-1] = sortedList[j]
				sortedList[j] = buf
			}
		}
	}

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}
