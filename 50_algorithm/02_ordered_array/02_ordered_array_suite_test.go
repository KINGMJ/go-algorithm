package _02_ordered_array_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test02OrderedArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "02OrderedArray Suite")
}
