package longest_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestPalindrome(t *testing.T) {
	t.Run("基本测试用例", func(t *testing.T) {
		assert.Equal(t, 7, longestPalindrome("abccccdd")) // 可以构造回文串 "dccaccd"
		assert.Equal(t, 1, longestPalindrome("a"))        // 单个字符
		assert.Equal(t, 2, longestPalindrome("bb"))       // 两个相同字符
		assert.Equal(t, 3, longestPalindrome("ccc"))      // 三个相同字符
	})

	t.Run("大小写混合测试", func(t *testing.T) {
		assert.Equal(t, 1, longestPalindrome("Aa"))       // 大小写不同，只能用其中一个
		assert.Equal(t, 7, longestPalindrome("aAbbccBB")) // 可以构造如 "BabccbaB"
	})

	t.Run("边界测试", func(t *testing.T) {
		assert.Equal(t, 0, longestPalindrome(""))       // 空字符串
		assert.Equal(t, 6, longestPalindrome("AAAAAA")) // 全部相同字符
		assert.Equal(t, 1, longestPalindrome("abcdef")) // 全部不同字符
	})
}
