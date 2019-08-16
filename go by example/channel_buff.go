package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "first msg"
	message <- "second msg"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
