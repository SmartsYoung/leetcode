package dfs

var dx = [4]int{-1, 0, 1, 0} // 控制遍历方向
var dy = [4]int{0, -1, 0, 1}

func judge(grid [][]byte, x, y int) bool { // 判断当前位置能不能走
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) { // 不能走
		return true
	}
	if grid[x][y] == '0' {
		return true
	}
	return false
}

func dfs(grid [][]byte, x, y int) { // 从当前位置开始遍历连接的土地
	if judge(grid, x, y) { // 遇到边界的时候返回
		return
	}
	grid[x][y] = '0'
	for i := 0; i < 4; i++ { // 遍历当前位置的四个方向，并累计与这四个方向连接的土地面积
		dfs(grid, x+dx[i], y+dy[i])
	}
}

func numIslands(grid [][]byte) int {
	cnt := 0                         // 岛屿数量
	for i := 0; i < len(grid); i++ { // 遍历地图
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' { // 如果遍历到土地，岛屿数量+1，同时用 dfs 将该岛屿连接的所有土地变成水
				dfs(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

/*作者：elliotxx
链接：https://leetcode-cn.com/problems/number-of-islands/solution/0msdfs-de-go-shi-xian-by-elliotxx/ */
