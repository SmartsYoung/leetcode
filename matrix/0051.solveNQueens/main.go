package _051_solveNQueens

func solveNQueens(n int) [][]string {
	board := []string{}
	ans := [][]string{}
	initBoard(&board, n)
	backTrack(0, n, &board, &ans)
	return ans

}

func initBoard(board *[]string, n int) {
	var s string
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < n; i++ {
		*board = append(*board, s)
	}
}

func backTrack(row int, n int, board *[]string, ans *[][]string) {
	if row == len(*board) {
		tmp := make([]string, len(*board))
		copy(tmp, *board)
		*ans = append(*ans, tmp)
	} else {
		for col := 0; col < n; col++ {
			if !isPlace(board, row, col, n) {
				continue
			}
			(*board)[row] = strRex((*board)[row], col, 'Q')
			backTrack(row+1, n, board, ans)
			(*board)[row] = strRex((*board)[row], col, '.')
		}
	}
}

func strRex(s string, i int, b byte) string {
	bytes := []byte(s)
	bytes[i] = b
	s = string(bytes)
	return s
}

func isPlace(board *[]string, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if (*board)[i][col] == byte('Q') {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*board)[i][j] == byte('Q') {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*board)[i][j] == byte('Q') {
			return false
		}
	}
	return true
}

/* 转载
作者：ru-shi-shuo-2
链接：https://leetcode-cn.com/problems/n-queens/solution/golang-hui-su-suan-fa-by-ru-shi-shuo-2/
*/
