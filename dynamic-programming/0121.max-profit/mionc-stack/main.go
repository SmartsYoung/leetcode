package main

import (
	"fmt"
	"math"
)

/**
121. 买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。


示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

// 单调栈
func maxProfit(prices []int) int {

	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, math.MinInt32)
	m := 0
	for i := 0; i < len(prices); i++ {

		for stack[len(stack)-1] != math.MinInt32 && stack[len(stack)-1] > prices[i] {
			m = int(math.Max(float64(stack[len(stack)-1]-stack[1]), float64(m)))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, prices[i])
	}
	m = int(math.Max(float64(stack[len(stack)-1]-stack[1]), float64(m)))
	return m
}

// dp
func maxProfit1(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = int(math.Min(float64(prices[i]), float64(minPrice)))
		dp[i] = int(math.Max(float64(dp[i-1]), float64(prices[i]-minPrice)))
	}
	return dp[len(prices)-1]
}

func main() {

	m := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(m)
	fmt.Println(res)

	n := []int{1, 2}
	rs := maxProfit1(n)
	fmt.Println(rs)
}
