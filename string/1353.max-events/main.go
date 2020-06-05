package _353_max_events

import "sort"

/**
1353. 最多可以参加的会议数目
给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于 endDayi 。

你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。注意，一天只能参加一个会议。

请你返回你可以参加的 最大 会议数目。


示例 1：

输入：events = [[1,2],[2,3],[3,4]]
输出：3
解释：你可以参加所有的三个会议。
安排会议的一种方案如上图。
第 1 天参加第一个会议。
第 2 天参加第二个会议。
第 3 天参加第三个会议。
示例 2：

输入：events= [[1,2],[2,3],[3,4],[1,2]]
输出：4
示例 3：

输入：events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
输出：4
示例 4：

输入：events = [[1,100000]]
输出：1
示例 5：

输入：events = [[1,1],[1,2],[1,3],[1,4],[1,5],[1,6],[1,7]]
输出：7


提示：

1 <= events.length <= 10^5
events[i].length == 2
1 <= events[i][0] <= events[i][1] <= 10^5
*/

//贪心的思想，对于第 i 天，如果有若干的会议都可以在这一天开，
//那么我们肯定是让 endDay小的会议先在这一天开才会使答案最优，因为 endDay 大的会议
//可选择的空间是比 endDay 小的多的，所以在满足条件的会议需要让 endDay 小的先开。

func maxEvents(events [][]int) int {

	eLen := len(events)
	if eLen <= 1 {
		return eLen
	}

	// minDay := events[0][0]
	sort.Slice(events, func(i int, j int) bool {
		// if events[i][0] < minDay {
		// 	minDay = events[i][0]
		// }
		// if events[j][0] < minDay {
		// 	minDay = events[j][0]
		// }
		return events[i][1] < events[j][1]
	})
	// fmt.Println(minDay)
	maxDay := events[eLen-1][1]

	days := make([]bool, maxDay+1)
	maxEvent := 0
	for i := 0; i < eLen; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if !days[j] { //每个会议只占用一天
				days[j] = true
				maxEvent++
				break
			}
		}
	}

	return maxEvent

}
