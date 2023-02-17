package _11_queue_on_array

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("QueueOnArray", func() {
	It("使用slice实现的队列", func() {
		sqQueue := NewSqQueue(24)
		// 进栈操作
		sqQueue = sqQueue.EnQueue(1)
		sqQueue = sqQueue.EnQueue(2)
		sqQueue = sqQueue.EnQueue(3)
		fmt.Println(sqQueue)

		sqQueue, element := sqQueue.DeQueue()

		fmt.Println(sqQueue)
		fmt.Println(element)
	})

	It("使用结构体实现的队列", func() {
		sqQueue := NewSqQueue2(24)
		sqQueue.EnQueue(1)
		sqQueue.EnQueue(2)
		sqQueue.EnQueue(3)
		fmt.Println(sqQueue)
		element := sqQueue.DeQueue()
		fmt.Println(sqQueue)
		fmt.Println(element)
	})
})
