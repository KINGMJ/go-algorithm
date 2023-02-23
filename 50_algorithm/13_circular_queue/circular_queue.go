package _13_circular_queue

type CircularQueue struct {
	data     []interface{}
	front    int // 队头指针
	rear     int // 队尾指针
	capacity int // 队列的容量
}

// 初始化一个循环队列
//
//	@param capacity
//	@return *CircularQueue
func NewCircularQueue(capacity int) *CircularQueue {
	if capacity == 0 {
		return nil
	}
	// 循环队列为空时，front和rear指针都指向data[0]
	return &CircularQueue{make([]interface{}, capacity), 0, 0, capacity}
}

func (this *CircularQueue) EnQueue(element interface{}) {
	// 判断队列是否已满
	if this.IsFull() {
		return
	}
	//将elem赋值给队尾
	this.data[this.rear] = element
	//队尾指针后移一位，若到最后则转到数组头部
	this.rear = (this.rear + 1) % this.capacity
}

func (this *CircularQueue) DeQueue() interface{} {
	// 判断队列是否为空
	if this.front == this.rear {
		return nil
	}
	// 出队列
	element := this.data[this.front]
	this.data[this.front] = nil
	// front指针向后移动一位，若到最后则转到数组头部
	this.front = (this.front + 1) % this.capacity
	return element
}

// 判断队列是否已满
//
//	@receiver this
//	@return bool
func (this *CircularQueue) IsFull() bool {
	return (this.rear+1)%this.capacity == this.front
}

// 返回队列长度，使用队列长度通用公式
//
//	@receiver this
//	@return int
func (this *CircularQueue) getLength() int {
	return (this.rear - this.front + this.capacity) % this.capacity
}
