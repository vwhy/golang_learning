package main

import "fmt"

func main() {
	fmt.Println(fact(7))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
