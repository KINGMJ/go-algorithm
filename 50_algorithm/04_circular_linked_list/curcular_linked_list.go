package _04_circular_linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
	next *Node
}

// 循环链表
type CircularLinkedList struct {
	rear   *Node // 尾指针，空链表指向头结点
	length int   // 链表长度
}

type CircularLinkedListInterface interface {
	// Append
	//  @param interface{}
	Append(interface{})

	Print()
}

// 模拟构造函数，初始化循环链表
//
//	@return *CircularLinkedList
func NewCircularLinkedList() *CircularLinkedList {
	circularLinkedList := &CircularLinkedList{rear: &Node{data: nil, next: nil}, length: 0}
	//带头结点的空链表，next指向自己
	circularLinkedList.rear.next = circularLinkedList.rear
	return circularLinkedList
}

func (this *CircularLinkedList) Append(element interface{}) {
	// 找到头结点
	head := this.rear.next
	current := head
	for current.next != head {
		current = current.next
	}
	node := &Node{data: element, next: head}
	current.next = node
	this.rear = node
	this.length++
}

// 打印链表结构
//
//	@receiver this
func (this *CircularLinkedList) Print() {
	head := this.rear.next
	current := head
	str := "链表："
	for current.next != head {
		current = current.next
		str += fmt.Sprintf("%v -> ", current.data)
	}
	str = strings.TrimRight(str, "-> ")
	str += fmt.Sprintf("，长度：%d", this.length)
	fmt.Println(str)
}

// 链表合并
//
//	@param listA
//	@param listB
func Merge(listA *CircularLinkedList, listB *CircularLinkedList) {
	// 保存A表的头结点
	headA := listA.rear.next
	//将本是指向B表的第一个节点（不是头结点）赋值给 listA.rear.next
	listA.rear.next = listB.rear.next.next
	//将原A表的头结点赋值给listB.rear.next
	listB.rear.next = headA
	//合并后rear为B的rear
	listA.rear = listB.rear
	// 长度更新
	listA.length = listA.length + listB.length
}
