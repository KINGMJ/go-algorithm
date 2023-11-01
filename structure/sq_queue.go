package structure

import "fmt"

// 顺序队列实现：数组本身就可以看做是一个顺序队列
type SqQueue []interface{}

type SqQueueInterface interface {
	// 进队列
	//  @param interface{}
	EnQueue(interface{}) SqQueue

	// 出队列
	//  @return SqQueue
	//  @return interface{}
	DeQueue() (SqQueue, interface{})
}

func NewSqQueue(capacity int) SqQueue {
	return SqQueue(make([]interface{}, 0, capacity))
}

func (this SqQueue) EnQueue(element interface{}) SqQueue {
	if len(this) == cap(this) {
		fmt.Println("队列已满")
		return this
	}
	return append(this, element)
}

func (this SqQueue) DeQueue() (SqQueue, interface{}) {
	if len(this) == 0 {
		fmt.Println("队列已空")
		return this, nil
	}
	return this[1:], this[0]
}
