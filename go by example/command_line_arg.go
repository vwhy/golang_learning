package main

import (
	"fmt"
	"os"
)

func main() {
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argWithProg, "\n", argWithoutProg, "\n", arg)
}
