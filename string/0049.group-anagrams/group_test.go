package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestGroup(t *testing.T) {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams3(str)
	fmt.Println(res)
}

//字母异位词分组
func groupAnagrams3(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}
	var helpMap = make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		var curStrArr = strings.Split(strs[i], "")
		sort.Strings(curStrArr)
		var curStrIndex = strings.Join(curStrArr, "")
		helpMap[curStrIndex] = append(helpMap[curStrIndex], strs[i])
	}
	for _, value := range helpMap {
		res = append(res, value)
	}
	return res
}

func groupAnagrams0(strs []string) [][]string {
	cnt := 0                                                                                                             // 有几个字母异位词
	hash := map[int]int{}                                                                                                // 记录每种异位词对应的下标数组
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103} // 用26个质数代表26个字母
	ans := [][]string{}
	for _, s := range strs {
		t := 1
		for _, x := range s { // 计算异位词对应的值
			t *= prime[x-'a']
		}
		if _, ok := hash[t]; !ok { // 这类异位词没出现过，记录新下标
			hash[t] = cnt
			cnt++
			ans = append(ans, []string{}) // 初始化新下标的数组
		}
		ans[hash[t]] = append(ans[hash[t]], s) // 将当前字符串加入该类异位词数组中
	}
	return ans
}
