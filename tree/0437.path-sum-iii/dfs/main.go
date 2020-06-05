package dfs

/**
437. 路径总和 III
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	pathImLeading := count(root, sum)        // 自己为开头的路径数
	leftPathSum := pathSum(root.Left, sum)   // 左边路径总数（相信他能算出来）
	rightPathSum := pathSum(root.Right, sum) // 右边路径总数（相信他能算出来）
	return leftPathSum + rightPathSum + pathImLeading
}

func count(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	// 我自己能不能独当一面，作为一条单独的路径呢？   这里能保证该路径不一定是叶子节点
	isMe := 0
	if root.Val == sum {
		isMe = 1
	}

	// 左边的小老弟，你那边能凑几个 sum - node.val 呀？
	leftBrother := count(root.Left, sum-root.Val)
	// 右边的小老弟，你那边能凑几个 sum - node.val 呀？
	rightBrother := count(root.Right, sum-root.Val)
	return isMe + leftBrother + rightBrother
}

// dfs
var s int
var prefixSum map[int]int

func pathSum1(root *TreeNode, sum int) int {
	s = sum
	prefixSum = map[int]int{}
	prefixSum[0] = 1
	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	res := prefixSum[cur-s] //负数索引忽略。刚好为0，或者刚好为前缀  则+1
	prefixSum[cur]++        //加入前缀
	res += dfs(root.Left, cur) + dfs(root.Right, cur)
	prefixSum[cur]-- //取消前缀
	return res
}
