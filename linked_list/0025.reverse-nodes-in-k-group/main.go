/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5



说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
通过次数43,258提交次数75,914

链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归反转以 head 为起点的 n 个节点，返回新的头结点
func reverseK(head *ListNode, k int) *ListNode {
	successor := &ListNode{}
	if k == 1 {
		// 记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// 以 head.next 为起点，需要反转后面的 n - 1 个节点
	last := reverseK(head.Next, k-1)
	head.Next.Next = head

	// 让反转之后的 head 节点和后面的节点连起来 只有 k==1时才和后面节点相连，其它情况下在循环开头又重新被赋值为nil
	head.Next = successor

	return last
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	l := 1
	start := head
	for nil != start.Next {
		start = start.Next
		l++
	}
	if l < k {
		return head
	}
	ti := l / k

	tmp := &ListNode{}
	res := head
	tmp = res
	//nStart = head                //原当前k链表的起点， 新链表中当前k链表的端点
	//res = reverseK(tmp.Next, k)  // 当前k链表反转之后返回当前k链表的起点

	for i := 0; i < ti; i++ {
		//	nStart = nextKNode(tmp.Next,k)  //原下一个k链表的起点， 新链表中下一个k链表的端点
		res = reverseK(res, k)
	}
	return tmp
}

// 非递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	empty := &ListNode{}
	empty.Next = head
	prev := empty
	for head != nil {
		for i := 0; i < k-1 && head != nil; i++ {
			head = head.Next
		}
		if head == nil {
			break
		}
		first := prev.Next
		next := head.Next
		head.Next = nil
		prev.Next = reverse(first)
		first.Next = next
		head = next
		prev = first
	}

	return empty.Next
}

func reverse(first *ListNode) *ListNode {
	if first == nil || first.Next == nil {
		return first
	}
	l := first
	r := first.Next
	for r != nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	first.Next = nil
	return l
}

/**

 */
