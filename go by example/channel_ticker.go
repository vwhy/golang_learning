package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for tick := range ticker.C {
			fmt.Println(tick)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("ticker stop")
}
