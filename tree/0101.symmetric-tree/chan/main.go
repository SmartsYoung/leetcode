package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodeCount(cur *TreeNode) int {
	if cur == nil {
		return 1
	}
	return getNodeCount(cur.Left) + getNodeCount(cur.Right) + 1
}

func LeftGen(lchan chan int, cur *TreeNode) {
	if cur == nil {
		lchan <- -1
		return
	}
	lchan <- cur.Val // 加入左遍历顺序通道
	LeftGen(lchan, cur.Left)
	LeftGen(lchan, cur.Right)
}

func RightGen(rchan chan int, cur *TreeNode) {
	if cur == nil {
		rchan <- -1
		return
	}
	rchan <- cur.Val // 加入左遍历顺序通道
	RightGen(rchan, cur.Right)
	RightGen(rchan, cur.Left)
}

func isSymmetric(root *TreeNode) bool { // 使用通道 chan 构造迭代器，判断每次结果是否一样
	nodeCount := getNodeCount(root) // 总节点数（加上nil节点）
	lchan := make(chan int)
	rchan := make(chan int)
	go LeftGen(lchan, root) // 开协程生成遍历顺序
	go RightGen(rchan, root)
	for nodeCount > 0 {
		lcur := <-lchan // 从通道中分别取出遍历顺序
		rcur := <-rchan
		if lcur != rcur { // 判断遍历顺序是否相同
			return false
		}
		nodeCount--
	}
	defer close(lchan) // 使用 defer 关键字在函数 return 之后关闭通道
	defer close(rchan)
	return true
}
