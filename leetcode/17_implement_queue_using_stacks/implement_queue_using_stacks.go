package implement_queue_using_stacks

// MyQueue 用栈实现队列的数据结构
type MyQueue struct {
	// 数组0为栈底，数组末尾为栈顶
	stackIn  []int // 用于入队的栈
	stackOut []int // 用于出队的栈
}

// Constructor 初始化数据结构
func Constructor() MyQueue {
	// 这里需要实现构造函数
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// Push 将元素 x 推到队列的末尾
func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

// Pop 从队列的开头移除并返回元素
func (q *MyQueue) Pop() int {
	q.transfer()
	if len(q.stackOut) == 0 {
		return 0
	}
	last := len(q.stackOut) - 1
	val := q.stackOut[last]
	q.stackOut = q.stackOut[:last]
	return val
}

// Peek 返回队列开头的元素
func (q *MyQueue) Peek() int {
	// 这里需要实现 peek 方法
	q.transfer()
	if len(q.stackOut) == 0 {
		return 0
	}
	return q.stackOut[len(q.stackOut)-1]
}

// Empty 返回队列是否为空
func (q *MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

// 如果stockOut为空，则将stockIn中的元素转移stockOut中
func (q *MyQueue) transfer() {
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			last := len(q.stackIn) - 1
			q.stackOut = append(q.stackOut, q.stackIn[last])
			q.stackIn = q.stackIn[:last]
		}
	}
}
