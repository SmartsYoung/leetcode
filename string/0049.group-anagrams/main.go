/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
链接：https://leetcode-cn.com/problems/group-anagrams
*/

package main

import (
	"sort"
	"strconv"
)

func GroupAnagrams(strs []string) [][]string {
	var str [][]string
	cnt := 0
	hash := map[int]int{}

	m := map[byte]int{'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41,
		'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101}

	arr := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		arr[i] = 1
	}

	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			arr[i] = arr[i] * m[strs[i][j]]
		}
		if _, ok := hash[arr[i]]; !ok {
			hash[arr[i]] = cnt // 如果not ok，更新cnt的值
			cnt++
			str = append(str, []string{})
		}
		str[hash[arr[i]]] = append(str[hash[arr[i]]], strs[i])
	}
	return str
}

/*Using characters frequency as key
下面这种解法没有用到排序，
我们用一个大小为26的int数组来统计每个单词中字符出现的次数，
然后将int数组转为一个唯一的字符串，跟字符串数组进行映射，
这样我们就不用给字符串排序了，代码如下：
*/
// 不用给字符串排序
// Using HashMap Using characters frequency as key
// Time Complexity: O(n*k) where k is the max length of string in strs
// Space Complexity: O(n*k)
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string) // 分配字符串切片map
	for _, str := range strs {
		key := getKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 用一个大小为26的int数组来统计每个单词中字符出现的次数，
// 然后将int数组转为一个唯一的字符串，跟字符串数组进行映射，
func getKey(s string) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	key := ""
	for _, n := range freq {
		key += strconv.Itoa(n) + "/"
	}
	return key
}

/*
字符串排序
这道题让我们群组给定字符串集中所有的错位词，
所谓的错位词就是两个字符串中字母出现的次数都一样，只是位置不同，
比如abc，bac, cba等它们就互为错位词，
那么我们如何判断两者是否是错位词呢，
我们发现如果把错位词的字符顺序重新排列，那么会得到相同的结果，
所以重新排序是判断是否互为错位词的方法，由于错位词重新排序后都会得到相同的字符串，
我们以此作为key，将所有错位词都保存到字符串数组中，建立key和字符串数组之间的映射，
最后再存入结果res中即可，擦巾代码如下：
*/
// 字符串排序 Using HashMap Using sorted string as key
// Time Complexity: O(n*klogk) where k is the max length of string in strs
// Space Complexity: O(n*k)
func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	record := make(map[string][]string)
	for _, str := range strs {
		s := []rune(str)            // 把错位词的字符顺序重新排列，那么会得到相同的结果
		sort.Sort(sortRunes(s))     // 以此作为key，将所有错位词都保存到字符串数组中
		val, _ := record[string(s)] // 建立key和字符串数组之间的映射
		val = append(val, str)
		record[string(s)] = val
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

/*原作者：custergo
链接：https://leetcode-cn.com/problems/group-anagrams/solution/custerxue-xi-bi-ji-zi-fu-chuan-pai-xu-hashmap-by-c/
*/
