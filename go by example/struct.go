package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 18})
	fmt.Println(person{age: 19, name: "alice"})
	fmt.Println(person{name: "caroline"})
	fmt.Println(&person{name: "david", age: 20})
	s := person{name: "why", age: 22}
	fmt.Println(s.age)
	sp := &s
	fmt.Println(sp.name)

	sp.age = 99
	fmt.Println(s.age)
}
