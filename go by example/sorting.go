package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("sorted strs:", strs)

	ints := []int{3, 5, 4}
	sort.Ints(ints)
	fmt.Println("sorted ints:", ints)
	fmt.Println("sorted?", sort.IntsAreSorted(ints))
}
