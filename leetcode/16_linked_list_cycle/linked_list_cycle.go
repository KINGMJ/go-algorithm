package linkedlistcycle

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 函数定义：判断链表是否有环
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
