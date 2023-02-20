package _12_queue_on_list

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("队列的链式存储结构", func() {
	listQueue := NewListQueue()
	// 入队列
	listQueue.EnQueue("a")
	listQueue.EnQueue("b")
	listQueue.Print()
	// 出队列
	fmt.Println(listQueue.DeQueue())
	fmt.Println(listQueue.DeQueue())

	listQueue.EnQueue("c")
	listQueue.Print()

})
