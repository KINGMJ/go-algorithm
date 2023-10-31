package _22_binary_tree

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 前序遍历，使用递归实现
func PreOrderTraversal(root *TreeNode) []interface{} {
	var res []interface{}
	var traversal func(root *TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Value)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// 前序遍历，使用栈实现
func PreOrderTraversal1(root *TreeNode) []interface{} {
	res := []interface{}{}
	stack := NewSqStack(64)
	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			res = append(res, root.Value)
			root = root.Left
		}
		root = (stack.Pop()).(*TreeNode)
		root = root.Right
	}
	return res
}
