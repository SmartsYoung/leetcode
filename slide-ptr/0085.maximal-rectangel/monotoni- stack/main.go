package main

import (
	"fmt"
	"math"
)

/**
85. 最大矩形   hard 单调栈
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6
*/

func maximalRectangle1(matrix [][]byte) int {
	var maxarea int
	var res int

	if matrix == nil || len(matrix) == 0 { //当matrix初始化并未赋值时，matrix不为nil
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])

	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			}
			height[j] += int(matrix[i][j] - '0')
		}
		res = maxRectangle(height)
		maxarea = int(math.Max(float64(res), float64(maxarea)))
	}

	return maxarea
}

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	mat := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	res := maximalRectangle1(matrix)
	fmt.Println(res)

	fmt.Println(maximalRectangle1(mat))

	arr := []int{2, 1, 5, 6, 2, 3}
	ret := maxRectangle(arr)
	fmt.Println(ret)

}

//单调栈实现
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	//保存最终结果
	max := 0
	//行数，列数
	m, n := len(matrix), len(matrix[0])
	//高度数组（统计每一行中1的高度）
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//每一行去找1的高度
			//如果不是‘1’，则将当前高度置为0
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				//是‘1’，则将高度加1
				height[j] = height[j] + 1
			}
		}
		//更新最大矩形的面积
		max = int(math.Max(float64(max), float64(maxRectangle(height))))
	}
	return max
}

//使用84题的结果
func maxRectangle(heights []int) int {
	//最大矩形面积
	maxarea := 0
	//定义栈
	stack := make([]int, 0)
	//放入-1在栈底是为了:如果矩形包括索引为0处的柱子，则左边界为0的左边，方便计算左边界的索引
	stack = append(stack, -1)
	for i := 0; i < len(heights); i++ {
		//当下一个柱子的高度小于当前栈顶柱子的高度
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			//得到当前栈顶元素的索引
			tmp := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//更新面积
			maxarea = int(math.Max(float64(maxarea), float64(heights[tmp]*(i-stack[len(stack)-1]-1))))
		}
		//当新加入柱子的高度大于栈顶柱子的高度（满足升序）
		stack = append(stack, i)
	}
	//当遍历完数组时，判断栈是否为空
	for stack[len(stack)-1] != -1 {
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
