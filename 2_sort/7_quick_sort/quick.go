package quick

import (
	"fmt"
	"math/rand"
	"time"
)

// QuickSort は
// 与えられた整数のスライスから、
// ソートの結果とクイックソート(じゃがいもソート)で並び替えた結果の整数のスライスを返します。
// c.f. No.2-7 クイックソート
func QuickSort(list []int) (success bool, sortedList []int) {
	// 返却値用のバッファを作る
	sortedList = make([]int, len(list))

	fmt.Printf("input  = %v\n", list)

	// リストが空なら空を返す
	if len(list) == 0 {
		fmt.Printf("output = %v\n", sortedList)
		return true, []int{}
	}

	// リストが一つならそのまま返す
	if len(list) == 1 {
		sortedList[0] = list[0]
		fmt.Printf("output = %v\n", sortedList)
		return true, sortedList
	}

	// 乱数から pivot を設定
	rand.Seed(time.Now().UnixNano())
	p := rand.Intn(len(list))
	pivot := list[p]
	fmt.Printf("pivot(%v) = %v\n", p, pivot)

	// 振り分け先の格納領域を用意する
	listA, listB := []int{}, []int{}
	// pivot を基準に小さいものと大きいものに振り分ける
	for i, v := range list {
		if i == p {
			continue
		} else if v < pivot {
			listA = append(listA, v)
		} else {
			listB = append(listB, v)
		}
	}

	// pivot より小さいリストをソートする
	_, sortedListA := QuickSort(listA)
	// pivot より大きいリストをソートする
	_, sortedListB := QuickSort(listB)

	// ソート済みのリストを連結する
	j := 0
	for _, v := range sortedListA {
		sortedList[j] = v
		j++
	}
	sortedList[j] = pivot
	j++
	for _, v := range sortedListB {
		sortedList[j] = v
		j++
	}

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}
