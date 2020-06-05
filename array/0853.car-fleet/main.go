package _853_car_fleet

import "sort"

/**
853. 车队
N  辆车沿着一条车道驶向位于 target 英里之外的共同目的地。

每辆车 i 以恒定的速度 speed[i] （英里/小时），从初始位置 position[i] （英里） 沿车道驶向目的地。

一辆车永远不会超过前面的另一辆车，但它可以追上去，并与前车以相同的速度紧接着行驶。

此时，我们会忽略这两辆车之间的距离，也就是说，它们被假定处于相同的位置。

车队 是一些由行驶在相同位置、具有相同速度的车组成的非空集合。注意，一辆车也可以是一个车队。

即便一辆车在目的地才赶上了一个车队，它们仍然会被视作是同一个车队。



会有多少车队到达目的地?



示例：

输入：target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
输出：3
解释：
从 10 和 8 开始的车会组成一个车队，它们在 12 处相遇。
从 0 处开始的车无法追上其它车，所以它自己就是一个车队。
从 5 和 3 开始的车会组成一个车队，它们在 6 处相遇。
请注意，在到达目的地之前没有其它车会遇到这些车队，所以答案是 3。

提示：

0 <= N <= 10 ^ 4
0 < target <= 10 ^ 6
0 < speed[i] <= 10 ^ 6
0 <= position[i] < target
所有车的初始位置各不相同。
*/

/*先算出每辆车辆到达目的地所需的时间,然后按照距目的地的距离排序.
遍历struct[],取出第一个值,表示当前车队到达需要的时间,判断后车到达时间是否比前车小:
是,则表示其会追上前车,二者视为一个车队.
否,则刷新最慢到达时间,此后车视为新车队的队头.
遍历完成,返回最慢到达时间的变化数,则是车队数量.*/

type node struct {
	pos   int
	times float64
}

func carFleet(target int, position []int, speed []int) int {
	node := make([]node, len(position))
	if len(position) == 0 {
		return 0
	}
	for i, _ := range position {
		node[i].pos = position[i]
		node[i].times = float64(target-position[i]) / float64(speed[i])
	}
	sort.Slice(node, func(i, j int) bool {
		return node[i].pos > node[j].pos
	})
	sign := node[0].times
	count := 1
	for _, time := range node {
		if time.times > sign {
			count++
			sign = time.times
		}
	}
	return count
}
