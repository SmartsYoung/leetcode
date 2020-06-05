package main

import "fmt"

/**
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

/*
终止条件、返回值和递归过程：
     当前节点root为空时，说明此树的高度为0，0也是最小值
     当前节点root的左子树和右子树都为空时，说明此树的高度为1，1也是最小值。
     如果为其他情况，则说明当前节点有值，且需要分别计算其左右子树的最小深度，返回最小深度+1，+1表示当前节点存在有1个深度
     时间复杂度是O(N)，N为树的节点数量
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func Min(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

// 模板
var min int

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = 1<<63 - 1
	dfs(root, 1)
	return min
}

func dfs(node *TreeNode, depth int) {
	if node.Right == nil && node.Left == nil { //虽然这里没有打印根，但我们停在这里（根）判断了一下，勉强当作前序遍历
		if depth < min {
			min = depth
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}

	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}

var node = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

func main() {

	res := minDepth(node)
	fmt.Println(res)

}
