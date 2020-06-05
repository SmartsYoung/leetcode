/*
99. 恢复二叉搜索树
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历 o(1)
var last *TreeNode
var wrong []*TreeNode

func recoverTree(root *TreeNode) {
	last = nil
	wrong = make([]*TreeNode, 0)
	dfs(root)
	wrong[0].Val, wrong[1].Val = wrong[1].Val, wrong[0].Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if last != nil && root.Val <= last.Val {
		if len(wrong) == 0 {
			wrong = append(wrong, last)
			wrong = append(wrong, root)
		} else {
			wrong[1] = root
		}
	}
	last = root
	dfs(root.Right)
}
