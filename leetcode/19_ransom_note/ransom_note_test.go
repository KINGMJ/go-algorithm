package ransom_note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	t.Run("基本测试用例", func(t *testing.T) {
		// 测试用例1: 单个字符不匹配
		assert.False(t, canConstruct("a", "b"))

		// 测试用例2: magazine中字符不足
		assert.False(t, canConstruct("aa", "ab"))

		// 测试用例3: 可以构造
		assert.True(t, canConstruct("aa", "aab"))
	})

	t.Run("空字符串测试", func(t *testing.T) {
		// 空字符串测试
		assert.True(t, canConstruct("", ""))
		assert.True(t, canConstruct("", "abc"))
		assert.False(t, canConstruct("a", ""))
	})

	t.Run("复杂测试用例", func(t *testing.T) {
		// 复杂字符串测试
		assert.False(t, canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"))
		assert.True(t, canConstruct("hello", "hello world"))
		assert.False(t, canConstruct("world", "hello"))
	})

	t.Run("重复字符测试", func(t *testing.T) {
		// 重复字符测试
		assert.True(t, canConstruct("aaa", "aaa"))
		assert.False(t, canConstruct("aaa", "aa"))
		assert.True(t, canConstruct("abc", "aabbcc"))
	})
}
