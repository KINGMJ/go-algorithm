package _11_queue_on_array

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

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 使用结构体实现的队列，这种方式有个缺陷，出队列后只是移动了head指针，数组空间没有被合理利用，可以演进为循环队列
type SqQueue2 struct {
	queue    []interface{}
	capacity int // 容量
	head     int // 队头指针，通过移动队头指针，出队列的时候不需要所有的元素往前移动一位
	tail     int // 队尾指针，队满条件 tail == capacity
}

func NewSqQueue2(n int) *SqQueue2 {
	return &SqQueue2{make([]interface{}, n), n, 0, 0}
}

func (this *SqQueue2) EnQueue(element interface{}) {
	// 满队列
	if this.tail == this.capacity {
		return
	}
	this.queue[this.tail] = element
	this.tail++
}

func (this *SqQueue2) DeQueue() interface{} {
	// 空队列
	if this.head == this.tail {
		return nil
	}
	element := this.queue[this.head]
	this.head++
	return element
}
