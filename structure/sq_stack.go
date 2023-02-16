package structure

import (
	"fmt"
)

// SqStack 顺序栈，数组实现
type SqStack struct {
	Data []interface{}
	Top  int //栈顶指针
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

	// 刷新栈
	Flush()
}

// 初始化操作，建立一个空栈，栈顶指针为-1
//
//	@param capacity
//	@return *SqStack
func NewSqStack(capacity int) *SqStack {
	return &SqStack{Data: make([]interface{}, capacity, capacity), Top: -1}
}

func (this *SqStack) IsEmpty() bool {
	return this.Top == -1
}

func (this *SqStack) Push(element interface{}) {
	if this.Top == cap(this.Data)-1 {
		panic("栈已满")
	}
	this.Top++
	this.Data[this.Top] = element
}

func (this *SqStack) Pop() interface{} {
	if this.IsEmpty() {
		panic("栈为空")
	}
	element := this.Data[this.Top]
	this.Data[this.Top] = nil
	this.Top--
	return element
}

func (this *SqStack) Print() {
	if this.IsEmpty() {
		fmt.Println("栈为空")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Println(this.Data[i])
	}
}

func (this *SqStack) Flush() {
	this.Top = -1
}
