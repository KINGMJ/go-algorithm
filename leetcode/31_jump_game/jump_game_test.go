package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "基本测试用例1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "基本测试用例2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "单个元素",
			nums: []int{0},
			want: true,
		},
		{
			name: "全零数组",
			nums: []int{0, 0, 0},
			want: false,
		},
		{
			name: "刚好能到达",
			nums: []int{2, 0, 0},
			want: true,
		},
		{
			name: "无法跨越零",
			nums: []int{1, 0, 1, 0},
			want: false,
		},
		{
			name: "能跨越零",
			nums: []int{2, 0, 1, 0, 1},
			want: true,
		},
		{
			name: "最大跳跃",
			nums: []int{5, 0, 0, 0, 0, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
