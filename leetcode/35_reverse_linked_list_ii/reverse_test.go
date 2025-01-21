package reverse_linked_list_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 辅助函数：将数组转换为链表
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

// 辅助函数：将链表转换为数组
func listToArray(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		left  int
		right int
		want  []int
	}{
		{
			name:  "基本测试用例",
			nums:  []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2, 5},
		},
		{
			name:  "反转整个链表",
			nums:  []int{1, 2, 3, 4, 5},
			left:  1,
			right: 5,
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "反转单个节点",
			nums:  []int{1, 2, 3},
			left:  2,
			right: 2,
			want:  []int{1, 2, 3},
		},
		{
			name:  "反转头部",
			nums:  []int{1, 2, 3, 4, 5},
			left:  1,
			right: 3,
			want:  []int{3, 2, 1, 4, 5},
		},
		{
			name:  "反转尾部",
			nums:  []int{1, 2, 3, 4, 5},
			left:  3,
			right: 5,
			want:  []int{1, 2, 5, 4, 3},
		},
		{
			name:  "空链表",
			nums:  []int{},
			left:  1,
			right: 1,
			want:  []int{},
		},
		{
			name:  "单节点链表",
			nums:  []int{1},
			left:  1,
			right: 1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.nums)
			got := reverseBetween(head, tt.left, tt.right)
			assert.Equal(t, tt.want, listToArray(got))
		})
	}
}
