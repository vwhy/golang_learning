package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, "is", v)
	}
	for k := range m {
		fmt.Println("key:", k)
	}
	for i, a := range "go" {
		fmt.Println(i, ":", a)
	}
}
