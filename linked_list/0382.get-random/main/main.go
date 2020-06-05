package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
382. 链表随机节点
给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:

// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用 time.Now().UnixNano())获得一个带纳秒的时间戳，形成一个新源保证每次获取的随机数都是随机的
	solution := Solution{
		head: head,
		r:    r,
	}
	return solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	cur := this.head.Next
	val := this.head.Val

	// 从第二个节点开始替换
	i := 2
	for cur != nil {
		// 第 i 个节点以 1/i 的概率替换当前值    r.Intn(i) 获取的随机数为[0, i)  this.r.Intn(i)+1 以 1/i 的概率获取[1, i]中的随机数
		if this.r.Intn(i)+1 == i {
			val = cur.Val
		}
		i++
		cur = cur.Next
	}
	return val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	obj := Constructor(head)
	param_1 := obj.GetRandom()
	fmt.Println(param_1)
}

var head = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
}
