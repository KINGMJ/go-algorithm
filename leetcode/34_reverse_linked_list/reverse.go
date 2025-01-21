package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var prev, nextTemp *ListNode
	curr := head
	for curr != nil {
		nextTemp = curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
