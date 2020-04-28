package main

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/

func isValidSudoku(board [][]byte) bool {
	row := [9]map[byte]bool{}
	col := [9]map[byte]bool{}
	box := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		row[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		box[i] = map[byte]bool{}
	}

	/**
	遍历数独。
	检查看到每个单元格值是否已经在当前的行 / 列 / 子数独中出现过：
	如果出现重复，返回 false。
	如果没有，则保留此值以进行进一步跟踪。
	返回 true。
	*/

	for rowK, rowV := range board {
		for colK, colV := range rowV {
			if board[rowK][colK] != '.' {
				if row[rowK][colV] == true || col[colK][colV] == true || box[(rowK/3)*3+colK/3][colV] == true {
					return false
				}
				row[rowK][colV] = true
				col[colK][colV] = true
				box[(rowK/3)*3+colK/3][colV] = true
			}
		}
	}
	return true
}
