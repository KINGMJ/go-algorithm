package _04_doubly_linked_list

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("DoublyLinkedList", func() {
	It("双向链表测试", func() {
		linkedList := NewDoublyLinkedList()
		linkedList.Append("a")
		linkedList.Append("b")
		linkedList.Append("c")
		linkedList.Print()

		// 插入测试
		linkedList.Insert(3, "d")
		linkedList.Print()

		// 测试前驱和后继结点
		node := linkedList.Get(2)
		fmt.Print(node.data)
		fmt.Print(node.next.data)
		fmt.Print(node.prev.data)

		// 删除测试
		linkedList.Delete(3)
		linkedList.Print()
		node = linkedList.Get(2)
		fmt.Print(node.data)
		fmt.Print(node.next.data)
		fmt.Print(node.prev.data)
	})
})
