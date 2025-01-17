package squares_of_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "基本测试用例1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "基本测试用例2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "全是负数",
			nums: []int{-5, -3, -2, -1},
			want: []int{1, 4, 9, 25},
		},
		{
			name: "全是正数",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "只有零",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "包含零",
			nums: []int{-3, -2, 0, 2, 3},
			want: []int{0, 4, 4, 9, 9},
		},
		{
			name: "对称数组",
			nums: []int{-3, -2, -1, 1, 2, 3},
			want: []int{1, 1, 4, 4, 9, 9},
		},
		{
			name: "空数组",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.nums)
			assert.Equal(t, tt.want, got)

			// 验证结果是否有序
			for i := 1; i < len(got); i++ {
				assert.LessOrEqual(t, got[i-1], got[i], "结果数组不是有序的")
			}
		})
	}
}
