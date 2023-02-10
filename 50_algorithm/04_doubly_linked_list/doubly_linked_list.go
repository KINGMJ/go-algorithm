package _04_doubly_linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{} // 数据域
	next *Node       // 后继指针
	prev *Node       // 前驱指针
}

// 双向链表
type DoublyLinkedList struct {
	head   *Node
	length int
}

type DoublyLinkedListInterface interface {
	Append(interface{})
	Insert(int, interface{})
	Delete(int) interface{}
	Print()
	Get(int) *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: &Node{data: nil, next: nil, prev: nil}, length: 0}
}

func (this *DoublyLinkedList) Append(element interface{}) {
	current := this.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{data: element, next: nil, prev: current}
	this.length++
}

func (this *DoublyLinkedList) Insert(index int, element interface{}) {
	j := 1
	current := this.head
	for current != nil && j < index {
		current = current.next
		j++
	}
	if current == nil || j > index {
		panic("index out of range")
	}
	current.next = &Node{data: element, next: current.next, prev: current}
	this.length++
}

func (this *DoublyLinkedList) Get(index int) *Node {
	j := 1
	current := this.head.next
	for current != nil && j < index {
		current = current.next
		j++
	}
	// 越界处理
	if current == nil || j > index {
		panic("index out of range")
	}
	return current
}

func (this *DoublyLinkedList) Delete(index int) interface{} {
	// 越界检查
	if index < 1 || index > this.length {
		panic("index out of range")
	}
	j := 1
	current := this.head
	for current != nil && j < index {
		current = current.next
		j++
	}
	node := current.next
	current.next = node.next
	node.next.prev = current
	this.length--
	return node.data

}

// 打印链表结构
//
//	@receiver this
func (this *DoublyLinkedList) Print() {
	current := this.head.next
	str := "链表："
	for current != nil {
		str += fmt.Sprintf("%v -> ", current.data)
		current = current.next
	}
	str = strings.TrimRight(str, "-> ")
	str += fmt.Sprintf("，长度：%d", this.length)
	fmt.Println(str)
}
