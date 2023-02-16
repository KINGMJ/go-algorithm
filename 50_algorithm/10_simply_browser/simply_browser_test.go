package _10_simply_browser

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("SimplyBrowser", func() {
	browser := NewBrowser()
	// 访问操作
	browser.Open("www.baidu.com")
	browser.Open("www.qq.com")
	browser.Open("www.google.com")

	// 回退
	browser.Back()
	browser.Back()

	// 访问新的地址
	browser.Open("www.facebook.com")

	fmt.Println("==========backStack============")
	browser.backStack.Print()
	fmt.Println("=======forwardStack============")
	browser.forwardStack.Print()
})
