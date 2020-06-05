package main

import "container/list"

/*
855. 考场就座
在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。

示例：

输入：["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
输出：[null,0,9,4,2,null,5]
解释：
ExamRoom(10) -> null
seat() -> 0，没有人在考场里，那么学生坐在 0 号座位上。
seat() -> 9，学生最后坐在 9 号座位上。
seat() -> 4，学生最后坐在 4 号座位上。
seat() -> 2，学生最后坐在 2 号座位上。
leave(4) -> null
seat() -> 5，学生最后坐在 5 号座位上。


提示：

1 <= N <= 10^9
在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。
*/

/**
这题使用链表就可以很简单的解决 时间空间复杂度都可以接受
插入和删除都是 o(n)

首先需要考虑的点有

新的位置需要放在两个已知点之间，如何求出这个点。因为要距离最大化，那么肯定是在中点，如果距离是奇数，那就是中点，如果是偶数，那就直接放在左侧。
如果要插入到头或者尾，需要特殊判断距离
在找到最大距离的同时，需要将对应的插入位置也保存下来
删除的时候只要遍历链表删除即可

作者：resara
链接：https://leetcode-cn.com/problems/exam-room/solution/golang-lian-biao-16ms-on-by-resara/
*/

type ExamRoom struct {
	seat *list.List // 表示坐着的同学的位置
	n    int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seat: list.New(),
		n:    N - 1,
	}
}

func (this *ExamRoom) Seat() int {
	// 还没有人入座，直接将0插入
	if this.seat.Len() == 0 {
		this.seat.PushFront(0)
		return 0
	}
	e := this.seat.Front()
	pre := e.Value.(int)
	max := pre // 头部需要特殊判断
	addVal := 0
	addFront := e
	e = e.Next()
	for ; e != nil; e = e.Next() {
		val := e.Value.(int)
		distant := (val - pre) / 2 // 两点之间的最远距离
		if distant > max {
			max = distant
			addFront = e           // 需要插入的点的后一个元素。方便找到后直接插入
			addVal = pre + distant // 需要插入的位置
		}
		pre = val
	}
	distant := this.n - pre // 尾部特殊判断
	if distant > max {
		this.seat.PushBack(this.n) // 直接插入到链表尾部
		return this.n
	}
	this.seat.InsertBefore(addVal, addFront) // 插入
	return addVal
}

// 遍历知道对应的位置删除
func (this *ExamRoom) Leave(p int) {
	for e := this.seat.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			this.seat.Remove(e)
			return
		}
	}
	return
}
