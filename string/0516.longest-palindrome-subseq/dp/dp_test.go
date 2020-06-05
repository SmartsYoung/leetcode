package dp

import (
	"fmt"
	"testing"
)

func Test_long(t *testing.T) {
	r := longestPalindromeSubseq("adacd")
	fmt.Println(r)
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq1("bbbab"))
}
