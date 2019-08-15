package main

import "fmt"

func main() {
	i := 1
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr", i)
	fmt.Println(&i)
}

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}
