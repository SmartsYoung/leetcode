package _305_get_all_elments

/*
1305. 两棵二叉搜索树中的所有元素
给你 root1 和 root2 这两棵二叉搜索树。

请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.



示例 1：

输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]
示例 2：

输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
输出：[-10,0,0,1,2,5,7,10]
示例 3：

输入：root1 = [], root2 = [5,1,7,0,2]
输出：[0,1,2,5,7]
示例 4：

输入：root1 = [0,-10,10], root2 = []
输出：[-10,0,10]
示例 5：


输入：root1 = [1,null,8], root2 = [8,1]
输出：[1,1,8,8]


提示：

每棵树最多有 5000 个节点。
每个节点的值在 [-10^5, 10^5] 之间。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	if root1 == nil && root2 == nil {
		return []int{}
	}
	var list1 []int
	var list2 []int
	midOrder(root1, &list1)
	midOrder(root2, &list2)
	return merge(list1, list2)
}

func merge(list1, list2 []int) []int {
	if len(list1) == 0 {
		return list2
	}
	if len(list2) == 0 {
		return list1
	}
	var res []int
	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] <= list2[j] {
			res = append(res, list1[i])
			i++
		} else {
			res = append(res, list2[j])
			j++
		}
	}
	if i < len(list1) {
		res = append(res, list1[i:]...)
	}
	if j < len(list2) {
		res = append(res, list2[j:]...)
	}
	return res
}

// 中序遍历
func midOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, list)
	*list = append(*list, root.Val)
	midOrder(root.Right, list)
}
