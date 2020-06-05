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

package bfs

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	visited := make([]int, len(s)) // 开始搜索的起始位置
	for len(stack) > 0 {
		begin := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[begin] == 0 {
			for end := begin + 1; end <= len(s); end++ {
				if m[s[begin:end]] {
					stack = append(stack, end) // 下一次搜索的起始位置
					if end == len(s) {         //已经搜索到字符串末尾
						return true
					}
				}
			}
			visited[begin] = 1 // 这个位置已经搜索过
		}
	}

	return false
}
