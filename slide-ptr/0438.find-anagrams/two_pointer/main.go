package main

/**
438. 找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例 1:

输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:

输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

// 超时 原因在于移动一次就得重新计算win窗口 因为没有用到双指针
func match(s string, wordMap map[byte]int) bool {
	ls := len(s)
	win := make(map[byte]int, 0)

	for i := 0; i < ls; i++ {
		if wordMap[s[i]] == 0 {
			return false
		}
		win[s[i]]++
	}

	for i := 0; i < ls; i++ {
		if wordMap[s[i]] != win[s[i]] {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) []int {
	ls := len(s)
	lp := len(p)

	if s == p {
		return []int{0}
	}
	if ls < lp {
		return []int{}
	}
	wordMap := make(map[byte]int, 0)
	//win := make(map[byte]int, lp)
	for i := 0; i < lp; i++ {
		wordMap[p[i]]++
	}
	var r []int

	right := 0

	for right < ls-lp+1 {
		res := match(s[right:right+lp], wordMap)
		if res == true {
			r = append(r, right)
		}
		right++
	}
	return []int{}
}

// 修改   仍然超时 原因是s过大，match函数执行多次循环，修改办法是使用 26 位字母或数字匹配
func match1(s string, wordMap, win map[byte]int) bool {
	ls := len(s)
	for i := 0; i < ls; i++ {
		if wordMap[s[i]] != win[s[i]] {
			return false
		}
	}

	return true
}

func findAnagrams1(s string, p string) []int {
	ls := len(s)
	lp := len(p)

	if s == p {
		return []int{0}
	}
	if ls < lp {
		return []int{}
	}
	wordMap := make(map[byte]int, 0)
	win := make(map[byte]int, 0)
	for i := 0; i < lp; i++ {
		wordMap[p[i]]++
	}
	var r []int

	left := 0
	right := 0

	for right < ls {

		for right-left+1 > lp {
			win[s[left]]--
			left++
		}
		res := match1(s[right:right+lp], wordMap, win)
		if right-left+1 == len(p) && res == true {
			r = append(r, right)
		}
		right++

	}
	return []int{}
}

func findAnagrams2(s string, p string) []int {
	ls := len(s)
	lp := len(p)

	match := 0
	if s == p {
		return []int{0}
	}
	if ls < lp {
		return []int{}
	}
	wordMap := make(map[byte]int, 0)
	win := make(map[byte]int, 0)
	for i := 0; i < lp; i++ {
		wordMap[p[i]]++
	}
	var r []int

	left := 0
	right := 0

	for right < ls {

		c := s[right]

		if wordMap[c] != 0 {
			win[c]++
			if win[c] == wordMap[c] {
				match++
			}
		}
		right++

		// window 中的字符串已符合 wordsmap 的要求了
		for match == lp {
			if right-left == lp {
				r = append(r, left)
			}

			c2 := s[left]

			if wordMap[c2] != 0 {
				win[c2]--
				if win[c2] < wordMap[c2] {
					match--
				}
				left++
			}
		}
	}
	return r
}

func main() {

}
