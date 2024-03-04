package contains_duplicate_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test02ContainsDuplicate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "02ContainsDuplicate Suite")
}
