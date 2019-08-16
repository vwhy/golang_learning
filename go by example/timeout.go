package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "ready!"
	}()

	select {
	case rea := <-c1:
		fmt.Println("received:", rea)
	case <-time.After(time.Second * 1):
		fmt.Println("time out")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "ready!!"
	}()

	select {
	case rea := <-c2:
		fmt.Println("received:", rea)
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
	}
}
