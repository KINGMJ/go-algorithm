package climbingstairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "2阶楼梯",
			input:    2,
			expected: 2,
		},
		{
			name:     "3阶楼梯",
			input:    3,
			expected: 3,
		},
		{
			name:     "4阶楼梯",
			input:    4,
			expected: 5,
		},
		{
			name:     "5阶楼梯",
			input:    5,
			expected: 8,
		},
		{
			name:     "边界条件: 1阶楼梯",
			input:    1,
			expected: 1,
		},
		{
			name:     "较大数值测试: 10阶楼梯",
			input:    10,
			expected: 89,
		},
		{
			name:     "较大数值测试: 20阶楼梯",
			input:    20,
			expected: 10946,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.input); got != tt.expected {
				t.Errorf("climbStairs() = %v, want %v", got, tt.expected)
			}
		})
	}
}
