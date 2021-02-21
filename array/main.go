package main

import "fmt"

var m []*Micro

type Micro struct {
	A int
}

func main() {

	add()

	fmt.Println(len(m))
	fmt.Println(m[0].A)

	a := make([]*Micro, len(m))
	for i := 0; i < len(m); i++ {
		a[0].A = 10
	}

	fmt.Println(len(a))
}

func add() {

	mi := &Micro{A: 5}

	m = append(m, mi)

}
