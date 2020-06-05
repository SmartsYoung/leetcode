package slide_window

/**
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。


示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/
// my answer
func findSubstring(s string, words []string) []int {
	var wordMap map[string]int

	if words == nil || s == "" {
		return nil
	}
	ls := len(s)
	lw := len(words)
	l := len(words[0])

	if ls < l*lw {
		return []int{}
	}
	for _, word := range words {
		if _, ok := wordMap[word]; ok {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}

	var r []int
	var m2 map[string]int

	for i := 0; i < ls-l*lw+1; i++ {
		j := i
		w := s[j+lw]
		for {
			if j-i == l*lw {
				r = append(r, j)
				break
			}
			if k, ok := wordMap[string(w)]; ok {
				if k2, ok := m2[string(w)]; ok {
					if k == k2 {
						break
					}
					m2[string(w)] += 1
				} else {
					m2[string(w)] = 1
				}
			} else {
				break
			}
			j += lw
		}
	}

	return []int{}
}

// 参考
func findSubstrings(s string, words []string) []int {
	var result []int
	if len(words) == 0 || len(s) == 0 {
		return result
	}
	wordMap := make(map[string]int, len(words)) // word出现的次数
	var slideMap map[string]int
	for _, word := range words {
		wordMap[word] += 1
	}
	wordLen := len(words[0])          // 单词的长度
	slideSize := wordLen * len(words) // 滑动窗口的大小
	var l, p int
	// 窗口滑动到新的边界
	slideWindow := func(newl int) {
		if newl > p { // 遇到了不是子串的单词, 才会滑动到p之后
			if len(slideMap) > 0 {
				slideMap = make(map[string]int, len(words))
			}
		} else {
			// 因为从旧窗口滑动到了新的窗口，所以要从slideMap中移除被滑出的元素
			for k := l; k < newl; k += wordLen {
				preWord := s[k : k+wordLen]
				slideMap[preWord] -= 1 // 移除元素
			}
		}
		l = newl
	}
	// 滑动窗口到某一个单词之后
	slideWindowToWord := func(word string) {
		k := l
		for ; k < l+slideSize; k += wordLen {
			preWord := s[k : k+wordLen]
			slideMap[preWord] -= 1
			if preWord == word {
				break
			}
		}
		l = k + wordLen
	}
	for i := 0; i < wordLen; i++ { // 遍历[0, wordLen-1]每一种开头的解
		l, p = i, i
		slideMap = make(map[string]int, len(words))
		for l+slideSize <= len(s) {
			if p == l+slideSize { // 找到一个解, p到了滑动窗口的右边界
				result = append(result, l)
				// 滑动一个字节
				slideWindow(l + wordLen)
				continue // 防止p越界
			}
			word := s[p : p+wordLen]
			if wordMap[word] > 0 { //是子串
				// 增加计数
				slideMap[word] += 1
				if slideMap[word] > wordMap[word] { // 次数超出范围
					slideWindowToWord(word) // 滑动窗口到该单词第一次出现的位置后
				}
			} else { // 不是子串
				slideWindow(p + wordLen) // 滑动到一个字节
			}
			p += wordLen
		}
	}
	return result
}

/**
思路
此处需要注意有一些重要条件：

words里的单词长度相同
寻找的子串是words里单词包含所有可能的拼接顺序
根据条件，我们可以每次从s中截取固定长度的子串，并判断在这个子串中是否包含了words里所有单词出现的次数（因为顺序不定，只要出现次数相同就可以构造出对应子串）。
再根据words里的单词长度相同, 我们每次从子串中截取一个单词长度的子串，判断它是否是子串中的单词，这里需要注意匹配到以后应该从这个单词末尾继续匹配，不然会出现单词重叠的情况。

作者：vouv
链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/chuan-lian-suo-you-dan-ci-de-zi-chuan-by-awlsx/
*/
func findSubstring2(s string, words []string) (result []int) {
	n := len(words)
	if n == 0 {
		return
	}
	dict := map[string]int{}
	wn := 0
	for _, w := range words {
		dict[w]++
		wn += len(w)
	}
	ns := len(s)
	for i := 0; i < ns-wn+1; i++ {
		if findIndexes(s[i:i+wn], len(words[0]), dict) {
			result = append(result, i)
		}
	}
	return
}

func findIndexes(s string, wl int, dict map[string]int) bool {
	ns := len(s)
	tmp := map[string]int{}
	for i := 0; i < ns-wl+1; i++ {
		k := s[i : i+wl]
		if dict[k] != 0 {
			tmp[k]++
			// 从末尾开始继续匹配，防止单词重叠
			i = i + wl - 1
		}
	}
	for k, v := range dict {
		if tmp[k] != v {
			return false
		}
	}
	return true
}
