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
		swapPairs(head1)
	}
}

func Test(t *testing.T) {
	h := swapPairs(head1)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}

func Test_1(t *testing.T) {
	h := swapPairs1(head2)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}
