package _064_min_path_sum

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

链接：https://leetcode-cn.com/problems/minimum-path-sum
*/

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var dp [][]int

	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	if m == 1 || n == 1 {
		return dp[m-1][n-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + Min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}
