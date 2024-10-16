package structure

import "fmt"

// 顺序队列实现：数组本身就可以看做是一个顺序队列
type SqQueue []any

type SqQueueInterface interface {
	// 进队列
	//  @param any
	EnQueue(any) SqQueue

	// 出队列
	//  @return SqQueue
	//  @return any
	DeQueue() (SqQueue, any)
}

func NewSqQueue(capacity int) SqQueue {
	return SqQueue(make([]any, 0, capacity))
}

func (q SqQueue) EnQueue(element any) SqQueue {
	if len(q) == cap(q) {
		fmt.Println("队列已满")
		return q
	}
	return append(q, element)
}

func (q SqQueue) DeQueue() (SqQueue, any) {
	if len(q) == 0 {
		fmt.Println("队列已空")
		return q, nil
	}
	return q[1:], q[0]
}
