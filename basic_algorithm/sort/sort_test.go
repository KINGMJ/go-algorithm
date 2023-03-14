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

var _ = Describe("冒泡排序", func() {
	It("冒泡排序", func() {
		Expect(BubbleSort([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	})

	It("冒泡排序算法改进", func() {
		Expect(BubbleSort1([]int{9, 1, 5, 8, 3, 7, 4, 6, 2})).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	})
})

var _ = Describe("插入排序", func() {
	It("插入排序", func() {
		Expect(InsertionSort([]int{6, 4, 1, 2, 5})).To(Equal([]int{1, 2, 4, 5, 6}))
	})
})
