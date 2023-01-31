package binary_search_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBinarySearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinarySearch Suite")
}
