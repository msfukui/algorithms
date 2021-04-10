package liner_search

// LinerSearch は
// 検索対象のスライスと検索したい値から
// 一致する値があれば、一致したことと、そのスライスの位置（添え字）を返します。
// 一致する値がなければ、一致しなかったことを返します。
// c.f. No.3-1 線形探索
func LinerSearch(array []int, value int) (result bool, index int) {
	for i, v := range array {
		if v == value {
			return true, i
		}
	}
	return false, 0
}
