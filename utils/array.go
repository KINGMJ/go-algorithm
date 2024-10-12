package utils

// Filter 过滤切片
//
//	@param s 输入切片
//	@param f 过滤函数
//	@return []E 过滤后的新切片
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	var r S
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
