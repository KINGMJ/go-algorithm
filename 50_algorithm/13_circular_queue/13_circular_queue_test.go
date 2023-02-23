package _13_circular_queue

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("循环队列测试", func() {
	circularQueue := NewCircularQueue(12)
	// 入队列
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	fmt.Println(circularQueue.data...)

	// 出队列
	element := circularQueue.DeQueue()
	fmt.Println(element)
	fmt.Println(circularQueue.data...)
	fmt.Println(circularQueue.getLength())
})
