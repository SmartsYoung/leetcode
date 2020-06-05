package main

import (
	"fmt"
	"testing"
)

func TestExamRoom_Seat(t *testing.T) {
	N := 10
	obj := Constructor(N)
	param_1 := obj.Seat()
	fmt.Println(param_1)
	param_2 := obj.Seat()
	fmt.Println(param_2)
	param_3 := obj.Seat()
	fmt.Println(param_3)
	param_4 := obj.Seat()
	fmt.Println(param_4)
	obj.Leave(4)
	param_5 := obj.Seat()
	fmt.Println(param_5)
}
