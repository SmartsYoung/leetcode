package problem0938

import "github.com/SmartsYoung/LeetCode-in-Go/kit"

// TreeNode is Definition for a binary tree node.
type TreeNode = kit.TreeNode

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0

	switch {
	case root.Val < L:
		sum = rangeSumBST(root.Right, L, R)
	case R < root.Val:
		sum = rangeSumBST(root.Left, L, R)
	default:
		sum = root.Val
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

/*遍历二叉搜索树，求出所有节点在[L,R]之间的值
优化：
1、当前节点值小于L，说明该节点的左子树，会在[L,R]的左边，则不进行遍历该节点的左子树
2、当前节点值大于R，说明该节点的右子树，会在[L,R]的右边，则不进行遍历该节点的右子树*/

func rangeSumBST1(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}
