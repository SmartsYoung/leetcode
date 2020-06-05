package bfs

import (
	"fmt"
	"testing"
)

func Test_word(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	res := wordBreak(s, wordDict)
	fmt.Println(res)

	s1 := "catsandog"
	w1 := []string{"cats", "dog", "sand", "and", "cat"}

	r := wordBreak(s1, w1)
	fmt.Println(r)
}
