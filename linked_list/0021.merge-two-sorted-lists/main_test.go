package main

import (
	"testing"
)

var head1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	},
}

var head2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	},
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTwoLists(head1, head2)
	}
}

func Test(t *testing.T) {
	h := mergeTwoLists(head1, head2)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}

func Test_1(t *testing.T) {
	h := mergeTwoLists1(head1, head2)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}

func Test_3(t *testing.T) {
	h := mergeTwoLists3(head1, head2)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}
