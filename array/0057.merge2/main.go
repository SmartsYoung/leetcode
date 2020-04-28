/**
57. 插入区间 hard
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
*/
package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
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
			} else if li == si && ri <= ei { // {1, 4}, {2, 5}
				intervals = append(intervals[:j], intervals[j+1:]...)
				intervals[i][1] = ei
				i--
				break
			} else if li == si && ei <= ri { //   {1, 4}, {2, 3}
				intervals = append(intervals[:j], intervals[j+1:]...)
			} else if ri < ei { //    {2, 3},{1, 4}
				intervals = append(intervals[:j], intervals[j+1:]...)
				intervals[i][1] = ei
				i--
				break
			} else if ri >= ei { // {1, 4}, {0, 2}     || {1, 4}, {0, 5}
				intervals = append(intervals[:j], intervals[j+1:]...)
			}
		}

		i++
	}

	return intervals

}
func main() {

	arr8 := [][]int{{0, 0}, {1, 4}, {6, 8}, {9, 11}}
	arr9 := []int{0, 9}
	fmt.Println(insert(arr8, arr9))

}
