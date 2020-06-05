package dfs

/**
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	path := []int{}
	find(root, sum, 0, path)
	return res
}

var res [][]int

func find(root *TreeNode, sum int, count int, path []int) {
	if root == nil {
		return
	}
	count += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && count == sum {
		// path是全局的，避免后面的路径把它污染，所以先把当前结果拷贝出来
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	find(root.Left, sum, count, path)
	find(root.Right, sum, count, path)
}
