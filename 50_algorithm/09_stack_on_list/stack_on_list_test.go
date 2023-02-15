package _09_stack_on_list

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("使用链表实现栈", func() {
	listStack := NewListStack()
	// 空栈pop报错
	// listStack.Pop()

	listStack.Push(1)
	listStack.Push(2)
	listStack.Push(3)
	listStack.Print()

	element := listStack.Pop()
	fmt.Printf("出栈元素: %d\n", element)
	listStack.Print()
})
