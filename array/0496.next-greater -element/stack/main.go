/*
496. 下一个更大元素 I
给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
示例 2:

输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
    对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
*/

package main

import (
	"fmt"
	"math"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil { // 这种写法有问题  空切片值不为nil？
		return []int{}
	}
	if len(nums1) == 0 || len(nums2) == 0 { //规范写法
		return []int{}
	}

	l := len(nums2)

	num := make([]int, l)
	stack := make([]int, 0) // 本题用hash表保存值更高效

	// 注意这里插入栈的顺序
	stack = append(stack, math.MaxInt32)
	stack = append(stack, nums2[l-1])

	num[l-1] = -1

	for i := l - 1; i > 0; i-- {
		if nums2[i-1] < nums2[i] {
			num[i-1] = nums2[i]
			stack = append(stack, nums2[i-1])
		} else {
			for len(stack) > 0 {
				top := stack[len(stack)-1] // 栈顶元素
				if len(stack) > 1 {
					if top < nums2[i-1] { // 栈顶元素小于当前值 出栈
						stack = stack[0 : len(stack)-1]
					} else {
						num[i-1] = top
						stack = append(stack, nums2[i-1])
						break
					}
				} else {
					num[i-1] = -1
					stack = append(stack, nums2[i-1])
					break
				}
			}
		}
	}

	for k, v := range nums1 {
		for k2, v2 := range nums2 {
			if v == v2 {
				nums1[k] = num[k2]
				break
			}
		}
	}
	return nums1
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
	fmt.Println(nextGreaterElement([]int{}, []int{}))
}

// 优化
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	if len(nums2) == 0 {
		return nil
	}

	var stack []int
	ans := make([]int, len(nums2))
	// k，v 表示 值和值在nums2中的index
	m := make(map[int]int, len(nums2))
	// 倒着来，构造单调递减栈
	for i := len(nums2) - 1; i >= 0; i-- {
		m[nums2[i]] = i
		// 单调递减，如果元素比栈顶的大，就一直pop，直到栈为空或者比栈顶小
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			// pop
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			ans[i] = stack[len(stack)-1]
		} else {
			ans[i] = -1
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = ans[m[nums1[i]]]
	}
	return res
}
