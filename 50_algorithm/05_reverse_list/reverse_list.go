package _05_reverse_list

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// ReverseList 双指针方式
//
//	@param head leetcode 这里的 head 指的是第一个结点，不是头结点，有点误导
//	@return *ListNode
func ReverseList(head *ListNode) *ListNode {
	// 定义两个指针
	var prev *ListNode     // 上一个结点
	curr := head           // 当前结点
	var nextTemp *ListNode // 下一个结点的值
	for curr != nil {
		//存储下一个节点
		nextTemp = curr.Next
		//把当前节点的next设置为prev
		curr.Next = prev
		//当前节点变成上一个节点
		prev = curr
		//下一个节点变成当前节点
		curr = nextTemp
	}
	// 最后一次循环时，因为 curr 变成了 nil，所以最终返回的结点应该是prev
	return prev
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
