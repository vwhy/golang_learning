package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readops uint64
	var writeops uint64

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	finalreadops := atomic.LoadUint64(&readops)
	finalwriteops := atomic.LoadUint64(&writeops)
	fmt.Print("readops: ", finalreadops, "\n", "writeops: ", finalwriteops, "\n")
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
