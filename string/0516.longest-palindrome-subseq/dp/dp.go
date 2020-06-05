package dp

/**
516. 最长回文子序列  medium
给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。

示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2
一个可能的最长回文子序列为 "bb"。
*/
/**
dp[i][j] 保存的是s[i:j+1]子串的最长回文子序列的长度， 只用到dp二维矩阵的左上部分
最小子问题为j-i=1时,dp[i][i+1] = dp[i+1][i] + 2或者dp[i][i+1] = max(dp[i+1][i+1], dp[i][i]) 此时全部子问题为已知解。

 s[i]==s[j] dp[i][j] = dp[i+1][j-1] +2
 s[i] != s[j] dp[i][j] = max(dp[i+1][j], dp[i][j-1])

*/
func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := len(s)
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		dp[i][i] = 1
	}
	for j := 1; j < m; j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化空间复杂度
func longestPalindromeSubseq1(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := len(s)

	pre := make([]int, m)
	cur := make([]int, m)

	for i := m - 1; i >= 0; i-- {
		cur[i] = 1
		for j := i + 1; j < m; j++ {
			if s[i] == s[j] {
				cur[j] = pre[j-1] + 2 // 上一个切片的状态dp[i+1][j-1]         但凡依赖 状态 i+1 就是上个切片pre， 但凡依赖  状态i 就是当前切片cur
			} else {
				cur[j] = max(cur[j-1], pre[j]) // 当前切片的上一个状态dp[i][j-1]  与 上一个切片的状态dp[i+1][j]
			}
		}
		pre, cur = cur, pre // 交换切片   当前切片在下一次循环中变成上个切片  为什么要交换 cur = pre ，只从一次循环来讲似乎不影响 pre 取值 多次循环后 cur 中产生的一些与当前状态无关的值也会被保存到 pre中
	}

	return pre[m-1]
}

// 优化空间复杂度
func longestPalindromeSubseq2(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := len(s)

	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}

	prev := 0
	temp := 0
	for i := m - 1; i >= 0; i-- {
		prev = 0
		for j := i + 1; j < m; j++ {
			temp = dp[j] // 更新当前长度
			if s[i] == s[j] {
				dp[j] = prev + 2 // 上一个状态dp[i+1][j-1]
			} else {
				dp[j] = max(dp[j-1], dp[j]) // 上一个内层for循环状态dp[i][j-1]  与 上一个外层for循环状态dp[i+1][j]，括号内的dp[j]即为temp
			}
			prev = temp // 保存当前长度
		}
	}

	return dp[m-1]
}
