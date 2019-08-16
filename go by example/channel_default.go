package main

import "fmt"

func main() {
	signal := make(chan bool)
	message := make(chan string)

	select {
	case msg := <-message:
		fmt.Println("received", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"

	select {
	case message <- msg:
		fmt.Println("message send")
	default:
		fmt.Println("no message send")
	}

	select {
	case sig := <-signal:
		fmt.Println("signal received", sig)
	case msg := <-message:
		fmt.Println("message received", msg)
	default:
		fmt.Println("no activity")
	}
}
