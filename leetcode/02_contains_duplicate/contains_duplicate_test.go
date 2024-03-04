package contains_duplicate

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("02ContainsDuplicate", func() {
	It("test1", func() {
		Expect(ContainsDuplicate([]int{1, 2, 3, 1})).To(Equal(true))
	})

	It("test2", func() {
		Expect(ContainsDuplicate([]int{1, 2, 3, 4})).To(Equal(false))
	})

	It("test3", func() {
		Expect(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})).To(Equal(true))
	})
})
