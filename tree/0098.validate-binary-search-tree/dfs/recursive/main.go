/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
dfs
1.左节点小于根节点
2.右节点大于根节点
3.左节点的右节点小于根节点
（因为 根节点的左节点下的所有值，必须全部小于根节点）
4.右节点的左节点大于根节点
（因为 根节点的右节点下的所有值，必须全部大于根节点）
*/

//func isValidBST(root *TreeNode) bool {
//	return dfs(root, nil, nil)
//}
//
//func dfs(root *TreeNode, min, max *int) bool {
//	if root == nil {
//		return true
//	}
//	if min != nil && root.Val <= *min || max != nil && *max <= root.Val {
//		return false
//	}
//	return dfs(root.Left, min, &root.Val) && dfs(root.Right, &root.Val, max)
//}

//中序遍历  左--根--右
var last *TreeNode

func isValidBST(root *TreeNode) bool {
	last = nil
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// last 为上一次访问节点
	if !dfs(root.Left) || last != nil && root.Val <= last.Val {
		return false
	}
	last = root
	return dfs(root.Right)
}

func isValidBST1(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		// 循环插入根节点以及左节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
