package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2("abc")
	}
}
