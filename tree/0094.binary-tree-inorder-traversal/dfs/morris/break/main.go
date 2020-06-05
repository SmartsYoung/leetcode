/*
94. 二叉树的中序遍历
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var last *TreeNode

	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			last = root.Left
			for last.Right != nil {
				last = last.Right
			}
			last.Right = root
			root, root.Left = root.Left, nil
		}
	}
	return res
}
