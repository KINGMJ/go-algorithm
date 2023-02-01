package sort

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("选择排序", func() {
	It("选择排序，从小到大排列", func() {
		Expect(SelectionSort([]int{6, 1, 4, 9, 5, 8})).To(Equal([]int{1, 4, 5, 6, 8, 9}))
	})
})

var _ = Describe("快速排序", func() {
	It("数组只有一个元素", func() {
		Expect(QuickSort([]int{1})).To(Equal([]int{1}))
	})
	It("大于一个元素", func() {
		Expect(QuickSort([]int{6, 1, 4, 9, 5, 8})).To(Equal([]int{1, 4, 5, 6, 8, 9}))
	})
})
