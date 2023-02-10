package _05_reverse_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test05ReverseList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "05ReverseList Suite")
}
