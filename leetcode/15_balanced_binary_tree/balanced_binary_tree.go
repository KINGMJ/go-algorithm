package balancedbinarytree

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 判断二叉树是否是平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 自底向上的递归
func isBalanced2(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := checkBalance(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := checkBalance(root.Right)
	if rightHeight == -1 {
		return -1
	}
	// 检查当前节点是否平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
