package _05_reverse_list

import (
	. "github.com/KINGMJ/go-algorithm/structure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("使用双指针逆转链表", func() {
	It("用例1", func() {
		linkedList := Array2List([]interface{}{1, 2, 3, 4, 5})
		head := ReverseList(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{5, 4, 3, 2, 1}))
	})

	It("用例2", func() {
		linkedList := Array2List([]interface{}{1, 2})
		head := ReverseList(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{2, 1}))
	})

	It("用例3", func() {
		linkedList := Array2List([]interface{}{})
		head := ReverseList(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{}))
	})
})

var _ = Describe("使用递归逆转链表", func() {
	It("用例1", func() {
		linkedList := Array2List([]interface{}{1, 2, 3, 4, 5})
		head := ReverseList2(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{5, 4, 3, 2, 1}))
	})

	It("用例2", func() {
		linkedList := Array2List([]interface{}{1, 2})
		head := ReverseList2(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{2, 1}))
	})

	It("用例3", func() {
		linkedList := Array2List([]interface{}{})
		head := ReverseList2(linkedList.Head.Next)
		Expect(Node2Array(head)).To(Equal([]interface{}{}))
	})
})
