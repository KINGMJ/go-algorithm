package _02_ordered_array

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("大小固定的有序数组，支持动态增删改操作", func() {
	It("test", func() {
		// 初始化一个大小固定位10的数组
		arr := NewArray(4)
		fmt.Println(arr)
		// 循环插入一个乱序的数组
		arr2 := []int{4, 1, 6, 2}
		for _, element := range arr2 {
			arr.Insert(element)
		}
		fmt.Println(arr)
		//数组已满，不能插入
		arr.Insert(9)
	})
})
