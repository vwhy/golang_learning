package main

import "fmt"

func main() {
	for {
		fmt.Println("666")
		break
	}

	for i := 5; i > 2; i-- {
		fmt.Println(i)
	}

	for m := 1; m <= 5; m++ {
		fmt.Println(m)
	}
}
