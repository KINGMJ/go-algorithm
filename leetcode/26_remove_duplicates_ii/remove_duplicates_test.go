package remove_duplicates_ii

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
			nums:     []int{1, 1, 1, 2, 2, 3},
			want:     5,
			wantNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "基本测试用例2",
			nums:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want:     7,
			wantNums: []int{0, 0, 1, 1, 2, 3, 3},
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
			name:     "两个相同元素",
			nums:     []int{1, 1},
			want:     2,
			wantNums: []int{1, 1},
		},
		{
			name:     "三个相同元素",
			nums:     []int{1, 1, 1},
			want:     2,
			wantNums: []int{1, 1},
		},
		{
			name:     "全部不同元素",
			nums:     []int{1, 2, 3, 4, 5},
			want:     5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "负数测试",
			nums:     []int{-1, -1, -1, 0, 0, 0, 1, 1, 1},
			want:     6,
			wantNums: []int{-1, -1, 0, 0, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, got, "返回值不符合预期")
			assert.Equal(t, tt.wantNums, tt.nums[:got], "数组内容不符合预期")
		})
	}
}
