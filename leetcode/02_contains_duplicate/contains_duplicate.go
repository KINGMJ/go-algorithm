package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	m1 := make(map[int]int)
	for _, v := range nums {
		if _, ok := m1[v]; ok {
			return true
		}
		m1[v] = 0
	}
	return false
}
