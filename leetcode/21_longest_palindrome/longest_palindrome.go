package longest_palindrome

// longestPalindrome 计算可以用字符串中的字符构造的最长回文串长度
// s 由大写和小写字母组成
func longestPalindrome(s string) int {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	length := 0
	hasOdd := false
	for _, v := range count {
		// 偶数次直接加入
		length += v / 2 * 2
		// 奇数次，如果之前没有奇数次，则加入1个，否则不加入
		if v%2 == 1 {
			hasOdd = true
		}
	}
	if hasOdd {
		length++
	}
	return length
}
