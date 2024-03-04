package n_queens

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NQueens", func() {
	It("test1", func() {
		expected := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
		Expect(SolveNQueens(4)).To(Equal(expected))
	})

	It("test2", func() {
		expected := [][]string{{"Q"}}
		Expect(SolveNQueens(1)).To(Equal(expected))
	})
})
