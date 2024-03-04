package towsum

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TowSum", func() {
	It("test1", func() {
		Expect(TwoSum([]int{2, 7, 11, 15}, 9)).To(Equal([]int{0, 1}))
	})

	It("test2", func() {
		Expect(TwoSum([]int{3, 2, 4}, 6)).To(Equal([]int{1, 2}))
	})

	It("test3", func() {
		Expect(TwoSum([]int{3, 3}, 6)).To(Equal([]int{0, 1}))
	})
})
