package assign_cookies

import "sort"

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	f1, f2 := 0, 0
	for f1 < len(g) && f2 < len(s) {
		if g[f1] <= s[f2] {
			f1++
		}
		f2++
	}
	return f1
}
