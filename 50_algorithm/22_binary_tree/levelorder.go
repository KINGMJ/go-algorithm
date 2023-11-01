package _22_binary_tree

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 队列实现层序遍历
func LeveLOrderTraversal(root *TreeNode) []interface{} {
	res := []interface{}{}
	queue := NewSqQueue(64)
	// root 进队列
	queue = queue.EnQueue(root)
	// 循环处理队列
	for len(queue) > 0 {
		var element interface{}
		queue, element := queue.DeQueue()
		root = element.(*TreeNode)
		res = append(res, root.Value)

		if root.Left != nil {
			queue = queue.EnQueue(root.Left)
		}
		if root.Right != nil {
			queue = queue.EnQueue(root.Right)
		}
	}
	return res
}
