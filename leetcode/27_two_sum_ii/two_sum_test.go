package two_sum_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "基本测试用例1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "基本测试用例2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "基本测试用例3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "两个相同的数",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "负数测试",
			numbers: []int{-3, -2, -1, 0},
			target:  -5,
			want:    []int{1, 2},
		},
		{
			name:    "较大数测试",
			numbers: []int{2, 3, 7, 11, 15},
			target:  26,
			want:    []int{4, 5},
		},
		{
			name:    "最小长度测试",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
