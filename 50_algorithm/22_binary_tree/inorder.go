package _22_binary_tree

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 中序遍历，使用递归实现
func InOrderTraversal(root *TreeNode) []interface{} {
	var res []interface{}
	var traversal func(root *TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Value)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// 中序遍历，栈实现
func InOrderTraversal1(root *TreeNode) []interface{} {
	res := []interface{}{}
	stack := NewSqStack(64)

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = (stack.Pop()).(*TreeNode)
		res = append(res, root.Value)
		root = root.Right
	}
	return res
}
