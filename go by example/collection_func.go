package main

import "fmt"
import "strings"

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}

func Map(strs []string, f func(string) string) []string {
	strsU := make([]string, len(strs))
	for i, v := range strs {
		strsU[i] = f(v)
	}
	return strsU
}

func Filter(strs []string, f func(string) bool) []string {
	strsf := make([]string, 0)
	for _, v := range strs {
		if f(v) {
			strsf = append(strsf, v)
		}
	}
	return strsf
}

func All(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

func Include(strs []string, str string) bool {
	return Index(strs, str) >= 0
}

func Index(strs []string, str string) int {
	for i, v := range strs {
		if v == str {
			return i
		}
	}
	return -1
}
