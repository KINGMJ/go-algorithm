package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "示例 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "单个元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "全部相同元素",
			nums: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "负数元素",
			nums: []int{-1, -1, 2, -1},
			want: -1,
		},
		{
			name: "大数组",
			nums: []int{1, 1, 1, 1, 2, 2, 3, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
