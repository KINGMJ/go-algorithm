package structure

import (
	"fmt"
)

// SqStack 顺序栈，数组实现
type SqStack struct {
	Data []any
	Top  int //栈顶指针
}

type SqStackInterface interface {
	// 栈是否为空
	//  @return bool
	IsEmpty() bool

	// 进栈操作
	//  @param any
	Push(any)

	// 出栈操作
	//  @return any
	Pop() any

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
	return &SqStack{Data: make([]any, capacity), Top: -1}
}

func (s *SqStack) IsEmpty() bool {
	return s.Top == -1
}

func (s *SqStack) Push(element any) {
	if s.Top == cap(s.Data)-1 {
		panic("栈已满")
	}
	s.Top++
	s.Data[s.Top] = element
}

func (s *SqStack) Pop() any {
	if s.IsEmpty() {
		panic("栈为空")
	}
	element := s.Data[s.Top]
	s.Data[s.Top] = nil
	s.Top--
	return element
}

func (s *SqStack) Print() {
	if s.IsEmpty() {
		fmt.Println("栈为空")
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.Data[i])
	}
}

func (s *SqStack) Flush() {
	s.Top = -1
}
