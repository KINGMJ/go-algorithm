package structure

type ListNode struct {
	Data any
	Next *ListNode
}

// 单链表数据结构
type LinkedList struct {
	Head   *ListNode // 头结点，一般数据域为空或存放链表的长度
	Length int       // 链表长度
}

// 模拟构造函数，构造一个单链表
//
//	@return *LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{Head: &ListNode{Data: nil, Next: nil}, Length: 0}
}

// 数组转链表
//
//	@param arr
func Array2List(arr []any) *LinkedList {
	linkedList := NewLinkedList()
	current := linkedList.Head
	for _, item := range arr {
		current.Next = &ListNode{Data: item, Next: nil}
		current = current.Next
	}
	return linkedList
}

// 链表转数组： 使用递归 + 闭包
//
//	@receiver linkList
//	@return []any
func (l *LinkedList) ToArray() []any {
	return Node2Array(l.Head.Next)
}

// node 结点转数组
//
//	@param node
//	@return []any
func Node2Array(node *ListNode) []any {
	if node == nil {
		return []any{}
	}
	//使用递归，先得到当前节点的数据
	current := []any{node.Data}
	//然后得到剩余部分的数据
	rest := Node2Array(node.Next)
	//合并结果
	return append(current, rest...)
}
