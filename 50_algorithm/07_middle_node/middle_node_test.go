package _07_middle_node

import (
	. "github.com/KINGMJ/go-algorithm/structure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("链表的中间结点", func() {
	It("用例1", func() {
		linkedList := Array2List([]interface{}{1, 2, 3, 4, 5})
		middleNode := MiddleNode(linkedList.Head.Next)
		Expect(Node2Array(middleNode)).To(Equal([]interface{}{3, 4, 5}))
	})

	It("用例2", func() {
		linkedList := Array2List([]interface{}{1, 2, 3, 4, 5, 6})
		middleNode := MiddleNode(linkedList.Head.Next)
		Expect(Node2Array(middleNode)).To(Equal([]interface{}{4, 5, 6}))
	})
})
