package main

import "fmt"

func main() {
	v1, v2 := validval()
	fmt.Println("validval:", v1, v2)
	_, v3 := validval()
	fmt.Println("v2:", v3)
}

func validval() (int, int) {
	return 3, 7
}
