package main

import "fmt"

func main() {
	if 7 == 7 {
		fmt.Println(7)
	} else {
		fmt.Println("not 7")
	}

	if a := 1; a > 0 {
		fmt.Println("big zero")
	}

	if num := 9; num >= 10 {
		fmt.Println("num > 10")
	} else if num <= 8 {
		fmt.Println("num < 8")
	} else {
		fmt.Println("num = 9")
	}
}
