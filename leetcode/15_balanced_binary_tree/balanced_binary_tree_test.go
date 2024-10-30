package balancedbinarytree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "平衡二叉树",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "非平衡二叉树",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name:     "空树",
			root:     nil,
			expected: true,
		},
		{
			name: "只有根节点",
			root: &TreeNode{
				Val: 1,
			},
			expected: true,
		},
		{
			name: "左右子树高度差为1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.expected {
				t.Errorf("isBalanced() = %v, want %v", got, tt.expected)
			}
		})
	}
}
