/**
139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/
package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1, len(s)+1)
	dp[0] = true

	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	// dp 表示以当前字符为止，之前的字符是否可拆
	// dp[j] && dict[s[j:i]] 表示以当前j为下标的字符串是否可拆和字典中是否含有s[i:j]
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
