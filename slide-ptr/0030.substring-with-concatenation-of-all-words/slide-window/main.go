package slide_window

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
