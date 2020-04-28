package main

import "math"
import "fmt"

/**
84. 柱状图中最大的矩形  hard
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:

输入: [2,1,5,6,2,3]
输出: 10
*/

func largestRectangleArea(heights []int) int {
	var maxarea int
	stack := make([]int, 1)
	stack = append(stack, -1) // 放入-1在栈底是为了:如果矩形包括索引为0处的柱子，则左边界为0的左边，方便计算左边界的索引

	for i := 0; i < len(heights); i++ {
		//当下一个柱子的高度小于当前栈顶柱子的高度
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] { //注意这里是for循环，满足此条件时一直执行循环
			//得到当前栈顶元素的索引
			tmp := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//更新面积                        弹出的栈顶元素  *  左右两边满足第一个小于弹出的栈顶元素值的索引（左边的即为当前的栈顶元素，右边的即为i)
			maxarea = int(math.Max(float64(maxarea), float64(heights[tmp]*(i-stack[len(stack)-1]-1))))
		}
		//当新加入柱子的高度大于栈顶柱子的高度（满足升序）
		stack = append(stack, i)
	}
	//当遍历完数组时，判断栈是否为空            剩下的是一个严格单调递增的数组
	for stack[len(stack)-1] != -1 { //for循环一直执行，指导栈顶元素为-1
		//得到当前栈顶元素索引
		tmp := stack[len(stack)-1]
		//出栈
		stack = stack[:len(stack)-1]
		//更新面积
		maxarea = int(math.Max(float64(maxarea), float64(heights[tmp]*(len(heights)-1-stack[len(stack)-1]))))
	}

	return maxarea
}

/*作者：shixuewei
链接：https://leetcode-cn.com/problems/maximal-rectangle/solution/golangjie-jue-si-lu-by-shixuewei/*/

func largestArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var maxarea int
	length := len(heights)
	leftMin := make([]int, length) // leftMin[i] 保存各自的左边第一个小的柱子。
	rightMin := make([]int, length)

	leftMin[0], rightMin[length-1] = -1, length

	for i := 1; i < length; i++ { // 时间复杂度为o(n)
		l := i - 1
		for l >= 0 && heights[l] >= heights[i] { //当前柱子更小一些，进行左移
			l = leftMin[l]
		}
		leftMin[i] = l
	}

	for i := length - 2; i >= 0; i-- {
		r := i + 1
		for r <= length-1 && heights[r] >= heights[i] { //当前柱子更小一些，进行右移
			r = rightMin[r]
		}
		rightMin[i] = r
	}

	for i := 0; i < length; i++ {
		res := heights[i] * (rightMin[i] - leftMin[i] - 1)
		maxarea = int(math.Max(float64(res), float64(maxarea)))
	}
	return maxarea
}

func main() {
	arr := []int{2, 1, 5, 6, 2, 3}
	ret := largestArea(arr)
	fmt.Println(ret)
}
