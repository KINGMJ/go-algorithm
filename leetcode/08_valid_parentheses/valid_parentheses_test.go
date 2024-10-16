package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"空字符串", "", true},
		{"有效的括号", "()", true},
		{"有效的括号组合", "()[]{}", true},
		{"无效的括号", "(]", false},
		{"无效的括号组合", "([)]", false},
		{"复杂有效括号", "{[]}", true},
		{"无效的括号", "(()", false},
		{"无效的括号", "((", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			assert.Equal(t, tt.expected, result, "对于输入 %s", tt.input)
		})
	}
}
