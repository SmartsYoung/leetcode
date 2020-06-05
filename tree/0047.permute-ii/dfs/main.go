/*
47. 全排列 II
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

package dfs

import "sort"

func permuteUnique(nums []int) [][]int {
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
		*res = append(*res, append([]int{}, arr...)) //这两语句有什么区别？为什么替换之后有时候结果会出错
		//*res = append(*res, arr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !num[i-1] { //排除相同的元素  !num[i-1] 保证 num[i-1]这个值的数没有被重复使用
			continue
		}
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
