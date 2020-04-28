package main

import "fmt"

/**
54. 螺旋矩阵
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
// 从外部向内部逐层遍历打印矩阵，最外面一圈打印完，里面仍然是一个矩阵

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	m := len(matrix)
	n := len(matrix[0])

	res := make([]int, 0)

	count := (min(m, n) + 1) >> 1 //确保最小值为奇数行时的值

	i := 0
	for i < count {
		for j := i; j < n-i; j++ {
			res = append(res, matrix[i][j])
		}

		for j := i + 1; j < m-i; j++ {
			res = append(res, matrix[j][n-i-1])
		}

		for j := n - i - 2; j >= i && m-1-i != i; j-- { // 从右往左，如果这一层只有1行，那么第一个循环已经将该行打印了，这里就不需要打印了，即 （m-1-i ）!= i
			res = append(res, matrix[m-1-i][j])
		}
		for j := m - i - 2; j >= i+1 && n-1-i != i; j-- { //从下往上，如果这一层只有1列，那么第2个循环已经将该列打印了，这里不需要打印，即(n-1-i) != i
			res = append(res, matrix[j][i])
		}
		i++
	}

	return res
}

func main() {
	m := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	test := [][]int{{2, 5, 8}, {-1, 0, 4}}
	n := spiralOrder(m)
	testRes := spiralOrder(test)

	fmt.Println(n, testRes)

}
