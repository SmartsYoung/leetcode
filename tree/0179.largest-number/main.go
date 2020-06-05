package _179_largest_number

import (
	"sort"
	"strconv"
	"strings"
)

/**
179. 最大数
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

//首先将nums中的元素转换为string
//然后利用自定义sort.Slice() 的比较函数，这里的比较函数为ss[i]+ss[j] >= ss[j]+ss[i]  比较两个数字拼接的字符串谁大
//最后将排完序的字符串打印出来。这里注意要将前面的0去除掉。

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	if o[0] == '0' {
		return "0"
	}
	return o
}
