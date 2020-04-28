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

package main

import "fmt"

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int
var parent []int
var count int
var rank []int

func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}

	col = len(grid[0])
	//path compress
	count = 0
	parent = make([]int, row*col)
	rank = make([]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				parent[i*col+j] = i*col + j
				count++
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					nr, nc := i+dx[k], j+dy[k]
					if isValid(grid, nr, nc) && grid[nr][nc] == '1' {
						union(i*col+j, nr*col+nc)
					}
				}
			}
		}
	}
	return count
}
func isValid(grid [][]byte, i, j int) bool {
	return !(i < 0 || i >= row || j < 0 || j >= col)
}

/**
*查找集合i（一个元素是一个集合）的源头（递归实现）。
 如果集合i的父亲是自己，说明自己就是源头，返回自己的标号；
 否则查找集合i的父亲的源头。
**/

func find(i int) int {
	if parent[i] != i {
		parent[i] = find(parent[i])
	}
	return parent[i]
}

func union(x, y int) {
	rootx := find(x)
	rooty := find(y)

	//rank
	if rootx != rooty {
		if rank[rootx] > rank[rooty] {
			parent[rooty] = rootx
		} else if rank[rootx] < rank[rooty] {
			parent[rootx] = rooty
		} else {
			parent[rooty] = rootx
			rank[rootx] += 1
		}
		count -= 1
	}

}

/**
定义一个struct：
包含三个interface：
union(element1, element2 int) //并查集中的“并”
find(input int) int //并查集中的“查”
getCount() //返回“不相交元素集合”个数
包含两个变量：
count int //记录“不相交元素集合”个数
parent []int //记录每个元素的parent，有相同parent的元素视为在同一个“不相交元素集合”

因为go没有构造函数的概念，所以另外写了一个接口：
NewUnionFind(grid [][]byte) *UnionFind
来实现并查集的初始化

总体思路就比较简单了：

生成并查集，并将所有值为'1'的元素放进相应的“不相交元素集合”中
查看一共有多少个“不相交元素集合”，该值即为题目定义的岛屿的数量

作者：todayweather
链接：https://leetcode-cn.com/problems/number-of-islands/solution/goyu-yan-bing-cha-ji-by-todayweather/
*/

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	ret := new(UnionFind)
	ret.count = 0
	m := len(grid)
	n := len(grid[0])
	parent := make([]int, m*n)
	// 并查集初始化，每一个元素的parent都是它自己，每一个元素单独构成一个不相交元素集合
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j //因为 j<n，因此这个值是唯一的
				ret.count++
			}
		}
	}
	ret.parent = parent
	return ret
}

func (uf *UnionFind) union(element1, element2 int) {
	parent1 := uf.find(element1)
	parent2 := uf.find(element2)
	if parent1 != parent2 { //两个节点的父节点不一样时，将查找节点设置为父节点
		uf.parent[parent1] = parent2
		uf.count--
	}
}

//在并查集里查找当前节点的父节点
func (uf *UnionFind) find(input int) int {
	if uf.parent[input] != input {
		uf.parent[input] = uf.find(uf.parent[input]) // 将当前节点的父节点直接设置为父节点
	}
	return uf.parent[input]
}

func (uf *UnionFind) getCount() int {
	return uf.count
}

func numIslands1(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	uf := NewUnionFind(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				cur := i*col + j
				// 只用查找右方和下方的点，左边的点已经在j - 1时处理过了，上边的点已经在i - 1时处理过了
				if i+1 < row && grid[i+1][j] == '1' {
					uf.union(cur, (i+1)*col+j) //将当前节点与下边节点联合
				}
				if j+1 < col && grid[i][j+1] == '1' {
					uf.union(cur, i*col+j+1) //将当前节点与右边节点联合
				}
			}
		}
	}
	return uf.getCount()
}

func main() {

	matrix := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res := numIslands(matrix)
	res1 := numIslands1(matrix)

	fmt.Println(res)
	fmt.Println(res1)

}
