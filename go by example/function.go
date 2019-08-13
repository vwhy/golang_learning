package main

import "fmt"

func main() {
	sum1 := plus(3, 4)
	fmt.Println("3 + 4 =", sum1)
	sum2 := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", sum2)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
