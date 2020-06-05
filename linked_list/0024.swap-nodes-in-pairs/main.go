package main

/**
24. 两两交换链表中的节点        medium
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3. */

type ListNode struct {
	Val  int
	Next *ListNode
}

// wrong
func swapPairs1(head *ListNode) *ListNode {

	empty := &ListNode{}
	empty.Next = head
	prev := empty

	for head != nil && head.Next != nil {

		one := head
		two := head.Next

		two.Next = one
		one.Next = head.Next.Next // 这里会无限循环，(画图）相当于one.Next = one  two.Next = one 一定不能写在最前面
		prev.Next = two

		prev = one
		head = head.Next.Next

	}
	return empty.Next
}

func swapPairs(head *ListNode) *ListNode {
	empty := &ListNode{}
	empty.Next = head
	prev := empty

	for head != nil && head.Next != nil {
		one := head
		two := head.Next

		prev.Next = two
		one.Next = two.Next
		two.Next = one

		prev = one
		head = one.Next
	}
	return empty.Next
}

/**
思路:三指针 prev,one,two
*/
