package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		{
			name:    "基本测试用例1",
			ratings: []int{1, 0, 2},
			want:    5, // [2,1,2]
		},
		{
			name:    "基本测试用例2",
			ratings: []int{1, 2, 2},
			want:    4, // [1,2,1]
		},
		{
			name:    "单个元素",
			ratings: []int{1},
			want:    1,
		},
		{
			name:    "两个相同分数",
			ratings: []int{1, 1},
			want:    2,
		},
		{
			name:    "递增序列",
			ratings: []int{1, 2, 3, 4, 5},
			want:    15, // [1,2,3,4,5]
		},
		{
			name:    "递减序列",
			ratings: []int{5, 4, 3, 2, 1},
			want:    15, // [5,4,3,2,1]
		},
		{
			name:    "山峰形状",
			ratings: []int{1, 3, 5, 3, 1},
			want:    9, // [1,2,3,2,1]
		},
		{
			name:    "谷底形状",
			ratings: []int{5, 3, 1, 3, 5},
			want:    11, // [3,2,1,2,3]
		},
		{
			name:    "相邻分数相等",
			ratings: []int{1, 2, 2, 3},
			want:    6, // [1,2,1,2]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy(tt.ratings)
			assert.Equal(t, tt.want, got)
		})
	}
}
