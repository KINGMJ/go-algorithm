package _12_queue_on_list

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
	next *Node
}

type ListQueue struct {
	front *Node // 头指针
	rear  *Node // 尾指针
}

type ListQueueInterface interface {
	// 进队列
	//  @param interface{}
	EnQueue(interface{})

	// 出队列
	//  @return interface{}
	DeQueue() interface{}
}

// 空队列时，头指针和尾指针都指向头结点
//
//	@return *ListQueue
func NewListQueue() *ListQueue {
	listQueue := &ListQueue{front: &Node{data: nil, next: nil}}
	listQueue.rear = listQueue.front
	return listQueue
}

// 进队列，链表尾部插入
//
//	@receiver this
//	@param element
func (this *ListQueue) EnQueue(element interface{}) {
	this.rear.next = &Node{data: element, next: nil}
	this.rear = this.rear.next
}

func (this *ListQueue) DeQueue() interface{} {
	// 队列为空判断
	if this.front == this.rear {
		return nil
	}
	// 获取队头元素
	p := this.front.next
	this.front.next = p.next
	// 如果出队列的元素刚好是队列最后一个元素，需要设置rear指针
	if this.rear == p {
		this.rear = this.front
	}
	return p.data
}

// 打印链表结构
//
//	@receiver this
func (this *ListQueue) Print() {
	current := this.front.next
	str := "链表："
	for current != nil {
		str += fmt.Sprintf("%v -> ", current.data)
		current = current.next
	}
	str = strings.TrimRight(str, "-> ")
	fmt.Println(str)
}
