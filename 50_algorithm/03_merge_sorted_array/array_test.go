package _03_merge_sorted_array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("合并两个有序数组", func() {
	It("示例1", func() {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		Merge(nums1, 3, nums2, 3)
		Expect(nums1).To(Equal([]int{1, 2, 2, 3, 5, 6}))
	})

	It("示例2", func() {
		nums1 := []int{1}
		nums2 := []int{}
		Merge(nums1, 1, nums2, 0)
		Expect(nums1).To(Equal([]int{1}))
	})

	It("示例3", func() {
		nums1 := []int{0}
		nums2 := []int{1}
		Merge(nums1, 0, nums2, 1)
		Expect(nums1).To(Equal([]int{1}))
	})
})
