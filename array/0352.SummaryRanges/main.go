/**
352. 将数据流变为多个不相交区间  hard
给定一个非负整数的数据流输入 a1，a2，…，an，…，将到目前为止看到的数字总结为不相交的区间列表。

例如，假设数据流中的整数为 1，3，7，2，6，…，每次的总结为：

[1, 1]
[1, 1], [3, 3]
[1, 1], [3, 3], [7, 7]
[1, 3], [7, 7]
[1, 3], [6, 7]


进阶：
如果有很多合并，并且与数据流的大小相比，不相交区间的数量很小，该怎么办?
*/

package main

import "fmt"

type Interval struct {
	start int
	end   int
}

type SummaryRanges struct {
	res []Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{
		res: make([]Interval, 0),
	}
}

func (this *SummaryRanges) AddNum(val int) {
	cur := Interval{val, val}
	res := make([]Interval, 0)
	pos := 0

	for _, it := range this.res {
		if cur.end+1 < it.start {
			res = append(res, it)
			continue
		}

		// 当前要插入的在 it 后面，先插入 it
		// pos 记录插入的位置
		if cur.start > it.end+1 {
			res = append(res, it)
			pos++
			continue
		}

		cur.start = mini(cur.start, it.start)
		cur.end = max(cur.end, it.end)
	}

	this.res = insert(res, cur, pos)
}

func (this *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, len(this.res))

	for k, v := range this.res {
		res[k] = make([]int, 2)
		res[k][0] = v.start
		res[k][1] = v.end
	}
	return res
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(a []Interval, c Interval, i int) []Interval {
	return append(a[:i], append([]Interval{c}, a[i:]...)...)
}

func main() {
	var val int
	_, _ = fmt.Scanf("%d", val)
	obj := Constructor()
	obj.AddNum(val)
	param_2 := obj.GetIntervals()
	fmt.Println(param_2)
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */

/**
作者：happyeveryday123
链接：https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/solution/golang-qu-jian-he-bing-wen-ti-shi-jian-fu-za-du-bi/
*/
