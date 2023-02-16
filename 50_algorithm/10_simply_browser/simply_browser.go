package _10_simply_browser

import (
	"fmt"

	. "github.com/KINGMJ/go-algorithm/structure"
)

type Browser struct {
	forwardStack *SqStack // 前进记录栈
	backStack    *SqStack // 后退记录栈
}

// 模拟浏览器的行为
type BrowserInterface interface {
	// 打开一个网址
	//  @param string
	Open(string)

	// 前进
	Forward()

	// 后退
	Back()
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewSqStack(128),
		backStack:    NewSqStack(128),
	}
}

func (this *Browser) Open(addr string) {
	fmt.Printf("访问： %v\n", addr)
	// 把当前正在访问的地址放进后退栈
	this.backStack.Push(addr)
	// 访问一个新的地址，前进栈就要刷新为空栈
	this.forwardStack.Flush()
}

func (this *Browser) Back() {
	// 判断是否可以后退
	// 因为我们把当前访问的addr也存放到了backStack，所以这里应该判断的不是空栈
	if this.backStack.Top == 0 {
		return
	}
	// 当前的addr出栈，存放到forwardStack里面
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	// 栈顶元素为我们回退访问的addr
	top = this.backStack.Data[this.backStack.Top]
	fmt.Printf("回退到：%v\n", top)
}

func (this *Browser) Forward() {
	// 判断是否可以前进
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("前进到：%v\n", top)
}
