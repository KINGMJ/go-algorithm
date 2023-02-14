package _07_middle_node

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 链表的中间结点，使用快慢指针
//
//	@param head
//	@return *ListNode
func MiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
