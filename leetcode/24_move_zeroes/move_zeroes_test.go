package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "基本测试用例",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "空数组",
			nums: []int{},
			want: []int{},
		},
		{
			name: "单个元素-零",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "单个元素-非零",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "全是零",
			nums: []int{0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "没有零",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "零在开头",
			nums: []int{0, 0, 1, 2, 3},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "零在结尾",
			nums: []int{1, 2, 3, 0, 0},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "零在中间",
			nums: []int{1, 0, 0, 2, 3},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "包含负数",
			nums: []int{-1, 0, 0, -2, 3},
			want: []int{-1, -2, 3, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
