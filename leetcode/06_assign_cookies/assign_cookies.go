package assign_cookies

import "slices"

func FindContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	// 孩子指针和饼干指针
	var gPtr, sPtr int = 0, 0

	for gPtr < len(g) && sPtr < len(s) {
		if g[gPtr] <= s[sPtr] {
			gPtr++
		}
		sPtr++
	}
	return gPtr
}
