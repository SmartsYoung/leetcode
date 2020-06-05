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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func minDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode       // 辅助队列
	queue = append(queue, root) // 根节点入队
	depth := 1                  // 初始化深度为1

	for len(queue) > 0 { // 当队列不为空时，将队列中的元素出队
		size := len(queue)          // 当前队列中元素个数size作为限定:当前层级中节点的数量
		for i := 0; i < size; i++ { // 每次只取当前层级中的节点
			s := queue[0]     // 获取队首元素
			queue = queue[1:] // 弹出队首元素
			if s.Left == nil && s.Right == nil {
				return depth // 叶子节点直接返回当前累计深度
			}
			if s.Left != nil { // 左子树不为空把左子树入队
				queue = append(queue, s.Left)
			}
			if s.Right != nil { // 右子树不为空把右子树入队
				queue = append(queue, s.Right)
			}
		}
		depth++ // 该层级节点已经访问完，深度+1
	}
	return depth // 一定会访问叶子节点并返回，不会走到这里
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

	res := minDepthIterative(node)
	fmt.Println(res)

}
