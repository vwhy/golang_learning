package main

import "fmt"
import "sync"
import "time"

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("worker", id, "begining")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "done")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
