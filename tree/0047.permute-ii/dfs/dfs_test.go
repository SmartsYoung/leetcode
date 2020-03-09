package dfs

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	arr := []int{1, 1, 2}
	res := permuteUnique(arr)

	fmt.Println(res)

}
