/**
820. 单词的压缩编码  medium
给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。

例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。

对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。

那么成功对给定单词列表进行编码的最小字符串长度是多少呢？



示例：

输入: words = ["time", "me", "bell"]
输出: 10
说明: S = "time#bell#" ， indexes = [0, 2, 5] 。


提示：

1 <= words.length <= 2000
1 <= words[i].length <= 7
每个单词都是小写字母 。
*/
package main

import (
	"fmt"
	"sort"
)

// 为了能够使用自定义函数来排序，我们需要一个
// 对应的排序类型，比如这里我们为内置的字符串
// 数组定义了一个别名ByLength
type ByLength []string

func (p ByLength) Len() int {
	return len(p)
}
func (p ByLength) Less(i, j int) bool {
	return len(p[i]) > len(p[j])
}
func (p ByLength) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func contain(str, substr string) bool {
	if len(substr) > len(str) {
		return false
	}
	str2 := str[len(str)-len(substr):]
	if substr == str2 {
		return true
	}
	return false
}

func minimumLengthEncoding(words []string) int {

	// 可以替换上面的排序函数
	/*sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})*/

	sort.Sort(ByLength(words)) // 实现的是别名ByLength类型的方法，因此通过ByLength调用

	length := len(words)
	res := 0

	hasMap := make([]int, length)

	if length == 1 {
		return len(words[0]) + 1
	}

	start := words[0]
	substr := words[1]

	for i := 0; i < length && hasMap[i] == 0; i++ {
		start = words[i]
		for j := i + 1; j < length && hasMap[j] == 0; j++ {
			substr = words[j]
			if contain(start, substr) == true {
				hasMap[j] = -1 // 标记后面被包含的字符串
				hasMap[i] += 1 //  标记包含后面字符串的起始字符串
			}
		}
	}

	for i, v := range hasMap {
		if v >= 0 {
			res += len(words[i])
			res += 1
		}
	}

	return res
}

func main() {
	words := []string{"time", "me", "bell"}
	words1 := []string{"time", "me"}

	res := minimumLengthEncoding(words)

	fmt.Println(res)
	fmt.Println(minimumLengthEncoding(words1))
	fmt.Println(minimumLengthEncoding1(words))
}

//
func minimumLengthEncoding1(words []string) int {
	m := make(map[string]int, 0)
	// 去重
	for _, w := range words {
		if _, ok := m[w]; !ok {
			m[w] = 1
		}
	}
	// 去子串
	for _, w := range words {
		for i := 1; i < len(w); i++ { //已经去重了，所以可以从子串的第二个位置开始
			delete(m, w[i:]) //若map中有m的子串就从map中删除，没有时不会出错，不会引发panic
		}
	}
	// 计算结果
	ans := 0
	for k, _ := range m {
		ans += len(k) + 1
	}

	return ans
}

/*
作者：Cluas
链接：https://leetcode-cn.com/problems/short-encoding-of-words/solution/qu-zhong-qu-zi-chuan-ji-suan-jie-guo-by-cluas/
*/
