package _08_stack_on_array

import (
	"fmt"
)

// SqStack 顺序栈，数组实现
type SqStack struct {
	data []interface{}
	top  int //栈顶指针
}

type SqStackInterface interface {
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

// 初始化操作，建立一个空栈，栈顶指针为-1
//
//	@param capacity
//	@return *SqStack
func NewSqStack(capacity int) *SqStack {
	return &SqStack{data: make([]interface{}, capacity, capacity), top: -1}
}

func (this *SqStack) IsEmpty() bool {
	return this.top == -1
}

func (this *SqStack) Push(element interface{}) {
	if this.top == cap(this.data)-1 {
		panic("栈已满")
	}
	this.top++
	this.data[this.top] = element
}

func (this *SqStack) Pop() interface{} {
	if this.IsEmpty() {
		panic("栈为空")
	}
	element := this.data[this.top]
	this.data[this.top] = nil
	this.top--
	return element
}

func (this *SqStack) Print() {
	if this.IsEmpty() {
		fmt.Println("栈为空")
	}
	for i := this.top; i >= 0; i-- {
		fmt.Println(this.data[i])
	}
}
