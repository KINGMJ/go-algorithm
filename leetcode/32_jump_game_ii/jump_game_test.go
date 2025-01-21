package jump_game_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "基本测试用例1",
			nums: []int{2, 3, 1, 1, 4},
			want: 2, // 2 -> 3 -> 4
		},
		{
			name: "基本测试用例2",
			nums: []int{2, 3, 0, 1, 4},
			want: 2, // 2 -> 3 -> 4
		},
		{
			name: "单个元素",
			nums: []int{0},
			want: 0,
		},
		{
			name: "最小步数",
			nums: []int{1, 2, 3},
			want: 2, // 1 -> 2 -> 3
		},
		{
			name: "最大跳跃",
			nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			want: 3,
		},
		{
			name: "相邻跳跃",
			nums: []int{1, 1, 1, 1},
			want: 3, // 1 -> 1 -> 1 -> 1
		},
		{
			name: "刚好到达",
			nums: []int{3, 2, 1},
			want: 1, // 直接跳到最后
		},
		{
			name: "需要精确跳跃",
			nums: []int{1, 2, 1, 1, 1},
			want: 3, // 1 -> 2 -> 1 -> 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
