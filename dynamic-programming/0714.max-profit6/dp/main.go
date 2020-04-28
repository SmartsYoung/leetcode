package main

import (
	"fmt"
)

/**
714. 买卖股票的最佳时机含手续费
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

示例 1:

输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
输出: 8
解释: 能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
注意:

0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0], dp[0][1] = 0, -prices[0]
	//	dp[1][0], dp[1][1] = max(0, prices[1]-prices[0]), max(-prices[0], -prices[1])

	for i := 1; i < len(prices); i++ {

		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
	}
	return dp[len(prices)-1][0]
}

//dp 测试用例较大时内存分配不够  原因是引入了较大变量 K，尽量不要引入变量 K ,除非有必要
func maxProfit1(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}
	K := len(prices) / 2
	res := 0
	dp := make([][][2]int, len(prices)) //初始化时注意第二维，第三维的维度

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, K+1)
	}

	// 初始化第0天 有股票 与 无股票的最大利润
	for k := 0; k < K+1; k++ {
		dp[0][k][0], dp[0][k][1] = 0, -prices[0]
	}

	// 第i(i>1)天 0次交易 有股票与无股票的最大利润
	for i := 1; i < len(prices); i++ {
		//第 i 天 0 次交易 没有股票最大利润 = 第 i-1 天 0 次交易 没有股票最大利润
		dp[i][0][0] = dp[i-1][0][0]
		//第 i 天 0 次交易 有股票最大利润 = max(第 i-1 天 0 次交易 有股票最大利润 , 第 i-1 天 0 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		dp[i][0][1] = max(dp[i-1][0][1], dp[i][0][0]-prices[i])
	}

	for i := 1; i < len(prices); i++ {
		for k := 1; k <= K; k++ {
			// #第 i 天 k 次交易 有股票最大利润 = max(第 i-1 天 k 次交易 有股票最大利润 , 第 i-1 天 k-1 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])

			//第 i 天 k 次交易 无股票最大利润 = max(第 i-1 天 k次交易 无股票最大利润 , 第 i-1 天 k次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
			dp[i][k][0] = max(dp[i-1][k][0], dp[i][k][1]+prices[i]-fee) // dp[i][k][0] 与 dp[i][k][1] 的顺序关系很重要，
			// 原因是dp[i][k][0]依赖于dp[i][k][1]
		}
	}

	for k := 0; k < K; k++ {
		res = max(dp[len(prices)-1][k][0], dp[len(prices)-1][k+1][0])
	}

	return res
}

func main() {
	fee := 2
	m := []int{1, 3, 2, 8, 4, 9}
	res := maxProfit(m, fee)
	fmt.Println(res)

	n := []int{3, 2, 6, 5, 0, 3}
	rs := maxProfit(n, fee)
	fmt.Println(rs)
}
