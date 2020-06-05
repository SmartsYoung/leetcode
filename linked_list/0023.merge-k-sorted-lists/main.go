package main

import (
	"container/heap"
)

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义类型为 ListNode 引用类型
type IntHeap []*ListNode

// Len方法返回集合中的元素个数
func (heap IntHeap) Len() int {
	return len(heap)
}

// Less方法报告索引i的元素是否比索引j的元素小
func (heap IntHeap) Less(i, j int) bool {
	return heap[i].Val < heap[j].Val
}

// Swap方法交换索引i和j的两个元素
func (heap IntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

// 向堆h中插入元素x，并保持堆的约束性。复杂度O(log(n))，其中n等于h.Len()。
func (heap *IntHeap) Push(x interface{}) {
	*heap = append(*heap, x.(*ListNode))
}

// 从末尾删除元素
func (heap *IntHeap) Pop() interface{} {
	var tmp IntHeap = *heap
	var n = len(tmp)
	var x = tmp[n-1]
	*heap = tmp[:n-1]
	return x
}

func mergeKLists1(lists []*ListNode) *ListNode {

	// 特判 链表为空或长度为0，直接返回空指针
	if lists == nil || len(lists) == 0 {
		return nil
	}

	var h *IntHeap = &IntHeap{}
	heap.Init(h)
	for _, list := range lists {

		if list == nil {
			continue
		}

		// 向堆中插入元素
		heap.Push(h, list)
	}

	// 表头
	var node *ListNode = &ListNode{}
	// 操作表
	var tmp *ListNode = node

	for h.Len() > 0 {
		// 从堆中取出最小值并赋予操作表
		tmp.Next = heap.Pop(h).(*ListNode)
		tmp = tmp.Next
		// 如果链表中的节点非空，那么再增加到堆中
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
	}
	// 带头表去头
	return node.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	for l > 1 {
		for i := 0; i < l/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[l-i-1])
		}
		l = (l + 1) / 2
	}
	return lists[0]
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
mergeTwoLists 基础上分治
*/
