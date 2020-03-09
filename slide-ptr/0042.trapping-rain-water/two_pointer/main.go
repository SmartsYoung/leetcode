package main

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

链接：https://leetcode-cn.com/problems/trapping-rain-water
*/
func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left] //设置左边最高柱子
			} else {
				res += leftMax - height[left] //右边必定有柱子挡水，所以，遇到所有值小于leftMax的，全部加入水池
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right] //设置右边最高柱子
			} else {
				res += rightMax - height[right] //左边必定有柱子挡水，所以，遇到所有值小于rightMax的，全部加入水池
			}
			right--
		}
	}
	return res
}
