package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println()
		fmt.Println(<-sigs)
		done <- true
	}()

	fmt.Println("waiting for signal")
	<-done
	fmt.Println("exiting")
}
