package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recived:", msg1)
		case msg2 := <-c2:
			fmt.Println("recived:", msg2)
		}
	}
}
