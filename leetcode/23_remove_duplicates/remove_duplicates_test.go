package remove_duplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantNums []int
	}{
		{
			name:     "基本测试用例1",
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name:     "基本测试用例2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "空数组",
			nums:     []int{},
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			want:     1,
			wantNums: []int{1},
		},
		{
			name:     "全部相同元素",
			nums:     []int{1, 1, 1, 1, 1},
			want:     1,
			wantNums: []int{1},
		},
		{
			name:     "没有重复元素",
			nums:     []int{1, 2, 3, 4, 5},
			want:     5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "负数测试",
			nums:     []int{-3, -3, -2, -1, -1, 0, 0},
			want:     4,
			wantNums: []int{-3, -2, -1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalNums := make([]int, len(tt.nums))
			copy(originalNums, tt.nums)

			got := removeDuplicates(tt.nums)

			assert.Equal(t, tt.want, got, "返回值不符合预期")
			assert.Equal(t, tt.wantNums, tt.nums[:got], "数组内容不符合预期")

			// 验证原数组后续元素未被修改
			if len(tt.nums) > got {
				assert.Equal(t, originalNums[got:], tt.nums[got:], "数组未使用部分被意外修改")
			}
		})
	}
}
