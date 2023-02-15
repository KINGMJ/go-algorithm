package _08_stack_on_array

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("使用数组模拟栈", func() {
	sqStack := NewSqStack(64)
	// 空栈pop报错
	// sqStack.Pop()

	// 进栈操作
	sqStack.Push(1)
	sqStack.Push(2)
	sqStack.Push(3)
	sqStack.Print()

	element := sqStack.Pop()
	fmt.Printf("出栈的元素：%d\n", element)
	sqStack.Print()
})
