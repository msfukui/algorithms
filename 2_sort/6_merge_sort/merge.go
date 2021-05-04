package merge

import "fmt"

// MergeSort は
// 与えられた整数のスライスから、
// ソートの結果とマージソート(しめじソート)で並び替えた結果の整数のスライスを返します。
// c.f. No.2-6 マージソート
func MergeSort(list []int) (success bool, sortedList []int) {
	fmt.Printf("input  = %v\n", list)

	if len(list) == 0 {
		fmt.Printf("output = %v\n", sortedList)
		return true, []int{}
	}

	if len(list) == 1 {
		sortedList = make([]int, len(list))
		copy(sortedList, list)
		fmt.Printf("output = %v\n", sortedList)
		return true, sortedList
	}

	f := len(list) / 2 // 分割する前半の個数
	l := len(list) - f // 分割する後半の個数

	// 前半、後半の値を格納する領域を確保して値をコピーする
	listA := make([]int, f)
	copy(listA, list[0:f])
	listB := make([]int, l)
	copy(listB, list[f:])

	_, sortedListA := MergeSort(listA)           // 前半部分を再帰で最小まで分割して並び替える
	_, sortedListB := MergeSort(listB)           // 後半部分を再帰で最小まで分割して並び替える
	sortedList = Merge(sortedListA, sortedListB) // 分割した値を並び替えて統合する

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}

// Merge は
// 与えられたソート済みの2つの整数のスライスから、
// 並び替えて1つに統合したスライスを生成して返します。
func Merge(listA []int, listB []int) (list []int) {
	// 返却値用のバッファを作る
	list = make([]int, len(listA)+len(listB))

	i, j := 0, 0
	for i < len(listA) || j < len(listB) {
		if j >= len(listB) || (i < len(listA) && listA[i] < listB[j]) {
			list[i+j] = listA[i]
			i++
		} else {
			list[i+j] = listB[j]
			j++
		}
	}

	return list
}
