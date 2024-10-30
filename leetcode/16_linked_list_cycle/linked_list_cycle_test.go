package linkedlistcycle

import (
	"testing"
)

// TestHasCycle 测试 hasCycle 函数
func TestHasCycle(t *testing.T) {
	// 创建测试用例
	tests := []struct {
		head     *ListNode
		expected bool
	}{
		// 示例测试用例
		{&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}, true}, // 有环
		{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, false},                                                   // 无环
		{
			// 创建一个有环的链表
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: true,
		},
	}

	// 创建环
	// 第三个测试用例的链表有环
	tests[2].head.Next.Next.Next = tests[2].head.Next // 使链表有环

	// 第一个测试用例的链表有环
	tests[0].head.Next.Next.Next = tests[0].head.Next // 使链表有环

	for _, test := range tests {
		result := hasCycle(test.head)
		if result != test.expected {
			t.Errorf("hasCycle(%v) = %v; expected %v", test.head, result, test.expected)
		}
	}
}
