package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		want     int
		wantNums []int
	}{
		{
			name:     "基本测试用例1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name:     "基本测试用例2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:     "空数组",
			nums:     []int{},
			val:      1,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "数组中没有要移除的元素",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			want:     4,
			wantNums: []int{1, 2, 3, 4},
		},
		{
			name:     "数组中全是要移除的元素",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "只有一个元素且需要移除",
			nums:     []int{1},
			val:      1,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "只有一个元素且不需要移除",
			nums:     []int{2},
			val:      1,
			want:     1,
			wantNums: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalLen := len(tt.nums)
			got := removeElement(tt.nums, tt.val)

			assert.Equal(t, tt.want, got, "返回值不符合预期")
			assert.Equal(t, tt.wantNums, tt.nums[:got], "数组内容不符合预期")

			// 验证原数组长度没有改变
			assert.Equal(t, originalLen, len(tt.nums), "数组长度被改变")
		})
	}
}
