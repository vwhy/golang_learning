package main

import (
	"fmt"
	"sort"
)

type bylength []string

func (s bylength) Len() int {
	return len(s)
}
func (s bylength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func (s bylength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fruit := bylength{"peach", "watermelon", "banana"}
	sort.Sort(fruit)
	fmt.Println(fruit)
}
