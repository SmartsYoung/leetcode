package bfs

import "container/list"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	marked := make([][]bool, m) // 标记与边界相连的‘O’们
	for i := range marked {
		marked[i] = make([]bool, n)
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs := func(r, c int) {
		queue := list.New()
		queue.PushBack([]int{r, c})
		marked[r][c] = true
		for queue.Len() > 0 {
			pos := queue.Remove(queue.Front()).([]int)
			for _, d := range dirs {
				nextR, nextC := pos[0]+d[0], pos[1]+d[1]
				if nextR < 0 || nextC < 0 || nextR >= m || nextC >= n ||
					board[nextR][nextC] == 'X' || marked[nextR][nextC] {
					continue
				}
				queue.PushBack([]int{nextR, nextC})
				marked[nextR][nextC] = true
			}
		}
	}
	// 上下边界
	for c := 0; c < n; c++ {
		if board[0][c] == 'O' {
			bfs(0, c)
		}
		if board[m-1][c] == 'O' {
			bfs(m-1, c)
		}
	}
	// 左右边界
	for r := 1; r < m-1; r++ {
		if board[r][0] == 'O' {
			bfs(r, 0)
		}
		if board[r][n-1] == 'O' {
			bfs(r, n-1)
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' && !marked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}
