package reverse_string

func reverseString(s []byte) {
	f1, f2 := 0, len(s)-1
	for f1 < f2 {
		s[f1], s[f2] = s[f2], s[f1]
		f1++
		f2--
	}
}
