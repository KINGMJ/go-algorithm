package structure

// 二叉树节点
type TreeNode struct {
	Value any
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value any) *TreeNode {
	return &TreeNode{Value: value, Left: nil, Right: nil}
}

// 数组转二叉树，使用leetcode紧凑表示法
func Array2Tree(arr []any) *TreeNode {
	len := len(arr)
	if len == 0 {
		return nil
	}
	// 根节点
	root := NewTreeNode(arr[0])
	isLChild := true

	// 根节点放入队列
	queue := NewSqQueue(len)
	queue = queue.EnQueue(root)

	// 从第二个元素开始遍历：
	// 如果是左节点直接追加，并且进队列；如果是右节点，处理完毕后需要出队列，进入下一层
	for i := 1; i < len; i++ {
		// 从队列中取出队首元素，作为父节点
		node := (queue[0]).(*TreeNode)
		if isLChild {
			if arr[i] != nil {
				node.Left = NewTreeNode(arr[i])
				// Left进队
				queue = queue.EnQueue(node.Left)
			}
			// Left处理完毕，开始处理Right
			isLChild = false
		} else {
			if arr[i] != nil {
				node.Right = NewTreeNode(arr[i])
				// Right进队
				queue = queue.EnQueue(node.Right)
			}
			// Right 处理完毕，开始出队列；处理下一层
			queue, _ = queue.DeQueue()
			isLChild = true
		}
	}
	return root
}
