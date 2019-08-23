package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	now := time.Now()
	p(time.Now())

	then := time.Date(
		2009, 11, 11, 11, 11, 11, 1111, time.UTC)
	p(then)
	p(then.Year())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Sub(now))
	p(then.Sub(now).Hours())
	p(then.Add(time.Second))
}
