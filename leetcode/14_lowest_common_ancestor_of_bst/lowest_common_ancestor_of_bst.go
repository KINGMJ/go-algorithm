package lowestcommonancestorofbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果 p 和 q 都小于根节点，则 LAC 在左子树
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// 如果 p 和 q 都大于根节点，则 LAC 在右子树
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// 如果 p 和 q 分别在根节点的两侧，则根节点就是 LAC
	return root
}

// 题解：

// 利用二叉查找树的性质，遍历整个二叉树：
// 1. 如果 p 和 q 都小于根节点，则 LAC 在左子树
// 2. 如果 p 和 q 都大于根节点，则 LAC 在右子树
// 3. 如果 p 和 q 分别在根节点的两侧，则根节点就是 LAC
