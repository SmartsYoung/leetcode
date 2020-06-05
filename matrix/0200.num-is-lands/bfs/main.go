/**
200. 岛屿数量  中等
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3
*/
package bfs

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int

func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}

	col = len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]byte, i, j int) {
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = '0'

	for len(queue) != 0 { //	循环搜索，直到队列不存在节点为1的值
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for m := 0; m < 4; m++ {
			tmp_i := i + dx[m]
			tmp_j := j + dy[m]

			if 0 <= tmp_i && tmp_i < row && 0 <= tmp_j && tmp_j < col && grid[tmp_i][tmp_j] == '1' {
				grid[tmp_i][tmp_j] = '0'
				queue = append(queue, tmp_i, tmp_j)
			}
		}
	}
}
