package main

import (
	"fmt"
	"math"
)

/**
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] < s) {
		return 0
	}

	var l, r = 0, 1
	res := 0
	m := math.MaxInt32

	res = nums[l]
	for r < len(nums) {
		if nums[r] >= s || nums[0] >= s {
			return 1
		}
		res += nums[r]
		for res >= s { // 若满足连续子数组之和大于等于s。则移动左指针
			m = min(m, r-l+1)
			res -= nums[l]
			l++
		}
		r++
	}
	if m == math.MaxInt32 {
		return 0
	}
	return m
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {

	arr := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(7, arr)
	fmt.Println(res)

	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
