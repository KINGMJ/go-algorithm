package _04_doubly_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test04DoublyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "04DoublyLinkedList Suite")
}
