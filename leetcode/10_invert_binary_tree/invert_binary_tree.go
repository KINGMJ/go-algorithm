package invertbinarytree

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 函数用于翻转二叉树，使用递归实现
func invertTree(root *TreeNode) *TreeNode {
	// 在此实现翻转二叉树的逻辑
	return deepInvertTree(root)
}

func deepInvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	deepInvertTree(root.Left)
	deepInvertTree(root.Right)

	return root
}

// 使用深度优先（栈）实现
func invertTreeWithStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node.Left, node.Right)
		}
	}

	return root
}

// 使用广度优先搜索（队列）实现
func invertTreeWithQueue(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			queue = append(queue, node.Left, node.Right)
		}
	}
	return root
}
