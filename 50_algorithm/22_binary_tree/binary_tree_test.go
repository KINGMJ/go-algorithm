package _22_binary_tree

import (
	. "github.com/KINGMJ/go-algorithm/structure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("前序遍历，递归实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(PreOrderTraversal(tree)).To(Equal([]interface{}{1, 2, 3}))
	})

	It("[1,2,3,4,5,6,null,null,7]", func() {
		tree := Array2Tree([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, 7})
		Expect(PreOrderTraversal(tree)).To(Equal([]interface{}{1, 2, 4, 7, 5, 3, 6}))
	})
})

var _ = Describe("前序遍历，栈实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(PreOrderTraversal1(tree)).To(Equal([]interface{}{1, 2, 3}))
	})

	It("[1,2,3,4,5,6,null,null,7]", func() {
		tree := Array2Tree([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, 7})
		Expect(PreOrderTraversal1(tree)).To(Equal([]interface{}{1, 2, 4, 7, 5, 3, 6}))
	})
})
