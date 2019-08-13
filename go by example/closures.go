package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	nextInt()
	fmt.Println(nextInt())

	newnextInt := intSeq()
	fmt.Println(newnextInt())
	fmt.Println(nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
