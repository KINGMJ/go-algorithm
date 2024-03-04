package n_queens_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test04NQueens(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "04NQueens Suite")
}
