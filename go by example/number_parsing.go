package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseFloat("1.234", 64)
	b, _ := strconv.ParseInt("123", 0, 64)
	c, _ := strconv.ParseInt("0x666", 0, 64)
	d, _ := strconv.ParseUint("666", 0, 64)
	e, _ := strconv.Atoi("123")
	_, err := strconv.Atoi("what")
	fmt.Println(a, b, c, d, e, err)
}
