package main

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 1

	start := head

	for start.Next != nil {
		l++
		start = start.Next
	}

	if n == l && l == 1 {
		return nil
	} else if n == l {
		head = head.Next
		return head
	}

	l = l - n - 1
	ln := 0
	begin := head
	res := begin
	for {
		if ln == l {
			break
		}
		begin = begin.Next
		ln++
	}

	if begin.Next.Next == nil {
		begin.Next = nil
		return res
	} else {
		begin.Next = begin.Next.Next
		return res
	}
	return res
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	zero := &ListNode{
		Next: head,
	}
	left := zero
	right := zero
	for i := 0; i <= n; i++ {
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return zero.Next
}

/**
思路:
1.双指针
2.头部加入空节点
*/
