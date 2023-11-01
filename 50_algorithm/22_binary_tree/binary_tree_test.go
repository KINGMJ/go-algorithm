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

	It("[1,2,3,4,5,6,nil,nil,7]", func() {
		tree := Array2Tree([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, 7})
		Expect(PreOrderTraversal(tree)).To(Equal([]interface{}{1, 2, 4, 7, 5, 3, 6}))
	})
})

var _ = Describe("前序遍历，栈实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(PreOrderTraversal1(tree)).To(Equal([]interface{}{1, 2, 3}))
	})

	It("[1,2,3,4,5,6,nil,nil,7]", func() {
		tree := Array2Tree([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, 7})
		Expect(PreOrderTraversal1(tree)).To(Equal([]interface{}{1, 2, 4, 7, 5, 3, 6}))
	})
})

var _ = Describe("中序遍历，递归实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(InOrderTraversal(tree)).To(Equal([]interface{}{1, 3, 2}))
	})

	It("[1, nil, 2]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2})
		Expect(InOrderTraversal(tree)).To(Equal([]interface{}{1, 2}))
	})
})

var _ = Describe("中序遍历，栈实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(InOrderTraversal1(tree)).To(Equal([]interface{}{1, 3, 2}))
	})

	It("[1, nil, 2]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2})
		Expect(InOrderTraversal1(tree)).To(Equal([]interface{}{1, 2}))
	})
})

var _ = Describe("后续遍历，递归实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(PostOrderTraversal(tree)).To(Equal([]interface{}{3, 2, 1}))
	})

	It("[1, nil, 2]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2})
		Expect(PostOrderTraversal(tree)).To(Equal([]interface{}{2, 1}))
	})
})

var _ = Describe("后续遍历，栈实现", func() {
	It("[1, nil, 2, 3]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2, 3})
		Expect(PostOrderTraversal1(tree)).To(Equal([]interface{}{3, 2, 1}))
	})

	It("[1, nil, 2]", func() {
		tree := Array2Tree([]interface{}{1, nil, 2})
		Expect(PostOrderTraversal1(tree)).To(Equal([]interface{}{2, 1}))
	})
})
