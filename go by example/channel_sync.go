package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool)
	go worker(done)

	<-done
}
