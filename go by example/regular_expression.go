package main

import "fmt"
import "bytes"
import "regexp"

var p = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	p(r.MatchString("peach"))
	p(r.FindString("peach punch"))
	p(r.FindStringIndex("peach punch"))
	p(r.FindStringSubmatch("peach"))
	p(r.FindStringSubmatchIndex("peach punch"))
	p(r.FindAllString("peach punch pinch", -1))
	p(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	p(r.FindAllString("peach punch pinch", 2))
	p(r.Match([]byte("match")))
	r = regexp.MustCompile("p([a-z]+)ch")
	p(r)
	p(r.ReplaceAllString("a peach", "fruit"))
	p(string(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)))
}
