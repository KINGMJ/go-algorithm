package _04_singly_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test04LinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "04LinkedList Suite")
}
