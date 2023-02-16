package _10_simply_browser_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test10SimplyBrowser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "10SimplyBrowser Suite")
}
