package dp

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

// 用动态规划的时间复杂度为o(n2） 与中心扩散法时间复杂度一样，但实际执行时间是下面所写中心扩散法的15倍
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	m := len(s)
	dp := make([][]bool, m)
	start := 0
	l := 1
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}
	// fmt.Println(dp)
	for j := 1; j < m; j++ {
		for i := j - 1; i >= 0; i-- {
			dynamic(dp, i, j, s)
			if dp[i][j] == true {
				if l < j-i+1 {
					l = j - i + 1
					start = i
				}
			}
		}
	}
	return s[start : start+l]
}

// 用动态规划的时间复杂度为o(n2） 降低空间复杂度 执行时间减少16ms，但执行时间仍是中心扩散法的10倍
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	m := len(s)
	dp := make([]bool, m)

	start := 0
	l := 1

	for i := m - 1; i >= 0; i-- {
		for j := m - 1; j >= i; j-- {
			// 因为 j 值是递减的，所以 当前 dp[j-1]的值取自于上一层 for循环 i=i+1，j=j-1  ,内层for循环dp[j]的值会存到下一层for循环时取值

			dp[j] = s[i] == s[j] && (j-i < 3 || dp[j-1]) //注意后面的大括号不能去掉 如 j-i  < 3 为 false 而dp[j-1]为 true

			if dp[j] == true && l < j-i+1 {
				l = j - i + 1
				start = i
			}
		}
	}

	/*for j := 0; j < m; j++ {
		for i := 0; i <= j; i++ {
	       dp[i] = s[i] == s[j] && (j-i < 3 || dp[i+1])
			 if dp[i] && j-i+1 > l {
					start = i
					l = j - i + 1
			}
		}
	}*/

	return s[start : start+l]
}

func dynamic(dp [][]bool, i, j int, s string) {
	if s[i] == s[j] {
		if j-i < 3 {
			dp[i][j] = true
		} else {
			dp[i][j] = dp[i+1][j-1]
		}
	} //else { // 可以去掉这个 else 语句
	// dp[i][j] = false
	//	}
}

func dynamic1(dp [][]bool, i, j int, s string) {
	res := s[i] == s[j] && (j-i) < 3 // 合并条件语句降低圈复杂度？ 貌似没有
	if res {
		dp[i][j] = true
	} else {
		dp[i][j] = dp[i+1][j-1]
	}
}

func dynamic2(dp [][]bool, i, j int, s string) {
	dp[i][j] = s[i] == s[j] && (j-i) < 3 || dp[i+1][j-1] // 合并条件语句降低圈复杂度
}

func isPalindrome(dp [][]bool, i, j, l, start int) bool {
	if dp[i][j] == true {
		if l < j-i+1 {
			l = j - i + 1
			start = i
		}
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
思路：
1.每扩散一次，设置中心与半径
2.每移动一次，判断是否有镜像
*/

func reverse(s string, l, r int) string {
	for l > 0 && r < len(s)-1 {
		l--
		r++
		if s[l] != s[r] {
			return s[l+1 : r]
		}
	}
	return ""
}

func longestPalindrome1(s string) string {
	// 字符串首尾添加字符 ^ # 便于边界处理
	str := string('^')
	str += s
	str += string('#')
	s1 := ""
	s2 := ""
	res := ""
	for i := 1; i < len(str)-1; i++ {
		s1 = reverse(str, i, i)
		if str[i] == str[i+1] {
			s2 = reverse(str, i, i+1)
		}
		temp := maxStr(s1, s2)
		res = maxStr(res, temp)

	}
	return res
}

func maxStr(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
