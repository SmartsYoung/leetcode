package _053_maxSubArray

import "math"

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

链接：https://leetcode-cn.com/problems/maximum-subarray
*/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	currSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(dp[i], res)
	}
	return res
}

func maxSubArrayDivide(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) >> 1
	leftMax := maxSubArrayDivide(nums, left, mid)
	rightMax := maxSubArrayDivide(nums, mid+1, right)

	twoSideMax := max(leftMax, rightMax)

	leftCrossMax := math.MinInt32
	leftCrossSum := 0

	for i := mid; i >= 0; i-- {
		leftCrossSum += nums[i]
		leftCrossMax = max(leftCrossMax, leftCrossSum)
	}

	rightCrossMax := math.MinInt32
	rightCrossSum := 0
	for i := mid + 1; i < len(nums); i++ {
		rightCrossSum += nums[i]
		rightCrossMax = max(rightCrossMax, rightCrossSum)
	}
	combineMax := rightCrossMax + leftCrossMax

	return max(combineMax, twoSideMax)
}
