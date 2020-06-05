package main

import "fmt"

type node struct {
	isStr bool
	next  []*node
}

func NewTree() *node {
	return &node{
		isStr: false,
		next:  make([]*node, 26),
	}
}

// 倒着插入  如 time 插入顺序 为 e - m - i -t   单词 me 的顺序为 e - m ,覆盖 time的部分路径
func insert(root *node, str string) {
	for i := len(str) - 1; i >= 0; i-- {
		if root.next[str[i]-'a'] == nil {
			root.next[str[i]-'a'] = &node{
				isStr: i == 0, // i == 0表示一条新的路径建立
				next:  make([]*node, 26),
			}
		}
		root = root.next[str[i]-'a']
	}
}

// 查找 root 中所有的叶子节点路径
func findLeaf(root *node, length int) int {
	var (
		isLeaf = true
		sum    = 0
	)
	nodes := root.next
	for i := 0; i < 26; i++ {
		if nodes[i] != nil {
			isLeaf = false
			sum += findLeaf(nodes[i], length+1)
		}
	}
	if isLeaf {
		return length + 1
	}
	return sum
}

func minimumLengthEncoding(words []string) int {
	var (
		root = NewTree()
	)
	for i := 0; i < len(words); i++ {
		temp := words[i]
		insert(root, temp)
	}
	return findLeaf(root, 0)
}

func main() {
	words := []string{"time", "me", "bell"}
	words1 := []string{"time", "me"}

	res := minimumLengthEncoding(words)

	fmt.Println(res)
	fmt.Println(minimumLengthEncoding(words1))
	fmt.Println(minimumLengthEncoding(words))
}
