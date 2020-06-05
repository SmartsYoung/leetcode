package dfs

import (
	"fmt"
	"testing"
)

func Test_word(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	res := wordBreak(s, wordDict)
	fmt.Println(res)
}
