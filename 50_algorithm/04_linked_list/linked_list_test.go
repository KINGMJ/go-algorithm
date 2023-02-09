package _04_linked_list

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	It("单链表测试", func() {
		linkedList := NewLinkedList()
		// append 操作
		linkedList.Append("a")
		linkedList.Append("b")
		linkedList.Append("c")
		linkedList.Print()

		// insert 操作
		linkedList.Insert(1, "d")
		linkedList.Print()

		// 删除操作
		data := linkedList.Delete(1)
		fmt.Println(data)
		linkedList.Print()

		// 获取元素
		data = linkedList.GetElem(2)
		fmt.Println(data)

		// 清空链表
		linkedList.Clear()
		linkedList.Print()

		linkedList.Append("x")
		linkedList.Append("y")
		linkedList.Append("z")
		// 链表转换为数组
		arr := linkedList.ToArray()
		fmt.Println(arr)

		// 头插法链表整表创建
		linkedList = CreateHead(10)
		linkedList.Print()

		// 尾插法链表整表创建
		linkedList = CreateTail(10)
		linkedList.Print()
	})
})
