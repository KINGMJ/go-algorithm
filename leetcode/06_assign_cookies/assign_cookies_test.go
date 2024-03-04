package assign_cookies

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AssignCookies", func() {
	It("test1", func() {
		g := []int{1, 2, 3}
		s := []int{1, 1}
		Expect(FindContentChildren(g, s)).To(Equal(1))
	})

	It("test2", func() {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		Expect(FindContentChildren(g, s)).To(Equal(2))
	})
})
