package problem0127

func ladderLength(beginWord string, endWord string, words []string) int {
	// 把 words 存入字典
	// 可以利用快速地添加，删除和查找单词
	dict := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}
	// 删除 dict 中的 beginWord
	delete(dict, beginWord)

	// queue 用于存放被 trans 到的 word
	queue := make([]string, 0, len(words))

	var trans func(string) bool
	// trans 用来把 words 中 word 能够 trans 到的单词，放入 queue
	// 如果 trans 过程中，遇到了 endWord，则返回 true
	trans = func(word string) bool {
		bytes := []byte(word)
		for i := 0; i < len(bytes); i++ {
			diffLetter := bytes[i]
			// 找出仅在索引 i 上不同的单词
			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				if b == diffLetter {
					continue
				}

				bytes[i] = b
				// 此时 bytes 为 word 所 trans 的单词

				// 令 temp := string(bytes)，会增加 70% 的时间

				if dict[string(bytes)] {
					// words 中存在 string(bytes)
					if string(bytes) == endWord {
						// trans 到了 endWord
						// 提前结束
						return true
					}

					// 把 string(bytes) 放入 queue 的尾部
					queue = append(queue, string(bytes))
					delete(dict, string(bytes))
				}
			}
			bytes[i] = diffLetter
		}

		return false
	}

	queue = append(queue, beginWord)
	dist := 1
	for len(queue) > 0 {
		qLen := len(queue)

		// 这个 for 循环，是按照每个 word 的 dist 值，来切分 queue 的
		for i := 0; i < qLen; i++ {
			// word 出列
			word := queue[0]
			queue = queue[1:]

			if trans(word) {
				// word 能够 trans 到 endWord
				// 提前结束
				return dist + 1
			}
		}

		dist++
	}

	return 0
}

// BFS Time Complexity: O(n*26^l), l = len(word), n=|wordList| Space Complexity: O(n)
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool) // 把word存入字典
	for _, word := range wordList {
		dict[word] = true // 可以利用字典快速添加、删除和查找单词
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	// queue := []string{beginWord} 定义辅助队列
	var queue []string
	queue = append(queue, beginWord)

	l := len(beginWord)
	steps := 0

	for len(queue) > 0 {
		steps++
		size := len(queue)          // 每一层的大小
		for i := size; i > 0; i-- { // 当前层级节点
			s := queue[0] // 原始单词
			queue = queue[1:]
			chs := []rune(s)
			for i := 0; i < l; i++ { // 对单词的每一位进行修改
				ch := chs[i]                  // 对当前单词的一位
				for c := 'a'; c <= 'z'; c++ { // 尝试从a-z
					if c == ch { // 跳过本身比如hot修改为hot
						continue
					}
					chs[i] = c
					t := string(chs)
					if t == endWord { // 找到结果
						return steps + 1
					}
					if _, ok := dict[t]; !ok { // 不在dict中，跳过
						continue
					}
					delete(dict, t)          // 从字典中删除该单词，因为已经访问过，若重复访问路径一定不是最短的
					queue = append(queue, t) // 将新的单词添加到队列
				}
				chs[i] = ch // 单词的第i位复位，再进行下面的操作
			}
		}
	}
	return 0
}
