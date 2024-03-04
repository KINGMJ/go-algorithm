package product_of_array_except_self

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductOfArrayExceptSelf", func() {
	It("test1", func() {
		Expect(ProductExceptSelf([]int{1, 2, 3, 4})).To(Equal([]int{24, 12, 8, 6}))
	})
	It("test2", func() {
		Expect(ProductExceptSelf([]int{-1, 1, 0, -3, 3})).To(Equal([]int{0, 0, 9, 0, 0}))
	})
})
