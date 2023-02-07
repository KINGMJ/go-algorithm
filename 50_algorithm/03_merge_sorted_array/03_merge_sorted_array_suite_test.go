package _03_merge_sorted_array_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test03MergeSortedArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "03MergeSortedArray Suite")
}
