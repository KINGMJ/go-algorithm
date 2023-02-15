package _09_stack_on_list

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type ListStack struct {
	top *Node // 栈顶节点
}

type ListStackInterface interface {
	// 栈是否为空
	//  @return bool
	IsEmpty() bool

	// 进栈操作
	//  @param interface{}
	Push(interface{})

	// 出栈操作
	//  @return interface{}
	Pop() interface{}

	// 打印栈
	Print()
}

// 初始化一个空栈，栈顶指针为nil
//
//	@return *ListStack
func NewListStack() *ListStack {
	return &ListStack{top: nil}
}

func (this *ListStack) IsEmpty() bool {
	return this.top == nil
}

// 入栈操作 ，链栈不存在栈满的情况，如果内存不足
//
//	@receiver this
//	@param interface{}
func (this *ListStack) Push(element interface{}) {
	//生成一个新的节点，把当前栈顶元素赋值给新结点的直接后继
	this.top = &Node{data: element, next: this.top}
}

func (this *ListStack) Pop() interface{} {
	if this.IsEmpty() {
		panic("empty stack")
	}
	element := this.top.data
	this.top = this.top.next
	return element
}

func (this *ListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.top
		for cur != nil {
			fmt.Println(cur.data)
			cur = cur.next
		}
	}
}
