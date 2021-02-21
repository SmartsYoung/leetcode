package _653_find_target

/**
653. 两数之和 IV - 输入 BST
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

案例 1:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

输出: True


案例 2:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

输出: False
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	num := make([]int, 0)
	tree(root, &num)

	if len(num) == 0 || len(num) == 1 {
		return false
	}
	m := make(map[int]int, len(num))
	for _, v := range num {
		if _, ok := m[k-v]; ok {
			return true
		}
		m[v] = v
	}
	return false
}

func tree(root *TreeNode, num *[]int) {
	if root == nil {
		return
	}
	tree(root.Left, num)
	*num = append(*num, root.Val)
	tree(root.Right, num)
}
