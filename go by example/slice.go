package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[2:5]
	fmt.Println("l1:", l)
	l = c[2:]
	fmt.Println("l2", l)
	l = c[:5]
	fmt.Println("l3", l)
	l = c[:]
	fmt.Println("l4", l)
	t := []string{"a", "b", "c"}
	fmt.Println("dcl", t)

	twos := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twos[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twos[i][j] = i + j
		}
	}
	fmt.Println("2d", twos)
}
