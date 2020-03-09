package dfs

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	arr := []int{5, 4, 2, 6}
	res := permute(arr)

	fmt.Println(res)
}
