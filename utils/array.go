package utils

// Filter 过滤数组
//
//	@param arr
//	@param cond
//	@return []int
func Filter(arr []int, cond func(int) bool) []int {
	var res []int
	for _, el := range arr {
		if cond(el) {
			res = append(res, el)
		}
	}
	return res
}
