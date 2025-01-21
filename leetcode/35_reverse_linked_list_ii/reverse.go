package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}

// 关键点解释：
// prev 保持不动，是反转区间的前一个节点
// curr 始终是反转区间的第一个节点（原始的2）
// next 是要插入到 prev 后面的节点
// 每次循环都是将 next 节点插入到 prev 的后面
