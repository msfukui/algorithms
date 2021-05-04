package heap

import "fmt"

// HeapSort は
// 与えられた整数のスライスから、
// ソートの結果とヒープソートで並び替えた結果の整数のスライスを返します。
// c.f. No.2-5 ヒープソート
func HeapSort(list []int) (success bool, sortedList []int) {
	// 返却値用の配列の領域を作成する
	sortedList = make([]int, len(list))

	fmt.Printf("input  = %v\n", list)

	// ヒープを構築する
	heap := []int{}
	for _, v := range list {
		_, heap = HeapAdd(heap, v)
	}

	// ヒープから順番に要素を取得して配列に詰め込む
	r := false
	for i := 0; i < len(list); i++ {
		fmt.Printf("  heap = %v\n", heap)
		sortedList[i] = heap[0]
		r, heap = HeapDel(heap)
		if !r {
			break
		}
	}

	fmt.Printf("output = %v\n", sortedList)

	return true, sortedList
}

// HeapAdd は
// 指定されたヒープに指定された値を追加して、追加した結果と追加した新しいヒープを返します。
// c.f. No.1-7 ヒープ
func HeapAdd(list []int, value int) (success bool, addedList []int) {
	// 返却用のヒープを作ってコピーする
	addedList = make([]int, len(list))
	copy(addedList, list)

	// ヒープ末尾に一旦要素を追加する
	addedList = append(addedList, value)

	i := len(addedList) - 1 // 自身の場所
	j := ((i + 1) / 2) - 1  // 自身の親の場所
	buf := 0
	for j >= 0 {
		// 自身と親を比較して親の方が大きければ置き換えする
		if addedList[j] > addedList[i] {
			buf = addedList[j]
			addedList[j] = addedList[i]
			addedList[i] = buf
			// さらに上の親について繰り返しチェックする
			i = j
			j = ((i + 1) / 2) - 1
		} else {
			// 親の方が小さければ置き換えは終了する
			break
		}
	}

	return true, addedList
}

// HeapDel は
// 指定されたヒープから一番上のノードを削除して、削除した結果と削除後の新しいヒープを返します。
// c.f. No.1-7 ヒープ
func HeapDel(list []int) (success bool, deletedList []int) {
	// 指定されたヒープが空の場合はエラーを返す
	if len(list) == 0 {
		return false, []int{}
	}
	// 指定されたヒープ要素が1つの場合は正常と空要素を返す
	if len(list) == 1 {
		return true, []int{}
	}

	// 先頭の要素を取り除いて末尾の要素を先頭に配置する
	deletedList = append(deletedList, list[len(list)-1])
	deletedList = append(deletedList, list[1:len(list)-1]...)

	i := 0                 // 自身の場所
	j := ((i + 1) * 2) - 1 // 自身の子の場所(左)
	k := 0
	buf := 0

	// 加工後のヒープ要素が2つの場合は
	// 自身と子を比較して子の方が小さければ置き換えをする
	if len(deletedList) == 2 {
		if deletedList[i] > deletedList[j] {
			buf = deletedList[i]
			deletedList[i] = deletedList[j]
			deletedList[j] = buf
		}
	}

	for j+1 < len(deletedList) {
		// 自身と子を比較して子の方が小さければ置き換えをする
		if deletedList[i] > deletedList[j] || deletedList[i] > deletedList[j+1] {
			// より小さい子の方に置き換える
			if deletedList[j] < deletedList[j+1] {
				k = j
			} else {
				k = j + 1
			}
			buf = deletedList[i]
			deletedList[i] = deletedList[k]
			deletedList[k] = buf
			// さらに下の子について繰り返しチェックする
			i = k
			j = ((i + 1) * 2) - 1
		} else {
			// 子の方が大きければ置き換えは終了する
			break
		}
	}

	return true, deletedList
}
