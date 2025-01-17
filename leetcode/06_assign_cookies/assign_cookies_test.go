package assign_cookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindContentChildren(t *testing.T) {
	tests := []struct {
		name string
		g    []int // 孩子的胃口值
		s    []int // 饼干的尺寸
		want int   // 期望满足的孩子数量
	}{
		{
			name: "基本测试用例1",
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			name: "基本测试用例2",
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
		{
			name: "空数组测试",
			g:    []int{},
			s:    []int{1, 2, 3},
			want: 0,
		},
		{
			name: "没有饼干",
			g:    []int{1, 2, 3},
			s:    []int{},
			want: 0,
		},
		{
			name: "饼干不能满足任何孩子",
			g:    []int{2, 3, 4},
			s:    []int{1},
			want: 0,
		},
		{
			name: "完全匹配",
			g:    []int{1, 2, 3},
			s:    []int{1, 2, 3},
			want: 3,
		},
		{
			name: "饼干数量多于孩子",
			g:    []int{1, 2},
			s:    []int{1, 2, 3, 4},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindContentChildren(tt.g, tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
