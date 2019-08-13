package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 5
	m["k2"] = 6
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("k1:", v1)

	fmt.Println("len", len(m))
	delete(m, "k2")
	fmt.Println("m now", m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"key0": 1, "key1": 2}
	fmt.Println(n)
}
