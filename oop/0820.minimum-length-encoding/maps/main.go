package main

import "fmt"

func minimumLengthEncoding(words []string) int {

	m := map[string]bool{}

	for _, v := range words {
		m[v] = true
	}

	for k, _ := range m {
		for i := 1; i < len(k); i++ {
			delete(m, k[i:])
		}
	}

	res := 0
	for k, _ := range m {

		res += len(k) + 1
	}

	return res
}

func main() {
	words := []string{"time", "me", "bell"}

	fmt.Println(minimumLengthEncoding(words))
}
