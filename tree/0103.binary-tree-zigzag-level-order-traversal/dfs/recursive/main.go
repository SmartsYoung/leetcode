/*
103. 二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		if level%2 == 0 {
			dfs(root.Left, level+1)
			dfs(root.Right, level+1)
			//res[level] = append(res[level], root.Val)
		} else {
			dfs(root.Right, level+1)
			dfs(root.Left, level+1)
			//res[level] = append([]int{root.Val}, res[level]...)
		}
	}
}
