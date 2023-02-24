package _14_permutations

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

// https://leetcode.cn/problems/permutations/
var _ = Describe("46. 全排列", func() {
	res := Permute([]int{1, 4, 5, 2})
	fmt.Println(res)
})
