/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

链接：https://leetcode-cn.com/problems/permutations
*/

package dfs

import "sort"

func permute(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	num := make([]bool, len(nums))

	for i := 0; i < length; i++ {
		num[i] = false
	}
	sort.Ints(nums)
	digui(nums, num, []int{}, &res)
	return res

}
func digui(nums []int, num []bool, arr []int, res *[][]int) {
	if len(arr) == len(nums) {
		//*res = append(*res, append([]int{}, arr...)) //这两语句有什么区别？为什么替换之后有时候结果会出错
		//*res = append(*res, arr)
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)

		return
	}
	for i := 0; i < len(nums); i++ {
		if num[i] == false {
			num[i] = true
			arr = append(arr, nums[i])
			digui(nums, num, arr, res)
			arr = arr[:len(arr)-1] //回溯
			num[i] = false
		}
	}
	return
}
