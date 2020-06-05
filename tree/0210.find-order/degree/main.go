/**
210. 课程表 II
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1:

输入: 2, [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2:

输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
说明:

输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
提示:

这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。
*/
package degree

/*
(1)初始化一个大小为numCourses的数组indegree；
（2）将图中个节点的入度保存到数组中，将数组中第一个入度为0的节点找出，
并将该节点在数组中的值设为-1，将数组中所有以该顶点为入度的顶点入度减一，将其push到vector中；
（3）重复（2）numCourses次，若期间在indegree数组中没有找到入度为0的顶点，则返回空;

*/

func findOrder(numCourses int, prerequisites [][]int) []int {

	degree := make(map[int][]int, numCourses) // 每个先导课程后面的课程列表

	outDegree := make([]int, numCourses) //存储每门课的入度

	for i := 0; i < len(prerequisites); i++ {
		degree[prerequisites[i][1]] = append(degree[prerequisites[i][1]], prerequisites[i][0])
		outDegree[prerequisites[i][0]]++
	}
	helper := []int{} //模拟辅助queue

	// 度为 0 的课程
	for i, v := range outDegree {
		if v == 0 {
			helper = append(helper, i) // v = 0时map中会有 i 值？
		}
	}

	for i := 0; i < len(helper); i++ { // 每一个度为 0 的课程
		for _, v := range degree[helper[i]] { // helper[i]为先导课程
			outDegree[v]--         // 课程的度减一
			if outDegree[v] == 0 { // 当课程的度减到 0 时， 说明该课程的先导课程已经全部学习完
				helper = append(helper, v)
			}
		}
	}
	if numCourses > len(helper) {
		return nil
	}

	return helper[0:numCourses]
}
