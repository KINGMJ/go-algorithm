package _22_binary_tree

import (
	. "github.com/KINGMJ/go-algorithm/structure"
)

// 后续遍历，递归实现
func PostOrderTraversal(root *TreeNode) interface{} {
	var res []interface{}
	var traversal func(root *TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Value)
	}
	traversal(root)
	return res
}

// 后续遍历，栈实现
func PostOrderTraversal1(root *TreeNode) interface{} {
	res := []interface{}{}
	stack := NewSqStack(64)
	var prev *TreeNode // 已经访问过的节点

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = (stack.Pop()).(*TreeNode)
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Value)
			prev = root
			root = nil
		} else {
			stack.Push(root)
			root = root.Right
		}
	}
	return res
}
