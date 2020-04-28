package main

import (
	"fmt"
)

/**
188. 买卖股票的最佳时机 IV
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
*/

// dp
func maxProfit(K int, prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	res := 0
	// 当数字K极大时，不用担心操作数目不够，但凡相邻两天价格大于0，就累积
	if K >= len(prices)/2 {
		for i := 1; i < len(prices); i++ {
			if prices[i]-prices[i-1] > 0 {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	// 防止过大测试用例
	/*if len(prices)+1>>2 < K {
		K = len(prices) >> 2
	}*/

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
			dp[i][k][0] = max(dp[i-1][k][0], dp[i][k][1]+prices[i]) // dp[i][k][0] 与 dp[i][k][1] 的顺序关系很重要，
			// 原因是dp[i][k][0]依赖于dp[i][k][1]
		}
	}

	for k := 0; k < K; k++ {
		res = max(dp[len(prices)-1][k][0], dp[len(prices)-1][k+1][0])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	k := 2
	m := []int{2, 4, 1}
	res := maxProfit(k, m)
	fmt.Println(res)

	n := []int{3, 2, 6, 5, 0, 3}
	rs := maxProfit(k, n)
	fmt.Println(rs)

	r := maxProfit(1, []int{1, 2})
	fmt.Println(r)
}
