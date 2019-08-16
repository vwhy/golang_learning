package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	bustryrequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bustryrequests <- i
	}
	close(bustryrequests)
	bustrylimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		bustrylimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			bustrylimiter <- t
		}
	}()
	for req := range bustryrequests {
		<-bustrylimiter
		fmt.Println("request", req, time.Now())
	}
}
