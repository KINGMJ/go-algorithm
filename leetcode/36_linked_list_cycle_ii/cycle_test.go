package linked_list_cycle_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 辅助函数：创建带环链表
func createCycleList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 创建链表
	head := &ListNode{Val: nums[0]}
	cur := head
	var cycleNode *ListNode

	// 记录环的入口节点
	if pos == 0 {
		cycleNode = head
	}

	// 构建链表
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
		if i == pos {
			cycleNode = cur
		}
	}

	// 创建环
	if pos != -1 {
		cur.Next = cycleNode
	}

	return head
}

func Test_detectCycle(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		pos       int // 环的入口位置，-1表示无环
		wantCycle bool
		wantPos   int
	}{
		{
			name:      "有环链表1",
			nums:      []int{3, 2, 0, -4},
			pos:       1,
			wantCycle: true,
			wantPos:   1,
		},
		{
			name:      "有环链表2",
			nums:      []int{1, 2},
			pos:       0,
			wantCycle: true,
			wantPos:   0,
		},
		{
			name:      "无环链表",
			nums:      []int{1},
			pos:       -1,
			wantCycle: false,
			wantPos:   -1,
		},
		{
			name:      "空链表",
			nums:      []int{},
			pos:       -1,
			wantCycle: false,
			wantPos:   -1,
		},
		{
			name:      "环在尾部",
			nums:      []int{1, 2, 3, 4},
			pos:       3,
			wantCycle: true,
			wantPos:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.nums, tt.pos)
			got := detectCycle(head)

			if !tt.wantCycle {
				assert.Nil(t, got)
				return
			}

			// 验证返回的节点是否是环的入口
			cur := head
			pos := 0
			for cur != got {
				cur = cur.Next
				pos++
			}
			assert.Equal(t, tt.wantPos, pos)
		})
	}
}
