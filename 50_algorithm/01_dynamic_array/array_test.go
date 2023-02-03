package _01_dynamic_array

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("支持动态扩容的数组", func() {
	It("test", func() {
		arr := NewArray(4)
		fmt.Println(arr)
		// push
		for i := 0; i < 10; i++ {
			arr.Push(i)
		}
		fmt.Println(arr)
		// 清空数组
		arr.Clear()
		fmt.Println(arr)

		// unshift
		for _, char := range "hello, world" {
			arr.Unshift(char)
		}
		fmt.Println(arr)

		// remove
		arr.Remove(1)
		fmt.Println(arr)

		arr.Clear()

		for i := 0; i < 10; i++ {
			arr.Push(i)
		}
		fmt.Println(arr)
		index := arr.FindIndex(func(val interface{}) bool {
			return val == 5
		})
		fmt.Println(index)

		arr.Set(0, "jerry")
		fmt.Println(arr)
	})
})
