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

package dfs

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	memo := make(map[int]bool)
	return dfs(s, m, 0, memo)
}

func dfs(s string, wordDict map[string]bool, start int, memo map[int]bool) bool {
	if start == len(s) { // start == len(s) 位置为字符串最后一个字母后的空格
		return true
	}
	if _, ok := memo[start]; ok {
		return memo[start]
	}

	//递归
	for i := start; i < len(s); i++ {
		//  if -ok 判断进行了dfs回溯优化
		if _, ok := wordDict[s[start:i+1]]; ok && dfs(s, wordDict, i+1, memo) {
			memo[start] = true // 以start下标开始的字符串可以拆分成字典中的单词
			return true
		}
	}
	// 没找到
	memo[start] = false
	return false
}
