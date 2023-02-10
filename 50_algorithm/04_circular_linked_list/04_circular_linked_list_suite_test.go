package _04_circular_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test04CircularLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "04CircularLinkedList Suite")
}
