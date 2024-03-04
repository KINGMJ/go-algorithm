package permutations

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Permutations", func() {
	It("test1", func() {
		expected := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}
		Expect(Permute([]int{1, 2, 3})).To(Equal(expected))
	})

	It("test2", func() {
		expected := [][]int{
			{0, 1},
			{1, 0},
		}
		Expect(Permute([]int{0, 1})).To(Equal(expected))
	})

	It("test3", func() {
		Expect(Permute([]int{1})).To(Equal([][]int{{1}}))
	})
})
