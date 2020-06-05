package main

import "fmt"

/**
112. 路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	// 对于其他情况则只需要递归判断左右子树即可,和sum则需要减去当前节点的值去更新
	// 左右子树只要能找到一个满足条件的路径即可
	return hasPathSum1(root.Left, sum-root.Val) || hasPathSum1(root.Right, sum-root.Val)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	if hasPathSum2(root.Left, sum-root.Val) {
		return true
	}
	if hasPathSum2(root.Right, sum-root.Val) {
		return true
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool

func hasPathSum(root *TreeNode, sum int) bool {
	res = false // res 申明为全局变量时应当在函数内再次赋值
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil && sum == root.Val {
		return true
	}
	dfs(root, sum)
	return res
}

func dfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	if dfs(root.Left, sum-root.Val) {
		res = true
		return true
	}
	if dfs(root.Right, sum-root.Val) {
		res = true
		return true
	}
	return false
}

var node = &TreeNode{
	Val:  1,
	Left: nil,
	Right: &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	},
}

func main() {

	var root *TreeNode = nil
	sum := 1
	res := hasPathSum(root, sum)
	fmt.Println(res)
	r := hasPathSum(node, 1)
	fmt.Println(r)
}
