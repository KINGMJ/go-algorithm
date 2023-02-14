package _06_merge_sorted_linked_list

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 合并链表：使用迭代，双指针方式
//
//	@param list1
//	@param list2
//	@return *ListNode
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 设置一个哨兵节点用于存放合并后的链表
	phead := &ListNode{Data: nil, Next: nil}
	// 维护一个 prev 指针，每次循环调整它的 next
	prev := phead
	// 通过移动 list1 和 list2 的指针来进行比较
	for list1 != nil && list2 != nil {
		if list1.Data.(int) <= list2.Data.(int) {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	// 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return phead.Next
}

// 合并链表：使用递归
//
//	@param list1
//	@param list2
//	@return *ListNode
func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 递归终止条件
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Data.(int) < list2.Data.(int) {
		list1.Next = MergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists2(list2.Next, list1)
		return list2
	}
}
