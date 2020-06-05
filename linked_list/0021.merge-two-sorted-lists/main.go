package main

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	zero := &ListNode{
		Val:  0,
		Next: nil,
	}

	if l1.Val > l2.Val {
		zero.Next = l2
	} else {
		zero.Next = l1
	}

	head := zero.Next

	for {
		if l1.Val > l2.Val {
			zero.Next = l2
			zero = zero.Next
			l2 = l2.Next
		} else {
			zero.Next = l1
			zero = zero.Next
			l1 = l1.Next
		}

		if l1 == nil {
			zero.Next = l2
			break
		}
		if l2 == nil {
			zero.Next = l1
			break
		}
	}
	return head
}

//
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := tmp

	for l1 != nil && l2 != nil {

		if l1.Val <= l2.Val {
			tmp.Next = l1
			tmp = tmp.Next //注意容易漏掉这一步
			if l1.Next == nil {
				tmp.Next = l2
				return head.Next
			}
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next //注意容易漏掉这一步
			if l2.Next == nil {
				tmp.Next = l1
				return head.Next
			}
			l2 = l2.Next
		}

	}

	return head.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	res := tmp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return res.Next
}

/**
思路:
根据大小，插入链表。
*/

// 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}

	return nil
}
