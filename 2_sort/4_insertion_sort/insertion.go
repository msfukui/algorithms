package insertion

import "fmt"

// InsertionSort は
// 与えられた整数のスライスから、
// ソートの結果と挿入ソートで並び替えた結果の整数のスライスを返します。
// c.f. No.2-4 挿入ソート
func InsertionSort(list []int) (success bool, sortedList []int) {
	// 返却値用のバッファを作ってコピー
	sortedList = make([]int, len(list))
	copy(sortedList, list)

	buf := 0 // 入れ替え用のバッファ

	fmt.Printf("input  = %v\n", list)

	for i, _ := range sortedList {
		fmt.Printf("  list(%v) = %v\n", i, sortedList)
		for j := i - 1; j >= 0; j-- {
			fmt.Printf("    list(%v,%v) = %v\n", i, j, sortedList)
			if sortedList[j] > sortedList[j+1] {
				buf = sortedList[j]
				sortedList[j] = sortedList[j+1]
				sortedList[j+1] = buf
			} else {
				break
			}
		}
	}

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}
