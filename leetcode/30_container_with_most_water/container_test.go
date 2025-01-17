package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "基本测试用例1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "基本测试用例2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "递增数组",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "递减数组",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "相同高度",
			height: []int{4, 4, 4, 4},
			want:   12,
		},
		{
			name:   "较大数值",
			height: []int{10000, 1, 1, 1, 10000},
			want:   40000,
		},
		{
			name:   "中间高两边低",
			height: []int{1, 2, 4, 3, 2, 1},
			want:   8,
		},
		{
			name:   "两边高中间低",
			height: []int{4, 3, 2, 1, 2, 3, 4},
			want:   24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			assert.Equal(t, tt.want, got)
		})
	}
}
