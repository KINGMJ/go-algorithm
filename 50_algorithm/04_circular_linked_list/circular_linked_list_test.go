package _04_circular_linked_list

import (
	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("CircularLinkedList", func() {
	It("循环链表测试", func() {
		circularLinkedList := NewCircularLinkedList()
		circularLinkedList.Append("a")
		circularLinkedList.Append("b")
		circularLinkedList.Append("c")
		circularLinkedList.Print()

		circularLinkedListB := NewCircularLinkedList()
		circularLinkedListB.Append(1)
		circularLinkedListB.Append(2)

		// 循环链表合并
		Merge(circularLinkedList, circularLinkedListB)
		circularLinkedList.Print()
	})
})
