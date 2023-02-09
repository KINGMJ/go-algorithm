package _04_linked_list

import (
	"fmt"
	"math/rand"
	"strings"
)

// 单链表结点
type Node struct {
	data interface{} // 数据域
	next *Node       // 指针域
}

// 单链表数据结构
type LinkedList struct {
	head   *Node // 头结点，一般数据域为空或存放链表的长度
	length int   // 链表长度
}

type LinkedListInterface interface {
	// 插入操作，在结点之前插入
	//  @param int 第几个结点
	//  @param interface{} 插入的元素
	Insert(int, interface{})

	// 追加操作
	//  @param interface{}
	Append(interface{})

	// 删除操作
	//  @param int
	//  @return interface{}
	Delete(int) interface{}

	// 获取元素
	//  @param int
	//  @return interface{}
	GetElem(int) interface{}

	// 清空链表
	Clear()

	// 转换为数组
	ToArray() []interface{}

	// 打印链表
	Print()
}

// 模拟构造函数，构造一个单链表
//
//	@return *LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{head: &Node{data: nil, next: nil}, length: 0}
}

// 头插法实现链表整表创建
//
//	@return *LinkedList
func CreateHead(nums int) *LinkedList {
	linkedList := NewLinkedList()
	for i := 0; i < nums; i++ {
		node := &Node{data: rand.Intn(100), next: nil}
		//核心逻辑，新创建的节点指向this.head.next，也就是null
		//this.head.next指向该节点，也就是第一个节点
		//每次创建的都是第一个节点
		node.next = linkedList.head.next
		linkedList.head.next = node
		linkedList.length++
	}
	return linkedList
}

// 尾插法实现链表整表创建
//
//	@return *LinkedList
func CreateTail(nums int) *LinkedList {
	linkedList := NewLinkedList()
	current := linkedList.head
	for i := 0; i < nums; i++ {
		//尾插法比较容易理解，就是append
		current.next = &Node{data: rand.Intn(100), next: nil}
		current = current.next
		linkedList.length++
	}
	return linkedList
}

func (this *LinkedList) Insert(index int, element interface{}) {
	j := 1
	current := this.head // 当前的结点
	//找到第 index-1 个节点：因为有头节点的缘故，找到的是 index 的前一个节点
	for current != nil && j < index {
		current = current.next
		j++
	}
	// 没有找到报错，
	// j>index 的逻辑：比如，index传入0等越界的情况判断
	if current == nil || j > index {
		panic("index out of range")
	}
	// 声明一个要插入的节点
	node := &Node{data: element, next: nil}
	// 插入核心逻辑
	node.next = current.next
	current.next = node
	// 链表长度+1
	this.length++
}

func (this *LinkedList) Append(element interface{}) {
	// 找到最后一个结点
	current := this.head
	for current.next != nil {
		current = current.next
	}
	node := &Node{data: element, next: nil}
	current.next = node
	this.length++
}

func (this *LinkedList) Delete(index int) interface{} {
	//如果链表存储了长度，可以借用该属性处理边界
	if index < 1 || index > this.length {
		panic("index out of range")
	}
	// 找到 index 的前一个结点
	j := 1
	current := this.head
	for current != nil && j < index {
		current = current.next
		j++
	}
	// 删除链表核心逻辑
	node := current.next
	current.next = node.next
	this.length--
	return node.data
}

func (this *LinkedList) GetElem(index int) interface{} {
	//这里跟前面的插入和删除不一样，是得到 index 处的节点，所以 p 为 第一个节点
	j := 1
	current := this.head.next
	for current != nil && j < index {
		current = current.next
		j++
	}
	// 处理越界
	if current == nil || j > index {
		panic("index out of range")
	}
	return current.data
}

func (this *LinkedList) Clear() {
	current := this.head.next
	// 为什么会需要q，对于C这种没有垃圾回收的语言，主动free变量会直接将它从内存中删除
	var q *Node
	for current != nil {
		q = current.next
		// 释放 current free(current)
		current = q
		this.length--
	}
	this.head.next = current
}

// ToArray 使用递归 + 闭包
//
//	@receiver this
//	@return []interface{}
func (this *LinkedList) ToArray() []interface{} {
	// 需要先声明一个函数，不然闭包里面获取不了
	var node2Array func(*Node) []interface{}
	node2Array = func(node *Node) []interface{} {
		if node == nil {
			return []interface{}{}
		}
		//使用递归，先得到当前节点的数据
		current := []interface{}{node.data}
		//然后得到剩余部分的数据
		rest := node2Array(node.next)
		//合并结果
		return append(current, rest...)
	}
	return node2Array(this.head.next)
}

// 打印链表结构
//
//	@receiver this
func (this *LinkedList) Print() {
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
