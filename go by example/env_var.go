package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("foo", "foooooo")
	fmt.Println("foo:", os.Getenv("foo"))
	fmt.Println("bar:", os.Getenv("bar"))
	fmt.Println()
	for _, v := range os.Environ() {
		key := strings.Split(v, "=")
		fmt.Println(key[0])
	}
}
