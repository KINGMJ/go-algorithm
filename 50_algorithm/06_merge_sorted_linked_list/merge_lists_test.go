package _06_merge_sorted_linked_list

import (
	. "github.com/KINGMJ/go-algorithm/structure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("合并两个有序链表，双指针解法", func() {
	It("用例1", func() {
		linkedList1 := Array2List([]interface{}{1, 2, 4})
		linkedList2 := Array2List([]interface{}{1, 3, 4})
		mergedList := MergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{1, 1, 2, 3, 4, 4}))
	})

	It("用例2", func() {
		linkedList1 := Array2List([]interface{}{})
		linkedList2 := Array2List([]interface{}{})
		mergedList := MergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{}))
	})

	It("用例3", func() {
		linkedList1 := Array2List([]interface{}{})
		linkedList2 := Array2List([]interface{}{0})
		mergedList := MergeTwoLists(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{0}))
	})
})

var _ = Describe("合并两个有序链表，递归解法", func() {
	It("用例1", func() {
		linkedList1 := Array2List([]interface{}{1, 2, 4})
		linkedList2 := Array2List([]interface{}{1, 3, 4})
		mergedList := MergeTwoLists2(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{1, 1, 2, 3, 4, 4}))
	})

	It("用例2", func() {
		linkedList1 := Array2List([]interface{}{})
		linkedList2 := Array2List([]interface{}{})
		mergedList := MergeTwoLists2(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{}))
	})

	It("用例3", func() {
		linkedList1 := Array2List([]interface{}{})
		linkedList2 := Array2List([]interface{}{0})
		mergedList := MergeTwoLists2(linkedList1.Head.Next, linkedList2.Head.Next)
		Expect(Node2Array(mergedList)).To(Equal([]interface{}{0}))
	})
})
