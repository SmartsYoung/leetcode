package dfs

/**
130. 被围绕的区域
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/

func dfs(board [][]byte, i, j int) {
	// 边界判断
	// board[i][j] == '#': 说明已经搜索过
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	// 和边界联通的O，先替换成#， 后续换回O
	board[i][j] = '#'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	// 先搜索并标记
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 从边缘开始搜索, 判断是否在边上，只从边上开始，如果边上都是X，那么里面的所有O都换成X就行
			// 如果边上有，那么之后就会被替换成#暂存，之后会被恢复成O，意味着不换成X
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
			if isEdge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	// 再替换，这里的O都应该是被包围的
	// # 是和边界联通的，即不被包围的
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
