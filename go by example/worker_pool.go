package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		fmt.Println("job", j, "finashed")
		result <- j * 2
	}
}
func main() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}

	for i := 1; i <= 9; i++ {
		fmt.Println(<-result)
	}
}
