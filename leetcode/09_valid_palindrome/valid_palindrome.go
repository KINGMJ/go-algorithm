package validpalindrome

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	// 移出非数字字母元素
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
	// 进行验证
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
