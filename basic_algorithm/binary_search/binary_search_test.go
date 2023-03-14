package binary_search

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearch", func() {
	It("查找的元素不在数组中", func() {
		position := BinarySearch([]int{1, 3, 5, 7, 9}, 4)
		Expect(position).To(Equal(-1))
	})

	It("查找的元素在数组中", func() {
		position := BinarySearch([]int{1, 3, 5, 7, 9}, 3)
		Expect(position).To(Equal(1))
	})
})
