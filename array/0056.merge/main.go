/**
56. 合并区间   Medium
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
//本地测试通过，但LeetCode测试数组{{0, 3}, {4, 5}, {5, 7}}未通过

package main

import (
	"fmt"
	"sort"
)

func merge1(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]

		if c[0] > m[1] {
			merged = append(merged, c) // 后面的区间对均大于当前区间对
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1] //替换merge切片的最后一个区间对的最后一个元素值
		}
	}

	return merged
}

func main() {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	arr1 := [][]int{{1, 4}, {0, 4}}
	arr2 := [][]int{{1, 4}, {0, 2}}
	arr3 := [][]int{{1, 4}, {2, 3}}
	arr4 := [][]int{{1, 4}, {3, 5}}
	arr5 := [][]int{{1, 4}, {0, 2}, {3, 5}}
	arr6 := [][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}}
	arr7 := [][]int{{0, 3}, {4, 5}, {5, 7}}
	fmt.Println(merge(arr))
	fmt.Println(merge(arr1))
	fmt.Println(merge(arr2))
	fmt.Println(merge(arr3))
	fmt.Println(merge(arr4))
	fmt.Println(merge(arr5))
	fmt.Println(merge(arr6))
	fmt.Println(merge(arr7))
}

// 时间复杂度稍高
func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool { //切片数组排序
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	for i < len(intervals) {
		li := intervals[i][0]
		ri := intervals[i][1]
		j := i + 1
		for j < len(intervals) {
			si := intervals[j][0]
			ei := intervals[j][1]

			if si > ri {
				j++
			} else if li == si && ri <= ei {
				intervals = append(intervals[:j], intervals[j+1:]...)
				intervals[i][1] = ei
				i--
				break
			} else if li == si && ei <= ri {
				intervals = append(intervals[:j], intervals[j+1:]...)
			} else if ri < ei {
				intervals = append(intervals[:j], intervals[j+1:]...)
				intervals[i][1] = ei
				i--
				break
			} else if ri >= ei {
				intervals = append(intervals[:j], intervals[j+1:]...)
			}
		}

		i++
	}

	return intervals
}

/*func merge(intervals [][]int) [][]int {

	m := &intervals
	sort.Slice(intervals, func(i, j int) bool { //切片数组排序
		return (*m)[i][0] < (*m)[j][0]
	})

	i := 0
	for i < len(intervals) {
		li := (*m)[i][0]
		ri := (*m)[i][1]
		j := i + 1
		for j < len(intervals) {
			si := (*m)[j][0]
			ei := (*m)[j][1]

			if si > ri {
				j++
			} else if li == si && ri <= ei { // {1, 4}, {2, 5}
				//res = append(res)
				*m = append((*m)[:j], (*m)[j+1:]...)
				intervals[i][1] = ei
			} else if li == si && ei <= ri { //   {1, 4}, {2, 3}
				*m = append((*m)[:j], (*m)[j+1:]...)
				//intervals = append(intervals[:j], intervals[j+1:]...)
			} else if ri < ei { //    {2, 3},{1, 4}
				*m = append((*m)[:j], (*m)[j+1:]...)
				//intervals = append(intervals[:j], intervals[j+1:]...)
				(*m)[i][1] = ei
				i--
				break
			} else if ri >= ei { // {1, 4}, {0, 2}     || {1, 4}, {0, 5}
				*m = append((*m)[:j], (*m)[j+1:]...)
				//intervals = append(intervals[:j], intervals[j+1:]...)
			}
		}

		i++
	}

	return *m
}
*/
