package structure

// 二叉树节点，在 binary_tree 中定义过
// type TreeNode struct {
// 	Value interface{}
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 二叉搜索树结构，里面保存一个根节点
type BinarySearchTree struct {
	Root *TreeNode
}

// 接口定义
type BinarySearchTreeInterface interface {
	// 插入操作
	// @param int 要插入的数据
	Insert(int)

	// 插入操作，支持批量插入
	Inserts(...int)

	// 查询操作
	// @param int 要查询的数据
	Search(int) *TreeNode

	// 中序遍历，输出有序序列
	InOrderTraversal() []int

	// 删除操作
	// @param int 要删除的数据
	Delete(int)
}

// 插入操作
func (this *BinarySearchTree) Insert(data int) {
	var insertRec func(*TreeNode, int) *TreeNode
	insertRec = func(root *TreeNode, data int) *TreeNode {
		// 如果 root不存在，插入
		if root == nil {
			return NewTreeNode(data)
		}
		// 如果要插入的数据比节点的数据大，递归右子树
		if data > (root.Value).(int) {
			root.Right = insertRec(root.Right, data)
			// 如果要插入的数据比节点的数据小，递归左子树
		} else if data < (root.Value).(int) {
			root.Left = insertRec(root.Left, data)
		}
		// 否则返回 root 节点
		return root
	}
	this.Root = insertRec(this.Root, data)
}

// 批量插入
func (this *BinarySearchTree) Inserts(args ...int) {
	for _, data := range args {
		this.Insert(data)
	}
}

// 查找操作
func (this *BinarySearchTree) Search(data int) *TreeNode {
	var searchRec func(*TreeNode, int) *TreeNode
	// 递归查询
	searchRec = func(root *TreeNode, data int) *TreeNode {
		// 先取根节点，如果等于我们要查找的数据，就返回
		if root == nil || (root.Value).(int) == data {
			return root
		}
		// 如果要查询的数据比根节点小，就在左子树中递归查找
		if data < (root.Value).(int) {
			return searchRec(root.Left, data)
		}
		// 如果查询的数据比根节点大，就在右子树中递归查找
		return searchRec(root.Right, data)
	}
	return searchRec(this.Root, data)
}

func (this *BinarySearchTree) Delete(key int) {
	var deleteRec func(*TreeNode, int) *TreeNode
	// 递归删除
	deleteRec = func(root *TreeNode, key int) *TreeNode {
		// 1. root 为空，代表未搜索到值为 key 的节点，返回空。
		if root == nil {
			return nil
		}
		// 2. root.val > key，表示值为 key 的节点可能存在于 root 的左子树中，
		// 需要递归地在 root.left 调用 deleteNode，并返回 root。
		if (root.Value).(int) > key {
			root.Left = deleteRec(root.Left, key)
			// 3. root.val < key，表示值为 key 的节点可能存在于 root 的右子树中，
			// 需要递归地在 root.right 调用 deleteNode，并返回 root。
		} else if (root.Value).(int) < key {
			root.Right = deleteRec(root.Right, key)
		} else {
			// 4. root.val=key，root 即为要删除的节点。此时要做的是删除 root，并将它的子树合并成一棵子树，保持有序性，并返回根节点。
			// 根据 root 的子树情况分成以下情况讨论：

			// 4a. root 为叶子节点，没有子树。此时可以直接将它删除，即返回空。（递归最开始已经处理了）
			// 4b. root 只有右子树，没有左子树。此时可以将它的右子树作为新的子树，返回它的右子节点。
			// 4c. root 只有左子树，没有右子树。此时可以将它的左子树作为新的子树，返回它的左子节点。
			if root.Left == nil {
				return root.Right
			} else if root.Right == nil {
				return root.Left
			}
			// 4d. root 有左右子树，这时可以将 root 的后继节点（比 root 大的最小节点，即它的右子树中的最小节点，
			// 记为 successor 作为新的根节点替代 root ，并将 successor 从 root 的右子树中删除，使得在保持有序性的情况下合并左右子树。

			// 找到后继结点
			successor := findSuccessor(root.Right)
			// 用 successor 作为新的根节点替代 root
			root.Value = successor.Value
			// 将 successor 从 root 的右子树中删除
			root.Right = deleteRec(root.Right, (successor.Value).(int))
		}
		return root
	}
	deleteRec(this.Root, key)
}

// 找到后继结点
func findSuccessor(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 中序遍历
func (this *BinarySearchTree) InOrderTraversal() []int {
	var res []int
	var traversal func(root *TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, (root.Value).(int))
		traversal(root.Right)
	}
	traversal(this.Root)
	return res
}
