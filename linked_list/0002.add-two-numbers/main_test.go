package main

import (
	"fmt"
	"testing"
)

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/

func TestList(t *testing.T) {
	var l1 *ListNode = new(ListNode)
	var l2 *ListNode = new(ListNode)
	var tmp1 *ListNode = l1
	var tmp2 *ListNode = l2
	l1.Val = 9
	l2.Val = 4

	for i := 0; i < 10; i++ {
		tmp1.Next = new(ListNode)
		tmp1 = tmp1.Next
		tmp1.Val = i

		tmp2.Next = new(ListNode)
		tmp2 = tmp2.Next
		tmp2.Val = i
	}

	sum := addTwoNumbers2(l1, l2)
	fmt.Print("sum = ")
	for sum != nil {
		fmt.Printf("[%d]", sum.Val)
		sum = sum.Next
	}
	fmt.Println("")
}

func Benchmark(b *testing.B) {
	l1 := fake([]int{5, 6, 7})
	l2 := fake([]int{5, 5, 5})
	for i := 0; i < b.N; i++ {
		addTwoNumbers2(l1, l2)
	}
}

func fake(list []int) *ListNode {
	l := &ListNode{}
	o := l
	for k, v := range list {
		o.Val = v
		if len(list) > k+1 {
			o.Next = &ListNode{}
			o = o.Next
		}
	}
	return l
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	next := &result
	num := 0
	for l1 != nil || l2 != nil || num > 0 {
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		*next = &ListNode{
			Val:  num % 10,
			Next: nil,
		}
		num = num / 10
		next = &((*next).Next)
	}
	return result
}
