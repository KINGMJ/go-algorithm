package _06_merge_sorted_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test06MergeSortedLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "06MergeSortedLinkedList Suite")
}
