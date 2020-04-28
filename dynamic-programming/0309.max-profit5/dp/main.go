package main

import (
	"fmt"
)

/**
309. 最佳买卖股票时机含冷冻期
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:

输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

*/

// 差分 大部分用例通过 考虑情况不完全
func maxProfit1(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	res := 0

	num := make([]int, len(prices)+1)
	num[0] = 0
	num[len(prices)] = 0
	for i := 1; i < len(prices); i++ {
		num[i] = prices[i] - prices[i-1]
		if num[i] > 0 {
			res += num[i]
		}
	}

	for i := 1; i < len(prices); i++ {
		if num[i] < 0 {
			res -= min(num[i-1], num[i+1])
		}
	}

	return res
}

func min(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//dp
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0], dp[0][1] = 0, -prices[0]
	dp[1][0], dp[1][1] = max(0, prices[1]-prices[0]), max(-prices[0], -prices[1])

	for i := 2; i < len(prices); i++ {
		//第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。 第 i-1天手上没有股票时第 i 天不一定能买
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	}
	return dp[len(prices)-1][0]
}
func main() {

	m := []int{1, 2, 3, 0, 2}
	res := maxProfit(m)
	fmt.Println(res)

	n := []int{3, 2, 6, 5, 0, 3}
	rs := maxProfit(n)
	fmt.Println(rs)

	r := maxProfit([]int{1, 2})
	fmt.Println(r)
}
