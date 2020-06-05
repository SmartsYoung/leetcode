package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2("abc")
	}
}

func Test_3(t *testing.T) {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring3(s))

}
